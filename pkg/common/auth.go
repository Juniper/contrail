package common

import "github.com/labstack/echo"

//AuthContext is used to represents AuthContext.
// API layer and DB layer depends on this.
type AuthContext struct {
	projectID string
	domainID  string
	userID    string
	roles     []string
}

const (
	//AdminRole string
	AdminRole = "admin"
)

//NewAuthContext makes a authentication context.
func NewAuthContext(domainID, projectID, userID string, roles []string) *AuthContext {
	return &AuthContext{
		projectID: projectID,
		domainID:  domainID,
		userID:    userID,
		roles:     roles,
	}
}

//IsAdmin is used to check if this is admin context
func (context *AuthContext) IsAdmin() bool {
	if context == nil {
		return true
	}
	return ContainsString(context.roles, AdminRole)
}

//ProjectID is used to get an id for project.
func (context *AuthContext) ProjectID() string {
	if context == nil {
		return "admin"
	}
	return context.projectID
}

//DomainID is used to get an id for domain.
func (context *AuthContext) DomainID() string {
	if context == nil {
		return "admin"
	}
	return context.domainID
}

//GetAuthContext is used to get an authentication from echo.Context.
func GetAuthContext(c echo.Context) *AuthContext {
	iAuth := c.Get("auth")
	auth, _ := iAuth.(*AuthContext)
	return auth
}
