package sink

import (
	"encoding/json"
	"errors"
	"fmt"
	"path"

	"github.com/Juniper/contrail/pkg/db"
	"github.com/gogo/protobuf/proto"
	"github.com/imdario/mergo"
)

type Codec interface {
	Encode(obj db.Object) ([]byte, error)
	Update(data []byte, obj db.Object) ([]byte, error)
	Key() string
}

type JSONCodec struct{}

func (j *JSONCodec) Encode(obj db.Object) ([]byte, error) {
	return json.Marshal(obj)
}

func (j *JSONCodec) Update(data []byte, obj db.Object) ([]byte, error) {
	if obj == nil {
		return nil, errors.New("got nil db.Object")
	}
	// Get new instance of the same type as obj.
	dst := proto.Clone(obj)
	dst.Reset()

	// Unmarshal old data into fresh instance.
	err := json.Unmarshal(data, dst)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling previous object data: %v", err)
	}

	// Merge new obj to old.
	if err := mergo.Merge(dst, obj, mergo.WithOverride); err != nil {
		return nil, fmt.Errorf("error merging structs: %v", err)
	}

	return j.Encode(dst)
}

func (j *JSONCodec) Key() string {
	return "json"
}

func resourceKey(c Codec, resourceName, pk string) string {
	return path.Join(c.Key(), resourceName, pk)
}
