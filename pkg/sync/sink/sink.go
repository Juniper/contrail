package sink

import (
	"context"
)

// Sink represents service that handler transfers data to.
type Sink interface {
	Put(ctx context.Context, key string, value []byte) error
	Delete(ctx context.Context, key string) error
}

type Txn interface {
	Get(key string) []byte
	Put(key string, val []byte)
}

// TxnSink represents Sink with transaction support.
type TxnSink interface {
	Sink
	InTransaction(ctx context.Context, do func(Txn) error) error
}
