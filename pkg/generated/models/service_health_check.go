package models

// ServiceHealthCheck

// ServiceHealthCheck
//proteus:generate
type ServiceHealthCheck struct {
	UUID                         string                  `json:"uuid,omitempty"`
	ParentUUID                   string                  `json:"parent_uuid,omitempty"`
	ParentType                   string                  `json:"parent_type,omitempty"`
	FQName                       []string                `json:"fq_name,omitempty"`
	IDPerms                      *IdPermsType            `json:"id_perms,omitempty"`
	DisplayName                  string                  `json:"display_name,omitempty"`
	Annotations                  *KeyValuePairs          `json:"annotations,omitempty"`
	Perms2                       *PermType2              `json:"perms2,omitempty"`
	ServiceHealthCheckProperties *ServiceHealthCheckType `json:"service_health_check_properties,omitempty"`

	ServiceInstanceRefs []*ServiceHealthCheckServiceInstanceRef `json:"service_instance_refs,omitempty"`
}

// ServiceHealthCheckServiceInstanceRef references each other
type ServiceHealthCheckServiceInstanceRef struct {
	UUID string   `json:"uuid"`
	To   []string `json:"to"` //FQDN

	Attr *ServiceInterfaceTag
}

// MakeServiceHealthCheck makes ServiceHealthCheck
func MakeServiceHealthCheck() *ServiceHealthCheck {
	return &ServiceHealthCheck{
		//TODO(nati): Apply default
		UUID:        "",
		ParentUUID:  "",
		ParentType:  "",
		FQName:      []string{},
		IDPerms:     MakeIdPermsType(),
		DisplayName: "",
		Annotations: MakeKeyValuePairs(),
		Perms2:      MakePermType2(),
		ServiceHealthCheckProperties: MakeServiceHealthCheckType(),
	}
}

// MakeServiceHealthCheckSlice() makes a slice of ServiceHealthCheck
func MakeServiceHealthCheckSlice() []*ServiceHealthCheck {
	return []*ServiceHealthCheck{}
}
