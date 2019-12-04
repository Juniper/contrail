package auth

import (
	"context"
)

type Identity interface {
	GetObjPerms() interface{}
	IsAdmin() bool
	IsGlobalRORole() bool
	IsCloudAdminRole() bool
	ProjectID() string
	DomainID() string
	UserID() string
	AuthToken() string
	Roles() []string
}

type defaultIdentity struct{}

// TODO(mblotniak): do not assume admin
func (d defaultIdentity) GetObjPerms() interface{} { return nil }
func (d defaultIdentity) IsAdmin() bool            { return true }
func (d defaultIdentity) IsGlobalRORole() bool     { return true }
func (d defaultIdentity) IsCloudAdminRole() bool   { return true }
func (d defaultIdentity) ProjectID() string        { return "admin" }
func (d defaultIdentity) DomainID() string         { return "admin" } // TODO(mblotniak): Verify correctness
func (d defaultIdentity) UserID() string           { return "admin" }
func (d defaultIdentity) AuthToken() string        { return "" }
func (d defaultIdentity) Roles() []string          { return nil }

type authContextKey string

const (
	authIdentityContextKey authContextKey = "auth"
)

// WithIdentity returns context with auth.Context stored.
func WithIdentity(ctx context.Context, auth Identity) context.Context {
	return context.WithValue(ctx, authIdentityContextKey, auth)
}

// GetIdentity is used to get an authentication from ctx.Context.
func GetIdentity(ctx context.Context) Identity {
	c, ok := ctx.Value(authIdentityContextKey).(Identity)
	if !ok || c == nil {
		return defaultIdentity{}
	}
	return c
}
