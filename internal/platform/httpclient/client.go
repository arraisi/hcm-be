package httpclient

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net"
	"net/http"
	"time"
)

// Client wraps a reusable HTTP client with sensible defaults.
// It implements the HTTPClient interface.
type Client struct {
	httpClient *http.Client
	headers    map[string]string
	retries    int
	backoff    time.Duration
	timeout    time.Duration
}

// New creates a new reusable HTTP client with given options.
func New(opt Options) *Client {
	if opt.Timeout == 0 {
		opt.Timeout = 10 * time.Second
	}
	if opt.Retries == 0 {
		opt.Retries = 2
	}
	if opt.Backoff == 0 {
		opt.Backoff = 300 * time.Millisecond
	}
	httpCli := &http.Client{
		Transport: &http.Transport{
			MaxIdleConns:        100,
			MaxIdleConnsPerHost: 10,
			IdleConnTimeout:     90 * time.Second,
			DialContext: (&net.Dialer{
				Timeout:   5 * time.Second,
				KeepAlive: 30 * time.Second,
			}).DialContext,
			TLSHandshakeTimeout: 5 * time.Second,
		},
	}
	return &Client{
		httpClient: httpCli,
		headers:    opt.Headers,
		retries:    opt.Retries,
		backoff:    opt.Backoff,
		timeout:    opt.Timeout,
	}
}

// DoJSON sends a JSON HTTP request and decodes the JSON response into respOut.
// It handles retries on transient errors and 5xx.
func (c *Client) DoJSON(ctx context.Context, method, url string, reqBody any, respOut any) error {
	var bodyBytes []byte
	if reqBody != nil {
		b, err := json.Marshal(reqBody)
		if err != nil {
			return fmt.Errorf("marshal request: %w", err)
		}
		bodyBytes = b
	}

	shouldRetryStatus := func(code int) bool {
		return code == http.StatusRequestTimeout || // 408
			code == http.StatusTooManyRequests || // 429
			(code >= 500 && code <= 599) // 5xx
	}

	var lastErr error

	for attempt := 0; attempt <= c.retries; attempt++ {
		if attempt > 0 {
			time.Sleep(time.Duration(attempt) * c.backoff)
		}

		// new context per attempt; do NOT defer cancel in a loop
		attemptCtx, cancel := context.WithTimeout(ctx, c.timeout)

		req, err := http.NewRequestWithContext(attemptCtx, method, url, bytes.NewReader(bodyBytes))
		if err != nil {
			cancel()
			return fmt.Errorf("build request: %w", err)
		}
		req.Header.Set("Content-Type", "application/json")
		for k, v := range c.headers {
			req.Header.Set(k, v)
		}

		res, err := c.httpClient.Do(req)
		if err != nil {
			cancel()
			if isTransient(err) {
				lastErr = err
				continue
			}
			return fmt.Errorf("http call: %w", err)
		}

		// Always close body in this iteration; no defer in loop
		respBytes, readErr := io.ReadAll(io.LimitReader(res.Body, 1<<20)) // 1MB cap
		_ = res.Body.Close()
		cancel()

		if readErr != nil {
			// reading body failed; consider retry if transient I/O
			if isTransient(readErr) {
				lastErr = readErr
				continue
			}
			return fmt.Errorf("read body: %w", readErr)
		}

		if res.StatusCode < 200 || res.StatusCode >= 300 {
			httpErr := &HTTPError{StatusCode: res.StatusCode, Body: string(respBytes)}
			if shouldRetryStatus(res.StatusCode) {
				lastErr = httpErr
				continue
			}
			return httpErr
		}

		if respOut != nil {
			if err := json.Unmarshal(respBytes, respOut); err != nil {
				return fmt.Errorf("decode response: %w", err)
			}
		}
		return nil
	}

	return fmt.Errorf("all retries failed: %w", lastErr)
}

func isTransient(err error) bool {
	var nerr net.Error
	if errors.As(err, &nerr) && nerr.Timeout() {
		return true
	}
	return false
}
