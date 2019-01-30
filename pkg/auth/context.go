package auth

import (
	"context"
)

type key string

// UserVisibleVerified used in context as additional property
const userVisibleVerified key = "isUserVisibleVerified"

// WithUserVisibleVerified creates child context with additional information that this context is verified
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
