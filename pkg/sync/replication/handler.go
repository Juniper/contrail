package replication

import (
	"context"
	"fmt"

	"github.com/kyleconroy/pgoutput"
	"github.com/pkg/errors"
	"github.com/siddontang/go-mysql/canal"
	"github.com/siddontang/go-mysql/mysql"
	"github.com/siddontang/go-mysql/replication"
	"github.com/siddontang/go-mysql/schema"
	"github.com/sirupsen/logrus"

	"github.com/Juniper/contrail/pkg/logutil"
	"github.com/Juniper/contrail/pkg/services"
)

type relationSet map[uint32]pgoutput.Relation

type EventDecoder interface {
	DecodeRowEvent(operation, resourceName string, pk []string, properties map[string]interface{}) (*services.Event, error)
}

// PgoutputEventHandler handles replication messages by decoding them as events and passing them to processor.
type PgoutputEventHandler struct {
	decoder   EventDecoder
	processor services.EventProcessor
	log       *logrus.Entry

	relations relationSet
}

// NewPgoutputEventHandler creates new ReplicationEventHandler with provided decoder and processor.
func NewPgoutputEventHandler(p services.EventProcessor, d EventDecoder) *PgoutputEventHandler {
	return &PgoutputEventHandler{
		decoder:   d,
		processor: p,
		log:       logutil.NewLogger("replication-event-handler"),
		relations: relationSet{},
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

func (h *PgoutputEventHandler) handleDataEvent(ctx context.Context, operation string, relationID uint32, row []pgoutput.Tuple) error {
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

// CanalEventHandler handles canal events by decoding them as events and passing them to processor.
type CanalEventHandler struct {
	decoder   EventDecoder
	processor services.EventProcessor
	log       *logrus.Entry
}

// NewCanalEventHandler creates new CanalEventHandler with given decoder and processor.
func NewCanalEventHandler(p services.EventProcessor, d EventDecoder) *CanalEventHandler {
	return &CanalEventHandler{
		decoder:   d,
		processor: p,
		log:       logutil.NewLogger("canal-event-handler"),
	}
}

// OnRow handles row changed modifying key-values in etcd based on action.
func (h *CanalEventHandler) OnRow(e *canal.RowsEvent) error {
	h.log.WithFields(logrus.Fields{
		"action":  e.Action,
		"table":   e.Table.Name,
		"rows":    e.Rows,                              // verbose logging for development phase
		"columns": fmt.Sprintf("%+v", e.Table.Columns), // verbose logging for development phase
	}).Debug("Received OnRow event")

	ctx := context.TODO()
	switch e.Action {
	case canal.InsertAction:
		return h.handleCreate(ctx, e.Rows, e.Table)
	case canal.UpdateAction:
		return h.handleUpdate(ctx, e.Rows, e.Table)
	case canal.DeleteAction:
		return h.handleDelete(ctx, e.Rows, e.Table)
	default:
		return fmt.Errorf("invalid ROW event action: %s", e.Action)
	}
}

type handlerFunc func(context.Context, string, []string, map[string]interface{}) error

func (h *CanalEventHandler) handleOperation(
	ctx context.Context,
	rows [][]interface{},
	t *schema.Table,
	operation string,
) error {
	for _, row := range rows {
		pk, err := getPrimaryKeyValue(row, t)
		if err != nil {
			return errors.Wrapf(err, "table %v", t.Name)
		}
		kvs, err := getKeyValues(row, t.Columns)
		if err != nil {
			return fmt.Errorf("table %s error: %s", t, err)
		}
		ev, err := h.decoder.DecodeRowEvent(operation, t.Name, []string{pk}, kvs)
		if err != nil {
			return err
		}
		if _, err = h.processor.Process(ctx, ev); err != nil {
			return err
		}
	}
	return nil

}

func (h *CanalEventHandler) handleCreate(ctx context.Context, rows [][]interface{}, t *schema.Table) error {
	return errors.Wrap(h.handleOperation(ctx, rows, t, services.OperationCreate), "operation CREATE")
}

func (h *CanalEventHandler) handleUpdate(ctx context.Context, rows [][]interface{}, t *schema.Table) error {
	return errors.Wrap(h.handleOperation(ctx, rows, t, services.OperationUpdate), "operation UPDATE")
}

func (h *CanalEventHandler) handleDelete(ctx context.Context, rows [][]interface{}, t *schema.Table) error {
	for _, row := range rows {
		pk, err := getPrimaryKeyValue(row, t)
		if err != nil {
			return errors.Wrapf(err, "delete from %v", t.Name)
		}

		ev, err := h.decoder.DecodeRowEvent(services.OperationDelete, t.Name, []string{pk}, nil)
		if err != nil {
			return err
		}
		if _, err = h.processor.Process(ctx, ev); err != nil {
			return err
		}
	}
	return nil
}

func ensureSinglePrimaryKey(keys []string) (string, error) {
	if pkLen := len(keys); pkLen != 1 {
		return "", errors.Errorf("expected single element primary key but got %v elements", pkLen)
	}
	return keys[0], nil
}

func getPrimaryKeyValue(row []interface{}, t *schema.Table) (string, error) {
	keys, err := getAllPrimaryKeys(row, t)
	if err != nil {
		return "", err
	}
	return ensureSinglePrimaryKey(keys)
}

func getAllPrimaryKeys(row []interface{}, t *schema.Table) ([]string, error) {
	var keys []string
	v, err := canal.GetPKValues(t, row)
	if err != nil {
		return nil, err
	}
	keys, err = primaryKeyToStringSlice(v)
	if err != nil {
		return nil, errors.Wrapf(err, "table: %s error", t.Name)
	}
	return keys, nil
}

func primaryKeyToStringSlice(keyValues []interface{}) ([]string, error) {
	keys := []string{}
	for i, pk := range keyValues {
		if pk == nil || pk == "" {
			return nil, fmt.Errorf("primary key value is nil or empty on key element at index %v", i)
		}
		keys = append(keys, fmt.Sprint(pk))
	}
	return keys, nil
}

// getKeyValues uses the fact that columns are in the same order as values in a row
func getKeyValues(row []interface{}, columns []schema.TableColumn) (map[string]interface{}, error) {
	kvs := make(map[string]interface{}, len(columns))
	for i, c := range columns {
		v, err := getValue(row[i], c.Type)
		if err != nil {
			return nil, fmt.Errorf("cannot retrieve value of field %q with DB type %q: %s",
				c.Name, c.RawType, err)
		}
		kvs[c.Name] = v
	}
	return kvs, nil
}

func getValue(value interface{}, columnType int) (interface{}, error) {
	if columnType == schema.TYPE_BIT || columnType == schema.TYPE_ENUM || columnType == schema.TYPE_SET {
		return "", fmt.Errorf("unsupported DB column type %v", columnType)
	}
	return value, nil
}

// OnRotate event is skipped.
func (h *CanalEventHandler) OnRotate(*replication.RotateEvent) error {
	h.log.Debug("Skipping OnRotate event")
	return nil
}

// OnDDL event is skipped.
func (h *CanalEventHandler) OnDDL(mysql.Position, *replication.QueryEvent) error {
	h.log.Debug("Skipping OnDDL event")
	return nil
}

// OnXID event is skipped.
func (h *CanalEventHandler) OnXID(mysql.Position) error {
	h.log.Debug("Skipping OnXID event")
	return nil
}

// OnGTID event is skipped.
func (h *CanalEventHandler) OnGTID(mysql.GTIDSet) error {
	h.log.Debug("Skipping OnGTID event")
	return nil
}

// OnPosSynced event is skipped.
func (h *CanalEventHandler) OnPosSynced(mysql.Position, bool) error {
	h.log.Debug("Skipping OnPosSynced event")
	return nil
}

// String adds support for stringer interface.
func (h *CanalEventHandler) String() string {
	return "canalEventHandler"
}
