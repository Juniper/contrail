package keystone

import (
	"fmt"

	"github.com/Juniper/asf/pkg/format"
)

//StaticAssignment is an implementation of Assignment based on static file.
type StaticAssignment struct {
	Domains  map[string]*Domain  `json:"domains"`
	Projects map[string]*Project `json:"projects"`
	Users    map[string]*User    `json:"users"`
}

//ListUsers is used to fetch all users
func (assignment *StaticAssignment) ListUsers() map[string]*User {
	return assignment.Users
}

//FetchUser is used to fetch a user by ID and Password.
func (assignment *StaticAssignment) FetchUser(name, password string) (*User, error) {
	user, ok := assignment.Users[name]
	if !ok {
		return nil, fmt.Errorf("user %s not found", name)
	}
	if user.Password != "" && format.InterfaceToString(user.Password) != password {
		return nil, fmt.Errorf("invalid credentials")
	}
	return user, nil
}

//ListDomains is used to list domains
func (assignment *StaticAssignment) ListDomains() []*Domain {
	domains := []*Domain{}
	for _, domain := range assignment.Domains {
		domains = append(domains, domain)
	}
	return domains
}

//ListProjects is used to list projects
func (assignment *StaticAssignment) ListProjects() []*Project {
	var projects []*Project
	for _, project := range assignment.Projects {
		projects = append(projects, project)
	}
	return projects
}
