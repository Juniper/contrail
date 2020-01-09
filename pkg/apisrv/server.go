package apisrv

import (
	"crypto/tls"
	"net/http"
	"net/url"

	"github.com/Juniper/asf/pkg/client"
	"github.com/Juniper/asf/pkg/db/basedb"
	"github.com/Juniper/asf/pkg/logutil"
	"github.com/Juniper/contrail/pkg/apisrv/baseapisrv"
	"github.com/Juniper/contrail/pkg/collector"
	"github.com/Juniper/contrail/pkg/collector/analytics"
	"github.com/Juniper/contrail/pkg/constants"
	"github.com/Juniper/contrail/pkg/db"
	"github.com/Juniper/contrail/pkg/db/cache"
	"github.com/Juniper/contrail/pkg/keystone"
	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/neutron"
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
)

// Server HTTP paths.
const (
	FQNameToIDPath  = "fqname-to-id"
	IDToFQNamePath  = "id-to-fqname"
	UserAgentKVPath = "useragent-kv"
	WatchPath       = "watch"
)

// Server represents Intent API Server.
type Server struct {
	Server                     *baseapisrv.Server
	endpointStore              endpointStore
	Keystone                   *keystone.Keystone
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
	log                        *logrus.Entry
}

// NewServer makes a server.
// nolint: gocyclo
func NewServer(es endpointStore) (*Server, error) {
	baseServer, err := baseapisrv.NewServer(authGRPCOpts())
	if err != nil {
		return nil, err
	}

	s := &Server{
		Server:        baseServer,
		endpointStore: es,
		log:           logutil.NewLogger("contrail-api-server"),
	}

	s.Collector, err = makeCollector()
	if err != nil {
		return nil, err
	}
	s.registerCollector(s.Collector)

	sqlDB, err := basedb.ConnectDB(analytics.WithCommitLatencyReporting(s.Collector))
	if err != nil {
		return nil, err
	}
	s.DBService = db.NewService(sqlDB)

	cs, err := s.setupService()
	if err != nil {
		return nil, err
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
		n := s.setupNeutronService(cs)
		// TODO(Witaut): Don't use Echo - an internal detail of Server.
		n.RegisterNeutronAPI(s.Server.Echo)
	}

	if err = s.registerStaticProxyEndpoints(); err != nil {
		return nil, errors.Wrap(err, "failed to register static proxy endpoints")
	}

	s.serveDynamicProxy(s.endpointStore)

	if err = s.setupAuthMiddleware(); err != nil {
		return nil, errors.Wrap(err, "failed to set up auth middleware")
	}

	if viper.GetBool("keystone.local") {
		// TODO(Witaut): Don't use Echo - an internal detail of Server.
		var k *keystone.Keystone
		k, err = keystone.Init(s.Server.Echo, s.endpointStore)
		if err != nil {
			return nil, errors.Wrap(err, "Failed to init local keystone server")
		}
		s.Keystone = k
	}

	if viper.GetBool("server.enable_grpc") {
		// TODO(Witaut): Don't use GRPCServer - an internal detail of Server.
		services.RegisterContrailServiceServer(s.Server.GRPCServer, s.Service)
		services.RegisterIPAMServer(s.Server.GRPCServer, s.IPAMServer)
		services.RegisterChownServer(s.Server.GRPCServer, s.ChownServer)
		services.RegisterSetTagServer(s.Server.GRPCServer, s.SetTagServer)
		services.RegisterRefRelaxServer(s.Server.GRPCServer, s.RefRelaxServer)
		services.RegisterFQNameToIDServer(s.Server.GRPCServer, s.FQNameToIDServer)
		services.RegisterIDToFQNameServer(s.Server.GRPCServer, s.IDToFQNameServer)
		services.RegisterUserAgentKVServer(s.Server.GRPCServer, s.UserAgentKVServer)
		services.RegisterPropCollectionUpdateServer(s.Server.GRPCServer, s.PropCollectionUpdateServer)
	}

	if viper.GetBool("homepage.enabled") {
		// TODO Move this to Server
		s.setupHomepage()
	}

	s.setupWatchAPI()
	s.setupActionResources(cs)

	return s, nil
}

func authGRPCOpts() (opts []grpc.ServerOption) {
	keystoneAuthURL, keystoneInsecure := viper.GetString("keystone.authurl"), viper.GetBool("keystone.insecure")
	if keystoneAuthURL != "" {
		opts = append(opts, grpc.UnaryInterceptor(keystone.AuthInterceptor(
			keystoneAuthURL,
			keystoneInsecure,
		)))
	} else if viper.GetBool("no_auth") {
		opts = append(opts, grpc.UnaryInterceptor(noAuthInterceptor()))
	}
	return opts
}

func (s *Server) setupAuthMiddleware() error {
	keystoneAuthURL, keystoneInsecure := viper.GetString("keystone.authurl"), viper.GetBool("keystone.insecure")
	if keystoneAuthURL != "" {
		var skipPaths []string
		skipPaths, err := keystone.GetAuthSkipPaths()
		if err != nil {
			return errors.Wrap(err, "failed to setup paths skipped from authentication")
		}
		// TODO(Witaut): Don't use Echo - an internal detail of Server.
		s.Server.Echo.Use(keystone.AuthMiddleware(keystoneAuthURL, keystoneInsecure, skipPaths))
	} else if viper.GetBool("no_auth") {
		// TODO(Witaut): Don't use Echo - an internal detail of Server.
		s.Server.Echo.Use(noAuthMiddleware())
	}
	return nil
}

func makeCollector() (c collector.Collector, err error) {
	cfg := &analytics.Config{}
	if err = viper.UnmarshalKey("collector", cfg); err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal collector config")
	}
	if c, err = analytics.NewCollector(cfg); err != nil {
		return nil, errors.Wrap(err, "failed to create collector")
	}
	return c, nil
}

func (s *Server) registerCollector(c collector.Collector) {
	analytics.AddLoggerHook(c)
	// TODO(Witaut): Don't use Echo - an internal detail of Server.
	s.Server.Echo.Use(middleware.BodyDump(func(
		ctx echo.Context, reqBody, resBody []byte,
	) {
		c.Send(analytics.RESTAPITrace(ctx, reqBody, resBody))
	}))
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

	// TODO(Witaut): Don't use Echo - an internal detail of Server.
	cs.RegisterRESTAPI(s.Server.Echo)
	return cs, nil
}

func (s *Server) setupNeutronService(cs services.Service) *neutron.Server {
	return &neutron.Server{
		ReadService:       cs,
		WriteService:      cs,
		UserAgentKV:       s.UserAgentKVServer,
		IDToFQNameService: s.IDToFQNameServer,
		FQNameToIDService: s.FQNameToIDServer,
		InTransactionDoer: s.DBService,
		Log:               logutil.NewLogger("neutron-server"),
	}
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

		// TODO(Witaut): Don't use Echo - an internal detail of Server.
		g := s.Server.Echo.Group(prefix)
		g.Use(removePathPrefixMiddleware(prefix))
		g.Use(proxyMiddleware(t, viper.GetBool("server.proxy.insecure")))
	}
	return nil
}

func (s *Server) serveDynamicProxy(es endpointStore) {
	config := loadDynamicProxyConfig()
	// TODO(Witaut): Don't use Echo - an internal detail of Server.
	s.Server.Echo.Group(config.Path, dynamicProxyMiddleware(es, config))

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

	// TODO(Witaut): Don't use Echo - an internal detail of Server.
	s.Server.Echo.GET("/", dh.Handle)
}

func (s *Server) setupWatchAPI() {
	if !viper.GetBool("cache.enabled") {
		return
	}
	// TODO(Witaut): Don't use Echo - an internal detail of Server.
	s.Server.Echo.GET(WatchPath, s.watchHandler)
}

func (s *Server) setupActionResources(cs *services.ContrailService) {
	// TODO(Witaut): Don't use Echo - an internal detail of Server.
	s.Server.Echo.POST(FQNameToIDPath, cs.RESTFQNameToUUID)
	s.Server.Echo.POST(IDToFQNamePath, cs.RESTIDToFQName)
	s.Server.Echo.POST(UserAgentKVPath, cs.RESTUserAgentKV)
	s.Server.Echo.POST(services.UploadCloudKeysPath, cs.RESTUploadCloudKeys)
}

// Run runs Server.
func (s *Server) Run() error {
	defer func() {
		if err := s.Close(); err != nil {
			s.log.WithError(err).Error("Closing DBService failed")
		}
	}()

	return s.Server.Run()
}

// Close closes Server.
func (s *Server) Close() error {
	s.log.Info("Closing server")
	s.Proxy.StopEndpointsSync()
	err := s.DBService.Close()
	s.log.Info("Server closed")
	return err
}
