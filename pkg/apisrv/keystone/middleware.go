package keystone

import (
	"context"
	"crypto/tls"
	"net"
	"net/http"
	"strings"
	"time"

	"github.com/Juniper/contrail/pkg/common"
	"github.com/databus23/keystone"
	"github.com/labstack/echo"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func newAuth(authURL string, insecure bool) *keystone.Auth {
	auth := keystone.New(authURL)
	tr := &http.Transport{
		Dial: (&net.Dialer{
			Timeout: 5 * time.Second,
		}).Dial,
		TLSHandshakeTimeout: 5 * time.Second,
		TLSClientConfig:     &tls.Config{InsecureSkipVerify: insecure}, // nolint: gas
	}
	client := &http.Client{
		Transport: tr,
		Timeout:   time.Second * 10,
	}
	auth.Client = client
	return auth
}

func authenticate(ctx context.Context, auth *keystone.Auth, tokenString string) (context.Context, error) {
	if tokenString == "" {
		return nil, common.ErrorUnauthenticated
	}
	validatedToken, err := auth.Validate(tokenString)
	if err != nil {
		return nil, common.ErrorUnauthenticated
	}
	log.WithField("token", validatedToken).Debug("Authenticated")
	roles := []string{}
	for _, r := range validatedToken.Roles {
		roles = append(roles, r.Name)
	}
	project := validatedToken.Project
	if project == nil {
		log.Debug("No project in a token")
		return nil, common.ErrorUnauthenticated
	}
	domain := validatedToken.Project.Domain.ID
	user := validatedToken.User
	authContext := common.NewAuthContext(domain, project.ID, user.ID, roles)

	var authKey interface{} = "auth"
	newCtx := context.WithValue(ctx, authKey, authContext)
	return newCtx, nil
}

//AuthMiddleware is a keystone v3 authentication middleware for REST API.
func AuthMiddleware(authURL string, insecure bool, skipPath []string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		auth := newAuth(authURL, insecure)
		return func(c echo.Context) error {
			for _, path := range skipPath {
				if strings.HasPrefix(c.Request().URL.Path, path) {
					return next(c)
				}
			}
			r := c.Request()
			if r.ProtoMajor == 2 && strings.Contains(r.Header.Get("Content-Type"), "application/grpc") {
				// Skip grpc
				return next(c)
			}
			tokenString := r.Header.Get("X-Auth-Token")
			ctx, err := authenticate(r.Context(), auth, tokenString)
			if err != nil {
				return common.ToHTTPError(err)
			}
			newRequest := r.WithContext(ctx)
			c.SetRequest(newRequest)
			return next(c)
		}
	}
}

//AuthInterceptor for Auth process for gRPC based apps.
func AuthInterceptor(authURL string, insecure bool) grpc.UnaryServerInterceptor {
	auth := newAuth(authURL, insecure)
	return func(ctx context.Context, req interface{},
		info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			return nil, common.ErrorUnauthenticated
		}
		token := md["x-auth-token"]
		if len(token) == 0 {
			return nil, common.ErrorUnauthenticated
		}
		newCtx, err := authenticate(ctx, auth, token[0])
		if err != nil {
			return nil, err
		}
		return handler(newCtx, req)
	}
}
