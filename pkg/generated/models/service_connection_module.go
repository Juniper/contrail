package models

// ServiceConnectionModule

import "encoding/json"

// ServiceConnectionModule
type ServiceConnectionModule struct {
	ServiceType ServiceConnectionType `json:"service_type"`
	E2Service   E2servicetype         `json:"e2_service"`
	DisplayName string                `json:"display_name"`
	UUID        string                `json:"uuid"`
	Annotations *KeyValuePairs        `json:"annotations"`
	Perms2      *PermType2            `json:"perms2"`
	ParentUUID  string                `json:"parent_uuid"`
	ParentType  string                `json:"parent_type"`
	FQName      []string              `json:"fq_name"`
	IDPerms     *IdPermsType          `json:"id_perms"`

	ServiceObjectRefs []*ServiceConnectionModuleServiceObjectRef `json:"service_object_refs"`
}

// ServiceConnectionModuleServiceObjectRef references each other
type ServiceConnectionModuleServiceObjectRef struct {
	UUID string   `json:"uuid"`
	To   []string `json:"to"` //FQDN

}

// String returns json representation of the object
func (model *ServiceConnectionModule) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakeServiceConnectionModule makes ServiceConnectionModule
func MakeServiceConnectionModule() *ServiceConnectionModule {
	return &ServiceConnectionModule{
		//TODO(nati): Apply default
		ServiceType: MakeServiceConnectionType(),
		E2Service:   MakeE2servicetype(),
		DisplayName: "",
		UUID:        "",
		Annotations: MakeKeyValuePairs(),
		Perms2:      MakePermType2(),
		ParentUUID:  "",
		ParentType:  "",
		FQName:      []string{},
		IDPerms:     MakeIdPermsType(),
	}
}

// MakeServiceConnectionModuleSlice() makes a slice of ServiceConnectionModule
func MakeServiceConnectionModuleSlice() []*ServiceConnectionModule {
	return []*ServiceConnectionModule{}
}
