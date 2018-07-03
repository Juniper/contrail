package sink

import (
	"context"
	"time"

	"github.com/coreos/etcd/clientv3"
	conc "github.com/coreos/etcd/clientv3/concurrency"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"

	"github.com/Juniper/contrail/pkg/db"
	"github.com/Juniper/contrail/pkg/log"
)

const (
	kvClientRequestTimeout = 60 * time.Second
)

// ETCDSink creates, updates and deletes data in etcd.
// It uses codec to create one etcd key with resource encoded in codec format.
type ETCDSink struct {
	kvClient      clientv3.KV
	codec         Codec
	inTransaction func(ctx context.Context, apply func(conc.STM) error) error
	log           *logrus.Entry
}

// NewETCDSink is a constructor.
func NewETCDSink(client *clientv3.Client, codec Codec) *ETCDSink {
	// *clientv3.Client required by conc.NewSTM due to bad library interface design
	return &ETCDSink{
		kvClient: clientv3.NewKV(client),
		inTransaction: func(ctx context.Context, apply func(conc.STM) error) error {
			_, err := conc.NewSTM(client, apply, conc.WithAbortContext(ctx))
			return errors.WithStack(err)
		},
		codec: codec,
		log:   log.NewLogger("etcd-sink"),
	}
}

// Create puts JSON-encoded object to etcd under "<resourceName>/json/<resourcePrimaryKey>" key.
func (s *ETCDSink) Create(resourceName string, pk string, object db.Object) error {
	s.log.WithFields(logrus.Fields{"key": resourceKey(s.codec, resourceName, pk), "object": object}).Debugf(
		"Creating %s-encoded resource in etcd", s.codec.Key())

	p, err := s.codec.Encode(object)
	if err != nil {
		return errors.Wrapf(err, "encode object to %s", s.codec.Key())
	}

	ctx, cancel := context.WithTimeout(context.Background(), kvClientRequestTimeout)
	defer cancel()
	_, err = s.kvClient.Put(ctx, resourceKey(s.codec, resourceName, pk), string(p))
	if err != nil {
		return errors.Wrapf(err, "put %s-encoded resource to etcd", s.codec.Key())
	}
	return nil
}

// Update puts JSON-encoded object to etcd under "<resourceName>/json/<resourcePrimaryKey>" key.
func (s *ETCDSink) Update(resourceName string, pk string, object db.Object) error {
	key := resourceKey(s.codec, resourceName, pk)
	s.log.WithFields(logrus.Fields{"key": key, "object": object}).Debugf(
		"Updating %s-encoded resource in etcd", s.codec.Key())

	ctx, cancel := context.WithTimeout(context.Background(), kvClientRequestTimeout)
	defer cancel()

	err := s.inTransaction(ctx, func(stm conc.STM) error {
		oldValue := stm.Get(key)
		newValue, err := s.codec.Update([]byte(oldValue), object)
		if err != nil {
			return err
		}
		stm.Put(key, string(newValue))
		return nil
	})

	return err
}

// Delete removes from etcd "<resourceName>/json/<resourcePrimaryKey>" key.
func (s *ETCDSink) Delete(resourceName string, pk string) error {
	s.log.WithFields(logrus.Fields{"key": resourceKey(s.codec, resourceName, pk)}).Debugf(
		"Deleting %s-encoded resource from etcd", s.codec.Key())
	ctx, cancel := context.WithTimeout(context.Background(), kvClientRequestTimeout)
	defer cancel()

	_, err := s.kvClient.Delete(ctx, resourceKey(s.codec, resourceName, pk))
	if err != nil {
		return errors.Wrapf(err, "delete %s-encoded resource in etcd", s.codec.Key())
	}
	return nil
}
