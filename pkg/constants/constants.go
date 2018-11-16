package constants

// list of constant strings.
const (
	OPERATION = "operation"
	ADD       = "add"
	UPDATE    = "update"
	DELETE    = "delete"
)

// OpCRUD CRUD opeartions type
type OpCRUD int

// CRUD operations
const (
	OpCreate OpCRUD = 1
	OpRead   OpCRUD = 2
	OpUpdate OpCRUD = 3
	OpDelete OpCRUD = 4
)
