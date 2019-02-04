package models

import (
	"encoding/json"
	"github.com/Juniper/contrail/pkg/format"
	"path"

	"github.com/gogo/protobuf/proto"
	"github.com/gogo/protobuf/types"
	"github.com/pkg/errors"
	"github.com/spf13/viper"

	"github.com/Juniper/contrail/pkg/constants"

	"github.com/Juniper/contrail/pkg/models/basemodels"
)

// Codec can encode objects and update encoded data with new data object.
type Codec interface {
	Encode(obj proto.Message) ([]byte, error)
	Decode(data []byte, obj proto.Message) error
	Key() string
}

// JSONCodec is Codec which uses JSON format for storing data.
var JSONCodec Codec = jsonCodec{}

type jsonCodec struct{}

// Encode serializes data to JSON format.
func (j jsonCodec) Encode(obj proto.Message) ([]byte, error) {
	return json.Marshal(obj)
}

// Decode serializes data to JSON format.
func (j jsonCodec) Decode(data []byte, obj proto.Message) error {
	return json.Unmarshal(data, obj)
}

// Key returns codec identifier.
func (j jsonCodec) Key() string {
	return "json"
}

// ProtoCodec is Codec which uses Proto format for storing data.
var ProtoCodec Codec = protoCodec{}

type protoCodec struct{}

// Encode serializes data to Proto format.
func (j protoCodec) Encode(obj proto.Message) ([]byte, error) {
	return proto.Marshal(obj)
}

// Decode serializes data to Proto format.
func (j protoCodec) Decode(data []byte, obj proto.Message) error {
	if len(data) == 0 {
		return nil
	}
	return proto.Unmarshal(data, obj)
}

// Key returns codec identifier.
func (j protoCodec) Key() string {
	return "proto"
}

// ResourceKey constructs key for given codec, resource name and pk.
func ResourceKey(resourceName, pk string) string {
	return path.Join("/", viper.GetString(constants.ETCDPathVK), resourceName, pk)
}

// UpdateData deserializes oldData into same type as object provided in update,
// applies an update and then serializes the result.
func UpdateData(c Codec, oldData []byte, update basemodels.Object, fm types.FieldMask) ([]byte, error) {
	if update == nil || len(fm.Paths) == 0 {
		return oldData, nil
	}
	if len(oldData) == 0 {
		return c.Encode(update)
	}
	oldObj := proto.Clone(update)
	if err := c.Decode(oldData, oldObj); err != nil {
		return nil, err
	}

	updateData := basemodels.ApplyFieldMask(update.ToMap(), fm)
	output, ok := oldObj.(basemodels.Object)
	if !ok {
		return nil, errors.Errorf("proto.Clone returned bad object type - %T (library bug)", oldObj)
	}
	format.ApplyMap(updateData, output)

	return c.Encode(output)
}
