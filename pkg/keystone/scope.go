package keystone

// GetDomain gets domain
func (s *Scope) GetDomain() *Domain {
	if s.Domain != nil {
		return s.Domain
	} else if s.Project != nil {
		return s.Project.Domain
	}
	return nil
}
