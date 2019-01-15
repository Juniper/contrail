package models

// EnableDomainSharing enables domain sharing for resource.
func (perms2 *PermType2) EnableDomainSharing(
	domainUUID string,
	accessLevel int64,
) error {

	if perms2 == nil {
		return nil
	}

	perms2.Share = append(perms2.Share, &ShareType{
		Tenant:       "domain:" + domainUUID,
		TenantAccess: accessLevel,
	})

	return nil
}
