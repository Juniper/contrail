package keystone

import (
	"fmt"
	"github.com/Juniper/contrail/pkg/cast"

	kscommon "github.com/Juniper/contrail/pkg/keystone"
)

//Assignment is used to manage domain, project and user information.
type Assignment interface {
	FetchUser(id, password string) (*kscommon.User, error)
	ListProjects() []*kscommon.Project
}

//StaticAssignment is an implementation of Assignment based on static file.
type StaticAssignment struct {
	Domains  map[string]*kscommon.Domain  `json:"domains"`
	Projects map[string]*kscommon.Project `json:"projects"`
	Users    map[string]*kscommon.User    `json:"users"`
}

//FetchUser is used to fetch a user by ID and Password.
func (assignment *StaticAssignment) FetchUser(name, password string) (*kscommon.User, error) {
	user, ok := assignment.Users[name]
	if !ok {
		return nil, fmt.Errorf("user %s not found", name)
	}
	if user.Password != "" && cast.InterfaceToString(user.Password) != password {
		return nil, fmt.Errorf("invalid credentials")
	}
	return user, nil
}

//ListProjects is used to list projects
func (assignment *StaticAssignment) ListProjects() []*kscommon.Project {
	projects := []*kscommon.Project{}
	for _, project := range assignment.Projects {
		projects = append(projects, project)
	}
	return projects
}
