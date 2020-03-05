package sync

import (
	"context"

	"github.com/Juniper/contrail/pkg/services"
	"github.com/Juniper/contrail/pkg/sync/replication"
)

type eventListProcessor interface {
	ProcessList(context.Context, *services.EventList) (*services.EventList, error)
}

type eventDecoder interface {
	DecodeRowEvent(operation, resourceName string, pk []string, properties map[string]interface{}) (*services.Event, error)
}

type EventChangeHandler struct {
	processor eventListProcessor
	decoder   eventDecoder
}

func (e *EventChangeHandler) Handle(ctx context.Context, changes []replication.Change) error {
	list := services.EventList{}
	for _, c := range changes {
		ev, err := e.decoder.DecodeRowEvent(
			changeOperationToServices(c.Operation()),
			c.Kind(),
			c.PK(),
			c.Data(),
		)
		if err != nil {
			return err
		}
		list.Events = append(list.Events, ev)
	}
	_, err := e.processor.ProcessList(ctx, &list)
	return err
}

func changeOperationToServices(op replication.ChangeOperation) string {
	switch op {
	case replication.CreateOperation:
		return services.OperationCreate
	case replication.UpdateOperation:
		return services.OperationUpdate
	case replication.DeleteOperation:
		return services.OperationDelete
	default:
		return services.OperationCreate
	}
}
