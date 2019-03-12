package db

import (
	"fmt"
	"strings"

	"github.com/gogo/protobuf/types"
	"github.com/pkg/errors"

	"github.com/Juniper/contrail/pkg/db/basedb"
	"github.com/Juniper/contrail/pkg/models/basemodels"
	"github.com/Juniper/contrail/pkg/schema"
	"github.com/Juniper/contrail/pkg/services"
)

// DecodeRowEvent transforms row data into *services.Event.
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
		return decodeResourceEvent(operation, resourceName, pk, obj, fields)
	case pkLen == 2 && isRef:
		return decodeReferenceEvent(operation, resourceName, pk, obj)
	}
	return nil, errors.Errorf(
		"%s row: unhandled case with table %v and primary key with %v elements", operation, resourceName, pkLen,
	)
}

func decodeResourceEvent(
	operation, resourceName string, pk []string, obj basedb.Object, fm *types.FieldMask,
) (*services.Event, error) {
	m := obj.ToMap()
	if fm != nil {
		m = basemodels.ApplyFieldMask(m, *fm)
	}
	return services.NewEvent(&services.EventOption{
		UUID:      pk[0],
		Kind:      resourceName,
		Data:      m,
		Operation: operation,
		FieldMask: fm,
	})
}

func decodeReferenceEvent(operation, resourceName string, pk []string, attr basedb.Object) (*services.Event, error) {
	if operation == services.OperationUpdate {
		return nil, errors.New("method UPDATE not available on ref_* resources - received ref-relax event")
	}
	return services.NewRefUpdateEvent(services.RefUpdateOption{
		Operation:     services.ParseRefOperation(operation),
		ReferenceType: refTableNameToReferenceName(resourceName),
		FromUUID:      pk[0],
		ToUUID:        pk[1],
		Attr:          attr,
	})
}

func refTableNameToReferenceName(name string) string {
	return strings.Replace(strings.TrimPrefix(name, "ref_"), "_", "-", -1)
}
