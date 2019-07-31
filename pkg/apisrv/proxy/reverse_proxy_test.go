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

func TestReverseProxy(t *testing.T) {
	require.NoError(t, logutil.Configure(logrus.DebugLevel.String())) // for debugging

	for _, tt := range []struct {
		name      string
		userAgent string
	}{
		{
			name:      "do request",
			userAgent: "test-client",
		},
		{
			name:      "do request without User-Agent header",
			userAgent: "",
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			bs := backendServer(t, tt.userAgent)
			defer bs.Close()

			rp := httptest.NewServer(proxy.NewReverseProxy(targetURL(t, bs.URL)))
			defer rp.Close()

			response, err := http.DefaultClient.Do(request(t, rp.URL, tt.userAgent))

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
		assert.Equal(t, requestMethod, r.Method)
		assert.Equal(t, "", r.URL.Scheme)
		assert.Equal(t, "", r.URL.Opaque)
		assert.Equal(t, "", r.URL.Host)
		assert.Equal(t, targetPath+resourcePath, r.URL.Path)
		assert.Equal(t, fmt.Sprintf("%s&%s", targetQuery, requestQuery), r.URL.RawQuery)
		assert.Equal(t, "", r.URL.Fragment)
		assert.Equal(t, expectedUserAgent, r.UserAgent())

		_, pErr := fmt.Fprint(w, message)
		assert.NoError(t, pErr)
	}))
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
