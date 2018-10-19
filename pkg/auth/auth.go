package auth

import (
	"context"

	"github.com/labstack/echo"

	"github.com/Juniper/contrail/pkg/strutil"
)

//Context is used to represents Context.
// API layer and DB layer depends on this.
type Context struct {
	projectID string
	domainID  string
	userID    string
	roles     []string
}

const (
	//AdminRole string
	AdminRole = "admin"
)

//NewContext makes a authentication context.
func NewContext(domainID, projectID, userID string, roles []string) *Context {
	return &Context{
		projectID: projectID,
		domainID:  domainID,
		userID:    userID,
		roles:     roles,
	}
}

//IsAdmin is used to check if this is admin context
func (context *Context) IsAdmin() bool {
	if context == nil {
		return true
	}
	return strutil.ContainsString(context.roles, AdminRole)
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
		"default-domain", "default-project", "admin", []string{"admin"})
	var authKey interface{} = "auth"
	return context.WithValue(ctx, authKey, Context)
}
