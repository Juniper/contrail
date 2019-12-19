package apisrv

import (
	"crypto/tls"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"sync"
	"time"

	"github.com/Juniper/asf/pkg/client"
	"github.com/Juniper/asf/pkg/db/basedb"
	"github.com/Juniper/asf/pkg/fileutil"
	"github.com/Juniper/asf/pkg/logutil"
	"github.com/Juniper/contrail/pkg/collector"
	"github.com/Juniper/contrail/pkg/collector/analytics"
	"github.com/Juniper/contrail/pkg/constants"
	"github.com/Juniper/contrail/pkg/db"
	"github.com/Juniper/contrail/pkg/db/cache"
	"github.com/Juniper/contrail/pkg/endpoint"
	"github.com/Juniper/contrail/pkg/keystone"
	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/neutron"
	"github.com/Juniper/contrail/pkg/replication"
	"github.com/Juniper/contrail/pkg/services"
	"github.com/Juniper/contrail/pkg/types"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"google.golang.org/grpc"

	asfclient "github.com/Juniper/asf/pkg/client"
	kstypes "github.com/Juniper/asf/pkg/keystone"
	etcdclient "github.com/Juniper/contrail/pkg/db/etcd"
	protocodec "github.com/gogo/protobuf/codec"
)

// Server HTTP paths.
const (
	FQNameToIDPath  = "fqname-to-id"
	IDToFQNamePath  = "id-to-fqname"
	UserAgentKVPath = "useragent-kv"
	WatchPath       = "watch"
)

type KeystoneController interface {
	GetAuthType(clusterID string) (string, error)
	Init(e *echo.Echo, es *endpoint.Store) error
}

// Server represents Intent API Server.
type Server struct {
	Echo                       *echo.Echo
	GRPCServer                 *grpc.Server
	Keystone                   KeystoneController
	DBService                  *db.Service
	Proxy                      *proxyService
	Service                    services.Service
	IPAMServer                 services.IPAMServer
	ChownServer                services.ChownServer
	SetTagServer               services.SetTagServer
	RefRelaxServer             services.RefRelaxServer
	UserAgentKVServer          services.UserAgentKVServer
	FQNameToIDServer           services.FQNameToIDServer
	IDToFQNameServer           services.IDToFQNameServer
	PropCollectionUpdateServer services.PropCollectionUpdateServer
	Cache                      *cache.DB
	Collector                  collector.Collector
	VNCReplicator              *replication.Replicator
	log                        *logrus.Entry
}

// NewServer makes a server.
func NewServer() (*Server, error) {
	server := &Server{
		Echo: echo.New(),
	}
	return server, nil
}

// Init setups the Server.
// nolint: gocyclo
func (s *Server) Init() (err error) {
	if err = logutil.Configure(viper.GetString("log_level")); err != nil {
		return err
	}
	s.log = logutil.NewLogger("api-server")

	// TODO: integrate Echo's logger with logrus
	if viper.GetBool("server.log_api") {
		s.Echo.Use(middleware.Logger())
	} else {
		s.Echo.Logger.SetOutput(ioutil.Discard) // Disables Echo's built-in logging.
	}

	if viper.GetBool("server.log_body") {
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

	if viper.GetBool("server.enable_gzip") {
		s.Echo.Use(middleware.Gzip())
	}

	s.Echo.Use(middleware.Recover())
	s.Echo.Binder = &customBinder{}

	sqlDB, err := basedb.ConnectDB(analytics.WithCommitLatencyReporting(s.Collector))
	if err != nil {
		return err
	}
	s.DBService = db.NewService(sqlDB)

	if err = s.setupCollector(); err != nil {
		return err
	}

	cs, err := s.setupService()
	if err != nil {
		return err
	}

	s.Service = cs
	s.IPAMServer = cs
	s.ChownServer = cs
	s.SetTagServer = cs
	s.RefRelaxServer = cs
	s.UserAgentKVServer = cs
	s.FQNameToIDServer = cs
	s.IDToFQNameServer = cs
	s.PropCollectionUpdateServer = cs

	if viper.GetBool("server.enable_vnc_neutron") {
		s.setupNeutronService(cs)
	}

	readTimeout := viper.GetInt("server.read_timeout")
	writeTimeout := viper.GetInt("server.write_timeout")
	s.Echo.Server.ReadTimeout = time.Duration(readTimeout) * time.Second
	s.Echo.Server.WriteTimeout = time.Duration(writeTimeout) * time.Second

	cors := viper.GetString("server.cors")

	if cors != "" {
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

	staticPath := viper.GetStringMapString("server.static_files")
	for prefix, root := range staticPath {
		s.Echo.Static(prefix, root)
	}

	if err = s.registerStaticProxyEndpoints(); err != nil {
		return errors.Wrap(err, "failed to register static proxy endpoints")
	}

	endpointStore := endpoint.NewStore()
	s.serveDynamicProxy(endpointStore)

	keystoneAuthURL, keystoneInsecure := viper.GetString("keystone.authurl"), viper.GetBool("keystone.insecure")
	if keystoneAuthURL != "" {
		var skipPaths []string
		skipPaths, err = GetAuthSkipPaths()
		if err != nil {
			return errors.Wrap(err, "failed to setup paths skipped from authentication")
		}
		s.Echo.Use(keystone.AuthMiddleware(keystoneAuthURL, keystoneInsecure, skipPaths))
	} else if viper.GetBool("no_auth") {
		s.Echo.Use(noAuthMiddleware())
	}

	if viper.GetBool("keystone.local") {
		var k KeystoneController
		err = k.Init(s.Echo, endpointStore)
		if err != nil {
			return errors.Wrap(err, "Failed to init local keystone server")
		}
		s.Keystone = k
	}

	if viper.GetBool("server.enable_vnc_replication") {
		if err = s.startVNCReplicator(endpointStore, s.Keystone); err != nil {
			return err
		}
	}

	if viper.GetBool("server.enable_grpc") {
		if !viper.GetBool("server.tls.enabled") {
			return errors.New("GRPC support requires TLS configuration")
		}
		s.log.Debug("Enabling gRPC server")
		opts := []grpc.ServerOption{
			// TODO(Michal): below option potentially breaks compatibility for non golang grpc clients.
			// Ensure it doesn't or find a better solution for un/marshaling `oneof` fields properly.
			grpc.CustomCodec(protocodec.New(0)),
		}
		if keystoneAuthURL != "" {
			opts = append(opts, grpc.UnaryInterceptor(keystone.AuthInterceptor(
				keystoneAuthURL,
				keystoneInsecure,
			)))
		} else if viper.GetBool("no_auth") {
			opts = append(opts, grpc.UnaryInterceptor(noAuthInterceptor()))
		}
		s.GRPCServer = grpc.NewServer(opts...)
		services.RegisterContrailServiceServer(s.GRPCServer, s.Service)
		services.RegisterIPAMServer(s.GRPCServer, s.IPAMServer)
		services.RegisterChownServer(s.GRPCServer, s.ChownServer)
		services.RegisterSetTagServer(s.GRPCServer, s.SetTagServer)
		services.RegisterRefRelaxServer(s.GRPCServer, s.RefRelaxServer)
		services.RegisterFQNameToIDServer(s.GRPCServer, s.FQNameToIDServer)
		services.RegisterIDToFQNameServer(s.GRPCServer, s.IDToFQNameServer)
		services.RegisterUserAgentKVServer(s.GRPCServer, s.UserAgentKVServer)
		services.RegisterPropCollectionUpdateServer(s.GRPCServer, s.PropCollectionUpdateServer)
		s.Echo.Use(gRPCMiddleware(s.GRPCServer))
	}

	if viper.GetBool("homepage.enabled") {
		s.setupHomepage()
	}

	s.setupWatchAPI()
	s.setupActionResources(cs)

	if viper.GetBool("recorder.enabled") {
		file := viper.GetString("recorder.file")
		scenario := &struct {
			Workflow []*recorderTask `yaml:"workflow,omitempty"`
		}{}
		var mutex sync.Mutex
		s.Echo.Use(middleware.BodyDump(func(c echo.Context, requestBody, responseBody []byte) {
			var data interface{}
			err := json.Unmarshal(requestBody, &data)
			if err != nil {
				s.log.WithError(err).Error("Malformed JSON input")
			}
			var expected interface{}
			err = json.Unmarshal(responseBody, &expected)
			if err != nil {
				s.log.WithError(err).Error("Malformed JSON response")
			}
			task := &recorderTask{
				Request: &client.Request{
					Method:   c.Request().Method,
					Path:     c.Request().URL.Path,
					Expected: []int{c.Response().Status},
					Data:     data,
				},
				Expect: expected,
			}
			mutex.Lock()
			defer mutex.Unlock()
			scenario.Workflow = append(scenario.Workflow, task)
			err = fileutil.SaveFile(file, scenario)
			if err != nil {
				s.log.WithError(err).Error("Failed to save scenario to file")
			}
		}))
	}

	return nil
}

func (s *Server) setupCollector() error {
	cfg := &analytics.Config{}
	var err error
	if err = viper.UnmarshalKey("collector", cfg); err != nil {
		return errors.Wrap(err, "failed to unmarshal collector config")
	}
	if s.Collector, err = analytics.NewCollector(cfg); err != nil {
		return errors.Wrap(err, "failed to create collector")
	}
	analytics.AddLoggerHook(s.Collector)
	s.Echo.Use(middleware.BodyDump(func(
		ctx echo.Context, reqBody, resBody []byte,
	) {
		s.Collector.Send(analytics.RESTAPITrace(ctx, reqBody, resBody))
	}))
	return nil
}

func (s *Server) setupService() (*services.ContrailService, error) {
	var serviceChain []services.Service

	cs, err := s.contrailService()
	if err != nil {
		return nil, err
	}
	serviceChain = append(serviceChain, cs)

	serviceChain = append(serviceChain, &services.RefUpdateToUpdateService{
		ReadService:       s.DBService,
		InTransactionDoer: s.DBService,
	})

	serviceChain = append(serviceChain, &services.SanitizerService{
		MetadataGetter: s.DBService,
	})

	serviceChain = append(serviceChain, &services.RBACService{
		ReadService: s.DBService,
		AAAMode:     viper.GetString("aaa_mode")})

	if viper.GetBool("server.enable_vnc_neutron") {
		serviceChain = append(serviceChain, &neutron.Service{
			Keystone: &kstypes.Client{
				URL: viper.GetString("keystone.authurl"),
				HTTPDoer: analytics.LatencyReportingDoer{
					Doer: &http.Client{
						Transport: &http.Transport{
							TLSClientConfig: &tls.Config{InsecureSkipVerify: viper.GetBool("keystone.insecure")},
						},
					},
					Collector:   s.Collector,
					Operation:   "VALIDATE",
					Application: "KEYSTONE",
				},
			},
			ReadService:    s.DBService,
			MetadataGetter: s.DBService,
			WriteService: &services.InternalContextWriteServiceWrapper{
				WriteService: serviceChain[0],
			},
			InTransactionDoer: s.DBService,
		})
	}

	serviceChain = append(serviceChain, &types.ContrailTypeLogicService{
		ReadService:       s.DBService,
		InTransactionDoer: s.DBService,
		AddressManager:    s.DBService,
		IntPoolAllocator:  s.DBService,
		MetadataGetter:    s.DBService,
		WriteService: &services.InternalContextWriteServiceWrapper{
			WriteService: serviceChain[0],
		},
	})

	serviceChain = append(serviceChain, services.NewQuotaCheckerService(s.DBService))

	if viper.GetBool("server.notify_etcd") {
		en := s.etcdNotifier()
		if en != nil {
			serviceChain = append(serviceChain, en)
		}
	}

	serviceChain = append(serviceChain, s.DBService)

	services.Chain(serviceChain...)
	return cs, nil
}

func (s *Server) contrailService() (*services.ContrailService, error) {
	tv, err := models.NewTypeValidatorWithFormat()
	if err != nil {
		return nil, err
	}

	cs := &services.ContrailService{
		BaseService:        services.BaseService{},
		DBService:          s.DBService,
		TypeValidator:      tv,
		MetadataGetter:     s.DBService,
		InTransactionDoer:  s.DBService,
		IntPoolAllocator:   s.DBService,
		RefRelaxer:         s.DBService,
		UserAgentKVService: s.DBService,
		Collector:          s.Collector,
	}

	cs.RegisterRESTAPI(s.Echo)
	return cs, nil
}

func (s *Server) setupNeutronService(cs services.Service) *neutron.Server {
	n := &neutron.Server{
		ReadService:       cs,
		WriteService:      cs,
		UserAgentKV:       s.UserAgentKVServer,
		IDToFQNameService: s.IDToFQNameServer,
		FQNameToIDService: s.FQNameToIDServer,
		InTransactionDoer: s.DBService,
		Log:               logutil.NewLogger("neutron-server"),
	}
	n.RegisterNeutronAPI(s.Echo)
	return n
}

func (s *Server) etcdNotifier() services.Service {
	// TODO(Michał): Make the codec configurable
	en, err := etcdclient.NewNotifierService(viper.GetString(constants.ETCDPathVK), models.JSONCodec)
	if err != nil {
		s.log.WithError(err).Error("Failed to add etcd Notifier Service - ignoring")
		return nil
	}
	return en
}

func (s *Server) registerStaticProxyEndpoints() error {
	for prefix, targetURLs := range viper.GetStringMapStringSlice("server.proxy") {
		if len(targetURLs) == 0 {
			return errors.Errorf("no target URLs provided for prefix %v", prefix)
		}

		// TODO(dfurman): proxy requests to all provided target URLs
		t, err := url.Parse(targetURLs[0])
		if err != nil {
			return errors.Wrapf(err, "bad proxy target URL: %s", targetURLs[0])
		}

		g := s.Echo.Group(prefix)
		g.Use(removePathPrefixMiddleware(prefix))
		g.Use(proxyMiddleware(t, viper.GetBool("server.proxy.insecure")))
	}
	return nil
}

func (s *Server) serveDynamicProxy(es *endpoint.Store) {
	config := loadDynamicProxyConfig()
	s.Echo.Group(config.Path, dynamicProxyMiddleware(es, config))

	s.Proxy = newProxyService(es, s.DBService, config)
	s.Proxy.StartEndpointsSync()
}

func loadDynamicProxyConfig() *DynamicProxyConfig {
	path := viper.GetString("server.dynamic_proxy_path")
	if path == "" {
		path = DefaultDynamicProxyPath
	}

	return &DynamicProxyConfig{
		Path:                         path,
		ServiceTokenEndpointPrefixes: viper.GetStringSlice("server.service_token_endpoint_prefixes"),
		ServiceUserClientConfig:      loadServiceUserClientConfig(),
	}
}

func loadServiceUserClientConfig() *asfclient.HTTPConfig {
	c := client.LoadHTTPConfig()
	c.SetCredentials(
		viper.GetString("keystone.service_user.id"),
		viper.GetString("keystone.service_user.password"),
	)
	c.Scope = kstypes.NewScope(
		viper.GetString("keystone.service_user.domain_id"),
		viper.GetString("keystone.service_user.domain_name"),
		viper.GetString("keystone.service_user.project_id"),
		viper.GetString("keystone.service_user.project_name"),
	)
	return c
}

func (s *Server) startVNCReplicator(endpointStore *endpoint.Store, auth KeystoneController) (err error) {
	s.VNCReplicator, err = replication.New(endpointStore, auth)
	if err != nil {
		return err
	}
	return s.VNCReplicator.Start()
}

func (s *Server) setupHomepage() {
	dh := NewHandler()

	services.RegisterSingularPaths(func(path string, name string) {
		dh.Register(path, "", name, "resource-base")
	})
	services.RegisterPluralPaths(func(path string, name string) {
		dh.Register(path, "", name, "collection")
	})

	dh.Register(FQNameToIDPath, "POST", "name-to-id", "action")
	dh.Register(IDToFQNamePath, "POST", "id-to-name", "action")
	dh.Register(UserAgentKVPath, "POST", UserAgentKVPath, "action")
	dh.Register(services.RefUpdatePath, "POST", services.RefUpdatePath, "action")
	dh.Register(services.RefRelaxForDeletePath, "POST", services.RefRelaxForDeletePath, "action")
	dh.Register(services.PropCollectionUpdatePath, "POST", services.PropCollectionUpdatePath, "action")
	dh.Register(services.SetTagPath, "POST", services.SetTagPath, "action")
	dh.Register(services.ChownPath, "POST", services.ChownPath, "action")
	dh.Register(services.IntPoolPath, "GET", services.IntPoolPath, "action")
	dh.Register(services.IntPoolPath, "POST", services.IntPoolPath, "action")
	dh.Register(services.IntPoolPath, "DELETE", services.IntPoolPath, "action")
	dh.Register(services.IntPoolsPath, "POST", services.IntPoolsPath, "action")
	dh.Register(services.IntPoolsPath, "DELETE", services.IntPoolsPath, "action")
	dh.Register(services.ObjPerms, "GET", services.ObjPerms, "action")

	// TODO: register sync?

	// TODO action resources
	// TODO documentation
	// TODO VN IP alloc
	// TODO VN IP free
	// TODO subnet IP count
	// TODO security policy draft

	s.Echo.GET("/", dh.Handle)
}

func (s *Server) setupWatchAPI() {
	if !viper.GetBool("cache.enabled") {
		return
	}
	s.Echo.GET(WatchPath, s.watchHandler)
}

func (s *Server) setupActionResources(cs *services.ContrailService) {
	s.Echo.POST(FQNameToIDPath, cs.RESTFQNameToUUID)
	s.Echo.POST(IDToFQNamePath, cs.RESTIDToFQName)
	s.Echo.POST(UserAgentKVPath, cs.RESTUserAgentKV)
	s.Echo.POST(services.UploadCloudKeysPath, cs.RESTUploadCloudKeys)
}

type recorderTask struct {
	Request *client.Request `yaml:"request,omitempty"`
	Expect  interface{}     `yaml:"expect,omitempty"`
}

// Run runs Server.
func (s *Server) Run() error {
	defer func() {
		if err := s.Close(); err != nil {
			s.log.WithError(err).Error("Closing DBService failed")
		}
	}()

	if viper.GetBool("server.tls.enabled") {
		return s.Echo.StartTLS(
			viper.GetString("server.address"),
			viper.GetString("server.tls.cert_file"),
			viper.GetString("server.tls.key_file"),
		)
	}

	return s.Echo.Start(viper.GetString("server.address"))
}

// Close closes Server.
func (s *Server) Close() error {
	if s.VNCReplicator != nil {
		s.VNCReplicator.Stop()
	}
	s.Proxy.StopEndpointsSync()
	return s.DBService.Close()
}
