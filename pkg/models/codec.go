package models

import (
	"encoding/json"
	"path"

	"github.com/gogo/protobuf/proto"
)

// Codec can encode objects and update encoded data with new data object.
type Codec interface {
	Encode(obj proto.Message) ([]byte, error)
	Decode(data []byte, obj proto.Message) error
	Key() string
}

// JSONCodec is Codec which uses JSON format for storing data.
type JSONCodec struct{}

// Encode serializes data to JSON format.
func (j JSONCodec) Encode(obj proto.Message) ([]byte, error) {
	return json.Marshal(obj)
}

// Decode serializes data to JSON format.
func (j JSONCodec) Decode(data []byte, obj proto.Message) error {
	return json.Unmarshal(data, obj)
}

// Key returns codec identifier.
func (j JSONCodec) Key() string {
	return "json"
}

// ProtoCodec is Codec which uses Proto format for storing data.
type ProtoCodec struct{}

// Encode serializes data to Proto format.
func (j ProtoCodec) Encode(obj proto.Message) ([]byte, error) {
	return proto.Marshal(obj)
}

// Decode serializes data to Proto format.
func (j ProtoCodec) Decode(data []byte, obj proto.Message) error {
	return proto.Unmarshal(data, obj)
}

// Key returns codec identifier.
func (j ProtoCodec) Key() string {
	return "proto"
}

// ResourceKey constructs key for given codec, resource name and pk.
func ResourceKey(c Codec, resourceName, pk string) string {
	return path.Join(c.Key(), resourceName, pk)
}
