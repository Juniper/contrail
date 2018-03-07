package models

import (
	"github.com/Juniper/contrail/pkg/common"
)

//To skip import error.
var _ = common.OPERATION

// MakeServiceTemplate makes ServiceTemplate
// nolint
func MakeServiceTemplate() *ServiceTemplate {
	return &ServiceTemplate{
		//TODO(nati): Apply default
		UUID:                      "",
		ParentUUID:                "",
		ParentType:                "",
		FQName:                    []string{},
		IDPerms:                   MakeIdPermsType(),
		DisplayName:               "",
		Annotations:               MakeKeyValuePairs(),
		Perms2:                    MakePermType2(),
		ConfigurationVersion:      0,
		ServiceTemplateProperties: MakeServiceTemplateType(),
	}
}

// MakeServiceTemplate makes ServiceTemplate
// nolint
func InterfaceToServiceTemplate(i interface{}) *ServiceTemplate {
	m, ok := i.(map[string]interface{})
	_ = m
	if !ok {
		return nil
	}
	return &ServiceTemplate{
		//TODO(nati): Apply default
		UUID:                      common.InterfaceToString(m["uuid"]),
		ParentUUID:                common.InterfaceToString(m["parent_uuid"]),
		ParentType:                common.InterfaceToString(m["parent_type"]),
		FQName:                    common.InterfaceToStringList(m["fq_name"]),
		IDPerms:                   InterfaceToIdPermsType(m["id_perms"]),
		DisplayName:               common.InterfaceToString(m["display_name"]),
		Annotations:               InterfaceToKeyValuePairs(m["annotations"]),
		Perms2:                    InterfaceToPermType2(m["perms2"]),
		ConfigurationVersion:      common.InterfaceToInt64(m["configuration_version"]),
		ServiceTemplateProperties: InterfaceToServiceTemplateType(m["service_template_properties"]),

		ServiceApplianceSetRefs: InterfaceToServiceTemplateServiceApplianceSetRefs(m["service_appliance_set_refs"]),
	}
}

func InterfaceToServiceTemplateServiceApplianceSetRefs(i interface{}) []*ServiceTemplateServiceApplianceSetRef {
	list, ok := i.([]interface{})
	if !ok {
		return nil
	}
	result := []*ServiceTemplateServiceApplianceSetRef{}
	for _, item := range list {
		m, ok := item.(map[string]interface{})
		_ = m
		if !ok {
			return nil
		}
		result = append(result, &ServiceTemplateServiceApplianceSetRef{
			UUID: common.InterfaceToString(m["uuid"]),
			To:   common.InterfaceToStringList(m["to"]),
		})
	}

	return result
}

// MakeServiceTemplateSlice() makes a slice of ServiceTemplate
// nolint
func MakeServiceTemplateSlice() []*ServiceTemplate {
	return []*ServiceTemplate{}
}

// InterfaceToServiceTemplateSlice() makes a slice of ServiceTemplate
// nolint
func InterfaceToServiceTemplateSlice(i interface{}) []*ServiceTemplate {
	list := common.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*ServiceTemplate{}
	for _, item := range list {
		result = append(result, InterfaceToServiceTemplate(item))
	}
	return result
}
