package models

// LoadbalancerHealthmonitor

import "encoding/json"

// LoadbalancerHealthmonitor
//proteus:generate
type LoadbalancerHealthmonitor struct {
	UUID                                string                         `json:"uuid,omitempty"`
	ParentUUID                          string                         `json:"parent_uuid,omitempty"`
	ParentType                          string                         `json:"parent_type,omitempty"`
	FQName                              []string                       `json:"fq_name,omitempty"`
	IDPerms                             *IdPermsType                   `json:"id_perms,omitempty"`
	DisplayName                         string                         `json:"display_name,omitempty"`
	Annotations                         *KeyValuePairs                 `json:"annotations,omitempty"`
	Perms2                              *PermType2                     `json:"perms2,omitempty"`
	LoadbalancerHealthmonitorProperties *LoadbalancerHealthmonitorType `json:"loadbalancer_healthmonitor_properties,omitempty"`
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
		UUID:        "",
		ParentUUID:  "",
		ParentType:  "",
		FQName:      []string{},
		IDPerms:     MakeIdPermsType(),
		DisplayName: "",
		Annotations: MakeKeyValuePairs(),
		Perms2:      MakePermType2(),
		LoadbalancerHealthmonitorProperties: MakeLoadbalancerHealthmonitorType(),
	}
}

// MakeLoadbalancerHealthmonitorSlice() makes a slice of LoadbalancerHealthmonitor
func MakeLoadbalancerHealthmonitorSlice() []*LoadbalancerHealthmonitor {
	return []*LoadbalancerHealthmonitor{}
}
