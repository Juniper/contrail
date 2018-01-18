package models

// LoadbalancerListener

import "encoding/json"

// LoadbalancerListener
type LoadbalancerListener struct {
	UUID                           string                    `json:"uuid,omitempty"`
	ParentType                     string                    `json:"parent_type,omitempty"`
	LoadbalancerListenerProperties *LoadbalancerListenerType `json:"loadbalancer_listener_properties,omitempty"`
	IDPerms                        *IdPermsType              `json:"id_perms,omitempty"`
	Perms2                         *PermType2                `json:"perms2,omitempty"`
	ParentUUID                     string                    `json:"parent_uuid,omitempty"`
	FQName                         []string                  `json:"fq_name,omitempty"`
	DisplayName                    string                    `json:"display_name,omitempty"`
	Annotations                    *KeyValuePairs            `json:"annotations,omitempty"`

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
		UUID:                           "",
		ParentType:                     "",
		LoadbalancerListenerProperties: MakeLoadbalancerListenerType(),
		Annotations:                    MakeKeyValuePairs(),
		Perms2:                         MakePermType2(),
		ParentUUID:                     "",
		FQName:                         []string{},
		DisplayName:                    "",
	}
}

// MakeLoadbalancerListenerSlice() makes a slice of LoadbalancerListener
func MakeLoadbalancerListenerSlice() []*LoadbalancerListener {
	return []*LoadbalancerListener{}
}
