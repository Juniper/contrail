package proxy

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/http/httputil"
	"net/url"
	"testing"

	"github.com/Juniper/contrail/pkg/logutil"
	"github.com/sirupsen/logrus"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewReverseProxy(t *testing.T) {
	for _, tt := range []struct {
		name          string
		rawTargetURLs []string
		expectedRP    *httputil.ReverseProxy
		fails         bool
	}{
		{
			name:          "fails with no target URLs",
			rawTargetURLs: []string{},
			fails:         true,
		},
		{
			name:          "fails with invalid target URLs",
			rawTargetURLs: []string{"::invalid-url::"},
			fails:         true,
		},
		{
			name:          "succeeds with mixed target URLs",
			rawTargetURLs: []string{"::invalid-url::", "127.0.0.1"},
		},
		{
			name:          "succeeds with valid URLs",
			rawTargetURLs: []string{"127.0.0.1", "127.0.0.2"},
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			rp, err := NewReverseProxy(tt.rawTargetURLs)

			if tt.fails {
				assert.Error(t, err)
				assert.Nil(t, rp)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, rp)
			}
		})
	}
}

func TestReverseProxy(t *testing.T) {
	require.NoError(t, logutil.Configure(logrus.DebugLevel.String())) // for debugging
	const (
		message                = "msg"
		sampleRequestPath      = "/resources?query1=one&query2=two"
		sampleTargetPath       = "/target?targetQuery=foo"
		sampleSwiftRequestPath = "/v1/AUTH_23436c71c0d148019cc5d7b6e066d514/contrail_container/test34"
		sampleSwiftTargetPath  = ""
	)

	for _, tt := range []struct {
		name           string
		userAgent      string
		requestPath    string
		targetPath     string
		method         string
		receivedURL    *url.URL
		receivedHeader http.Header
	}{{
		name:   "simple proxy",
		method: http.MethodGet,
		receivedURL: &url.URL{
			Path: "/",
		},
	}, {
		name:        "proxy adds query",
		method:      http.MethodGet,
		requestPath: "/resource",
		targetPath:  "?param=true",
		receivedURL: &url.URL{
			Path:     "/resource",
			RawQuery: "param=true",
		},
	}, {
		name:        "proxy adds path segment",
		method:      http.MethodGet,
		requestPath: "?param=true",
		targetPath:  "/extra",
		receivedURL: &url.URL{
			Path:     "/extra",
			RawQuery: "param=true",
		},
	}, {
		name:        "proxies GET",
		method:      http.MethodGet,
		requestPath: sampleRequestPath,
		targetPath:  sampleTargetPath,
		receivedURL: &url.URL{
			Path:     "/target/resources",
			RawQuery: "targetQuery=foo&query1=one&query2=two",
		},
	}, {
		name:        "proxies POST",
		method:      http.MethodPost,
		requestPath: sampleRequestPath,
		targetPath:  sampleTargetPath,
		receivedURL: &url.URL{
			Path:     "/target/resources",
			RawQuery: "targetQuery=foo&query1=one&query2=two",
		},
	}, {
		name:        "proxies request with User-Agent header",
		method:      http.MethodGet,
		requestPath: sampleRequestPath,
		targetPath:  sampleTargetPath,
		userAgent:   "test-client",
		receivedURL: &url.URL{
			Path:     "/target/resources",
			RawQuery: "targetQuery=foo&query1=one&query2=two",
		},
	}, {
		name:        "adds X-Service-Token to a swift request",
		method:      http.MethodGet,
		requestPath: sampleSwiftRequestPath,
		targetPath:  sampleSwiftTargetPath,
		userAgent:   "test-client",
		receivedURL: &url.URL{
			Path:     "/v1/AUTH_23436c71c0d148019cc5d7b6e066d514/contrail_container/test34",
			RawQuery: "",
		},
		receivedHeader: http.Header{
			"X-Service-Token": []string{"4793e2594eb54e538fbe0a6265b2e0bf"},
		},
	}} {
		t.Run(tt.name, func(t *testing.T) {
			bs := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, tt.method, r.Method)
				assert.Equal(t, tt.receivedURL, r.URL)
				assert.Equal(t, tt.userAgent, r.UserAgent())
				for key, expectedValues := range tt.receivedHeader {
					values, ok := r.Header[key]
					assert.Truef(t, ok, "an %q header should be added by the request", key)
					assert.Equalf(t, expectedValues, values, "header %q should have values: %v", key, values)
				}

				_, pErr := fmt.Fprint(w, message)
				if pErr != nil {
					w.WriteHeader(http.StatusInternalServerError)
				}
			}))
			defer bs.Close()

			rp, err := NewReverseProxy([]string{bs.URL + tt.targetPath})
			require.NoError(t, err)
			rps := httptest.NewServer(rp)
			defer rps.Close()

			response, err := http.DefaultClient.Do(request(t, tt.method, rps.URL+tt.requestPath, tt.userAgent))

			assert.NoError(t, err)
			assert.Equal(t, http.StatusOK, response.StatusCode)

			b, err := ioutil.ReadAll(response.Body)
			assert.NoError(t, err)
			assert.Equal(t, message, string(b))
		})
	}
}

func request(t *testing.T, method, requestURL, userAgent string) *http.Request {
	r, err := http.NewRequest(method, requestURL, nil)
	require.NoError(t, err)

	r.Header.Set(UserAgentHeader, userAgent)

	return r
}
