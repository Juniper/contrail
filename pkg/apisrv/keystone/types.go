package keystone

import (
	types "github.com/Juniper/contrail/pkg/common/keystone"
)

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
	UUID string `json:"uuid,omitempty"`
	Name string `json:"name,omitempty"`
}
