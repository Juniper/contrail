package keystone

import (
	"context"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"

	"github.com/Juniper/contrail/pkg/common"
	"github.com/databus23/keystone"
	"github.com/labstack/echo"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"

	apicommon "github.com/Juniper/contrail/pkg/apisrv/common"
	log "github.com/sirupsen/logrus"
)

const (
	keystoneService = "keystone"
	private         = apicommon.Private
	pathSep         = "/"
)

func authenticate(ctx context.Context, auth *keystone.Auth, tokenString string) (context.Context, error) {
	if tokenString == "" {
		log.Debug("No auth token in request")
		return nil, common.ErrorUnauthenticated
	}
	validatedToken, err := auth.Validate(tokenString)
	if err != nil {
		log.Errorf("Invalid Token: %s", err)
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

func getKeystoneEndpoint(endpoints *apicommon.EndpointStore) (authEndpoint string, err error) {
	endpointCount := 0
	authEndpoint = ""
	endpoints.Data.Range(func(key, targets interface{}) bool {
		keyString, _ := key.(string)
		keyParts := strings.Split(keyString, pathSep)
		if keyParts[3] != keystoneService || keyParts[4] != private {
			return true // continue iterating the endpoints
		}
		endpointCount++
		if endpointCount > 1 {
			err = fmt.Errorf("Ambiguious, more than one cluster found")
			return false
		}
		authEndpoints, _ := targets.(*apicommon.TargetStore)
		authEndpoint = authEndpoints.Next(private)
		return false

	})

	return authEndpoint, err
}

// GetAuthSkipPaths returns the list of paths which need not be authenticated.
func GetAuthSkipPaths() []string {
	skipPaths := []string{
		"/keystone/v3/auth/tokens",
		"/proxy/keystone/v3/auth/tokens",
		"/keystone/v3/auth/projects",
		"/v3/auth/tokens",
	}
	// skip auth for all the static files
	for prefix, root := range viper.GetStringMap("server.static_files") {
		if prefix == "/" {
			staticFiles, err := ioutil.ReadDir(root.(string))
			if err != nil {
				log.Fatal(err)
			}
			for _, staticFile := range staticFiles {
				skipPaths = append(skipPaths,
					filepath.Join(prefix, staticFile.Name()))
			}
		} else {
			skipPaths = append(skipPaths, prefix)
		}
	}
	return skipPaths
}

//AuthMiddleware is a keystone v3 authentication middleware for REST API.
//nolint: gocyclo
func AuthMiddleware(keystoneClient *KeystoneClient, skipPath []string,
	endpoints *apicommon.EndpointStore) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		keystoneClient.AuthURL = keystoneClient.LocalAuthURL
		auth := keystoneClient.NewAuth()
		return func(c echo.Context) error {
			for _, path := range skipPath {
				switch c.Request().URL.Path {
				case "/":
					return next(c)
				default:
					if strings.Contains(c.Request().URL.Path, path) {
						return next(c)
					}
				}
			}
			r := c.Request()
			if r.ProtoMajor == 2 && strings.Contains(r.Header.Get("Content-Type"), "application/grpc") {
				// Skip grpc
				return next(c)
			}
			keystoneEndpoint, err := getKeystoneEndpoint(endpoints)
			if err != nil {
				log.Errorf("Unable to get keystone endpoint: %s", err)
				return common.ToHTTPError(common.ErrorUnauthenticated)
			}
			if keystoneEndpoint != "" {
				keystoneClient.SetAuthURL(keystoneEndpoint)
				auth = keystoneClient.NewAuth()
			}
			tokenString := r.Header.Get("X-Auth-Token")
			if tokenString == "" {
				cookie, _ := r.Cookie("x-auth-token")
				if cookie != nil {
					tokenString = cookie.Value
				}
				if tokenString == "" {
					tokenString = c.QueryParam("auth_token")
				}
			}
			ctx, err := authenticate(r.Context(), auth, tokenString)
			if err != nil {
				log.Errorf("Authentication failure: %s", err)
				return common.ToHTTPError(err)
			}
			newRequest := r.WithContext(ctx)
			c.SetRequest(newRequest)
			return next(c)
		}
	}
}

//AuthInterceptor for Auth process for gRPC based apps.
func AuthInterceptor(keystoneClient *KeystoneClient,
	endpoints *apicommon.EndpointStore) grpc.UnaryServerInterceptor {
	keystoneClient.AuthURL = keystoneClient.LocalAuthURL
	auth := keystoneClient.NewAuth()
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
		keystoneEndpoint, err := getKeystoneEndpoint(endpoints)
		if err != nil {
			log.Error(err)
			return nil, common.ErrorUnauthenticated
		}
		if keystoneEndpoint != "" {
			keystoneClient.SetAuthURL(keystoneEndpoint)
			auth = keystoneClient.NewAuth()
		}
		newCtx, err := authenticate(ctx, auth, token[0])
		if err != nil {
			return nil, err
		}
		return handler(newCtx, req)
	}
}
