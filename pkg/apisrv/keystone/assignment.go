package keystone

import "fmt"

//Assignment is used to manage domain, project and user information.
type Assignment interface {
	FetchUser(id, password string) (*User, error)
}

//StaticAssignment is an implementation of Assignment based on static file.
type StaticAssignment struct {
	Domains  map[string]*Domain  `json:"domains"`
	Projects map[string]*Project `json:"projects"`
	Users    map[string]*User    `json:"users"`
}

//FetchUser is used to fetch a user by ID and Password.
func (assignment *StaticAssignment) FetchUser(id, password string) (*User, error) {
	user, ok := assignment.Users[id]
	if !ok {
		return nil, fmt.Errorf("User %s not found", id)
	}
	return user, nil
}
