package models

// ServiceHealthCheck

import "encoding/json"

// ServiceHealthCheck
type ServiceHealthCheck struct {
	Annotations                  *KeyValuePairs          `json:"annotations,omitempty"`
	ServiceHealthCheckProperties *ServiceHealthCheckType `json:"service_health_check_properties,omitempty"`
	ParentType                   string                  `json:"parent_type,omitempty"`
	IDPerms                      *IdPermsType            `json:"id_perms,omitempty"`
	DisplayName                  string                  `json:"display_name,omitempty"`
	FQName                       []string                `json:"fq_name,omitempty"`
	Perms2                       *PermType2              `json:"perms2,omitempty"`
	UUID                         string                  `json:"uuid,omitempty"`
	ParentUUID                   string                  `json:"parent_uuid,omitempty"`

	ServiceInstanceRefs []*ServiceHealthCheckServiceInstanceRef `json:"service_instance_refs,omitempty"`
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
		FQName:                       []string{},
		Perms2:                       MakePermType2(),
		UUID:                         "",
		ParentUUID:                   "",
		Annotations:                  MakeKeyValuePairs(),
		ServiceHealthCheckProperties: MakeServiceHealthCheckType(),
		ParentType:                   "",
		IDPerms:                      MakeIdPermsType(),
		DisplayName:                  "",
	}
}

// MakeServiceHealthCheckSlice() makes a slice of ServiceHealthCheck
func MakeServiceHealthCheckSlice() []*ServiceHealthCheck {
	return []*ServiceHealthCheck{}
}
