package models

// ServiceGroup

import "encoding/json"

// ServiceGroup
type ServiceGroup struct {
	ServiceGroupFirewallServiceList *FirewallServiceGroupType `json:"service_group_firewall_service_list,omitempty"`
	Annotations                     *KeyValuePairs            `json:"annotations,omitempty"`
	ParentType                      string                    `json:"parent_type,omitempty"`
	IDPerms                         *IdPermsType              `json:"id_perms,omitempty"`
	DisplayName                     string                    `json:"display_name,omitempty"`
	Perms2                          *PermType2                `json:"perms2,omitempty"`
	UUID                            string                    `json:"uuid,omitempty"`
	ParentUUID                      string                    `json:"parent_uuid,omitempty"`
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
		Annotations: MakeKeyValuePairs(),
		ParentType:  "",
		IDPerms:     MakeIdPermsType(),
		ServiceGroupFirewallServiceList: MakeFirewallServiceGroupType(),
		Perms2:      MakePermType2(),
		UUID:        "",
		ParentUUID:  "",
		FQName:      []string{},
		DisplayName: "",
	}
}

// MakeServiceGroupSlice() makes a slice of ServiceGroup
func MakeServiceGroupSlice() []*ServiceGroup {
	return []*ServiceGroup{}
}
