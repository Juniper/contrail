package replication

import (
	"context"
	"fmt"

	"github.com/kyleconroy/pgoutput"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"

	"github.com/Juniper/contrail/pkg/logutil"
	"github.com/Juniper/contrail/pkg/services"
)

// PgoutputEventHandler handles replication messages by decoding them as events and passing them to processor.
type PgoutputEventHandler struct {
	decoder   EventDecoder
	processor services.EventProcessor
	log       *logrus.Entry

	relations relationSet

	idToFQName fqNameCache
}

type fqNameCache map[string][]string

type relationSet map[uint32]pgoutput.Relation

// NewPgoutputEventHandler creates new ReplicationEventHandler with provided decoder and processor.
func NewPgoutputEventHandler(p services.EventProcessor, d EventDecoder) *PgoutputEventHandler {
	return &PgoutputEventHandler{
		decoder:    d,
		processor:  p,
		log:        logutil.NewLogger("replication-event-handler"),
		relations:  relationSet{},
		idToFQName: fqNameCache{},
	}
}

// Handle handles provided message by passing decoding its contents passing them to processor.
func (h *PgoutputEventHandler) Handle(ctx context.Context, msg pgoutput.Message) error {
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

func (h *PgoutputEventHandler) handleDataEvent(
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
	h.updateFQNameCache(ev)

	if err = h.sanitizeCreateRefEvent(ev); err != nil {
		return errors.Wrapf(err, "failed to sanitize reference fqName, event: %v", ev)
	}
	_, err = h.processor.Process(ctx, ev)
	return err
}

func (h *PgoutputEventHandler) updateFQNameCache(event *services.Event) {
	switch er := event.GetRequest().(type) {
	case services.CreateEventRequest:
		r := er.GetRequest().GetResource()
		h.idToFQName[r.GetUUID()] = r.GetFQName()
	case services.UpdateEventRequest:
		r := er.GetRequest().GetResource()
		h.idToFQName[r.GetUUID()] = r.GetFQName()
	case services.DeleteEventRequest:
		delete(h.idToFQName, er.GetRequest().GetID())
	}
}

func (h *PgoutputEventHandler) sanitizeCreateRefEvent(event *services.Event) error {
	refEvent, ok := event.GetRequest().(services.CreateRefEventRequest)
	if !ok {
		return nil
	}
	reference := refEvent.GetRequest().GetReference()
	fqName, ok := h.idToFQName[reference.GetUUID()]
	if !ok {
		return errors.Errorf("failed to fetched reference fq_name for uuid: '%v'", reference.GetUUID())
	}
	reference.SetTo(fqName)
	return nil
}

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
