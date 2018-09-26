package models

import (
	"github.com/gogo/protobuf/proto"
)

// SerializedObject represents serialized data and its object form.
type SerializedObject struct {
	data []byte
	obj  proto.Message
}

// NewSerializedObject constructs serialized object.
func NewSerializedObject(d []byte, o proto.Message) *SerializedObject {
	return &SerializedObject{
		data: d,
		obj:  o,
	}
}

// GetData returns serialized data from object.
func (s *SerializedObject) GetData() []byte {
	if s != nil {
		return s.data
	}
	return nil
}

// Map could be used to apply action on serialized object.
func (s *SerializedObject) Map(c Codec, f func()) error {
	err := c.Decode(s.data, s.obj)
	if err != nil {
		return err
	}
	f()
	s.data, err = c.Encode(s.obj)
	return err
}
