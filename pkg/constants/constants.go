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
	OpCreate OpCRUD = iota + 1
	OpRead     
	OpUpdate   
	OpDelete  
	OpInvalid
)

// CRUD operations rune constants
const (
	UpperCaseC rune = 'C'
	UpperCaseR rune = 'R'  
	UpperCaseU rune = 'U' 
	UpperCaseD rune = 'D' 
)
