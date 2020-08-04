package httputil

import (
	"net/http"

	"golang.org/x/time/rate"
)

// Doer is a simplest http client.
type Doer interface {
	Do(req *http.Request) (*http.Response, error)
}

// DoerFunc is a function that implements Doer interface.
type DoerFunc func(req *http.Request) (*http.Response, error)

// Do executes the function.
func (f DoerFunc) Do(req *http.Request) (*http.Response, error) {
	return f(req)
}

// RateLimitingDoer is a doer that limits request throughput.
type RateLimitingDoer struct {
	Doer

	limiter *rate.Limiter
}

// NewRateLimitingDoer creates a new doer that limits the throughput to match given rate.
func NewRateLimitingDoer(d Doer, rps rate.Limit, burst int) *RateLimitingDoer {
	return &RateLimitingDoer{Doer: d, limiter: rate.NewLimiter(rps, burst)}
}

// Do executes the request.
func (d *RateLimitingDoer) Do(req *http.Request) (*http.Response, error) {
	if err := d.limiter.Wait(req.Context()); err != nil {
		return &http.Response{}, err
	}
	return d.Doer.Do(req)
}
