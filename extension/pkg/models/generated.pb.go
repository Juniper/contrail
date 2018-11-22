// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/Juniper/contrail/extension/pkg/models/generated.proto

/*
Package models is a generated protocol buffer package.

It is generated from these files:
	github.com/Juniper/contrail/extension/pkg/models/generated.proto

It has these top-level messages:
	Sample
	SampleTagRef
	Tag
	TagTagRef
	PermType2
	KeyValuePair
	KeyValuePairs
	PermType
	UuidType
	IdPermsType
	ShareType
*/
package models

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"

import strings "strings"
import reflect "reflect"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type Sample struct {
	UUID                 string          `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty" yaml:"sample"`
	Name                 string          `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty" yaml:"sample"`
	ParentUUID           string          `protobuf:"bytes,3,opt,name=parent_uuid,json=parentUuid,proto3" json:"parent_uuid,omitempty" yaml:"sample"`
	ParentType           string          `protobuf:"bytes,4,opt,name=parent_type,json=parentType,proto3" json:"parent_type,omitempty" yaml:"sample"`
	FQName               []string        `protobuf:"bytes,5,rep,name=fq_name,json=fqName" json:"fq_name,omitempty" yaml:"sample"`
	IDPerms              *IdPermsType    `protobuf:"bytes,6,opt,name=id_perms,json=idPerms" json:"id_perms,omitempty" yaml:"sample"`
	DisplayName          string          `protobuf:"bytes,7,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty" yaml:"sample"`
	Annotations          *KeyValuePairs  `protobuf:"bytes,8,opt,name=annotations" json:"annotations,omitempty" yaml:"sample"`
	Perms2               *PermType2      `protobuf:"bytes,9,opt,name=perms2" json:"perms2,omitempty" yaml:"sample"`
	ConfigurationVersion int64           `protobuf:"varint,10,opt,name=configuration_version,json=configurationVersion,proto3" json:"configuration_version,omitempty" yaml:"sample"`
	ContainerConfig      string          `protobuf:"bytes,11,opt,name=container_config,json=containerConfig,proto3" json:"container_config,omitempty" yaml:"sample"`
	ContentConfig        string          `protobuf:"bytes,12,opt,name=content_config,json=contentConfig,proto3" json:"content_config,omitempty" yaml:"sample"`
	LayoutConfig         string          `protobuf:"bytes,13,opt,name=layout_config,json=layoutConfig,proto3" json:"layout_config,omitempty" yaml:"sample"`
	TagRefs              []*SampleTagRef `protobuf:"bytes,1014,rep,name=tag_refs,json=tagRefs" json:"tag_refs,omitempty" yaml:"tag_refs"`
}

func (m *Sample) Reset()                    { *m = Sample{} }
func (*Sample) ProtoMessage()               {}
func (*Sample) Descriptor() ([]byte, []int) { return fileDescriptorGenerated, []int{0} }

func (m *Sample) GetUUID() string {
	if m != nil {
		return m.UUID
	}
	return ""
}

func (m *Sample) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Sample) GetParentUUID() string {
	if m != nil {
		return m.ParentUUID
	}
	return ""
}

func (m *Sample) GetParentType() string {
	if m != nil {
		return m.ParentType
	}
	return ""
}

func (m *Sample) GetFQName() []string {
	if m != nil {
		return m.FQName
	}
	return nil
}

func (m *Sample) GetIDPerms() *IdPermsType {
	if m != nil {
		return m.IDPerms
	}
	return nil
}

func (m *Sample) GetDisplayName() string {
	if m != nil {
		return m.DisplayName
	}
	return ""
}

func (m *Sample) GetAnnotations() *KeyValuePairs {
	if m != nil {
		return m.Annotations
	}
	return nil
}

func (m *Sample) GetPerms2() *PermType2 {
	if m != nil {
		return m.Perms2
	}
	return nil
}

func (m *Sample) GetConfigurationVersion() int64 {
	if m != nil {
		return m.ConfigurationVersion
	}
	return 0
}

func (m *Sample) GetContainerConfig() string {
	if m != nil {
		return m.ContainerConfig
	}
	return ""
}

func (m *Sample) GetContentConfig() string {
	if m != nil {
		return m.ContentConfig
	}
	return ""
}

func (m *Sample) GetLayoutConfig() string {
	if m != nil {
		return m.LayoutConfig
	}
	return ""
}

func (m *Sample) GetTagRefs() []*SampleTagRef {
	if m != nil {
		return m.TagRefs
	}
	return nil
}

type SampleTagRef struct {
	UUID string   `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty" yaml:"uuid"`
	To   []string `protobuf:"bytes,2,rep,name=to" json:"to,omitempty" yaml:"to"`
}

func (m *SampleTagRef) Reset()                    { *m = SampleTagRef{} }
func (*SampleTagRef) ProtoMessage()               {}
func (*SampleTagRef) Descriptor() ([]byte, []int) { return fileDescriptorGenerated, []int{1} }

func (m *SampleTagRef) GetUUID() string {
	if m != nil {
		return m.UUID
	}
	return ""
}

func (m *SampleTagRef) GetTo() []string {
	if m != nil {
		return m.To
	}
	return nil
}

type Tag struct {
	UUID                 string         `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty" yaml:"tag"`
	Name                 string         `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty" yaml:"tag"`
	ParentUUID           string         `protobuf:"bytes,3,opt,name=parent_uuid,json=parentUuid,proto3" json:"parent_uuid,omitempty" yaml:"tag"`
	ParentType           string         `protobuf:"bytes,4,opt,name=parent_type,json=parentType,proto3" json:"parent_type,omitempty" yaml:"tag"`
	FQName               []string       `protobuf:"bytes,5,rep,name=fq_name,json=fqName" json:"fq_name,omitempty" yaml:"tag"`
	IDPerms              *IdPermsType   `protobuf:"bytes,6,opt,name=id_perms,json=idPerms" json:"id_perms,omitempty" yaml:"tag"`
	DisplayName          string         `protobuf:"bytes,7,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty" yaml:"tag"`
	Annotations          *KeyValuePairs `protobuf:"bytes,8,opt,name=annotations" json:"annotations,omitempty" yaml:"tag"`
	Perms2               *PermType2     `protobuf:"bytes,9,opt,name=perms2" json:"perms2,omitempty" yaml:"tag"`
	ConfigurationVersion int64          `protobuf:"varint,10,opt,name=configuration_version,json=configurationVersion,proto3" json:"configuration_version,omitempty" yaml:"tag"`
	TagValue             string         `protobuf:"bytes,11,opt,name=tag_value,json=tagValue,proto3" json:"tag_value,omitempty" yaml:"tag"`
	TagRefs              []*TagTagRef   `protobuf:"bytes,1012,rep,name=tag_refs,json=tagRefs" json:"tag_refs,omitempty" yaml:"tag_refs"`
	SampleBackRefs       []*Sample      `protobuf:"bytes,3013,rep,name=sample_back_refs,json=sampleBackRefs" json:"sample_back_refs,omitempty" yaml:"sample_back_refs"`
	TagBackRefs          []*Tag         `protobuf:"bytes,3014,rep,name=tag_back_refs,json=tagBackRefs" json:"tag_back_refs,omitempty" yaml:"tag_back_refs"`
}

func (m *Tag) Reset()                    { *m = Tag{} }
func (*Tag) ProtoMessage()               {}
func (*Tag) Descriptor() ([]byte, []int) { return fileDescriptorGenerated, []int{2} }

func (m *Tag) GetUUID() string {
	if m != nil {
		return m.UUID
	}
	return ""
}

func (m *Tag) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Tag) GetParentUUID() string {
	if m != nil {
		return m.ParentUUID
	}
	return ""
}

func (m *Tag) GetParentType() string {
	if m != nil {
		return m.ParentType
	}
	return ""
}

func (m *Tag) GetFQName() []string {
	if m != nil {
		return m.FQName
	}
	return nil
}

func (m *Tag) GetIDPerms() *IdPermsType {
	if m != nil {
		return m.IDPerms
	}
	return nil
}

func (m *Tag) GetDisplayName() string {
	if m != nil {
		return m.DisplayName
	}
	return ""
}

func (m *Tag) GetAnnotations() *KeyValuePairs {
	if m != nil {
		return m.Annotations
	}
	return nil
}

func (m *Tag) GetPerms2() *PermType2 {
	if m != nil {
		return m.Perms2
	}
	return nil
}

func (m *Tag) GetConfigurationVersion() int64 {
	if m != nil {
		return m.ConfigurationVersion
	}
	return 0
}

func (m *Tag) GetTagValue() string {
	if m != nil {
		return m.TagValue
	}
	return ""
}

func (m *Tag) GetTagRefs() []*TagTagRef {
	if m != nil {
		return m.TagRefs
	}
	return nil
}

func (m *Tag) GetSampleBackRefs() []*Sample {
	if m != nil {
		return m.SampleBackRefs
	}
	return nil
}

func (m *Tag) GetTagBackRefs() []*Tag {
	if m != nil {
		return m.TagBackRefs
	}
	return nil
}

type TagTagRef struct {
	UUID string   `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty" yaml:"uuid"`
	To   []string `protobuf:"bytes,2,rep,name=to" json:"to,omitempty" yaml:"to"`
}

func (m *TagTagRef) Reset()                    { *m = TagTagRef{} }
func (*TagTagRef) ProtoMessage()               {}
func (*TagTagRef) Descriptor() ([]byte, []int) { return fileDescriptorGenerated, []int{3} }

func (m *TagTagRef) GetUUID() string {
	if m != nil {
		return m.UUID
	}
	return ""
}

func (m *TagTagRef) GetTo() []string {
	if m != nil {
		return m.To
	}
	return nil
}

type PermType2 struct {
	Owner        string       `protobuf:"bytes,1,opt,name=owner,proto3" json:"owner,omitempty" yaml:""`
	OwnerAccess  int64        `protobuf:"varint,2,opt,name=owner_access,json=ownerAccess,proto3" json:"owner_access,omitempty" yaml:""`
	GlobalAccess int64        `protobuf:"varint,3,opt,name=global_access,json=globalAccess,proto3" json:"global_access,omitempty" yaml:""`
	Share        []*ShareType `protobuf:"bytes,4,rep,name=share" json:"share,omitempty" yaml:""`
}

func (m *PermType2) Reset()                    { *m = PermType2{} }
func (*PermType2) ProtoMessage()               {}
func (*PermType2) Descriptor() ([]byte, []int) { return fileDescriptorGenerated, []int{4} }

func (m *PermType2) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *PermType2) GetOwnerAccess() int64 {
	if m != nil {
		return m.OwnerAccess
	}
	return 0
}

func (m *PermType2) GetGlobalAccess() int64 {
	if m != nil {
		return m.GlobalAccess
	}
	return 0
}

func (m *PermType2) GetShare() []*ShareType {
	if m != nil {
		return m.Share
	}
	return nil
}

// Omitempty tag is removed from fields of KeyValuePair type, because it caused issues in REST API clients
// which expected all fields to be present. To achieve that "gogoproto.jsontag" extension is used.
type KeyValuePair struct {
	Value string `protobuf:"bytes,1,opt,name=value,proto3" json:"" yaml:""`
	Key   string `protobuf:"bytes,2,opt,name=key,proto3" json:"" yaml:""`
}

func (m *KeyValuePair) Reset()                    { *m = KeyValuePair{} }
func (*KeyValuePair) ProtoMessage()               {}
func (*KeyValuePair) Descriptor() ([]byte, []int) { return fileDescriptorGenerated, []int{5} }

func (m *KeyValuePair) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func (m *KeyValuePair) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

type KeyValuePairs struct {
	KeyValuePair []*KeyValuePair `protobuf:"bytes,1,rep,name=key_value_pair,json=keyValuePair" json:"key_value_pair,omitempty" yaml:""`
}

func (m *KeyValuePairs) Reset()                    { *m = KeyValuePairs{} }
func (*KeyValuePairs) ProtoMessage()               {}
func (*KeyValuePairs) Descriptor() ([]byte, []int) { return fileDescriptorGenerated, []int{6} }

func (m *KeyValuePairs) GetKeyValuePair() []*KeyValuePair {
	if m != nil {
		return m.KeyValuePair
	}
	return nil
}

type PermType struct {
	Owner       string `protobuf:"bytes,1,opt,name=owner,proto3" json:"owner,omitempty" yaml:""`
	OwnerAccess int64  `protobuf:"varint,2,opt,name=owner_access,json=ownerAccess,proto3" json:"owner_access,omitempty" yaml:""`
	OtherAccess int64  `protobuf:"varint,3,opt,name=other_access,json=otherAccess,proto3" json:"other_access,omitempty" yaml:""`
	Group       string `protobuf:"bytes,4,opt,name=group,proto3" json:"group,omitempty" yaml:""`
	GroupAccess int64  `protobuf:"varint,5,opt,name=group_access,json=groupAccess,proto3" json:"group_access,omitempty" yaml:""`
}

func (m *PermType) Reset()                    { *m = PermType{} }
func (*PermType) ProtoMessage()               {}
func (*PermType) Descriptor() ([]byte, []int) { return fileDescriptorGenerated, []int{7} }

func (m *PermType) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *PermType) GetOwnerAccess() int64 {
	if m != nil {
		return m.OwnerAccess
	}
	return 0
}

func (m *PermType) GetOtherAccess() int64 {
	if m != nil {
		return m.OtherAccess
	}
	return 0
}

func (m *PermType) GetGroup() string {
	if m != nil {
		return m.Group
	}
	return ""
}

func (m *PermType) GetGroupAccess() int64 {
	if m != nil {
		return m.GroupAccess
	}
	return 0
}

type UuidType struct {
	UUIDMslong uint64 `protobuf:"varint,1,opt,name=uuid_mslong,json=uuidMslong,proto3" json:"uuid_mslong,omitempty" yaml:""`
	UUIDLslong uint64 `protobuf:"varint,2,opt,name=uuid_lslong,json=uuidLslong,proto3" json:"uuid_lslong,omitempty" yaml:""`
}

func (m *UuidType) Reset()                    { *m = UuidType{} }
func (*UuidType) ProtoMessage()               {}
func (*UuidType) Descriptor() ([]byte, []int) { return fileDescriptorGenerated, []int{8} }

func (m *UuidType) GetUUIDMslong() uint64 {
	if m != nil {
		return m.UUIDMslong
	}
	return 0
}

func (m *UuidType) GetUUIDLslong() uint64 {
	if m != nil {
		return m.UUIDLslong
	}
	return 0
}

type IdPermsType struct {
	Enable       bool      `protobuf:"varint,1,opt,name=enable,proto3" json:"enable,omitempty" yaml:""`
	Description  string    `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty" yaml:""`
	Created      string    `protobuf:"bytes,3,opt,name=created,proto3" json:"created,omitempty" yaml:""`
	Creator      string    `protobuf:"bytes,4,opt,name=creator,proto3" json:"creator,omitempty" yaml:""`
	UserVisible  bool      `protobuf:"varint,5,opt,name=user_visible,json=userVisible,proto3" json:"user_visible,omitempty" yaml:""`
	LastModified string    `protobuf:"bytes,6,opt,name=last_modified,json=lastModified,proto3" json:"last_modified,omitempty" yaml:""`
	Permissions  *PermType `protobuf:"bytes,7,opt,name=permissions" json:"permissions,omitempty" yaml:""`
	UUID         *UuidType `protobuf:"bytes,8,opt,name=uuid" json:"uuid,omitempty" yaml:""`
}

func (m *IdPermsType) Reset()                    { *m = IdPermsType{} }
func (*IdPermsType) ProtoMessage()               {}
func (*IdPermsType) Descriptor() ([]byte, []int) { return fileDescriptorGenerated, []int{9} }

func (m *IdPermsType) GetEnable() bool {
	if m != nil {
		return m.Enable
	}
	return false
}

func (m *IdPermsType) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *IdPermsType) GetCreated() string {
	if m != nil {
		return m.Created
	}
	return ""
}

func (m *IdPermsType) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *IdPermsType) GetUserVisible() bool {
	if m != nil {
		return m.UserVisible
	}
	return false
}

func (m *IdPermsType) GetLastModified() string {
	if m != nil {
		return m.LastModified
	}
	return ""
}

func (m *IdPermsType) GetPermissions() *PermType {
	if m != nil {
		return m.Permissions
	}
	return nil
}

func (m *IdPermsType) GetUUID() *UuidType {
	if m != nil {
		return m.UUID
	}
	return nil
}

type ShareType struct {
	TenantAccess int64  `protobuf:"varint,1,opt,name=tenant_access,json=tenantAccess,proto3" json:"tenant_access,omitempty" yaml:""`
	Tenant       string `protobuf:"bytes,2,opt,name=tenant,proto3" json:"tenant,omitempty" yaml:""`
}

func (m *ShareType) Reset()                    { *m = ShareType{} }
func (*ShareType) ProtoMessage()               {}
func (*ShareType) Descriptor() ([]byte, []int) { return fileDescriptorGenerated, []int{10} }

func (m *ShareType) GetTenantAccess() int64 {
	if m != nil {
		return m.TenantAccess
	}
	return 0
}

func (m *ShareType) GetTenant() string {
	if m != nil {
		return m.Tenant
	}
	return ""
}

func init() {
	proto.RegisterType((*Sample)(nil), "github.com.Juniper.contrail.extension.pkg.models.Sample")
	proto.RegisterType((*SampleTagRef)(nil), "github.com.Juniper.contrail.extension.pkg.models.SampleTagRef")
	proto.RegisterType((*Tag)(nil), "github.com.Juniper.contrail.extension.pkg.models.Tag")
	proto.RegisterType((*TagTagRef)(nil), "github.com.Juniper.contrail.extension.pkg.models.TagTagRef")
	proto.RegisterType((*PermType2)(nil), "github.com.Juniper.contrail.extension.pkg.models.PermType2")
	proto.RegisterType((*KeyValuePair)(nil), "github.com.Juniper.contrail.extension.pkg.models.KeyValuePair")
	proto.RegisterType((*KeyValuePairs)(nil), "github.com.Juniper.contrail.extension.pkg.models.KeyValuePairs")
	proto.RegisterType((*PermType)(nil), "github.com.Juniper.contrail.extension.pkg.models.PermType")
	proto.RegisterType((*UuidType)(nil), "github.com.Juniper.contrail.extension.pkg.models.UuidType")
	proto.RegisterType((*IdPermsType)(nil), "github.com.Juniper.contrail.extension.pkg.models.IdPermsType")
	proto.RegisterType((*ShareType)(nil), "github.com.Juniper.contrail.extension.pkg.models.ShareType")
}
func (m *Sample) Size() (n int) {
	var l int
	_ = l
	l = len(m.UUID)
	if l > 0 {
		n += 1 + l + sovGenerated(uint64(l))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovGenerated(uint64(l))
	}
	l = len(m.ParentUUID)
	if l > 0 {
		n += 1 + l + sovGenerated(uint64(l))
	}
	l = len(m.ParentType)
	if l > 0 {
		n += 1 + l + sovGenerated(uint64(l))
	}
	if len(m.FQName) > 0 {
		for _, s := range m.FQName {
			l = len(s)
			n += 1 + l + sovGenerated(uint64(l))
		}
	}
	if m.IDPerms != nil {
		l = m.IDPerms.Size()
		n += 1 + l + sovGenerated(uint64(l))
	}
	l = len(m.DisplayName)
	if l > 0 {
		n += 1 + l + sovGenerated(uint64(l))
	}
	if m.Annotations != nil {
		l = m.Annotations.Size()
		n += 1 + l + sovGenerated(uint64(l))
	}
	if m.Perms2 != nil {
		l = m.Perms2.Size()
		n += 1 + l + sovGenerated(uint64(l))
	}
	if m.ConfigurationVersion != 0 {
		n += 1 + sovGenerated(uint64(m.ConfigurationVersion))
	}
	l = len(m.ContainerConfig)
	if l > 0 {
		n += 1 + l + sovGenerated(uint64(l))
	}
	l = len(m.ContentConfig)
	if l > 0 {
		n += 1 + l + sovGenerated(uint64(l))
	}
	l = len(m.LayoutConfig)
	if l > 0 {
		n += 1 + l + sovGenerated(uint64(l))
	}
	if len(m.TagRefs) > 0 {
		for _, e := range m.TagRefs {
			l = e.Size()
			n += 2 + l + sovGenerated(uint64(l))
		}
	}
	return n
}

func (m *SampleTagRef) Size() (n int) {
	var l int
	_ = l
	l = len(m.UUID)
	if l > 0 {
		n += 1 + l + sovGenerated(uint64(l))
	}
	if len(m.To) > 0 {
		for _, s := range m.To {
			l = len(s)
			n += 1 + l + sovGenerated(uint64(l))
		}
	}
	return n
}

func (m *Tag) Size() (n int) {
	var l int
	_ = l
	l = len(m.UUID)
	if l > 0 {
		n += 1 + l + sovGenerated(uint64(l))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovGenerated(uint64(l))
	}
	l = len(m.ParentUUID)
	if l > 0 {
		n += 1 + l + sovGenerated(uint64(l))
	}
	l = len(m.ParentType)
	if l > 0 {
		n += 1 + l + sovGenerated(uint64(l))
	}
	if len(m.FQName) > 0 {
		for _, s := range m.FQName {
			l = len(s)
			n += 1 + l + sovGenerated(uint64(l))
		}
	}
	if m.IDPerms != nil {
		l = m.IDPerms.Size()
		n += 1 + l + sovGenerated(uint64(l))
	}
	l = len(m.DisplayName)
	if l > 0 {
		n += 1 + l + sovGenerated(uint64(l))
	}
	if m.Annotations != nil {
		l = m.Annotations.Size()
		n += 1 + l + sovGenerated(uint64(l))
	}
	if m.Perms2 != nil {
		l = m.Perms2.Size()
		n += 1 + l + sovGenerated(uint64(l))
	}
	if m.ConfigurationVersion != 0 {
		n += 1 + sovGenerated(uint64(m.ConfigurationVersion))
	}
	l = len(m.TagValue)
	if l > 0 {
		n += 1 + l + sovGenerated(uint64(l))
	}
	if len(m.TagRefs) > 0 {
		for _, e := range m.TagRefs {
			l = e.Size()
			n += 2 + l + sovGenerated(uint64(l))
		}
	}
	if len(m.SampleBackRefs) > 0 {
		for _, e := range m.SampleBackRefs {
			l = e.Size()
			n += 3 + l + sovGenerated(uint64(l))
		}
	}
	if len(m.TagBackRefs) > 0 {
		for _, e := range m.TagBackRefs {
			l = e.Size()
			n += 3 + l + sovGenerated(uint64(l))
		}
	}
	return n
}

func (m *TagTagRef) Size() (n int) {
	var l int
	_ = l
	l = len(m.UUID)
	if l > 0 {
		n += 1 + l + sovGenerated(uint64(l))
	}
	if len(m.To) > 0 {
		for _, s := range m.To {
			l = len(s)
			n += 1 + l + sovGenerated(uint64(l))
		}
	}
	return n
}

func (m *PermType2) Size() (n int) {
	var l int
	_ = l
	l = len(m.Owner)
	if l > 0 {
		n += 1 + l + sovGenerated(uint64(l))
	}
	if m.OwnerAccess != 0 {
		n += 1 + sovGenerated(uint64(m.OwnerAccess))
	}
	if m.GlobalAccess != 0 {
		n += 1 + sovGenerated(uint64(m.GlobalAccess))
	}
	if len(m.Share) > 0 {
		for _, e := range m.Share {
			l = e.Size()
			n += 1 + l + sovGenerated(uint64(l))
		}
	}
	return n
}

func (m *KeyValuePair) Size() (n int) {
	var l int
	_ = l
	l = len(m.Value)
	if l > 0 {
		n += 1 + l + sovGenerated(uint64(l))
	}
	l = len(m.Key)
	if l > 0 {
		n += 1 + l + sovGenerated(uint64(l))
	}
	return n
}

func (m *KeyValuePairs) Size() (n int) {
	var l int
	_ = l
	if len(m.KeyValuePair) > 0 {
		for _, e := range m.KeyValuePair {
			l = e.Size()
			n += 1 + l + sovGenerated(uint64(l))
		}
	}
	return n
}

func (m *PermType) Size() (n int) {
	var l int
	_ = l
	l = len(m.Owner)
	if l > 0 {
		n += 1 + l + sovGenerated(uint64(l))
	}
	if m.OwnerAccess != 0 {
		n += 1 + sovGenerated(uint64(m.OwnerAccess))
	}
	if m.OtherAccess != 0 {
		n += 1 + sovGenerated(uint64(m.OtherAccess))
	}
	l = len(m.Group)
	if l > 0 {
		n += 1 + l + sovGenerated(uint64(l))
	}
	if m.GroupAccess != 0 {
		n += 1 + sovGenerated(uint64(m.GroupAccess))
	}
	return n
}

func (m *UuidType) Size() (n int) {
	var l int
	_ = l
	if m.UUIDMslong != 0 {
		n += 1 + sovGenerated(uint64(m.UUIDMslong))
	}
	if m.UUIDLslong != 0 {
		n += 1 + sovGenerated(uint64(m.UUIDLslong))
	}
	return n
}

func (m *IdPermsType) Size() (n int) {
	var l int
	_ = l
	if m.Enable {
		n += 2
	}
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovGenerated(uint64(l))
	}
	l = len(m.Created)
	if l > 0 {
		n += 1 + l + sovGenerated(uint64(l))
	}
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovGenerated(uint64(l))
	}
	if m.UserVisible {
		n += 2
	}
	l = len(m.LastModified)
	if l > 0 {
		n += 1 + l + sovGenerated(uint64(l))
	}
	if m.Permissions != nil {
		l = m.Permissions.Size()
		n += 1 + l + sovGenerated(uint64(l))
	}
	if m.UUID != nil {
		l = m.UUID.Size()
		n += 1 + l + sovGenerated(uint64(l))
	}
	return n
}

func (m *ShareType) Size() (n int) {
	var l int
	_ = l
	if m.TenantAccess != 0 {
		n += 1 + sovGenerated(uint64(m.TenantAccess))
	}
	l = len(m.Tenant)
	if l > 0 {
		n += 1 + l + sovGenerated(uint64(l))
	}
	return n
}

func sovGenerated(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozGenerated(x uint64) (n int) {
	return sovGenerated(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *Sample) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&Sample{`,
		`UUID:` + fmt.Sprintf("%v", this.UUID) + `,`,
		`Name:` + fmt.Sprintf("%v", this.Name) + `,`,
		`ParentUUID:` + fmt.Sprintf("%v", this.ParentUUID) + `,`,
		`ParentType:` + fmt.Sprintf("%v", this.ParentType) + `,`,
		`FQName:` + fmt.Sprintf("%v", this.FQName) + `,`,
		`IDPerms:` + strings.Replace(fmt.Sprintf("%v", this.IDPerms), "IdPermsType", "IdPermsType", 1) + `,`,
		`DisplayName:` + fmt.Sprintf("%v", this.DisplayName) + `,`,
		`Annotations:` + strings.Replace(fmt.Sprintf("%v", this.Annotations), "KeyValuePairs", "KeyValuePairs", 1) + `,`,
		`Perms2:` + strings.Replace(fmt.Sprintf("%v", this.Perms2), "PermType2", "PermType2", 1) + `,`,
		`ConfigurationVersion:` + fmt.Sprintf("%v", this.ConfigurationVersion) + `,`,
		`ContainerConfig:` + fmt.Sprintf("%v", this.ContainerConfig) + `,`,
		`ContentConfig:` + fmt.Sprintf("%v", this.ContentConfig) + `,`,
		`LayoutConfig:` + fmt.Sprintf("%v", this.LayoutConfig) + `,`,
		`TagRefs:` + strings.Replace(fmt.Sprintf("%v", this.TagRefs), "SampleTagRef", "SampleTagRef", 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *SampleTagRef) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&SampleTagRef{`,
		`UUID:` + fmt.Sprintf("%v", this.UUID) + `,`,
		`To:` + fmt.Sprintf("%v", this.To) + `,`,
		`}`,
	}, "")
	return s
}
func (this *Tag) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&Tag{`,
		`UUID:` + fmt.Sprintf("%v", this.UUID) + `,`,
		`Name:` + fmt.Sprintf("%v", this.Name) + `,`,
		`ParentUUID:` + fmt.Sprintf("%v", this.ParentUUID) + `,`,
		`ParentType:` + fmt.Sprintf("%v", this.ParentType) + `,`,
		`FQName:` + fmt.Sprintf("%v", this.FQName) + `,`,
		`IDPerms:` + strings.Replace(fmt.Sprintf("%v", this.IDPerms), "IdPermsType", "IdPermsType", 1) + `,`,
		`DisplayName:` + fmt.Sprintf("%v", this.DisplayName) + `,`,
		`Annotations:` + strings.Replace(fmt.Sprintf("%v", this.Annotations), "KeyValuePairs", "KeyValuePairs", 1) + `,`,
		`Perms2:` + strings.Replace(fmt.Sprintf("%v", this.Perms2), "PermType2", "PermType2", 1) + `,`,
		`ConfigurationVersion:` + fmt.Sprintf("%v", this.ConfigurationVersion) + `,`,
		`TagValue:` + fmt.Sprintf("%v", this.TagValue) + `,`,
		`TagRefs:` + strings.Replace(fmt.Sprintf("%v", this.TagRefs), "TagTagRef", "TagTagRef", 1) + `,`,
		`SampleBackRefs:` + strings.Replace(fmt.Sprintf("%v", this.SampleBackRefs), "Sample", "Sample", 1) + `,`,
		`TagBackRefs:` + strings.Replace(fmt.Sprintf("%v", this.TagBackRefs), "Tag", "Tag", 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *TagTagRef) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&TagTagRef{`,
		`UUID:` + fmt.Sprintf("%v", this.UUID) + `,`,
		`To:` + fmt.Sprintf("%v", this.To) + `,`,
		`}`,
	}, "")
	return s
}
func (this *PermType2) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&PermType2{`,
		`Owner:` + fmt.Sprintf("%v", this.Owner) + `,`,
		`OwnerAccess:` + fmt.Sprintf("%v", this.OwnerAccess) + `,`,
		`GlobalAccess:` + fmt.Sprintf("%v", this.GlobalAccess) + `,`,
		`Share:` + strings.Replace(fmt.Sprintf("%v", this.Share), "ShareType", "ShareType", 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *KeyValuePair) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&KeyValuePair{`,
		`Value:` + fmt.Sprintf("%v", this.Value) + `,`,
		`Key:` + fmt.Sprintf("%v", this.Key) + `,`,
		`}`,
	}, "")
	return s
}
func (this *KeyValuePairs) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&KeyValuePairs{`,
		`KeyValuePair:` + strings.Replace(fmt.Sprintf("%v", this.KeyValuePair), "KeyValuePair", "KeyValuePair", 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *PermType) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&PermType{`,
		`Owner:` + fmt.Sprintf("%v", this.Owner) + `,`,
		`OwnerAccess:` + fmt.Sprintf("%v", this.OwnerAccess) + `,`,
		`OtherAccess:` + fmt.Sprintf("%v", this.OtherAccess) + `,`,
		`Group:` + fmt.Sprintf("%v", this.Group) + `,`,
		`GroupAccess:` + fmt.Sprintf("%v", this.GroupAccess) + `,`,
		`}`,
	}, "")
	return s
}
func (this *UuidType) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&UuidType{`,
		`UUIDMslong:` + fmt.Sprintf("%v", this.UUIDMslong) + `,`,
		`UUIDLslong:` + fmt.Sprintf("%v", this.UUIDLslong) + `,`,
		`}`,
	}, "")
	return s
}
func (this *IdPermsType) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&IdPermsType{`,
		`Enable:` + fmt.Sprintf("%v", this.Enable) + `,`,
		`Description:` + fmt.Sprintf("%v", this.Description) + `,`,
		`Created:` + fmt.Sprintf("%v", this.Created) + `,`,
		`Creator:` + fmt.Sprintf("%v", this.Creator) + `,`,
		`UserVisible:` + fmt.Sprintf("%v", this.UserVisible) + `,`,
		`LastModified:` + fmt.Sprintf("%v", this.LastModified) + `,`,
		`Permissions:` + strings.Replace(fmt.Sprintf("%v", this.Permissions), "PermType", "PermType", 1) + `,`,
		`UUID:` + strings.Replace(fmt.Sprintf("%v", this.UUID), "UuidType", "UuidType", 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *ShareType) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&ShareType{`,
		`TenantAccess:` + fmt.Sprintf("%v", this.TenantAccess) + `,`,
		`Tenant:` + fmt.Sprintf("%v", this.Tenant) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringGenerated(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}

func init() {
	proto.RegisterFile("github.com/Juniper/contrail/extension/pkg/models/generated.proto", fileDescriptorGenerated)
}

var fileDescriptorGenerated = []byte{
	// 1465 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x58, 0xcf, 0x73, 0xdb, 0xd4,
	0x16, 0xf6, 0x8f, 0xf8, 0xd7, 0x95, 0x9d, 0x76, 0xee, 0x4b, 0x5b, 0xb7, 0xaf, 0xb5, 0xf2, 0xf4,
	0xde, 0x3c, 0x42, 0x99, 0xd8, 0x25, 0xc0, 0xd0, 0x29, 0x04, 0x52, 0x27, 0xb4, 0x13, 0xda, 0xd2,
	0x90, 0x3a, 0x5d, 0xc0, 0x80, 0xb9, 0x91, 0xaf, 0x15, 0x8d, 0x65, 0x49, 0x95, 0xe4, 0xb4, 0x66,
	0x58, 0x14, 0xf8, 0x0f, 0x60, 0x58, 0x03, 0x3b, 0xfe, 0x0a, 0x86, 0x0d, 0x0c, 0xcb, 0x2e, 0x59,
	0x69, 0x1a, 0xb1, 0x63, 0xc5, 0x74, 0x18, 0x86, 0x25, 0x73, 0xcf, 0x95, 0xac, 0xab, 0x58, 0xed,
	0x4c, 0x3c, 0x2d, 0x3b, 0xe9, 0x9c, 0xf3, 0x7d, 0xe7, 0xdc, 0x73, 0xcf, 0xfd, 0xae, 0x6c, 0xb4,
	0xa6, 0xe9, 0xde, 0xde, 0x68, 0xb7, 0xa9, 0x5a, 0xc3, 0xd6, 0xdb, 0x23, 0x53, 0xb7, 0xa9, 0xd3,
	0x52, 0x2d, 0xd3, 0x73, 0x88, 0x6e, 0xb4, 0xe8, 0x3d, 0x8f, 0x9a, 0xae, 0x6e, 0x99, 0x2d, 0x7b,
	0xa0, 0xb5, 0x86, 0x56, 0x8f, 0x1a, 0x6e, 0x4b, 0xa3, 0x26, 0x75, 0x88, 0x47, 0x7b, 0x4d, 0xdb,
	0xb1, 0x3c, 0x0b, 0x5f, 0x88, 0x19, 0x9a, 0x21, 0x43, 0x33, 0x62, 0x68, 0x4e, 0x18, 0x9a, 0xf6,
	0x40, 0x6b, 0x72, 0x86, 0x33, 0xcb, 0x42, 0x4e, 0xcd, 0xd2, 0xac, 0x16, 0x10, 0xed, 0x8e, 0xfa,
	0xf0, 0x06, 0x2f, 0xf0, 0xc4, 0x13, 0x28, 0x5f, 0x55, 0x50, 0xf1, 0x16, 0x19, 0xda, 0x06, 0xc5,
	0xcb, 0x68, 0x6e, 0x34, 0xd2, 0x7b, 0xf5, 0xec, 0x62, 0x76, 0xa9, 0xd2, 0x3e, 0x1d, 0xf8, 0xf2,
	0xdc, 0xce, 0xce, 0xe6, 0xc6, 0x23, 0x5f, 0xae, 0x8d, 0xc9, 0xd0, 0xb8, 0xa4, 0xb8, 0x10, 0xa7,
	0x6c, 0x43, 0x18, 0x0b, 0x37, 0xc9, 0x90, 0xd6, 0x73, 0x71, 0xf8, 0x3b, 0x64, 0x48, 0x53, 0xc2,
	0x59, 0x18, 0x5e, 0x43, 0x92, 0x4d, 0x1c, 0x6a, 0x7a, 0x5d, 0x48, 0x92, 0x07, 0x94, 0x1c, 0xf8,
	0x32, 0xda, 0x02, 0x73, 0x7a, 0x2a, 0xc4, 0x31, 0x3b, 0x2c, 0x61, 0xcc, 0xe0, 0x8d, 0x6d, 0x5a,
	0x9f, 0x3b, 0xcc, 0xd0, 0x19, 0xdb, 0xf4, 0xb1, 0x0c, 0xcc, 0x89, 0x5f, 0x46, 0xa5, 0xfe, 0x9d,
	0x2e, 0x54, 0x5d, 0x58, 0xcc, 0x2f, 0x55, 0xda, 0xff, 0x0e, 0x7c, 0xb9, 0x78, 0xe5, 0xdd, 0xf4,
	0xba, 0x8b, 0xfd, 0x3b, 0xcc, 0x81, 0xf7, 0x51, 0x59, 0xef, 0x75, 0x6d, 0xea, 0x0c, 0xdd, 0x7a,
	0x71, 0x31, 0xbb, 0x24, 0xad, 0xac, 0x36, 0x8f, 0xba, 0x2d, 0xcd, 0xcd, 0xde, 0x16, 0x23, 0x60,
	0x65, 0xb4, 0xcf, 0x06, 0xbe, 0x5c, 0xda, 0xdc, 0x00, 0xc3, 0x74, 0xda, 0x92, 0xce, 0x43, 0xf1,
	0x3a, 0xaa, 0xf6, 0x74, 0xd7, 0x36, 0xc8, 0x98, 0x97, 0x5c, 0x82, 0x05, 0x2f, 0x06, 0xbe, 0x2c,
	0x6d, 0x70, 0x7b, 0x7a, 0xdd, 0x52, 0x2f, 0xf6, 0xe2, 0xcf, 0xb3, 0x48, 0x22, 0xa6, 0x69, 0x79,
	0xc4, 0xd3, 0x2d, 0xd3, 0xad, 0x97, 0x61, 0x01, 0x6f, 0x1e, 0x7d, 0x01, 0xd7, 0xe8, 0xf8, 0x36,
	0x31, 0x46, 0x74, 0x8b, 0xe8, 0x8e, 0xcb, 0xab, 0xb8, 0x1c, 0xf3, 0xa6, 0x54, 0x21, 0x64, 0xc5,
	0x36, 0x2a, 0x42, 0xff, 0x56, 0xea, 0x15, 0xc8, 0xff, 0xda, 0xd1, 0xf3, 0xb3, 0x9e, 0xb0, 0xee,
	0xad, 0xf0, 0x4d, 0x83, 0x16, 0xad, 0xa4, 0x6c, 0x1a, 0xcf, 0x83, 0x3f, 0x44, 0x27, 0x54, 0xcb,
	0xec, 0xeb, 0xda, 0xc8, 0x81, 0x1a, 0xba, 0xfb, 0xd4, 0x61, 0x94, 0x75, 0xb4, 0x98, 0x5d, 0xca,
	0xb7, 0x9f, 0x0f, 0x7c, 0x79, 0x61, 0x5d, 0x0c, 0xb8, 0xcd, 0xfd, 0xd3, 0x8c, 0x0b, 0x6a, 0x4a,
	0x18, 0xbe, 0x89, 0x8e, 0xb3, 0x7a, 0x89, 0x6e, 0x52, 0xa7, 0xcb, 0x23, 0xea, 0x12, 0x6c, 0xd0,
	0xff, 0x02, 0x5f, 0x3e, 0xb6, 0x1e, 0xf9, 0x78, 0x8e, 0x69, 0xd6, 0x63, 0x6a, 0x32, 0x02, 0x6f,
	0xa2, 0x79, 0x66, 0x62, 0xe3, 0x1d, 0xd2, 0x55, 0x81, 0x4e, 0x09, 0x7c, 0xb9, 0xb6, 0xce, 0x3d,
	0x8f, 0x23, 0xab, 0xa9, 0xa2, 0x1f, 0x5f, 0x41, 0x35, 0x83, 0x8c, 0xad, 0xd1, 0x84, 0xa9, 0x06,
	0x4c, 0xff, 0x09, 0x7c, 0xb9, 0x7a, 0x1d, 0x1c, 0x8f, 0x23, 0xaa, 0x1a, 0x82, 0x1b, 0x7f, 0x8c,
	0xca, 0x1e, 0xd1, 0xba, 0x0e, 0xed, 0xbb, 0xf5, 0x3f, 0x4b, 0x8b, 0xf9, 0x25, 0x69, 0xe5, 0x8d,
	0xa3, 0x6f, 0x1c, 0x57, 0x97, 0x0e, 0xd1, 0xb6, 0x69, 0xbf, 0xdd, 0x60, 0xa3, 0xcf, 0x9f, 0xd9,
	0xcc, 0x1c, 0xe3, 0xe9, 0xa3, 0x2c, 0xca, 0x76, 0xc9, 0xe3, 0x3e, 0xe5, 0x23, 0x54, 0x15, 0x81,
	0xf8, 0x85, 0x84, 0x38, 0x9d, 0x12, 0xc4, 0x49, 0xe2, 0x1c, 0xcc, 0x1b, 0x49, 0xd3, 0x7f, 0x51,
	0xce, 0xb3, 0xea, 0x39, 0x38, 0xe2, 0xff, 0x0a, 0x7c, 0x39, 0xd7, 0xb1, 0x1e, 0xf9, 0x72, 0x25,
	0x4c, 0x66, 0x29, 0xdb, 0x39, 0xcf, 0x52, 0x0e, 0x2a, 0x28, 0xdf, 0x21, 0x1a, 0x3e, 0x9f, 0x60,
	0x3e, 0x29, 0x30, 0xa3, 0x49, 0x75, 0x11, 0xf1, 0xf9, 0x84, 0xe6, 0x9d, 0x14, 0x34, 0x2f, 0x11,
	0x0b, 0x82, 0xb7, 0x9a, 0x26, 0x78, 0x67, 0xa7, 0x04, 0x4f, 0x04, 0x8a, 0x6a, 0xb7, 0x9a, 0xa6,
	0x76, 0x67, 0xa7, 0xd4, 0x2e, 0x05, 0x0e, 0x52, 0xf7, 0xe2, 0x61, 0xa9, 0xab, 0x27, 0xa4, 0x4e,
	0x84, 0x45, 0x3a, 0xe7, 0x3e, 0x6d, 0x9d, 0x3b, 0x9d, 0xd4, 0x39, 0x31, 0xe7, 0x44, 0xe4, 0xd6,
	0x52, 0x45, 0xee, 0xdc, 0xb4, 0xc8, 0x89, 0xe8, 0x84, 0xc2, 0xdd, 0x7f, 0x36, 0x0a, 0x77, 0x6e,
	0x5a, 0xe1, 0x12, 0x25, 0x88, 0xf2, 0x36, 0x7c, 0x9a, 0xf2, 0x56, 0x4f, 0xc8, 0x5b, 0x62, 0xa3,
	0x42, 0x6d, 0x7b, 0xff, 0xc9, 0xda, 0xf6, 0xff, 0x27, 0x68, 0x9b, 0x48, 0x97, 0x2e, 0x6c, 0xaf,
	0xa2, 0x0a, 0x3b, 0x8e, 0xfb, 0xac, 0x13, 0xa1, 0xa2, 0x9d, 0x09, 0x7c, 0xb9, 0xdc, 0x21, 0x1a,
	0x74, 0xe7, 0x10, 0x09, 0x53, 0x08, 0xb0, 0xe3, 0xbb, 0x82, 0x5a, 0xfc, 0xc1, 0xd5, 0x62, 0x86,
	0x3e, 0x74, 0x88, 0x76, 0x44, 0xa9, 0xc0, 0x5f, 0x64, 0xd1, 0x71, 0x2e, 0x60, 0xdd, 0x5d, 0xa2,
	0x0e, 0x78, 0x05, 0x3f, 0x9e, 0x82, 0x0a, 0x2e, 0xce, 0xaa, 0x57, 0xed, 0xe5, 0xc0, 0x97, 0xe7,
	0xc3, 0x67, 0xa2, 0x0e, 0xc2, 0x2a, 0x4e, 0x89, 0x7a, 0x19, 0xa7, 0x53, 0xb6, 0xe7, 0xdd, 0x44,
	0x28, 0xbb, 0x77, 0x6b, 0xac, 0xd6, 0xb8, 0xa2, 0x9f, 0x78, 0x45, 0xaf, 0xcc, 0xd4, 0x93, 0xf6,
	0x73, 0x6c, 0x1c, 0xd9, 0x43, 0x5c, 0xcb, 0x42, 0xdc, 0x11, 0xa1, 0x10, 0xc9, 0x8b, 0x83, 0x94,
	0x0f, 0x50, 0x65, 0xd2, 0xd0, 0x67, 0x20, 0xa1, 0xdf, 0xe6, 0x50, 0x65, 0x32, 0xb8, 0xf8, 0x3c,
	0x2a, 0x58, 0x77, 0x4d, 0xea, 0x84, 0x09, 0x16, 0x02, 0x5f, 0x2e, 0xdc, 0x64, 0x86, 0x47, 0xbe,
	0x5c, 0xe2, 0x40, 0x65, 0x9b, 0x87, 0xe0, 0x55, 0x54, 0x85, 0x87, 0x2e, 0x51, 0x55, 0xea, 0xba,
	0x20, 0xa8, 0x79, 0x18, 0x34, 0x09, 0x20, 0x97, 0xc1, 0x2c, 0x02, 0x25, 0x2b, 0xb6, 0xe3, 0x35,
	0x54, 0xd3, 0x0c, 0x6b, 0x97, 0x18, 0x11, 0x3e, 0x0f, 0x78, 0xf6, 0x65, 0x50, 0xbd, 0x0a, 0x8e,
	0x69, 0x82, 0xaa, 0x26, 0x38, 0xb0, 0x86, 0x0a, 0xee, 0x1e, 0x71, 0x98, 0xb0, 0xce, 0x38, 0xa9,
	0xb7, 0x18, 0x1c, 0x74, 0x0e, 0x56, 0x0a, 0xaf, 0x89, 0x95, 0x02, 0xbf, 0xb2, 0x87, 0xaa, 0xa2,
	0xb0, 0xe0, 0x26, 0x2a, 0xf0, 0xb3, 0xc5, 0xbb, 0xc4, 0x4e, 0x7b, 0x01, 0xbc, 0xbf, 0xf9, 0x72,
	0x26, 0x81, 0x87, 0x30, 0xbc, 0x84, 0xf2, 0x03, 0x3a, 0x16, 0x6e, 0x9c, 0xfc, 0x35, 0x3a, 0x3e,
	0x1c, 0xcb, 0x42, 0x94, 0x2f, 0xb3, 0xa8, 0x96, 0xd0, 0x30, 0xfc, 0x69, 0x16, 0xcd, 0x0f, 0xe8,
	0x98, 0x1f, 0xe6, 0xae, 0x4d, 0x74, 0xb6, 0x37, 0x33, 0x5e, 0xe3, 0x22, 0x33, 0x6f, 0xb4, 0x68,
	0x49, 0x34, 0x7a, 0x20, 0x38, 0x94, 0xaf, 0x73, 0xa8, 0x1c, 0xcd, 0xc8, 0x3f, 0x39, 0x22, 0x0c,
	0xee, 0xed, 0xc5, 0xf0, 0xbc, 0x00, 0x67, 0xf6, 0x34, 0x78, 0x6c, 0x67, 0x95, 0x6a, 0x8e, 0x35,
	0xb2, 0xc3, 0x8b, 0x17, 0x2a, 0xbd, 0xca, 0x0c, 0x89, 0x4a, 0x21, 0x84, 0xa5, 0x82, 0x87, 0x28,
	0x55, 0x21, 0x4e, 0x05, 0x90, 0x94, 0x54, 0x5a, 0x6c, 0x57, 0x3e, 0xcb, 0xa2, 0x32, 0xbb, 0xf2,
	0xa1, 0x43, 0x97, 0x90, 0xc4, 0xce, 0x5f, 0x77, 0xe8, 0x1a, 0x96, 0xa9, 0x41, 0x9f, 0xe6, 0xe0,
	0x22, 0x45, 0xec, 0xac, 0xde, 0x00, 0xab, 0xc8, 0x84, 0x58, 0x34, 0x37, 0x4f, 0xb0, 0x06, 0xc7,
	0xe6, 0x92, 0xd8, 0xeb, 0xe9, 0x58, 0x6e, 0x56, 0xbe, 0x9f, 0x43, 0x92, 0x70, 0x77, 0xe3, 0x65,
	0x54, 0xa4, 0x26, 0xd9, 0x35, 0xf8, 0x9c, 0x96, 0xdb, 0x27, 0xd8, 0xad, 0xf4, 0x16, 0x58, 0x44,
	0x8a, 0x30, 0x08, 0xbf, 0x8e, 0xa4, 0x1e, 0x75, 0x55, 0x47, 0xb7, 0xd9, 0x5d, 0x12, 0x4e, 0x2b,
	0x74, 0x60, 0x23, 0x36, 0x27, 0x3a, 0x20, 0x84, 0xe3, 0x0b, 0xa8, 0xa4, 0x3a, 0x94, 0xfd, 0xec,
	0x0d, 0x3f, 0x93, 0xd8, 0x9c, 0x97, 0xd6, 0xb9, 0x49, 0x44, 0x45, 0x61, 0x13, 0x84, 0xe5, 0x84,
	0x1b, 0x14, 0x23, 0x2c, 0x67, 0x1a, 0x61, 0xc1, 0x38, 0x8d, 0x5c, 0xea, 0x74, 0xf7, 0x75, 0x57,
	0x67, 0xcb, 0x2a, 0xc0, 0xb2, 0xa0, 0xc4, 0x1d, 0x97, 0x3a, 0xb7, 0xb9, 0x39, 0x51, 0xe2, 0x28,
	0xb6, 0x33, 0xc5, 0x31, 0x88, 0xeb, 0x75, 0x87, 0x56, 0x4f, 0xef, 0xeb, 0xb4, 0x07, 0x5f, 0x48,
	0xfc, 0x07, 0x64, 0xf5, 0x3a, 0x71, 0xbd, 0x1b, 0xa1, 0x3d, 0x71, 0x10, 0x0c, 0xc1, 0x81, 0xef,
	0x21, 0x89, 0xdd, 0xdf, 0xba, 0xeb, 0xc2, 0x67, 0x4a, 0x09, 0xbe, 0x14, 0x2e, 0xcd, 0xfe, 0xa5,
	0xc0, 0x6b, 0xdf, 0x8a, 0x29, 0x13, 0xb5, 0x0b, 0xa9, 0xb0, 0x1a, 0x0a, 0x7f, 0x79, 0xd6, 0x94,
	0xd1, 0x74, 0xc2, 0x4d, 0x10, 0x5d, 0x1a, 0x93, 0x5c, 0x40, 0xae, 0x7c, 0x82, 0x2a, 0x13, 0x45,
	0x64, 0xdd, 0xf2, 0xa8, 0x49, 0x4c, 0x2f, 0x3a, 0x12, 0xd9, 0x58, 0x9f, 0x3b, 0xe0, 0x48, 0xd1,
	0x67, 0x4f, 0x70, 0xb0, 0xf9, 0xe3, 0xef, 0xe1, 0x2c, 0xc1, 0xfc, 0x71, 0x68, 0x62, 0xfe, 0x78,
	0x50, 0xfb, 0xe2, 0x83, 0x83, 0x46, 0xe6, 0x97, 0x83, 0x46, 0xe6, 0xe1, 0x41, 0x23, 0xf3, 0xfb,
	0x41, 0x23, 0xf3, 0xd7, 0x41, 0x23, 0x73, 0x3f, 0x68, 0x64, 0xbf, 0x0b, 0x1a, 0x99, 0x1f, 0x82,
	0x46, 0xe6, 0xe7, 0xa0, 0x91, 0x79, 0x10, 0x34, 0x32, 0x0f, 0x83, 0x46, 0xf6, 0x9b, 0x5f, 0x1b,
	0x99, 0xf7, 0x8a, 0x7c, 0x6d, 0xbb, 0x45, 0xf8, 0x1f, 0xe4, 0xa5, 0xbf, 0x03, 0x00, 0x00, 0xff,
	0xff, 0x79, 0x7c, 0x4c, 0x83, 0xac, 0x11, 0x00, 0x00,
}
