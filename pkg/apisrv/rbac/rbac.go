package rbac

import (
	"context"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	"github.com/Juniper/contrail/pkg/auth"
	"github.com/Juniper/contrail/pkg/constants"
	"github.com/Juniper/contrail/pkg/errutil"
	"github.com/Juniper/contrail/pkg/models"
)

type rbac struct{}

//APISrvRbac , rbac singleton instance
var APISrvRbac = &rbac{}

type crudSet struct {
	cMap map[constants.OpCrud]bool
}

func newcrudSet() *crudSet {

	var n crudSet
	n.cMap = make(map[constants.OpCrud]bool)
	return &n
}

func (s *crudSet) setCrud(crud string) {

	for _, c := range crud {
		op := APISrvRbac.rbacCrudToOpMap(c)
		s.cMap[op] = true
	}
	return
}

// Return whether rbac is enabled .
func (r *rbac) isRbacEnabled() bool {

	return viper.GetString("aaa_mode") == "rbac"
}

// Convert user roles from ctx to a map.
func (r *rbac) getUserRoles(ctx context.Context) map[string]bool {

	rset := make(map[string]bool)

	if ctx == nil {
		return nil
	}

	aCtx := auth.GetAuthCTX(ctx)

	if aCtx == nil {
		return nil
	}

	roles := aCtx.Roles()

	for _, role := range roles {
		rset[role] = true
	}

	return rset
}

//  Form a hash key from objName,objField,role
func (r *rbac) rbacGetHashKey(objName string, _ string, role string) (key string) {

	// Not implemented field name based rules yet
	return objName + ".*." + role

}

// Add rules to the rule hash table. Rule hash table will have entries of the form
// key => CRUD . This will merge the object_type CRUD in global,domain and
// project level.

func (r *rbac) rbacUpdateRuleHash(hash map[string]*crudSet, objName string,
	fieldName string, role string, crud string) {

	key := r.rbacGetHashKey(objName, fieldName, role)

	if val, ok := hash[key]; ok {

		val.setCrud(crud)

	} else {

		val := newcrudSet()
		val.setCrud(crud)
		hash[key] = val
	}

	return

}

//  Add objType or wildcard RBAC rules to hash.
//  Input "perm" can be of the following form
//   Member:CRUD
//

func (r *rbac) objTypePermAddToHash(perm *models.RbacPermType, hash map[string]*crudSet,
	roles map[string]bool, objType string, objField string) {

	if perm == nil {
		return
	}
	role := perm.GetRoleName()
	crud := perm.GetRoleCrud()

	if _, ok := roles[role]; ok {

		r.rbacUpdateRuleHash(hash, objType, objField, role, crud)
	}
	return

}

//  Add objType or wildcard RBAC rules to hash.
//  Input "perms" can be of the following form
//   Member:CRUD ,Development:CRUD
//

func (r *rbac) objTypePermsAddToHash(perms []*models.RbacPermType, hash map[string]*crudSet,
	roles map[string]bool, objType string, objField string) {

	for _, perm := range perms {

		r.objTypePermAddToHash(perm, hash, roles, objType, objField)
	}
	return

}

//  Add obj_type or wildcard RBAC rules to hash.
//   Input rbacRule can be of the following form
//   < virtual-network, *> => Member:CRUD ,Development:CRUD
//

func (r *rbac) objTypeRuleAddToHash(rbacRule *models.RbacRuleType, hash map[string]*crudSet,
	userRoles map[string]bool, objType string) {

	if rbacRule == nil {
		return
	}
	ruleObjType := rbacRule.GetRuleObject()
	ruleObjField := rbacRule.GetRuleField()

	if ruleObjType != objType && ruleObjType != "*" {
		return
	}

	perms := rbacRule.GetRulePerms()

	r.objTypePermsAddToHash(perms, hash, userRoles, ruleObjType, ruleObjField)

	return

}

// Add  matching the objType or wildcard RBAC rules to hash.
//   Input can be of the following form
//   < virtual-network, *> => Member:CRUD ,Development:CRUD
//   < virtual-ip, *> => Member:CRUD
//

func (r *rbac) rbacRulesAddToHash(rbacRules []*models.RbacRuleType, hash map[string]*crudSet,
	roles map[string]bool, objType string) {

	for _, obj := range rbacRules {

		r.objTypeRuleAddToHash(obj, hash, roles, objType)

	}
	return
}

// Get rules matching the objectType from the rule list .
// For example input apiAccessRules may contains rules from global-config, default-domain & project
// different all object types. There can be multiple roles and CRUD for every object type .
// Following  is an input example.
//
// Rule 1
// 1)  <global-config, virtual-network, network-policy> => Member:CRUD
// Rule 2
// 2)  <default-domain,virtual-network, network-ipam> => Development:CRUD
// Rule 3
// 1)  <project,virtual-ip, *>		 => Member:CRUD
// 2)  <project,virtual-network, *>     => Member:CRUD, Development:CRUD
//
// Function will return a map in the following format
//   objtype.*.role ==> {'C','R','U','D}
// For example result could be
//  virtual-network.*Member ==> {'C','R','U','D}
//

func (r *rbac) rbacGetObjTypeRules(apiAccessRules []*models.APIAccessList,
	roles map[string]bool, objType string) map[string]*crudSet {

	hash := make(map[string]*crudSet)

	for _, apiRule := range apiAccessRules {

		rbacRules := apiRule.GetAPIRbacRules()
		r.rbacRulesAddToHash(rbacRules, hash, roles, objType)
	}

	return hash

}

//'CRUD' to operation  mapping

func (r *rbac) rbacCrudToOpMap(crud rune) constants.OpCrud {

	var res constants.OpCrud

	switch crud {
	case 'C':
		res = constants.OpCreate
	case 'R':
		res = constants.OpRead
	case 'U':
		res = constants.OpUpdate
	case 'D':
		res = constants.OpDelete
	}
	return res
}

// Checks whether is Resource access is allowed purely based on  RBAC config
// or user roles.

func (r *rbac) IsROAccessAllowed(ctx *auth.Context, op constants.OpCrud) bool {

	if op != constants.OpRead {
		return false
	}

	if ctx.IsGlobalRORole() {
		return true
	}

	return false
}

// Checks whether is Resource access is allowed purely based on  RBAC config
// or user roles.

func (r *rbac) validateRoleAndConfig(ctx *auth.Context) bool {

	// Rbac is not configured
	if !r.isRbacEnabled() {
		return true
	}

	// If User is cloud_admin allow access

	if ctx.IsAdmin() {
		return true
	}

	return false
}

//  Checks whether any obj type rules or wildcard allow this operation
//  for a particular role.

func (r *rbac) validateRoleAPILevel(hash map[string]*crudSet, role string,
	objType string, op constants.OpCrud) bool {

	key := r.rbacGetHashKey(objType, "_", role)
	if val, ok := hash[key]; ok {

		if _, ok := val.cMap[op]; ok {

			return true
		}
	}
	return false

}

//  Checks whether any obj type rules or wildcard allow this operation
//  for any roles which user is part of .

func (r *rbac) validateAPILevel(hash map[string]*crudSet, roleSet map[string]bool,
	objtype string, op constants.OpCrud) bool {

	for role := range roleSet {

		// check for objtype rule .
		if r.validateRoleAPILevel(hash, role, objtype, op) {
			return true
		}
		// check for wild card role rule
		if r.validateRoleAPILevel(hash, "*", objtype, op) {

			return true
		}

		// check for wild card rule
		if r.validateRoleAPILevel(hash, role, "*", op) {
			return true
		}

		// TODO: need to handle rule object '/'
	}

	return false
}

func (r *rbac) ValidateRequest(ctx context.Context, apiAccessRules []*models.APIAccessList,
	objType string, op constants.OpCrud) (bool, error) {

	var hash map[string]*crudSet

	if ctx == nil {
		return false, errutil.ErrorPermissionDenied
	}

	authCtx := auth.GetAuthCTX(ctx)

	roleSet := r.getUserRoles(ctx)

	// Check whether access is allowed based on role or config

	if r.validateRoleAndConfig(authCtx) {
		return true, nil
	}

	// Check whether access is allowed based on global read only role.

	if r.IsROAccessAllowed(authCtx, op) {
		return true, nil
	}

	//Convert all RBAC rules to per object_type hash which map to CRUD

	hash = r.rbacGetObjTypeRules(apiAccessRules, roleSet, objType)

	// If no rules to allow access, deny permission

	if len(apiAccessRules) == 0 || authCtx == nil {

		return false, errutil.ErrorPermissionDenied
	}

	if r.validateAPILevel(hash, roleSet, objType, op) {

		return true, nil
	}

	log.Debug("Rbac : Access Denied ")
	return false, errutil.ErrorPermissionDenied

}
