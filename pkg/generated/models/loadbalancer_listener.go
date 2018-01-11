package models

// LoadbalancerListener

import "encoding/json"

// LoadbalancerListener
type LoadbalancerListener struct {
	LoadbalancerListenerProperties *LoadbalancerListenerType `json:"loadbalancer_listener_properties"`
	ParentType                     string                    `json:"parent_type"`
	FQName                         []string                  `json:"fq_name"`
	IDPerms                        *IdPermsType              `json:"id_perms"`
	DisplayName                    string                    `json:"display_name"`
	Annotations                    *KeyValuePairs            `json:"annotations"`
	UUID                           string                    `json:"uuid"`
	ParentUUID                     string                    `json:"parent_uuid"`
	Perms2                         *PermType2                `json:"perms2"`

	LoadbalancerRefs []*LoadbalancerListenerLoadbalancerRef `json:"loadbalancer_refs"`
}

// LoadbalancerListenerLoadbalancerRef references each other
type LoadbalancerListenerLoadbalancerRef struct {
	UUID string   `json:"uuid"`
	To   []string `json:"to"` //FQDN

}

// String returns json representation of the object
func (model *LoadbalancerListener) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakeLoadbalancerListener makes LoadbalancerListener
func MakeLoadbalancerListener() *LoadbalancerListener {
	return &LoadbalancerListener{
		//TODO(nati): Apply default
		UUID:       "",
		ParentUUID: "",
		Perms2:     MakePermType2(),
		LoadbalancerListenerProperties: MakeLoadbalancerListenerType(),
		ParentType:                     "",
		FQName:                         []string{},
		IDPerms:                        MakeIdPermsType(),
		DisplayName:                    "",
		Annotations:                    MakeKeyValuePairs(),
	}
}

// MakeLoadbalancerListenerSlice() makes a slice of LoadbalancerListener
func MakeLoadbalancerListenerSlice() []*LoadbalancerListener {
	return []*LoadbalancerListener{}
}
