package apisrv

import (
	"database/sql"
	"net/url"
	"strings"
	"time"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"google.golang.org/grpc"

	"github.com/Juniper/contrail/pkg/apisrv/keystone"
	"github.com/Juniper/contrail/pkg/common"
	"github.com/Juniper/contrail/pkg/db"
	"github.com/Juniper/contrail/pkg/serviceif"
	"github.com/Juniper/contrail/pkg/services"

	etcdclient "github.com/Juniper/contrail/pkg/db/etcd"
	log "github.com/sirupsen/logrus"
)

//Server represents Intent API Server.
type Server struct {
	Echo     *echo.Echo
	DB       *sql.DB
	Keystone *keystone.Keystone
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
	dbService := db.NewService(s.DB, viper.GetString("database.dialect"))
	serviceChain = append(serviceChain, dbService)

	serviceif.Chain(serviceChain)

	return service
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
	e.Use(middleware.Logger())
	//e.Use(middleware.Recover())
	//e.Use(middleware.BodyLimit("10M"))

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
	if staticPath != nil {
		for prefix, root := range staticPath {
			e.Static(prefix, root)
		}
	}

	proxy := viper.GetStringMapStringSlice("proxy")
	if proxy != nil {
		for prefix, target := range proxy {
			g := e.Group(prefix)
			g.Use(removePathPrefixMiddleware(prefix))

			t, err := url.Parse(target[0])
			if err != nil {
				return errors.Wrapf(err, "bad proxy target URL: %s", target[0])
			}
			g.Use(proxyMiddleware(t, viper.GetBool("server.proxy.insecure")))
		}
	}
	keystoneAuthURL := viper.GetString("keystone.authurl")
	if keystoneAuthURL != "" {
		e.Use(keystone.AuthMiddleware(keystoneAuthURL,
			viper.GetBool("keystone.insecure"),
			[]string{
				"/v3/auth/tokens",
				"/public"}))
	} else if viper.GetBool("no_auth") {
		e.Use(noAuthMiddleware())
	}
	localKeystone := viper.GetBool("keystone.local")
	if localKeystone {
		keystone, err := keystone.Init(e)
		if err != nil {
			return errors.Wrap(err, "Failed to init local keystone server")
		}
		s.Keystone = keystone
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
					keystone.AuthInterceptor(keystoneAuthURL, viper.GetBool("keystone.insecure"))))
		} else if viper.GetBool("no_auth") {
			grpcServer = grpc.NewServer(
				grpc.UnaryInterceptor(
					noAuthInterceptor()))
		}
		services.RegisterContrailServiceServer(grpcServer, service)
		e.Use(gRPCMiddleware(grpcServer))
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
	return s.DB.Close()
}
