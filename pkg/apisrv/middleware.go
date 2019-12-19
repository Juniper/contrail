package apisrv

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"net/http/httputil"
	"net/url"
	"path/filepath"
	"strings"
	"time"

	"github.com/Juniper/asf/pkg/errutil"
	"github.com/Juniper/contrail/pkg/auth"
	"github.com/databus23/keystone"
	"github.com/labstack/echo"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"

	auth2 "github.com/Juniper/contrail/pkg/auth"
	cleanhttp "github.com/hashicorp/go-cleanhttp"
)

func removePathPrefixMiddleware(prefix string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			req := c.Request()
			req.URL.Path = strings.TrimPrefix(req.URL.Path, prefix)
			return next(c)
		}
	}
}

func proxyMiddleware(target *url.URL, insecure bool) func(next echo.HandlerFunc) echo.HandlerFunc {
	server := httputil.NewSingleHostReverseProxy(target)
	if target.Scheme == "https" {
		server.Transport = &http.Transport{
			Dial: (&net.Dialer{
				Timeout:   30 * time.Second,
				KeepAlive: 30 * time.Second,
			}).Dial,
			TLSClientConfig:     &tls.Config{InsecureSkipVerify: insecure},
			TLSHandshakeTimeout: 10 * time.Second,
		}
	}
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			r := c.Request()
			w := c.Response()
			server.ServeHTTP(w, r)
			return nil
		}
	}
}

func noAuthMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			r := c.Request()
			ctx := auth.NoAuth(r.Context())
			newRequest := r.WithContext(ctx)
			c.SetRequest(newRequest)
			return next(c)
		}
	}
}

func noAuthInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{},
		info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		newCtx := auth.NoAuth(ctx)
		return handler(newCtx, req)
	}
}

func gRPCMiddleware(grpcServer http.Handler) func(next echo.HandlerFunc) echo.HandlerFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			r := c.Request()
			w := c.Response()
			if r.ProtoMajor == 2 && strings.Contains(r.Header.Get("Content-Type"), "application/grpc") {
				grpcServer.ServeHTTP(w, r)
				return nil
			}
			if err := next(c); err != nil {
				c.Error(err)
			}
			return nil
		}
	}
}

type customBinder struct{}

func (*customBinder) Bind(i interface{}, c echo.Context) (err error) {
	rq := c.Request()
	ct := rq.Header.Get(echo.HeaderContentType)
	err = echo.ErrUnsupportedMediaType
	if !strings.HasPrefix(ct, echo.MIMEApplicationJSON) {
		db := new(echo.DefaultBinder)
		return db.Bind(i, c)
	}

	dec := json.NewDecoder(rq.Body)
	dec.UseNumber()
	err = dec.Decode(i)
	if err == io.EOF {
		return nil
	}
	return err
}

func GetAuthSkipPaths() ([]string, error) {
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

func httpTransport(insecure bool) *http.Transport {
	t := cleanhttp.DefaultPooledTransport()
	t.TLSClientConfig = &tls.Config{InsecureSkipVerify: insecure}
	return t
}

func authenticate(ctx context.Context, auth *keystone.Auth, tokenString string) (context.Context, error) {
	if tokenString == "" {
		return nil, errors.Wrap(errutil.ErrorUnauthenticated, "no auth token in request")
	}
	validatedToken, err := auth.Validate(tokenString)
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

	objPerms := auth2.NewObjPerms(validatedToken)
	authContext := auth2.NewContext(domain, project.ID, user.ID, roles, tokenString, objPerms)

	var authKey interface{} = "auth"
	newCtx := context.WithValue(ctx, authKey, authContext)
	return newCtx, nil
}

func newKeystoneAuth(authURL string, insecure bool) *keystone.Auth {
	a := keystone.New(authURL)
	a.Client = &http.Client{
		Transport: httpTransport(insecure),
	}
	return a
}

func AuthMiddleware(authURL string, insecure bool, skipPath []string) echo.MiddlewareFunc {
	auth := newKeystoneAuth(authURL, insecure)

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			for _, pathQuery := range skipPath {
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
			ctx, err := authenticate(r.Context(), auth, tokenString)
			if err != nil {
				logrus.Errorf("Authentication failure: %s", err)
				return errutil.ToHTTPError(err)
			}
			newRequest := r.WithContext(ctx)
			c.SetRequest(newRequest)
			return next(c)
		}
	}
}

func AuthInterceptor(authURL string, insecure bool) grpc.UnaryServerInterceptor {
	auth := newKeystoneAuth(authURL, insecure)

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
		newCtx, err := authenticate(ctx, auth, token[0])
		if err != nil {
			return nil, err
		}
		return handler(newCtx, req)
	}
}
