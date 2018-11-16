package models

//GetRBACRules retreives RBAC rules from API access list entry.
func (m *APIAccessList) GetRBACRules() []*RbacRuleType {
	return m.GetAPIAccessListEntries().GetRbacRule()
}
