package models

import (
	"fmt"
	"strings"

	"github.com/Juniper/asf/pkg/format"
)

type stubObject struct {
	UUID        string   `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty" yaml:"uuid,omitempty"`
	Name        string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty" yaml:"name,omitempty"`
	ParentUUID  string   `protobuf:"bytes,3,opt,name=parent_uuid,json=parentUuid,proto3" json:"parent_uuid,omitempty" yaml:"parent_uuid,omitempty"`
	ParentType  string   `protobuf:"bytes,4,opt,name=parent_type,json=parentType,proto3" json:"parent_type,omitempty" yaml:"parent_type,omitempty"`
	FQName      []string `protobuf:"bytes,5,rep,name=fq_name,json=fqName,proto3" json:"fq_name,omitempty" yaml:"fq_name,omitempty"`
	DisplayName string   `protobuf:"bytes,7,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty" yaml:"display_name,omitempty"`
}

func (o *stubObject) Reset() {
	*o = stubObject{}
}

func (o *stubObject) String() string {
	if o == nil {
		return "nil"
	}
	s := strings.Join([]string{`&stubObject{`,
		`UUID:` + fmt.Sprintf("%v", o.UUID) + `,`,
		`Name:` + fmt.Sprintf("%v", o.Name) + `,`,
		`ParentUUID:` + fmt.Sprintf("%v", o.ParentUUID) + `,`,
		`ParentType:` + fmt.Sprintf("%v", o.ParentType) + `,`,
		`FQName:` + fmt.Sprintf("%v", o.FQName) + `,`,
		`DisplayName:` + fmt.Sprintf("%v", o.DisplayName) + `,`,
		`}`,
	}, "")
	return s
}

func (o *stubObject) ProtoMessage() {
}

func (o *stubObject) GetUUID() string {
	if o != nil {
		return o.UUID
	}
	return ""
}

func (o *stubObject) SetUUID(uuid string) {
	o.UUID = uuid
}

func (o *stubObject) GetFQName() []string {
	if o != nil {
		return o.FQName
	}
	return nil
}

func (o *stubObject) GetParentUUID() string {
	if o != nil {
		return o.ParentUUID
	}
	return ""
}

func (o *stubObject) GetParentType() string {
	if o != nil {
		return o.ParentType
	}
	return ""
}

func (o *stubObject) Kind() string {
	return "stub-object"
}

func (o *stubObject) GetReferences() References {
	return nil
}

func (o *stubObject) GetTagReferences() References {
	return nil
}

func (o *stubObject) GetBackReferences() []Object {
	return nil
}

func (o *stubObject) GetChildren() []Object {
	return nil
}

func (o *stubObject) SetHref(string) {
}

func (o *stubObject) AddReference(interface{}) {
}

func (o *stubObject) AddBackReference(interface{}) {
}

func (o *stubObject) AddChild(interface{}) {
}

func (o *stubObject) RemoveReference(interface{}) {
}

func (o *stubObject) RemoveBackReference(interface{}) {
}

func (o *stubObject) RemoveChild(interface{}) {
}

func (o *stubObject) RemoveReferences() {
}

func (o *stubObject) ToMap() map[string]interface{} {
	if o == nil {
		return nil
	}
	return map[string]interface{}{
		"uuid":         o.UUID,
		"name":         o.Name,
		"parent_uuid":  o.ParentUUID,
		"parent_type":  o.ParentType,
		"fq_name":      o.FQName,
		"display_name": o.DisplayName,
	}
}

func (o *stubObject) ApplyMap(m map[string]interface{}) error {
	var err error
	if len(m) == 0 || o == nil {
		return nil
	}

	if val, ok := m["uuid"]; ok && val != nil {
		o.UUID, err = format.InterfaceToStringE(val)
	}
	if val, ok := m["name"]; ok && val != nil {
		o.Name, err = format.InterfaceToStringE(val)
	}
	if val, ok := m["parent_uuid"]; ok && val != nil {
		o.ParentUUID, err = format.InterfaceToStringE(val)
	}
	if val, ok := m["parent_type"]; ok && val != nil {
		o.ParentType, err = format.InterfaceToStringE(val)
	}
	if val, ok := m["fq_name"]; ok && val != nil {
		o.FQName, err = format.InterfaceToStringListE(val)
	}
	if val, ok := m["display_name"]; ok && val != nil {
		o.DisplayName, err = format.InterfaceToStringE(val)
	}
	return err
}
