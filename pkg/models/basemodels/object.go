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
	TypeName() string
	Depends() []string
	AddDependency(interface{})
	RemoveDependency(interface{})
	ToMap() map[string]interface{}
	ApplyMap(map[string]interface{})
	ApplyPropCollectionUpdate(*PropCollectionUpdate) (updated map[string]interface{}, err error)
}

//Ownable is generic interface which have a owner.
type Ownable interface {
	GetPerms2Owner() string
}
