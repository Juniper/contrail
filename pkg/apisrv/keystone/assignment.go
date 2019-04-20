package keystone

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"

	"github.com/Juniper/contrail/pkg/apisrv/client"
	"github.com/Juniper/contrail/pkg/format"
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

//VNCAPIAssignment is an implementation of Assignment based on vnc api-server.
type VNCAPIAssignment struct {
	Domains   map[string]*kscommon.Domain  `json:"domains"`
	Projects  map[string]*kscommon.Project `json:"projects"`
	Users     map[string]*kscommon.User    `json:"users"`
	vncClient *client.HTTP
}

//ListUsers is used to fetch all users
func (assignment *StaticAssignment) ListUsers() map[string]*kscommon.User {
	return assignment.Users
}

//FetchUser is used to fetch a user by ID and Password.
func (assignment *StaticAssignment) FetchUser(name, password string) (*kscommon.User, error) {
	user, ok := assignment.Users[name]
	if !ok {
		return nil, fmt.Errorf("user %s not found", name)
	}
	if user.Password != "" && format.InterfaceToString(user.Password) != password {
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

//FetchUser is used to fetch a user by ID and Password.
func (assignment *VNCAPIAssignment) FetchUser(name, password string) (*kscommon.User, error) {
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
func (assignment *VNCAPIAssignment) ListDomains() []*kscommon.Domain {
	domains := []*kscommon.Domain{}
	for _, domain := range assignment.Domains {
		domain = append(domains, domain)
	}
	return projects
}

//ListProjects is used to list projects
func (assignment *VNCAPIAssignment) ListProjects() []*kscommon.Project {
	projects := []*kscommon.Project{}
	for _, project := range assignment.Projects {
		projects = append(projects, project)
	}
	return projects
}

//Init the VNCAPI assignment with vnc api-server projects/domains
func (assignment *VNCAPIAssignment) Init(
	configEndpoint string, staticUsers map[string]*kscommon.User) error {
	if assignment.vncClient == nil {
		assignment.vncClient = client.NewHTTP("", "", "", "", true, nil)
	}
	assignment.vncClient.Endpoint = configEndpoint
	assignment.vncClient.Init()
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
	VNCAPIUsers := make(map[string]*kscommon.User)
	for name, user := range staticUsers {
		VNCAPIUser := &kscommon.User{
			Domain:   user.Domain,
			ID:       user.ID,
			Name:     user.Name,
			Password: user.Password,
			Email:    user.Email,
			Roles:    make([]*kscommon.Role, len(roles)),
		}
		copy(VNCAPIUser.Roles, roles)
		VNCAPIUser.Roles = append(VNCAPIUser.Roles, user.Roles...)
		VNCAPIUsers[name] = VNCAPIUser
	}
	assignment.Users = VNCAPIUsers

	return nil
}

func (assignment *VNCAPIAssignment) getVncProjects(
	domains map[string]*kscommon.Domain) (map[string]*kscommon.Project, []*kscommon.Role, error) {
	projectURI := "/projects"
	query := url.Values{"detail": []string{"True"}}
	vncProjectsResponse := &VncProjectListResponse{}
	_, err := assignment.vncClient.ReadWithQuery(
		context.Background(), projectURI, query, vncProjectsResponse)
	if err != nil {
		return nil, nil, err
	}
	projects := map[string]*kscommon.Project{}
	roles := []*kscommon.Role{}
	for _, vncProject := range vncProjectsResponse.Projects {
		project := &kscommon.Project{
			Name:   vncProject.Project.Name,
			ID:     strings.Replace(vncProject.Project.UUID, "-", "", -1),
			Domain: domains[vncProject.Project.FQName[0]],
		}

		projects[vncProject.Project.Name] = project
		roles = append(roles, &kscommon.Role{
			ID:      "admin",
			Name:    "admin",
			Project: project,
		})
	}
	return projects, roles, nil
}

func (assignment *VNCAPIAssignment) getVncDomains() (map[string]*kscommon.Domain, error) {
	domainURI := "/domains"
	query := url.Values{"detail": []string{"True"}}
	vncDomainsResponse := &VncDomainListResponse{}
	_, err := assignment.vncClient.ReadWithQuery(
		context.Background(), domainURI, query, vncDomainsResponse)
	if err != nil {
		return nil, err
	}
	domains := map[string]*kscommon.Domain{}
	for _, vncDomain := range vncDomainsResponse.Domains {
		domains[vncDomain.Domain.Name] = &kscommon.Domain{
			Name: vncDomain.Domain.Name,
			ID:   strings.Replace(vncDomain.Domain.UUID, "-", "", -1),
		}
	}
	return domains, nil
}
