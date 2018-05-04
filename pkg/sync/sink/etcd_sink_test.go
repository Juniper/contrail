package sink

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/Juniper/contrail/pkg/log"
	"github.com/Juniper/contrail/pkg/models"
	"github.com/coreos/etcd/clientv3"
	conc "github.com/coreos/etcd/clientv3/concurrency"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

var c = &JSONCodec{}

func TestETCDSinkCreateFailsWhenKVClientFails(t *testing.T) {
	m := givenKVClientMock()
	m.onPut("json/test-resource/test-pk", mock.AnythingOfType("string")).Return(nil, assert.AnError).Once()
	s := givenETCDSink(m, nil, c)

	err := s.Create("test-resource", "test-pk", models.MakeLogicalInterface())

	assert.Error(t, err)
	m.AssertExpectations(t)
}

func TestETCDSinkCreateFailsWhenPropertiesEncodingFails(t *testing.T) {
	s := givenETCDSink(nil, nil, c)

	err := s.Create("test-resource", "test-pk", &errorObject{})

	assert.Error(t, err)
}

func TestETCDSinkCreatePutsJSONEncodedResourceUnderOneEtcdKey(t *testing.T) {
	resource := models.MakeLogicalInterface()
	m := givenKVClientMock()
	m.onPut("json/test-resource/test-pk", string(toJSON(t, resource))).Return(&clientv3.PutResponse{}, nil).Once()
	stm := &mockSTM{}
	s := givenETCDSink(m, stm, c)

	err := s.Create("test-resource", "test-pk", resource)

	assert.NoError(t, err)
	m.AssertExpectations(t)
}

func TestETCDSinkUpdateFailsWhenKVClientFails(t *testing.T) {
	stm := &mockSTM{}
	s := givenETCDSink(nil, stm, c)
	s.inTransaction = func(ctx context.Context, apply func(conc.STM) error) error {
		return assert.AnError
	}

	err := s.Update("test-resource", "test-pk", models.MakeLogicalInterface())

	assert.Error(t, err)
}

func TestETCDSinkUpdatePutsJSONEncodedResourceUnderOneEtcdKey(t *testing.T) {
	resource := models.MakeLogicalInterface()
	stm := &mockSTM{}
	stm.On("Get", []string{"json/test-resource/pk"}).Return("{}").Once()
	stm.On("Put", "json/test-resource/pk", string(toJSON(t, resource)), mock.Anything).Return(&clientv3.PutResponse{}, nil).Once()
	s := givenETCDSink(nil, stm, c)

	err := s.Update("test-resource", "pk", resource)

	assert.NoError(t, err)
	stm.AssertExpectations(t)
}

func TestETCDSinkDeleteFailsWhenKVClientFails(t *testing.T) {
	m := givenKVClientMock()
	m.onDelete("json/test-resource/test-pk").Return(nil, assert.AnError).Once()
	s := givenETCDSink(m, nil, c)

	err := s.Delete("test-resource", "test-pk")

	assert.Error(t, err)
	m.AssertExpectations(t)
}

func TestETCDSinkDeleteRemovesEtcdKeyWithGivenResource(t *testing.T) {
	m := givenKVClientMock()
	m.onDelete("json/test-resource/pk").Return(&clientv3.DeleteResponse{}, nil).Once()
	s := givenETCDSink(m, nil, c)

	err := s.Delete("test-resource", "pk")

	assert.NoError(t, err)
	m.AssertExpectations(t)
}

func givenETCDSink(kv clientv3.KV, stm conc.STM, codec Codec) *ETCDSink {
	return &ETCDSink{
		kvClient: kv,
		inTransaction: func(ctx context.Context, apply func(conc.STM) error) error {
			return apply(stm)
		},
		codec: codec,
		log:   log.NewLogger("etcd-sink"),
	}
}

func toJSON(t *testing.T, data interface{}) []byte {
	d, err := json.Marshal(data)
	require.NoError(t, err)
	return d
}

type mockSTM struct {
	mock.Mock
	conc.STM
}

func (m *mockSTM) Get(key ...string) string {
	args := m.MethodCalled("Get", key)
	return args.String(0)
}

func (m *mockSTM) Put(key string, val string, opts ...clientv3.OpOption) {
	_ = m.MethodCalled("Put", key, val, opts)
}

func (m *mockSTM) Rev(key string) int64 {
	args := m.MethodCalled("Rev", key)
	val, _ := args.Get(0).(int64)
	return val
}

func (m *mockSTM) Del(key string) {
	_ = m.MethodCalled("Del", key)
}

type errorObject struct{}

func (e *errorObject) Reset()                       {}
func (e *errorObject) String() string               { return "" }
func (e *errorObject) ProtoMessage()                {}
func (e *errorObject) UnmarshalJSON([]byte) error   { return assert.AnError }
func (e *errorObject) MarshalJSON() ([]byte, error) { return nil, assert.AnError }
