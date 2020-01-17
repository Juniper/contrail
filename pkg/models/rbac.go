package models

import (
	"github.com/Juniper/asf/pkg/rbac"
)

func ToRBACAPIAcesssList(modelsList []*APIAccessList) []*rbac.APIAccessList {
	rbacList := make([]*rbac.APIAccessList, 0, len(modelsList))
	for _, v := range modelsList {
		rbacList = append(rbacList, &rbac.APIAccessList{
			ParentType:           v.ParentType,
			FQName:               v.FQName,
			APIAccessListEntries: toRBACRuleEntriesType(v.GetAPIAccessListEntries()),
		})
	}
	return rbacList
}

func toRBACRuleEntriesType(modelsType *RbacRuleEntriesType) *rbac.RuleEntriesType {
	var modelsRules []*RbacRuleType = modelsType.GetRbacRule()
	rbacRules := make([]*rbac.RuleType, 0, len(modelsRules))
	for _, v := range modelsRules {
		rbacRules = append(rbacRules, &rbac.RuleType{
			RuleObject : v.RuleObject,
			RulePerms  : toRBACRulePerms(v.GetRulePerms()),
		})
	}
	return &rbac.RuleEntriesType{
		Rule: rbacRules,
	}
}

func toRBACRulePerms(modelsTypes []*RbacPermType) []*rbac.PermType {
	rbacTypes := make([]*rbac.PermType, 0, len(modelsTypes))
	for _, v := range modelsTypes {
		rbacTypes = append(rbacTypes, &rbac.PermType{
			RoleCrud: v.GetRoleCrud(),
			RoleName: v.GetRoleName(),
		})
	}
	return rbacTypes
}

func ToRBACPermType2(modelsType *PermType2) *rbac.PermType2 {
	return &rbac.PermType2{
			Owner: modelsType.Owner,
			OwnerAccess: modelsType.OwnerAccess,
			GlobalAccess: modelsType.GlobalAccess,
			Share: toRBACShares(modelsType.GetShare()),
		}
}

func toRBACShares(modelsShares []*ShareType) []*rbac.ShareType {
	rbacShares := make([]*rbac.ShareType, 0, len(modelsShares))
	for _, v := range modelsShares {
		rbacShares = append(rbacShares, &rbac.ShareType{
			TenantAccess: v.GetTenantAccess(),
			Tenant: v.GetTenant(),
		})
	}
	return rbacShares
}
