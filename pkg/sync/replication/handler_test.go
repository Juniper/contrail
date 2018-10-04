package replication

import (
	"context"
	"fmt"
	"testing"

	"github.com/jackc/pgx/pgtype"
	"github.com/kyleconroy/pgoutput"
	"github.com/siddontang/go-mysql/canal"
	"github.com/siddontang/go-mysql/mysql"
	"github.com/siddontang/go-mysql/replication"
	"github.com/siddontang/go-mysql/schema"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestPgoutputEventHandlerHandle(t *testing.T) {
	exampleRelation := pgoutput.Relation{
		Name: "test-resource",
		Columns: []pgoutput.Column{
			{Name: "string-property", Key: true, Type: pgtype.VarcharOID},
			{Name: "int-property", Type: pgtype.Int4OID},
			{Name: "float-property", Type: pgtype.Float8OID},
		},
	}

	exampleRow := []pgoutput.Tuple{
		{Value: []byte(`foo`)},
		{Value: []byte(`1337`)},
		{Value: []byte(`1.337`)},
	}

	exampleRowData := map[string]interface{}{
		"string-property": "foo",
		"int-property":    int32(1337),
		"float-property":  1.337,
	}

	tests := []struct {
		name         string
		initMock     func(oner)
		initialRels  relationAddGetter
		message      pgoutput.Message
		fails        bool
		expectedRels relationAddGetter
	}{
		{name: "nil message", message: nil},
		{name: "insert unknown relation", message: pgoutput.Insert{}, fails: true},
		{name: "update unknown relation", message: pgoutput.Update{}, fails: true},
		{name: "delete unknown relation", message: pgoutput.Delete{}, fails: true},
		{name: "insert malformed relation", message: pgoutput.Insert{RelationID: 1}, fails: true},
		{name: "update malformed relation", message: pgoutput.Update{RelationID: 1}, fails: true},
		{name: "delete malformed relation", message: pgoutput.Delete{RelationID: 1}, fails: true},
		{
			name:        "insert no primary key",
			initialRels: &relationSet{1: pgoutput.Relation{Name: "rel"}},
			message:     pgoutput.Insert{RelationID: 1},
			fails:       true,
		},
		{
			name:        "update no primary key",
			initialRels: &relationSet{1: pgoutput.Relation{Name: "rel"}},
			message:     pgoutput.Update{RelationID: 1},
			fails:       true,
		},
		{
			name:        "delete no primary key",
			initialRels: &relationSet{1: pgoutput.Relation{Name: "rel"}},
			message:     pgoutput.Delete{RelationID: 1},
			fails:       true,
		},
		{
			name:         "new relation",
			message:      pgoutput.Relation{ID: 1337},
			expectedRels: &relationSet{1337: pgoutput.Relation{ID: 1337}},
		},
		{
			name:         "already stored relation",
			initialRels:  &relationSet{1337: pgoutput.Relation{Name: "old"}},
			message:      pgoutput.Relation{ID: 1337, Name: "new"},
			expectedRels: &relationSet{1337: pgoutput.Relation{ID: 1337, Name: "new"}},
		},
		{
			name: "correct insert message",
			initMock: func(m oner) {
				m.On("Create", "test-resource", []string{"foo"}, exampleRowData).Return(nil).Once()
			},
			initialRels: &relationSet{1: exampleRelation},
			message:     pgoutput.Insert{RelationID: 1, Row: exampleRow},
		},
		{
			name: "correct update message",
			initMock: func(m oner) {
				m.On("Update", "test-resource", []string{"foo"}, exampleRowData).Return(nil).Once()
			},
			initialRels: &relationSet{1: exampleRelation},
			message:     pgoutput.Update{RelationID: 1, Row: exampleRow},
		},
		{
			name: "correct delete message",
			initMock: func(m oner) {
				m.On("Delete", "test-resource", []string{"foo"}).Return(nil).Once()
			},
			initialRels: &relationSet{1: exampleRelation},
			message:     pgoutput.Delete{RelationID: 1, Row: exampleRow},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// given
			m := newRowSinkMock()
			if tt.initMock != nil {
				tt.initMock(m)
			}

			h := NewPgoutputEventHandler(m)
			if tt.initialRels != nil {
				h.relations = tt.initialRels
			}

			// when
			err := h.Handle(context.Background(), tt.message)

			// then
			if tt.fails {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}

			if tt.expectedRels != nil {
				assert.Equal(t, tt.expectedRels, h.relations)
			}

			m.AssertExpectations(t)
		})
	}
}

func TestOnRowFailsWhenInvalidActionGiven(t *testing.T) {
	h := NewCanalEventHandler(&rowSinkMock{})
	err := h.OnRow(givenRowsEvent("invalid-action"))
	assert.Error(t, err)
}

func TestOnRowFailsWhenInvalidTablePrimaryKeyGiven(t *testing.T) {
	for _, action := range []string{canal.InsertAction, canal.UpdateAction, canal.DeleteAction} {
		t.Run(action, func(t *testing.T) {
			h := NewCanalEventHandler(&rowSinkMock{})
			e := givenRowsEvent(action)
			e.Table.PKColumns = []int{}

			err := h.OnRow(e)

			assert.Error(t, err)
		})
	}
}

func TestOnRowFailsWhenTableWithMultiColumnPrimaryKeyGiven(t *testing.T) {
	for _, action := range []string{canal.InsertAction, canal.UpdateAction, canal.DeleteAction} {
		t.Run(action, func(t *testing.T) {
			h := NewCanalEventHandler(&rowSinkMock{})
			e := givenRowsEvent(action)
			e.Table.PKColumns = []int{0, 1}

			err := h.OnRow(e)

			assert.Error(t, err)
		})
	}
}

func TestOnRowFailsWhenEmptyPrimaryKeyValueGiven(t *testing.T) {
	for _, action := range []string{canal.InsertAction, canal.UpdateAction, canal.DeleteAction} {
		t.Run(action, func(t *testing.T) {
			h := NewCanalEventHandler(&rowSinkMock{})
			e := givenRowsEvent(action)
			e.Rows = [][]interface{}{{"", 1337, 1.337}}

			err := h.OnRow(e)

			assert.Error(t, err)
		})
	}
}

func TestOnRowFailsWhenInvalidTableColumnTypeGiven(t *testing.T) {
	tests := []struct {
		action     string
		columnType int
	}{
		{canal.InsertAction, schema.TYPE_BIT},
		{canal.InsertAction, schema.TYPE_ENUM},
		{canal.InsertAction, schema.TYPE_SET},
		{canal.UpdateAction, schema.TYPE_BIT},
		{canal.UpdateAction, schema.TYPE_ENUM},
		{canal.UpdateAction, schema.TYPE_SET},
	}
	for _, test := range tests {
		t.Run(fmt.Sprint(test.action, test.columnType), func(t *testing.T) {
			h := NewCanalEventHandler(&rowSinkMock{})
			e := givenRowsEvent(test.action)
			e.Table.Columns = []schema.TableColumn{
				{Name: "string-property", Type: schema.TYPE_STRING},
				{Name: "property", Type: test.columnType},
			}
			e.Rows = [][]interface{}{{"foo", "property-value"}}

			err := h.OnRow(e)

			assert.Error(t, err)
		})
	}
}

func TestOnRow(t *testing.T) {
	exampleRows := [][]interface{}{{"foo", 1337, 1.337}, {"bar", 0, 0.1}, {"baz", -1337, -1.337}}
	exampleRowsData := []map[string]interface{}{
		{"string-property": "foo", "int-property": 1337, "float-property": 1.337},
		{"string-property": "bar", "int-property": 0, "float-property": 0.1},
		{"string-property": "baz", "int-property": -1337, "float-property": -1.337},
	}

	tests := []struct {
		name     string
		initMock func(oner)
		action   string
		rows     [][]interface{}
		fails    bool
	}{
		{
			name: "sink create fails",
			initMock: func(m oner) {
				m.On("Create", mock.Anything, mock.AnythingOfType("[]string"), mock.Anything).Return(assert.AnError).Once()
			},
			action: canal.InsertAction,
			fails:  true,
		},
		{
			name: "insert 3 rows correctly",
			initMock: func(m oner) {
				m.On("Create", "test-resource", []string{"foo"}, exampleRowsData[0]).Return(nil).Once()
				m.On("Create", "test-resource", []string{"bar"}, exampleRowsData[1]).Return(nil).Once()
				m.On("Create", "test-resource", []string{"baz"}, exampleRowsData[2]).Return(nil).Once()
			},
			action: canal.InsertAction,
			rows:   exampleRows,
		},
		{
			name: "sink update fails",
			initMock: func(m oner) {
				m.On("Update", mock.Anything, mock.Anything, mock.Anything).Return(assert.AnError).Once()
			},
			action: canal.UpdateAction,
			fails:  true,
		},
		{
			name: "update 3 rows correctly",
			initMock: func(m oner) {
				m.On("Update", "test-resource", []string{"foo"}, exampleRowsData[0]).Return(nil).Once()
				m.On("Update", "test-resource", []string{"bar"}, exampleRowsData[1]).Return(nil).Once()
				m.On("Update", "test-resource", []string{"baz"}, exampleRowsData[2]).Return(nil).Once()
			},
			action: canal.UpdateAction,
			rows:   exampleRows,
		},
		{
			name: "sink delete fails",
			initMock: func(m oner) {
				m.On("Delete", mock.Anything, mock.Anything, mock.Anything).Return(assert.AnError).Once()
			},
			action: canal.DeleteAction,
			fails:  true,
		},
		{
			name: "delete 3 rows correctly",
			initMock: func(m oner) {
				m.On("Delete", "test-resource", []string{"foo"}).Return(nil).Once()
				m.On("Delete", "test-resource", []string{"bar"}).Return(nil).Once()
				m.On("Delete", "test-resource", []string{"baz"}).Return(nil).Once()
			},
			action: canal.DeleteAction,
			rows:   exampleRows,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &rowSinkMock{}
			tt.initMock(m)
			h := NewCanalEventHandler(m)

			e := givenRowsEvent(tt.action)
			if tt.rows != nil {
				e.Rows = tt.rows
			}

			err := h.OnRow(e)

			if tt.fails {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			m.AssertExpectations(t)
		})
	}

}

func givenRowsEvent(action string) *canal.RowsEvent {
	return &canal.RowsEvent{
		Action: action,
		Table: &schema.Table{
			Name: "test-resource",
			Columns: []schema.TableColumn{
				{Name: "string-property", Type: schema.TYPE_STRING},
				{Name: "int-property", Type: schema.TYPE_NUMBER},
				{Name: "float-property", Type: schema.TYPE_FLOAT},
			},
			PKColumns: []int{0},
		},
		Rows: [][]interface{}{{"foo", 1337, 1.337}},
	}
}

func TestOnRotateEventIsSkipped(t *testing.T) {
	h := NewCanalEventHandler(nil)
	err := h.OnRotate(&replication.RotateEvent{})
	assert.NoError(t, err)
}

func TestOnDDLEventIsSkipped(t *testing.T) {
	h := NewCanalEventHandler(nil)
	err := h.OnDDL(mysql.Position{}, &replication.QueryEvent{})
	assert.NoError(t, err)
}

func TestOnXIDEventIsSkipped(t *testing.T) {
	h := NewCanalEventHandler(nil)
	err := h.OnXID(mysql.Position{})
	assert.NoError(t, err)
}

func TestOnGTIDEventIsSkipped(t *testing.T) {
	h := NewCanalEventHandler(nil)
	err := h.OnGTID(&mysql.MysqlGTIDSet{})
	assert.NoError(t, err)
}

func TestOnPosSyncedEventIsSkipped(t *testing.T) {
	h := NewCanalEventHandler(nil)
	err := h.OnPosSynced(mysql.Position{}, false)
	assert.NoError(t, err)
}

func TestStringerReturnsHandlerName(t *testing.T) {
	h := NewCanalEventHandler(nil)
	assert.Equal(t, "canalEventHandler", h.String())
}

type rowSinkMock struct {
	mock.Mock
}

func newRowSinkMock() *rowSinkMock {
	return &rowSinkMock{}
}

func (m *rowSinkMock) Create(
	ctx context.Context, resourceName string, pk []string, properties map[string]interface{},
) error {
	args := m.Called(resourceName, pk, properties)
	return args.Error(0)
}

func (m *rowSinkMock) Update(
	ctx context.Context, resourceName string, pk []string, properties map[string]interface{},
) error {
	args := m.Called(resourceName, pk, properties)
	return args.Error(0)
}

func (m *rowSinkMock) Delete(
	ctx context.Context, resourceName string, pk []string,
) error {
	args := m.Called(resourceName, pk)
	return args.Error(0)
}
