package models

import (
	"github.com/Juniper/asf/pkg/rbac"
)

func ToRBACAPIAcesssList(modelsList []*APIAccessList) []*rbac.APIAccessList {
	rbacList := make([]*rbac.APIAccessList, len(modelsList))
	for i, v := range modelsList {
		rbacList[i] = &rbac.APIAccessList{
			ParentType:           v.ParentType,
			FQName:               v.FQName,
			APIAccessListEntries: toRBACRuleEntriesType(v.GetAPIAccessListEntries()),
		}
	}
	return rbacList
}

func toRBACRuleEntriesType(modelsType *RbacRuleEntriesType) *rbac.RuleEntriesType {
	var modelsRules []*RbacRuleType = modelsType.GetRbacRule()
	rbacRules := make([]*rbac.RuleType, len(modelsRules))
	for i, v := range modelsRules {
		rbacRules[i] = &rbac.RuleType{
			RuleObject : v.RuleObject,
			RulePerms  : toRBACRulePerms(v.GetRulePerms()),
		}
	}
	return &rbac.RuleEntriesType{
		Rule: rbacRules,
	}
}

func toRBACRulePerms(modelsTypes []*RbacPermType) []*rbac.PermType {
	rbacTypes := make([]*rbac.PermType, len(modelsTypes))
	for i, v := range modelsTypes {
		rbacTypes[i] = &rbac.PermType{
			RoleCrud: v.GetRoleCrud(),
			RoleName: v.GetRoleName(),
		}
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
	rbacShares := make([]*rbac.ShareType, len(modelsShares))
	for i, v := range modelsShares {
		rbacShares[i] = &rbac.ShareType{
			TenantAccess: v.GetTenantAccess(),
			Tenant: v.GetTenant(),
		}
	}
	return rbacShares
}
