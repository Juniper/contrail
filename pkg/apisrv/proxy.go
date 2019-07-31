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
	"github.com/Juniper/contrail/pkg/logutil"
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
	ProxySyncInterval       = 2 * time.Second
	XClusterIDKey           = "X-Cluster-ID"

	limit         = 100
	pathSeparator = "/"
)

// proxyService provides dynamic HTTP and WebSockets proxy capabilities.
// TODO(Daniel): move to subpackage to improve design
// TODO(Daniel): proxy to other targets when 502/503 received from target
type proxyService struct {
	endpointStore    *endpoint.Store
	dbService        services.Service
	dynamicProxyPath string
	stopService      context.CancelFunc
	serviceWaitGroup *sync.WaitGroup
	forceUpdateChan  chan chan struct{}
	log              *logrus.Entry
}

func newProxyService(
	e *echo.Echo, es *endpoint.Store, dbService services.Service, dynamicProxyPath string,
) *proxyService {
	log := logutil.NewLogger("proxy-service")

	if dynamicProxyPath == "" {
		dynamicProxyPath = DefaultDynamicProxyPath
	}
	e.Group(dynamicProxyPath, dynamicProxyMiddleware(es, dynamicProxyPath, log))

	return &proxyService{
		endpointStore:    es,
		dbService:        dbService,
		dynamicProxyPath: dynamicProxyPath,
		forceUpdateChan:  make(chan chan struct{}),
		log:              log,
	}
}

func dynamicProxyMiddleware(
	es *endpoint.Store, dynamicProxyPath string, log *logrus.Entry,
) func(next echo.HandlerFunc) echo.HandlerFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			r := c.Request()
			pp := proxyPrefix(r.URL.Path, scope(r.URL.Path))
			rp, err := reverseProxy(es, r.URL.Path, pp, log)
			if err != nil {
				return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
			}

			if err = setClusterIDKeyHeader(r, dynamicProxyPath); err != nil {
				return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
			}

			r.URL.Path = strings.TrimPrefix(r.URL.Path, strings.TrimSuffix(pp, endpoint.PublicURLScope))

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

func reverseProxy(
	es *endpoint.Store, urlPath string, proxyPrefix string, log *logrus.Entry,
) (*httputil.ReverseProxy, error) {
	ts := es.Read(proxyPrefix)
	if ts == nil {
		return nil, errors.Errorf("target store not found for given proxy prefix: %v", proxyPrefix)
	}

	targets := ts.ReadAll(scope(urlPath))
	if targets == nil {
		return nil, errors.New("no targets in proxy target store")
	}

	return proxy.NewReverseProxy(parseTargetURLs(targets, log))
}

func scope(urlPath string) string {
	if strings.Contains(urlPath, endpoint.PrivateURLScope) {
		return endpoint.PrivateURLScope
	}
	return endpoint.PublicURLScope
}

func parseTargetURLs(targets []*endpoint.Endpoint, log *logrus.Entry) []*url.URL {
	var targetURLs []*url.URL
	for _, target := range targets {
		tURL, err := url.Parse(target.URL)
		if err != nil {
			log.WithError(err).WithField("target-url", target.URL).Error("Failed to parse target URL - ignoring")
		} else {
			targetURLs = append(targetURLs, tURL)
		}
	}
	return targetURLs
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
				p.log.Info("Stopping dynamic proxy service")
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
		p.log.WithError(err).Error("Endpoints read failed")
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
			p.manageProxyEndpoint(e, endpoint.PublicURLScope)
			p.manageProxyEndpoint(e, endpoint.PrivateURLScope)
		}
	}
}

func (p *proxyService) checkDeleted(endpoints map[string]*models.Endpoint) {
	p.endpointStore.Data.Range(func(prefix, proxy interface{}) bool {
		s, ok := proxy.(*endpoint.TargetStore)
		if !ok {
			p.log.WithField("prefix", prefix).Error("Unable to read cluster's proxy data from in-memory store")
			return true
		}
		s.Data.Range(func(id, endpoint interface{}) bool {
			_, ok := endpoint.(*models.Endpoint)
			if !ok {
				p.log.WithField("id", id).Error("Unable to Read endpoint data from in-memory store")
				return true
			}
			ids, ok := id.(string)
			if !ok {
				p.log.WithField("id", id).Error("Unable to convert id to string when looking endpointStore")
				return true
			}
			_, ok = endpoints[ids]
			if !ok {
				s.Remove(ids)
				p.log.WithField("id", ids).Debug("Deleting dynamic proxy endpoint")
			}
			return true
		})
		if s.Count() == 0 {
			prefixStr, ok := prefix.(string)
			if !ok {
				p.log.WithField("prefix", prefix).Error("Unable to convert prefix to string")
			}
			p.endpointStore.Remove(prefixStr)
			p.log.WithField("prefix", prefixStr).Debug("Deleting dynamic proxy endpoint prefix")
		}
		return true
	})
}

func (p *proxyService) initProxyTargetStore(e *models.Endpoint) {
	if e.PublicURL != "" {
		proxyPrefix := p.getProxyPrefix(e, endpoint.PublicURLScope)
		targetStore := p.endpointStore.Read(proxyPrefix)
		if targetStore == nil {
			p.endpointStore.Write(proxyPrefix, endpoint.NewTargetStore())
		}
		privateProxyPrefix := p.getProxyPrefix(e, endpoint.PrivateURLScope)
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
		p.log.WithField("prefix", proxyPrefix).Error("Endpoint store for prefix is not found in-memory store")
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
		p.log.WithFields(logrus.Fields{
			"prefix":        prefix,
			"endpoint-uuid": endpoint.UUID,
		}).Error("Parent UUID missing for endpoint")
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
