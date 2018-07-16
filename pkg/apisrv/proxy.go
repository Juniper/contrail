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
	"github.com/Juniper/contrail/pkg/common"
	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/services"
)

const (
	public  = apicommon.Public
	private = apicommon.Private
	pathSep = "/"
	limit   = 100
)

type proxyService struct {
	group         string
	echoServer    *echo.Echo
	dbService     services.Service
	EndpointStore *apicommon.EndpointStore
	// context to stop servicing proxy endpoints
	serviceContext     context.Context
	stopServiceContext context.CancelFunc
	serviceWaitGroup   *sync.WaitGroup
	forceUpdateChan    chan chan struct{}
}

func newProxyService(e *echo.Echo, endpointStore *apicommon.EndpointStore,
	dbService services.Service) *proxyService {
	group := viper.GetString("server.dynamic_proxy_path")
	if group == "" {
		group = "proxy"
	}
	p := &proxyService{
		group:         group,
		dbService:     dbService,
		echoServer:    e,
		EndpointStore: endpointStore,
	}
	return p
}

func (p *proxyService) readEndpoints() (map[string]*models.Endpoint, error) {
	endpoints := make(map[string]*models.Endpoint)
	ctx := common.NoAuth(context.Background())
	spec := services.ListSpec{Limit: limit}
	for {
		request := &services.ListEndpointRequest{Spec: &spec}
		response, err := p.dbService.ListEndpoint(ctx, request)
		if err != nil {
			return nil, err
		}
		for _, e := range response.Endpoints {
			endpoints[e.UUID] = e
		}
		if len(response.Endpoints) != limit {
			// less than 100 records present in DB
			break
		}
		// more than 100 records present in DB, continue to read
		offset := int64(len(endpoints) + 1)
		spec.Offset = offset
	}
	return endpoints, nil
}

func (p *proxyService) getProxyPrefixFromURL(urlPath string, scope string) (proxyPrefix string) {
	paths := strings.Split(urlPath, pathSep)
	prefixes := make([]string, 4)
	copy(prefixes, paths[:4])
	prefixes = append(prefixes, scope)
	return strings.Join(prefixes, pathSep)
}

func (p *proxyService) getReverseProxy(urlPath string) (
	prefix string, server *httputil.ReverseProxy) {
	var scope string
	if strings.Contains(urlPath, private) {
		scope = private
	} else {
		scope = public
	}
	proxyPrefix := p.getProxyPrefixFromURL(urlPath, scope)
	proxyEndpoint := p.EndpointStore.Read(proxyPrefix)
	if proxyEndpoint == nil {
		log.Errorf("Endpoint targets not found for %s", proxyPrefix)
		return strings.TrimSuffix(proxyPrefix, public), nil
	}
	target := proxyEndpoint.Next(scope)
	if target == "" {
		return strings.TrimSuffix(proxyPrefix, public), nil
	}
	insecure := true          //TODO:(ijohnson) add insecure to endpoint schema
	u, _ := url.Parse(target) // nolint
	server = httputil.NewSingleHostReverseProxy(u)
	if u.Scheme == "https" {
		server.Transport = &http.Transport{
			Dial: (&net.Dialer{
				Timeout:   30 * time.Second,
				KeepAlive: 30 * time.Second,
			}).Dial,
			TLSClientConfig:     &tls.Config{InsecureSkipVerify: insecure}, // nolint
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

func (p *proxyService) checkDeleted(endpoints map[string]*models.Endpoint) {
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
			ids := id.(string)
			_, ok = endpoints[ids]
			if !ok {
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

func (p *proxyService) syncProxyEndpoints(endpoints map[string]*models.Endpoint) {
	// delete stale proxy endpoints in-memory
	p.checkDeleted(endpoints)
	// create/update proxy middleware
	for _, endpoint := range endpoints {
		p.initProxyTargetStore(endpoint)
		if endpoint.PublicURL != "" {
			p.manageProxyEndpoint(endpoint, public)
			p.manageProxyEndpoint(endpoint, private)
		}
	}
}

func (p *proxyService) forceUpdate() {
	wait := make(chan struct{})
	p.forceUpdateChan <- wait
	<-wait
}

func (p *proxyService) serve() {
	// add dynamic proxy middleware
	g := p.echoServer.Group(p.group)
	g.Use(p.dynamicProxyMiddleware())
	p.forceUpdateChan = make(chan chan struct{})
	p.serviceContext, p.stopServiceContext = context.WithCancel(context.Background())
	p.serviceWaitGroup = &sync.WaitGroup{}
	p.serviceWaitGroup.Add(1)
	go func() {
		defer p.serviceWaitGroup.Done()
		ticker := time.NewTicker(2 * time.Second)
		defer ticker.Stop()
		for {
			select {
			case <-p.serviceContext.Done():
				log.Info("stopping dynamic proxy server")
				return
			case wait := <-p.forceUpdateChan:
				p.updateEndpoints()
				close(wait)
			case <-ticker.C:
				p.updateEndpoints()
			}
		}
	}()
}

func (p *proxyService) updateEndpoints() {
	endpoints, err := p.readEndpoints()
	if err != nil {
		log.WithError(err).Error("Endpoints read failed")
		return
	}
	p.syncProxyEndpoints(endpoints)
}

func (p *proxyService) stop() {
	// stop proxy server poll
	p.stopServiceContext()
	// wait for the proxy server poll to complete
	p.serviceWaitGroup.Wait()
}
