package models

// Reference is a generic reference instance.
type Reference interface {
	SetUUID(uuid string)
	SetTo(to []string)
	GetUUID() string
	GetTo() []string
	GetReferredKind() string
}
