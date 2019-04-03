package common

import (
	"crypto/tls"
	"errors"
	"net"
	"net/http"
	"net/http/httputil"
	"net/url"
	"time"

	"github.com/sirupsen/logrus"
)

// roundRobin tries connecting to the endpoints
// in round robin fashion in case of connection failure
func roundRobin(network string, targets []*Endpoint) (net.Conn, error) {
	for _, target := range targets {
		u, err := url.Parse(target.URL)
		if err != nil {
			logrus.WithError(err).WithField("target", target.URL).Info("Failed to parse target - ignoring")
		}

		conn, err := net.Dial(network, u.Host)
		if err == nil {
			return conn, nil
		}
	}
	return nil, errors.New("No targets available")
}

// that will randomly select a host from the passed
func NewReverseProxyPool(targets []*Endpoint) *httputil.ReverseProxy {
	u, err := url.Parse(targets[0].URL)
	if err != nil {
		logrus.WithError(err).WithField("target", targets[0].URL).Info("Failed to parse target - ignoring")
	}
	director := func(req *http.Request) {
		req.URL.Scheme = u.Scheme
		req.URL.Host = u.Host
	}
	insecure := true //TODO:(ijohnson) add insecure to endpoint schema
	transport := &http.Transport{
		Proxy: http.ProxyFromEnvironment,
		Dial: func(network, addr string) (net.Conn, error) {
			return roundRobin(network, targets)
		},
		TLSClientConfig:     &tls.Config{InsecureSkipVerify: insecure},
		TLSHandshakeTimeout: 10 * time.Second,
	}
	return &httputil.ReverseProxy{
		Director:  director,
		Transport: transport,
	}
}
