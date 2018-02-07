package integration

import (
	"testing"

	"github.com/Juniper/contrail/pkg/apisrv"
	"github.com/Juniper/contrail/pkg/apisrv/keystone"
	pkglog "github.com/Juniper/contrail/pkg/log"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/require"
)

// HTTPAPIClient is API Server client for tests purposes.
type HTTPAPIClient struct {
	*apisrv.Client
	log *logrus.Entry
}

// NewHTTPAPIClient creates HTTP client of API Server.
func NewHTTPAPIClient(t *testing.T, apiServerURL string) *HTTPAPIClient {
	l := pkglog.NewLogger("api-server-client")
	l.WithFields(logrus.Fields{"endpoint": apiServerURL}).Debug("Connecting to API Server")
	c := apisrv.NewClient(
		apiServerURL,
		apiServerURL+authEndpointSuffix,
		aliceUserID,
		alicePassword,
		defaultDomainID,
		true,
		&keystone.Scope{
			Project: &keystone.Project{
				ID:   adminProjectID,
				Name: adminProjectName,
				Domain: &keystone.Domain{
					ID: defaultDomainID,
				},
			},
		},
	)
	c.Debug = true

	err := c.Login()
	require.NoError(t, err, "connecting API Server failed")

	return &HTTPAPIClient{
		Client: c,
		log:    l,
	}
}
