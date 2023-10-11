package throttler

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"time"
)

type Throttler struct {
	transport  http.RoundTripper
	retries    int
	ticker     *time.Ticker
	errContent string
}

const defaultErrContent = "exceeded"

type Option func(*Throttler)

func WithErrContent(errContent string) Option {
	return func(h *Throttler) {
		h.errContent = errContent
	}
}
func WithRoundTripper(transport http.RoundTripper) Option {
	return func(h *Throttler) {
		h.transport = transport
	}
}

// NewThrottler create a new rate limiter with request per second(rps) and retry times
func NewThrottler(rps int, retry int, options ...Option) *Throttler {
	interval := time.Second / time.Duration(rps)
	handler := Throttler{
		retries:    retry,
		errContent: defaultErrContent,
		transport:  http.DefaultTransport,
		ticker:     time.NewTicker(interval),
	}

	for _, option := range options {
		option(&handler)
	}
	return &handler
}
func (r *Throttler) shouldRetry(resp *http.Response, err error) bool {
	if resp == nil {
		return true
	}

	if err != nil {
		return false
	}
	if resp.StatusCode == http.StatusTooManyRequests || resp.StatusCode == http.StatusBadGateway {
		return true
	}
	var buf bytes.Buffer
	_, err = io.Copy(&buf, resp.Body)
	if err != nil {
		return false
	}
	defer resp.Body.Close()
	if bytes.Contains(buf.Bytes(), []byte(r.errContent)) {
		return true
	}
	resp.Body = io.NopCloser(&buf)
	return false
}

// RoundTrip implements the RoundTripper interface.
// Exponential backoff is used to retry requests that fail due to rate limiting.
func (r *Throttler) RoundTrip(req *http.Request) (*http.Response, error) {
	retryInternal := time.Second
	backoff := func() {
		// exponential backoff strategy
		retryInternal *= 2
		if retryInternal > time.Minute {
			retryInternal = time.Second
		}
		fmt.Println("backoff", retryInternal)
	}

	for i := 0; i < r.retries; i++ {
		<-r.ticker.C
		resp, err := r.transport.RoundTrip(req)
		if !r.shouldRetry(resp, err) {
			return resp, nil
		}

		timer := time.NewTimer(retryInternal)
		<-timer.C
		timer.Stop()
		backoff()
	}
	return nil, fmt.Errorf("retries exceeded %d", r.retries)
}
