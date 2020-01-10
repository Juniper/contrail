package keystone



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
