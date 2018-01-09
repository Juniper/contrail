package models

// LoadbalancerHealthmonitor

import "encoding/json"

// LoadbalancerHealthmonitor
type LoadbalancerHealthmonitor struct {
	FQName                              []string                       `json:"fq_name"`
	DisplayName                         string                         `json:"display_name"`
	UUID                                string                         `json:"uuid"`
	ParentType                          string                         `json:"parent_type"`
	Perms2                              *PermType2                     `json:"perms2"`
	ParentUUID                          string                         `json:"parent_uuid"`
	IDPerms                             *IdPermsType                   `json:"id_perms"`
	LoadbalancerHealthmonitorProperties *LoadbalancerHealthmonitorType `json:"loadbalancer_healthmonitor_properties"`
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
		Annotations: MakeKeyValuePairs(),
		Perms2:      MakePermType2(),
		ParentUUID:  "",
		IDPerms:     MakeIdPermsType(),
		LoadbalancerHealthmonitorProperties: MakeLoadbalancerHealthmonitorType(),
		ParentType:                          "",
		FQName:                              []string{},
		DisplayName:                         "",
		UUID:                                "",
	}
}

// InterfaceToLoadbalancerHealthmonitor makes LoadbalancerHealthmonitor from interface
func InterfaceToLoadbalancerHealthmonitor(iData interface{}) *LoadbalancerHealthmonitor {
	data := iData.(map[string]interface{})
	return &LoadbalancerHealthmonitor{
		Annotations: InterfaceToKeyValuePairs(data["annotations"]),

		//{"type":"object","properties":{"key_value_pair":{"type":"array","item":{"type":"object","properties":{"key":{"type":"string"},"value":{"type":"string"}}}}}}
		Perms2: InterfaceToPermType2(data["perms2"]),

		//{"type":"object","properties":{"global_access":{"type":"integer","minimum":0,"maximum":7},"owner":{"type":"string"},"owner_access":{"type":"integer","minimum":0,"maximum":7},"share":{"type":"array","item":{"type":"object","properties":{"tenant":{"type":"string"},"tenant_access":{"type":"integer","minimum":0,"maximum":7}}}}}}
		ParentUUID: data["parent_uuid"].(string),

		//{"type":"string"}
		IDPerms: InterfaceToIdPermsType(data["id_perms"]),

		//{"type":"object","properties":{"created":{"type":"string"},"creator":{"type":"string"},"description":{"type":"string"},"enable":{"type":"boolean"},"last_modified":{"type":"string"},"permissions":{"type":"object","properties":{"group":{"type":"string"},"group_access":{"type":"integer","minimum":0,"maximum":7},"other_access":{"type":"integer","minimum":0,"maximum":7},"owner":{"type":"string"},"owner_access":{"type":"integer","minimum":0,"maximum":7}}},"user_visible":{"type":"boolean"}}}
		LoadbalancerHealthmonitorProperties: InterfaceToLoadbalancerHealthmonitorType(data["loadbalancer_healthmonitor_properties"]),

		//{"description":"Configuration parameters for health monitor like type, method, retries etc.","type":"object","properties":{"admin_state":{"type":"boolean"},"delay":{"type":"integer"},"expected_codes":{"type":"string"},"http_method":{"type":"string"},"max_retries":{"type":"integer"},"monitor_type":{"type":"string","enum":["PING","TCP","HTTP","HTTPS"]},"timeout":{"type":"integer"},"url_path":{"type":"string"}}}
		ParentType: data["parent_type"].(string),

		//{"type":"string"}
		FQName: data["fq_name"].([]string),

		//{"type":"array","item":{"type":"string"}}
		DisplayName: data["display_name"].(string),

		//{"type":"string"}
		UUID: data["uuid"].(string),

		//{"type":"string"}

	}
}

// InterfaceToLoadbalancerHealthmonitorSlice makes a slice of LoadbalancerHealthmonitor from interface
func InterfaceToLoadbalancerHealthmonitorSlice(data interface{}) []*LoadbalancerHealthmonitor {
	list := data.([]interface{})
	result := MakeLoadbalancerHealthmonitorSlice()
	for _, item := range list {
		result = append(result, InterfaceToLoadbalancerHealthmonitor(item))
	}
	return result
}

// MakeLoadbalancerHealthmonitorSlice() makes a slice of LoadbalancerHealthmonitor
func MakeLoadbalancerHealthmonitorSlice() []*LoadbalancerHealthmonitor {
	return []*LoadbalancerHealthmonitor{}
}
