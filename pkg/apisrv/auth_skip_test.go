package apisrv_test

import (
	"crypto/tls"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/Juniper/contrail/pkg/testutil/integration"
)

const (
	testAuthSkipFile = "./test_data/test_auth_skip.yml"
)

func TestContrailClusterAuthSkip(t *testing.T) {
	var testScenario integration.TestScenario
	err := integration.LoadTestScenario(&testScenario, testAuthSkipFile, nil)
	assert.NoError(t, err, "failed to load test data")
	cleanup := integration.RunDirtyTestScenario(t, &testScenario, server)
	defer cleanup()

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	// request without auth for a path not in skip path
	url := server.TestServer.URL + "/contrail-clusters"
	res, err := client.Get(url)
	assert.NoError(t, err, "failed to read contrail-clusters")
	assert.Equal(t, res.StatusCode, 401, "unexpected status code")
	// request without auth for a path/query not in skip path
	url = server.TestServer.URL + "/contrail-clusters?fields=uuid,name,fq_name"
	res, err = client.Get(url)
	assert.NoError(t, err, "failed to read contrail-clusters with query")
	assert.Equal(t, res.StatusCode, 401, "unexpected status code when using incorrect query")
	// request without auth for a path/query present in skip path
	url = server.TestServer.URL + "/contrail-clusters?fields=uuid,name"
	res, err = client.Get(url)
	assert.NoError(t, err, "failed to read contrail-clusters with query")
	defer res.Body.Close() // nolint: errcheck
	assert.Equal(t, res.StatusCode, 200, "unexpected status code when using correct query")
	assert.NoError(t, err, "failed to list contrail-clusters")
}
