// +build integration

package apisrv

import (
	"testing"
	"time"
)

// TestProxyEndpointWithSleep tests the first part of TestProxyEndpoint,
// but verifies that endpoint updates are triggered every 2 seconds.
// TODO: Remove this test when proxyService switches to using events instead of Ticker.
func TestProxyEndpointWithSleep(t *testing.T) {
	// Create a cluster and its neutron endpoint
	clusterAName := "clusterA"
	testScenario, clusterANeutronPublic, clusterANeutronPrivate, cleanup1 := runEndpointTest(
		t, clusterAName, true)
	defer cleanup1()
	// remove tempfile after test
	defer clusterANeutronPrivate.Close()
	defer clusterANeutronPublic.Close()

	// wait for proxy endpoints to update
	time.Sleep(2 * time.Second)

	verifyProxyAndTestIt(t, testScenario, clusterAName, true)

	// create one more cluster/neutron endpoint for new cluster
	clusterBName := "clusterB"
	testScenario, neutronPublic, neutronPrivate, cleanup2 := runEndpointTest(
		t, clusterBName, false)
	defer cleanup2()
	// remove tempfile after test
	defer neutronPrivate.Close()
	defer neutronPublic.Close()

	// wait for proxy endpoints to update
	time.Sleep(2 * time.Second)

	// verify new proxies
	verifyProxyAndTestIt(t, testScenario, clusterBName, true)

	// verify existing proxies, make sure the proxy prefix is updated with cluster id
	verifyProxyAndTestIt(t, testScenario, clusterAName, true)
}
