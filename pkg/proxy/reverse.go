package proxy

import (
	"context"
	"crypto/tls"
	"fmt"
	"net"
	"net/http"
	"net/http/httputil"
	"net/url"
	"path"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"

	cleanhttp "github.com/hashicorp/go-cleanhttp"
)

// Proxy constants.
const (
	skipServerCertificateVerification = true // TODO: add "insecure" field to endpoint schema
	userAgentHeader                   = "User-Agent"
	xServiceTokenHeader               = "X-Service-Token"
)

type serviceTokener interface {
	// ServiceToken returns a service token that can be added to openstack requests.
	// It returns an empty string if no service token should be added.
	// TODO Add an error return value.
	ServiceToken() string
}

// NullServiceTokener makes ReverseProxy not add any service token to requests.
type NullServiceTokener struct{}

// ServiceToken signifies that no service token should be added to the request.
func (NullServiceTokener) ServiceToken() string {
	return ""
}

// NewReverseProxy returns a new ReverseProxy that routes URLs to the scheme, host, and base path
// provided in target. If the target's path is "/base" and the incoming request was for "/dir",
// the target request will be for /base/dir.
// If addServiceToken is true, the proxy will add the server's service token to all requests.
func NewReverseProxy(rawTargetURLs []string, serviceTokener serviceTokener) (*httputil.ReverseProxy, error) {
	targetURLs := parseTargetURLs(rawTargetURLs)
	if len(targetURLs) == 0 {
		return nil, errors.New("no valid target URLs given")
	}

	return &httputil.ReverseProxy{
		Director:  director(targetURLs[0], serviceTokener),
		Transport: transport(targetURLs),
	}, nil
}

func parseTargetURLs(rawTargetURLs []string) []*url.URL {
	var targetURLs []*url.URL
	for _, t := range rawTargetURLs {
		tURL, err := url.Parse(t)
		if err != nil {
			logrus.WithError(err).WithField("target-url", t).Error("Failed to parse target URL - ignoring")
		} else {
			targetURLs = append(targetURLs, tURL)
		}
	}
	return targetURLs
}

func director(firstTargetURL *url.URL, st serviceTokener) func(r *http.Request) {
	return func(r *http.Request) {
		r.URL.Scheme = firstTargetURL.Scheme
		r.URL.Host = firstTargetURL.Host // request host might be reassigned in ReverseProxy.Transport.DialContext.
		r.URL.Path = path.Join("/", firstTargetURL.Path, r.URL.Path)
		r.URL.RawQuery = mergeQueries(r.URL.RawQuery, firstTargetURL.RawQuery)
		r.Header = withNoDefaultUserAgent(r.Header)
		r.Header = withServiceToken(r.Header, st)

		logrus.WithField("url", r.URL).Debug("Reverse proxy: proxying request")
	}
}

func mergeQueries(requestQuery, targetQuery string) string {
	if targetQuery == "" || requestQuery == "" {
		return targetQuery + requestQuery
	}
	return fmt.Sprintf("%s&%s", targetQuery, requestQuery)
}

func withServiceToken(h http.Header, st serviceTokener) http.Header {
	if token := st.ServiceToken(); token != "" {
		h.Set(xServiceTokenHeader, token)
	}
	return h
}

func withNoDefaultUserAgent(h http.Header) http.Header {
	if _, ok := h[userAgentHeader]; !ok {
		// explicitly disable User-Agent so it's not set to default value
		h.Set(userAgentHeader, "")
	}
	return h
}

func transport(targetURLs []*url.URL) *http.Transport {
	t := cleanhttp.DefaultPooledTransport()
	t.TLSClientConfig = &tls.Config{InsecureSkipVerify: skipServerCertificateVerification}
	t.DialContext = func(ctx context.Context, network, _ string) (net.Conn, error) {
		return roundRobinDial(ctx, network, targetURLs)
	}
	return t
}

func roundRobinDial(ctx context.Context, network string, targetURLs []*url.URL) (net.Conn, error) {
	var errs []error
	var d net.Dialer
	for _, targetURL := range targetURLs {
		c, err := d.DialContext(ctx, network, targetURL.Host)
		if err == nil {
			return c, nil
		}

		errs = append(errs, err)
		logrus.WithError(err).WithFields(logrus.Fields{
			"network":    network,
			"target-url": targetURL,
		}).Debug("Reverse proxy: failed to dial to target - trying next target")
	}

	return nil, errors.Errorf("proxy: failed to dial to all target URLs (%v): %v", targetURLs, errs)
}
