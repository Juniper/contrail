package keystone

import (
	"context"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"

	"github.com/databus23/keystone"
	"github.com/labstack/echo"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"

	apicommon "github.com/Juniper/contrail/pkg/apisrv/common"
	auth2 "github.com/Juniper/contrail/pkg/auth"
	"github.com/Juniper/contrail/pkg/errutil"
)

const (
	keystoneService = "keystone"
)

func authenticate(ctx context.Context, auth *keystone.Auth, tokenString string) (context.Context, error) {
	if tokenString == "" {
		return nil, errors.Wrap(errutil.ErrorUnauthenticated, "no auth token in request")
	}
	validatedToken, err := auth.Validate(tokenString)
	if err != nil {
		log.Errorf("Invalid Token: %s", err)
		return nil, errutil.ErrorUnauthenticated
	}
	roles := []string{}
	for _, r := range validatedToken.Roles {
		roles = append(roles, r.Name)
	}
	project := validatedToken.Project
	if project == nil {
		log.Debug("No project in a token")
		return nil, errutil.ErrorUnauthenticated
	}
	domain := validatedToken.Project.Domain.ID
	user := validatedToken.User
	authContext := auth2.NewContext(domain, project.ID, user.ID, roles)

	var authKey interface{} = "auth"
	newCtx := context.WithValue(ctx, authKey, authContext)
	return newCtx, nil
}

func getKeystoneEndpoint(clusterID string, endpoints *apicommon.EndpointStore) (authEndpoint string, err error) {
	if clusterID != "" {
		scope = "private"
		endpointKey := strings.Join([]string{"/proxy", clusterID, keystoneService, scope})
		keystoneTargets := endpoints.Read(endpointKey)
		if keystoneTargets == nil {
			return "", fmt.Errorf("keystone targets not found for: %s", endpointKey)
		}
		authEndpoint = keystoneTargets.Next(scope)
		if target == "" {
			return "", fmt.Errorf("unable to get keystone endpoint for: %s", endpointKey)
		}
	} else {
		authEndpoint, err = endpoints.GetEndpoint(keystoneService)
		if err != nil {
			return "", fmt.Errorf("unable to get keystone endpoint: %s", err)
		}
	}
	return authEndpoint, err

}

// GetAuthSkipPaths returns the list of paths which need not be authenticated.
func GetAuthSkipPaths() []string {
	skipPaths := []string{
		"/contrail-clusters?fields=uuid,name",
		"/keystone/v3/auth/tokens",
		"/proxy/keystone/v3/auth/tokens",
		"/keystone/v3/auth/projects",
		"/v3/auth/tokens",
	}
	// skip auth for all the static fileutil
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
func AuthMiddleware(keystoneClient *Client, skipPath []string,
	endpoints *apicommon.EndpointStore) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		keystoneClient.AuthURL = keystoneClient.LocalAuthURL
		auth := keystoneClient.NewAuth()
		return func(c echo.Context) error {
			for _, pathQuery := range skipPath {
				switch c.Request().URL.Path {
				case "/":
					return next(c)
				default:
					if strings.Contains(pathQuery, "?") {
						paths := strings.Split(pathQuery, "?")
						if strings.Contains(c.Request().URL.Path, paths[0]) {
							if strings.Compare(c.Request().URL.RawQuery, paths[1]) == 0 {
								return next(c)
							}
						}
					} else if strings.Contains(c.Request().URL.Path, pathQuery) {
						return next(c)
					}
				}
			}
			r := c.Request()
			if r.ProtoMajor == 2 && strings.Contains(r.Header.Get("Content-Type"), "application/grpc") {
				// Skip grpc
				return next(c)
			}
			clusterID := r.Header.Get("X-CLUSTER-ID")
			keystoneEndpoint, err := getKeystoneEndpoint(clusterID, endpoints)
			if err != nil {
				log.Errorf("unable to get keystone endpoint: %s", err)
				return errutil.ToHTTPError(errutil.ErrorUnauthenticated)
			}
			if keystoneEndpoint != "" {
				keystoneClient.SetAuthURL(keystoneEndpoint)
				auth = keystoneClient.NewAuth()
			}
			tokenString := r.Header.Get("X-Auth-Token")
			if tokenString == "" {
				cookie, _ := r.Cookie("x-auth-token") // nolint: errcheck
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
				return errutil.ToHTTPError(err)
			}
			newRequest := r.WithContext(ctx)
			c.SetRequest(newRequest)
			return next(c)
		}
	}
}

//AuthInterceptor for Auth process for gRPC based apps.
func AuthInterceptor(keystoneClient *Client,
	endpoints *apicommon.EndpointStore) grpc.UnaryServerInterceptor {
	keystoneClient.AuthURL = keystoneClient.LocalAuthURL
	auth := keystoneClient.NewAuth()
	return func(ctx context.Context, req interface{},
		info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			return nil, errutil.ErrorUnauthenticated
		}
		token := md["x-auth-token"]
		if len(token) == 0 {
			return nil, errutil.ErrorUnauthenticated
		}
		clusterID := md["X-CLUSTER-ID"]
		keystoneEndpoint, err := getKeystoneEndpoint(clusterID, endpoints)
		if err != nil {
			log.Error(err)
			return nil, errutil.ErrorUnauthenticated
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
