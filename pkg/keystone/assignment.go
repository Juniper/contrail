package keystone

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/Juniper/asf/pkg/format"
	"github.com/Juniper/asf/pkg/keystone"
	"github.com/Juniper/contrail/pkg/client"
	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"

	asfclient "github.com/Juniper/asf/pkg/client"
)

//Assignment is used to manage domain, project and user information.
type Assignment interface {
	FetchUser(id, password string) (*keystone.User, error)
	ListProjects() []*keystone.Project
	ListDomains() []*keystone.Domain
	ListUsers() []*keystone.User
}

//StaticAssignment is an implementation of Assignment based on static file.
type StaticAssignment struct {
	Domains  map[string]*keystone.Domain  `json:"domains"`
	Projects map[string]*keystone.Project `json:"projects"`
	Users    map[string]*keystone.User    `json:"users"`
}

//ListUsers is used to fetch all users
func (assignment *StaticAssignment) ListUsers() (users []*keystone.User) {
	for _, user := range assignment.Users {
		users = append(users, user)
	}
	return users
}

//FetchUser is used to fetch a user by ID and Password.
func (assignment *StaticAssignment) FetchUser(name, password string) (*keystone.User, error) {
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
func (assignment *StaticAssignment) ListDomains() (domains []*keystone.Domain) {
	for _, domain := range assignment.Domains {
		domains = append(domains, domain)
	}
	return domains
}

//ListProjects is used to list projects
func (assignment *StaticAssignment) ListProjects() (projects []*keystone.Project) {
	for _, project := range assignment.Projects {
		projects = append(projects, project)
	}
	return projects
}

//VNCAPIAssignment is an implementation of Assignment based on vnc api-server.
type VNCAPIAssignment struct {
	Domains   map[string]*keystone.Domain  `json:"domains"`
	Projects  map[string]*keystone.Project `json:"projects"`
	Users     map[string]*keystone.User    `json:"users"`
	vncClient *client.HTTP
}

//Init the VNCAPI assignment with vnc api-server projects/domains
func (assignment *VNCAPIAssignment) Init(configEndpoint string, staticUsers map[string]*keystone.User) error {
	if assignment.vncClient == nil {
		assignment.vncClient = client.NewHTTP(&asfclient.HTTPConfig{Insecure: true})
	}
	assignment.vncClient.Endpoint = configEndpoint

	domains, err := assignment.getVncDomains()
	if err != nil {
		logrus.Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	projects, roles, err := assignment.getVncProjects(domains)
	if err != nil {
		logrus.Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	assignment.Domains = domains
	assignment.Projects = projects
	VNCAPIUsers := map[string]*keystone.User{}
	for name, user := range staticUsers {
		VNCAPIUser := &keystone.User{
			Domain:   user.Domain,
			ID:       user.ID,
			Name:     user.Name,
			Password: user.Password,
			Email:    user.Email,
			Roles:    make([]*keystone.Role, len(roles)),
		}
		copy(VNCAPIUser.Roles, roles)
		VNCAPIUser.Roles = append(VNCAPIUser.Roles, user.Roles...)
		VNCAPIUsers[name] = VNCAPIUser
	}
	assignment.Users = VNCAPIUsers

	return nil
}

func (assignment *VNCAPIAssignment) getVncDomains() (map[string]*keystone.Domain, error) {
	domainURI := "/domains"
	query := url.Values{"detail": []string{"True"}}
	vncDomainsResponse := &VncDomainListResponse{}
	_, err := assignment.vncClient.ReadWithQuery(context.Background(), domainURI, query, vncDomainsResponse)
	if err != nil {
		return nil, err
	}
	domains := map[string]*keystone.Domain{}
	for _, vncDomain := range vncDomainsResponse.Domains {
		domains[vncDomain.Domain.Name] = &keystone.Domain{
			Name: vncDomain.Domain.Name,
			ID:   strings.Replace(vncDomain.Domain.UUID, "-", "", -1),
		}
	}
	return domains, nil
}

func (assignment *VNCAPIAssignment) getVncProjects(
	domains map[string]*keystone.Domain) (map[string]*keystone.Project, []*keystone.Role, error) {
	projectURI := "/projects"
	query := url.Values{"detail": []string{"True"}}
	vncProjectsResponse := &VncProjectListResponse{}
	_, err := assignment.vncClient.ReadWithQuery(context.Background(), projectURI, query, vncProjectsResponse)
	if err != nil {
		return nil, nil, err
	}
	projects := map[string]*keystone.Project{}
	var roles []*keystone.Role
	for _, vncProject := range vncProjectsResponse.Projects {
		domain := domains[vncProject.Project.FQName[0]]
		project := &keystone.Project{
			Name:     vncProject.Project.Name,
			ID:       strings.Replace(vncProject.Project.UUID, "-", "", -1),
			Domain:   domain,
			ParentID: domain.ID,
		}

		projects[vncProject.Project.Name] = project
		roles = append(roles, &keystone.Role{
			ID:      "admin",
			Name:    "admin",
			Project: project,
		})
	}
	return projects, roles, nil
}

//FetchUser is used to fetch a user by ID and Password.
func (assignment *VNCAPIAssignment) FetchUser(name, password string) (*keystone.User, error) {
	user, ok := assignment.Users[name]
	if !ok {
		return nil, fmt.Errorf("user %s not found", name)
	}
	if user.Password != "" && format.InterfaceToString(user.Password) != password {
		return nil, fmt.Errorf("invalid Credentials")
	}
	return user, nil
}

//ListDomains is used to list domains
func (assignment *VNCAPIAssignment) ListDomains() (domains []*keystone.Domain) {
	for _, domain := range assignment.Domains {
		domains = append(domains, domain)
	}
	return domains
}

//ListProjects is used to list projects
func (assignment *VNCAPIAssignment) ListProjects() (projects []*keystone.Project) {
	for _, project := range assignment.Projects {
		projects = append(projects, project)
	}
	return projects
}

//ListUsers is used to list users
func (assignment *VNCAPIAssignment) ListUsers() (users []*keystone.User) {
	for _, user := range assignment.Users {
		users = append(users, user)
	}
	return users
}
