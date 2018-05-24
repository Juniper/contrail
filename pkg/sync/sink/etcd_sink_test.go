package sink

import (
	"context"
	"testing"

	"github.com/Juniper/contrail/pkg/log"
	"github.com/coreos/etcd/clientv3"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestETCDSinkCreateFailsWhenKVClientFails(t *testing.T) {
	m := givenKVClientMock()
	m.onPut("json/test-resource/test-pk", mock.AnythingOfType("string")).Return(nil, assert.AnError).Once()
	s := givenETCDSink(m)

	err := s.Put(context.Background(), "json/test-resource/test-pk", []byte(`{}`))

	assert.Error(t, err)
	m.AssertExpectations(t)
}

func TestETCDSinkCreatePutsJSONEncodedResourceUnderOneEtcdKey(t *testing.T) {
	m := givenKVClientMock()
	m.onPut("json/test-resource/test-pk", "{}").Return(&clientv3.PutResponse{}, nil).Once()
	s := givenETCDSink(m)

	err := s.Put(context.Background(), "json/test-resource/test-pk", []byte(`{}`))

	assert.NoError(t, err)
	m.AssertExpectations(t)
}

func TestETCDSinkDeleteFailsWhenKVClientFails(t *testing.T) {
	m := givenKVClientMock()
	m.onDelete("json/test-resource/test-pk").Return(nil, assert.AnError).Once()
	s := givenETCDSink(m)

	err := s.Delete(context.Background(), "json/test-resource/test-pk")

	assert.Error(t, err)
	m.AssertExpectations(t)
}

func TestETCDSinkDeleteRemovesEtcdKeyWithGivenResource(t *testing.T) {
	m := givenKVClientMock()
	m.onDelete("json/test-resource/test-pk").Return(&clientv3.DeleteResponse{}, nil).Once()
	s := givenETCDSink(m)

	err := s.Delete(context.Background(), "json/test-resource/test-pk")

	assert.NoError(t, err)
	m.AssertExpectations(t)
}

func givenETCDSink(kv clientv3.KV) *ETCDSink {
	return &ETCDSink{
		kvClient: kv,
		log:      log.NewLogger("etcd-sink"),
	}
}
