package proxy

import (
	"crypto/tls"
	"fmt"
	"net"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"

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
func NewReverseProxy(targetURLs []*url.URL) *httputil.ReverseProxy {
	fmt.Println("NewReverseProxy")
	return &httputil.ReverseProxy{
		Director:  director(targetURLs),
		Transport: transport(targetURLs),
	}
}

func director(targetURLs []*url.URL) func(r *http.Request) {
	targetURL := targetURLs[0] // TODO

	return func(r *http.Request) {
		fmt.Println("rp.director")
		fmt.Printf("r.URL: %v %#v\n", r.URL, r.URL)
		fmt.Printf("targetURL: %v %#v\n", targetURL, targetURL)
		fmt.Printf("r.Header: %v\n", r.Header)

		r.URL.Scheme = targetURL.Scheme
		r.URL.Host = targetURL.Host
		r.URL.Path = singleJoiningSlash(targetURL.Path, r.URL.Path)
		r.URL.RawQuery = mergeQueries(r.URL.RawQuery, targetURL.RawQuery)
		r.Header = withNoDefaultUserAgent(r.Header)

		logrus.WithField("url", r.URL).Debug("Reverse proxy: proxying request")
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

func transport(targetURLs []*url.URL) *http.Transport {
	t := cleanhttp.DefaultPooledTransport()
	t.TLSClientConfig = &tls.Config{InsecureSkipVerify: true} // TODO: add "insecure" field to endpoint schema
	t.Dial = func(network, addr string) (net.Conn, error) {
		return roundRobinDial(network, targetURLs)
	}
	return t
}

func roundRobinDial(network string, targetURLs []*url.URL) (net.Conn, error) {
	fmt.Printf("rp.transport.Dial; network %v, targetURLs %v\n", network, targetURLs)

	var errs []error
	for _, targetURL := range targetURLs {
		c, err := net.Dial(network, targetURL.Host)
		if err == nil {
			return c, nil
		}

		fmt.Printf("rp.transport.Dial: c %+v and err: %v\n", c, err)
		errs = append(errs, err)
	}

	return nil, errors.Errorf("proxy: failed to dial to all target URLs (%v): %v", targetURLs, errs)
}
