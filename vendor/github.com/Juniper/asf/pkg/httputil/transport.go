package httputil

import (
	"crypto/tls"
	"net"
	"net/http"
	"time"
)

func DefaultTransport(skipTLSCertVerification bool) *http.Transport {
	return &http.Transport{
		Proxy: http.ProxyFromEnvironment,
		DialContext: (&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
		}).DialContext,
		ForceAttemptHTTP2: true,
		TLSClientConfig:       &tls.Config{InsecureSkipVerify: skipTLSCertVerification},
		TLSHandshakeTimeout:   30 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
	}
}
