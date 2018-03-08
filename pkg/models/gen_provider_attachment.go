package models

import (
	"github.com/Juniper/contrail/pkg/common"
)

//To skip import error.
var _ = common.OPERATION

// MakeProviderAttachment makes ProviderAttachment
// nolint
func MakeProviderAttachment() *ProviderAttachment {
	return &ProviderAttachment{
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

// MakeProviderAttachment makes ProviderAttachment
// nolint
func InterfaceToProviderAttachment(i interface{}) *ProviderAttachment {
	m, ok := i.(map[string]interface{})
	_ = m
	if !ok {
		return nil
	}
	return &ProviderAttachment{
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

		VirtualRouterRefs: InterfaceToProviderAttachmentVirtualRouterRefs(m["virtual_router_refs"]),
	}
}

func InterfaceToProviderAttachmentVirtualRouterRefs(i interface{}) []*ProviderAttachmentVirtualRouterRef {
	list, ok := i.([]interface{})
	if !ok {
		return nil
	}
	result := []*ProviderAttachmentVirtualRouterRef{}
	for _, item := range list {
		m, ok := item.(map[string]interface{})
		_ = m
		if !ok {
			return nil
		}
		result = append(result, &ProviderAttachmentVirtualRouterRef{
			UUID: common.InterfaceToString(m["uuid"]),
			To:   common.InterfaceToStringList(m["to"]),
		})
	}

	return result
}

// MakeProviderAttachmentSlice() makes a slice of ProviderAttachment
// nolint
func MakeProviderAttachmentSlice() []*ProviderAttachment {
	return []*ProviderAttachment{}
}

// InterfaceToProviderAttachmentSlice() makes a slice of ProviderAttachment
// nolint
func InterfaceToProviderAttachmentSlice(i interface{}) []*ProviderAttachment {
	list := common.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*ProviderAttachment{}
	for _, item := range list {
		result = append(result, InterfaceToProviderAttachment(item))
	}
	return result
}
