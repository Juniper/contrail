package replication

import (
	"context"
	"errors"
	"fmt"

	pkglog "github.com/Juniper/contrail/pkg/log"
	"github.com/kyleconroy/pgoutput"
	"github.com/siddontang/go-mysql/canal"
	"github.com/siddontang/go-mysql/mysql"
	"github.com/siddontang/go-mysql/replication"
	"github.com/siddontang/go-mysql/schema"
	"github.com/sirupsen/logrus"
)

type relationAddGetter interface {
	Add(pgoutput.Relation)
	Get(id uint32) (pgoutput.Relation, error)
}

type rowData map[string]interface{}
type rowKey []interface{}

func (r rowKey) String() string {
	v := ([]interface{})(r)
	return fmt.Sprint(v...)
}

// RowSink is data consumer capable of processing row data.
type RowSink interface {
	HandleCreate(ctx context.Context, schemaID string, pk string, data map[string]interface{}) error
	HandleUpdate(ctx context.Context, schemaID string, pk string, data map[string]interface{}) error
	HandleDelete(ctx context.Context, schemaID string, pk string) error
}

// PgoutputEventHandler handles replication messages by pushing it to sink.
type PgoutputEventHandler struct {
	sink RowSink
	log  *logrus.Entry

	relations relationAddGetter
}

// NewPgoutputEventHandler creates new ReplicationEventHandler using sink provided as an argument.
func NewPgoutputEventHandler(s RowSink) *PgoutputEventHandler {
	return &PgoutputEventHandler{
		sink:      s,
		log:       pkglog.NewLogger("replication-event-handler"),
		relations: &relationSet{},
	}
}

// Handle handles provided message by passing its contents to sink or returns an error
func (h *PgoutputEventHandler) Handle(ctx context.Context, msg pgoutput.Message) error {

	switch v := msg.(type) {
	case pgoutput.Relation:
		h.log.Debug("received RELATION message")
		h.relations.Add(v)
	case pgoutput.Insert:
		h.log.Debug("received INSERT message")
		tableName, pk, data, err := h.resolveRow(v.RelationID, v.Row)
		if err != nil {
			return err
		}

		return h.sink.HandleCreate(ctx, tableName, pk.String(), data)
	case pgoutput.Update:
		h.log.Debug("received UPDATE message")
		tableName, pk, data, err := h.resolveRow(v.RelationID, v.Row)
		if err != nil {
			return err
		}

		return h.sink.HandleUpdate(ctx, tableName, pk.String(), data)
	case pgoutput.Delete:
		h.log.Debug("received DELETE message")
		tableName, pk, _, err := h.resolveRow(v.RelationID, v.Row)
		if err != nil {
			return err
		}

		return h.sink.HandleDelete(ctx, tableName, pk.String())
	}
	return nil
}

func (h *PgoutputEventHandler) resolveRow(
	relationID uint32,
	row []pgoutput.Tuple,
) (tableName string, pk rowKey, data rowData, err error) {
	relation, err := h.relations.Get(relationID)
	if err != nil {
		return "", nil, nil, err
	}

	key, data, err := decodeRowData(relation, row)
	if err != nil {
		return "", nil, nil, fmt.Errorf("error decoding row: %v", err)
	}

	return relation.Name, key, data, nil

}

func decodeRowData(
	relation pgoutput.Relation,
	row []pgoutput.Tuple,
) (key rowKey, data rowData, err error) {
	key, data = rowKey{}, rowData{}

	if t, c := len(row), len(relation.Columns); t != c {
		return nil, nil, fmt.Errorf("malformed message or relation columns, got %d values but relation has %d columns", t, c)
	}

	for i, tuple := range row {
		col := relation.Columns[i]
		decoder := getDecoder(col)
		if err = decoder.DecodeText(nil, tuple.Value); err != nil {
			return nil, nil, fmt.Errorf("error decoding column '%v': %s", col.Name, err)
		}
		value := decoder.Get()
		data[col.Name] = value
		if col.Key {
			key = append(key, value)
		}

	}

	if len(key) == 0 {
		return nil, nil, errors.New("no key values provided")
	}

	return key, data, nil
}

// CanalEventHandler handles canal events by pushing it to sink.
type CanalEventHandler struct {
	sink RowSink
	log  *logrus.Entry
}

// NewCanalEventHandler creates new CanalEventHandler with given sink.
func NewCanalEventHandler(s RowSink) *CanalEventHandler {
	return &CanalEventHandler{
		sink: s,
		log:  pkglog.NewLogger("canal-event-handler"),
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

	switch e.Action {
	case canal.InsertAction:
		return h.handleCreate(e.Rows, e.Table)
	case canal.UpdateAction:
		return h.handleUpdate(e.Rows, e.Table)
	case canal.DeleteAction:
		return h.handleDelete(e.Rows, e.Table)
	default:
		return fmt.Errorf("invalid ROW event action: %s", e.Action)
	}
}

// TODO(daniel): remove duplication
// nolint: dupl
func (h *CanalEventHandler) handleCreate(rows [][]interface{}, t *schema.Table) error {
	for _, row := range rows {
		pk, err := getPrimaryKeyValue(row, t)
		if err != nil {
			return err
		}
		kvs, err := getKeyValues(row, t.Columns)
		if err != nil {
			return fmt.Errorf("table %s error: %s", t, err)
		}
		_, _, _ = t, pk, kvs
		if err := h.sink.HandleCreate(context.Background(), t.Name, pk, kvs); err != nil {
			return err
		}
	}
	return nil
}

// TODO(daniel): remove duplication
// nolint: dupl
func (h *CanalEventHandler) handleUpdate(rows [][]interface{}, t *schema.Table) error {
	for _, row := range rows {
		pk, err := getPrimaryKeyValue(row, t)
		if err != nil {
			return err
		}
		kvs, err := getKeyValues(row, t.Columns)
		if err != nil {
			return fmt.Errorf("table %s error: %s", t, err)
		}
		_, _, _ = t, pk, kvs
		if err := h.sink.HandleUpdate(context.Background(), t.Name, pk, kvs); err != nil {
			return err
		}
	}
	return nil
}

func (h *CanalEventHandler) handleDelete(rows [][]interface{}, t *schema.Table) error {
	for _, row := range rows {
		pk, err := getPrimaryKeyValue(row, t)
		if err != nil {
			return err
		}

		_, _ = t, pk
		if err := h.sink.HandleDelete(context.Background(), t.Name, pk); err != nil {
			return err
		}
	}
	return nil
}

func getPrimaryKeyValue(row []interface{}, t *schema.Table) (string, error) {
	v, err := canal.GetPKValues(t, row)
	if err != nil {
		return "", err
	}
	key, err := primaryKeyToString(v)
	if err != nil {
		return "", fmt.Errorf("table %s: %v", t.Name, err)
	}
	return key, nil
}

func primaryKeyToString(keyValues rowKey) (string, error) {
	if len(keyValues) == 0 {
		return "", errors.New("no key values provided")
	}
	if len(keyValues) > 1 {
		return "", errors.New("multi-column primary key is not supported")
	}
	pk := keyValues[0]
	if pk == nil || pk == "" {
		return "", errors.New("primary key value is nil or empty")
	}

	return fmt.Sprint(pk), nil
}

// getKeyValues uses the fact that columns are in the same order as values in a row
func getKeyValues(row []interface{}, columns []schema.TableColumn) (rowData, error) {
	kvs := make(rowData, len(columns))
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
