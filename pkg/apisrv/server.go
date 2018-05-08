package apisrv

import (
	"database/sql"
	"encoding/json"
	"net/url"
	"strings"
	"sync"
	"time"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"google.golang.org/grpc"

	"github.com/Juniper/contrail/pkg/apisrv/keystone"
	"github.com/Juniper/contrail/pkg/apisrv/types"
	"github.com/Juniper/contrail/pkg/common"
	"github.com/Juniper/contrail/pkg/db"
	"github.com/Juniper/contrail/pkg/serviceif"
	"github.com/Juniper/contrail/pkg/services"

	apicommon "github.com/Juniper/contrail/pkg/apisrv/common"
	etcdclient "github.com/Juniper/contrail/pkg/db/etcd"
	log "github.com/sirupsen/logrus"
)

//Server represents Intent API Server.
type Server struct {
	Echo      *echo.Echo
	DB        *sql.DB
	Keystone  *keystone.Keystone
	dbService serviceif.Service
	Proxy     *proxyService
}

// NewServer makes a server
func NewServer() (*Server, error) {
	server := &Server{
		Echo: echo.New(),
	}
	return server, nil
}

//SetupService setup service.
//Application with custom logics can embed server struct, and overwrite
//this method.
func (s *Server) SetupService() serviceif.Service {
	var serviceChain []serviceif.Service
	service := &services.ContrailService{
		BaseService: serviceif.BaseService{},
	}

	serviceChain = append(serviceChain, service)
	service.RegisterRESTAPI(s.Echo)

	serviceChain = append(serviceChain, &types.ContrailTypeLogicService{
		BaseService: serviceif.BaseService{},
	})

	etcdNotifierEnabled := viper.GetBool("etcd_notifier.enabled")
	if etcdNotifierEnabled {
		etcdNotifierServers := strings.Split(viper.GetString("etcd_notifier.servers"), ",")
		etcdNotifierPath := viper.GetString("etcd_notifier.path")
		etcdNotifierService, err := etcdclient.NewEtcdNotifierService(etcdNotifierServers, etcdNotifierPath)
		if err == nil {
			log.Println("Adding ETCD Notifier Service.")
			serviceChain = append(serviceChain, etcdNotifierService)
		}
	}

	// Put DB Service at the end
	serviceChain = append(serviceChain, s.dbService)

	serviceif.Chain(serviceChain)

	return service
}

func (s *Server) serveDynamicProxy(endpointStore *apicommon.EndpointStore) {
	s.Proxy = newProxyService(s.Echo, endpointStore, s.dbService)
	s.Proxy.serve()
}

//Init setup the server.
func (s *Server) Init() error {
	common.SetLogLevel()
	sqlDB, err := db.ConnectDB()
	if err != nil {
		return errors.Wrap(err, "Init DB failed")
	}
	s.DB = sqlDB
	e := s.Echo
	if viper.GetBool("server.log_api") {
		e.Use(middleware.Logger())
	}
	//e.Use(middleware.Recover())
	//e.Use(middleware.BodyLimit("10M"))

	s.dbService = db.NewService(s.DB, viper.GetString("database.dialect"))
	service := s.SetupService()

	readTimeout := viper.GetInt("server.read_timeout")
	writeTimeout := viper.GetInt("server.write_timeout")
	e.Server.ReadTimeout = time.Duration(readTimeout) * time.Second
	e.Server.WriteTimeout = time.Duration(writeTimeout) * time.Second

	cors := viper.GetString("cors")

	if cors != "" {
		log.Printf("Enabling CORS for %s", cors)
		if cors == "*" {
			log.Printf("cors for * have security issue. DO NOT USE THIS IN PRODUCTION")
		}
		e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
			AllowOrigins:  []string{cors},
			AllowMethods:  []string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
			AllowHeaders:  []string{"X-Auth-Token", "Content-Type"},
			ExposeHeaders: []string{"X-Total-Count"},
		}))
	}

	staticPath := viper.GetStringMapString("static_files")
	for prefix, root := range staticPath {
		e.Static(prefix, root)
	}

	proxy := viper.GetStringMapStringSlice("proxy")
	for prefix, target := range proxy {
		g := e.Group(prefix)
		g.Use(removePathPrefixMiddleware(prefix))

		t, err := url.Parse(target[0])
		if err != nil {
			return errors.Wrapf(err, "bad proxy target URL: %s", target[0])
		}
		g.Use(proxyMiddleware(t, viper.GetBool("server.proxy.insecure")))
	}
	// serve dynamic proxy based on configured endpoints
	endpointStore := apicommon.MakeEndpointStore() // sync map to store proxy endpoints
	s.serveDynamicProxy(endpointStore)

	keystoneAuthURL := viper.GetString("keystone.authurl")
	var keystoneClient *keystone.KeystoneClient
	if keystoneAuthURL != "" {
		keystoneClient = keystone.NewKeystoneClient(keystoneAuthURL,
			viper.GetBool("keystone.insecure"))
		e.Use(keystone.AuthMiddleware(keystoneClient,
			[]string{
				"/keystone/v3/auth/tokens",
				"/proxy/keystone/v3/auth/tokens",
				"/keystone/v3/auth/projects",
				"/v3/auth/tokens",
				"/public"},
			endpointStore))
	} else if viper.GetBool("no_auth") {
		e.Use(noAuthMiddleware())
	}
	localKeystone := viper.GetBool("keystone.local")
	if localKeystone {
		k, err := keystone.Init(e, endpointStore, keystoneClient)
		if err != nil {
			return errors.Wrap(err, "Failed to init local keystone server")
		}
		s.Keystone = k
	}
	if viper.GetBool("enable_grpc") {
		if !viper.GetBool("tls.enabled") {
			log.Fatal("GRPC support requires TLS configuraion.")
		}
		log.Debug("enabling grpc")
		var grpcServer *grpc.Server
		if keystoneAuthURL != "" {
			grpcServer = grpc.NewServer(
				grpc.UnaryInterceptor(
					keystone.AuthInterceptor(keystoneClient, endpointStore)))
		} else if viper.GetBool("no_auth") {
			grpcServer = grpc.NewServer(
				grpc.UnaryInterceptor(
					noAuthInterceptor()))
		}
		services.RegisterContrailServiceServer(grpcServer, service)
		e.Use(gRPCMiddleware(grpcServer))
	}

	if viper.GetBool("recorder.enabled") {
		file := viper.GetString("recorder.file")
		scenario := &TestScenario{
			Workflow: []*Task{},
		}
		var mutex sync.Mutex
		e.Use(middleware.BodyDump(func(c echo.Context, requestBody, responseBody []byte) {
			var data interface{}
			err := json.Unmarshal(requestBody, &data)
			if err != nil {
				log.Debug("malformed json input")
			}
			var expected interface{}
			err = json.Unmarshal(responseBody, &expected)
			if err != nil {
				log.Debug("malformed json response")
			}
			task := &Task{
				Request: &Request{
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
			err = common.SaveFile(file, scenario)
			if err != nil {
				log.Warn(err)
			}
		}))
	}
	return nil
}

// Run runs server.
func (s *Server) Run() error {
	defer func() {
		err := s.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()
	e := s.Echo
	address := viper.GetString("address")
	tlsEnabled := viper.GetBool("tls.enabled")
	var keyFile, certFile string
	if tlsEnabled {
		keyFile = viper.GetString("tls.key_file")
		certFile = viper.GetString("tls.cert_file")

		e.Logger.Fatal(e.StartTLS(address, certFile, keyFile))
	} else {
		e.Logger.Fatal(e.Start(address))
	}
	return nil
}

//Close closes server resources
func (s *Server) Close() error {
	s.Proxy.stop()
	return s.DB.Close()
}
