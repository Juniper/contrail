package models

// ServiceHealthCheck

import "encoding/json"

// ServiceHealthCheck
type ServiceHealthCheck struct {
	ParentType                   string                  `json:"parent_type"`
	FQName                       []string                `json:"fq_name"`
	DisplayName                  string                  `json:"display_name"`
	Perms2                       *PermType2              `json:"perms2"`
	ServiceHealthCheckProperties *ServiceHealthCheckType `json:"service_health_check_properties"`
	IDPerms                      *IdPermsType            `json:"id_perms"`
	Annotations                  *KeyValuePairs          `json:"annotations"`
	UUID                         string                  `json:"uuid"`
	ParentUUID                   string                  `json:"parent_uuid"`

	ServiceInstanceRefs []*ServiceHealthCheckServiceInstanceRef `json:"service_instance_refs"`
}

// ServiceHealthCheckServiceInstanceRef references each other
type ServiceHealthCheckServiceInstanceRef struct {
	UUID string   `json:"uuid"`
	To   []string `json:"to"` //FQDN

	Attr *ServiceInterfaceTag
}

// String returns json representation of the object
func (model *ServiceHealthCheck) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakeServiceHealthCheck makes ServiceHealthCheck
func MakeServiceHealthCheck() *ServiceHealthCheck {
	return &ServiceHealthCheck{
		//TODO(nati): Apply default
		Perms2: MakePermType2(),
		ServiceHealthCheckProperties: MakeServiceHealthCheckType(),
		ParentType:                   "",
		FQName:                       []string{},
		DisplayName:                  "",
		ParentUUID:                   "",
		IDPerms:                      MakeIdPermsType(),
		Annotations:                  MakeKeyValuePairs(),
		UUID:                         "",
	}
}

// InterfaceToServiceHealthCheck makes ServiceHealthCheck from interface
func InterfaceToServiceHealthCheck(iData interface{}) *ServiceHealthCheck {
	data := iData.(map[string]interface{})
	return &ServiceHealthCheck{
		ParentUUID: data["parent_uuid"].(string),

		//{"type":"string"}
		IDPerms: InterfaceToIdPermsType(data["id_perms"]),

		//{"type":"object","properties":{"created":{"type":"string"},"creator":{"type":"string"},"description":{"type":"string"},"enable":{"type":"boolean"},"last_modified":{"type":"string"},"permissions":{"type":"object","properties":{"group":{"type":"string"},"group_access":{"type":"integer","minimum":0,"maximum":7},"other_access":{"type":"integer","minimum":0,"maximum":7},"owner":{"type":"string"},"owner_access":{"type":"integer","minimum":0,"maximum":7}}},"user_visible":{"type":"boolean"}}}
		Annotations: InterfaceToKeyValuePairs(data["annotations"]),

		//{"type":"object","properties":{"key_value_pair":{"type":"array","item":{"type":"object","properties":{"key":{"type":"string"},"value":{"type":"string"}}}}}}
		UUID: data["uuid"].(string),

		//{"type":"string"}
		ServiceHealthCheckProperties: InterfaceToServiceHealthCheckType(data["service_health_check_properties"]),

		//{"description":"Service health check has following fields.","type":"object","properties":{"delay":{"type":"integer"},"delayUsecs":{"type":"integer"},"enabled":{"type":"boolean"},"expected_codes":{"type":"string"},"health_check_type":{"type":"string","enum":["link-local","end-to-end","segment"]},"http_method":{"type":"string"},"max_retries":{"type":"integer"},"monitor_type":{"type":"string","enum":["PING","HTTP","BFD"]},"timeout":{"type":"integer"},"timeoutUsecs":{"type":"integer"},"url_path":{"type":"string"}}}
		ParentType: data["parent_type"].(string),

		//{"type":"string"}
		FQName: data["fq_name"].([]string),

		//{"type":"array","item":{"type":"string"}}
		DisplayName: data["display_name"].(string),

		//{"type":"string"}
		Perms2: InterfaceToPermType2(data["perms2"]),

		//{"type":"object","properties":{"global_access":{"type":"integer","minimum":0,"maximum":7},"owner":{"type":"string"},"owner_access":{"type":"integer","minimum":0,"maximum":7},"share":{"type":"array","item":{"type":"object","properties":{"tenant":{"type":"string"},"tenant_access":{"type":"integer","minimum":0,"maximum":7}}}}}}

	}
}

// InterfaceToServiceHealthCheckSlice makes a slice of ServiceHealthCheck from interface
func InterfaceToServiceHealthCheckSlice(data interface{}) []*ServiceHealthCheck {
	list := data.([]interface{})
	result := MakeServiceHealthCheckSlice()
	for _, item := range list {
		result = append(result, InterfaceToServiceHealthCheck(item))
	}
	return result
}

// MakeServiceHealthCheckSlice() makes a slice of ServiceHealthCheck
func MakeServiceHealthCheckSlice() []*ServiceHealthCheck {
	return []*ServiceHealthCheck{}
}
