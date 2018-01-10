package models

// LoadbalancerListener

import "encoding/json"

// LoadbalancerListener
type LoadbalancerListener struct {
	ParentType                     string                    `json:"parent_type"`
	FQName                         []string                  `json:"fq_name"`
	IDPerms                        *IdPermsType              `json:"id_perms"`
	Perms2                         *PermType2                `json:"perms2"`
	LoadbalancerListenerProperties *LoadbalancerListenerType `json:"loadbalancer_listener_properties"`
	ParentUUID                     string                    `json:"parent_uuid"`
	DisplayName                    string                    `json:"display_name"`
	Annotations                    *KeyValuePairs            `json:"annotations"`
	UUID                           string                    `json:"uuid"`

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
		IDPerms: MakeIdPermsType(),
		Perms2:  MakePermType2(),
		LoadbalancerListenerProperties: MakeLoadbalancerListenerType(),
		ParentType:                     "",
		FQName:                         []string{},
		Annotations:                    MakeKeyValuePairs(),
		UUID:                           "",
		ParentUUID:                     "",
		DisplayName:                    "",
	}
}

// InterfaceToLoadbalancerListener makes LoadbalancerListener from interface
func InterfaceToLoadbalancerListener(iData interface{}) *LoadbalancerListener {
	data := iData.(map[string]interface{})
	return &LoadbalancerListener{
		UUID: data["uuid"].(string),

		//{"type":"string"}
		ParentUUID: data["parent_uuid"].(string),

		//{"type":"string"}
		DisplayName: data["display_name"].(string),

		//{"type":"string"}
		Annotations: InterfaceToKeyValuePairs(data["annotations"]),

		//{"type":"object","properties":{"key_value_pair":{"type":"array","item":{"type":"object","properties":{"key":{"type":"string"},"value":{"type":"string"}}}}}}
		Perms2: InterfaceToPermType2(data["perms2"]),

		//{"type":"object","properties":{"global_access":{"type":"integer","minimum":0,"maximum":7},"owner":{"type":"string"},"owner_access":{"type":"integer","minimum":0,"maximum":7},"share":{"type":"array","item":{"type":"object","properties":{"tenant":{"type":"string"},"tenant_access":{"type":"integer","minimum":0,"maximum":7}}}}}}
		LoadbalancerListenerProperties: InterfaceToLoadbalancerListenerType(data["loadbalancer_listener_properties"]),

		//{"type":"object","properties":{"admin_state":{"type":"boolean"},"connection_limit":{"type":"integer"},"default_tls_container":{"type":"string"},"protocol":{"type":"string","enum":["HTTP","HTTPS","TCP","UDP","TERMINATED_HTTPS"]},"protocol_port":{"type":"integer"},"sni_containers":{"type":"array","item":{"type":"string"}}}}
		ParentType: data["parent_type"].(string),

		//{"type":"string"}
		FQName: data["fq_name"].([]string),

		//{"type":"array","item":{"type":"string"}}
		IDPerms: InterfaceToIdPermsType(data["id_perms"]),

		//{"type":"object","properties":{"created":{"type":"string"},"creator":{"type":"string"},"description":{"type":"string"},"enable":{"type":"boolean"},"last_modified":{"type":"string"},"permissions":{"type":"object","properties":{"group":{"type":"string"},"group_access":{"type":"integer","minimum":0,"maximum":7},"other_access":{"type":"integer","minimum":0,"maximum":7},"owner":{"type":"string"},"owner_access":{"type":"integer","minimum":0,"maximum":7}}},"user_visible":{"type":"boolean"}}}

	}
}

// InterfaceToLoadbalancerListenerSlice makes a slice of LoadbalancerListener from interface
func InterfaceToLoadbalancerListenerSlice(data interface{}) []*LoadbalancerListener {
	list := data.([]interface{})
	result := MakeLoadbalancerListenerSlice()
	for _, item := range list {
		result = append(result, InterfaceToLoadbalancerListener(item))
	}
	return result
}

// MakeLoadbalancerListenerSlice() makes a slice of LoadbalancerListener
func MakeLoadbalancerListenerSlice() []*LoadbalancerListener {
	return []*LoadbalancerListener{}
}
