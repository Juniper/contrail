package models

// ServiceGroup

import "encoding/json"

// ServiceGroup
type ServiceGroup struct {
	DisplayName                     string                    `json:"display_name,omitempty"`
	Annotations                     *KeyValuePairs            `json:"annotations,omitempty"`
	UUID                            string                    `json:"uuid,omitempty"`
	ParentUUID                      string                    `json:"parent_uuid,omitempty"`
	ParentType                      string                    `json:"parent_type,omitempty"`
	ServiceGroupFirewallServiceList *FirewallServiceGroupType `json:"service_group_firewall_service_list,omitempty"`
	IDPerms                         *IdPermsType              `json:"id_perms,omitempty"`
	Perms2                          *PermType2                `json:"perms2,omitempty"`
	FQName                          []string                  `json:"fq_name,omitempty"`
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
		ServiceGroupFirewallServiceList: MakeFirewallServiceGroupType(),
		IDPerms:     MakeIdPermsType(),
		Perms2:      MakePermType2(),
		FQName:      []string{},
		DisplayName: "",
		Annotations: MakeKeyValuePairs(),
		UUID:        "",
		ParentUUID:  "",
		ParentType:  "",
	}
}

// MakeServiceGroupSlice() makes a slice of ServiceGroup
func MakeServiceGroupSlice() []*ServiceGroup {
	return []*ServiceGroup{}
}
