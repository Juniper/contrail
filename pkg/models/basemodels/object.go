package basemodels

import (
	"github.com/gogo/protobuf/proto"
)

// Object is generic model instance.
type Object interface {
	proto.Message
	GetUUID() string
	GetFQName() []string
	GetParentUUID() string
	Kind() string
	GetReferences() []Reference
	GetBackReferences() []Object
	GetChildren() []Object
	AddBackReference(interface{})
	AddChild(interface{})
	RemoveBackReference(interface{})
	RemoveChild(interface{})
	ToMap() map[string]interface{}
	ApplyMap(map[string]interface{})
	ApplyPropCollectionUpdate(*PropCollectionUpdate) (updated map[string]interface{}, err error)
}
