package auth

import (
	"context"
)

type key string

const (
	// InternalRequestKey used in context as additional propetry
	InternalRequestKey key = "isInternal"

	// UserVisibleVerified used in context as additional property
	userVisibleVerified key = "isUserVisibleVerified"
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

// WithUserVisibleVerified creates child context with additional information
// that this context is verified
func WithUserVisibleVerified(ctx context.Context) context.Context {
	if v := ctx.Value(userVisibleVerified); v == nil {
		return context.WithValue(ctx, userVisibleVerified, true)
	}
	return ctx
}

// IsUserVisibleVerified checks if context is verified
func IsUserVisibleVerified(ctx context.Context) bool {
	if v := ctx.Value(userVisibleVerified); v != nil {
		return v.(bool)
	}
	return false
}
