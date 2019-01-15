package models

// EnableDomainSharing enables domain sharing for resource.
func (p *PermType2) EnableDomainSharing(
	domainUUID string,
	accessLevel int64,
) error {

	if p == nil {
		return nil
	}

	p.Share = append(p.Share, &ShareType{
		Tenant:       "domain:" + domainUUID,
		TenantAccess: accessLevel,
	})

	return nil
}
