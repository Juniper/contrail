package apisrv

import "github.com/Juniper/contrail/pkg/apisrv/keystone"

//CreateTestProject in keystone.
func CreateTestProject(s *Server, testID string) {
	assignment := s.Keystone.Assignment.(*keystone.StaticAssignment)
	assignment.Projects[testID] = &keystone.Project{
		Domain: assignment.Domains["default"],
		ID:     testID,
		Name:   testID,
	}

	assignment.Users[testID] = &keystone.User{
		Domain:   assignment.Domains["default"],
		ID:       testID,
		Name:     testID,
		Password: testID,
		Roles: []*keystone.Role{
			{
				ID:      "member",
				Name:    "Member",
				Project: assignment.Projects[testID],
			},
		},
	}
}

//Task has API request and expected response.
type Task struct {
	Name    string      `yaml:"name"`
	Client  string      `yaml:"client"`
	Request *Request    `yaml:"request"`
	Expect  interface{} `yaml:"expect"`
}

//TestScenario has a list of tasks.
type TestScenario struct {
	Name        string              `yaml:"name"`
	Description string              `yaml:"description"`
	Tables      []string            `yaml:"tables"`
	Clients     map[string]*Client  `yaml:"clients"`
	Cleanup     []map[string]string `yaml:"cleanup"`
	Workflow    []*Task             `yaml:"workflow"`
}
