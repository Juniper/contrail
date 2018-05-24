package sink

import (
	"encoding/json"
	"fmt"
	"path"

	"github.com/Juniper/contrail/pkg/db"
	"github.com/gogo/protobuf/proto"
	"github.com/imdario/mergo"
)

// Codec can encode objects and update encoded data with new data object.
type Codec interface {
	Encode(obj db.Object) ([]byte, error)
	Decode(data []byte, dst db.Object) error
	Key() string
}

func ResourceKey(c Codec, resourceName, pk interface{}) string {
	return path.Join(c.Key(), fmt.Sprint(resourceName), fmt.Sprint(pk))
}

func UpdateResourceData(c Codec, oldData []byte, updateObject db.Object) ([]byte, error) {
	if updateObject == nil {
		return oldData, nil
	}
	dst := proto.Clone(updateObject)
	dst.Reset()

	if err := c.Decode(oldData, dst); err != nil {
		return nil, err
	}
	if err := mergo.Merge(dst, updateObject, mergo.WithOverride); err != nil {
		return nil, err
	}
	return c.Encode(dst)
}

// JSONCodec is Codec which uses JSON format for storing data.
var JSONCodec Codec = jsonCodec{}

type jsonCodec struct{}

// Encode serializes data to JSON format.
func (j jsonCodec) Encode(obj db.Object) ([]byte, error) {
	return json.Marshal(obj)
}

// Decode deserializes data to map[string]interface{}.
func (j jsonCodec) Decode(data []byte, dst db.Object) error {
	return json.Unmarshal(data, &dst)
}

// Key returns codec identifier.
func (j jsonCodec) Key() string {
	return "json"
}

// ProtoCodec is Codec which uses protobuf format for storing data.
var ProtoCodec Codec // TODO(Michal): implement
