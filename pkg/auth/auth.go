package auth

import (
	"context"

	"github.com/labstack/echo"

	"github.com/Juniper/contrail/pkg/format"
)

//Context is used to represents Context.
// API layer and DB layer depends on this.
type Context struct {
	projectID string
	domainID  string
	userID    string
	roles     []string
	authToken string
}

const (
	//AdminRole string
	AdminRole = "admin"
	//GlobalReadOnlyRole string
	GlobalReadOnlyRole = "RO"
)

//NewContext makes a authentication context.
func NewContext(domainID, projectID, userID string, roles []string, authToken string) *Context {
	return &Context{
		projectID: projectID,
		domainID:  domainID,
		userID:    userID,
		roles:     roles,
		authToken: authToken,
	}
}

//IsAdmin is used to check if this is admin context
func (context *Context) IsAdmin() bool {
	if context == nil {
		return true
	}
	return format.ContainsString(context.roles, AdminRole)
}

//IsGlobalRORole is used to check if this context is  global read only role
func (context *Context) IsGlobalRORole() bool {
	if context == nil {
		return true
	}
	return format.ContainsString(context.roles, GlobalReadOnlyRole)
}

//ProjectID is used to get an id for project.
func (context *Context) ProjectID() string {
	if context == nil {
		return "admin"
	}
	return context.projectID
}

//DomainID is used to get an id for domain.
func (context *Context) DomainID() string {
	if context == nil {
		return "admin"
	}
	return context.domainID
}

//UserID is used to get an id for User.
func (context *Context) UserID() string {
	if context == nil {
		return AdminRole
	}
	return context.userID
}

//AuthToken is used to get an auth token of request.
func (context *Context) AuthToken() string {
	if context == nil {
		return ""
	}
	return context.authToken
}

//Roles  is used to get the roles of a user
func (context *Context) Roles() []string {
	if context == nil {
		return nil
	}
	return context.roles
}

//GetContext is used to get an authentication from echo.Context.
func GetContext(c echo.Context) *Context {
	ctx := c.Request().Context()
	return GetAuthCTX(ctx)
}

//GetAuthCTX is used to get an authentication from ctx.Context.
func GetAuthCTX(ctx context.Context) *Context {
	iAuth := ctx.Value("auth")
	auth, _ := iAuth.(*Context) //nolint: errcheck
	return auth
}

// NoAuth is used to create new no auth context
func NoAuth(ctx context.Context) context.Context {
	Context := NewContext(
		"default-domain", "default-project", "admin", []string{"admin"}, "")
	var authKey interface{} = "auth"
	return context.WithValue(ctx, authKey, Context)
}
