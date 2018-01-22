package models

// LoadbalancerListener

import "encoding/json"

// LoadbalancerListener
type LoadbalancerListener struct {
	LoadbalancerListenerProperties *LoadbalancerListenerType `json:"loadbalancer_listener_properties,omitempty"`
	ParentUUID                     string                    `json:"parent_uuid,omitempty"`
	ParentType                     string                    `json:"parent_type,omitempty"`
	DisplayName                    string                    `json:"display_name,omitempty"`
	FQName                         []string                  `json:"fq_name,omitempty"`
	IDPerms                        *IdPermsType              `json:"id_perms,omitempty"`
	Annotations                    *KeyValuePairs            `json:"annotations,omitempty"`
	Perms2                         *PermType2                `json:"perms2,omitempty"`
	UUID                           string                    `json:"uuid,omitempty"`

	LoadbalancerRefs []*LoadbalancerListenerLoadbalancerRef `json:"loadbalancer_refs,omitempty"`
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
		IDPerms:                        MakeIdPermsType(),
		Annotations:                    MakeKeyValuePairs(),
		Perms2:                         MakePermType2(),
		UUID:                           "",
		FQName:                         []string{},
		ParentUUID:                     "",
		ParentType:                     "",
		DisplayName:                    "",
		LoadbalancerListenerProperties: MakeLoadbalancerListenerType(),
	}
}

// MakeLoadbalancerListenerSlice() makes a slice of LoadbalancerListener
func MakeLoadbalancerListenerSlice() []*LoadbalancerListener {
	return []*LoadbalancerListener{}
}
