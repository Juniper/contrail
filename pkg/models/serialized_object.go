package models

import (
	"github.com/gogo/protobuf/proto"
)

// SerializedObject represents serialized data and its object form.
type SerializedObject struct {
	data  []byte
	obj   proto.Message
	codec Codec
}

// NewSerializedObject constructs serialized object.
func NewSerializedObject(d []byte, o proto.Message, c Codec) *SerializedObject {
	return &SerializedObject{
		data:  d,
		obj:   o,
		codec: c,
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
func (s *SerializedObject) Map(f func()) error {
	err := s.codec.Decode(s.data, s.obj)
	if err != nil {
		return err
	}
	f()
	s.data, err = s.codec.Encode(s.obj)
	return err
}
