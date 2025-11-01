package httpclient

import "time"

type Options struct {
	Headers map[string]string
	Retries int
	Backoff time.Duration
	Timeout time.Duration
}
