package baseapisrv

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"strings"
	"time"

	"github.com/Juniper/asf/pkg/keystone"
	"github.com/Juniper/asf/pkg/logutil"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"google.golang.org/grpc"

	protocodec "github.com/gogo/protobuf/codec"
)

// Server is an HTTP and GRPC API server.
type Server struct {
	Echo            *echo.Echo
	HomepageHandler *HomepageHandler
	log             *logrus.Entry
}

// APIPlugin registers HTTP endpoints and GRPC services in Server.
type APIPlugin interface {
	RegisterHTTPAPI(HTTPRouter)
	RegisterGRPCAPI(GRPCRouter)
}

// HTTPRouter allows registering HTTP endpoints.
type HTTPRouter interface {
	GET(path string, h HandlerFunc, m ...MiddlewareFunc)
	POST(path string, h HandlerFunc, m ...MiddlewareFunc)
	PUT(path string, h HandlerFunc, m ...MiddlewareFunc)
	DELETE(path string, h HandlerFunc, m ...MiddlewareFunc)
	Use(m ...MiddlewareFunc)
	Group(prefix string, m ...MiddlewareFunc)

	AddNoAuthPaths(paths ...string)

	// TODO Rename to RegisterHomepage
	// TODO Merge into GET, ...
	Register(path string, method string, name string, rel string)
}

// HandlerFunc handles an HTTP request.
// This is a type alias so that users can make one without importing this package if needed.
// TODO(Witaut): Don't require importing echo to make a HandlerFunc.
type HandlerFunc = func(echo.Context) error

// MiddlewareFunc returns a HandlerFunc that processes a request, possibly leaving further processing to next.
// This is a type alias so that users can make one without importing this package if needed.
type MiddlewareFunc = func(next HandlerFunc) HandlerFunc

// GRPCRouter allows registering GRPC services.
type GRPCRouter interface {
	RegisterService(description *grpc.ServiceDesc, service interface{})
}

// NewServer makes a new Server.
func NewServer(plugins []APIPlugin, noAuthPaths []string) (*Server, error) {
	s := &Server{
		Echo:            echo.New(),
		HomepageHandler: NewHomepageHandler(),
	}

	if err := logutil.Configure(viper.GetString("log_level")); err != nil {
		return nil, err
	}
	s.log = logutil.NewLogger("api-server")

	s.setupLoggingMiddleware()

	if viper.GetBool("server.enable_gzip") {
		s.Echo.Use(middleware.Gzip())
	}

	s.Echo.Use(middleware.Recover())
	s.Echo.Binder = &customBinder{}

	readTimeout := viper.GetInt("server.read_timeout")
	writeTimeout := viper.GetInt("server.write_timeout")
	s.Echo.Server.ReadTimeout = time.Duration(readTimeout) * time.Second
	s.Echo.Server.WriteTimeout = time.Duration(writeTimeout) * time.Second

	s.setupCORS()

	staticPath := viper.GetStringMapString("server.static_files")
	for prefix, root := range staticPath {
		s.Echo.Static(prefix, root)
	}

	r := &httpRouter{
		Echo:            s.Echo,
		HomepageHandler: s.HomepageHandler,
	}
	for _, plugin := range plugins {
		plugin.RegisterHTTPAPI(r)
	}
	noAuthPaths = append(noAuthPaths, r.noAuthPaths...)

	httpMiddleware, authGRPCOpts := s.authMiddleware(noAuthPaths)
	r.Use(httpMiddleware...)

	if err := s.setupGRPC(authGRPCOpts, plugins); err != nil {
		return nil, err
	}

	if viper.GetBool("homepage.enabled") {
		s.Echo.GET("/", s.HomepageHandler.Handle)
	}

	if viper.GetBool("recorder.enabled") {
		s.Echo.Use(recorderMiddleware(s.log))
	}

	return s, nil
}

func (s *Server) setupLoggingMiddleware() {
	// TODO: integrate Echo's logger with logrus
	if viper.GetBool("server.log_api") {
		s.Echo.Use(middleware.Logger())
	} else {
		s.Echo.Logger.SetOutput(ioutil.Discard) // Disables Echo's built-in logging.
	}

	if !viper.GetBool("server.log_body") {
		return
	}
	s.Echo.Use(middleware.BodyDump(func(c echo.Context, requestBody, responseBody []byte) {
		if len(responseBody) > 10000 {
			responseBody = responseBody[0:10000] // trim too long entries
		}
		s.log.WithFields(logrus.Fields{
			"request-body":  string(requestBody),
			"response-body": string(responseBody),
		}).Debug("HTTP request handled")
	}))
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

func (s *Server) setupCORS() {
	cors := viper.GetString("server.cors")
	if cors == "" {
		return
	}

	s.log.WithField("cors", cors).Debug("Enabling CORS")
	if cors == "*" {
		s.log.Warn("cors for * have security issue. DO NOT USE THIS IN PRODUCTION")
	}
	s.Echo.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:  []string{cors},
		AllowMethods:  []string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
		AllowHeaders:  []string{"X-Auth-Token", "Content-Type"},
		ExposeHeaders: []string{"X-Total-Count"},
	}))
}

func (s *Server) setupGRPC(grpcOpts []grpc.ServerOption, plugins []APIPlugin) error {
	if !viper.GetBool("server.enable_grpc") {
		return nil
	}
	if !viper.GetBool("server.tls.enabled") {
		return errors.New("GRPC support requires TLS configuration")
	}

	s.log.Debug("Enabling gRPC server")

	opts := []grpc.ServerOption{
		// TODO(Michal): below option potentially breaks compatibility for non golang grpc clients.
		// Ensure it doesn't or find a better solution for un/marshaling `oneof` fields properly.
		grpc.CustomCodec(protocodec.New(0)),
	}
	opts = append(opts, grpcOpts...)
	server := grpc.NewServer(opts...)

	for _, plugin := range plugins {
		plugin.RegisterGRPCAPI(server)
	}

	s.Echo.Use(gRPCMiddleware(server))

	return nil
}

func (s *Server) authMiddleware(noAuthPaths []string) (httpMiddleware []MiddlewareFunc, grpcOpts []grpc.ServerOption) {
	authURL := viper.GetString("keystone.authurl")
	if authURL != "" {
		insecure := viper.GetBool("keystone.insecure")

		m := keystone.NewAuthMiddleware(authURL, insecure, noAuthPaths)
		httpMiddleware = append(httpMiddleware, m.HTTPMiddleware)
		grpcOpts = append(grpcOpts, grpc.UnaryInterceptor(m.GRPCInterceptor))
	} else if viper.GetBool("no_auth") {
		httpMiddleware = append(httpMiddleware, noAuthHTTPMiddleware)
		grpcOpts = append(grpcOpts, grpc.UnaryInterceptor(noAuthGRPCInterceptor))
	}
	return httpMiddleware, grpcOpts
}

type httpRouter struct {
	*echo.Echo
	*HomepageHandler
	noAuthPaths []string
}

// GET registers a GET handler.
func (r *httpRouter) GET(path string, h HandlerFunc, m ...MiddlewareFunc) {
	r.Echo.GET(path, echo.HandlerFunc(h), echoMiddleware(m)...)
}

// POST registers a POST handler.
func (r *httpRouter) POST(path string, h HandlerFunc, m ...MiddlewareFunc) {
	r.Echo.POST(path, echo.HandlerFunc(h), echoMiddleware(m)...)
}

// PUT registers a PUT handler.
func (r *httpRouter) PUT(path string, h HandlerFunc, m ...MiddlewareFunc) {
	r.Echo.PUT(path, echo.HandlerFunc(h), echoMiddleware(m)...)
}

// DELETE registers a DELETE handler.
func (r *httpRouter) DELETE(path string, h HandlerFunc, m ...MiddlewareFunc) {
	r.Echo.DELETE(path, echo.HandlerFunc(h), echoMiddleware(m)...)
}

// Use makes middleware run for all requests.
func (r *httpRouter) Use(m ...MiddlewareFunc) {
	r.Echo.Use(echoMiddleware(m)...)
}

// Group makes the middleware run for all requests under prefix.
func (r *httpRouter) Group(prefix string, m ...MiddlewareFunc) {
	r.Echo.Group(prefix, echoMiddleware(m)...)
}

// AddNoAuthPaths makes requests to paths skip authentication.
func (r *httpRouter) AddNoAuthPaths(paths ...string) {
	r.noAuthPaths = append(r.noAuthPaths, paths...)
}

func echoMiddleware(ms []MiddlewareFunc) []echo.MiddlewareFunc {
	echoMiddleware := make([]echo.MiddlewareFunc, 0, len(ms))
	for _, m := range ms {
		echoMiddleware = append(echoMiddleware, echoMiddlewareFunc(m))
	}
	return echoMiddleware
}

func echoMiddlewareFunc(m MiddlewareFunc) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return echo.HandlerFunc(m(HandlerFunc(next)))
	}
}

// Run starts serving the APIs to clients.
func (s *Server) Run() error {
	if viper.GetBool("server.tls.enabled") {
		return s.Echo.StartTLS(
			viper.GetString("server.address"),
			viper.GetString("server.tls.cert_file"),
			viper.GetString("server.tls.key_file"),
		)
	}

	return s.Echo.Start(viper.GetString("server.address"))
}
