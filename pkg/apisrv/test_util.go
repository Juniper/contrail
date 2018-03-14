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
			&keystone.Role{
				ID:      "member",
				Name:    "Member",
				Project: assignment.Projects[testID],
			},
		},
	}
}
