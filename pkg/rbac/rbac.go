package rbac

import (
	"context"
	"strings"

	"github.com/Juniper/asf/pkg/auth"
	"github.com/Juniper/asf/pkg/errutil"
	"github.com/Juniper/asf/pkg/models/basemodels"
	"github.com/Juniper/contrail/pkg/models"
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
func CheckPermissions(ctx context.Context, l []*models.APIAccessList, aaaMode string, kind string, op Action) error {
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
func CheckObjectPermissions(ctx context.Context, p *models.PermType2, aaaMode string, kind string, op Action) error {
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
	return errutil.ErrorForbiddenf("object access not allowed. access denied ")
}

func tenantInfo(tenant string) (projectType string, uuid string) {
	uuid = strings.Split(tenant, ":")[1]
	if tenantType(tenant) == models.KindDomain {
		return models.KindDomain, uuid
	}
	return models.KindProject, uuid
}

func tenantType(tenant string) string {
	if strings.HasPrefix(tenant, "domain") {
		return models.KindDomain
	}
	return models.KindProject
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
