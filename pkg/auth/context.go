package auth

import (
	"context"
)

type key string

// UserVisibleVerified used in context as additional property
const UserVisibleVerified key = "isUserVisibleVerified"

// IsUserVisibleVerified creates child context with additional information that this context is verified
func IsUserVisibleVerified(ctx context.Context) (bool, context.Context) {
	if v := ctx.Value(UserVisibleVerified); v != nil {
		return v.(bool), ctx
	}
	return false, context.WithValue(ctx, UserVisibleVerified, true)
}
