package constants

// list of constant strings.
const (
	OPERATION = "operation"
	ADD       = "add"
	UPDATE    = "update"
	DELETE    = "delete"
)

// OperationCRUD CRUD opeartions type
type OperationCRUD int

// CRUD operations
const (
	OpCreate OperationCRUD = iota
	OpRead
	OpUpdate
	OpDelete
	OpInvalid
)
