package rbac

import (
	"context"
	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	"github.com/Juniper/contrail/pkg/auth"
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

// Action CRUD opeartions type
type Action int

// CRUD operations
const (
	OpInvalid Action = iota
	OpCreate
	OpRead
	OpUpdate
	OpDelete
)

const aaaModeRBAC = "rbac"

type actionSet map[Action]struct{}

// allowedActions is a hash table of roles and it's allowed actions.
type allowedActions map[string]actionSet

type resourceKey string

// resourceID is the 2st level hash table to get objectType specific rules.
type resourceID map[resourceKey]allowedActions

type tenantKey string

// tenantPermission is the 1st level hash table to get tenant specific rules.
type tenantPermission map[tenantKey]resourceID

// globalPermission is the global hash table with RBAC rules. Currently it will be populated for every request. Later
// it will be enhanced to get updated only with a database change for APIAccessList config.
var globalPermission = make(tenantPermission)

// request structure.
type request struct {
	op      Action          // Requested CRUD operation.
	roles   map[string]bool // Roles of the user making request.
	kind    string          // kind of the resource.
	domain  string          // Domain name of the user.
	project string          // project name of the user.
}

// ValidateRequest checks whether resource operation is allowed based on RBAC config.
func ValidateRequest(ctx context.Context, l []*models.APIAccessList, kind string, op Action) error {

	if ctx == nil {

		return errutil.ErrorForbiddenf("invalid context. access denied")
	}

	authCtx := auth.GetAuthCTX(ctx)

	if authCtx == nil {

		return errutil.ErrorForbiddenf("invalid auth context. access denied ")
	}

	if validateRoleAndConfig(authCtx) {
		return nil
	}

	if isROAccessAllowed(authCtx, op) {
		return nil
	}

	// If no rules to allow access, deny permission

	if len(l) == 0 {

		return errutil.ErrorForbiddenf("no API access list rules present.access denied ")
	}

	globalPermission.accessListAdd(l)

	if globalPermission.validateAPILevel(newRequest(ctx, kind, op)) {
		return nil
	}

	return errutil.ErrorForbiddenf("RBAC : no matching  API access list rules. access denied ")

}

// Hash key for first level hash.
func getTenantKey(projectName string, domainName string) tenantKey {
	return tenantKey(projectName + "." + domainName)
}

// addTenantLevelPolicy create a project specific policy.
func (t tenantPermission) addTenantLevelPolicy(
	global []*models.APIAccessList,
	domain map[string][]*models.APIAccessList,
	project map[tenantKey][]*models.APIAccessList,
) {
	for k, list := range project {

		s := strings.Split(string(k), ".")
		domainName := s[0]
		domainList := domain[domainName]
		list = append(list, global...)
		list = append(list, domainList...)
		t[k] = t.addAPIAccessRulesToHash(list)
	}
}

// addDomainLevelPolicy create a domain specific policy.
func (t tenantPermission) addDomainLevelPolicy(
	global []*models.APIAccessList, domain map[string][]*models.APIAccessList) {
	for k, list := range domain {
		list = append(list, global...)
		key := getTenantKey(k, "*")
		t[key] = t.addAPIAccessRulesToHash(list)
	}
}

// addGlobalLevelPolicy create a  global policy.
func (t tenantPermission) addGlobalLevelPolicy(global []*models.APIAccessList) {

	key := getTenantKey("*", "*")
	t[key] = t.addAPIAccessRulesToHash(global)
}

// accessListAdd creates project, domain and  global policies  .
func (t tenantPermission) accessListAdd(lists []*models.APIAccessList) {

	global := make([]*models.APIAccessList, 0)
	domains := make(map[string][]*models.APIAccessList)
	projects := make(map[tenantKey][]*models.APIAccessList)

	for _, list := range lists {

		parentType := list.GetParentType()
		fqName := list.GetFQName()

		if parentType == models.KindGlobalSystemConfig {
			global = append(global, list)
		} else if parentType == models.KindDomain {
			domainName := fqName[0]
			domains[domainName] = append(domains[domainName], list)

		} else if parentType == models.KindProject {

			domainName, projectName := fqName[0], fqName[1]
			key := getTenantKey(domainName, projectName)
			projects[key] = append(projects[key], list)
		}
	}
	t.addTenantLevelPolicy(global, domains, projects)
	t.addDomainLevelPolicy(global, domains)
	t.addGlobalLevelPolicy(global)

}

// Add rules matching objectType and wild card rules from the rule list to resourceID table.
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
//  kind.* ==>{ role1 : actionSet1  ,  role2 : actionSet2 }
// For example result could be
// virtual-network.* ==> { Member 	 : { OpCreate:true,OpDelete:true  }  ,
//			    Development  : { OpUpdate:true } }
func (t tenantPermission) addAPIAccessRulesToHash(apiAccessRules []*models.APIAccessList) resourceID {

	res := make(resourceID)
	for _, apiRule := range apiAccessRules {

		rbacRules := getRBACRules(apiRule)
		res.rulesAddToHash(rbacRules)
	}

	return res
}

func (t tenantPermission) getProjectHashTable(rq *request) resourceID {

	key := getTenantKey(rq.domain, rq.project)

	if val, ok := t[key]; ok {
		return val
	}
	key = getTenantKey("*", rq.project)

	if val, ok := t[key]; ok {
		return val
	}
	key = getTenantKey("*", "*")

	if val, ok := t[key]; ok {
		return val
	}

	return nil
}

// Checks whether any hash entry is matching the key.
func (t *tenantPermission) validateResourceAPILevel(rq *request, kind string) bool {

	projectHashTable := t.getProjectHashTable(rq)

	if projectHashTable == nil {
		return false
	}

	key := getResourceHashKey(kind)
	if val, ok := projectHashTable[key]; ok {

		return val.isRuleMatching(rq.roles, rq.op)

	}
	return false
}

// Checks whether any obj type rules or wildcard allow this operation
// for any roles which user is part of.
func (t *tenantPermission) validateAPILevel(rq *request) bool {

	// check for objtype rule .
	if t.validateResourceAPILevel(rq, rq.kind) {
		return true
	}

	// check for wild card rule
	if t.validateResourceAPILevel(rq, "*") {
		return true
	}

	return false
}

// crudToOpMap converts crud string to op slice.
func crudToOpMap(crud string) []Action {

	var res []Action

	for _, c := range crud {
		switch c {
		case UpperCaseC:
			res = append(res, OpCreate)
		case UpperCaseR:
			res = append(res, OpRead)
		case UpperCaseU:
			res = append(res, OpUpdate)
		case UpperCaseD:
			res = append(res, OpDelete)
		default:
			log.Debug("Invalid crud operation")
		}
	}
	return res
}

// Checks whether is Read only  Resource access is allowed if operation is READ.
func isROAccessAllowed(ctx *auth.Context, op Action) bool {

	if op != OpRead {
		return false
	}

	if ctx.IsGlobalRORole() {
		return true
	}

	return false
}

func newAllowedAction() allowedActions {

	return make(map[string]actionSet)
}

// Add a new role to action  mapping entry. For example, "Member" ==> "CR".
func (m allowedActions) updateEntry(role string, crud []Action) {

	if cs, ok := m[role]; ok {

		cs.setCRUD(crud)

	} else {
		n := newCRUDSet()
		n.setCRUD(crud)
		m[role] = n
	}
	return
}

// Check whether there is any rule is allowing this op.
func (m allowedActions) isRuleMatching(roleSet map[string]bool, op Action) bool {

	// Check against all user roles
	for role := range roleSet {
		if cs, ok := m[role]; ok {

			return cs.isOpAllowed(op)
		}
	}
	// Check against wild card role rule
	if cs, ok := m["*"]; ok {

		return cs.isOpAllowed(op)
	}
	return false

}

// newCRUDSet create a new actionSet.
func newCRUDSet() actionSet {
	return make(map[Action]struct{})
}

// setCRUD add an operation to actionSet.
func (s actionSet) setCRUD(crud []Action) {

	for _, op := range crud {
		s[op] = struct{}{}
	}
	return
}

// Check whether an operation is is present in actionSet.
func (s actionSet) isOpAllowed(op Action) bool {

	if _, ok := s[op]; ok {
		return true

	}
	return false
}

// Create a new request Object. This object will be initialized with user roles
// project,domain ,objectType info.
func newRequest(
	ctx context.Context, kind string, op Action,
) *request {

	authCtx := auth.GetAuthCTX(ctx)
	n := request{
		kind:    kind,
		op:      op,
		domain:  authCtx.DomainID(),
		project: authCtx.ProjectID(),
		roles:   getUserRoles(ctx),
	}

	return &n
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

// Form a hash key from objName,objField,role
func getResourceHashKey(kind string) (key resourceKey) {

	// Not implemented field name based rules yet
	return resourceKey(kind + "." + "*")

}

// Add rules to the resourceID hash table. resourceID hash table will have entries of the form
// key => {{role1:actionSet1},{role2:actionSet2}} . This will merge the object_type CRUD in global,domain and
// project level.
func (r resourceID) updateRuleToHash(ruleKind string, role string, crud []Action) {

	key := getResourceHashKey(ruleKind)

	if val, ok := r[key]; ok {

		val.updateEntry(role, crud)

	} else {
		val := newAllowedAction()
		val.updateEntry(role, crud)
		r[key] = val
	}

	return

}

// Add objectType or wildcard RBAC rules to hash.
// Input "perm" is in the following form
// Member:CRUD
func (r resourceID) objectTypePermAddToHash(ruleKind string, perm *models.RbacPermType) {

	if perm == nil {
		return
	}
	role := perm.GetRoleName()
	crud := crudToOpMap(perm.GetRoleCrud())
	r.updateRuleToHash(ruleKind, role, crud)
	return

}

// Add objectType or wildcard RBAC rules to hash.
// Input "perms" is in the  following form
// Member:CRUD ,Development:CRUD
func (r resourceID) objectTypePermsAddToHash(ruleKind string, perms []*models.RbacPermType) {

	for _, perm := range perms {

		r.objectTypePermAddToHash(ruleKind, perm)
	}
	return

}

// Add objectType or wildcard RBAC rules to hash.
// Input rbacRule is in the the following form
// < virtual-network, *> => Member:CRUD ,Development:CRUD
func (r resourceID) objectTypeRuleAddToHash(rbacRule *models.RbacRuleType) {

	if rbacRule == nil {
		return
	}

	ruleKind := pluralNameToKind(rbacRule.GetRuleObject())
	perms := rbacRule.GetRulePerms()

	r.objectTypePermsAddToHash(ruleKind, perms)

	return

}

// Add  matching the objectType or wildcard RBAC rules to hash.
// Input is in the  following form
// < virtual-network, *> => Member:CRUD ,Development:CRUD
// < virtual-ip, *> => Member:CRUD
func (r resourceID) rulesAddToHash(rbacRules []*models.RbacRuleType) {

	for _, obj := range rbacRules {

		r.objectTypeRuleAddToHash(obj)

	}
	return
}

//getRBACRules retrieves RBAC rules from API access list entry.
func getRBACRules(l *models.APIAccessList) []*models.RbacRuleType {
	return l.GetAPIAccessListEntries().GetRbacRule()
}

// Convert plural object name to object kind.  eg.virtual-networks => virtual-network
func pluralNameToKind(name string) string {

	return strings.TrimSuffix(name, "s")
}

// Return whether RBAC is enabled.
func isRBACEnabled() bool {

	return viper.GetString("aaa_mode") == aaaModeRBAC
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
