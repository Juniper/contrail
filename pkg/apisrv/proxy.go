package apisrv

import (
	"context"
	"crypto/tls"
	"database/sql"
	"fmt"
	"net"
	"net/http"
	"net/http/httputil"
	"net/url"
	"path/filepath"
	"time"

	"github.com/labstack/echo"
	"github.com/spf13/viper"

	"github.com/Juniper/contrail/pkg/common"
	"github.com/Juniper/contrail/pkg/db"
	"github.com/Juniper/contrail/pkg/models"
)

const (
	public  = "public"
	private = "private"
)

type proxyEndpoint struct {
	clusterID map[string]*proxyEndpoint
}

type proxyServer struct {
	echoServer    *echo.Echo
	dbConn        *sql.DB
	endpointStore *common.EndpointStore
}

func (p *proxyServer) readEndpoints() (map[string]*models.Endpoint, error) {
	dbService := db.NewService(p.dbConn, viper.GetString("database.dialect"))
	ctx := common.NoAuth(context.Background())
	spec := models.ListSpec{}
	request := &models.ListEndpointRequest{Spec: &spec}
	response, err := dbService.ListEndpoint(ctx, request)
	if err != nil {
		return nil, err
	}
	var endpoints map[string]*models.Endpoint
	for _, e := range response.Endpoints {
		endpoints[e.UUID] = e
	}
	return endpoints, nil
}

func (p *proxyServer) dynamicProxyMiddleware(id string, urlType string) func(next echo.HandlerFunc) echo.HandlerFunc {
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

func (p *proxyServer) createProxyMiddleware(clusterID string, endpointID string, name string) error {
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

	return nil
}

func (p *proxyServer) isMultiCluster(endpoints map[string]*models.Endpoint, clusterIDS map[string]bool) bool {
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

func (p *proxyServer) isMultiClusterNow() bool {
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

func (p *proxyServer) checkDeleteProxyMiddlewares(endpoints map[string]*models.Endpoint) {
	p.endpointStore.Data.Range(func(id, endpoint interface{}) bool {
		// call N go routines to delete proxy middleware
		go func() {
			_, ok := endpoints[id.(string)]
			if !ok {
				p.endpointStore.Data.Delete(id)
				// TODO:(ijohnson) Find a way to remove the middleware from echo server
			}
		}()
		return true
	})
}

func (p *proxyServer) manageProxyMiddlewares(endpoints map[string]*models.Endpoint, multiCluster bool) error {
	for _, endpoint := range endpoints {
		// call N go routines to create/update proxy middleware
		go func() error {
			ep, ok := p.endpointStore.Data.Load(endpoint.UUID)
			if !ok {
				return fmt.Errorf("Unable to load endpoint(%s) data from memory store", endpoint.UUID)
			}
			var clusterID string
			if multiCluster {
				clusterID = endpoint.ParentUUID
			} else {
				clusterID = ""
			}
			e, ok := ep.(*models.Endpoint)
			if !ok { // proxy endpoint not in memory store
				// store and create proxy middleware
				p.endpointStore.Data.Store(endpoint.UUID, &endpoint)
				p.createProxyMiddleware(clusterID, endpoint.UUID, endpoint.Name)
				return nil
			}
			if e.Name != endpoint.Name { // proxy endpoint name changed
				// store and create proxy middleware
				p.endpointStore.Data.Store(endpoint.UUID, &endpoint)
				p.createProxyMiddleware(clusterID, endpoint.UUID, endpoint.Name)
				return nil
			}
			// check whether multi cluster proxy served already
			alreadyMultiCluster := p.isMultiClusterNow()
			if !alreadyMultiCluster && multiCluster { // multi cluster proxies, prefix proxy with cluster id
				// store and create proxy middleware with cluster id prefix
				p.endpointStore.Data.Store(endpoint.UUID, &endpoint)
				p.createProxyMiddleware(clusterID, endpoint.UUID, endpoint.Name)
			}
			if alreadyMultiCluster && !multiCluster { // single cluster proxies, remove cluster id prefix from proxy
				// store and create proxy middleware without cluster id prefix
				p.endpointStore.Data.Store(endpoint.UUID, &endpoint)
				p.createProxyMiddleware(clusterID, endpoint.UUID, endpoint.Name)
			} else if e.PrivateURL != endpoint.PrivateURL ||
				e.PublicURL != endpoint.PublicURL {
				// update the endpoint store
				p.endpointStore.Data.Store(endpoint.UUID, &endpoint)
			}
			return nil
		}()
		// delete stale proxy middleware
		p.checkDeleteProxyMiddlewares(endpoints)
	}
	return nil
}

func (p *proxyServer) serve() error {
	// serve forever
	for {
		// poll db for the endpoint resource
		endpoints, err := p.readEndpoints()
		if err != nil {
			return err
		}
		clusterIDS := map[string]bool{}
		// find if the cluster id should be a prefix to proxy
		multiCluster := p.isMultiCluster(endpoints, clusterIDS)
		// create/delete proxy middleware
		p.manageProxyMiddlewares(endpoints, multiCluster)
	}
	return nil
}
