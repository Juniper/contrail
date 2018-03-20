package apisrv

import (
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/flosch/pongo2"
	"github.com/stretchr/testify/assert"
)

func TestProxyEndpoint(t *testing.T) {
	// test server to serve keystone private url
	ksPrivate := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "cluster_a_keystone_private_url_token")
	}))
	defer ksPrivate.Close()
	ksPrivate.URL = ksPrivate.URL + "/v3"
	// test server to serve keystone public url
	ksPublic := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "cluster_a_keystone_public_url_token")
	}))
	defer ksPublic.Close()
	ksPublic.URL = ksPublic.URL + "/v3"

	context := pongo2.Context{
		"cluster_a_keystone_private_url": ksPrivate.URL,
		"cluster_a_keystone_public_url":  ksPublic.URL,
	}

	testFile := GetTestFromTemplate(t, "./test_data/test_endpoint.tmpl", context)
	defer os.Remove(testFile) // remove tempfile after test

	var testScenario TestScenario
	LoadTestScenario(&testScenario, testFile)
	RunTestScenario(t, &testScenario)

	// Make sure endpoints are created
	e, err := server.Proxy.readEndpoints()
	assert.NoError(t, err, "failed to read the created endpoints")
	if len(e) != 1 {
		assert.NoError(t, errors.New("Endpoint NotFound"),
			"Expected number of endpoints not present")
	}

	for _, client := range testScenario.Clients {
		var publicRes map[string]interface{}
		client.Read("/keystone/v3", &publicRes)
		var privateRes map[string]interface{}
		client.Read("/keystone/private/v3", &privateRes)

	}
	//res, _ := http.Get(ksPrivate.URL)
	//token, _ := ioutil.ReadAll(res.Body)
	//res.Body.Close()
	//fmt.Printf("%s", token)

}
