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
	"github.com/spf13/viper"

	"github.com/Juniper/contrail/pkg/common"
	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/serviceif"

	apicommon "github.com/Juniper/contrail/pkg/apisrv/common"
	log "github.com/sirupsen/logrus"
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
	dbService     serviceif.Service
	EndpointStore *apicommon.EndpointStore
	// context to stop servicing proxy endpoints
	serviceContext     context.Context
	stopServiceContext context.CancelFunc
	serviceWaitGroup   *sync.WaitGroup
}

func newProxyService(e *echo.Echo, endpointStore *apicommon.EndpointStore,
	dbService serviceif.Service) *proxyService {
	group := viper.GetString("proxy.group")
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
	spec := models.ListSpec{Limit: limit}
	for {
		request := &models.ListEndpointRequest{Spec: &spec}
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

func (p *proxyService) getReverseService(urlPath string) (
	prefix string, server *httputil.ReverseProxy) {
	var scope string
	if strings.Contains(urlPath, private) {
		scope = private
	} else {
		scope = public
	}
	proxyPrefix := p.getProxyPrefixFromURL(urlPath, scope)
	proxyEndpoint := p.EndpointStore.Read(proxyPrefix)
	target := proxyEndpoint.Next(scope)
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
			prefix, server := p.getReverseService(r.URL.Path)
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
			log.Fatalf("Unable to Read cluster(%s)'s proxy data from in-memory Write",
				prefix)
			return true
		}
		s.Data.Range(func(id, endpoint interface{}) bool {
			_, ok := endpoint.(*models.Endpoint)
			if !ok {
				log.Fatalf("Unable to Read endpoint(%s) data from in-memory Write",
					id)
				return true
			}
			ids := id.(string)
			_, ok = endpoints[ids]
			if !ok {
				s.Data.Delete(id)
				log.Debugf("deleting dynamic proxy endpoint for id: %s", ids)
			}
			return true
		})
		return true
	})
}

func (p *proxyService) getProxyPrefix(endpoint *models.Endpoint, scope string) (proxyPrefix string) {
	prefixes := []string{"", p.group, endpoint.ParentUUID, endpoint.Name, scope}
	return strings.Join(prefixes, pathSep)
}

func (p *proxyService) initProxyTargetStore(endpoint *models.Endpoint) {
	if endpoint.PublicURL != "" {
		proxyPrefix := p.getProxyPrefix(endpoint, public)
		_, ok := p.EndpointStore.Data.Load(proxyPrefix)
		if !ok {
			p.EndpointStore.Data.Store(proxyPrefix, apicommon.MakeTargetStore())
		}
	}
	if endpoint.PrivateURL != "" {
		proxyPrefix := p.getProxyPrefix(endpoint, private)
		_, ok := p.EndpointStore.Data.Load(proxyPrefix)
		if !ok {
			p.EndpointStore.Data.Store(proxyPrefix, apicommon.MakeTargetStore())
		}
	}
}

func (p *proxyService) manageProxyEndpoint(endpoint *models.Endpoint, scope string) {
	proxyPrefix := p.getProxyPrefix(endpoint, scope)
	s := p.EndpointStore.Read(proxyPrefix)
	if s == nil {
		log.Fatalf("endpoint Write for %s is not found in-memory Write",
			proxyPrefix)
	}
	e := s.Read(endpoint.UUID)
	if e != endpoint {
		// proxy endpoint not in memory Write or
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
		}
		if endpoint.PrivateURL != "" {
			p.manageProxyEndpoint(endpoint, private)
		}
	}
}

func (p *proxyService) serve() {
	// add dynamic proxy middleware
	g := p.echoServer.Group(p.group)
	g.Use(p.dynamicProxyMiddleware())

	p.serviceContext, p.stopServiceContext = context.WithCancel(context.Background())
	p.serviceWaitGroup = &sync.WaitGroup{}
	p.serviceWaitGroup.Add(1)
	go func() {
		// serve forever
		defer p.serviceWaitGroup.Done()
		var err error
		var endpoints map[string]*models.Endpoint
		for {
			select {
			case <-p.serviceContext.Done():
				log.Info("stopping dynamic proxy server")
				return
			default:
				// poll db for the endpoint resource
				endpoints, err = p.readEndpoints()
				if err != nil {
					// log and continue during DB read
					log.Debug("Endpoints read failed")
					log.Error(err)
				}
				if endpoints != nil {
					// create/update/delete proxy endpoints in-memory
					p.syncProxyEndpoints(endpoints)
				}
				time.Sleep(2 * time.Second)
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
