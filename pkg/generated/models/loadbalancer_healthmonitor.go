package models

// LoadbalancerHealthmonitor

import "encoding/json"

// LoadbalancerHealthmonitor
type LoadbalancerHealthmonitor struct {
	FQName                              []string                       `json:"fq_name,omitempty"`
	UUID                                string                         `json:"uuid,omitempty"`
	ParentType                          string                         `json:"parent_type,omitempty"`
	LoadbalancerHealthmonitorProperties *LoadbalancerHealthmonitorType `json:"loadbalancer_healthmonitor_properties,omitempty"`
	IDPerms                             *IdPermsType                   `json:"id_perms,omitempty"`
	DisplayName                         string                         `json:"display_name,omitempty"`
	Annotations                         *KeyValuePairs                 `json:"annotations,omitempty"`
	Perms2                              *PermType2                     `json:"perms2,omitempty"`
	ParentUUID                          string                         `json:"parent_uuid,omitempty"`
}

// String returns json representation of the object
func (model *LoadbalancerHealthmonitor) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakeLoadbalancerHealthmonitor makes LoadbalancerHealthmonitor
func MakeLoadbalancerHealthmonitor() *LoadbalancerHealthmonitor {
	return &LoadbalancerHealthmonitor{
		//TODO(nati): Apply default
		FQName:                              []string{},
		UUID:                                "",
		ParentType:                          "",
		LoadbalancerHealthmonitorProperties: MakeLoadbalancerHealthmonitorType(),
		IDPerms:     MakeIdPermsType(),
		DisplayName: "",
		Annotations: MakeKeyValuePairs(),
		Perms2:      MakePermType2(),
		ParentUUID:  "",
	}
}

// MakeLoadbalancerHealthmonitorSlice() makes a slice of LoadbalancerHealthmonitor
func MakeLoadbalancerHealthmonitorSlice() []*LoadbalancerHealthmonitor {
	return []*LoadbalancerHealthmonitor{}
}
