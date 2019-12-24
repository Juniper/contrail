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

	asfclient "github.com/Juniper/asf/pkg/client"
	kstypes "github.com/Juniper/asf/pkg/keystone"
	etcdclient "github.com/Juniper/contrail/pkg/db/etcd"
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
	s := &Server{
		endpointStore: es,
		log:           logutil.NewLogger("contrail-api-server"),
	}

	var plugins []baseapisrv.APIPlugin

	var err error
	s.Collector, err = makeCollector()
	if err != nil {
		return nil, err
	}
	analytics.AddLoggerHook(s.Collector)
	plugins = append(plugins, bodyDumpPlugin{s.Collector})

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

	plugins = append(plugins, cs)

	if viper.GetBool("server.enable_vnc_neutron") {
		plugins = append(plugins, s.setupNeutronService(cs))
	}

	plugins = append(plugins, staticProxyPlugin{})

	config := loadDynamicProxyConfig()
	s.Proxy = newProxyService(s.endpointStore, s.DBService, config)
	s.Proxy.StartEndpointsSync()
	plugins = append(plugins, newDynamicProxyPlugin(s.endpointStore, config))

	authPlugins, err := authPlugins()
	if err != nil {
		return nil, errors.Wrap(err, "failed to set up auth middleware")
	}
	plugins = append(plugins, authPlugins...)

	if viper.GetBool("keystone.local") {
		var k *keystone.Keystone
		k, err = keystone.Init(s.endpointStore)
		if err != nil {
			return nil, errors.Wrap(err, "Failed to init local keystone server")
		}
		s.Keystone = k
		plugins = append(plugins, k)
	}

	if viper.GetBool("cache.enabled") {
		plugins = append(plugins, s.Cache)
	}

	s.Server, err = baseapisrv.NewServer(plugins)
	if err != nil {
		return nil, err
	}

	return s, nil
}

func authPlugins() (plugins []baseapisrv.APIPlugin, err error) {
	if viper.GetString("keystone.authurl") != "" {
		k, err := keystone.NewAuthPluginByViper()
		if err != nil {
			return nil, err
		}
		plugins = append(plugins, k)
	} else if viper.GetBool("no_auth") {
		plugins = append(plugins, noAuthPlugin{})
	}
	return plugins, nil
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

type bodyDumpPlugin struct {
	c collector.Collector
}

func (p bodyDumpPlugin) RegisterHTTPAPI(r baseapisrv.HTTPRouter) error {
	r.Use(fromEchoMiddlewareFunc(middleware.BodyDump(func(ctx echo.Context, reqBody, resBody []byte) {
		p.c.Send(analytics.RESTAPITrace(ctx, reqBody, resBody))
	})))
	return nil
}

// fromEchoMiddlewareFunc makes a baseapisrv.MiddlewareFunc from echo.MiddlewareFunc.
func fromEchoMiddlewareFunc(m echo.MiddlewareFunc) baseapisrv.MiddlewareFunc {
	return func(next baseapisrv.HandlerFunc) baseapisrv.HandlerFunc {
		return baseapisrv.HandlerFunc(m(echo.HandlerFunc(next)))
	}
}

func (p bodyDumpPlugin) RegisterGRPCAPI(r baseapisrv.GRPCRouter) error {
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

	return &services.ContrailService{
		BaseService:        services.BaseService{},
		DBService:          s.DBService,
		TypeValidator:      tv,
		MetadataGetter:     s.DBService,
		InTransactionDoer:  s.DBService,
		IntPoolAllocator:   s.DBService,
		RefRelaxer:         s.DBService,
		UserAgentKVService: s.DBService,
		Collector:          s.Collector,
	}, nil
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

type staticProxyPlugin struct{}

func (staticProxyPlugin) RegisterHTTPAPI(r baseapisrv.HTTPRouter) error {
	for prefix, targetURLs := range viper.GetStringMapStringSlice("server.proxy") {
		if len(targetURLs) == 0 {
			return errors.Errorf("no target URLs provided for prefix %v", prefix)
		}

		// TODO(dfurman): proxy requests to all provided target URLs
		t, err := url.Parse(targetURLs[0])
		if err != nil {
			return errors.Wrapf(err, "bad proxy target URL: %s", targetURLs[0])
		}

		r.Group(
			prefix,
			removePathPrefixMiddleware(prefix),
			proxyMiddleware(t, viper.GetBool("server.proxy.insecure")),
		)
	}
	return nil
}

func (staticProxyPlugin) RegisterGRPCAPI(r baseapisrv.GRPCRouter) error {
	return nil
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
