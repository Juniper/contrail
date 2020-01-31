package apisrv

import (
	"crypto/tls"
	"net/http"
	"strings"

	"github.com/Juniper/asf/pkg/apisrv/baseapisrv"
	"github.com/Juniper/asf/pkg/db/basedb"
	"github.com/Juniper/asf/pkg/logutil"
	"github.com/Juniper/contrail/pkg/collector"
	"github.com/Juniper/contrail/pkg/collector/analytics"
	"github.com/Juniper/contrail/pkg/constants"
	"github.com/Juniper/contrail/pkg/db"
	"github.com/Juniper/contrail/pkg/db/cache"
	"github.com/Juniper/contrail/pkg/keystone"
	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/neutron"
	"github.com/Juniper/contrail/pkg/proxy"
	"github.com/Juniper/contrail/pkg/services"
	"github.com/Juniper/contrail/pkg/types"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	asfkeystone "github.com/Juniper/asf/pkg/keystone"
	etcdclient "github.com/Juniper/contrail/pkg/db/etcd"
)

// Server represents Intent API Server.
type Server struct {
	Server            *baseapisrv.Server
	Keystone          *keystone.Keystone
	DBService         *db.Service
	Proxy             *proxy.Proxy
	Service           services.Service
	UserAgentKVServer services.UserAgentKVServer
	FQNameToIDServer  services.FQNameToIDServer
	IDToFQNameServer  services.IDToFQNameServer
	Collector         collector.Collector
	log               *logrus.Entry
}

type endpointStore interface {
	RemoveDeleted(endpoints map[string]*models.Endpoint)
	Contains(prefix string) bool
	GetEndpointURLs(prefix string, endpointKey string) []string
	UpdateEndpoint(prefix string, endpoint *models.Endpoint) error
	InitScope(prefix string)
	GetEndpointURL(clusterID, prefix string) (string, bool)
	GetPassword(prefix string, endpointKey string) string
	GetUsername(prefix string, endpointKey string) string
	Remove(prefix string)
}

// NewServer makes a server.
// nolint: gocyclo
func NewServer(es endpointStore, cache *cache.DB) (*Server, error) {
	s := &Server{
		log: logutil.NewLogger("contrail-api-server"),
	}

	var plugins []baseapisrv.APIPlugin

	var err error
	s.Collector, err = makeCollector()
	if err != nil {
		return nil, err
	}
	analytics.AddLoggerHook(s.Collector)
	plugins = append(plugins, analytics.BodyDumpPlugin{Collector: s.Collector})

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
	s.UserAgentKVServer = cs
	s.FQNameToIDServer = cs
	s.IDToFQNameServer = cs

	plugins = append(plugins, cs)

	if viper.GetBool("server.enable_vnc_neutron") {
		plugins = append(plugins, s.setupNeutronService(cs))
	}

	staticProxyPlugin, err := proxy.NewStaticProxyPluginByViper()
	if err != nil {
		return nil, err
	}
	plugins = append(plugins, staticProxyPlugin)

	s.Proxy = proxy.NewProxyFromViper(es, s.DBService)
	s.Proxy.StartEndpointsSync()
	plugins = append(plugins, s.Proxy)

	if viper.GetBool("keystone.local") {
		var k *keystone.Keystone
		k, err = keystone.Init(es)
		if err != nil {
			return nil, errors.Wrap(err, "Failed to init local keystone server")
		}
		s.Keystone = k
		plugins = append(plugins, k)
	}

	if viper.GetBool("cache.enabled") {
		plugins = append(plugins, cache)
	}

	plugins = append(plugins, services.UploadCloudKeysPlugin{})

	s.Server, err = baseapisrv.NewServer(plugins, noAuthPaths())
	if err != nil {
		return nil, err
	}

	return s, nil
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
			Keystone: &asfkeystone.Client{
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

func noAuthPaths() []string {
	return []string{
		"/v3/auth/tokens", // TODO(mblotniak): Is this ever used?
		strings.Join([]string{
			models.ContrailClusterPluralPath,
			"?fields=",
			models.ContrailClusterFieldUUID,
			",",
			models.ContrailClusterFieldName,
		}, ""),
	}
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
