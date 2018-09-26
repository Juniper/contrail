package models

import (
	"github.com/gogo/protobuf/proto"
)

type SerializedObject struct {
	data []byte
	obj  proto.Message
}

func NewSerializedObject(d []byte, o proto.Message) *SerializedObject {
	return &SerializedObject{
		data: d,
		obj:  o,
	}
}

func (s *SerializedObject) GetData() []byte {
	if s != nil {
		return s.data
	}
	return nil
}

func (s *SerializedObject) Map(c Codec, f func()) error {
	err := c.Decode(s.data, s.obj)
	if err != nil {
		return err
	}
	f()
	s.data, err = c.Encode(s.obj)
	return err
}
