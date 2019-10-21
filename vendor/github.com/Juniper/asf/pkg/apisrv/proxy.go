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

	"github.com/Juniper/asf/pkg/apisrv/endpoint"

	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	"github.com/Juniper/asf/pkg/auth"
	"github.com/Juniper/asf/pkg/models"
	"github.com/Juniper/asf/pkg/services"
	"github.com/Juniper/asf/pkg/services/baseservices"
)

const (
	// DefaultDynamicProxyPath default value for server.dynamic_proxy_path
	DefaultDynamicProxyPath = "proxy"
	public                  = endpoint.Public
	private                 = endpoint.Private
	pathSep                 = "/"
	limit                   = 100
	xClusterIDKey           = "X-Cluster-ID"
)

type proxyService struct {
	group         string
	echoServer    *echo.Echo
	dbService     services.Service
	EndpointStore *endpoint.Store
	// context to stop servicing proxy endpoints
	serviceContext     context.Context
	stopServiceContext context.CancelFunc
	serviceWaitGroup   *sync.WaitGroup
	forceUpdateChan    chan chan struct{}
}

func newProxyService(e *echo.Echo, endpointStore *endpoint.Store,
	dbService services.Service) *proxyService {
	group := viper.GetString("server.dynamic_proxy_path")
	if group == "" {
		group = DefaultDynamicProxyPath
	}
	p := &proxyService{
		group:         group,
		dbService:     dbService,
		echoServer:    e,
		EndpointStore: endpointStore,
	}
	return p
}

// Serve starts proxy service.
func (p *proxyService) Serve() {
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
				logrus.Info("Stopping dynamic proxy server")
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

func (p *proxyService) dynamicProxyMiddleware() func(next echo.HandlerFunc) echo.HandlerFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			r := c.Request()
			clusterID := getClusterIDFromProxyURL(r.URL.Path)
			if clusterID != "" {
				//set clusterID in proxy request, so that the
				//proxy endpoints can use it to get server's
				//keystone token.
				r.Header.Set(xClusterIDKey, clusterID)
			} else {
				return echo.NewHTTPError(http.StatusInternalServerError,
					"Cluster ID not found in proxy URL")
			}
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

// getClusterIDFromProxyURL parses the proxy url to retrieve clusterID.
func getClusterIDFromProxyURL(url string) (clusterID string) {
	paths := strings.Split(url, pathSep)
	if len(paths) > 3 && paths[1] == "proxy" {
		clusterID = paths[2]
	}
	return clusterID
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
		logrus.WithField("proxy-prefix", proxyPrefix).Info("Endpoint targets not found for given proxy prefix")
		return strings.TrimSuffix(proxyPrefix, public), nil
	}
	target := proxyEndpoint.Next(scope)
	if target == nil {
		return strings.TrimSuffix(proxyPrefix, public), nil
	}
	insecure := true //TODO:(ijohnson) add insecure to endpoint schema

	u, err := url.Parse(target.URL)
	if err != nil {
		logrus.WithError(err).WithField("target", target.URL).Info("Failed to parse target - ignoring")
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

func (p *proxyService) getProxyPrefixFromURL(urlPath string, scope string) (proxyPrefix string) {
	paths := strings.Split(urlPath, pathSep)
	prefixes := make([]string, 4)
	copy(prefixes, paths[:4])
	prefixes = append(prefixes, scope)
	return strings.Join(prefixes, pathSep)
}

func (p *proxyService) updateEndpoints() {
	endpoints, err := p.readEndpoints()
	if err != nil {
		logrus.WithError(err).Error("Endpoints read failed")
		return
	}
	p.syncProxyEndpoints(endpoints)
}

func (p *proxyService) readEndpoints() (map[string]*models.Endpoint, error) {
	endpoints := make(map[string]*models.Endpoint)
	ctx := auth.NoAuth(context.Background())
	spec := baseservices.ListSpec{Limit: limit}
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
		marker := response.Endpoints[len(response.Endpoints)-1].UUID
		spec.Marker = marker
	}
	return endpoints, nil
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

func (p *proxyService) checkDeleted(endpoints map[string]*models.Endpoint) {
	p.EndpointStore.Data.Range(func(prefix, proxy interface{}) bool {
		s, ok := proxy.(*endpoint.TargetStore)
		if !ok {
			logrus.Errorf("Unable to Read cluster(%s)'s proxy data from in-memory store",
				prefix)
			return true
		}
		s.Data.Range(func(id, endpoint interface{}) bool {
			_, ok := endpoint.(*models.Endpoint)
			if !ok {
				logrus.Errorf("Unable to Read endpoint(%s) data from in-memory store",
					id)
				return true
			}
			ids, ok := id.(string)
			if !ok {
				logrus.Errorf("Unable to convert id %v to string when looking EndpointStore", id)
				return true
			}
			_, ok = endpoints[ids]
			if !ok {
				s.Remove(ids)
				logrus.Debugf("deleting dynamic proxy endpoint for id: %s", ids)
			}
			return true
		})
		if s.Count() == 0 {
			prefixStr, ok := prefix.(string)
			if !ok {
				logrus.Errorf("Unable to convert prefix %v to string", prefix)
			}
			p.EndpointStore.Remove(prefixStr)
			logrus.Debugf("deleting dynamic proxy endpoint prefix: %s", prefixStr)
		}
		return true
	})
}

func (p *proxyService) initProxyTargetStore(e *models.Endpoint) {
	if e.PublicURL != "" {
		proxyPrefix := p.getProxyPrefix(e, public)
		targetStore := p.EndpointStore.Read(proxyPrefix)
		if targetStore == nil {
			p.EndpointStore.Write(proxyPrefix, endpoint.NewTargetStore())
		}
		privateProxyPrefix := p.getProxyPrefix(e, private)
		privateTargetStore := p.EndpointStore.Read(privateProxyPrefix)
		if privateTargetStore == nil {
			p.EndpointStore.Write(privateProxyPrefix, endpoint.NewTargetStore())
		}
	}
}

func (p *proxyService) manageProxyEndpoint(endpoint *models.Endpoint, scope string) {
	proxyPrefix := p.getProxyPrefix(endpoint, scope)
	s := p.EndpointStore.Read(proxyPrefix)
	if s == nil {
		logrus.Errorf("endpoint store for %s is not found in-memory store",
			proxyPrefix)
	}
	e := s.Read(endpoint.UUID)
	if !reflect.DeepEqual(e, endpoint) {
		// proxy endpoint not in memory store or
		// proxy endpoint updated
		s.Write(endpoint.UUID, endpoint)
	}
}

func (p *proxyService) getProxyPrefix(endpoint *models.Endpoint, scope string) (proxyPrefix string) {
	prefix := endpoint.Prefix
	// TODO(ijohnson) remove using DisplayName as prefix
	// once UI takes prefix as input.
	if prefix == "" {
		prefix = endpoint.DisplayName
	}

	if endpoint.ParentUUID == "" {
		logrus.Errorf("Parent uuid missing for endpoint %s(%s)", prefix, endpoint.UUID)
	}
	prefixes := []string{"", p.group, endpoint.ParentUUID, prefix, scope}
	return strings.Join(prefixes, pathSep)
}

// ForceUpdate requests an immediate update of endpoints and waits for its completion.
func (p *proxyService) ForceUpdate() {
	wait := make(chan struct{})
	p.forceUpdateChan <- wait
	<-wait
}

// Stop stop proxy service.
func (p *proxyService) Stop() {
	// stop proxy server poll
	p.stopServiceContext()
	// wait for the proxy server poll to complete
	p.serviceWaitGroup.Wait()
}
