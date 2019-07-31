package proxy

import (
	"crypto/tls"
	"fmt"
	"net"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
	"time"

	"github.com/Juniper/contrail/pkg/apisrv/endpoint"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

// NewReverseProxy returns a new ReverseProxy that routes URLs to the scheme, host, and base path
// provided in target. If the target's path is "/base" and the incoming request was for "/dir",
// the target request will be for /base/dir.
func NewReverseProxy(targetURL *url.URL) *httputil.ReverseProxy {
	logrus.Warn("Hoge NewReverseProxy")
	return &httputil.ReverseProxy{
		Director: func(req *http.Request) {
			logrus.Warn("Hoge rp.director")
			req.URL.Scheme = targetURL.Scheme
			req.URL.Host = targetURL.Host
			req.URL.Path = singleJoiningSlash(targetURL.Path, req.URL.Path)

			if targetURL.RawQuery == "" || req.URL.RawQuery == "" {
				req.URL.RawQuery = targetURL.RawQuery + req.URL.RawQuery
			} else {
				req.URL.RawQuery = targetURL.RawQuery + "&" + req.URL.RawQuery
			}

			if _, ok := req.Header["User-Agent"]; !ok {
				// explicitly disable User-Agent so it's not set to default value
				req.Header.Set("User-Agent", "")
			}
		},
		Transport: transport(targetURL),
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

func transport(targetURL *url.URL) *http.Transport {
	// TODO: use cleanhttp package
	t := &http.Transport{
		Proxy: http.ProxyFromEnvironment,
		DialContext: (&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
		}).DialContext,
		MaxIdleConns:          100,
		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
	}

	if targetURL.Scheme == "https" {
		t.TLSClientConfig = &tls.Config{InsecureSkipVerify: true} // TODO: add insecure to endpoint schema
	}
	return t
}

// NewReverseProxy create reverse proxy capable of handling multiple target hosts.
func NewReverseProxy2(targets []*endpoint.Endpoint, targetURL *url.URL) *httputil.ReverseProxy {
	logrus.Warn("Hoge NewReverseProxy")
	return &httputil.ReverseProxy{
		Director: func(req *http.Request) {
			logrus.Warn("Hoge rp.director")
			// TODO(Daniel): move that logic to Transport.Dial, because here it is called only once at initialization
			req.URL.Scheme = targetURL.Scheme
			req.URL.Host = targetURL.Host
			req.URL.Path = targetURL.Path
			req.URL.RawQuery = combineQueries(req.URL.RawQuery, targetURL.RawQuery)
		},
		Transport: &http.Transport{
			Proxy: http.ProxyFromEnvironment,
			Dial: func(network, addr string) (net.Conn, error) {
				logrus.Warn("Hoge constructing rp.transport.roundRobinDial")
				return roundRobinDial(network, targets)
			},
			TLSClientConfig:     &tls.Config{InsecureSkipVerify: true}, // TODO: add insecure to endpoint schema
			TLSHandshakeTimeout: 10 * time.Second,
		},
	}
}

func combineQueries(requestQuery, targetQuery string) string {
	if targetQuery == "" || requestQuery == "" {
		return targetQuery + requestQuery
	} else {
		return fmt.Sprintf("%s&%s", targetQuery, requestQuery)
	}
}

// roundRobinDial tries connecting to the endpoints in round robin fashion in case of connection failure.
func roundRobinDial(network string, targets []*endpoint.Endpoint) (net.Conn, error) {
	logrus.Warn("Hoge rp.transport.roundRobinDial")
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

	logrus.Warn("Hoge rp.transport.roundRobinDial returns no targets available error")
	return nil, errors.New("no targets available")
}
