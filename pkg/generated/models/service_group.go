package models

// ServiceGroup

import "encoding/json"

// ServiceGroup
type ServiceGroup struct {
	ParentType                      string                    `json:"parent_type,omitempty"`
	FQName                          []string                  `json:"fq_name,omitempty"`
	IDPerms                         *IdPermsType              `json:"id_perms,omitempty"`
	Annotations                     *KeyValuePairs            `json:"annotations,omitempty"`
	Perms2                          *PermType2                `json:"perms2,omitempty"`
	UUID                            string                    `json:"uuid,omitempty"`
	ParentUUID                      string                    `json:"parent_uuid,omitempty"`
	ServiceGroupFirewallServiceList *FirewallServiceGroupType `json:"service_group_firewall_service_list,omitempty"`
	DisplayName                     string                    `json:"display_name,omitempty"`
}

// String returns json representation of the object
func (model *ServiceGroup) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakeServiceGroup makes ServiceGroup
func MakeServiceGroup() *ServiceGroup {
	return &ServiceGroup{
		//TODO(nati): Apply default
		Annotations: MakeKeyValuePairs(),
		Perms2:      MakePermType2(),
		UUID:        "",
		ParentUUID:  "",
		ParentType:  "",
		FQName:      []string{},
		IDPerms:     MakeIdPermsType(),
		ServiceGroupFirewallServiceList: MakeFirewallServiceGroupType(),
		DisplayName:                     "",
	}
}

// MakeServiceGroupSlice() makes a slice of ServiceGroup
func MakeServiceGroupSlice() []*ServiceGroup {
	return []*ServiceGroup{}
}
