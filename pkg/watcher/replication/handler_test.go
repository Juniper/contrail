package replication

import (
	"fmt"
	"testing"

	"github.com/siddontang/go-mysql/canal"
	"github.com/siddontang/go-mysql/mysql"
	"github.com/siddontang/go-mysql/replication"
	"github.com/siddontang/go-mysql/schema"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestEventHandlerIsNoopByDefault(t *testing.T) {
	for _, action := range []string{canal.InsertAction, canal.UpdateAction, canal.DeleteAction} {
		t.Run(action, func(t *testing.T) {
			h := givenEventHandler(nil)
			err := h.OnRow(givenRowsEvent(action))
			assert.NoError(t, err)
		})
	}
}

func TestOnRowFailsWhenInvalidActionGiven(t *testing.T) {
	h := givenEventHandler(&sinkMock{})
	err := h.OnRow(givenRowsEvent("invalid-action"))
	assert.Error(t, err)
}

func TestOnRowFailsWhenInvalidTablePrimaryKeyGiven(t *testing.T) {
	for _, action := range []string{canal.InsertAction, canal.UpdateAction, canal.DeleteAction} {
		t.Run(action, func(t *testing.T) {
			h := givenEventHandler(&sinkMock{})
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
			h := givenEventHandler(&sinkMock{})
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
			h := givenEventHandler(&sinkMock{})
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
			h := givenEventHandler(&sinkMock{})
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

func TestOnRowFailsWhenSinkCreateFails(t *testing.T) {
	m := &sinkMock{}
	m.On("Create", mock.Anything, mock.Anything, mock.Anything).Return(assert.AnError).Once()
	h := givenEventHandler(m)
	e := givenRowsEvent(canal.InsertAction)

	err := h.OnRow(e)

	assert.Error(t, err)
	m.AssertExpectations(t)
}

// TODO(daniel): remove duplication
// nolint: dupl
func TestOnRowCreatesResourceInSinkForEveryRowInInsertEvent(t *testing.T) {
	m := &sinkMock{}
	m.On("Create", "test-resource", "foo", map[string]interface{}{
		"string-property": "foo",
		"int-property":    1337,
		"float-property":  1.337,
	}).Return(nil).Once()
	m.On("Create", "test-resource", "bar", map[string]interface{}{
		"string-property": "bar",
		"int-property":    0,
		"float-property":  0.1,
	}).Return(nil).Once()
	m.On("Create", "test-resource", "baz", map[string]interface{}{
		"string-property": "baz",
		"int-property":    -1337,
		"float-property":  -1.337,
	}).Return(nil).Once()
	h := givenEventHandler(m)
	e := givenRowsEvent(canal.InsertAction)
	e.Rows = [][]interface{}{
		{"foo", 1337, 1.337},
		{"bar", 0, 0.1},
		{"baz", -1337, -1.337},
	}

	err := h.OnRow(e)

	assert.NoError(t, err)
	m.AssertExpectations(t)
}

func TestOnRowFailsWhenSinkUpdateFails(t *testing.T) {
	m := &sinkMock{}
	m.On("Update", mock.Anything, mock.Anything, mock.Anything).Return(assert.AnError).Once()
	h := givenEventHandler(m)
	e := givenRowsEvent(canal.UpdateAction)

	err := h.OnRow(e)

	assert.Error(t, err)
	m.AssertExpectations(t)
}

// TODO(daniel): remove duplication
// nolint: dupl
func TestOnRowUpdatesResourceInSinkForEveryRowInUpdateEvent(t *testing.T) {
	m := &sinkMock{}
	m.On("Update", "test-resource", "foo", map[string]interface{}{
		"string-property": "foo",
		"int-property":    1337,
		"float-property":  1.337,
	}).Return(nil).Once()
	m.On("Update", "test-resource", "bar", map[string]interface{}{
		"string-property": "bar",
		"int-property":    0,
		"float-property":  0.1,
	}).Return(nil).Once()
	m.On("Update", "test-resource", "baz", map[string]interface{}{
		"string-property": "baz",
		"int-property":    -1337,
		"float-property":  -1.337,
	}).Return(nil).Once()
	h := givenEventHandler(m)
	e := givenRowsEvent(canal.UpdateAction)
	e.Rows = [][]interface{}{
		{"foo", 1337, 1.337},
		{"bar", 0, 0.1},
		{"baz", -1337, -1.337},
	}

	err := h.OnRow(e)

	assert.NoError(t, err)
	m.AssertExpectations(t)
}

func TestOnRowFailsWhenSinkDeleteFails(t *testing.T) {
	m := &sinkMock{}
	m.On("Delete", mock.Anything, mock.Anything, mock.Anything).Return(assert.AnError).Once()
	h := givenEventHandler(m)
	e := givenRowsEvent(canal.DeleteAction)

	err := h.OnRow(e)

	assert.Error(t, err)
	m.AssertExpectations(t)
}

func TestOnRowDeletesResourcesInSinkOnDeleteEvent(t *testing.T) {
	m := &sinkMock{}
	m.On("Delete", "test-resource", "foo").Return(nil).Once()
	m.On("Delete", "test-resource", "bar").Return(nil).Once()
	m.On("Delete", "test-resource", "baz").Return(nil).Once()
	h := givenEventHandler(m)
	e := givenRowsEvent(canal.DeleteAction)
	e.Rows = [][]interface{}{
		{"foo", 1337, 1.337},
		{"bar", 0, 0.1},
		{"baz", -1337, -1.337},
	}

	err := h.OnRow(e)

	assert.NoError(t, err)
	m.AssertExpectations(t)
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
	h := givenEventHandler(nil)
	err := h.OnRotate(&replication.RotateEvent{})
	assert.NoError(t, err)
}

func TestOnDDLEventIsSkipped(t *testing.T) {
	h := givenEventHandler(nil)
	err := h.OnDDL(mysql.Position{}, &replication.QueryEvent{})
	assert.NoError(t, err)
}

func TestOnXIDEventIsSkipped(t *testing.T) {
	h := givenEventHandler(nil)
	err := h.OnXID(mysql.Position{})
	assert.NoError(t, err)
}

func TestOnGTIDEventIsSkipped(t *testing.T) {
	h := givenEventHandler(nil)
	err := h.OnGTID(&mysql.MysqlGTIDSet{})
	assert.NoError(t, err)
}

func TestOnPosSyncedEventIsSkipped(t *testing.T) {
	h := givenEventHandler(nil)
	err := h.OnPosSynced(mysql.Position{}, false)
	assert.NoError(t, err)
}

func TestStringerReturnsHandlerName(t *testing.T) {
	h := givenEventHandler(nil)
	assert.Equal(t, "eventHandler", h.String())
}

func givenEventHandler(s Sink) *CanalEventHandler {
	return NewCanalEventHandler(s)
}

type sinkMock struct {
	mock.Mock
}

func (m *sinkMock) Create(resourceName string, pk string, properties map[string]interface{}) error {
	args := m.Called(resourceName, pk, properties)
	return args.Error(0)
}

func (m *sinkMock) Update(resourceName string, pk string, properties map[string]interface{}) error {
	args := m.Called(resourceName, pk, properties)
	return args.Error(0)
}

func (m *sinkMock) Delete(resourceName string, pk string) error {
	args := m.Called(resourceName, pk)
	return args.Error(0)
}
