package baseclient

import (
	"context"
	"net/http"
)

type clientContextKey string

const (
	headersClientContextKey clientContextKey = "headers"
)

// WithHTTPHeader creates creates child context provided header.
func WithHTTPHeader(ctx context.Context, key, value string) context.Context {
	headers := http.Header{}
	if v, ok := ctx.Value(headersClientContextKey).(http.Header); ok && v != nil {
		headers = v
	}
	headers.Add(key, value)
	return context.WithValue(ctx, headersClientContextKey, headers)
}

// SetContextHeaders sets extra headers that stored in context.
func SetContextHeaders(request *http.Request) {
	if request == nil {
		return
	}
	ctx := request.Context()
	if headers, ok := ctx.Value(headersClientContextKey).(http.Header); ok && len(headers) > 0 {
		if request.Header == nil {
			request.Header = http.Header{}
		}
		for key := range headers {
			request.Header.Set(key, headers.Get(key))
		}
	}
}
