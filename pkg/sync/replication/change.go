package replication

import "context"

type ChangeOperation int

const (
	CreateOperation ChangeOperation = iota
	UpdateOperation
	DeleteOperation
)

type Change interface {
	PK() []string
	Kind() string
	Operation() ChangeOperation
	Data() map[string]interface{}
}

type ChangeHandler interface {
	Handle(ctx context.Context, changes []Change) error
}

type ChangeHandlerFunc func(ctx context.Context, changes []Change) error

func (f ChangeHandlerFunc) Handle(ctx context.Context, changes []Change) error {
	return f(ctx, changes)
}
