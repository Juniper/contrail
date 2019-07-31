package proxy

import (
	"fmt"
	"net"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"

	"github.com/Juniper/contrail/pkg/apisrv/endpoint"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"

	cleanhttp "github.com/hashicorp/go-cleanhttp"
)

// Proxy constants.
const (
	UserAgentHeader = "User-Agent"
)

// NewReverseProxy returns a new ReverseProxy that routes URLs to the scheme, host, and base path
// provided in target. If the target's path is "/base" and the incoming request was for "/dir",
// the target request will be for /base/dir.
func NewReverseProxy(targetURL *url.URL) *httputil.ReverseProxy {
	fmt.Println("Hoge NewReverseProxy")
	return &httputil.ReverseProxy{
		Director: func(req *http.Request) {
			fmt.Println("Hoge rp.director")
			fmt.Printf("req.URL: %v %#v\n", req.URL, req.URL)
			fmt.Printf("targetURL: %v %#v\n", targetURL, targetURL)
			fmt.Printf("req.Header: %v\n", req.Header)

			req.URL.Scheme = targetURL.Scheme
			req.URL.Host = targetURL.Host
			req.URL.Path = singleJoiningSlash(targetURL.Path, req.URL.Path)
			req.URL.RawQuery = mergeQueries(req.URL.RawQuery, targetURL.RawQuery)
			req.Header = withNoDefaultUserAgent(req.Header)

			logrus.WithField("url", req.URL).Debug("Reverse proxy: proxying request")
		},
		Transport: cleanhttp.DefaultPooledTransport(),
	}
}

func singleJoiningSlash(a, b string) string {
	aSlash := strings.HasSuffix(a, "/")
	bSlash := strings.HasPrefix(b, "/")
	switch {
	case aSlash && bSlash:
		return a + b[1:]
	case !aSlash && !bSlash:
		return a + "/" + b
	}
	return a + b
}

func mergeQueries(requestQuery, targetQuery string) string {
	if targetQuery == "" || requestQuery == "" {
		return targetQuery + requestQuery
	}
	return fmt.Sprintf("%s&%s", targetQuery, requestQuery)
}

func withNoDefaultUserAgent(h http.Header) http.Header {
	if _, ok := h[UserAgentHeader]; !ok {
		// explicitly disable User-Agent so it's not set to default value
		h.Set(UserAgentHeader, "")
	}
	return h
}

// NewReverseProxy2 creates reverse proxy capable of handling multiple target hosts.
func NewReverseProxy2(targets []*endpoint.Endpoint, targetURL *url.URL) *httputil.ReverseProxy {
	fmt.Println("Hoge NewReverseProxy")
	return &httputil.ReverseProxy{
		Director: func(req *http.Request) {
			fmt.Println("Hoge rp.director")
			// TODO(Daniel): move that logic to Transport.Dial, because here it is called only once at initialization
			req.URL.Scheme = targetURL.Scheme
			req.URL.Host = targetURL.Host
			req.URL.Path = targetURL.Path
			req.URL.RawQuery = mergeQueries(req.URL.RawQuery, targetURL.RawQuery)
		},
		Transport: &http.Transport{
			Proxy: http.ProxyFromEnvironment,
			Dial: func(network, addr string) (net.Conn, error) {
				fmt.Println("Hoge constructing rp.transport.roundRobinDial")
				return roundRobinDial(network, targets)
			},
		},
	}
}

// roundRobinDial tries connecting to the endpoints in round robin fashion in case of connection failure.
func roundRobinDial(network string, targets []*endpoint.Endpoint) (net.Conn, error) {
	fmt.Println("Hoge rp.transport.roundRobinDial")
	for _, target := range targets {
		u, err := url.Parse(target.URL)
		if err != nil {
			logrus.WithError(err).WithField(
				"target", target.URL,
			).Debug("Failed to parse target - ignoring and trying next one")
		}

		conn, err := net.Dial(network, u.Host)
		if err == nil {
			return conn, nil
		}
	}

	fmt.Println("Hoge rp.transport.roundRobinDial returns no targets available error")
	return nil, errors.New("no targets available")
}
