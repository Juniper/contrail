package testutil

import (
	"crypto/tls"
	"net/http"
	"net/http/httptest"

	"github.com/golang/mock/gomock"
)

// NewTestHTTPServer starts and returns new test HTTP 2 Server.
func NewTestHTTPServer(h http.Handler) *httptest.Server {
	s := httptest.NewUnstartedServer(h)
	s.TLS = new(tls.Config)
	s.TLS.NextProtos = append(s.TLS.NextProtos, "h2")
	s.StartTLS()
	return s
}

// NotNil matches any non-nil value.
func NotNil() gomock.Matcher {
	return gomock.Not(gomock.Nil())
}
