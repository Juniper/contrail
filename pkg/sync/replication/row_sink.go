package replication

import (
	"context"
	"fmt"

	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/sync/sink"
)

// RowSink is data consumer capable of processing row data.
type RowSink interface {
	Create(ctx context.Context, resourceName string, pk string, properties map[string]interface{}) error
	Update(ctx context.Context, resourceName string, pk string, properties map[string]interface{}) error
	Delete(ctx context.Context, resourceName string, pk string) error
}

type rowScanner interface {
	ScanRow(schemaID string, rowData map[string]interface{}) (models.Object, error)
}

type objectMappingAdapter struct {
	sink.Sink
	rs rowScanner
}

// NewObjectMappingAdapter creates adapter that maps row data into objects using
// provided rowScanner and pushes that object to sink.
func NewObjectMappingAdapter(s sink.Sink, rs rowScanner) RowSink {
	return &objectMappingAdapter{Sink: s, rs: rs}
}

func (o *objectMappingAdapter) Create(
	ctx context.Context, resourceName string, pk string, properties map[string]interface{},
) error {
	obj, err := o.rs.ScanRow(resourceName, properties)
	if err != nil {
		return fmt.Errorf("error scanning row: %v", err)
	}
	return o.Sink.Create(ctx, resourceName, pk, obj)
}

func (o *objectMappingAdapter) Update(
	ctx context.Context, resourceName string, pk string, properties map[string]interface{},
) error {
	obj, err := o.rs.ScanRow(resourceName, properties)
	if err != nil {
		return fmt.Errorf("error scanning row: %v", err)
	}
	return o.Sink.Update(ctx, resourceName, pk, obj)
}
