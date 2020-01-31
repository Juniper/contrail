package proxy

import (
	"crypto/tls"
	"net"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
	"time"

	"github.com/Juniper/asf/pkg/apisrv/baseapisrv"
	"github.com/labstack/echo"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

// StaticProxyPlugin provides statically configured endpoints that proxy requests.
type StaticProxyPlugin struct {
	proxies []staticProxy
}

// NewStaticProxyPluginByViper creates a StaticProxyPlugin based on global Viper configuration.
func NewStaticProxyPluginByViper() (p StaticProxyPlugin, err error) {
	for prefix, rawTargetURLs := range viper.GetStringMapStringSlice("server.proxy") {
		targetURLs, err := parseTargetURLs(rawTargetURLs)
		if err != nil {
			return StaticProxyPlugin{}, errors.Wrapf(err, "invalid target URLs for prefix %v", prefix)
		}
		p.proxies = append(p.proxies, newStaticProxy(
			prefix,
			targetURLs,
			viper.GetBool("server.proxy.insecure"),
		))
	}
	return p, nil
}

func parseTargetURLs(rawTargetURLs []string) (targetURLs []*url.URL, err error) {
	if len(rawTargetURLs) == 0 {
		return nil, errors.New("no target URLs provided")
	}
	for _, rawURL := range rawTargetURLs {
		url, err := url.Parse(rawURL)
		if err != nil {
			return nil, errors.Wrapf(err, "bad proxy target URL: %s", rawURL)
		}
		targetURLs = append(targetURLs, url)
	}
	return targetURLs, nil
}

// RegisterHTTPAPI registers the proxy endpoints.
func (p StaticProxyPlugin) RegisterHTTPAPI(r baseapisrv.HTTPRouter) {
	for _, pr := range p.proxies {
		r.Group(pr.prefix, baseapisrv.WithMiddleware(pr.middleware))
	}
}

// RegisterGRPCAPI does nothing.
func (StaticProxyPlugin) RegisterGRPCAPI(r baseapisrv.GRPCRouter) {
}

type staticProxy struct {
	prefix string
	server *httputil.ReverseProxy
}

func newStaticProxy(prefix string, targetURLs []*url.URL, insecure bool) staticProxy {
	// TODO(dfurman): proxy requests to all provided target URLs
	target := targetURLs[0]

	pr := staticProxy{
		prefix: prefix,
		server: httputil.NewSingleHostReverseProxy(target),
	}
	if target.Scheme == "https" {
		pr.server.Transport = &http.Transport{
			Dial: (&net.Dialer{
				Timeout:   30 * time.Second,
				KeepAlive: 30 * time.Second,
			}).Dial,
			TLSClientConfig:     &tls.Config{InsecureSkipVerify: insecure},
			TLSHandshakeTimeout: 10 * time.Second,
		}
	}
	return pr
}

func (pr staticProxy) middleware(next baseapisrv.HandlerFunc) baseapisrv.HandlerFunc {
	return func(c echo.Context) error {
		r := c.Request()
		w := c.Response()
		r.URL.Path = strings.TrimPrefix(r.URL.Path, pr.prefix)
		pr.server.ServeHTTP(w, r)
		return nil
	}
}
