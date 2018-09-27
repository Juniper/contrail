package replication

import (
	"context"
	"fmt"
	"strings"

	"github.com/Juniper/contrail/pkg/models/basemodels"
	"github.com/Juniper/contrail/pkg/schema"
	"github.com/Juniper/contrail/pkg/sync/sink"
	"github.com/pkg/errors"
)

// RowSink is data consumer capable of processing row data (aka db logic).
type RowSink interface {
	Create(ctx context.Context, resourceName string, pk []string, properties map[string]interface{}) error
	Update(ctx context.Context, resourceName string, pk []string, properties map[string]interface{}) error
	Delete(ctx context.Context, resourceName string, pk []string) error
}

type rowScanner interface {
	ScanRow(schemaID string, rowData map[string]interface{}) (basemodels.Object, error)
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
	ctx context.Context, resourceName string, pk []string, properties map[string]interface{},
) error {
	obj, err := o.rs.ScanRow(resourceName, properties)
	if err != nil {
		return fmt.Errorf("error scanning row: %v", err)
	}
	pkLen := len(pk)
	if strings.HasPrefix(resourceName, schema.RefPrefix) {
		if pkLen != 2 {
			return errors.Errorf("CreateRef expect two-element primary key but got %v elements", pkLen)
		}
		return o.Sink.CreateRef(ctx, resourceName, pk, obj)
	}
	if pkLen != 1 {
		return errors.Errorf("single element primary key expected for resource %v, %v elems provided", resourceName, pkLen)
	}
	return o.Sink.Create(ctx, resourceName, pk[0], obj)
}

func (o *objectMappingAdapter) Update(
	ctx context.Context, resourceName string, pk []string, properties map[string]interface{},
) error {
	pkLen := len(pk)
	if strings.HasPrefix(resourceName, schema.RefPrefix) {
		return errors.New("method UPDATE not available on ref_* resources - this is a bug")
	}
	if k := len(pk); k != 1 {
		return errors.Errorf("single element primary key expected for resource %v, %v elems provided", resourceName, pkLen)
	}
	obj, err := o.rs.ScanRow(resourceName, properties)
	if err != nil {
		return fmt.Errorf("error scanning row: %v", err)
	}
	return o.Sink.Update(ctx, resourceName, pk[0], obj)
}

func (o *objectMappingAdapter) Delete(
	ctx context.Context, resourceName string, pk []string) error {
	pkLen := len(pk)
	if strings.HasPrefix(resourceName, schema.RefPrefix) {
		if pkLen != 2 {
			return errors.Errorf("CreateRef expect two-element primary key but got %v elements", pkLen)
		}
		return o.Sink.DeleteRef(ctx, resourceName, pk)
	}
	if k := len(pk); k != 1 {
		return errors.Errorf("single element primary key expected for resource %v, %v elems provided", resourceName, pkLen)
	}
	return o.Sink.Delete(ctx, resourceName, pk[0])
}
