package models

// LoadbalancerHealthmonitor

import "encoding/json"

// LoadbalancerHealthmonitor
type LoadbalancerHealthmonitor struct {
	Perms2                              *PermType2                     `json:"perms2,omitempty"`
	UUID                                string                         `json:"uuid,omitempty"`
	FQName                              []string                       `json:"fq_name,omitempty"`
	IDPerms                             *IdPermsType                   `json:"id_perms,omitempty"`
	Annotations                         *KeyValuePairs                 `json:"annotations,omitempty"`
	ParentUUID                          string                         `json:"parent_uuid,omitempty"`
	ParentType                          string                         `json:"parent_type,omitempty"`
	LoadbalancerHealthmonitorProperties *LoadbalancerHealthmonitorType `json:"loadbalancer_healthmonitor_properties,omitempty"`
	DisplayName                         string                         `json:"display_name,omitempty"`
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
		LoadbalancerHealthmonitorProperties: MakeLoadbalancerHealthmonitorType(),
		DisplayName:                         "",
		ParentUUID:                          "",
		ParentType:                          "",
		IDPerms:                             MakeIdPermsType(),
		Annotations:                         MakeKeyValuePairs(),
		Perms2:                              MakePermType2(),
		UUID:                                "",
		FQName:                              []string{},
	}
}

// MakeLoadbalancerHealthmonitorSlice() makes a slice of LoadbalancerHealthmonitor
func MakeLoadbalancerHealthmonitorSlice() []*LoadbalancerHealthmonitor {
	return []*LoadbalancerHealthmonitor{}
}
