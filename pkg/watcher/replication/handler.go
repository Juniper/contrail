package replication

import (
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

// Sink represents service that handler transfers data to.
type Sink interface {
	Create(resourceName string, pk string, properties map[string]interface{}) error
	Update(resourceName string, pk string, properties map[string]interface{}) error
	Delete(resourceName string, pk string) error
}

type relationAddGetter interface {
	Add(pgoutput.Relation)
	Get(id uint32) (pgoutput.Relation, error)
}

// PgoutputEventHandler handles replication messages by pushing it to sink.
type PgoutputEventHandler struct {
	sink Sink
	log  *logrus.Entry

	relations relationAddGetter
}

// NewPgoutputEventHandler creates new ReplicationEventHandler using sink provided as an argument.
func NewPgoutputEventHandler(s Sink) *PgoutputEventHandler {
	if s == nil {
		s = &noopSink{}
	}

	return &PgoutputEventHandler{
		sink:      s,
		log:       pkglog.NewLogger("replication-event-handler"),
		relations: &relationSet{},
	}
}

// Handle handles provided message by passing its contents to sink or returns an error
func (h *PgoutputEventHandler) Handle(msg pgoutput.Message) error {

	switch v := msg.(type) {
	case pgoutput.Relation:
		h.log.Debug("received RELATION message")
		h.relations.Add(v)
	case pgoutput.Insert:
		h.log.Debug("received INSERT message")
		return h.handleCreate(v.RelationID, v.Row)
	case pgoutput.Update:
		h.log.Debug("received UPDATE message")
		return h.handleUpdate(v.RelationID, v.Row)
	case pgoutput.Delete:
		h.log.Debug("received DELETE message")
		return h.handleDelete(v.RelationID, v.Row)
	}
	return nil
}

func (h *PgoutputEventHandler) handleCreate(relationID uint32, row []pgoutput.Tuple) error {
	relation, err := h.relations.Get(relationID)
	if err != nil {
		return err
	}

	pk, data, err := decodeRowData(relation, row)
	if err != nil {
		return fmt.Errorf("error decoding row: %v", err)
	}

	return h.sink.Create(relation.Name, pk, data)
}

func (h *PgoutputEventHandler) handleUpdate(relationID uint32, row []pgoutput.Tuple) error {
	relation, err := h.relations.Get(relationID)
	if err != nil {
		return err
	}

	pk, data, err := decodeRowData(relation, row)
	if err != nil {
		return fmt.Errorf("error decoding row: %v", err)
	}

	return h.sink.Update(relation.Name, pk, data)
}

func (h *PgoutputEventHandler) handleDelete(relationID uint32, row []pgoutput.Tuple) error {
	relation, err := h.relations.Get(relationID)
	if err != nil {
		return err
	}

	pk, _, err := decodeRowData(relation, row)
	if err != nil {
		return fmt.Errorf("error decoding row: %v", err)
	}

	return h.sink.Delete(relation.Name, pk)
}

func decodeRowData(relation pgoutput.Relation, row []pgoutput.Tuple) (string, map[string]interface{}, error) {
	keys, data := []interface{}{}, map[string]interface{}{}

	for i, tuple := range row {
		col := relation.Columns[i]
		decoder := col.Decoder()
		if err := decoder.DecodeText(nil, tuple.Value); err != nil {
			return "", nil, fmt.Errorf("error decoding tuple %d: %s", i, err)
		}
		value := decoder.Get()
		data[col.Name] = value
		if col.Key {
			keys = append(keys, value)
		}

	}

	pk, err := primaryKeyToString(keys)
	if err != nil {
		return "", nil, fmt.Errorf("error creating PK: %v", err)
	}

	return pk, data, nil
}

// CanalEventHandler handles canal events by pushing it to sink.
type CanalEventHandler struct {
	sink Sink
	log  *logrus.Entry
}

// NewCanalEventHandler creates new CanalEventHandler with given sink.
func NewCanalEventHandler(s Sink) *CanalEventHandler {
	if s == nil {
		s = &noopSink{}
	}

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
		if err := h.sink.Create(t.Name, pk, kvs); err != nil {
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
		if err := h.sink.Update(t.Name, pk, kvs); err != nil {
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

		if err := h.sink.Delete(t.Name, pk); err != nil {
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

func primaryKeyToString(keyValues []interface{}) (string, error) {
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
	return "eventHandler"
}

type noopSink struct{}

func (s *noopSink) Create(resourceName string, pk string, properties map[string]interface{}) error {
	return nil
}

func (s *noopSink) Update(resourceName string, pk string, properties map[string]interface{}) error {
	return nil
}

func (s *noopSink) Delete(resourceName string, pk string) error { return nil }
