package rbac

import (
	"context"
	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	"github.com/Juniper/contrail/pkg/auth"
	"github.com/Juniper/contrail/pkg/constants"
	"github.com/Juniper/contrail/pkg/errutil"
	"github.com/Juniper/contrail/pkg/models"
)

// CRUD operations rune constants
const (
	UpperCaseC rune = 'C'
	UpperCaseR rune = 'R'
	UpperCaseU rune = 'U'
	UpperCaseD rune = 'D'
)

type crudSet map[constants.OperationCRUD]struct{}

type roleToCRUDMap map[string]crudSet

type resourceHashKey string

// request structure.
type request struct {
	rolePermissions map[resourceHashKey]roleToCRUDMap // Hash table with objtype key  to roleToCRUDMap
	op              constants.OperationCRUD           // Requested CRUD operation
	roles           map[string]bool                   // Roles of the user making request
	kind            string                            // Objtype of the resource
	domain          string                            // Domain name of the user.
	project         string                            // project name of the user.
}

// ValidateRequest checks whether resource operation is allowed based on RBAC config.
func ValidateRequest(ctx context.Context, l []*models.APIAccessList, kind string, op constants.OperationCRUD) error {

	if ctx == nil {

		return errutil.ErrorForbiddenf("Invalid Context. Access Denied ")
	}

	authCtx := auth.GetAuthCTX(ctx)

	if authCtx == nil {

		return errutil.ErrorForbiddenf("Invalid Auth Context. Access Denied ")
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

	if len(l) == 0 {

		return errutil.ErrorForbiddenf("No API Access List rules present.Access Denied ")
	}
	// Create RBAC request object based on request parameters and rbac rule configuration and do  RBAC check
	if newRequest(ctx, l, kind, op).validateAPILevel() {
		return nil
	}

	return errutil.ErrorForbiddenf("RBAC : No Matching  API Access List rules. Access Denied ")

}

// Create a new crudSet.
func newCRUDSet() crudSet {
	return make(map[constants.OperationCRUD]struct{})
}

// Add an operation to crudSet.
func (s *crudSet) setCRUD(crud string) {

	for _, c := range crud {
		op := crudToOpMap(c)
		if op != constants.OpInvalid {
			(*s)[op] = struct{}{}
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

func newRoleToCRUDMap() roleToCRUDMap {

	return make(map[string]crudSet)
}

// Add a new role to crud mapping entry. For example, "Member" ==> "CR".
func (m *roleToCRUDMap) updateEntry(role string, crud string) {

	if cs, ok := (*m)[role]; ok {

		cs.setCRUD(crud)

	} else {
		n := newCRUDSet()
		n.setCRUD(crud)
		(*m)[role] = n
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

// Create a new request Object .This object will be initailized with user roles
// and relevant RBAC rolePermissions hash.
func newRequest(
	ctx context.Context, apiAccessRules []*models.APIAccessList, kind string, op constants.OperationCRUD,
) *request {

	var n request

	n.kind = kind
	n.op = op
	authCtx := auth.GetAuthCTX(ctx)

	n.domain = authCtx.DomainID()
	n.project = authCtx.ProjectID()

	n.roles = getUserRoles(ctx)
	n.rolePermissions = make(map[resourceHashKey]roleToCRUDMap)

	//Add relevant RBAC rules for this object type to request.rolePermissions
	n.addAPIAccessRulesToHash(apiAccessRules)
	return &n
}

// Return whether RBAC is enabled.
func isRBACEnabled() bool {

	return viper.GetString("aaa_mode") == "rbac"
}

// Convert user roles from ctx to a map.
func getUserRoles(ctx context.Context) map[string]bool {

	if ctx == nil {
		return nil
	}

	aCtx := auth.GetAuthCTX(ctx)

	if aCtx == nil {
		return nil
	}

	rset := make(map[string]bool)
	for _, role := range aCtx.Roles() {
		rset[role] = true
	}

	return rset
}

// Filter applicable RBAC rules based on attachment point . Include all RBAC rules from global-system config.
// Include rules only from matching domain and project
func (rq *request) isMatchingAttachmentPoint(apiAccessRule *models.APIAccessList) bool {

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

// Form a hash key from objName,objField,role
func (rq *request) getHashKey(kind string) (key resourceHashKey) {

	// Not implemented field name based rules yet
	return resourceHashKey(kind + "." + "*")

}

// Add rules to the rule hash table. Rule hash table will have entries of the form
// key => {{role1:crudSet1},{role2:crudSet2}} . This will merge the object_type CRUD in global,domain and
// project level.
func (rq *request) updateRuleToHash(role string, crud string) {

	key := rq.getHashKey(rq.kind)

	if val, ok := rq.rolePermissions[key]; ok {

		val.updateEntry(role, crud)

	} else {
		val := newRoleToCRUDMap()
		val.updateEntry(role, crud)
		rq.rolePermissions[key] = val
	}

	return

}

// Add objectType or wildcard RBAC rules to hash.
// Input "perm" is in the following form
// Member:CRUD
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

// Add objectType or wildcard RBAC rules to hash.
// Input "perms" is in the  following form
// Member:CRUD ,Development:CRUD
func (rq *request) objectTypePermsAddToHash(perms []*models.RbacPermType) {

	for _, perm := range perms {

		rq.objectTypePermAddToHash(perm)
	}
	return

}

// Add obj_type or wildcard RBAC rules to hash.
// Input rbacRule is in the the following form
// < virtual-network, *> => Member:CRUD ,Development:CRUD
func (rq *request) objectTypeRuleAddToHash(rbacRule *models.RbacRuleType) {

	if rbacRule == nil {
		return
	}
	ruleKind := pluralNameToKind(rbacRule.GetRuleObject())

	if ruleKind != rq.kind && ruleKind != "*" {
		return
	}

	perms := rbacRule.GetRulePerms()

	rq.objectTypePermsAddToHash(perms)

	return

}

// Add  matching the objectType or wildcard RBAC rules to hash.
// Input is in the  following form
// < virtual-network, *> => Member:CRUD ,Development:CRUD
// < virtual-ip, *> => Member:CRUD
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
// i) Filter Rules based on project and domain
// ii) Create request.hash in the following form
//  objtype.* ==>{ role1 : crudSet1  ,  role2 : crudSet2 }
// For example result could be
// virtual-network.* ==> { Member 	 : { OpCreate:true,OpDelete:true  }  ,
//			    Development  : { OpUpdate:true } }
func (rq *request) addAPIAccessRulesToHash(apiAccessRules []*models.APIAccessList) {

	// TODO Filter rules based on domain and project
	for _, apiRule := range apiAccessRules {

		if rq.isMatchingAttachmentPoint(apiRule) {
			rbacRules := apiRule.GetRBACRules()
			rq.rulesAddToHash(rbacRules)
		}
	}

	return

}

// Checks whether any hash entry is matching the key.
func (rq *request) validateResourceAPILevel(kind string) bool {

	key := rq.getHashKey(kind)
	if val, ok := rq.rolePermissions[key]; ok {

		return val.isRuleMatching(rq.roles, rq.op)

	}
	return false

}

// Checks whether any obj type rules or wildcard allow this operation
// for any roles which user is part of.
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

// 'CRUD' to operation  mapping
func crudToOpMap(crud rune) constants.OperationCRUD {

	var res constants.OperationCRUD

	switch crud {
	case UpperCaseC:
		res = constants.OpCreate
	case UpperCaseR:
		res = constants.OpRead
	case UpperCaseU:
		res = constants.OpUpdate
	case UpperCaseD:
		res = constants.OpDelete
	default:
		res = constants.OpInvalid
		log.Debug("Invalid operation")
	}
	return res
}

// Checks whether is Read only  Resource access is allowed if operation is READ.
func isROAccessAllowed(ctx *auth.Context, op constants.OperationCRUD) bool {

	if op != constants.OpRead {
		return false
	}

	if ctx.IsGlobalRORole() {
		return true
	}

	return false
}

// Convert plural object name to object kind.  eg.virtual-networks => virtual-network
func pluralNameToKind(name string) string {

	return strings.TrimSuffix(name, "s")
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
