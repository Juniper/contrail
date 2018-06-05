package etcd

import (
	"context"
	"strconv"

	conc "github.com/coreos/etcd/clientv3/concurrency"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

// Txn is a transaction object allowing to perform operations in it.
type Txn interface {
	Get(key string) []byte
	Put(key string, val []byte)
}

var txnKey interface{} = "etcd-txn"

// GetTxn get a txn from context.
func GetTxn(ctx context.Context) Txn {
	iTxn := ctx.Value(txnKey)
	t, _ := iTxn.(Txn)
	return t
}

// WithTxn returns new context with Txn object.
func WithTxn(ctx context.Context, t Txn) context.Context {
	return context.WithValue(ctx, txnKey, t)
}

// GetInt64InTxn gets value from specified key and parses it to int.
func GetInt64InTxn(t Txn, key string) (int64, error) {
	valueData := t.Get(key)

	value, err := strconv.ParseInt(string(valueData), 10, 64)
	if err != nil {
		return 0, errors.Wrapf(err, `parsing "%s" to int64`, string(valueData))
	}
	return value, nil
}

// PutInt64InTxn puts int value in specified key.
func PutInt64InTxn(t Txn, key string, value int64) {
	valueString := strconv.FormatInt(value, 10)
	t.Put(key, []byte(valueString))
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
