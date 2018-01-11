package models

// LoadbalancerHealthmonitor

import "encoding/json"

// LoadbalancerHealthmonitor
type LoadbalancerHealthmonitor struct {
	ParentType                          string                         `json:"parent_type"`
	FQName                              []string                       `json:"fq_name"`
	IDPerms                             *IdPermsType                   `json:"id_perms"`
	DisplayName                         string                         `json:"display_name"`
	Perms2                              *PermType2                     `json:"perms2"`
	UUID                                string                         `json:"uuid"`
	LoadbalancerHealthmonitorProperties *LoadbalancerHealthmonitorType `json:"loadbalancer_healthmonitor_properties"`
	ParentUUID                          string                         `json:"parent_uuid"`
	Annotations                         *KeyValuePairs                 `json:"annotations"`
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
		Annotations:                         MakeKeyValuePairs(),
		LoadbalancerHealthmonitorProperties: MakeLoadbalancerHealthmonitorType(),
		ParentUUID:                          "",
		IDPerms:                             MakeIdPermsType(),
		DisplayName:                         "",
		Perms2:                              MakePermType2(),
		UUID:                                "",
		ParentType:                          "",
		FQName:                              []string{},
	}
}

// MakeLoadbalancerHealthmonitorSlice() makes a slice of LoadbalancerHealthmonitor
func MakeLoadbalancerHealthmonitorSlice() []*LoadbalancerHealthmonitor {
	return []*LoadbalancerHealthmonitor{}
}
