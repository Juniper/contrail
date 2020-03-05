package replication

import "context"

// ChangeOperation is an enum for Operation types.
type ChangeOperation int

const (
	// CreateOperation is a value for Create Operation.
	CreateOperation ChangeOperation = iota
	// UpdateOperation is a value for Update Operation.
	UpdateOperation
	// DeleteOperation is a value for Delete Operation.
	DeleteOperation
)

// Change contains change data.
type Change interface {
	PK() []string
	Kind() string
	Operation() ChangeOperation
	Data() map[string]interface{}
}

// ChangeHandler is a Change Handler.
type ChangeHandler interface {
	Handle(ctx context.Context, changes []Change) error
}

// ChangeHandlerFunc is an alias for Change function.
type ChangeHandlerFunc func(ctx context.Context, changes []Change) error

// Handle is a method for handling Change.
func (f ChangeHandlerFunc) Handle(ctx context.Context, changes []Change) error {
	return f(ctx, changes)
}
