package models

// GetAPIRbacRules retreives RBAC rules from API access list entry
func (m *APIAccessList) GetAPIRBACRules() []*RbacRuleType {
	if m == nil {
		return nil
	}

	return m.GetAPIAccessListEntries().GetRbacRule()
}
