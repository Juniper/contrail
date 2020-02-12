package proxy

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/Juniper/asf/pkg/apiserver"
	"github.com/Juniper/asf/pkg/auth"
	"github.com/Juniper/asf/pkg/format"
	"github.com/Juniper/asf/pkg/logutil"
	"github.com/Juniper/asf/pkg/proxy"
	"github.com/Juniper/asf/pkg/services/baseservices"
	"github.com/Juniper/contrail/pkg/client"
	"github.com/Juniper/contrail/pkg/keystone"
	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/services"
	"github.com/labstack/echo"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	asfclient "github.com/Juniper/asf/pkg/client"
	asfkeystone "github.com/Juniper/asf/pkg/keystone"
)

// Dynamic proxy service related constants.
const (
	DefaultPath  = "proxy"
	SyncInterval = 2 * time.Second

	PrivateURLScope  = "private"
	PublicURLScope   = "public"
	XClusterIDKey    = "X-Cluster-ID"
	XServiceTokenKey = "X-Service-Token"

	limit         = 100
	pathSeparator = "/"
)

type endpointStore interface {
	RemoveDeleted(endpoints map[string]*models.Endpoint)
	Contains(prefix string) bool
	GetEndpointURLs(prefix string, endpointKey string) []string
	UpdateEndpoint(prefix string, endpoint *models.Endpoint) error
	InitScope(prefix string)
	GetEndpointURL(clusterID, prefix string) (string, bool)
}

// Dynamic provides dynamic HTTP and WebSocket proxy capabilities.
type Dynamic struct {
	es               endpointStore
	dbService        services.Service
	config           *Config
	stopService      context.CancelFunc
	serviceWaitGroup *sync.WaitGroup
	forceUpdateChan  chan chan struct{}
	log              *logrus.Entry
}

// Config configures the dynamic proxy service and middleware.
type Config struct {
	// Path is the path the proxy is served at.
	Path string
	// ServiceTokenEndpointPrefixes are prefixes of the endpoints that should have a service token injected into requests.
	ServiceTokenEndpointPrefixes []string
	// ServiceUserClientConfig is an HTTP config that allows logging in as the service user.
	ServiceUserClientConfig *asfclient.HTTPConfig
}

// NewDynamicFromViper creates a Dynamic proxy from the global Viper configuration.
func NewDynamicFromViper(es endpointStore, dbService services.Service) *Dynamic {
	return NewDynamic(es, dbService, ConfigFromViper())
}

// NewDynamic creates a new Dynamic proxy.
func NewDynamic(es endpointStore, dbService services.Service, config *Config) *Dynamic {
	return &Dynamic{
		es:              es,
		dbService:       dbService,
		config:          config,
		forceUpdateChan: make(chan chan struct{}),
		log:             logutil.NewLogger("dynamic-proxy"),
	}
}

// ConfigFromViper creates a Config from the global Viper configuration.
func ConfigFromViper() *Config {
	return &Config{
		Path:                         pathFromViper(),
		ServiceTokenEndpointPrefixes: viper.GetStringSlice("server.service_token_endpoint_prefixes"),
		ServiceUserClientConfig:      serviceUserClientConfigFromViper(),
	}
}

func pathFromViper() string {
	if path := viper.GetString("server.dynamic_proxy_path"); path != "" {
		return path
	}
	return DefaultPath
}

func serviceUserClientConfigFromViper() *asfclient.HTTPConfig {
	c := asfclient.LoadHTTPConfig()
	c.SetCredentials(
		viper.GetString("keystone.service_user.id"),
		viper.GetString("keystone.service_user.password"),
	)
	c.Scope = asfkeystone.NewScope(
		viper.GetString("keystone.service_user.domain_id"),
		viper.GetString("keystone.service_user.domain_name"),
		viper.GetString("keystone.service_user.project_id"),
		viper.GetString("keystone.service_user.project_name"),
	)
	return c
}

// RegisterHTTPAPI registers the proxy endpoint at the configured path.
func (p *Dynamic) RegisterHTTPAPI(r apiserver.HTTPRouter) {
	r.Group(
		p.config.Path,
		apiserver.WithMiddleware(p.middleware),
		apiserver.WithNoAuth(),
		apiserver.WithHomepageType(apiserver.ProxyEndpoint),
	)
}

// RegisterGRPCAPI does nothing.
func (p *Dynamic) RegisterGRPCAPI(r apiserver.GRPCRouter) {}

func (p *Dynamic) middleware(next apiserver.HandlerFunc) apiserver.HandlerFunc {
	return func(ctx echo.Context) error {
		r := ctx.Request()

		clusterID, err := p.clusterID(r.URL.Path)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		setClusterIDHeader(r, clusterID)

		pp := proxyPrefix(r.URL.Path, urlScope(r.URL.Path))
		scope := urlScope(r.URL.Path)
		r.URL.Path = withNoProxyPrefix(r.URL.Path, pp)

		if p.shouldInjectServiceToken(pp) {
			var token string
			token, err = p.obtainServiceToken(r.Context(), clusterID)
			if err != nil {
				p.log.WithError(err).Error("Failed to obtain service token - not adding it to the request")
			} else {
				setServiceTokenHeader(r, token)
			}
		}

		t, err := p.readTargets(scope, pp)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		if err = proxy.HandleRequest(ctx, t, p.log); err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("Failed to handle proxy request: %v", err))
		}

		return nil
	}
}

func proxyPrefix(urlPath string, scope string) string {
	prefixes := strings.Split(urlPath, pathSeparator)[:4]
	return strings.Join(append(prefixes, scope), pathSeparator)
}

func withNoProxyPrefix(path, pp string) string {
	return strings.TrimPrefix(path, strings.TrimSuffix(pp, PublicURLScope))
}

func (p *Dynamic) readTargets(scope string, proxyPrefix string) ([]string, error) {
	if !p.es.Contains(proxyPrefix) {
		return nil, errors.Errorf("target store not found for given proxy prefix: %v", proxyPrefix)
	}
	targets := p.es.GetEndpointURLs(scope, proxyPrefix)
	if len(targets) == 0 {
		return nil, errors.New("no targets in proxy target store")
	}
	return targets, nil
}

func urlScope(urlPath string) string {
	if strings.Contains(urlPath, PrivateURLScope) {
		return PrivateURLScope
	}
	return PublicURLScope
}

func setClusterIDHeader(r *http.Request, clusterID string) {
	r.Header.Set(XClusterIDKey, clusterID)
}

func (p *Dynamic) clusterID(url string) (string, error) {
	paths := strings.Split(url, pathSeparator)
	if len(paths) <= 3 || paths[1] != p.config.Path {
		return "", errors.Errorf("cluster ID not found in proxy URL: %v", url)
	}
	return paths[2], nil
}

func (p *Dynamic) shouldInjectServiceToken(proxyPrefix string) bool {
	paths := strings.Split(proxyPrefix, pathSeparator)
	if len(paths) < 4 {
		return false
	}
	endpointPrefix := paths[3]
	allowedPrefixes := p.config.ServiceTokenEndpointPrefixes
	return format.ContainsString(allowedPrefixes, endpointPrefix)
}

func (p *Dynamic) obtainServiceToken(ctx context.Context, clusterID string) (string, error) {
	apiClient := client.NewHTTP(p.config.ServiceUserClientConfig)

	ctx = keystone.WithXClusterID(ctx, clusterID)
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
func (p *Dynamic) StartEndpointsSync() {
	serviceCtx, cancel := context.WithCancel(context.Background())
	p.stopService = cancel

	p.serviceWaitGroup = &sync.WaitGroup{}
	p.serviceWaitGroup.Add(1)

	go func() {
		defer p.serviceWaitGroup.Done()
		ticker := time.NewTicker(SyncInterval)
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

func (p *Dynamic) updateEndpoints() {
	endpoints, err := p.readEndpoints()
	if err != nil {
		p.log.WithError(err).Error("Endpoints read failed")
		return
	}
	p.syncEndpoints(endpoints)
}

func (p *Dynamic) readEndpoints() (map[string]*models.Endpoint, error) {
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

func (p *Dynamic) syncEndpoints(endpoints map[string]*models.Endpoint) {
	// delete stale  endpoints in-memory
	p.es.RemoveDeleted(endpoints)
	// create/update proxy middleware
	for _, e := range endpoints {
		p.initTargetStore(e)
		if e.PublicURL != "" {
			p.manageEndpoint(e, PublicURLScope)
			p.manageEndpoint(e, PrivateURLScope)
		}
	}
}

func (p *Dynamic) initTargetStore(e *models.Endpoint) {
	if e.PublicURL != "" {
		proxyPrefix := p.getProxyPrefix(e, PublicURLScope)
		p.es.InitScope(proxyPrefix)

		privateProxyPrefix := p.getProxyPrefix(e, PrivateURLScope)
		p.es.InitScope(privateProxyPrefix)
	}
}

func (p *Dynamic) manageEndpoint(endpoint *models.Endpoint, scope string) {
	proxyPrefix := p.getProxyPrefix(endpoint, scope)

	err := p.es.UpdateEndpoint(proxyPrefix, endpoint)
	if err != nil {
		p.log.WithField("prefix", proxyPrefix).Error(err)
	}
}

func (p *Dynamic) getProxyPrefix(endpoint *models.Endpoint, scope string) (proxyPrefix string) {
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
	prefixes := []string{"", p.config.Path, endpoint.ParentUUID, prefix, scope}
	return strings.Join(prefixes, pathSeparator)
}

// ForceUpdate requests an immediate update of endpoints and waits for its completion.
func (p *Dynamic) ForceUpdate() {
	wait := make(chan struct{})
	p.forceUpdateChan <- wait
	<-wait
}

// StopEndpointsSync stops synchronization of proxy endpoints.
func (p *Dynamic) StopEndpointsSync() {
	// stop proxy server poll
	p.stopService()
	// wait for the proxy server poll to complete
	p.serviceWaitGroup.Wait()
}
