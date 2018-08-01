package sink

import (
	"context"

	"github.com/Juniper/contrail/pkg/models/basemodels"
)

// Sink represents service that handler transfers data to.
type Sink interface {
	Create(ctx context.Context, resourceName string, pk string, obj basemodels.Object) error
	Update(ctx context.Context, resourceName string, pk string, obj basemodels.Object) error
	Delete(ctx context.Context, resourceName string, pk string) error
}
