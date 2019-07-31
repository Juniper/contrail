package apisrv

import (
	"context"
	"net/http"
	"net/http/httputil"
	"net/url"
	"reflect"
	"strings"
	"sync"
	"time"

	"github.com/Juniper/contrail/pkg/apisrv/endpoint"
	"github.com/Juniper/contrail/pkg/apisrv/proxy"
	"github.com/Juniper/contrail/pkg/auth"
	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/services"
	"github.com/Juniper/contrail/pkg/services/baseservices"
	"github.com/labstack/echo"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

// Proxy service related constants.
const (
	DefaultDynamicProxyPath = "proxy"
	ProxySyncInterval       = time.Second
	XClusterIDKey           = "X-Cluster-ID"

	pathSeparator = "/"
	limit         = 100
)

// proxyService provides dynamic HTTP and WebSockets proxy capabilities.
// TODO(Daniel): move to subpackage to improve design
type proxyService struct {
	endpointStore    *endpoint.Store
	dbService        services.Service
	dynamicProxyPath string
	stopService      context.CancelFunc
	serviceWaitGroup *sync.WaitGroup
	forceUpdateChan  chan chan struct{}
}

func newProxyService(
	e *echo.Echo, es *endpoint.Store, dbService services.Service, dynamicProxyPath string,
) *proxyService {
	if dynamicProxyPath == "" {
		dynamicProxyPath = DefaultDynamicProxyPath
	}
	e.Group(dynamicProxyPath, dynamicProxyMiddleware(es, dynamicProxyPath))

	return &proxyService{
		endpointStore:    es,
		dbService:        dbService,
		dynamicProxyPath: dynamicProxyPath,
		forceUpdateChan:  make(chan chan struct{}),
	}
}

func dynamicProxyMiddleware(es *endpoint.Store, dynamicProxyPath string) func(next echo.HandlerFunc) echo.HandlerFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			r := c.Request()
			pp := proxyPrefix(r.URL.Path, scope(r.URL.Path))

			rp, err := reverseProxy(es, r.URL.Path, pp)
			if err != nil {
				return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
			}

			if err = setClusterIDKeyHeader(r, dynamicProxyPath); err != nil {
				return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
			}

			r.URL.Path = strings.TrimPrefix(r.URL.Path, strings.TrimSuffix(pp, endpoint.Public))

			rp.ServeHTTP(c.Response(), r)
			return nil
		}
	}
}

func proxyPrefix(urlPath string, scope string) string {
	prefixes := make([]string, 4)
	copy(prefixes, strings.Split(urlPath, pathSeparator)[:4])
	return strings.Join(append(prefixes, scope), pathSeparator)
}

func reverseProxy(es *endpoint.Store, urlPath string, proxyPrefix string) (*httputil.ReverseProxy, error) {
	logrus.Warn("Hoge func reverseProxy")
	ts := es.Read(proxyPrefix)
	if ts == nil {
		return nil, errors.Errorf("target store not found for given proxy prefix: %v", proxyPrefix)
	}

	targets := ts.ReadAll(scope(urlPath))
	if targets == nil {
		return nil, errors.Errorf("failed to read endpoint targets from proxy target store; urlPath: %v", urlPath)
	}

	targetURL, err := url.Parse(targets[0].URL) // TODO: use all targets
	if err != nil {
		return nil, errors.Wrapf(err, "failed to parse target URL: %v", targets[0].URL)
	}

	//return newMultipleHostReverseProxy(targets, targetURL), nil
	return proxy.NewReverseProxy(targetURL), nil
}

func scope(urlPath string) string {
	if strings.Contains(urlPath, endpoint.Private) {
		return endpoint.Private
	}
	return endpoint.Public
}

// setClusterIDKeyHeader Sets cluster ID in proxy request, so that the proxy endpoints can use it
// to get server's Keystone token.
func setClusterIDKeyHeader(r *http.Request, dynamicProxyPath string) error {
	cid := clusterID(r.URL.Path, dynamicProxyPath)
	if cid == "" {
		return errors.Errorf("cluster ID not found in proxy URL: %v", r.URL.Path)
	}

	r.Header.Set(XClusterIDKey, cid)
	return nil
}

func clusterID(url, dynamicProxyPath string) (clusterID string) {
	paths := strings.Split(url, pathSeparator)
	if len(paths) > 3 && paths[1] == dynamicProxyPath {
		return paths[2]
	}
	return ""
}

// Serve starts proxy service.
func (p *proxyService) Serve() {
	serviceCtx, cancel := context.WithCancel(context.Background())
	p.stopService = cancel

	p.serviceWaitGroup = &sync.WaitGroup{}
	p.serviceWaitGroup.Add(1)

	go func() {
		defer p.serviceWaitGroup.Done()
		ticker := time.NewTicker(ProxySyncInterval)
		defer ticker.Stop()
		for {
			select {
			case <-serviceCtx.Done():
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
	for _, e := range endpoints {
		p.initProxyTargetStore(e)
		if e.PublicURL != "" {
			p.manageProxyEndpoint(e, endpoint.Public)
			p.manageProxyEndpoint(e, endpoint.Private)
		}
	}
}

func (p *proxyService) checkDeleted(endpoints map[string]*models.Endpoint) {
	p.endpointStore.Data.Range(func(prefix, proxy interface{}) bool {
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
				logrus.Errorf("Unable to convert id %v to string when looking endpointStore", id)
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
			p.endpointStore.Remove(prefixStr)
			logrus.Debugf("deleting dynamic proxy endpoint prefix: %s", prefixStr)
		}
		return true
	})
}

func (p *proxyService) initProxyTargetStore(e *models.Endpoint) {
	if e.PublicURL != "" {
		proxyPrefix := p.getProxyPrefix(e, endpoint.Public)
		targetStore := p.endpointStore.Read(proxyPrefix)
		if targetStore == nil {
			p.endpointStore.Write(proxyPrefix, endpoint.NewTargetStore())
		}
		privateProxyPrefix := p.getProxyPrefix(e, endpoint.Private)
		privateTargetStore := p.endpointStore.Read(privateProxyPrefix)
		if privateTargetStore == nil {
			p.endpointStore.Write(privateProxyPrefix, endpoint.NewTargetStore())
		}
	}
}

func (p *proxyService) manageProxyEndpoint(endpoint *models.Endpoint, scope string) {
	proxyPrefix := p.getProxyPrefix(endpoint, scope)
	s := p.endpointStore.Read(proxyPrefix)
	if s == nil {
		logrus.Errorf("endpoint store for %s is not found in-memory store", proxyPrefix)
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
	prefixes := []string{"", p.dynamicProxyPath, endpoint.ParentUUID, prefix, scope}
	return strings.Join(prefixes, pathSeparator)
}

// ForceUpdate requests an immediate update of endpoints and waits for its completion.
func (p *proxyService) ForceUpdate() {
	wait := make(chan struct{})
	p.forceUpdateChan <- wait
	<-wait
}

// Stop stops proxy service.
func (p *proxyService) Stop() {
	// stop proxy server poll
	p.stopService()
	// wait for the proxy server poll to complete
	p.serviceWaitGroup.Wait()
}
