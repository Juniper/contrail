package keystone

import (
	"context"
	"strings"

	"github.com/Juniper/contrail/pkg/common"
	"github.com/databus23/keystone"
	"github.com/labstack/echo"
	log "github.com/sirupsen/logrus"
)

//AuthMiddleware is a keystone v3 authentication middleware.
func AuthMiddleware(authURL string, skipPath []string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		auth := keystone.New(authURL)
		return func(c echo.Context) error {
			for _, path := range skipPath {
				if strings.HasPrefix(c.Request().URL.Path, path) {
					return next(c)
				}
			}
			tokenString := c.Request().Header.Get("X-Auth-Token")
			if tokenString == "" {
				return echo.ErrUnauthorized
			}
			validatedToken, err := auth.Validate(tokenString)
			if err != nil {
				log.Debug(err)
				return echo.ErrUnauthorized
			}
			log.WithField("token", validatedToken).Debug("Authenticated")
			roles := []string{}
			for _, r := range validatedToken.Roles {
				roles = append(roles, r.Name)
			}
			project := validatedToken.Project
			if project == nil {
				log.Debug("No project in a token")
				return echo.ErrUnauthorized
			}
			domain := validatedToken.Domain
			if domain == nil {
				log.Debug("No domain in a token")
				return echo.ErrUnauthorized
			}
			user := validatedToken.User
			auth := common.NewAuthContext(domain.ID, project.ID, user.ID, roles)
			request := c.Request()
			var authKey interface{}
			authKey = "auth"
			ctx := context.WithValue(request.Context(), authKey, auth)
			newRequest := request.WithContext(ctx)
			c.SetRequest(newRequest)
			return next(c)
		}
	}
}
