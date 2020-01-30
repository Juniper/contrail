package httputil

import (
	"context"
	"net/http"
)

type httputilContextKey string

const (
	headersClientContextKey httputilContextKey = "headers"
)

// WithHTTPHeader creates child context with provided header.
func WithHTTPHeader(ctx context.Context, key, value string) context.Context {
	headers := http.Header{}
	if v, ok := ctx.Value(headersClientContextKey).(http.Header); ok && v != nil {
		headers = v.Clone()
	}
	headers.Set(key, value)
	return context.WithValue(ctx, headersClientContextKey, headers)
}

// SetContextHeaders sets extra headers that are stored in context.
func SetContextHeaders(r *http.Request) {
	if r == nil {
		return
	}
	if headers, ok := r.Context().Value(headersClientContextKey).(http.Header); ok && len(headers) > 0 {
		if r.Header == nil {
			r.Header = http.Header{}
		}
		for key := range headers {
			r.Header.Set(key, headers.Get(key))
		}
	}
}
