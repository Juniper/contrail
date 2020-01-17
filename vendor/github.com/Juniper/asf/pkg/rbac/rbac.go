package rbac

import (
	"context"
	"strings"

	"github.com/Juniper/asf/pkg/auth"
	"github.com/Juniper/asf/pkg/errutil"
	"github.com/Juniper/asf/pkg/models/basemodels"
)

// CRUD operations rune constants
const (
	UpperCaseC rune = 'C'
	UpperCaseR rune = 'R'
	UpperCaseU rune = 'U'
	UpperCaseD rune = 'D'
)

const aaaModeRBAC = "rbac"

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

// Permission scopes
const (
	ScopeDomain             = "domain"
	ScopeProject            = "project"
	ScopeGlobalSystemConfig = "global-system-config"
)

// PermType2 is used to check object level (perms2) permissions.
type PermType2 struct {
	Owner        string
	OwnerAccess  int64
	GlobalAccess int64
	Share        []*ShareType
}

// ShareType is used to check whether a resource is shared with the tenant
type ShareType struct {
	TenantAccess int64
	Tenant       string
}

func (s *ShareType) getTenantAccess() int64 {
	if s != nil {
		return s.TenantAccess
	}
	return 0
}

func (s *ShareType) getTenant() string {
	if s != nil {
		return s.Tenant
	}
	return ""
}

func (p *PermType2) getOwner() string {
	if p != nil{
		return p.Owner
	}
	return ""
}

func (p *PermType2) getOwnerAccess() int64 {
	if p != nil {
		return p.OwnerAccess
	}
	return 0
}

func (p *PermType2) getGlobalAccess() int64 {
	if p != nil {
		return p.GlobalAccess
	}
	return 0
}

func (p *PermType2) getShare() []*ShareType {
	if p != nil {
		return p.Share
	}
	return nil
}

// APIAccessList structure
type APIAccessList struct {
	ParentType           string
	FQName               []string
	APIAccessListEntries *RuleEntriesType
}

// RuleEntriesType used for entries
type RuleEntriesType struct {
	Rule []*RuleType
}

// RuleType used for entry
type RuleType struct {
	RuleObject string
	RulePerms  []*PermType
}

// PermType is used to check object level permissions.
type PermType struct {
	RoleCrud string
	RoleName string
}

func (r *PermType) getRoleCrud() string {
	if r != nil {
		return r.RoleCrud
	}
	return ""
}

func (r *PermType) getRoleName() string {
	if r != nil {
		return r.RoleName
	}
	return ""
}

func (r *RuleType) getRuleObject() string {
	if r != nil {
		return r.RuleObject
	}
	return ""
}

func (m *RuleType) getRulePerms() []*PermType {
	if m != nil {
		return m.RulePerms
	}
	return nil
}

func (r *RuleEntriesType) getRuleType() []*RuleType {
	if r != nil {
		return r.Rule
	}
	return nil
}

func (a *APIAccessList) getParentType() string {
	if a != nil {
		return a.ParentType
	}
	return ""
}

func (a *APIAccessList) getFQName() []string {
	if a != nil {
		return a.FQName
	}
	return nil
}

func (a *APIAccessList) getAPIAccessListEntries() *RuleEntriesType {
	if a != nil {
		return a.APIAccessListEntries
	}
	return nil
}

// CheckCommonPermissions checks checks whether access is based on RBAC configuration or user roles..
func CheckCommonPermissions(ctx context.Context, aaaMode string, kind string, op Action) (isAllowed bool, err error) {
	if !isRBACEnabled(aaaMode) {
		return true, nil
	}
	if ctx == nil {
		return false, errutil.ErrorForbidden("invalid context. access denied")
	}

	authCtx := auth.GetIdentity(ctx)

	if authCtx == nil {
		return false, errutil.ErrorForbidden("invalid auth context. access denied ")
	}

	if authCtx.IsAdmin() {
		return true, nil
	}

	if isROAccessAllowed(authCtx, op) {
		return true, nil
	}
	return false, nil
}

// isRBACEnabled returns whether RBAC is enabled.
func isRBACEnabled(aaaMode string) bool {
	return aaaMode == aaaModeRBAC
}

// isROAccessAllowed checks whether is Read only  Resource access is allowed if operation is READ.
func isROAccessAllowed(ctx auth.Identity, op Action) bool {
	if op != ActionRead {
		return false
	}

	if ctx.IsGlobalRORole() {
		return true
	}
	return false
}

// CheckPermissions checks whether resource operation is allowed based on RBAC config.
func CheckPermissions(ctx context.Context, l []*APIAccessList, aaaMode string, kind string, op Action) error {
	if len(l) == 0 {
		return errutil.ErrorForbidden("no API access list rules present.access denied ")
	}

	// hash table with RBAC rules. Currently it will be populated for every request. Later
	// it will be enhanced to get updated only with a database change for APIAccessList config.

	g := make(tenantPermissions)
	err := g.AddAccessList(l)
	if err != nil {
		return err
	}

	if !g.ValidateAPILevel(newRequest(ctx, kind, op)) {
		return errutil.ErrorForbidden("no matching  API access list rules. access denied ")
	}
	return nil
}

// CheckObjectPermissions checks object level (perms2) permissions.
func CheckObjectPermissions(ctx context.Context, p *PermType2, aaaMode string, kind string, op Action) error {
	rq := newRequest(ctx, kind, op)

	if p == nil {
		return errutil.ErrorForbidden("invalid perms2 field.  access denied ")
	}

	// Check whether resource global access permissions allows  the Action.
	if permsAccessAllowed(rq, p.getGlobalAccess()) {
		return nil
	}

	// Check whether resource owner access  permissions allows the Action.
	if rq.project == p.getOwner() {
		if !permsAccessAllowed(rq, p.getOwnerAccess()) {
			return errutil.ErrorForbidden("object access not allowed. access denied ")
		}
		return nil
	}
	return checkObjectSharePermissions(ctx, p, rq)
}

// request structure.
type request struct {
	op      Action          // Requested CRUD operation.
	roles   map[string]bool // Roles of the user making request.
	kind    string          // kind of the resource.
	domain  string          // Domain name of the user.
	project string          // project name of the user.
}

// newRequest creates a new request Object. This object will be initialized with user roles
// project,domain ,objectKind info.
func newRequest(ctx context.Context, kind string, op Action) *request {
	authCtx := auth.GetIdentity(ctx)
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

	aCtx := auth.GetIdentity(ctx)
	if aCtx == nil {
		return nil
	}

	rset := make(map[string]bool)
	for _, role := range aCtx.Roles() {
		rset[role] = true
	}
	return rset
}

func checkObjectSharePermissions(ctx context.Context, p *PermType2, rq *request) error {
	share := p.getShare()
	// Check whether resource is shared with the tenant and Action is allowed.
	for _, st := range share {
		tenant := st.getTenant()
		tAccess := st.getTenantAccess()
		shareType, uuid := tenantInfo(tenant)

		if (shareType == ScopeDomain && uuid == rq.domain) ||
			(shareType == ScopeProject && uuid == rq.project) {
			if permsAccessAllowed(rq, tAccess) {
				return nil
			}
		}
	}
	return errutil.ErrorForbiddenf("object access not allowed. access denied ")
}

func tenantInfo(tenant string) (projectType string, uuid string) {
	uuid = strings.Split(tenant, ":")[1]
	if tenantType(tenant) == ScopeDomain {
		return ScopeDomain, uuid
	}
	return ScopeProject, uuid
}

func tenantType(tenant string) string {
	if strings.HasPrefix(tenant, "domain") {
		return ScopeDomain
	}
	return ScopeProject
}

func permsAccessAllowed(rq *request, access int64) bool {
	p := actionToPerms(rq.op)
	if p&access != 0 {
		return true
	}
	return false
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
