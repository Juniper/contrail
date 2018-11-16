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

type roleToCrudMap struct {
	cMap map[string]*crudSet
}

//Rbac request structure

type rbacRequest struct {
	hash      map[string]*roleToCrudMap // Hash table with objtype key  to roleToCrudMap
	op        constants.OpCrud          // Requested CRUD operation
	roles     map[string]bool           // Roles of the user making request
	objType   string                    // Objtype of the resource
	fieldName string                    // Object Feild name
	_         string                    // Domain name of the user.
	_         string                    // project name of the user.
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

func (s *crudSet) isOpAllowed(op constants.OpCrud) bool {

	if _, ok := s.cMap[op]; ok {
		return true

	}
	return false
}

func newRoleToCrudMap() *roleToCrudMap {

	var n roleToCrudMap
	n.cMap = make(map[string]*crudSet)
	return &n
}

func (m *roleToCrudMap) addEntry(role string, crud string) {

	if cs, ok := m.cMap[role]; ok {

		cs.setCrud(crud)

	} else {

		cs := newcrudSet()
		cs.setCrud(crud)
		m.cMap[role] = cs
	}
	return
}

//

func (m *roleToCrudMap) isRuleMatching(roleSet map[string]bool, op constants.OpCrud) bool {

	// Check against all user roles
	for role := range roleSet {
		if cs, ok := m.cMap[role]; ok {

			return cs.isOpAllowed(op)
		}
	}
	// Check against wild card role rule
	if cs, ok := m.cMap["*"]; ok {

		return cs.isOpAllowed(op)
	}
	return false

}

// Create a new rbacRequest Object .  This object will be initilized with user roles
// and relevant rbac rules hash .
func newRbacRequest(ctx context.Context, apiAccessRules []*models.APIAccessList,
	objType string, op constants.OpCrud) *rbacRequest {

	var n rbacRequest

	n.objType = objType
	n.op = op
	n.roles = APISrvRbac.getUserRoles(ctx)
	n.hash = make(map[string]*roleToCrudMap)

	//Add relevant RBAC rules for this object type to rbacRequest.hash
	n.rbacGetObjTypeRules(apiAccessRules)
	return &n
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
func (rq *rbacRequest) rbacGetHashKey(objType string, _ string) (key string) {

	// Not implemented field name based rules yet
	return objType + "." + "*"

}

// Add rules to the rule hash table. Rule hash table will have entries of the form
// key => CRUD . This will merge the object_type CRUD in global,domain and
// project level.

func (rq *rbacRequest) rbacUpdateRuleHash(role string, crud string) {

	key := rq.rbacGetHashKey(rq.objType, rq.fieldName)

	if val, ok := rq.hash[key]; ok {

		val.addEntry(role, crud)

	} else {

		val := newRoleToCrudMap()
		val.addEntry(role, crud)
		rq.hash[key] = val
	}

	return

}

//  Add objType or wildcard RBAC rules to hash.
//  Input "perm" can be of the following form
//   Member:CRUD
//

func (rq *rbacRequest) objTypePermAddToHash(perm *models.RbacPermType) {

	if perm == nil {
		return
	}
	role := perm.GetRoleName()
	crud := perm.GetRoleCrud()

	if _, ok := rq.roles[role]; ok {

		rq.rbacUpdateRuleHash(role, crud)
	}
	return

}

//  Add objType or wildcard RBAC rules to hash.
//  Input "perms" can be of the following form
//  Member:CRUD ,Development:CRUD
//

func (rq *rbacRequest) objTypePermsAddToHash(perms []*models.RbacPermType) {

	for _, perm := range perms {

		rq.objTypePermAddToHash(perm)
	}
	return

}

//  Add obj_type or wildcard RBAC rules to hash.
//  Input rbacRule can be of the following form
//  < virtual-network, *> => Member:CRUD ,Development:CRUD
//

func (rq *rbacRequest) objTypeRuleAddToHash(rbacRule *models.RbacRuleType) {

	if rbacRule == nil {
		return
	}
	ruleObjType := rbacRule.GetRuleObject()

	if ruleObjType != rq.objType && ruleObjType != "*" {
		return
	}

	perms := rbacRule.GetRulePerms()

	rq.objTypePermsAddToHash(perms)

	return

}

// Add  matching the objType or wildcard RBAC rules to hash.
//   Input can be of the following form
//   < virtual-network, *> => Member:CRUD ,Development:CRUD
//   < virtual-ip, *> => Member:CRUD
//

func (rq *rbacRequest) rbacRulesAddToHash(rbacRules []*models.RbacRuleType) {

	for _, obj := range rbacRules {

		rq.objTypeRuleAddToHash(obj)

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
// Function will create rq.hash in the following form
//   objtype.* ==> {'C','R','U','D}
// For example result could be
//  virtual-network.* ==> {'C','R','U','D}
//

func (rq *rbacRequest) rbacGetObjTypeRules(apiAccessRules []*models.APIAccessList) {

	// TODO Filter rules based on domain and project
	for _, apiRule := range apiAccessRules {

		rbacRules := apiRule.GetAPIRbacRules()
		rq.rbacRulesAddToHash(rbacRules)
	}

	return

}

//  Checks whether any obj type rules
//

func (rq *rbacRequest) validateResourceAPILevel(objType string) bool {

	key := rq.rbacGetHashKey(objType, "_")
	if val, ok := rq.hash[key]; ok {

		return val.isRuleMatching(rq.roles, rq.op)

	}
	return false

}

//  Checks whether any obj type rules or wildcard allow this operation
//  for any roles which user is part of .

func (rq *rbacRequest) validateAPILevel() bool {

	// check for objtype rule .
	if rq.validateResourceAPILevel(rq.objType) {
		return true
	}

	// check for wild card rule
	if rq.validateResourceAPILevel("*") {
		return true
	}

	return false
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

// Checks whether is Read only  Resource access is allowed if operation is READ.
//

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

//  ValidateRequest checks whether resource operation is allowed based on RBAC config.
func (r *rbac) ValidateRequest(ctx context.Context, apiAccessRules []*models.APIAccessList,
	objType string, op constants.OpCrud) (bool, error) {

	if ctx == nil {
		return false, errutil.ErrorPermissionDenied
	}

	authCtx := auth.GetAuthCTX(ctx)

	// Check whether access is allowed based on role or config

	if r.validateRoleAndConfig(authCtx) {
		return true, nil
	}

	// Check whether access is allowed based on global read only role.

	if r.IsROAccessAllowed(authCtx, op) {
		return true, nil
	}

	// If no rules to allow access, deny permission

	if len(apiAccessRules) == 0 || authCtx == nil {

		return false, errutil.ErrorPermissionDenied
	}
	// Create RBAC request object based on request parameters and rbac rule configuration
	rbacReq := newRbacRequest(ctx, apiAccessRules, objType, op)

	// Do the RBAC check
	if rbacReq.validateAPILevel() {

		return true, nil
	}

	log.Debug("Rbac : Access Denied ")
	return false, errutil.ErrorPermissionDenied

}
