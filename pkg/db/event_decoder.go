package db

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/pkg/errors"

	"github.com/Juniper/contrail/pkg/format"
	"github.com/Juniper/contrail/pkg/schema"
	"github.com/Juniper/contrail/pkg/services"
)

// DecodeRowEvent transforms row data into *services.Event.
// TODO(Michal): Unit test this.
func (db *Service) DecodeRowEvent(
	operation, resourceName string, pk []string, properties map[string]interface{},
) (*services.Event, error) {
	obj, fields, err := db.ScanRow(resourceName, properties)
	if err != nil {
		return nil, fmt.Errorf("error scanning row: %v", err)
	}

	pkLen := len(pk)
	isRef := strings.HasPrefix(resourceName, schema.RefPrefix)
	switch {
	case pkLen == 1 && !isRef:
		return services.NewEvent(&services.EventOption{
			UUID:      pk[0],
			Kind:      resourceName,
			Data:      obj.ToMap(),
			Operation: operation,
			FieldMask: fields,
		})
	case pkLen == 2 && isRef:
		if operation == services.OperationUpdate {
			return nil, errors.New("method UPDATE not available on ref_* resources - received ref-relax event")
		}
		return services.NewRefUpdateEvent(services.RefUpdateOption{
			Operation:     services.ParseRefOperation(operation),
			ReferenceType: refTableNameToReferenceName(resourceName),
			FromUUID:      pk[0],
			ToUUID:        pk[1],
			AttrData:      json.RawMessage(format.MustJSON(obj)),
		})
	}
	return nil, errors.Errorf(
		"%s row: unhandled case with table %v and primary key with %v elements", operation, resourceName, pkLen,
	)
}

func refTableNameToReferenceName(name string) string {
	return strings.Replace(strings.TrimPrefix(name, "ref_"), "_", "-", -1)
}
