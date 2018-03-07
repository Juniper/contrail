package models

import (
	"github.com/Juniper/contrail/pkg/common"
)

//To skip import error.
var _ = common.OPERATION

// MakeCustomerAttachment makes CustomerAttachment
// nolint
func MakeCustomerAttachment() *CustomerAttachment {
	return &CustomerAttachment{
		//TODO(nati): Apply default
		UUID:                 "",
		ParentUUID:           "",
		ParentType:           "",
		FQName:               []string{},
		IDPerms:              MakeIdPermsType(),
		DisplayName:          "",
		Annotations:          MakeKeyValuePairs(),
		Perms2:               MakePermType2(),
		ConfigurationVersion: 0,
	}
}

// MakeCustomerAttachment makes CustomerAttachment
// nolint
func InterfaceToCustomerAttachment(i interface{}) *CustomerAttachment {
	m, ok := i.(map[string]interface{})
	_ = m
	if !ok {
		return nil
	}
	return &CustomerAttachment{
		//TODO(nati): Apply default
		UUID:                 common.InterfaceToString(m["uuid"]),
		ParentUUID:           common.InterfaceToString(m["parent_uuid"]),
		ParentType:           common.InterfaceToString(m["parent_type"]),
		FQName:               common.InterfaceToStringList(m["fq_name"]),
		IDPerms:              InterfaceToIdPermsType(m["id_perms"]),
		DisplayName:          common.InterfaceToString(m["display_name"]),
		Annotations:          InterfaceToKeyValuePairs(m["annotations"]),
		Perms2:               InterfaceToPermType2(m["perms2"]),
		ConfigurationVersion: common.InterfaceToInt64(m["configuration_version"]),
	}
}

// MakeCustomerAttachmentSlice() makes a slice of CustomerAttachment
// nolint
func MakeCustomerAttachmentSlice() []*CustomerAttachment {
	return []*CustomerAttachment{}
}

// InterfaceToCustomerAttachmentSlice() makes a slice of CustomerAttachment
// nolint
func InterfaceToCustomerAttachmentSlice(i interface{}) []*CustomerAttachment {
	list := common.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*CustomerAttachment{}
	for _, item := range list {
		result = append(result, InterfaceToCustomerAttachment(item))
	}
	return result
}
