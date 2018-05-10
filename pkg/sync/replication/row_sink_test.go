package replication

import (
	"testing"

	"github.com/gogo/protobuf/proto"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestObjectMappingAdapterCreate(t *testing.T) {
	resourceName, pk, props := "resource", "1", map[string]interface{}{}
	message := &dummyMessage{}

	tests := []struct {
		name           string
		initRowScanner func(o oner)
		initSink       func(o oner)
		fails          bool
	}{
		{name: "scanner fails", initRowScanner: func(o oner) {
			o.On("ScanRow", resourceName, props).Return(nil, assert.AnError).Once()
		}, fails: true},
		{
			name: "sink fails",
			initRowScanner: func(o oner) {
				o.On("ScanRow", resourceName, props).Return(message, nil).Once()
			},
			initSink: func(o oner) {
				o.On("Create", resourceName, pk, message).Return(assert.AnError).Once()
			},
			fails: true,
		},
		{
			name: "correct message",
			initRowScanner: func(o oner) {
				o.On("ScanRow", resourceName, props).Return(message, nil).Once()
			},
			initSink: func(o oner) {
				o.On("Create", resourceName, pk, message).Return(nil).Once()
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

			err := a.Create(resourceName, pk, props)

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

func (d *dummyMessage) Reset() {}

func (d *dummyMessage) String() string {
	return "dummy"
}

func (d *dummyMessage) ProtoMessage() {}

type rowScannerMock mock.Mock

func (m *rowScannerMock) ScanRow(schemaID string, rowData map[string]interface{}) (proto.Message, error) {
	args := (*mock.Mock)(m).MethodCalled("ScanRow", schemaID, rowData)
	if args.Error(1) != nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(proto.Message), nil
}

type sinkMock struct {
	mock.Mock
}

func (s *sinkMock) Create(resourceName string, pk string, obj interface{}) error {
	args := s.MethodCalled("Create", resourceName, pk, obj)
	return args.Error(0)
}

func (s *sinkMock) Update(resourceName string, pk string, obj interface{}) error {
	args := s.MethodCalled("Update", resourceName, pk, obj)
	return args.Error(0)
}

func (s *sinkMock) Delete(resourceName string, pk string) error {
	args := s.MethodCalled("Delete", resourceName, pk)
	return args.Error(0)
}
