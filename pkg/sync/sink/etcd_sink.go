package sink

import (
	"context"
	"fmt"
	"time"

	"github.com/Juniper/contrail/pkg/log"
	"github.com/coreos/etcd/clientv3"
	conc "github.com/coreos/etcd/clientv3/concurrency"
	"github.com/sirupsen/logrus"
)

const (
	kvClientRequestTimeout = 60 * time.Second
)

// ETCDSink creates, updates and deletes data in etcd.
// It uses codec to create one etcd key with resource encoded in codec format.
type ETCDSink struct {
	kvClient      clientv3.KV
	inTransaction func(ctx context.Context, apply func(conc.STM) error) error
	log           *logrus.Entry
}

// NewETCDSink is a constructor.
func NewETCDSink(client *clientv3.Client) *ETCDSink {
	// *clientv3.Client required by conc.NewSTM due to bad library interface design
	return &ETCDSink{
		kvClient: clientv3.NewKV(client),
		inTransaction: func(ctx context.Context, apply func(conc.STM) error) error {
			_, err := conc.NewSTM(client, apply, conc.WithAbortContext(ctx))
			return err
		},
		log: log.NewLogger("etcd-sink"),
	}
}

// Put puts value to etcd with timeout.
func (e *ETCDSink) Put(ctx context.Context, key string, value []byte) error {
	e.log.WithFields(logrus.Fields{"key": key}).Debugf(
		"Put resource in etcd")
	ctx, cancel := context.WithTimeout(ctx, kvClientRequestTimeout)
	defer cancel()

	_, err := e.kvClient.Put(ctx, key, string(value))
	if err != nil {
		return fmt.Errorf("put resource to etcd: %s", err)
	}

	return nil
}

// Delete deletes value from etcd with timeout.
func (e *ETCDSink) Delete(ctx context.Context, key string) error {
	e.log.WithFields(logrus.Fields{"key": key}).Debugf(
		"Deleting resource from etcd")
	ctx, cancel := context.WithTimeout(ctx, kvClientRequestTimeout)
	defer cancel()

	_, err := e.kvClient.Delete(ctx, key)
	if err != nil {
		return fmt.Errorf("delete JSON-encoded resource in etcd: %s", err)
	}

	return nil
}

// InTransaction wraps clientv3 transaction and wraps conc.STM with own sink.Txn.
func (e *ETCDSink) InTransaction(ctx context.Context, do func(Txn) error) error {
	ctx, cancel := context.WithTimeout(context.Background(), kvClientRequestTimeout)
	defer cancel()

	return e.inTransaction(ctx, func(stm conc.STM) error {
		return do(stmTxn{stm, e.log})
	})
}

type stmTxn struct {
	conc.STM
	log *logrus.Entry
}

func (s stmTxn) Get(key string) []byte {
	s.log.WithFields(logrus.Fields{"key": key}).Debugf(
		"Get resource from etcd in transaction")
	return []byte(s.STM.Get(key))
}

func (s stmTxn) Put(key string, val []byte) {
	s.log.WithFields(logrus.Fields{"key": key}).Debugf(
		"Put resource in etcd in transaction")
	s.STM.Put(key, string(val))
}
