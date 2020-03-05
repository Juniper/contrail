package replication

type ChangeOperation string

const (
	CreateOperation ChangeOperation = "CREATE"
	UpdateOperation ChangeOperation = "UPDATE"
	DeleteOperation ChangeOperation = "DELETE"
)

type Change interface {
	PK() []string
	Kind() string
	Operation() ChangeOperation
	Data() map[string]interface{}
}

type ChangeHandler interface {
	Handle(changes []Change) error
}
