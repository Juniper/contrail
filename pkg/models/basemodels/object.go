package basemodels

import (
	"github.com/gogo/protobuf/proto"
)

// Object is generic model instance.
type Object interface {
	proto.Message
	GetUUID() string
	SetUUID(string)
	GetFQName() []string
	GetParentUUID() string
	GetParentType() string
	Kind() string
	GetReferences() References
	GetTagReferences() References
	GetBackReferences() []Object
	GetChildren() []Object
	AddBackReference(interface{})
	AddChild(interface{})
	RemoveBackReference(interface{})
	RemoveChild(interface{})
	RemoveReferences()
	ToMap() map[string]interface{}
	ApplyMap(map[string]interface{})
	ApplyPropCollectionUpdate(*PropCollectionUpdate) (updated map[string]interface{}, err error)
}
