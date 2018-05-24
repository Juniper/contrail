package sink

import (
	"context"
)

// Sink represents service that handler transfers data to.
type Sink interface {
	Put(ctx context.Context, key string, value []byte) error
	Delete(ctx context.Context, key string) error
	InTransaction(ctx context.Context, do func(context.Context) error) error
}
