package apisrv

import (
	"crypto/tls"
	"net"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
	"time"

	"github.com/Juniper/contrail/pkg/apisrv/baseapisrv"
	"github.com/labstack/echo"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

type staticProxyPlugin struct{}

func (staticProxyPlugin) RegisterHTTPAPI(r baseapisrv.HTTPRouter) error {
	for prefix, targetURLs := range viper.GetStringMapStringSlice("server.proxy") {
		if len(targetURLs) == 0 {
			return errors.Errorf("no target URLs provided for prefix %v", prefix)
		}

		// TODO(dfurman): proxy requests to all provided target URLs
		t, err := url.Parse(targetURLs[0])
		if err != nil {
			return errors.Wrapf(err, "bad proxy target URL: %s", targetURLs[0])
		}

		r.Group(
			prefix,
			removePathPrefixMiddleware(prefix),
			proxyMiddleware(t, viper.GetBool("server.proxy.insecure")),
		)
	}
	return nil
}

func (staticProxyPlugin) RegisterGRPCAPI(r baseapisrv.GRPCRouter) error {
	return nil
}

func removePathPrefixMiddleware(prefix string) baseapisrv.MiddlewareFunc {
	return func(next baseapisrv.HandlerFunc) baseapisrv.HandlerFunc {
		return func(c echo.Context) error {
			req := c.Request()
			req.URL.Path = strings.TrimPrefix(req.URL.Path, prefix)
			return next(c)
		}
	}
}

func proxyMiddleware(target *url.URL, insecure bool) baseapisrv.MiddlewareFunc {
	server := httputil.NewSingleHostReverseProxy(target)
	if target.Scheme == "https" {
		server.Transport = &http.Transport{
			Dial: (&net.Dialer{
				Timeout:   30 * time.Second,
				KeepAlive: 30 * time.Second,
			}).Dial,
			TLSClientConfig:     &tls.Config{InsecureSkipVerify: insecure},
			TLSHandshakeTimeout: 10 * time.Second,
		}
	}
	return func(next baseapisrv.HandlerFunc) baseapisrv.HandlerFunc {
		return func(c echo.Context) error {
			r := c.Request()
			w := c.Response()
			server.ServeHTTP(w, r)
			return nil
		}
	}
}
