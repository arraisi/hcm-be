package utils

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

	"github.com/arraisi/hcm-be/internal/platform/httpclient"
)

// HttpUtil provides a flexible HTTP client interface with token support
type HttpUtil interface {
	Post(ctx context.Context, url string, body any, token string, headerValue ...map[string]string) (response []byte, err error)
	Put(ctx context.Context, url string, body any, token string, headerValue ...map[string]string) (response []byte, err error)
	Get(ctx context.Context, url string, token string, headerValue ...map[string]string) (response []byte, err error)
	Delete(ctx context.Context, url string, token string, headerValue ...map[string]string) (response []byte, err error)
}

// HttpUtilClient implements HttpUtil interface
type HttpUtilClient struct {
	httpClient *http.Client
	retries    int
	backoff    time.Duration
	timeout    time.Duration
}

// NewHttpUtil creates a new HttpUtil client with given options
func NewHttpUtil(opt httpclient.Options) HttpUtil {
	if opt.Timeout == 0 {
		opt.Timeout = 10 * time.Second
	}
	if opt.Retries == 0 {
		opt.Retries = 2
	}
	if opt.Backoff == 0 {
		opt.Backoff = 300 * time.Millisecond
	}

	// Create transport with reasonable defaults
	transport := &http.Transport{
		MaxIdleConns:        100,
		MaxIdleConnsPerHost: 10,
		IdleConnTimeout:     90 * time.Second,
		DialContext: (&net.Dialer{
			Timeout:   5 * time.Second,
			KeepAlive: 30 * time.Second,
		}).DialContext,
		TLSHandshakeTimeout:   5 * time.Second,
		ResponseHeaderTimeout: opt.Timeout, // Use configured timeout for response headers
	}

	httpCli := &http.Client{
		Transport: transport,
		Timeout:   opt.Timeout, // Set overall request timeout from options
	}

	return &HttpUtilClient{
		httpClient: httpCli,
		retries:    opt.Retries,
		backoff:    opt.Backoff,
		timeout:    opt.Timeout,
	}
}

// Post sends a POST request with optional token and headers
func (c *HttpUtilClient) Post(ctx context.Context, url string, body any, token string, headerValue ...map[string]string) (response []byte, err error) {
	return c.doRequest(ctx, http.MethodPost, url, body, token, headerValue...)
}

// Put sends a PUT request with optional token and headers
func (c *HttpUtilClient) Put(ctx context.Context, url string, body any, token string, headerValue ...map[string]string) (response []byte, err error) {
	return c.doRequest(ctx, http.MethodPut, url, body, token, headerValue...)
}

// Get sends a GET request with optional token and headers
func (c *HttpUtilClient) Get(ctx context.Context, url string, token string, headerValue ...map[string]string) (response []byte, err error) {
	return c.doRequest(ctx, http.MethodGet, url, nil, token, headerValue...)
}

// Delete sends a DELETE request with optional token and headers
func (c *HttpUtilClient) Delete(ctx context.Context, url string, token string, headerValue ...map[string]string) (response []byte, err error) {
	return c.doRequest(ctx, http.MethodDelete, url, nil, token, headerValue...)
}

// doRequest is the internal method that handles all HTTP requests
func (c *HttpUtilClient) doRequest(ctx context.Context, method, url string, body any, token string, headerValue ...map[string]string) ([]byte, error) {
	var bodyBytes []byte
	if body != nil {
		b, err := json.Marshal(body)
		if err != nil {
			return nil, fmt.Errorf("marshal request: %w", err)
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

		// new context per attempt
		attemptCtx, cancel := context.WithTimeout(ctx, c.timeout)

		req, err := http.NewRequestWithContext(attemptCtx, method, url, bytes.NewReader(bodyBytes))
		if err != nil {
			cancel()
			return nil, fmt.Errorf("build request: %w", err)
		}

		// Set default headers
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Accept", "application/json")

		// Set authorization token if provided
		if token != "" {
			req.Header.Set("Authorization", "Bearer "+token)
		}

		// Set additional headers if provided
		for _, headers := range headerValue {
			for k, v := range headers {
				req.Header.Set(k, v)
			}
		}

		res, err := c.httpClient.Do(req)
		if err != nil {
			cancel()
			if isTransientError(err) {
				lastErr = err
				continue
			}
			return nil, fmt.Errorf("http call: %w", err)
		}

		// Always close body in this iteration
		respBytes, readErr := io.ReadAll(io.LimitReader(res.Body, 1<<20)) // 1MB cap
		_ = res.Body.Close()
		cancel()

		if readErr != nil {
			// reading body failed; consider retry if transient I/O
			if isTransientError(readErr) {
				lastErr = readErr
				continue
			}
			return nil, fmt.Errorf("read body: %w", readErr)
		}

		if res.StatusCode < 200 || res.StatusCode >= 300 {
			httpErr := &httpclient.HTTPError{StatusCode: res.StatusCode, Body: string(respBytes)}
			if shouldRetryStatus(res.StatusCode) {
				lastErr = httpErr
				continue
			}
			return respBytes, httpErr
		}

		return respBytes, nil
	}

	return nil, fmt.Errorf("all retries failed: %w", lastErr)
}

// isTransientError checks if an error is transient and should be retried
func isTransientError(err error) bool {
	var nerr net.Error
	if errors.As(err, &nerr) && nerr.Timeout() {
		return true
	}
	return false
}
