package models

// ServiceGroup

import "encoding/json"

// ServiceGroup
type ServiceGroup struct {
	FQName                          []string                  `json:"fq_name"`
	IDPerms                         *IdPermsType              `json:"id_perms"`
	DisplayName                     string                    `json:"display_name"`
	Annotations                     *KeyValuePairs            `json:"annotations"`
	Perms2                          *PermType2                `json:"perms2"`
	ServiceGroupFirewallServiceList *FirewallServiceGroupType `json:"service_group_firewall_service_list"`
	UUID                            string                    `json:"uuid"`
	ParentUUID                      string                    `json:"parent_uuid"`
	ParentType                      string                    `json:"parent_type"`
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
		IDPerms:     MakeIdPermsType(),
		DisplayName: "",
		Annotations: MakeKeyValuePairs(),
		Perms2:      MakePermType2(),
		ServiceGroupFirewallServiceList: MakeFirewallServiceGroupType(),
		UUID:       "",
		FQName:     []string{},
		ParentUUID: "",
		ParentType: "",
	}
}

// MakeServiceGroupSlice() makes a slice of ServiceGroup
func MakeServiceGroupSlice() []*ServiceGroup {
	return []*ServiceGroup{}
}
