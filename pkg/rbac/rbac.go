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

const (
	KindDomain             = "domain"
	KindProject            = "project"
	KindGlobalSystemConfig = "global-system-config"
)

type ShareType struct {
	TenantAccess int64  `protobuf:"varint,1,opt,name=tenant_access,json=tenantAccess,proto3" json:"tenant_access,omitempty" yaml:"tenant_access,omitempty"`
	Tenant       string `protobuf:"bytes,2,opt,name=tenant,proto3" json:"tenant,omitempty" yaml:"tenant,omitempty"`
}

func (s *ShareType) GetTenantAccess() int64 {
	if s != nil {
		return s.TenantAccess
	}
	return 0
}

func (s *ShareType) GetTenant() string {
	if s != nil {
		return s.Tenant
	}
	return ""
}

type PermType2 struct {
	Owner        string       `protobuf:"bytes,1,opt,name=owner,proto3" json:"owner" yaml:"owner,omitempty"`
	OwnerAccess  int64        `protobuf:"varint,2,opt,name=owner_access,json=ownerAccess,proto3" json:"owner_access" yaml:"owner_access,omitempty"`
	GlobalAccess int64        `protobuf:"varint,3,opt,name=global_access,json=globalAccess,proto3" json:"global_access" yaml:"global_access,omitempty"`
	Share        []*ShareType `protobuf:"bytes,4,rep,name=share,proto3" json:"share" yaml:"share,omitempty"`
}

func (p *PermType2) GetOwner() string {
	if p != nil{
		return p.Owner
	}
	return ""
}

func (p *PermType2) GetOwnerAccess() int64 {
	if p != nil {
		return p.OwnerAccess
	}
	return 0
}

func (p *PermType2) GetGlobalAccess() int64 {
	if p != nil {
		return p.GlobalAccess
	}
	return 0
}

func (p *PermType2) GetShare() []*ShareType {
	if p != nil {
		return p.Share
	}
	return nil
}

type PermType struct {
	RoleCrud string `protobuf:"bytes,1,opt,name=role_crud,json=roleCrud,proto3" json:"role_crud,omitempty" yaml:"role_crud,omitempty"`
	RoleName string `protobuf:"bytes,2,opt,name=role_name,json=roleName,proto3" json:"role_name,omitempty" yaml:"role_name,omitempty"`
}

func (r *PermType) GetRoleCrud() string {
	if r != nil {
		return r.RoleCrud
	}
	return ""
}

func (r *PermType) GetRoleName() string {
	if r != nil {
		return r.RoleName
	}
	return ""
}

type RuleType struct {
	RuleObject string      `protobuf:"bytes,2,opt,name=rule_object,json=ruleObject,proto3" json:"rule_object,omitempty" yaml:"rule_object,omitempty"`
	RulePerms  []*PermType `protobuf:"bytes,3,rep,name=rule_perms,json=rulePerms,proto3" json:"rule_perms,omitempty" yaml:"rule_perms,omitempty"`
}

func (r *RuleType) GetRuleObject() string {
	if r != nil {
		return r.RuleObject
	}
	return ""
}

func (m *RuleType) GetRulePerms() []*PermType {
	if m != nil {
		return m.RulePerms
	}
	return nil
}

type RuleEntriesType struct {
	Rule []*RuleType `protobuf:"bytes,1,rep,name=rbac_rule,json=rbacRule,proto3" json:"rbac_rule,omitempty" yaml:"rbac_rule,omitempty"`
}

func (r *RuleEntriesType) GetRbacRule() []*RuleType {
	if r != nil {
		return r.Rule
	}
	return nil
}

type APIAccessList struct {
	ParentType           string           `protobuf:"bytes,4,opt,name=parent_type,json=parentType,proto3" json:"parent_type,omitempty" yaml:"parent_type,omitempty"`
	FQName               []string         `protobuf:"bytes,5,rep,name=fq_name,json=fqName,proto3" json:"fq_name,omitempty" yaml:"fq_name,omitempty"`
	APIAccessListEntries *RuleEntriesType `protobuf:"bytes,12,opt,name=api_access_list_entries,json=apiAccessListEntries,proto3" json:"api_access_list_entries,omitempty" yaml:"api_access_list_entries,omitempty"`
}

func (a *APIAccessList) GetParentType() string {
	if a != nil {
		return a.ParentType
	}
	return ""
}

func (a *APIAccessList) GetFQName() []string {
	if a != nil {
		return a.FQName
	}
	return nil
}

func (a *APIAccessList) GetAPIAccessListEntries() *RuleEntriesType {
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
	if permsAccessAllowed(rq, p.GetGlobalAccess()) {
		return nil
	}

	// Check whether resource owner access  permissions allows the Action.
	if rq.project == p.GetOwner() {
		if !permsAccessAllowed(rq, p.GetOwnerAccess()) {
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
	share := p.GetShare()
	// Check whether resource is shared with the tenant and Action is allowed.
	for _, st := range share {
		tenant := st.GetTenant()
		tAccess := st.GetTenantAccess()
		shareType, uuid := tenantInfo(tenant)

		if (shareType == KindDomain && uuid == rq.domain) ||
			(shareType == KindProject && uuid == rq.project) {
			if permsAccessAllowed(rq, tAccess) {
				return nil
			}
		}
	}
	return errutil.ErrorForbiddenf("object access not allowed. access denied ")
}

func tenantInfo(tenant string) (projectType string, uuid string) {
	uuid = strings.Split(tenant, ":")[1]
	if tenantType(tenant) == KindDomain {
		return KindDomain, uuid
	}
	return KindProject, uuid
}

func tenantType(tenant string) string {
	if strings.HasPrefix(tenant, "domain") {
		return KindDomain
	}
	return KindProject
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
