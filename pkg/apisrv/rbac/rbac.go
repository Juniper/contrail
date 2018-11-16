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

type crudSet map[constants.OperationCRUD]bool

type roleToCRUDMap map[string]*crudSet

//RBAC request structure .
type request struct {
	ruleHashTbl map[string]*roleToCRUDMap // Hash table with objtype key  to roleToCRUDMap
	op          constants.OperationCRUD   // Requested CRUD operation
	roles       map[string]bool           // Roles of the user making request
	kind        string                    // Objtype of the resource
	domain      string                    // Domain name of the user.
	project     string                    // project name of the user.
}

// Create a new crudSet.
func newCRUDSet() *crudSet {

	var n crudSet
	n = make(map[constants.OperationCRUD]bool)
	return &n
}

// Add an operation to crudSet.
func (s *crudSet) setCRUD(crud string) {

	for _, c := range crud {
		op := crudToOpMap(c)
		if op != constants.OpInvalid {
			(*s)[op] = true
		}
	}
	return
}

// Check whether an operation is is present in crudSet.
func (s *crudSet) isOpAllowed(op constants.OperationCRUD) bool {

	if _, ok := (*s)[op]; ok {
		return true

	}
	return false
}

func newRoleToCRUDMap() *roleToCRUDMap {

	var n roleToCRUDMap
	n = make(map[string]*crudSet)
	return &n
}

//Add a new role to crud mapping entry. For example, "Member" ==> "CR".
func (m *roleToCRUDMap) updateEntry(role string, crud string) {

	if cs, ok := (*m)[role]; ok {

		cs.setCRUD(crud)

	} else {
		cs := newCRUDSet()
		cs.setCRUD(crud)
		(*m)[role] = cs
	}
	return
}

// Check whether there is any rule is allowing this op.
func (m *roleToCRUDMap) isRuleMatching(roleSet map[string]bool, op constants.OperationCRUD) bool {

	// Check against all user roles
	for role := range roleSet {
		if cs, ok := (*m)[role]; ok {

			return cs.isOpAllowed(op)
		}
	}
	// Check against wild card role rule
	if cs, ok := (*m)["*"]; ok {

		return cs.isOpAllowed(op)
	}
	return false

}

// Create a new request Object .This object will be initilized with user roles
// and relevant rbac rules hash.
func newRequest(ctx context.Context, apiAccessRules []*models.APIAccessList,
	kind string, op constants.OperationCRUD) *request {

	var n request

	n.kind = kind
	n.op = op
	authCtx := auth.GetAuthCTX(ctx)

	//Initialize project and domain information.
	n.domain = authCtx.DomainID()
	n.project = authCtx.ProjectID()

	//Initialize the roles of the user from context .
	n.roles = getUserRoles(ctx)
	n.ruleHashTbl = make(map[string]*roleToCRUDMap)

	//Add relevant RBAC rules for this object type to request.ruleHashTbl
	n.addAPIAccessRulesToHash(apiAccessRules)
	return &n
}

// Return whether RBAC is enabled.
func isRBACEnabled() bool {

	return viper.GetString("aaa_mode") == "rbac"
}

// Convert user roles from ctx to a map.
func getUserRoles(ctx context.Context) map[string]bool {

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

// Filter applicable RBAC rules based on attachment point . Include all RBAC rules from global-system config.
// Include rules only from matching domain and project
func (rq *request) IsMatchingAttachmentPoint(apiAccessRule *models.APIAccessList) bool {

	if apiAccessRule == nil {
		return false
	}

	parentType := apiAccessRule.GetParentType()
	fqName := apiAccessRule.GetFQName()

	if parentType == models.KindGlobalSystemConfig {
		return true
	}

	if parentType == models.KindDomain {

		if rq.domain == fqName[0] {
			return true
		}

		return false
	}

	if parentType == models.KindProject {

		if rq.project == fqName[1] {
			return true
		}
		return false
	}

	return false
}

//  Form a hash key from objName,objField,role
func (rq *request) getHashKey(kind string) (key string) {

	// Not implemented field name based rules yet
	return kind + "." + "*"

}

// Add rules to the rule hash table. Rule hash table will have entries of the form
// key => {{role1:crudSet1},{role2:crudSet2}} . This will merge the object_type CRUD in global,domain and
// project level.
func (rq *request) updateRuleToHash(role string, crud string) {

	key := rq.getHashKey(rq.kind)

	if val, ok := rq.ruleHashTbl[key]; ok {

		val.updateEntry(role, crud)

	} else {
		val := newRoleToCRUDMap()
		val.updateEntry(role, crud)
		rq.ruleHashTbl[key] = val
	}

	return

}

//  Add objectType or wildcard RBAC rules to hash.
//  Input "perm" is in the following form
//   Member:CRUD
//
func (rq *request) objectTypePermAddToHash(perm *models.RbacPermType) {

	if perm == nil {
		return
	}
	role := perm.GetRoleName()
	crud := perm.GetRoleCrud()

	if _, ok := rq.roles[role]; ok {

		rq.updateRuleToHash(role, crud)
	}
	return

}

//  Add objectType or wildcard RBAC rules to hash.
//  Input "perms" is in the  following form
//  Member:CRUD ,Development:CRUD
//
func (rq *request) objectTypePermsAddToHash(perms []*models.RbacPermType) {

	for _, perm := range perms {

		rq.objectTypePermAddToHash(perm)
	}
	return

}

//  Add obj_type or wildcard RBAC rules to hash.
//  Input rbacRule is in the the following form
//  < virtual-network, *> => Member:CRUD ,Development:CRUD
//
func (rq *request) objectTypeRuleAddToHash(rbacRule *models.RbacRuleType) {

	if rbacRule == nil {
		return
	}
	ruleObjType := rbacRule.GetRuleObject()

	if ruleObjType != rq.kind && ruleObjType != "*" {
		return
	}

	perms := rbacRule.GetRulePerms()

	rq.objectTypePermsAddToHash(perms)

	return

}

// Add  matching the objectType or wildcard RBAC rules to hash.
//   Input is in the  following form
//   < virtual-network, *> => Member:CRUD ,Development:CRUD
//   < virtual-ip, *> => Member:CRUD
//
func (rq *request) rulesAddToHash(rbacRules []*models.RbacRuleType) {

	for _, obj := range rbacRules {

		rq.objectTypeRuleAddToHash(obj)

	}
	return
}

// Add rules matching objectType and wild card rules from the rule list to request Hash.
// For example input apiAccessRules may contains rules from global-config, default-domain & project
// different all object types. There can be multiple roles and CRUD for every object type.
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
// This Function will do the following
//  i) Filter Rules based on project and domain
// ii) Create request.hash in the following form
//   objtype.* ==>{ role1 : crudSet1  ,  role2 : crudSet2 }
// For example result could be
//  virtual-network.* ==> { Member 	 : { OpCreate:true,OpDelete:true  }  ,
//			    Development  : { OpUpdate:true } }
//
func (rq *request) addAPIAccessRulesToHash(apiAccessRules []*models.APIAccessList) {

	// TODO Filter rules based on domain and project
	for _, apiRule := range apiAccessRules {

		if rq.IsMatchingAttachmentPoint(apiRule) {
			rbacRules := apiRule.GetRBACRules()
			rq.rulesAddToHash(rbacRules)
		}
	}

	return

}

//  Checks whether any hash entry is matching the key.
//
func (rq *request) validateResourceAPILevel(kind string) bool {

	key := rq.getHashKey(kind)
	if val, ok := rq.ruleHashTbl[key]; ok {

		return val.isRuleMatching(rq.roles, rq.op)

	}
	return false

}

//  Checks whether any obj type rules or wildcard allow this operation
//  for any roles which user is part of.
func (rq *request) validateAPILevel() bool {

	// check for objtype rule .
	if rq.validateResourceAPILevel(rq.kind) {
		return true
	}

	// check for wild card rule
	if rq.validateResourceAPILevel("*") {
		return true
	}

	return false
}

//'CRUD' to operation  mapping
func crudToOpMap(crud rune) constants.OperationCRUD {

	var res constants.OperationCRUD

	switch crud {
	case constants.UpperCaseC:
		res = constants.OpCreate
	case constants.UpperCaseR:
		res = constants.OpRead
	case constants.UpperCaseU:
		res = constants.OpUpdate
	case constants.UpperCaseD:
		res = constants.OpDelete
	default:
		res = constants.OpInvalid
		log.Debug("Invalid operation")
	}
	return res
}

// Checks whether is Read only  Resource access is allowed if operation is READ.
//
func isROAccessAllowed(ctx *auth.Context, op constants.OperationCRUD) bool {

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
func validateRoleAndConfig(ctx *auth.Context) bool {

	// RBAC is not configured
	if !isRBACEnabled() {
		return true
	}

	// If User is cloud_admin allow access

	if ctx.IsAdmin() {
		return true
	}

	return false
}

//ValidateRequest checks whether resource operation is allowed based on RBAC config.
func ValidateRequest(ctx context.Context, apiAccessRules []*models.APIAccessList,
	kind string, op constants.OperationCRUD) error {

	if ctx == nil {
		return errutil.ErrorPermissionDenied
	}

	authCtx := auth.GetAuthCTX(ctx)

	if authCtx == nil {
		return errutil.ErrorPermissionDenied
	}

	// Check whether access is allowed based on role or config

	if validateRoleAndConfig(authCtx) {
		return nil
	}

	// Check whether access is allowed based on global read only role.

	if isROAccessAllowed(authCtx, op) {
		return nil
	}

	// If no rules to allow access, deny permission

	if len(apiAccessRules) == 0 || authCtx == nil {

		return errutil.ErrorPermissionDenied
	}
	// Create RBAC request object based on request parameters and rbac rule configuration
	rbacReq := newRequest(ctx, apiAccessRules, kind, op)

	// Do the RBAC check
	if rbacReq.validateAPILevel() {

		return nil
	}

	log.Debug("RBAC : Access Denied ")
	return errutil.ErrorPermissionDenied

}
