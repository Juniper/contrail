package auth

import (
	"context"
	"time"

	"github.com/databus23/keystone"
	"github.com/labstack/echo"

	"github.com/Juniper/contrail/pkg/format"
)

// Context is used to represents Context.
// API layer and DB layer depends on this.
type Context struct {
	projectID   string
	domainID    string
	userID      string
	roles       []string
	authToken   string
	objectPerms ObjectPerms // *services.ObjectPerms
}

// identification is struct that describe the identity of resource
type identification struct {
	ID   string `yaml:"id" json:"id"`
	Name string `yaml:"name" json:"name"`
}

// token is used in ObjectPerms to store token related information.
type token struct {
	IsDomain  bool             `json:"is_domain"`
	AuthToken string           `json:"auth_token"`
	ExpiresAt string           `json:"expires_at"`
	IssuedAt  string           `json:"issued_at"`
	Version   string           `json:"version"`
	Roles     []identification `json:"roles"`
	Project   struct {
		identification
		Domain identification `json:"domain"`
	} `json:"project"`
	User struct {
		identification
		Domain identification `json:"domain"`
	} `json:"user"`
}

// ObjectPerms holds information get from Keystone module.
type ObjectPerms struct {
	IsGlobalReadOnlyRole bool `json:"is_global_read_only_role"`
	IsCloudAdminRole     bool `json:"is_cloud_admin_role"`
	TokenInfo            struct {
		Token token `json:"token"`
	} `json:"token_info"`
}

// NewObjPerms inits ObjectPerms structure using keystone token
func NewObjPerms(kt *keystone.Token) ObjectPerms {
	if kt == nil {
		return ObjectPerms{}
	}

	// nolint: lll
	// TODO: implement rest of logic: https://github.com/Juniper/contrail-controller/blob/691559e3cbfa9d9db227b4ee55f7eced141c4498/src/config/api-server/vnc_cfg_api_server/vnc_cfg_api_server.py#L2332
	objPerms := ObjectPerms{
		//  part of parameters are set while creating Context in NewContext() method
		TokenInfo: struct {
			Token token `json:"token"`
		}{
			Token: token{
				ExpiresAt: kt.ExpiresAt.Format(time.RFC3339),
				IssuedAt:  kt.IssuedAt.Format(time.RFC3339),
				Version:   "", // TODO(pawel.drapiewski): find the way to get this information if needed
				Roles:     tokenRolesToObjectPermsRoles(kt.Roles),
				Project: struct {
					identification
					Domain identification `json:"domain"`
				}{
					identification: identification{
						ID:   kt.Project.ID,
						Name: kt.Project.Name,
					},
					Domain: identification{
						ID:   kt.Project.Domain.ID,
						Name: kt.Project.Domain.Name,
					},
				},
				User: struct {
					identification
					Domain identification `json:"domain"`
				}{
					identification: identification{
						ID:   kt.User.ID,
						Name: kt.User.Name,
					},
					Domain: identification{
						ID:   kt.User.Domain.ID,
						Name: kt.User.Domain.Name,
					},
				},
			},
		},
	}

	if kt.Domain != nil {
		objPerms.TokenInfo.Token.IsDomain = kt.Domain.Enabled
	}
	return objPerms
}

func tokenRolesToObjectPermsRoles(tokenRoles []struct {
	ID   string
	Name string
}) []identification {
	var identifications []identification
	for _, role := range tokenRoles {
		identifications = append(identifications, identification{ID: role.ID, Name: role.Name})
		_ = role
	}
	return identifications
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
func (context *Context) GetObjPerms() ObjectPerms {
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

// GetContext is used to get an authentication from echo.Context.
func GetContext(c echo.Context) *Context {
	ctx := c.Request().Context()
	return GetAuthCTX(ctx)
}

// GetAuthCTX is used to get an authentication from ctx.Context.
func GetAuthCTX(ctx context.Context) *Context {
	iAuth := ctx.Value("auth")
	auth, _ := iAuth.(*Context) // nolint: errcheck
	return auth
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
	var authKey interface{} = "auth"
	return context.WithValue(ctx, authKey, Context)
}
