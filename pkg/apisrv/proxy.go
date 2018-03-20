package apisrv

import (
	"context"
	"crypto/tls"
	"net"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
	"sync"
	"time"

	"github.com/labstack/echo"

	"github.com/Juniper/contrail/pkg/common"
	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/serviceif"

	log "github.com/sirupsen/logrus"
)

const (
	prefix  = "proxy"
	public  = "public"
	private = "private"
	pathSep = "/"
)

// ProxyServer holds info about the proxy server and endpoint store
type ProxyServer struct {
	echoServer   *echo.Echo
	dbService    serviceif.Service
	proxyStore   *proxyStore
	multiCluster bool
	// context to stop producing proxy messages
	serverContext     context.Context
	stopServerContext context.CancelFunc
	serverWaitGroup   *sync.WaitGroup
}

type proxyData interface {
	load()
	store()
}

type proxyEndpoint struct {
	common.EndpointStore
	nextTarget string
}

type proxyStore struct {
	Data *sync.Map
}

func makeProxyStore() *proxyStore {
	return &proxyStore{
		Data: new(sync.Map),
	}
}

func (e *proxyEndpoint) next(scope string) (endpointURL string) {
	endpointURL = ""
	e.Data.Range(func(id, endpoint interface{}) bool {
		ids := id.(string)
		if e.nextTarget == "" {
			e.nextTarget = ids
		}
		if endpointURL != "" {
			e.nextTarget = ids
			// exit Range iteration as next target is identified
			return false
		}
		switch scope {
		case public:
			if ids == e.nextTarget {
				endpointURL = endpoint.(*models.Endpoint).PublicURL
				// let Range iterate till nextServer is identified
				return true
			}
		case private:
			if ids == e.nextTarget {
				endpointURL = endpoint.(*models.Endpoint).PrivateURL
				// let Range iterate till nextServer is identified
				return true
			}
		}
		return true
	})
	return endpointURL
}

func (e *proxyEndpoint) load(id string) *models.Endpoint {
	ep, ok := e.Data.Load(id)
	if !ok {
		return nil
	}
	endpoint, ok := ep.(*models.Endpoint)
	return endpoint
}

func (e *proxyEndpoint) store(id string, endpoint *models.Endpoint) {
	e.Data.Store(id, endpoint)
}

func (s *proxyStore) load(proxyPrefix string) *proxyEndpoint {
	p, ok := s.Data.Load(proxyPrefix)
	if !ok {
		return nil
	}
	proxyEndpoint, ok := p.(*proxyEndpoint)
	return proxyEndpoint

}

func (s *proxyStore) store(proxyPrefix string, proxyEndpoint *proxyEndpoint) {
	s.Data.Store(proxyPrefix, proxyEndpoint)
}

// NewProxyServer creates new proxy server
func NewProxyServer(e *echo.Echo, proxyStore *proxyStore, dbService serviceif.Service) *ProxyServer {
	p := &ProxyServer{
		dbService:  dbService,
		echoServer: e,
		proxyStore: proxyStore,
	}
	return p
}

func (p *ProxyServer) readEndpoints() (map[string]*models.Endpoint, error) {
	ctx := common.NoAuth(context.Background())
	spec := models.ListSpec{Limit: 100}
	request := &models.ListEndpointRequest{Spec: &spec}
	response, err := p.dbService.ListEndpoint(ctx, request)
	if err != nil {
		return nil, err
	}
	endpoints := make(map[string]*models.Endpoint)
	for _, e := range response.Endpoints {
		endpoints[e.UUID] = e
	}
	return endpoints, nil
}

func (p *ProxyServer) getProxyPrefixFromURL(urlPath string, scope string) (proxyPrefix string) {
	paths := strings.Split(urlPath, pathSep)
	prefixes := make([]string, 4)
	copy(prefixes, paths[:4])
	prefixes = append(prefixes, scope)
	return strings.Join(prefixes, pathSep)
}

func (p *ProxyServer) getReverseProxyServer(urlPath string) (
	prefix string, server *httputil.ReverseProxy) {
	var scope string
	if strings.Contains(urlPath, private) {
		scope = private
	} else {
		scope = public
	}
	proxyPrefix := p.getProxyPrefixFromURL(urlPath, scope)
	proxyEndpoint := p.proxyStore.load(proxyPrefix)
	target := proxyEndpoint.next(scope)
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

func (p *ProxyServer) dynamicProxyMiddleware() func(next echo.HandlerFunc) echo.HandlerFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			r := c.Request()
			prefix, server := p.getReverseProxyServer(r.URL.Path)
			r.URL.Path = strings.TrimPrefix(r.URL.Path, prefix)
			w := c.Response()
			server.ServeHTTP(w, r)
			return nil
		}
	}
}

func (p *ProxyServer) checkDeleteProxyEndpoints(endpoints map[string]*models.Endpoint) {
	wg := &sync.WaitGroup{}
	p.proxyStore.Data.Range(func(prefix, proxy interface{}) bool {
		s, ok := proxy.(*proxyEndpoint)
		if !ok {
			log.Fatalf("Unable to load cluster(%s)'s proxy data from in-memory store",
				prefix)
			return true
		}
		s.Data.Range(func(id, endpoint interface{}) bool {
			_, ok := endpoint.(*models.Endpoint)
			if !ok {
				log.Fatalf("Unable to load endpoint(%s) data from in-memory store",
					id)
				return true
			}
			wg.Add(1)
			// call N go routines to delete proxy middleware
			go func(s *proxyEndpoint) {
				defer wg.Done()
				ids := id.(string)
				_, ok := endpoints[ids]
				if !ok {
					s.Data.Delete(id)
					log.Debugf("deleting dynamic proxy endpoint for id: %s", ids)
				}
			}(s)
			return true
		})
		return true
	})
	wg.Wait()
}

func (p *ProxyServer) getProxyPrefix(endpoint *models.Endpoint, scope string) (proxyPrefix string) {
	prefixes := []string{"", prefix, endpoint.ParentUUID, endpoint.Name, scope}
	return strings.Join(prefixes, pathSep)
}

func (p *ProxyServer) initProxyEndpointStore(endpoint *models.Endpoint) {
	if endpoint.PublicURL != "" {
		proxyPrefix := p.getProxyPrefix(endpoint, public)
		_, ok := p.proxyStore.Data.Load(proxyPrefix)
		if !ok {
			e := &proxyEndpoint{
				EndpointStore: common.EndpointStore{
					Data: new(sync.Map),
				},
				nextTarget: "",
			}
			p.proxyStore.Data.Store(proxyPrefix, e)
		}
	}
	if endpoint.PrivateURL != "" {
		proxyPrefix := p.getProxyPrefix(endpoint, private)
		_, ok := p.proxyStore.Data.Load(proxyPrefix)
		if !ok {
			e := &proxyEndpoint{
				EndpointStore: common.EndpointStore{
					Data: new(sync.Map),
				},
				nextTarget: "",
			}
			p.proxyStore.Data.Store(proxyPrefix, e)
		}
	}
}

func (p *ProxyServer) manageProxyEndpoint(endpoint *models.Endpoint, scope string) {
	proxyPrefix := p.getProxyPrefix(endpoint, scope)
	s := p.proxyStore.load(proxyPrefix)
	if s == nil {
		log.Fatalf("proxy endpoint for %s is not found in-memory store",
			proxyPrefix)
	}
	e := s.load(endpoint.UUID)
	if e != nil { // proxy endpoint in memory store
		if e != endpoint {
			// update the endpoint store
			s.store(endpoint.UUID, endpoint)
		}
	} else { // proxy endpoint not in memory store
		// store and create proxy middleware
		log.Debugf("new endpoint(%s, %s) created",
			endpoint.Name, endpoint.UUID)
		s.store(endpoint.UUID, endpoint)
	}
}

func (p *ProxyServer) syncProxyEndpoints(endpoints map[string]*models.Endpoint) {
	// delete stale proxy endpoints in-memory
	p.checkDeleteProxyEndpoints(endpoints)
	// create/update proxy middleware
	wg := &sync.WaitGroup{}
	for _, endpoint := range endpoints {
		p.initProxyEndpointStore(endpoint)
		wg.Add(1)
		// call N go routines to create/update proxy middleware
		go func(endpoint *models.Endpoint) {
			defer wg.Done()
			if endpoint.PublicURL != "" {
				p.manageProxyEndpoint(endpoint, public)
			}
			if endpoint.PrivateURL != "" {
				p.manageProxyEndpoint(endpoint, private)
			}
		}(endpoint)
	}
	wg.Wait()
}

func (p *ProxyServer) serve() {
	// add dynamic proxy middleware
	g := p.echoServer.Group(prefix)
	//g.Use(removePathPrefixMiddleware(prefix))
	g.Use(p.dynamicProxyMiddleware())

	p.serverContext, p.stopServerContext = context.WithCancel(context.Background())
	p.serverWaitGroup = &sync.WaitGroup{}
	p.serverWaitGroup.Add(1)
	go func() {
		// serve forever
		defer p.serverWaitGroup.Done()
		var err error
		var endpoints map[string]*models.Endpoint
		for {
			select {
			case <-p.serverContext.Done():
				log.Info("stopping dynamic proxy server")
				return
			default:
				// poll db for the endpoint resource
				endpoints, err = p.readEndpoints()
				if err != nil {
					log.Error("not able to read endpoints")
					log.Fatal(err)
				}
				if endpoints != nil {
					// create/update/delete proxy endpoints in-memory
					p.syncProxyEndpoints(endpoints)
				}
			}
		}
	}()
}

func (p *ProxyServer) stop() {
	// stop proxy server poll
	p.stopServerContext()
	// wait for the proxy server poll to complete
	p.serverWaitGroup.Wait()
}
