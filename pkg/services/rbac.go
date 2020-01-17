package services

import (
	"github.com/Juniper/asf/pkg/rbac"
	"github.com/Juniper/contrail/pkg/models"
)

func toRBACAPIAcesssList(modelsList []*models.APIAccessList) []*rbac.APIAccessList {
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

func toRBACRuleEntriesType(modelsType *models.RbacRuleEntriesType) *rbac.RuleEntriesType {
	var modelsRules []*models.RbacRuleType = modelsType.GetRbacRule()
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

func toRBACRulePerms(modelsTypes []*models.RbacPermType) []*rbac.PermType {
	rbacTypes := make([]*rbac.PermType, len(modelsTypes))
	for i, v := range modelsTypes {
		rbacTypes[i] = &rbac.PermType{
			RoleCrud: v.GetRoleCrud(),
			RoleName: v.GetRoleName(),
		}
	}
	return rbacTypes
}

func toRBACPermType2(modelsType *models.PermType2) *rbac.PermType2 {
	return &rbac.PermType2{
			Owner: modelsType.Owner,
			OwnerAccess: modelsType.OwnerAccess,
			GlobalAccess: modelsType.GlobalAccess,
			Share: toRBACShares(modelsType.GetShare()),
		}
}

func toRBACShares(modelsShares []*models.ShareType) []*rbac.ShareType {
	rbacShares := make([]*rbac.ShareType, len(modelsShares))
	for i, v := range modelsShares {
		rbacShares[i] = &rbac.ShareType{
			TenantAccess: v.GetTenantAccess(),
			Tenant: v.GetTenant(),
		}
	}
	return rbacShares
}
