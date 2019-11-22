package client

import (
	"github.com/Juniper/contrail/pkg/services"
)

// GRPC is a GRPC API server client.
type GRPC struct {
	c services.ContrailServiceClient
}

// NewGRPC returns a new GRPC.
func NewGRPC(c services.ContrailServiceClient) *GRPC {
	return &GRPC{c: c}
}
