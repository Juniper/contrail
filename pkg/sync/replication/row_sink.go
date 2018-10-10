package replication

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/pkg/errors"

	"github.com/Juniper/contrail/pkg/db/basedb"
	"github.com/Juniper/contrail/pkg/schema"
	"github.com/Juniper/contrail/pkg/sync/sink"
)

// RowSink is data consumer capable of processing row data (aka db logic).
type RowSink interface {
	Create(ctx context.Context, resourceName string, pk []string, properties map[string]interface{}) error
	Update(ctx context.Context, resourceName string, pk []string, properties map[string]interface{}) error
	Delete(ctx context.Context, resourceName string, pk []string) error
}

type rowScanner interface {
	ScanRow(schemaID string, rowData map[string]interface{}) (basedb.Object, error)
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
	pkLen := len(pk)
	obj, err := o.rs.ScanRow(resourceName, properties)
	if err != nil {
		return errors.Wrap(err, "error scanning row")
	}
	isRef := strings.HasPrefix(resourceName, schema.RefPrefix)
	switch {
	case pkLen == 1 && !isRef:
		return o.Sink.Create(ctx, resourceName, pk[0], obj)
	case pkLen == 2 && isRef:
		return o.Sink.CreateRef(ctx, resourceName, pk, obj)
	}
	return errors.Errorf("create row: unhandled case with table %v and primary key with %v elements", resourceName, pkLen)
}

func scanAttr(properties map[string]interface{}) (map[string]interface{}, error) {
	attr := map[string]interface{}{}
	exclude := map[string]bool{"from": true, "to": true}

	for key, value := range properties {
		if exclude[key] {
			continue
		}
		switch b := value.(type) {
		case []byte:
			var v interface{}
			err := json.Unmarshal(b, &v)
			if err != nil {
				return nil, err
			}
			attr[key] = v
		default:
			attr[key] = value
		}
	}
	return attr, nil
}

func (o *objectMappingAdapter) Update(
	ctx context.Context, resourceName string, pk []string, properties map[string]interface{},
) error {
	pkLen := len(pk)
	isRef := strings.HasPrefix(resourceName, schema.RefPrefix)
	switch {
	case pkLen == 1 && !isRef:
		obj, err := o.rs.ScanRow(resourceName, properties)
		if err != nil {
			return fmt.Errorf("error scanning row: %v", err)
		}
		return o.Sink.Update(ctx, resourceName, pk[0], obj)
	case pkLen == 2 && isRef:
		return errors.New("method UPDATE not available on ref_* resources - this is a bug")
	}
	return errors.Errorf("update row: unhandled case with table %v and primary key with %v elements", resourceName, pkLen)
}

func (o *objectMappingAdapter) Delete(
	ctx context.Context, resourceName string, pk []string) error {
	pkLen := len(pk)
	isRef := strings.HasPrefix(resourceName, schema.RefPrefix)
	switch {
	case pkLen == 1 && !isRef:
		return o.Sink.Delete(ctx, resourceName, pk[0])
	case pkLen == 2 && isRef:
		return o.Sink.DeleteRef(ctx, resourceName, pk)
	}
	return errors.Errorf("delete row: unhandled case with table %v and primary key with %v elements", resourceName, pkLen)
}
