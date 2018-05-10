package replication

import (
	"fmt"

	"github.com/Juniper/contrail/pkg/sync/etcd"
	"github.com/gogo/protobuf/proto"
)

// RowSink is data consumer capable of processing row data.
type RowSink interface {
	Create(resourceName string, pk string, properties map[string]interface{}) error
	Update(resourceName string, pk string, properties map[string]interface{}) error
	Delete(resourceName string, pk string) error
}

type rowScanner interface {
	ScanRow(schemaID string, rowData map[string]interface{}) (proto.Message, error)
}

type objectMappingAdapter struct {
	etcd.Sink
	rs rowScanner
}

// NewObjectMappingAdapter creates adapter that maps row data into objects using
// provided rowScanner and pushes that object to sink.
func NewObjectMappingAdapter(s etcd.Sink, rs rowScanner) RowSink {
	return &objectMappingAdapter{Sink: s, rs: rs}
}

func (o *objectMappingAdapter) Create(resourceName string, pk string, properties map[string]interface{}) error {
	obj, err := o.rs.ScanRow(resourceName, properties)
	if err != nil {
		return fmt.Errorf("error scanning row: %v", err)
	}
	return o.Sink.Create(resourceName, pk, obj)
}

func (o *objectMappingAdapter) Update(resourceName string, pk string, properties map[string]interface{}) error {
	obj, err := o.rs.ScanRow(resourceName, properties)
	if err != nil {
		return fmt.Errorf("error scanning row: %v", err)
	}
	return o.Sink.Update(resourceName, pk, obj)
}
