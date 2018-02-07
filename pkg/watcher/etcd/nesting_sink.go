package etcd

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"path"

	pkglog "github.com/Juniper/contrail/pkg/log"
	"github.com/coreos/etcd/clientv3"
	"github.com/sirupsen/logrus"
	"golang.org/x/net/context"
)

const nestedPrefix = "nested"

// NestingSink creates, updates and deletes data in etcd.
// It uses nested strategy by creating etcd key for each resource property.
// Value of each fixed-sized number is binary-encoded with big endian byte order.
type NestingSink struct {
	kvClient clientv3.KV
	log      *logrus.Entry
}

// NewNestingSink is a constructor.
func NewNestingSink(kv clientv3.KV) *NestingSink {
	if kv == nil {
		kv = &noopKVClient{}
	}

	return &NestingSink{
		kvClient: kv,
		log:      pkglog.NewLogger("nesting-sink"),
	}
}

// Create puts given properties to etcd using "nested/<resourceName>/<resourcePrimaryKey>" key prefix.
func (s *NestingSink) Create(resourceName string, pk string, properties map[string]interface{}) error {
	s.log.WithFields(logrus.Fields{"key-prefix": nestedKeyPrefix(resourceName, pk), "properties": properties}).Debug(
		"Creating etcd key-value pairs")
	return s.putEachProperty(resourceName, pk, properties)
}

// Update puts given properties to etcd using "nested/<resourceName>/<resourcePrimaryKey>" key prefix.
func (s *NestingSink) Update(resourceName string, pk string, properties map[string]interface{}) error {
	s.log.WithFields(logrus.Fields{"key-prefix": nestedKeyPrefix(resourceName, pk), "properties": properties}).Debug(
		"Updating etcd key-value pairs")
	// TODO(daniel): update only changed fields
	return s.putEachProperty(resourceName, pk, properties)
}

func (s *NestingSink) putEachProperty(resourceName string, pk string, properties map[string]interface{}) error {
	for k, v := range properties {
		s.log.WithFields(logrus.Fields{"pk": pk, "key": k, "value": v}).Debug("Put")
		value, err := encodeValue(v)
		if err != nil {
			return fmt.Errorf("encode value %#v of key %s: %s", v, k, err)
		}

		s.log.WithFields(logrus.Fields{"pk": pk, "key": k, "encoded-value": value}).Debug("Put")
		ctx, cancel := context.WithTimeout(context.Background(), kvClientRequestTimeout)
		_, err = s.kvClient.Put(ctx, path.Join(nestedKeyPrefix(resourceName, pk), k), value)
		cancel()
		if err != nil {
			return fmt.Errorf("put property to etcd: %s", err)
		}
	}
	return nil
}

func encodeValue(value interface{}) (string, error) {
	switch v := value.(type) {
	case string:
		return v, nil
	case int:
		return binaryEncode(int64(v))
	default:
		return binaryEncode(v)
	}
}

func binaryEncode(value interface{}) (string, error) {
	var b bytes.Buffer
	err := binary.Write(&b, binary.BigEndian, value)
	if err != nil {
		return "", err
	}
	return b.String(), nil
}

// Delete removes from etcd all keys that start with "nested/<resourceName>/<resourcePrimaryKey>".
func (s *NestingSink) Delete(resourceName string, pk string) error {
	s.log.WithFields(logrus.Fields{"key-prefix": nestedKeyPrefix(resourceName, pk)}).Debug(
		"Deleting etcd key-value pairs starting with prefix")
	ctx, cancel := context.WithTimeout(context.Background(), kvClientRequestTimeout)
	_, err := s.kvClient.Delete(ctx, nestedKeyPrefix(resourceName, pk), clientv3.WithFromKey())
	cancel()
	if err != nil {
		return fmt.Errorf("delete resource from etcd: %s", err)
	}
	return nil
}

func nestedKeyPrefix(resourceName, pk string) string {
	return path.Join(nestedPrefix, resourceName, pk)
}
