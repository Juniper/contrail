package keystone

import (
	"context"

	"github.com/Juniper/asf/pkg/httputil"
)

const (
	// xClusterIDHeader is a header used to select cluster specific keystone for auth
	xClusterIDHeader = "X-Cluster-ID"
)

// WithXClusterID creates child context with cluster ID
func WithXClusterID(ctx context.Context, clusterID string) context.Context {
	return httputil.WithHTTPHeader(ctx, xClusterIDHeader, clusterID)
}
