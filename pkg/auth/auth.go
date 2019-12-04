package auth

import (
	"context"

	"github.com/Juniper/asf/pkg/auth"
	"github.com/Juniper/asf/pkg/format"
)

// Context is used to represents Context.
// API layer and DB layer depends on this.
type Context struct {
	projectID   string
	domainID    string
	userID      string
	roles       []string
	authToken   string
	objectPerms ObjectPerms
}

const (
	// AdminRole string
	AdminRole = "admin"
	// GlobalReadOnlyRole string
	GlobalReadOnlyRole = "RO"
	// CloudAdminRole string
	CloudAdminRole = "cloud_admin"
)

// NewContext makes a authentication context.
func NewContext(
	domainID, projectID, userID string, roles []string, authToken string, objectPerms ObjectPerms,
) *Context {
	ctx := &Context{
		projectID:   projectID,
		domainID:    domainID,
		userID:      userID,
		roles:       roles,
		authToken:   authToken,
		objectPerms: objectPerms,
	}
	ctx.substituteObjectPerms()
	return ctx
}

func (context *Context) substituteObjectPerms() {
	context.objectPerms.IsGlobalReadOnlyRole = context.IsGlobalRORole()
	context.objectPerms.IsCloudAdminRole = context.IsCloudAdminRole()
	context.objectPerms.TokenInfo.Token.AuthToken = context.AuthToken()
}

// GetObjPerms returns object perms
func (context *Context) GetObjPerms() interface{} {
	return context.objectPerms
}

// IsAdmin is used to check if this is admin context
func (context *Context) IsAdmin() bool {
	if context == nil {
		return true
	}
	return format.ContainsString(context.roles, AdminRole)
}

// IsGlobalRORole is used to check if this context is  global read only role
func (context *Context) IsGlobalRORole() bool {
	if context == nil {
		return true
	}
	return format.ContainsString(context.roles, GlobalReadOnlyRole)
}

// IsCloudAdminRole is used to check if this context is cloud admin role
func (context *Context) IsCloudAdminRole() bool {
	if context == nil {
		return true
	}
	return format.ContainsString(context.roles, CloudAdminRole)
}

// ProjectID is used to get an id for project.
func (context *Context) ProjectID() string {
	if context == nil {
		return "admin"
	}
	return context.projectID
}

// DomainID is used to get an id for domain.
func (context *Context) DomainID() string {
	if context == nil {
		return "admin"
	}
	return context.domainID
}

// UserID is used to get an id for User.
func (context *Context) UserID() string {
	if context == nil {
		return AdminRole
	}
	return context.userID
}

// AuthToken is used to get an auth token of request.
func (context *Context) AuthToken() string {
	if context == nil {
		return ""
	}
	return context.authToken
}

// Roles  is used to get the roles of a user
func (context *Context) Roles() []string {
	if context == nil {
		return nil
	}
	return context.roles
}

// NoAuth is used to create new no auth context
func NoAuth(ctx context.Context) context.Context {
	Context := NewContext(
		"default-domain",
		"default-project",
		AdminRole,
		[]string{AdminRole},
		"",
		NewObjPerms(nil),
	)
	return auth.WithIdentity(ctx, Context)
}
