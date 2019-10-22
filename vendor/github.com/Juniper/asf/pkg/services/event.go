package services

import (
	"github.com/Juniper/asf/pkg/models"
)

// Possible operations of events.
const (
	OperationCreate = "CREATE"
	OperationUpdate = "UPDATE"
	OperationDelete = "DELETE"
	OperationMixed  = "MIXED"

	EmptyEventList = "EMPTY"
)

// ResourceEvent is an event that relates to a resource.
type ResourceEvent interface {
	GetResource() models.Object
	Operation() string
}

// ReferenceEvent is an event that relates to a reference.
type ReferenceEvent interface {
	GetID() string
	GetReference() models.Reference
	Operation() string
}

// NewRefUpdateFromEvent creates RefUpdate from ReferenceEvent.
func NewRefUpdateFromEvent(e ReferenceEvent) RefUpdate {
	ref := e.GetReference()
	u := RefUpdate{
		Operation: ParseRefOperation(e.Operation()),
		Type:      ref.GetFromKind(),
		UUID:      e.GetID(),
		RefType:   ref.GetToKind(),
		RefUUID:   ref.GetUUID(),
	}

	if attr := ref.GetAttribute(); attr != nil {
		u.Attr = attr.ToMap()
	}
	return u
}

// ParseRefOperation parses RefOperation from string value.
func ParseRefOperation(s string) (op RefOperation) {
	switch s {
	case OperationCreate, string(RefOperationAdd):
		return RefOperationAdd
	case OperationDelete:
		return RefOperationDelete
	default:
		return RefOperation(s)
	}
}
