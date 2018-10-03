package keystone

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/labstack/echo"
	log "github.com/sirupsen/logrus"

	"github.com/Juniper/contrail/pkg/apisrv/client"
	"github.com/Juniper/contrail/pkg/common"
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

//DynamicAssignment is an implementation of Assignment based on vnc api-server.
type DynamicAssignment struct {
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
		return nil, fmt.Errorf("User %s not found", name)
	}
	if user.Password != "" && common.InterfaceToString(user.Password) != password {
		return nil, fmt.Errorf("Invalid Credentials")
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
func (assignment *DynamicAssignment) FetchUser(name, password string) (*kscommon.User, error) {
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
func (assignment *DynamicAssignment) ListProjects() []*kscommon.Project {
	projects := []*kscommon.Project{}
	for _, project := range assignment.Projects {
		projects = append(projects, project)
	}
	return projects
}

//Init the dynamic assignment with vnc api-server projects/domains
func (assignment *DynamicAssignment) Init(
	configEndpoint string, staticUsers map[string]*kscommon.User) error {
	if assignment.vncClient == nil {
		assignment.vncClient = client.NewHTTP("", "", "", "", true, nil)
	}
	assignment.vncClient.Endpoint = configEndpoint
	assignment.vncClient.Init()
	domains, err := assignment.getVncDomains()
	if err != nil {
		log.Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	projects, roles, err := assignment.getVncProjects(domains)
	if err != nil {
		log.Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	assignment.Domains = domains
	assignment.Projects = projects
	dynamicUsers := make(map[string]*kscommon.User)
	for name, user := range staticUsers {
		dynamicUser := user
		copy(roles, dynamicUser.Roles)
		dynamicUser.Roles = roles
		roles = append(roles, user.Roles...)
		dynamicUsers[name] = dynamicUser
	}
	assignment.Users = dynamicUsers

	return nil
}

func (assignment *DynamicAssignment) getVncProjects(
	domains map[string]*kscommon.Domain) (map[string]*kscommon.Project, []*kscommon.Role, error) {
	projectURI := "/projects"
	query := url.Values{"detail": []string{"True"}}
	vncProjectsResponse := &VncProjectListResponse{}
	_, err := assignment.vncClient.ReadWithQuery(
		projectURI, query, vncProjectsResponse)
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
			Name:    "Admin",
			Project: project,
		})
	}
	return projects, roles, nil
}

func (assignment *DynamicAssignment) getVncDomains() (map[string]*kscommon.Domain, error) {
	domainURI := "/domains"
	query := url.Values{"detail": []string{"True"}}
	vncDomainsResponse := &VncDomainListResponse{}
	_, err := assignment.vncClient.ReadWithQuery(
		domainURI, query, vncDomainsResponse)
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
