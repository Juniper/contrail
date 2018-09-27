package sink

import (
	"context"
	"strings"

	"github.com/Juniper/contrail/pkg/models/basemodels"
	"github.com/Juniper/contrail/pkg/schema"
	"github.com/Juniper/contrail/pkg/services"
	"github.com/Juniper/contrail/pkg/sync"
	"github.com/pkg/errors"
)

// EventProcessorSink is a Sink that dispatches events to processor.
type EventProcessorSink struct {
	services.EventProcessor
}

// Create dispatches OperationCreate event to processor.
func (e *EventProcessorSink) Create(ctx context.Context, resourceName string, pk []string, obj basemodels.Object) error {
	var err error
	var ev *services.Event
	//func resolveReferenceTable(name string) (typeName, refType string) {
	if strings.HasPrefix(resourceName, schema.RefPrefix) {
		if len(pk) != 2 {
			return errors.Errorf("expecting primary key with 2 items, got %d instead", len(pk))
		}
		typeName, typeRef := resolveReferenceTable(resourceName)
		ev, err = services.NewEventFromRefUpdate(&services.RefUpdate{
			Operation: services.RefOperationAdd,
			Type:      typeName,
			UUID:      pk[0],
			RefType:   typeRef,
			RefUUID:   pk[1],
		})
	} else {
		ev, err = services.NewEvent(&services.EventOption{
			UUID:      pk,
			Kind:      resourceName,
			Data:      obj.ToMap(),
			Operation: services.OperationCreate,
		})
	}
	if err != nil {
		return err
	}
	return e.process(ctx, ev)
}

// Update dispatches OperationUpdate event to processor.
func (e *EventProcessorSink) Update(ctx context.Context, resourceName string, pk []string, obj basemodels.Object) error {
	ev, err := services.NewEvent(&services.EventOption{
		UUID:      pk,
		Kind:      resourceName,
		Data:      obj.ToMap(),
		Operation: services.OperationUpdate,
	})
	if err != nil {
		return err
	}
	return e.process(ctx, ev)
}

// Delete dispatches OperationDelete event to processor.
func (e *EventProcessorSink) Delete(ctx context.Context, resourceName string, pk []string) error {
	var err error
	var ev *services.Event
	//func resolveReferenceTable(name string) (typeName, refType string) {
	if strings.HasPrefix(resourceName, schema.RefPrefix) {
		if len(pk) != 2 {
			return errors.Errorf("expecting primary key with 2 items, got %d instead", len(pk))
		}
		typeName, typeRef := resolveReferenceTable(resourceName)
		ev, err = services.NewEventFromRefUpdate(&services.RefUpdate{
			Operation: services.RefOperationAdd,
			Type:      typeName,
			UUID:      pk[0],
			RefType:   typeRef,
			RefUUID:   pk[1],
		})
	} else {
		ev, err = services.NewEvent(&services.EventOption{
			UUID:      pk,
			Kind:      resourceName,
			Operation: services.OperationDelete,
		})
	}
	if err != nil {
		return err
	}
	return e.process(ctx, ev)
}

func (e *EventProcessorSink) process(ctx context.Context, ev *services.Event) error {
	_, err := e.Process(ctx, ev)
	return err
}
