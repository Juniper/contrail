package constants

// list of constant strings.
const (
	OPERATION = "operation"
	ADD       = "add"
	UPDATE    = "update"
	DELETE    = "delete"
)

// OpCrud CRUD opeartions type
type OpCrud int

// CRUD operations
const (
	OpCreate OpCrud = 1
	OpRead   OpCrud = 2
	OpUpdate OpCrud = 3
	OpDelete OpCrud = 4
)
