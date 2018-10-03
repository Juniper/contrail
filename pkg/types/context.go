package types

import (
	"context"
)

type key string

const (
	// InternalRequestKey used in context as additional propetry
	InternalRequestKey key = "isInternal"
)

// WithInternalRequest creates child context with additional information
// that this context is for internal requests
func WithInternalRequest(ctx context.Context) context.Context {
	return context.WithValue(ctx, InternalRequestKey, true)
}

// IsInternalRequest checks if context is for internal request
func IsInternalRequest(ctx context.Context) bool {
	if v := ctx.Value(InternalRequestKey); v != nil {
		return v.(bool)
	}

	return false
}
