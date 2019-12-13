package apisrv

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"strings"
	"sync"
	"time"

	"github.com/Juniper/asf/pkg/format"
	"github.com/Juniper/asf/pkg/logutil"
	"github.com/Juniper/asf/pkg/proxy"
	"github.com/Juniper/asf/pkg/services/baseservices"
	"github.com/Juniper/contrail/pkg/auth"
	"github.com/Juniper/contrail/pkg/client"
	"github.com/Juniper/contrail/pkg/endpoint"
	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/services"
	"github.com/labstack/echo"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"

	asfclient "github.com/Juniper/asf/pkg/client"
)

// Proxy service related constants.
const (
	DefaultDynamicProxyPath = "proxy"
	ProxySyncInterval       = 2 * time.Second
	XClusterIDKey           = "X-Cluster-ID"
	XServiceTokenKey        = "X-Service-Token"

	limit         = 100
	pathSeparator = "/"
)

// proxyService provides dynamic HTTP and WebSockets proxy capabilities.
// TODO(dfurman): move to subpackage to improve design
type proxyService struct {
	endpointStore    *endpoint.Store
	dbService        services.Service
	dynamicProxyPath string
	stopService      context.CancelFunc
	serviceWaitGroup *sync.WaitGroup
	forceUpdateChan  chan chan struct{}
	log              *logrus.Entry
}

// DynamicProxyConfig configures the dynamic proxy service and middleware.
type DynamicProxyConfig struct {
	// Path is the path the proxy is served at.
	Path string
	// ServiceTokenEndpointPrefixes are prefixes of the endpoints that should have a service token injected into requests.
	ServiceTokenEndpointPrefixes []string
	// ServiceUserClientConfig is an HTTP config that allows logging in as the service user.
	ServiceUserClientConfig *asfclient.HTTPConfig
}

func newProxyService(es *endpoint.Store, dbService services.Service, config *DynamicProxyConfig) *proxyService {
	return &proxyService{
		endpointStore:    es,
		dbService:        dbService,
		dynamicProxyPath: config.Path,
		forceUpdateChan:  make(chan chan struct{}),
		log:              logutil.NewLogger("proxy-service"),
	}
}

func dynamicProxyMiddleware(
	es *endpoint.Store, config *DynamicProxyConfig,
) func(next echo.HandlerFunc) echo.HandlerFunc {
	log := logutil.NewLogger("dynamic-proxy-mw")

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			r := ctx.Request()

			clusterID, err := clusterID(r.URL.Path, config.Path)
			if err != nil {
				return echo.NewHTTPError(http.StatusBadRequest, err.Error())
			}
			setClusterIDHeader(r, clusterID)

			pp := proxyPrefix(r.URL.Path, urlScope(r.URL.Path))
			scope := urlScope(r.URL.Path)
			r.URL.Path = withNoProxyPrefix(r.URL.Path, pp)

			if shouldInjectServiceToken(pp, config) {
				var token string
				token, err = obtainServiceToken(r.Context(), clusterID, config)
				if err != nil {
					log.WithError(err).Error("Failed to obtain service token - not adding it to the request")
				} else {
					setServiceTokenHeader(r, token)
				}
			}

			t, err := readTargets(es, scope, pp)
			if err != nil {
				return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
			}

			if err = proxy.HandleRequest(ctx, rawTargetURLs(t), log); err != nil {
				return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("Failed to handle proxy request: %v", err))
			}

			return nil
		}
	}
}

func proxyPrefix(urlPath string, scope string) string {
	prefixes := strings.Split(urlPath, pathSeparator)[:4]
	return strings.Join(append(prefixes, scope), pathSeparator)
}

func withNoProxyPrefix(path, pp string) string {
	return strings.TrimPrefix(path, strings.TrimSuffix(pp, endpoint.PublicURLScope))
}

func readTargets(es *endpoint.Store, scope string, proxyPrefix string) ([]*endpoint.Endpoint, error) {
	ts := es.Read(proxyPrefix)
	if ts == nil {
		return nil, errors.Errorf("target store not found for given proxy prefix: %v", proxyPrefix)
	}

	targets := ts.ReadAll(scope)
	if len(targets) == 0 {
		return nil, errors.New("no targets in proxy target store")
	}

	return targets, nil
}

func urlScope(urlPath string) string {
	if strings.Contains(urlPath, endpoint.PrivateURLScope) {
		return endpoint.PrivateURLScope
	}
	return endpoint.PublicURLScope
}

func rawTargetURLs(targets []*endpoint.Endpoint) []string {
	var u []string
	for _, target := range targets {
		u = append(u, target.URL)
	}
	return u
}

func setClusterIDHeader(r *http.Request, clusterID string) {
	r.Header.Set(XClusterIDKey, clusterID)
}

func clusterID(url, dynamicProxyPath string) (string, error) {
	paths := strings.Split(url, pathSeparator)
	if len(paths) <= 3 || paths[1] != dynamicProxyPath {
		return "", errors.Errorf("cluster ID not found in proxy URL: %v", url)
	}
	return paths[2], nil
}

func shouldInjectServiceToken(proxyPrefix string, config *DynamicProxyConfig) bool {
	paths := strings.Split(proxyPrefix, pathSeparator)
	if len(paths) < 4 {
		return false
	}
	endpointPrefix := paths[3]
	allowedPrefixes := config.ServiceTokenEndpointPrefixes
	return format.ContainsString(allowedPrefixes, endpointPrefix)
}

func obtainServiceToken(ctx context.Context, clusterID string, config *DynamicProxyConfig) (string, error) {
	apiClient := client.NewHTTP(config.ServiceUserClientConfig)

	ctx = auth.WithXClusterID(ctx, clusterID)
	if err := apiClient.Login(ctx); err != nil {
		return "", errors.Wrap(err, "failed to log in as service user")
	}

	if apiClient.AuthToken == "" {
		return "", errors.New("keystone returned no service token for service user")
	}
	return apiClient.AuthToken, nil
}

func setServiceTokenHeader(r *http.Request, token string) {
	r.Header.Set(XServiceTokenKey, token)
}

// StartEndpointsSync starts synchronization of proxy endpoints.
func (p *proxyService) StartEndpointsSync() {
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

// StopEndpointsSync stops synchronization of proxy endpoints.
func (p *proxyService) StopEndpointsSync() {
	// stop proxy server poll
	p.stopService()
	// wait for the proxy server poll to complete
	p.serviceWaitGroup.Wait()
}
