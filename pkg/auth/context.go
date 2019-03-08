package auth

import (
	"context"
	"net/http"

	"github.com/Juniper/contrail/pkg/format"
)

type key string

const (
	// internalRequestKey used in context as additional propetry
	internalRequestKey key = "isInternal"
	// xClusterID used in context, which will be set in HEADER
	// to select cluster specific keystone for auth
	xClusterID key = "X-Cluster-ID"
	// xAuthToktn used in context, which will be set in HEADER
	// to fetch cluster keystone token with infra(superuser) token
	xAuthToken key = "X-Auth-Token"
)

// WithInternalRequest creates child context with additional information
// that this context is for internal requests
func WithInternalRequest(ctx context.Context) context.Context {
	return context.WithValue(ctx, internalRequestKey, true)
}

// IsInternalRequest checks if context is for internal request
func IsInternalRequest(ctx context.Context) bool {
	if v := ctx.Value(internalRequestKey); v != nil {
		return v.(bool)
	}

	return false
}

// WithXClusterID creates child context with cluster ID
func WithXClusterID(ctx context.Context, clusterID string) context.Context {
	if v := ctx.Value(xClusterID); v == nil {
		return context.WithValue(ctx, xClusterID, clusterID)
	}
	return ctx
}

// SetXClusterIDInHeader sets X-Cluster-ID in the HEADER.
func SetXClusterIDInHeader(
	ctx context.Context, request *http.Request) *http.Request {
	if v := ctx.Value(xClusterID); v != nil {
		request.Header.Set(string(xClusterID), format.InterfaceToString(v))
	}
	return request
}

// WithXAuthToken creates child context with Auth Token
func WithXAuthToken(ctx context.Context, token string) context.Context {
	if v := ctx.Value(xAuthToken); v == nil {
		return context.WithValue(ctx, xAuthToken, token)
	}
	return ctx
}

// SetXAuthTokenInHeader sets X-Auth-Token in the HEADER.
func SetXAuthTokenInHeader(
	ctx context.Context, request *http.Request) *http.Request {
	if v := ctx.Value(xAuthToken); v != nil {
		request.Header.Set(string(xAuthToken), format.InterfaceToString(v))
	}
	return request
}
