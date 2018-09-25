package types

import (
	"context"
)

type key string

const (
	// InternalRequestKey used in context as additional propetry
	InternalRequestKey key = "isInternal"
)

// GetInternalRequestContext TODO
func GetInternalRequestContext(ctx context.Context) context.Context {
	return context.WithValue(ctx, InternalRequestKey, true)
}

// IsInternalRequest TODO
func IsInternalRequest(ctx context.Context) bool {
	if v := ctx.Value(InternalRequestKey); v != nil {
		return v.(bool)
	}

	return false
}
