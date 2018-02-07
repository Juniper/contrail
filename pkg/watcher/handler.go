package watcher

import (
	"errors"
	"fmt"

	pkglog "github.com/Juniper/contrail/pkg/log"
	"github.com/siddontang/go-mysql/canal"
	"github.com/siddontang/go-mysql/mysql"
	"github.com/siddontang/go-mysql/replication"
	"github.com/siddontang/go-mysql/schema"
	"github.com/sirupsen/logrus"
)

// sink represents service that handler transfers data to.
type sink interface {
	Create(resourceName string, pk string, properties map[string]interface{}) error
	Update(resourceName string, pk string, properties map[string]interface{}) error
	Delete(resourceName string, pk string) error
}

type eventHandler struct {
	sink sink
	log  *logrus.Entry
}

func newHandler(s sink) *eventHandler {
	if s == nil {
		s = &noopSink{}
	}

	return &eventHandler{
		sink: s,
		log:  pkglog.NewLogger("event-handler"),
	}
}

// OnRow handles row changed modifying key-values in etcd based on action.
func (h *eventHandler) OnRow(e *canal.RowsEvent) error {
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
func (h *eventHandler) handleCreate(rows [][]interface{}, t *schema.Table) error {
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
func (h *eventHandler) handleUpdate(rows [][]interface{}, t *schema.Table) error {
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

func (h *eventHandler) handleDelete(rows [][]interface{}, t *schema.Table) error {
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

	if len(v) != 1 {
		return "", fmt.Errorf("table %s has multi-column primary key which is not supported", t.Name)
	}
	pk := v[0]
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
func (h *eventHandler) OnRotate(*replication.RotateEvent) error {
	h.log.Debug("Skipping OnRotate event")
	return nil
}

// OnDDL event is skipped.
func (h *eventHandler) OnDDL(mysql.Position, *replication.QueryEvent) error {
	h.log.Debug("Skipping OnDDL event")
	return nil
}

// OnXID event is skipped.
func (h *eventHandler) OnXID(mysql.Position) error {
	h.log.Debug("Skipping OnXID event")
	return nil
}

// OnGTID event is skipped.
func (h *eventHandler) OnGTID(mysql.GTIDSet) error {
	h.log.Debug("Skipping OnGTID event")
	return nil
}

// OnPosSynced event is skipped.
func (h *eventHandler) OnPosSynced(mysql.Position, bool) error {
	h.log.Debug("Skipping OnPosSynced event")
	return nil
}

// String adds support for stringer interface.
func (h *eventHandler) String() string {
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
