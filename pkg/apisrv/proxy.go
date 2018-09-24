package apisrv

import (
	"context"
	"crypto/tls"
	"net"
	"net/http"
	"net/http/httputil"
	"net/url"
	"reflect"
	"strings"
	"sync"
	"time"

	"github.com/labstack/echo"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	apicommon "github.com/Juniper/contrail/pkg/apisrv/common"
	"github.com/Juniper/contrail/pkg/db/cache"
	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/services"
)

const (
	// DefaultDynamicProxyPath default value for server.dynamic_proxy.path
	DefaultDynamicProxyPath = "proxy"
	public                  = apicommon.Public
	private                 = apicommon.Private
	pathSep                 = "/"
)

type proxyService struct {
	group         string
	echoServer    *echo.Echo
	dbService     services.Service
	EndpointStore *apicommon.EndpointStore
	Cache         *cache.DB
	// context to stop servicing proxy endpoints
	serviceContext     context.Context
	stopServiceContext context.CancelFunc
	serviceWaitGroup   *sync.WaitGroup
}

func newProxyService(e *echo.Echo, endpointStore *apicommon.EndpointStore,
	dbService services.Service, cache *cache.DB) *proxyService {
	group := viper.GetString("server.dynamic_proxy.path")
	if group == "" {
		group = DefaultDynamicProxyPath
	}
	p := &proxyService{
		group:         group,
		dbService:     dbService,
		echoServer:    e,
		EndpointStore: endpointStore,
		Cache:         cache,
	}
	return p
}

func (p *proxyService) getProxyPrefixFromURL(urlPath string, scope string) (proxyPrefix string) {
	paths := strings.Split(urlPath, pathSep)
	prefixes := make([]string, 4)
	copy(prefixes, paths[:4])
	prefixes = append(prefixes, scope)
	return strings.Join(prefixes, pathSep)
}

func (p *proxyService) getReverseProxy(urlPath string) (prefix string, server *httputil.ReverseProxy) {
	var scope string
	if strings.Contains(urlPath, private) {
		scope = private
	} else {
		scope = public
	}
	proxyPrefix := p.getProxyPrefixFromURL(urlPath, scope)
	proxyEndpoint := p.EndpointStore.Read(proxyPrefix)
	if proxyEndpoint == nil {
		log.WithField("proxy-prefix", proxyPrefix).Info("Endpoint targets not found for given proxy prefix")
		return strings.TrimSuffix(proxyPrefix, public), nil
	}
	target := proxyEndpoint.Next(scope)
	if target == "" {
		return strings.TrimSuffix(proxyPrefix, public), nil
	}
	insecure := true //TODO:(ijohnson) add insecure to endpoint schema

	u, err := url.Parse(target)
	if err != nil {
		log.WithError(err).WithField("target", target).Info("Failed to parse target - ignoring")
	}

	server = httputil.NewSingleHostReverseProxy(u)
	if u.Scheme == "https" {
		server.Transport = &http.Transport{
			Dial: (&net.Dialer{
				Timeout:   30 * time.Second,
				KeepAlive: 30 * time.Second,
			}).Dial,
			TLSClientConfig:     &tls.Config{InsecureSkipVerify: insecure},
			TLSHandshakeTimeout: 10 * time.Second,
		}
	}
	return strings.TrimSuffix(proxyPrefix, public), server
}

func (p *proxyService) dynamicProxyMiddleware() func(next echo.HandlerFunc) echo.HandlerFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			r := c.Request()
			prefix, server := p.getReverseProxy(r.URL.Path)
			if server == nil {
				return echo.NewHTTPError(http.StatusInternalServerError,
					"Proxy endpoint not found in endpoint store")
			}
			r.URL.Path = strings.TrimPrefix(r.URL.Path, prefix)
			w := c.Response()
			server.ServeHTTP(w, r)
			return nil
		}
	}
}

func (p *proxyService) deleteProxyEndpoint(endpointID string) {
	p.EndpointStore.Data.Range(func(prefix, proxy interface{}) bool {
		s, ok := proxy.(*apicommon.TargetStore)
		if !ok {
			log.Errorf("Unable to Read cluster(%s)'s proxy data from in-memory store",
				prefix)
			return true
		}
		s.Data.Range(func(id, endpoint interface{}) bool {
			_, ok := endpoint.(*models.Endpoint)
			if !ok {
				log.Errorf("Unable to Read endpoint(%s) data from in-memory store",
					id)
				return true
			}
			ids, ok := id.(string)
			if !ok {
				log.Errorf("Unable to convert id %v to string when looking EndpointStore", id)
				return true
			}
			if ids == endpointID {
				s.Remove(ids)
				log.Debugf("deleting dynamic proxy endpoint for id: %s", ids)
			}
			return true
		})
		return true
	})
}

func (p *proxyService) getProxyPrefix(endpoint *models.Endpoint, scope string) (proxyPrefix string) {
	prefix := endpoint.Prefix
	// TODO(ijohnson) remove using DisplayName as prefix
	// once UI takes prefix as input.
	if prefix == "" {
		prefix = endpoint.DisplayName
	}

	if endpoint.ParentUUID == "" {
		log.Errorf("Parent uuid missing for endpoint %s(%s)", prefix, endpoint.UUID)
	}
	prefixes := []string{"", p.group, endpoint.ParentUUID, prefix, scope}
	return strings.Join(prefixes, pathSep)
}

func (p *proxyService) initProxyTargetStore(endpoint *models.Endpoint) {
	if endpoint.PublicURL != "" {
		proxyPrefix := p.getProxyPrefix(endpoint, public)
		targetStore := p.EndpointStore.Read(proxyPrefix)
		if targetStore == nil {
			p.EndpointStore.Write(proxyPrefix, apicommon.MakeTargetStore())
		}
		privateProxyPrefix := p.getProxyPrefix(endpoint, private)
		privateTargetStore := p.EndpointStore.Read(privateProxyPrefix)
		if privateTargetStore == nil {
			p.EndpointStore.Write(privateProxyPrefix, apicommon.MakeTargetStore())
		}
	}
}

func (p *proxyService) manageProxyEndpoint(endpoint *models.Endpoint, scope string) {
	proxyPrefix := p.getProxyPrefix(endpoint, scope)
	s := p.EndpointStore.Read(proxyPrefix)
	if s == nil {
		log.Errorf("endpoint store for %s is not found in-memory store",
			proxyPrefix)
	}
	e := s.Read(endpoint.UUID)
	if !reflect.DeepEqual(e, endpoint) {
		// proxy endpoint not in memory store or
		// proxy endpoint updated
		s.Write(endpoint.UUID, endpoint)
	}
}

func (p *proxyService) syncProxyEndpoint(endpoint *models.Endpoint) {
	// create/update proxy middleware
	p.initProxyTargetStore(endpoint)
	if endpoint.PublicURL != "" {
		p.manageProxyEndpoint(endpoint, public)
		p.manageProxyEndpoint(endpoint, private)
	}
}

func (p *proxyService) process(e *services.Event) {
	// process the received event.
	switch event := e.Request.(type) {
	case *services.Event_CreateEndpointRequest:
		p.syncProxyEndpoint(event.CreateEndpointRequest.Endpoint)
	case *services.Event_UpdateEndpointRequest:
		p.syncProxyEndpoint(event.UpdateEndpointRequest.Endpoint)
	case *services.Event_DeleteEndpointRequest:
		p.deleteProxyEndpoint(event.DeleteEndpointRequest.ID)
	}
	log.Debugf("proxy: Processed event %v", e)
}

func (p *proxyService) serve() {
	// add dynamic proxy middleware
	g := p.echoServer.Group(p.group)
	g.Use(p.dynamicProxyMiddleware())
	p.serviceContext, p.stopServiceContext = context.WithCancel(context.Background())
	watcher, _ := p.Cache.AddWatcher(p.serviceContext, 0)
	p.serviceWaitGroup = &sync.WaitGroup{}
	p.serviceWaitGroup.Add(1)
	go func() {
		defer p.serviceWaitGroup.Done()
		for {
			select {
			case <-p.serviceContext.Done():
				log.Info("Stopping dynamic proxy server")
				return
			case e := <-watcher.Chan():
				p.process(e)
			}
		}
	}()
}

func (p *proxyService) stop() {
	// stop proxy server poll
	p.stopServiceContext()
	// wait for the proxy server poll to complete
	p.serviceWaitGroup.Wait()
}
