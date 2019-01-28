package keystone

import (
	types "github.com/Juniper/contrail/pkg/keystone"
)

//ProjectResponse represents a project list response.
type ProjectResponse struct {
	Project *types.Project `json:"project"`
}

//ProjectListResponse represents a project list response.
type ProjectListResponse struct {
	Projects []*types.Project `json:"projects"`
}

//VncProjectListResponse represents a project list response.
type VncProjectListResponse struct {
	Projects []*VncProject `json:"projects"`
}

//VncProject represents a vnc config project object.
type VncProject struct {
	Project *ConfigProject `json:"project"`
}

//ConfigProject represents project object.
type ConfigProject struct {
	UUID   string   `json:"uuid,omitempty"`
	Name   string   `json:"name,omitempty"`
	FQName []string `json:"fq_name,omitempty"`
}

//VncDomainListResponse represents a domain list response.
type VncDomainListResponse struct {
	Domains []*VncDomain `json:"domains"`
}

//VncDomain represents a vnc config domain object.
type VncDomain struct {
	Domain *ConfigDomain `json:"domain"`
}

//ConfigDomain represents domain object.
type ConfigDomain struct {
	UUID string `json:"uuid,omitempty"`
	Name string `json:"name,omitempty"`
}
