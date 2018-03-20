package apisrv

import (
	"context"
	"crypto/tls"
	"net"
	"net/http"
	"net/http/httputil"
	"net/url"
	"path/filepath"
	"sync"
	"time"

	"github.com/labstack/echo"

	"github.com/Juniper/contrail/pkg/common"
	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/serviceif"

	log "github.com/sirupsen/logrus"
)

const (
	public  = "public"
	private = "private"
)

// ProxyServer holds info about the proxy server and endpoint store
type ProxyServer struct {
	echoServer    *echo.Echo
	dbService     serviceif.Service
	endpointStore *common.EndpointStore
}

// NewProxyServer creates new proxy server
func NewProxyServer(e *echo.Echo, endpointStore *common.EndpointStore, dbService serviceif.Service) *ProxyServer {
	p := &ProxyServer{
		dbService:     dbService,
		echoServer:    e,
		endpointStore: endpointStore,
	}
	return p
}

func (p *ProxyServer) readEndpoints() (map[string]*models.Endpoint, error) {
	ctx := common.NoAuth(context.Background())
	spec := models.ListSpec{}
	request := &models.ListEndpointRequest{Spec: &spec}
	response, err := p.dbService.ListEndpoint(ctx, request)
	if err != nil {
		return nil, err
	}
	var endpoints map[string]*models.Endpoint
	for _, e := range response.Endpoints {
		endpoints[e.UUID] = e
	}
	return endpoints, nil
}

func (p *ProxyServer) dynamicProxyMiddleware(id string, urlType string) func(next echo.HandlerFunc) echo.HandlerFunc {
	e, ok := p.endpointStore.Data.Load(id)
	if !ok {
		return nil
	}
	ep, ok := e.(*models.Endpoint)
	var target string
	if urlType == private {
		target = ep.PrivateURL
	} else {
		target = ep.PublicURL
	}
	insecure := true          //TODO:(ijohnson) add insecure to endpoint schema
	u, _ := url.Parse(target) // nolint
	server := httputil.NewSingleHostReverseProxy(u)
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
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			r := c.Request()
			w := c.Response()
			server.ServeHTTP(w, r)
			return nil
		}
	}
}

func (p *ProxyServer) createProxyMiddleware(clusterID string, endpointID string, name string) {
	// proxy public url
	publicPrefix := filepath.Join("/", clusterID, name)
	g := p.echoServer.Group(publicPrefix)
	g.Use(removePathPrefixMiddleware(publicPrefix))
	g.Use(p.dynamicProxyMiddleware(endpointID, public))

	// proxy private url
	privatePrefix := filepath.Join("/", clusterID, name, private)
	g = p.echoServer.Group(privatePrefix)
	g.Use(removePathPrefixMiddleware(privatePrefix))
	g.Use(p.dynamicProxyMiddleware(endpointID, private))
}

func (p *ProxyServer) isMultiCluster(endpoints map[string]*models.Endpoint, clusterIDS map[string]bool) bool {
	// assume one cluster, no need to prefix proxy with  cluster id
	for _, endpoint := range endpoints {
		if _, ok := clusterIDS[endpoint.ParentUUID]; !ok {
			clusterIDS[endpoint.ParentUUID] = true
		}
		if len(clusterIDS) > 1 {
			// more than one cluster, need to prefix proxy with cluster id
			return true
		}
	}
	return false
}

func (p *ProxyServer) isMultiClusterNow() bool {
	var clusterIDS map[string]bool
	var multiCluster bool
	p.endpointStore.Data.Range(func(id, endpoint interface{}) bool {
		e, _ := endpoint.(*models.Endpoint)
		if _, ok := clusterIDS[e.ParentUUID]; !ok {
			clusterIDS[e.ParentUUID] = true
		}
		if len(clusterIDS) > 1 {
			// more than one cluster, need to prefix proxy with cluster id
			multiCluster = true
			return true
		}
		multiCluster = false
		return false
	})
	return multiCluster
}

func (p *ProxyServer) checkDeleteProxyMiddlewares(endpoints map[string]*models.Endpoint) {
	wg := &sync.WaitGroup{}
	p.endpointStore.Data.Range(func(id, endpoint interface{}) bool {
		wg.Add(1)
		// call N go routines to delete proxy middleware
		go func() {
			_, ok := endpoints[id.(string)]
			if !ok {
				p.endpointStore.Data.Delete(id)
				// TODO:(ijohnson) Find a way to remove the middleware from echo server
			}
		}()
		wg.Wait()
		return true
	})
}

func (p *ProxyServer) manageProxyMiddlewares(endpoints map[string]*models.Endpoint, multiCluster bool) {
	wg := &sync.WaitGroup{}
	for _, endpoint := range endpoints {
		wg.Add(1)
		// call N go routines to create/update proxy middleware
		go func() {
			ep, ok := p.endpointStore.Data.Load(endpoint.UUID)
			if !ok {
				log.Fatalf("Unable to load endpoint(%s) data from memory store", endpoint.UUID)
				wg.Done()
				return
			}
			var clusterID string
			if multiCluster {
				clusterID = endpoint.ParentUUID
			} else {
				clusterID = ""
			}
			// check whether multi cluster proxy served already
			alreadyMultiCluster := p.isMultiClusterNow()
			e, ok := ep.(*models.Endpoint)
			if !ok { // proxy endpoint not in memory store
				// store and create proxy middleware
				p.endpointStore.Data.Store(endpoint.UUID, &endpoint)
				p.createProxyMiddleware(clusterID, endpoint.UUID, endpoint.Name)
			} else if e.Name != endpoint.Name { // proxy endpoint name changed
				// store and create proxy middleware
				p.endpointStore.Data.Store(endpoint.UUID, &endpoint)
				p.createProxyMiddleware(clusterID, endpoint.UUID, endpoint.Name)
			} else if !alreadyMultiCluster && multiCluster { // multi cluster proxies, prefix proxy with cluster id
				// store and create proxy middleware with cluster id prefix
				p.endpointStore.Data.Store(endpoint.UUID, &endpoint)
				p.createProxyMiddleware(clusterID, endpoint.UUID, endpoint.Name)
			} else if alreadyMultiCluster && !multiCluster { // single cluster proxies, remove cluster id prefix from proxy
				// store and create proxy middleware without cluster id prefix
				p.endpointStore.Data.Store(endpoint.UUID, &endpoint)
				p.createProxyMiddleware(clusterID, endpoint.UUID, endpoint.Name)
			} else if e.PrivateURL != endpoint.PrivateURL ||
				e.PublicURL != endpoint.PublicURL {
				// update the endpoint store
				p.endpointStore.Data.Store(endpoint.UUID, &endpoint)
			}
			wg.Done()
		}()
		wg.Wait()
		// delete stale proxy middleware
		p.checkDeleteProxyMiddlewares(endpoints)
	}
}

func (p *ProxyServer) serve(taskCtx context.Context) {
	// serve forever
	var err error
	var multiCluster bool
	var clusterIDS map[string]bool
	var endpoints map[string]*models.Endpoint
	for {
		select {
		case <-taskCtx.Done():
			log.Info("stopping dynamic proxy server")
			return
		default:
			// poll db for the endpoint resource
			endpoints, err = p.readEndpoints()
			if err != nil {
				log.Error("not able to read endpoints")
				log.Fatal(err)
			}
			clusterIDS = map[string]bool{}
			// find if the cluster id should be a prefix to proxy
			multiCluster = p.isMultiCluster(endpoints, clusterIDS)
			// create/delete proxy middleware
			p.manageProxyMiddlewares(endpoints, multiCluster)
		}
		// poll at 5 seconds interval
		time.Sleep(5 * time.Second)
	}
}
