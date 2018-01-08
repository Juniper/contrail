package apisrv

import (
	"database/sql"
	"net/url"
	"time"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/pkg/errors"
	"github.com/spf13/viper"

	"github.com/Juniper/contrail/pkg/apisrv/keystone"
	"github.com/Juniper/contrail/pkg/common"
	"github.com/Juniper/contrail/pkg/generated/api"

	log "github.com/sirupsen/logrus"
)

const (
	retryDB     = 10
	retryDBWait = 10
)

//Server represents Intent API Server.
type Server struct {
	Echo *echo.Echo
	DB   *sql.DB
}

// NewServer makes a server
func NewServer() (*Server, error) {
	server := &Server{
		Echo: echo.New(),
	}
	return server, nil
}

//Init setup the server.
func (s *Server) Init() error {
	common.SetLogLevel()
	db, err := common.ConnectDB()
	if err != nil {
		return errors.Wrap(err, "Init DB failed")
	}
	s.DB = db

	e := s.Echo

	//e.Use(middleware.Recover())
	e.Use(middleware.BodyLimit("10M"))

	for _, a := range api.APIs {
		a.SetDB(s.DB)
		common.RegisterAPI(a)
	}
	common.Routes(e)
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
		for prefix, targetStrings := range proxy {
			targets := []*middleware.ProxyTarget{}
			for _, targetString := range targetStrings {
				targetURL, err := url.Parse(targetString)
				if err != nil {
					e.Logger.Fatal(err)
				}
				targets = append(targets,
					&middleware.ProxyTarget{
						URL: targetURL,
					})
			}

			g := e.Group(prefix)
			g.Use(removePathPrefixMiddleware(prefix))
			g.Use(middleware.Proxy(&middleware.RoundRobinBalancer{
				Targets: targets}))
		}

	}
	keystoneAuthURL := viper.GetString("keystone.authurl")
	if keystoneAuthURL != "" {
		e.Use(keystone.AuthMiddleware(keystoneAuthURL,
			[]string{
				"/v3/auth/tokens",
				"/public"}))
	}
	localKeystone := viper.GetBool("keystone.local")
	if localKeystone {
		err := keystone.Init(e)
		if err != nil {
			return errors.Wrap(err, "Failed to init local keystone server")
		}
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
