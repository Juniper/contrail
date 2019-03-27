package replication

import (
	"context"
	"fmt"

	"github.com/pkg/errors"
	"github.com/siddontang/go-mysql/canal"
	"github.com/siddontang/go-mysql/mysql"
	"github.com/siddontang/go-mysql/replication"
	"github.com/siddontang/go-mysql/schema"
	"github.com/sirupsen/logrus"

	"github.com/Juniper/contrail/pkg/logutil"
	"github.com/Juniper/contrail/pkg/services"
)

// CanalHandler handles canal events by decoding them as events and passing them to processor.
type CanalHandler struct {
	decoder   EventDecoder
	processor services.EventProcessor
	log       *logrus.Entry
}

// NewCanalHandler creates new CanalHandler with given decoder and processor.
func NewCanalHandler(p services.EventProcessor, d EventDecoder) *CanalHandler {
	return &CanalHandler{
		decoder:   d,
		processor: p,
		log:       logutil.NewLogger("canal-event-handler"),
	}
}

// OnRow handles row changed modifying key-values in etcd based on action.
func (h *CanalHandler) OnRow(e *canal.RowsEvent) error {
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

func (h *CanalHandler) handleOperation(
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

func (h *CanalHandler) handleCreate(ctx context.Context, rows [][]interface{}, t *schema.Table) error {
	return errors.Wrap(h.handleOperation(ctx, rows, t, services.OperationCreate), "operation CREATE")
}

func (h *CanalHandler) handleUpdate(ctx context.Context, rows [][]interface{}, t *schema.Table) error {
	return errors.Wrap(h.handleOperation(ctx, rows, t, services.OperationUpdate), "operation UPDATE")
}

func (h *CanalHandler) handleDelete(ctx context.Context, rows [][]interface{}, t *schema.Table) error {
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
func (h *CanalHandler) OnRotate(*replication.RotateEvent) error {
	h.log.Debug("Skipping OnRotate event")
	return nil
}

// OnDDL event is skipped.
func (h *CanalHandler) OnDDL(mysql.Position, *replication.QueryEvent) error {
	h.log.Debug("Skipping OnDDL event")
	return nil
}

// OnXID event is skipped.
func (h *CanalHandler) OnXID(mysql.Position) error {
	h.log.Debug("Skipping OnXID event")
	return nil
}

// OnGTID event is skipped.
func (h *CanalHandler) OnGTID(mysql.GTIDSet) error {
	h.log.Debug("Skipping OnGTID event")
	return nil
}

// OnPosSynced event is skipped.
func (h *CanalHandler) OnPosSynced(mysql.Position, bool) error {
	h.log.Debug("Skipping OnPosSynced event")
	return nil
}

// String adds support for stringer interface.
func (h *CanalHandler) String() string {
	return "canalEventHandler"
}
