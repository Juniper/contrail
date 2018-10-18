package replication

import (
	"context"
	"testing"

	"github.com/Juniper/contrail/pkg/models/basemodels"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestObjectMappingAdapterCreate(t *testing.T) {
	resourceName, pk, props := "resource", []string{"1"}, map[string]interface{}{}
	message := &dummyMessage{}

	tests := []struct {
		name           string
		initRowScanner func(o oner)
		initSink       func(o oner)
		fails          bool
	}{
		{
			name: "scanner fails",
			initRowScanner: func(o oner) {
				o.On("ScanRow", resourceName, props).Return(nil, assert.AnError).Once()
			},
			fails: true},
		{
			name: "sink fails",
			initRowScanner: func(o oner) {
				o.On("ScanRow", resourceName, props).Return(message, nil).Once()
			},
			initSink: func(o oner) {
				o.On("Create", resourceName, pk[0], message).Return(assert.AnError).Once()
			},
			fails: true,
		},
		{
			name: "correct message",
			initRowScanner: func(o oner) {
				o.On("ScanRow", resourceName, props).Return(message, nil).Once()
			},
			initSink: func(o oner) {
				o.On("Create", resourceName, pk[0], message).Return(nil).Once()
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sMock, rsMock := &sinkMock{}, &mock.Mock{}
			if tt.initRowScanner != nil {
				tt.initRowScanner(rsMock)
			}
			if tt.initSink != nil {
				tt.initSink(sMock)
			}

			a := NewObjectMappingAdapter(sMock, (*rowScannerMock)(rsMock))

			err := a.Create(context.Background(), resourceName, pk, props)

			if tt.fails {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			sMock.AssertExpectations(t)
			rsMock.AssertExpectations(t)
		})
	}
}

func TestObjectMappingAdapterRefCreate(t *testing.T) {
	resourceName, correctPK, props := "ref_resource", []string{"1", "2"}, map[string]interface{}{}

	sMock, rsMock := &sinkMock{}, &mock.Mock{}
	adapter := NewObjectMappingAdapter(sMock, (*rowScannerMock)(rsMock))

	tests := []struct {
		name           string
		initRowScanner func(o oner)
		initSink       func(o oner)
		fails          bool
		operationFunc  func(context.Context, string, []string, map[string]interface{}) error
		pk             []string
	}{
		{
			name:          "missing multiple primary key in ref_ message",
			fails:         true,
			pk:            []string{"1"},
			operationFunc: adapter.Create,
			initRowScanner: func(o oner) {
			},
		},
		{
			name:          "update for refs is not handled",
			fails:         true,
			pk:            correctPK,
			operationFunc: adapter.Update,
		},
		{
			name:          "correct ref_ message",
			pk:            correctPK,
			operationFunc: adapter.Create,
			initRowScanner: func(o oner) {
			},
			initSink: func(o oner) {
				o.On("CreateRef", resourceName, correctPK, map[string]interface{}(nil)).Return(nil).Once()
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.initRowScanner != nil {
				tt.initRowScanner(rsMock)
			}
			if tt.initSink != nil {
				tt.initSink(sMock)
			}

			err := tt.operationFunc(context.Background(), resourceName, tt.pk, props)

			if tt.fails {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			sMock.AssertExpectations(t)
			rsMock.AssertExpectations(t)
		})
	}
}

type dummyMessage struct{}

func (d *dummyMessage) Reset()                                 {}
func (d *dummyMessage) String() string                         { return "dummy" }
func (d *dummyMessage) ProtoMessage()                          {}
func (d *dummyMessage) ToMap() map[string]interface{}          { return nil }
func (d *dummyMessage) Kind() string                           { return "" }
func (d *dummyMessage) GetParentUUID() string                  { return "" }
func (d *dummyMessage) GetUUID() string                        { return "" }
func (d *dummyMessage) GetReferences() []basemodels.Reference  { return nil }
func (d *dummyMessage) GetBackReferences() []basemodels.Object { return nil }
func (d *dummyMessage) GetChildren() []basemodels.Object       { return nil }
func (d *dummyMessage) AddBackReference(interface{})           {}
func (d *dummyMessage) AddChild(interface{})                   {}
func (d *dummyMessage) RemoveBackReference(interface{})        {}
func (d *dummyMessage) RemoveChild(interface{})                {}
func (d *dummyMessage) GetFQName() []string                    { return nil }
func (d *dummyMessage) TypeName() string                       { return "" }

func (d *dummyMessage) ApplyPropCollectionUpdate(
	*basemodels.PropCollectionUpdate,
) (updated map[string]interface{}, err error) {
	return nil, nil
}

func (d *dummyMessage) ApplyMap(_ map[string]interface{}) {}

type rowScannerMock mock.Mock

func (m *rowScannerMock) ScanRow(schemaID string, rowData map[string]interface{}) (basemodels.Object, error) {
	args := (*mock.Mock)(m).MethodCalled("ScanRow", schemaID, rowData)
	if args.Error(1) != nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(basemodels.Object), nil
}

type sinkMock struct {
	mock.Mock
}

func (s *sinkMock) Create(ctx context.Context, resourceName string, pk string, obj basemodels.Object) error {
	args := s.MethodCalled("Create", resourceName, pk, obj)
	return args.Error(0)
}

func (s *sinkMock) Update(ctx context.Context, resourceName string, pk string, obj basemodels.Object) error {
	args := s.MethodCalled("Update", resourceName, pk, obj)
	return args.Error(0)
}

func (s *sinkMock) Delete(ctx context.Context, resourceName string, pk string) error {
	args := s.MethodCalled("Delete", resourceName, pk)
	return args.Error(0)
}

func (s *sinkMock) CreateRef(ctx context.Context, resourceName string, pk []string, attr map[string]interface{}) error {
	args := s.MethodCalled("CreateRef", resourceName, pk, attr)
	return args.Error(0)
}

func (s *sinkMock) DeleteRef(ctx context.Context, resourceName string, pk []string) error {
	args := s.MethodCalled("DeleteRef", resourceName, pk)
	return args.Error(0)
}
