package keystone

import (
	"context"
	"crypto/tls"
	"io/ioutil"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/Juniper/asf/pkg/errutil"
	"github.com/Juniper/contrail/pkg/apisrv/baseapisrv"
	"github.com/Juniper/contrail/pkg/auth"
	"github.com/databus23/keystone"
	"github.com/labstack/echo"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"

	asfauth "github.com/Juniper/asf/pkg/auth"
	cleanhttp "github.com/hashicorp/go-cleanhttp"
)

// AuthPlugin authenticates requests to the server with Keystone.
type AuthPlugin struct {
	auth      *keystone.Auth
	skipPaths []string
}

// NewAuthPluginByViper creates an AuthPlugin based on global Viper configuration.
func NewAuthPluginByViper() (*AuthPlugin, error) {
	skipPaths, err := getAuthSkipPaths()
	if err != nil {
		return nil, errors.Wrap(err, "failed to setup paths skipped from authentication")
	}

	authURL := viper.GetString("keystone.authurl")
	insecure := viper.GetBool("keystone.insecure")

	return &AuthPlugin{
		auth:      newKeystoneAuth(authURL, insecure),
		skipPaths: skipPaths,
	}, nil
}

// getAuthSkipPaths returns the list of paths which need not be authenticated.
func getAuthSkipPaths() ([]string, error) {
	skipPaths := []string{
		"/contrail-clusters?fields=uuid,name",
		"/keystone/v3/auth/tokens",
		// TODO(dfurman): "server.dynamic_proxy_path" or DefaultDynamicProxyPath should be used
		"/proxy/",
		"/keystone/v3/projects",
		"/keystone/v3/auth/projects", // TODO: Remove this, since "/keystone/v3/projects" is a keystone endpoint
		"/v3/auth/tokens",
	}
	// skip auth for all the static fileutil
	for prefix, root := range viper.GetStringMap("server.static_files") {
		if prefix == "/" {
			staticFiles, err := ioutil.ReadDir(root.(string))
			if err != nil {
				return nil, errors.WithStack(err)
			}
			for _, staticFile := range staticFiles {
				skipPaths = append(skipPaths,
					filepath.Join(prefix, staticFile.Name()))
			}
		} else {
			skipPaths = append(skipPaths, prefix)
		}
	}
	return skipPaths, nil
}

// RegisterHTTPAPI registers authentication middleware for most endpoints in the server.
func (p *AuthPlugin) RegisterHTTPAPI(r baseapisrv.HTTPRouter) {
	r.Use(p.middleware)
}

//AuthMiddleware is a keystone v3 authentication middleware for REST API.
//nolint: gocyclo
func (p *AuthPlugin) middleware(next baseapisrv.HandlerFunc) baseapisrv.HandlerFunc {
	return func(c echo.Context) error {
		for _, pathQuery := range p.skipPaths {
			switch c.Request().URL.Path {
			case "/":
				return next(c)
			default:
				if strings.Contains(pathQuery, "?") {
					paths := strings.Split(pathQuery, "?")
					if strings.Contains(c.Request().URL.Path, paths[0]) &&
						strings.Compare(c.Request().URL.RawQuery, paths[1]) == 0 {
						return next(c)
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
		ctx, err := authenticate(r.Context(), p.auth, tokenString)
		if err != nil {
			logrus.Errorf("Authentication failure: %s", err)
			return errutil.ToHTTPError(err)
		}
		newRequest := r.WithContext(ctx)
		c.SetRequest(newRequest)
		return next(c)
	}
}

// RegisterGRPCAPI registers an authentication interceptor for GRPC.
func (p *AuthPlugin) RegisterGRPCAPI(r baseapisrv.GRPCRouter) {
	r.AddServerOptions(grpc.UnaryInterceptor(p.interceptor))
}

// interceptor for Auth process for gRPC based apps.
func (p *AuthPlugin) interceptor(
	ctx context.Context, req interface{},
	info *grpc.UnaryServerInfo, handler grpc.UnaryHandler,
) (interface{}, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, errutil.ErrorUnauthenticated
	}
	token := md["x-auth-token"]
	if len(token) == 0 {
		return nil, errutil.ErrorUnauthenticated
	}
	newCtx, err := authenticate(ctx, p.auth, token[0])
	if err != nil {
		return nil, err
	}
	return handler(newCtx, req)
}

func newKeystoneAuth(authURL string, insecure bool) *keystone.Auth {
	a := keystone.New(authURL)
	a.Client = &http.Client{
		Transport: httpTransport(insecure),
	}
	return a
}

func httpTransport(insecure bool) *http.Transport {
	t := cleanhttp.DefaultPooledTransport()
	t.TLSClientConfig = &tls.Config{InsecureSkipVerify: insecure}
	return t
}

func authenticate(ctx context.Context, ka *keystone.Auth, tokenString string) (context.Context, error) {
	if tokenString == "" {
		return nil, errors.Wrap(errutil.ErrorUnauthenticated, "no auth token in request")
	}
	validatedToken, err := ka.Validate(tokenString)
	if err != nil {
		logrus.Errorf("Invalid Token: %s", err)
		return nil, errutil.ErrorUnauthenticated
	}
	roles := []string{}
	for _, r := range validatedToken.Roles {
		roles = append(roles, r.Name)
	}
	project := validatedToken.Project
	if project == nil {
		logrus.Debug("No project in a token")
		return nil, errutil.ErrorUnauthenticated
	}
	domain := validatedToken.Project.Domain.ID
	user := validatedToken.User

	objPerms := auth.NewObjPerms(validatedToken)
	authContext := auth.NewContext(domain, project.ID, user.ID, roles, tokenString, objPerms)

	return asfauth.WithIdentity(ctx, authContext), nil
}
