package keystone

import (
	"fmt"

	"github.com/Juniper/contrail/pkg/common"
	types "github.com/Juniper/contrail/pkg/common/keystone"
)

//Assignment is used to manage domain, project and user information.
type Assignment interface {
	FetchUser(id, password string) (*types.User, error)
	ListProjects() []*types.Project
}

//StaticAssignment is an implementation of Assignment based on static file.
type StaticAssignment struct {
	Domains  map[string]*types.Domain  `json:"domains"`
	Projects map[string]*types.Project `json:"projects"`
	Users    map[string]*types.User    `json:"users"`
}

//FetchUser is used to fetch a user by ID and Password.
func (assignment *StaticAssignment) FetchUser(name, password string) (*types.User, error) {
	user, ok := assignment.Users[name]
	if !ok {
		return nil, fmt.Errorf("user %s not found", name)
	}
	if user.Password != "" && common.InterfaceToString(user.Password) != password {
		return nil, fmt.Errorf("invalid credentials")
	}
	return user, nil
}

//ListProjects is used to list projects
func (assignment *StaticAssignment) ListProjects() []*types.Project {
	projects := []*types.Project{}
	for _, project := range assignment.Projects {
		projects = append(projects, project)
	}
	return projects
}
