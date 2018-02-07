package integration

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/Juniper/contrail/pkg/apisrv"
	"github.com/Juniper/contrail/pkg/apisrv/keystone"
	pkglog "github.com/Juniper/contrail/pkg/log"
	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// HTTPAPIClient is API Server client for tests purposes.
type HTTPAPIClient struct {
	*apisrv.Client
	log *logrus.Entry
}

// NewHTTPAPIClient creates HTTP client of API Server.
func NewHTTPAPIClient(t *testing.T, apiServerURL string) *HTTPAPIClient {
	l := pkglog.NewLogger("http-api-client")
	l.WithFields(logrus.Fields{"endpoint": apiServerURL}).Debug("Connecting to API Server")
	c := apisrv.NewClient(
		apiServerURL,
		apiServerURL+authEndpointSuffix,
		adminUserID,
		adminUserPassword,
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

// CreateResource creates resource in API Server and checks operation error.
func (c *HTTPAPIClient) CreateResource(t *testing.T, path string, requestData interface{}) {
	var responseData interface{}
	r, err := c.Create(
		path,
		requestData,
		&responseData,
	)
	c.log.WithFields(logrus.Fields{
		"requestData":  requestData,
		"response":     r,
		"responseData": responseData,
	}).Debug("Got Create response")
	assert.NoError(t, err, fmt.Sprintf("creating resource failed\n requestData: %v\n response: %+v", requestData, r))
}

// GetResource gets resource from API Server and checks operation error.
func (c *HTTPAPIClient) GetResource(t *testing.T, path string, responseData interface{}) {
	r, err := c.Read(path, &responseData)
	c.log.WithFields(logrus.Fields{
		"response":     r,
		"responseData": responseData,
	}).Debug("Got Get response")
	assert.NoError(t, err, fmt.Sprintf("getting resource failed\n response: %+v", r))
}

// DeleteResource deletes resource from API Server and checks operation error.
func (c *HTTPAPIClient) DeleteResource(t *testing.T, path string) {
	r, err := c.Delete(path, nil)
	c.log.WithField("response", r).Debug("Got Delete response")
	assert.NoError(t, err, "deleting resource failed\n response: %+v", r)
}

// CheckResourceDoesNotExist checks that there is no resource with given path.
func (c *HTTPAPIClient) CheckResourceDoesNotExist(t *testing.T, path string) {
	r, err := c.Do(
		echo.GET,
		path,
		nil,
		nil,
		[]int{http.StatusNotFound},
	)
	assert.NoError(t, err, "getting resource failed\n response: %+v", r)
}
