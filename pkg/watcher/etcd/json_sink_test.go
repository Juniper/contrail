package etcd

import (
	"encoding/json"
	"testing"

	"github.com/coreos/etcd/clientv3"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestJSONSinkIsNoopByDefault(t *testing.T) {
	s := givenJSONSink(nil)

	err := s.Create("test-resource", "test-pk", map[string]interface{}{"property": "foo"})
	assert.NoError(t, err)

	err = s.Update("test-resource", "test-pk", map[string]interface{}{"property": 1337})
	assert.NoError(t, err)

	err = s.Delete("test-resource", "test-pk")
	assert.NoError(t, err)
}

func TestJSONSinkCreateFailsWhenKVClientFails(t *testing.T) {
	m := givenKVClientMock()
	m.onPut("json/test-resource/test-pk", mock.AnythingOfType("string")).Return(nil, assert.AnError).Once()
	s := givenJSONSink(m)

	err := s.Create("test-resource", "test-pk", map[string]interface{}{"property": 1337})

	assert.Error(t, err)
	m.AssertExpectations(t)
}

func TestJSONSinkCreateFailsWhenPropertiesEncodingFails(t *testing.T) {
	s := givenJSONSink(nil)

	err := s.Create("test-resource", "test-pk", map[string]interface{}{
		"key": map[float64]interface{}{1.337: "value"},
	})

	assert.Error(t, err)
}

func TestJSONSinkCreatePutsJSONEncodedResourceUnderOneEtcdKey(t *testing.T) {
	resource := map[string]interface{}{
		"foo-property":   "foo",
		"1337-property":  1337,
		"1.337-property": 1.337,
		"true-property":  true,
	}
	m := givenKVClientMock()
	m.onPut("json/test-resource/test-pk", string(toJSON(t, resource))).Return(&clientv3.PutResponse{}, nil).Once()
	s := givenJSONSink(m)

	err := s.Create("test-resource", "test-pk", resource)

	assert.NoError(t, err)
	m.AssertExpectations(t)
}

func TestJSONSinkUpdateFailsWhenKVClientFails(t *testing.T) {
	m := givenKVClientMock()
	m.onPut("json/test-resource/test-pk", mock.AnythingOfType("string")).Return(nil, assert.AnError).Once()
	s := givenJSONSink(m)

	err := s.Update("test-resource", "test-pk", map[string]interface{}{"property": "foo"})

	assert.Error(t, err)
	m.AssertExpectations(t)
}

func TestJSONSinkUpdatePutsJSONEncodedResourceUnderOneEtcdKey(t *testing.T) {
	resource := map[string]interface{}{
		"foo-property":   "new-value",
		"1337-property":  1337,
		"1.337-property": 1.337,
		"true-property":  true,
	}
	m := givenKVClientMock()
	m.onPut("json/test-resource/test-pk", string(toJSON(t, resource))).Return(&clientv3.PutResponse{}, nil).Once()
	s := givenJSONSink(m)

	err := s.Update("test-resource", "test-pk", resource)

	assert.NoError(t, err)
	m.AssertExpectations(t)
}

func TestJSONSinkDeleteFailsWhenKVClientFails(t *testing.T) {
	m := givenKVClientMock()
	m.onDelete("json/test-resource/test-pk").Return(nil, assert.AnError).Once()
	s := givenJSONSink(m)

	err := s.Delete("test-resource", "test-pk")

	assert.Error(t, err)
	m.AssertExpectations(t)
}

func TestJSONSinkDeleteRemovesEtcdKeyWithGivenResource(t *testing.T) {
	m := givenKVClientMock()
	m.onDelete("json/test-resource/test-pk").Return(&clientv3.DeleteResponse{}, nil).Once()
	s := givenJSONSink(m)

	err := s.Delete("test-resource", "test-pk")

	assert.NoError(t, err)
	m.AssertExpectations(t)
}

func givenJSONSink(kv clientv3.KV) *JSONSink {
	return NewJSONSink(kv)
}

func toJSON(t *testing.T, data interface{}) []byte {
	d, err := json.Marshal(data)
	require.NoError(t, err)
	return d
}
