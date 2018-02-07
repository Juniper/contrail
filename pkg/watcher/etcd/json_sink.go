package etcd

import (
	"context"
	"encoding/json"
	"fmt"
	"path"

	"github.com/Juniper/contrail/pkg/log"
	"github.com/coreos/etcd/clientv3"
	"github.com/sirupsen/logrus"
)

const jsonPrefix = "json"

// JSONSink creates, updates and deletes data in etcd.
// It uses JSON strategy by creating one etcd key with JSON-encoded resource.
type JSONSink struct {
	kvClient clientv3.KV
	log      *logrus.Entry
}

// NewJSONSink is a constructor.
func NewJSONSink(kv clientv3.KV) *JSONSink {
	if kv == nil {
		kv = &noopKVClient{}
	}

	return &JSONSink{
		kvClient: kv,
		log:      log.NewLogger("json-sink"),
	}
}

// Create puts JSON-encoded properties to etcd under "<resourceName>/json/<resourcePrimaryKey>" key.
func (s *JSONSink) Create(resourceName string, pk string, properties map[string]interface{}) error {
	s.log.WithFields(logrus.Fields{"key": jsonKey(resourceName, pk), "properties": properties}).Debug(
		"Creating JSON-encoded resource in etcd")
	return s.putJSONEncodedProperties(resourceName, pk, properties)
}

// Update puts JSON-encoded properties to etcd under "<resourceName>/json/<resourcePrimaryKey>" key.
func (s *JSONSink) Update(resourceName string, pk string, properties map[string]interface{}) error {
	s.log.WithFields(logrus.Fields{"key": jsonKey(resourceName, pk), "properties": properties}).Debug(
		"Updating JSON-encoded resource in etcd")
	// TODO(daniel): check if leaving unchanged fields works
	return s.putJSONEncodedProperties(resourceName, pk, properties)
}

func (s *JSONSink) putJSONEncodedProperties(resourceName, pk string, properties map[string]interface{}) error {
	p, err := json.Marshal(properties)
	if err != nil {
		return fmt.Errorf("encode properties to JSON: %s", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), kvClientRequestTimeout)
	_, err = s.kvClient.Put(ctx, jsonKey(resourceName, pk), string(p))
	cancel()
	if err != nil {
		return fmt.Errorf("put JSON-encoded resource to etcd: %s", err)
	}
	return nil
}

// Delete removes from etcd "<resourceName>/json/<resourcePrimaryKey>" key.
func (s *JSONSink) Delete(resourceName string, pk string) error {
	s.log.WithFields(logrus.Fields{"key": jsonKey(resourceName, pk)}).Debug("Deleting JSON-encoded resource from etcd")
	ctx, cancel := context.WithTimeout(context.Background(), kvClientRequestTimeout)
	_, err := s.kvClient.Delete(ctx, jsonKey(resourceName, pk))
	cancel()
	if err != nil {
		return fmt.Errorf("delete JSON-encoded resource in etcd: %s", err)
	}
	return nil
}

func jsonKey(resourceName, pk string) string {
	return path.Join(jsonPrefix, resourceName, pk)
}
