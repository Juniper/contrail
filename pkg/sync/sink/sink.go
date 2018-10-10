package sink

import (
	"context"

	"github.com/Juniper/contrail/pkg/db/basedb"
)

// Sink represents service that handles resources object logic rather than db loic.
type Sink interface {
	Create(ctx context.Context, resourceName string, pk string, obj basedb.Object) error
	Update(ctx context.Context, resourceName string, pk string, obj basedb.Object) error
	Delete(ctx context.Context, resourceName string, pk string) error

	CreateRef(ctx context.Context, resourceName string, pk []string, obj basedb.Object) error
	DeleteRef(ctx context.Context, resourceName string, pk []string) error
}
