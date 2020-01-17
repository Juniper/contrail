package models

import (
	"github.com/Juniper/asf/pkg/rbac"
)

type APIAccessLists []*APIAccessList

func (a APIAccessLists) ToRBAC() []*rbac.APIAccessList {
	rbacList := make([]*rbac.APIAccessList, 0, len(a))
	for _, v := range a {
		rbacList = append(rbacList, &rbac.APIAccessList{
			ParentType:           v.ParentType,
			FQName:               v.FQName,
			APIAccessListEntries: v.GetAPIAccessListEntries().toRBAC(),
		})
	}
	return rbacList
}

func (m *RbacRuleEntriesType) toRBAC() *rbac.RuleEntriesType {
	var modelsRules []*RbacRuleType = m.GetRbacRule()
	rbacRules := make([]*rbac.RuleType, 0, len(modelsRules))
	for _, v := range modelsRules {
		rbacRules = append(rbacRules, &rbac.RuleType{
			RuleObject : v.RuleObject,
			RulePerms  : RbacPermTypes(v.GetRulePerms()).toRBAC(),
		})
	}
	return &rbac.RuleEntriesType{
		Rule: rbacRules,
	}
}

type RbacPermTypes []*RbacPermType

func (r RbacPermTypes) toRBAC() []*rbac.PermType {
	rbacTypes := make([]*rbac.PermType, 0, len(r))
	for _, v := range r {
		rbacTypes = append(rbacTypes, &rbac.PermType{
			RoleCrud: v.GetRoleCrud(),
			RoleName: v.GetRoleName(),
		})
	}
	return rbacTypes
}

func (p *PermType2) ToRBAC() *rbac.PermType2 {
	return &rbac.PermType2{
			Owner: p.Owner,
			OwnerAccess: p.OwnerAccess,
			GlobalAccess: p.GlobalAccess,
			Share: ShareTypes(p.GetShare()).toRBAC(),
		}
}

type ShareTypes []*ShareType

func (s ShareTypes) toRBAC() []*rbac.ShareType {
	rbacShares := make([]*rbac.ShareType, 0, len(s))
	for _, v := range s {
		rbacShares = append(rbacShares, &rbac.ShareType{
			TenantAccess: v.GetTenantAccess(),
			Tenant: v.GetTenant(),
		})
	}
	return rbacShares
}
