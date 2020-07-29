package services

import "context"

// InTransactionDoer executes do function atomically.
type InTransactionDoer interface {
	DoInTransaction(ctx context.Context, do func(context.Context) error) error
}

// NoTransaction executes do function non-atomically.
var NoTransaction = noTransaction{}

type noTransaction struct{}

// DoInTransaction just runs do.
func (noTransaction) DoInTransaction(ctx context.Context, do func(context.Context) error) error {
	return do(ctx)
}
