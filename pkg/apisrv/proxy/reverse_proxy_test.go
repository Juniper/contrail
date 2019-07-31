package proxy_test

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/Juniper/contrail/pkg/logutil"
	"github.com/sirupsen/logrus"

	"github.com/Juniper/contrail/pkg/apisrv/proxy"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	message       = "msg"
	requestMethod = "GET"
	resourcePath  = "/resources"
	requestQuery  = "query1=one&query2=two"
	targetPath    = "/target"
	targetQuery   = "targetQuery=foo"
)

func TestNewReverseProxyFailsWithNoTargetURLs(t *testing.T) {
	rp, err := proxy.NewReverseProxy([]*url.URL{})

	assert.Error(t, err)
	assert.Nil(t, rp)
}

func TestReverseProxy(t *testing.T) {
	require.NoError(t, logutil.Configure(logrus.DebugLevel.String())) // for debugging

	for _, tt := range []struct {
		name      string
		userAgent string
	}{
		{
			name: "proxies request",
		},
		{
			name:      "proxies request with User-Agent header",
			userAgent: "test-client",
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			bs := backendServer(t, tt.userAgent)
			defer bs.Close()

			rp, err := proxy.NewReverseProxy([]*url.URL{targetURL(t, bs.URL)})
			require.NoError(t, err)
			rps := httptest.NewServer(rp)
			defer rps.Close()

			response, err := http.DefaultClient.Do(request(t, rps.URL, tt.userAgent))

			assert.NoError(t, err)
			assert.Equal(t, http.StatusOK, response.StatusCode)

			b, err := ioutil.ReadAll(response.Body)
			assert.NoError(t, err)
			assert.Equal(t, message, string(b))
		})
	}
}

func backendServer(t *testing.T, expectedUserAgent string) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		checkRequest(t, r, expectedUserAgent)

		_, pErr := fmt.Fprint(w, message)
		if pErr != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
	}))
}

func checkRequest(t *testing.T, r *http.Request, expectedUserAgent string) {
	assert.Equal(t, requestMethod, r.Method)
	assert.Equal(t, "", r.URL.Scheme)
	assert.Equal(t, "", r.URL.Opaque)
	assert.Equal(t, "", r.URL.Host)
	assert.Equal(t, targetPath+resourcePath, r.URL.Path)
	assert.Equal(t, fmt.Sprintf("%s&%s", targetQuery, requestQuery), r.URL.RawQuery)
	assert.Equal(t, "", r.URL.Fragment)
	assert.Equal(t, expectedUserAgent, r.UserAgent())
}

func targetURL(t *testing.T, bsURL string) *url.URL {
	tURL, err := url.Parse(bsURL)
	require.NoError(t, err)

	tURL.Path = targetPath
	tURL.RawQuery = targetQuery
	return tURL
}

func request(t *testing.T, rpURL string, userAgent string) *http.Request {
	requestURL, err := url.Parse(rpURL)
	require.NoError(t, err)
	requestURL.Path = resourcePath
	requestURL.RawQuery = requestQuery
	requestURL.Fragment = "test-fragment"

	r, err := http.NewRequest(requestMethod, requestURL.String(), nil)
	require.NoError(t, err)

	r.Header.Set(proxy.UserAgentHeader, userAgent)

	return r
}
