package testutil

import (
	"testing"

	"github.com/Juniper/contrail/pkg/apisrv"
	pkglog "github.com/Juniper/contrail/pkg/log"
	"github.com/sirupsen/logrus"
)

// APIServerClient is API Server client for integration tests purposes.
type APIServerClient struct {
	*apisrv.Client
	log *logrus.Entry
}

// NewAPIServerClient is APIServerClient constructor.
func NewAPIServerClient(t *testing.T) *APIServerClient {
	l := pkglog.NewLogger("api-server-client")
	l.WithFields(logrus.Fields{"endpoint": apiServerEndpoint}).Debug("Connecting to API Server")
	//c := apisrv.NewClient(serverEndpoint, serverID, serverPassword)
	c := &apisrv.Client{}
	err := c.Login()
	if err != nil {
		t.Fatal("Cannot connect to API APIServer")
	}
	return &APIServerClient{
		Client: c,
		log:    l,
	}
}

// Create creates resource in server and returns its ID.
func (s *APIServer) Create(kvs map[string]interface{}) {
	s.log.WithField("key-values", kvs).Debug("Creating resource in API Server")
	// TODO: implement
}
