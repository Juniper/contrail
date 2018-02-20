package models

// ShareType

// ShareType
//proteus:generate
type ShareType struct {
	TenantAccess AccessType `json:"tenant_access,omitempty"`
	Tenant       string     `json:"tenant,omitempty"`
}

// MakeShareType makes ShareType
func MakeShareType() *ShareType {
	return &ShareType{
		//TODO(nati): Apply default
		TenantAccess: MakeAccessType(),
		Tenant:       "",
	}
}

// MakeShareTypeSlice() makes a slice of ShareType
func MakeShareTypeSlice() []*ShareType {
	return []*ShareType{}
}
