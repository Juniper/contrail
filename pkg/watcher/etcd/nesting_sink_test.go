package etcd

import (
	"bytes"
	"encoding/binary"
	"testing"

	"github.com/coreos/etcd/clientv3"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestNestingSinkIsNoopByDefault(t *testing.T) {
	s := givenNestingSink(nil)

	err := s.Create("test-resource", "test-pk", map[string]interface{}{"property": "foo"})
	assert.NoError(t, err)

	err = s.Update("test-resource", "test-pk", map[string]interface{}{"property": "foo"})
	assert.NoError(t, err)

	err = s.Delete("test-resource", "test-pk")
	assert.NoError(t, err)
}

func TestNestingSinkCreateFailsWhenInvalidPropertyValuePassed(t *testing.T) {
	s := givenNestingSink(givenKVClientMock())
	err := s.Create("test-resource", "test-pk", map[string]interface{}{"property": make(chan struct{})})
	assert.Error(t, err)
}

func TestNestingSinkCreateFailsWhenKVClientFails(t *testing.T) {
	m := givenKVClientMock()
	m.onPut("nested/test-resource/test-pk/property", "foo").Return(nil, assert.AnError).Once()
	s := givenNestingSink(m)

	err := s.Create("test-resource", "test-pk", map[string]interface{}{"property": "foo"})

	assert.Error(t, err)
	m.AssertExpectations(t)
}

// TODO(daniel): remove duplication
// nolint: dupl
func TestNestingSinkCreatePutsEtcdKeysForAllParameters(t *testing.T) {
	m := givenKVClientMock()
	m.onPut("nested/test-resource/test-pk/foo-property", "foo").Return(&clientv3.PutResponse{}, nil).Once()
	m.onPut("nested/test-resource/test-pk/1337-property", mock.Anything).Return(&clientv3.PutResponse{}, nil).Once()
	m.onPut("nested/test-resource/test-pk/1.337-property", mock.Anything).Return(&clientv3.PutResponse{}, nil).Once()
	m.onPut("nested/test-resource/test-pk/true-property", mock.Anything).Return(&clientv3.PutResponse{}, nil).Once()
	s := givenNestingSink(m)

	err := s.Create("test-resource", "test-pk", map[string]interface{}{
		"foo-property":   "foo",
		"1337-property":  1337,
		"1.337-property": 1.337,
		"true-property":  true,
	})

	assert.NoError(t, err)
	assert.Equal(t, int64(1337), decodeInt64Value(t, m.valuesPut["nested/test-resource/test-pk/1337-property"]))
	assert.Equal(t, 1.337, decodeFloat64Value(t, m.valuesPut["nested/test-resource/test-pk/1.337-property"]))
	assert.Equal(t, true, decodeBoolValue(t, m.valuesPut["nested/test-resource/test-pk/true-property"]))
	m.AssertExpectations(t)
}

func TestNestingSinkUpdateFailsWhenInvalidPropertyValuePassed(t *testing.T) {
	s := givenNestingSink(givenKVClientMock())
	err := s.Update("test-resource", "test-pk", map[string]interface{}{"property": make(chan struct{})})
	assert.Error(t, err)
}

func TestNestingSinkUpdateFailsWhenKVClientFails(t *testing.T) {
	m := givenKVClientMock()
	m.onPut("nested/test-resource/test-pk/property", "foo").Return(nil, assert.AnError).Once()
	s := givenNestingSink(m)

	err := s.Update("test-resource", "test-pk", map[string]interface{}{
		"property": "foo",
	})

	assert.Error(t, err)
	m.AssertExpectations(t)
}

// TODO(daniel): remove duplication
// nolint: dupl
func TestNestingSinkUpdatePutsEtcdKeysForAllParameters(t *testing.T) {
	m := givenKVClientMock()
	m.onPut("nested/test-resource/test-pk/foo-property", "foo").Return(&clientv3.PutResponse{}, nil).Once()
	m.onPut("nested/test-resource/test-pk/1337-property", mock.Anything).Return(&clientv3.PutResponse{}, nil).Once()
	m.onPut("nested/test-resource/test-pk/1.337-property", mock.Anything).Return(&clientv3.PutResponse{}, nil).Once()
	m.onPut("nested/test-resource/test-pk/true-property", mock.Anything).Return(&clientv3.PutResponse{}, nil).Once()
	s := givenNestingSink(m)

	err := s.Update("test-resource", "test-pk", map[string]interface{}{
		"foo-property":   "foo",
		"1337-property":  1337,
		"1.337-property": 1.337,
		"true-property":  true,
	})

	assert.NoError(t, err)
	assert.Equal(t, int64(1337), decodeInt64Value(t, m.valuesPut["nested/test-resource/test-pk/1337-property"]))
	assert.Equal(t, 1.337, decodeFloat64Value(t, m.valuesPut["nested/test-resource/test-pk/1.337-property"]))
	assert.Equal(t, true, decodeBoolValue(t, m.valuesPut["nested/test-resource/test-pk/true-property"]))
	m.AssertExpectations(t)
}

func decodeInt64Value(t *testing.T, encodedValue string) interface{} {
	var value int64
	err := binary.Read(bytes.NewBufferString(encodedValue), binary.BigEndian, &value)
	require.NoError(t, err)

	return value
}

func decodeFloat64Value(t *testing.T, encodedValue string) interface{} {
	var value float64
	err := binary.Read(bytes.NewBufferString(encodedValue), binary.BigEndian, &value)
	require.NoError(t, err)

	return value
}

func decodeBoolValue(t *testing.T, encodedValue string) interface{} {
	var value bool
	err := binary.Read(bytes.NewBufferString(encodedValue), binary.BigEndian, &value)
	require.NoError(t, err)

	return value
}

func TestNestingSinkDeleteFailsWhenKVClientFails(t *testing.T) {
	m := givenKVClientMock()
	m.onDelete("nested/test-resource/test-pk").Return(nil, assert.AnError).Once()
	s := givenNestingSink(m)

	err := s.Delete("test-resource", "test-pk")

	assert.Error(t, err)
	m.AssertExpectations(t)
}

func TestNestingSinkDeleteRemovesAllEtcdKeysOfGivenResource(t *testing.T) {
	m := givenKVClientMock()
	m.onDelete("nested/test-resource/test-pk").Return(&clientv3.DeleteResponse{}, nil).Once()
	s := givenNestingSink(m)

	err := s.Delete("test-resource", "test-pk")

	assert.NoError(t, err)
	m.AssertExpectations(t)
}

func givenNestingSink(kv clientv3.KV) *NestingSink {
	return NewNestingSink(kv)
}
