package services

import (
	"context"
	"net/http"

	uuid "github.com/satori/go.uuid"
)

type servicesContextKey string

const (
	// internalRequestKey used in context as additional propetry
	internalRequestKey servicesContextKey = "isInternal"

	requestIDKey servicesContextKey = "requestIDKey"

	xRequestIDHeader = "X-Request-Id"
)

// WithInternalRequest creates child context with additional information
// that this context is for internal requests
func WithInternalRequest(ctx context.Context) context.Context {
	return context.WithValue(ctx, internalRequestKey, true)
}

// IsInternalRequest checks if context is for internal request
func IsInternalRequest(ctx context.Context) bool {
	value, ok := ctx.Value(internalRequestKey).(bool)
	return ok && value
}

// WithRequestID assign given request_id to context if there is no one in.
func WithRequestID(ctx context.Context, requestID string) context.Context {
	if ctx.Value(requestIDKey) != nil {
		return ctx
	}

	return context.WithValue(ctx, requestIDKey, requestID)
}

// WithGeneratedRequestID assign new generated request_id to context if there is no one in.
func WithGeneratedRequestID(ctx context.Context) context.Context {
	requestID := "req-" + uuid.NewV4().String()

	return context.WithValue(ctx, requestIDKey, requestID)
}

// RequestID retrieves request id from context.
func RequestID(ctx context.Context) string {
	requestID, ok := ctx.Value(requestIDKey).(string)
	if !ok {
		return "NO-REQUESTID"
	}

	return requestID
}

// ContextFromRequest gets context.Context from a received request.
func ContextFromRequest(r *http.Request) context.Context {
	return WithRequestID(r.Context(), r.Header.Get(xRequestIDHeader))
}
