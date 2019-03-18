package etcd

import (
	"context"

	conc "github.com/coreos/etcd/clientv3/concurrency"
	"github.com/sirupsen/logrus"
)

// Txn is a transaction object allowing to perform operations in it.
type Txn interface {
	Get(key string) []byte
	Put(key string, val []byte)
	Delete(key string)
}

var txnKey interface{} = "etcd-txn"

// GetTxn get a txn from context.
func GetTxn(ctx context.Context) Txn {
	iTxn := ctx.Value(txnKey)
	t, _ := iTxn.(Txn) //nolint: errcheck
	return t
}

// WithTxn returns new context with Txn object.
func WithTxn(ctx context.Context, t Txn) context.Context {
	return context.WithValue(ctx, txnKey, t)
}

type stmTxn struct {
	conc.STM
	log *logrus.Entry
}

func (s stmTxn) Get(key string) []byte {
	s.log.WithFields(logrus.Fields{"key": key}).Debugf(
		"Getting resource from etcd in transaction")
	return []byte(s.STM.Get(key))
}

func (s stmTxn) Put(key string, val []byte) {
	s.log.WithFields(logrus.Fields{"key": key}).Debugf(
		"Putting resource in etcd in transaction")
	s.STM.Put(key, string(val))
}

func (s stmTxn) Delete(key string) {
	s.log.WithFields(logrus.Fields{"key": key}).Debugf(
		"Deletting resource in etcd in transaction")
	s.STM.Del(key)
}
