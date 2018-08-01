package models

import (
	"github.com/gogo/protobuf/proto"
)

// Object is generic model instance.
type Object interface {
	proto.Message
	ToMap() map[string]interface{}
	Kind() string
	Depends() []string
	ApplyPropCollectionUpdate(*PropCollectionUpdate) (updated map[string]interface{}, err error)
}
