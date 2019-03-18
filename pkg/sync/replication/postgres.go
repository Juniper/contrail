package replication

import (
	"context"
	"fmt"
	"github.com/Juniper/contrail/pkg/logutil"
	"github.com/Juniper/contrail/pkg/services"
	"github.com/kyleconroy/pgoutput"
	"github.com/sirupsen/logrus"
)

// EventDecoder is capable of decoding row data in form of map into an Event.
type EventDecoder interface {
	DecodeRowEvent(operation, resourceName string, pk []string, properties map[string]interface{}) (*services.Event, error)
}

// PgOutputEventHandler handles replication messages by decoding them as events and passing them to processor.
type PgOutputEventHandler struct {
	decoder   EventDecoder
	processor services.EventProcessor
	log       *logrus.Entry

	relations relationSet

	idToFQName map[string][]string
}

type relationSet map[uint32]pgoutput.Relation

// NewPgOutputEventHandler creates new ReplicationEventHandler with provided decoder and processor.
func NewPgOutputEventHandler(p services.EventProcessor, d EventDecoder) *PgOutputEventHandler {
	return &PgOutputEventHandler{
		decoder:   d,
		processor: p,
		log:       logutil.NewLogger("replication-event-handler"),
		relations: relationSet{},
	}
}

// Handle handles provided message by passing decoding its contents passing them to processor.
func (h *PgOutputEventHandler) Handle(ctx context.Context, msg pgoutput.Message) error {
	switch v := msg.(type) {
	case pgoutput.Relation:
		h.log.Debug("received RELATION message")
		h.relations[v.ID] = v
	case pgoutput.Insert:
		h.log.Debug("received INSERT message")
		return h.handleDataEvent(ctx, services.OperationCreate, v.RelationID, v.Row)
	case pgoutput.Update:
		h.log.Debug("received UPDATE message")
		return h.handleDataEvent(ctx, services.OperationUpdate, v.RelationID, v.Row)
	case pgoutput.Delete:
		h.log.Debug("received DELETE message")
		return h.handleDataEvent(ctx, services.OperationDelete, v.RelationID, v.Row)
	}
	return nil
}

func (h *PgOutputEventHandler) handleDataEvent(
	ctx context.Context, operation string, relationID uint32, row []pgoutput.Tuple,
) error {
	relation, ok := h.relations[relationID]
	if !ok {
		return fmt.Errorf("no relation for %d", relationID)
	}

	pk, data, err := decodeRowData(relation, row)
	if err != nil {
		return fmt.Errorf("error decoding row: %v", err)
	}
	if len(pk) == 0 {
		return fmt.Errorf("no primary key specified for row: %v", row)
	}

	ev, err := h.decoder.DecodeRowEvent(operation, relation.Name, pk, data)
	if err != nil {
		return err
	}
	_, err = h.processor.Process(ctx, ev)
	return err
}

func (h *PgOutputEventHandler) updateCache(event *services.Event) {
	switch er := event.Request.(type) {
	case services.CreateEventRequest:
		r := er.GetRequest().GetResource()
		h.idToFQName[r.GetUUID()] = r.GetFQName()
	}
}

func

func decodeRowData(
	relation pgoutput.Relation,
	row []pgoutput.Tuple,
) (pk []string, data map[string]interface{}, err error) {
	keys, data := []interface{}{}, map[string]interface{}{}

	if t, c := len(row), len(relation.Columns); t != c {
		return nil, nil, fmt.Errorf("malformed message or relation columns, got %d values but relation has %d columns", t, c)
	}

	for i, tuple := range row {
		col := relation.Columns[i]
		decoder := col.Decoder()
		if err = decoder.DecodeText(nil, tuple.Value); err != nil {
			return nil, nil, fmt.Errorf("error decoding column '%v': %s", col.Name, err)
		}
		value := decoder.Get()
		data[col.Name] = value
		if col.Key {
			keys = append(keys, value)
		}

	}

	pk, err = primaryKeyToStringSlice(keys)
	if err != nil {
		return nil, nil, fmt.Errorf("error creating PK: %v", err)
	}

	return pk, data, nil
}
