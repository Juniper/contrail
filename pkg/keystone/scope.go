package keystone

// GetDomain gets domain
func (s *Scope) GetDomain() *Domain {
	if s == nil {
		return nil
	}
	if s.Domain != nil {
		return s.Domain
	} else if s.Project != nil {
		return s.Project.Domain
	}
	return nil
}

// GetScope returns the project/domain scope
func GetScope(domainID, domainName, projectID, projectName string) *Scope {
	scope := &Scope{
		Project: &Project{
			Domain: &Domain{},
		},
	}
	if domainID != "" {
		scope.Project.Domain.ID = domainID
	} else if domainName != "" {
		scope.Project.Domain.Name = domainName
	}
	if projectID != "" {
		scope.Project.ID = projectID
	} else if projectName != "" {
		scope.Project.Name = projectName
	}
	return scope
}
