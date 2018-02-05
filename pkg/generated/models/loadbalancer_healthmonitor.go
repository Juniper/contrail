package models

// LoadbalancerHealthmonitor

import "encoding/json"

// LoadbalancerHealthmonitor
type LoadbalancerHealthmonitor struct {
	ParentType                          string                         `json:"parent_type,omitempty"`
	DisplayName                         string                         `json:"display_name,omitempty"`
	Annotations                         *KeyValuePairs                 `json:"annotations,omitempty"`
	UUID                                string                         `json:"uuid,omitempty"`
	FQName                              []string                       `json:"fq_name,omitempty"`
	IDPerms                             *IdPermsType                   `json:"id_perms,omitempty"`
	LoadbalancerHealthmonitorProperties *LoadbalancerHealthmonitorType `json:"loadbalancer_healthmonitor_properties,omitempty"`
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
		LoadbalancerHealthmonitorProperties: MakeLoadbalancerHealthmonitorType(),
		Perms2:      MakePermType2(),
		ParentUUID:  "",
		FQName:      []string{},
		IDPerms:     MakeIdPermsType(),
		DisplayName: "",
		Annotations: MakeKeyValuePairs(),
		UUID:        "",
		ParentType:  "",
	}
}

// MakeLoadbalancerHealthmonitorSlice() makes a slice of LoadbalancerHealthmonitor
func MakeLoadbalancerHealthmonitorSlice() []*LoadbalancerHealthmonitor {
	return []*LoadbalancerHealthmonitor{}
}
