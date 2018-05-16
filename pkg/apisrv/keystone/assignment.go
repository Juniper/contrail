package keystone

import (
	"fmt"

	"github.com/Juniper/contrail/pkg/common"
)

//Assignment is used to manage domain, project and user information.
type Assignment interface {
	FetchUser(id, password string) (*User, error)
	ListProjects() []*Project
}

//StaticAssignment is an implementation of Assignment based on static file.
type StaticAssignment struct {
	Domains  map[string]*Domain  `json:"domains"`
	Projects map[string]*Project `json:"projects"`
	Users    map[string]*User    `json:"users"`
}

//FetchUser is used to fetch a user by ID and Password.
func (assignment *StaticAssignment) FetchUser(name, password string) (*User, error) {
	user, ok := assignment.Users[name]
	if !ok {
		return nil, fmt.Errorf("User %s not found", name)
	}
	if user.Password != "" && common.InterfaceToString(user.Password) != password {
		return nil, fmt.Errorf("Invalid Credentials")
	}
	return user, nil
}

//ListProjects is used to list projects
func (assignment *StaticAssignment) ListProjects() []*Project {
	projects := []*Project{}
	for _, project := range assignment.Projects {
		projects = append(projects, project)
	}
	return projects
}
