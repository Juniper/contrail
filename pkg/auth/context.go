package auth

import (
	"context"

	"github.com/Juniper/contrail/pkg/client/baseclient"
)

type key string

const (
	// internalRequestKey used in context as additional propetry
	internalRequestKey key = "isInternal"

	// xClusterIDHeader is a header used to select cluster specific keystone for auth
	xClusterIDHeader = "X-Cluster-ID"
	// xAuthTokenHeader is a header used by keystone to store user auth tokens
	xAuthTokenHeader = "X-Auth-Token"
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
	return baseclient.WithHTTPHeader(ctx, xClusterIDHeader, clusterID)
}

// WithXAuthToken creates child context with Auth Token
func WithXAuthToken(ctx context.Context, token string) context.Context {
	return baseclient.WithHTTPHeader(ctx, xAuthTokenHeader, token)
}
