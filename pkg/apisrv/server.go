package apisrv

import (
	"database/sql"
	"time"

	"github.com/labstack/echo"
	"github.com/pkg/errors"
	"github.com/spf13/viper"

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
	err := server.init()
	if err != nil {
		return nil, err
	}
	return server, nil
}

func (s *Server) init() error {
	common.SetLogLevel()
	db, err := common.ConnectDB()
	if err != nil {
		return errors.Wrap(err, "Init DB failed")
	}
	s.DB = db

	e := s.Echo

	for _, a := range api.APIs {
		a.SetDB(s.DB)
		common.RegisterAPI(a)
	}
	common.Routes(e)
	readTimeout := viper.GetInt("server.read_timeout")
	writeTimeout := viper.GetInt("server.write_timeout")
	e.Server.ReadTimeout = time.Duration(readTimeout) * time.Second
	e.Server.WriteTimeout = time.Duration(writeTimeout) * time.Second
	return nil
}

// Run run's application
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

func (s *Server) Close() error {
	return s.DB.Close()
}
