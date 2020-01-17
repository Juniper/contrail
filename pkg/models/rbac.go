package models

import (
	"github.com/Juniper/asf/pkg/rbac"
)

// APIAccessLists is a type in models package corresponding to []*rbac.APIAccessList.
type APIAccessLists []*APIAccessList

// ToRBAC translates objects to []*rbac.APIAccessList.
func (a APIAccessLists) ToRBAC() []*rbac.APIAccessList {
	rbacLists := make([]*rbac.APIAccessList, 0, len(a))
	for _, v := range a {
		rbacLists = append(rbacLists, &rbac.APIAccessList{
			ParentType:           v.ParentType,
			FQName:               v.FQName,
			APIAccessListEntries: v.GetAPIAccessListEntries().toRBAC(),
		})
	}
	return rbacLists
}

func (m *RbacRuleEntriesType) toRBAC() *rbac.RuleEntriesType {
	rbacRules := make([]*rbac.RuleType, 0, len(m.GetRbacRule()))
	for _, v := range m.GetRbacRule() {
		rbacRules = append(rbacRules, &rbac.RuleType{
			RuleObject: v.RuleObject,
			RulePerms:  rbacPermTypes(v.GetRulePerms()).toRBAC(),
		})
	}
	return &rbac.RuleEntriesType{
		Rule: rbacRules,
	}
}

type rbacPermTypes []*RbacPermType

func (r rbacPermTypes) toRBAC() []*rbac.PermType {
	rbacTypes := make([]*rbac.PermType, 0, len(r))
	for _, v := range r {
		rbacTypes = append(rbacTypes, &rbac.PermType{
			RoleCrud: v.GetRoleCrud(),
			RoleName: v.GetRoleName(),
		})
	}
	return rbacTypes
}

// ToRBAC translates the object to *rbac.PermType2.
func (p *PermType2) ToRBAC() *rbac.PermType2 {
	return &rbac.PermType2{
		Owner:        p.Owner,
		OwnerAccess:  p.OwnerAccess,
		GlobalAccess: p.GlobalAccess,
		Share:        shareTypes(p.GetShare()).toRBAC(),
	}
}

type shareTypes []*ShareType

func (s shareTypes) toRBAC() []*rbac.ShareType {
	rbacShares := make([]*rbac.ShareType, 0, len(s))
	for _, v := range s {
		rbacShares = append(rbacShares, &rbac.ShareType{
			TenantAccess: v.GetTenantAccess(),
			Tenant:       v.GetTenant(),
		})
	}
	return rbacShares
}
