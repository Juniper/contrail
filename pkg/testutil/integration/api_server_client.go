package integration

import (
	"testing"

	"github.com/Juniper/contrail/pkg/apisrv"
	"github.com/Juniper/contrail/pkg/apisrv/keystone"
	pkglog "github.com/Juniper/contrail/pkg/log"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/require"
)

// APIServerClient is API Server client for integration tests purposes.
type APIServerClient struct {
	*apisrv.Client
	log *logrus.Entry
}

// NewAPIServerClient creates APIServerClient.
func NewAPIServerClient(t *testing.T, apiServerURL string) *APIServerClient {
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

	err := c.Login()
	require.NoError(t, err, "connecting API Server failed")

	return &APIServerClient{
		Client: c,
		log:    l,
	}
}
