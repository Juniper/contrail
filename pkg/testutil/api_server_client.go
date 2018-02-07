package testutil

import (
	"testing"

	"github.com/Juniper/contrail/pkg/apisrv"
	"github.com/Juniper/contrail/pkg/apisrv/keystone"
	pkglog "github.com/Juniper/contrail/pkg/log"
	"github.com/sirupsen/logrus"
)

// APIServerClient is API Server client for integration tests purposes.
type APIServerClient struct {
	*apisrv.Client
	log *logrus.Entry
}

// NewAPIServerClient is APIServerClient constructor.
func NewAPIServerClient(t *testing.T, apiServerURL string) *APIServerClient {
	l := pkglog.NewLogger("api-server-client")
	l.WithFields(logrus.Fields{"endpoint": apiServerURL}).Debug("Connecting to API Server")
	c := apisrv.NewClient(
		apiServerURL,
		apiServerURL+AuthEndpointSuffix,
		AdminUserID,
		AdminPassword,
		&keystone.Scope{
			Project: &keystone.Project{
				ID: AdminProjectID,
			},
		},
	)
	err := c.Login()
	if err != nil {
		t.Fatal("Cannot connect to API Server")
	}
	return &APIServerClient{
		Client: c,
		log:    l,
	}
}

// Create creates resource in server and returns its ID.
//func (s *APIServer) Create(kvs map[string]interface{}) {
//	s.log.WithField("key-values", kvs).Debug("Creating resource in API Server")
//	// TODO: implement
//}
