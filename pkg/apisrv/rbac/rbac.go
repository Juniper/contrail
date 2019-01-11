package rbac

import (
	"context"
	"strings"

	"github.com/Juniper/contrail/pkg/auth"
	"github.com/Juniper/contrail/pkg/errutil"
	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/models/basemodels"
)

// CRUD operations rune constants
const (
	UpperCaseC rune = 'C'
	UpperCaseR rune = 'R'
	UpperCaseU rune = 'U'
	UpperCaseD rune = 'D'
)

// Action enumerates CRUD actions.
type Action int

// Enumeration of Actions.
const (
	ActionInvalid Action = iota
	ActionCreate
	ActionRead
	ActionUpdate
	ActionDelete
)

const aaaModeRBAC = "rbac"

type actionSet map[Action]struct{}

// allowedActions is a hash table of roles and it's allowed actions.
type allowedActions map[string]actionSet

type resourceKey string

// resourcePermission is the 2st level hash table to get objectKind specific rules.
type resourcePermission map[resourceKey]allowedActions

type tenantKey string

// tenantPermission is the 1st level hash table to get tenant specific rules.
type tenantPermission map[tenantKey]resourcePermission

// request structure.
type request struct {
	op      Action          // Requested CRUD operation.
	roles   map[string]bool // Roles of the user making request.
	kind    string          // kind of the resource.
	domain  string          // Domain name of the user.
	project string          // project name of the user.
}

// CheckPermissions checks whether resource operation is allowed based on RBAC config.
func CheckPermissions(ctx context.Context, l []*models.APIAccessList, aaaMode string, kind string, op Action) error {

	allowed, err := checkConfig(ctx, aaaMode, kind, op)
	if err != nil || allowed == true {
		return err
	}

	if len(l) == 0 {
		return errutil.ErrorForbidden("no API access list rules present.access denied ")
	}

	// hash table with RBAC rules. Currently it will be populated for every request. Later
	// it will be enhanced to get updated only with a database change for APIAccessList config.

	g := make(tenantPermission)
	err = g.addAccessList(l)
	if err != nil {
		return err
	}

	if !g.validateAPILevel(newRequest(ctx, kind, op)) {
		return errutil.ErrorForbidden("no matching  API access list rules. access denied ")
	}
	return nil
}

func checkConfig(ctx context.Context, aaaMode string, kind string, op Action) (isAllowed bool, err error) {

	if !isRBACEnabled(aaaMode) {
		return true, nil
	}
	if ctx == nil {
		return false, errutil.ErrorForbidden("invalid context. access denied")
	}

	authCtx := auth.GetAuthCTX(ctx)

	if authCtx == nil {
		return false, errutil.ErrorForbidden("invalid auth context. access denied ")
	}

	if validateRole(authCtx) {
		return true, nil
	}

	if isROAccessAllowed(authCtx, op) {
		return true, nil
	}
	return false, nil
}

// getTenantKey get the Hash key for first level hash.
func getTenantKey(domainName string, projectName string) tenantKey {
	return tenantKey(domainName + "." + projectName)
}

// addTenantLevelPolicy create a project specific policy.
func (t tenantPermission) addTenantLevelPolicy(
	global []*models.APIAccessList,
	domain map[string][]*models.APIAccessList,
	project map[tenantKey][]*models.APIAccessList,
) error {
	for k, list := range project {

		s := strings.Split(string(k), ".")
		domainName := s[0]
		domainList := domain[domainName]
		list = append(list, global...)
		list = append(list, domainList...)
		entry, err := t.addAPIAccessRules(list)
		if err != nil {
			return err
		}
		t[k] = entry
	}
	return nil
}

// addDomainLevelPolicy create a domain specific policy.
func (t tenantPermission) addDomainLevelPolicy(
	global []*models.APIAccessList, domain map[string][]*models.APIAccessList) error {
	for k, list := range domain {
		list = append(list, global...)
		key := getTenantKey(k, "*")
		entry, err := t.addAPIAccessRules(list)
		if err != nil {
			return err
		}
		t[key] = entry
	}
	return nil
}

// addGlobalLevelPolicy create a  global policy.
func (t tenantPermission) addGlobalLevelPolicy(global []*models.APIAccessList) error {
	key := getTenantKey("*", "*")
	entry, err := t.addAPIAccessRules(global)
	if err != nil {
		return err
	}
	t[key] = entry
	return nil
}

// addAccessList creates project, domain and  global policies.
func (t tenantPermission) addAccessList(lists []*models.APIAccessList) error {
	global := make([]*models.APIAccessList, 0)
	domains := make(map[string][]*models.APIAccessList)
	projects := make(map[tenantKey][]*models.APIAccessList)

	for _, list := range lists {

		parentType := list.GetParentType()
		fqName := list.GetFQName()

		switch parentType {
		case models.KindGlobalSystemConfig:
			global = append(global, list)
		case models.KindDomain:
			domainName := fqName[0]
			domains[domainName] = append(domains[domainName], list)
		case models.KindProject:
			domainName, projectName := fqName[0], fqName[1]
			key := getTenantKey(domainName, projectName)
			projects[key] = append(projects[key], list)
		default:
			return errutil.ErrorInternalf("invalid parent type present in RBAC rule.")
		}
	}

	err := t.addTenantLevelPolicy(global, domains, projects)
	if err != nil {
		return err
	}

	err = t.addDomainLevelPolicy(global, domains)
	if err != nil {
		return err
	}

	err = t.addGlobalLevelPolicy(global)
	if err != nil {
		return err
	}
	return nil
}

// addAPIAccessRules adds rules matching (resource )objectKind and wild card rules from the rule list to
// resourcePermission table.
// For example input apiAccessRules may contains rules from global-config, default-domain & project
// different all resource types. There can be multiple roles and CRUD for every resource type.
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
//  objectKind.* ==>{ role1 : actionSet1  ,  role2 : actionSet2 }
// For example result could be
// virtual-network.* ==> { Member 	 : { ActionCreate:true,ActionDelete:true  }  ,
//			    Development  : { ActionUpdate:true } }
func (t tenantPermission) addAPIAccessRules(apiAccessRules []*models.APIAccessList) (resourcePermission, error) {
	res := make(resourcePermission)
	for _, apiRule := range apiAccessRules {

		rbacRules := getRBACRules(apiRule)
		err := res.rulesAdd(rbacRules)
		if err != nil {
			return nil, err
		}
	}

	return res, nil
}

// getResourcePermission retrieves matching first level hash table entry .
func (t tenantPermission) getResourcePermission(rq *request) resourcePermission {
	key := getTenantKey(rq.domain, rq.project)

	if val, ok := t[key]; ok {
		return val
	}
	key = getTenantKey(rq.domain, "*")

	if val, ok := t[key]; ok {
		return val
	}
	key = getTenantKey("*", "*")

	if val, ok := t[key]; ok {
		return val
	}

	return nil
}

// validateResourceAPILevel checks whether any hash entry is matching the key.
func (t *tenantPermission) validateResourceAPILevel(rq *request, kind string) bool {
	r := t.getResourcePermission(rq)

	if r == nil {
		return false
	}

	key := getResourceKey(kind)
	if val, ok := r[key]; ok {

		return val.isRuleMatching(rq.roles, rq.op)

	}
	return false
}

// CheckObjectPermissions checks object level (perms2) permissions.
func CheckObjectPermissions(ctx context.Context, p *models.PermType2, aaaMode string, kind string, op Action) error {
	allowed, err := checkConfig(ctx, aaaMode, kind, op)
	if err != nil || allowed == true {
		return err
	}

	rq := newRequest(ctx, kind, op)

	if p == nil {
		return errutil.ErrorForbidden("Invalid perms2 field.  access denied ")
	}

	// Check whether resource global access permissions allows  the Action.
	if permsAccessAllowed(rq, p.GetGlobalAccess()) {
		return nil
	}

	// Check whether resource owner access  permissions allows the Action.
	if rq.project == p.GetOwner() {
		if !permsAccessAllowed(rq, p.GetOwnerAccess()) {
			return errutil.ErrorForbidden("Object access not allowed. access denied ")
		}
		return nil
	}
	return checkObjectSharePermissions(ctx, p, rq)
}

func checkObjectSharePermissions(ctx context.Context, p *models.PermType2, rq *request) error {
	share := p.GetShare()
	// Check whether resource is shared with the tenant and Action is allowed.
	for _, st := range share {
		tenant := st.GetTenant()
		tAccess := st.GetTenantAccess()
		shareType, uuid := tenantInfo(tenant)

		if (shareType == models.KindDomain && uuid == rq.domain) ||
			(shareType == models.KindProject && uuid == rq.project) {
			if permsAccessAllowed(rq, tAccess) {
				return nil
			}
		}
	}
	return errutil.ErrorForbiddenf("Object access not allowed. access denied ")
}

func tenantType(tenant string) string {
	if strings.HasPrefix(tenant, "domain") {
		return models.KindDomain
	}
	return models.KindProject
}

func tenantInfo(tenant string) (projectType string, uuid string) {
	uuid = strings.Split(tenant, ":")[1]
	if tenantType(tenant) == models.KindDomain {
		return models.KindDomain, uuid
	}
	return models.KindProject, uuid
}

func actionToPerms(a Action) int64 {
	switch a {
	case ActionCreate:
		return basemodels.PermsW
	case ActionDelete:
		return basemodels.PermsW
	case ActionUpdate:
		return basemodels.PermsW
	case ActionRead:
		return basemodels.PermsR
	}
	return basemodels.PermsR
}

func permsAccessAllowed(rq *request, access int64) bool {
	p := actionToPerms(rq.op)
	if p&access != 0 {
		return true
	}
	return false
}

// validateAPILevel checks whether any resource rules or wildcard allow this operation
// for any roles which user is part of.
func (t *tenantPermission) validateAPILevel(rq *request) bool {
	// check for Resource  rule .
	if t.validateResourceAPILevel(rq, rq.kind) {
		return true
	}

	// check for wild card rule
	if t.validateResourceAPILevel(rq, "*") {
		return true
	}
	return false
}

// crudToActionSet converts crud string to op slice.
func crudToActionSet(crud string) ([]Action, error) {
	var res []Action

	for _, c := range crud {
		switch c {
		case UpperCaseC:
			res = append(res, ActionCreate)
		case UpperCaseR:
			res = append(res, ActionRead)
		case UpperCaseU:
			res = append(res, ActionUpdate)
		case UpperCaseD:
			res = append(res, ActionDelete)
		default:
			return nil, errutil.ErrorInternalf("invalid crud action present in RBAC rule.")
		}
	}
	return res, nil
}

// isROAccessAllowed checks whether is Read only  Resource access is allowed if operation is READ.
func isROAccessAllowed(ctx *auth.Context, op Action) bool {
	if op != ActionRead {
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

// updateEntry adds a new role to action  mapping entry. For example, "Member" ==> "CR".
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

// isRuleMatching checks whether there is any rule is allowing this op.
func (m allowedActions) isRuleMatching(roleSet map[string]bool, op Action) bool {
	// Check against all user roles
	for role := range roleSet {
		if cs, ok := m[role]; ok {

			return cs.isActionAllowed(op)
		}
	}
	// Check against wild card role rule
	if cs, ok := m["*"]; ok {

		return cs.isActionAllowed(op)
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

// isActionAllowed checks whether an operation is is present in actionSet.
func (s actionSet) isActionAllowed(op Action) bool {
	if _, ok := s[op]; ok {
		return true

	}
	return false
}

// newRequest creates a new request Object. This object will be initialized with user roles
// project,domain ,objectKind info.
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

// getUserRoles converts user roles from ctx to a map.
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

// getResourceKey forms a hash key from objName,objField.
func getResourceKey(kind string) (key resourceKey) {
	// Not implemented field name based rules yet
	return resourceKey(kind + "." + "*")
}

// updateRule adds rules to the resourcePermission hash table. resourcePermission hash table will have
// entries of the form  key => {{role1:actionSet1},{role2:actionSet2}}. This will merge the resource CRUD in
// global,domain and  project level.
func (r resourcePermission) updateRule(ruleKind string, role string, crud []Action) {
	key := getResourceKey(ruleKind)

	if val, ok := r[key]; ok {
		val.updateEntry(role, crud)

	} else {
		val := newAllowedAction()
		val.updateEntry(role, crud)
		r[key] = val
	}

	return
}

// objectKindPermsAdd adds objectKind or wildcard RBAC rules to hash.
// Input "perm" is in the following form
// Member:CRUD
func (r resourcePermission) objectKindPermAdd(ruleKind string, perm *models.RbacPermType) error {
	if perm == nil {
		return nil
	}
	role := perm.GetRoleName()
	crud, err := crudToActionSet(perm.GetRoleCrud())
	if err != nil {
		return err
	}
	r.updateRule(ruleKind, role, crud)
	return nil
}

// objectKindPermsAdd adds objectKind or wildcard RBAC rules to hash.
// Input "perms" is in the  following form
// Member:CRUD ,Development:CRUD
func (r resourcePermission) objectKindPermsAdd(ruleKind string, perms []*models.RbacPermType) error {
	for _, perm := range perms {
		err := r.objectKindPermAdd(ruleKind, perm)
		if err != nil {
			return err
		}
	}
	return nil
}

// objectKindRuleAdd add objectKind or wildcard RBAC rules to hash.
// Input rbacRule is in the the following form
// < virtual-network, *> => Member:CRUD ,Development:CRUD.
func (r resourcePermission) objectKindRuleAdd(rbacRule *models.RbacRuleType) error {
	if rbacRule == nil {
		return nil
	}

	ruleKind := pluralNameToKind(rbacRule.GetRuleObject())
	perms := rbacRule.GetRulePerms()

	err := r.objectKindPermsAdd(ruleKind, perms)
	if err != nil {
		return err
	}
	return nil
}

// rulesAdd adds  matching the objectKind or wildcard RBAC rules to hash.
// Input is in the  following form
// < virtual-network, *> => Member:CRUD ,Development:CRUD
// < virtual-ip, *> => Member:CRUD
func (r resourcePermission) rulesAdd(rbacRules []*models.RbacRuleType) error {
	for _, obj := range rbacRules {

		err := r.objectKindRuleAdd(obj)
		if err != nil {
			return err
		}
	}
	return nil
}

// getRBACRules retrieves RBAC rules from API access list entry.
func getRBACRules(l *models.APIAccessList) []*models.RbacRuleType {
	return l.GetAPIAccessListEntries().GetRbacRule()
}

// pluralNameToKind converts plural object name to object kind.  eg.virtual-networks => virtual-network.
func pluralNameToKind(name string) string {
	return strings.TrimSuffix(name, "s")
}

// isRBACEnabled returns whether RBAC is enabled.
func isRBACEnabled(aaaMode string) bool {

	return aaaMode == aaaModeRBAC
}

// validateRole checks whether is Resource access is allowed purely based on  Role.
func validateRole(ctx *auth.Context) bool {

	// If User is cloud_admin allow access

	if ctx.IsAdmin() {
		return true
	}
	return false
}
