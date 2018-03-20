package apisrv

import (
	"crypto/tls"
	"database/sql"
	"net"
	"net/http"
	"net/http/httputil"
	"net/url"
	"path/filepath"
	"time"

	"github.com/labstack/echo"

	"github.com/Juniper/contrail/pkg/common"
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

/*
func (p *proxyServer) readEndpoints() (*map[string]models.Endpoint, error) {
	q := "SELECT * FROM endpoint"
	rows, err := p.dbConn.Query(q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if err := rows.Err(); err != nil {
		return nil, err
	}
	var endpoints *map[string]models.Endpoint
	for rows.Next() {
		if err := rows.Scan(); err != nil {
			return nil, err
		}
		endpoints = append(endpoints, i)
	}
	return endpoints, nil
}
*/

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

func (p *proxyServer) checkIDPrefix(endpoints *map[string]models.Endpoint, clusterIDS *map[string]bool) bool {
	// assume one cluster, no need to prefix proxy with  cluster id
	idPrefix := false
	for _, endpoint := range *endpoints {
		if _, ok := *clusterIDS[endpoint.ParentUUID]; ok {
			continue
		} else {
			*clusterIDS[endpoint.ParentUUID] = true
		}
		if len(*clusterIDS) > 1 {
			// more than one cluster, need to prefix proxy with cluster id
			idPrefix = true
			break
		}
	}
	return idPrefix
}

func (p *proxyServer) checkDeleteProxyMiddlewares(endpoints *map[string]models.Endpoint) {
	p.endpointStore.Data.Range(func(id, endpoint interface{}) bool {
		// call N go routines to delete proxy middleware
		go func() {
			_, ok := *endpoints[id]
			if !ok {
				p.endpointStore.Data.Delete(id)
				// TODO:(ijohnson) Find a way to remove the middleware from echo server
			}
		}()
		return true
	})
}

func (p *proxyServer) manageProxyMiddlewares(endpoints *map[string]models.Endpoint, idPrefix bool) bool {
	for _, endpoint := range *endpoints {
		// call N go routines to create/update proxy middleware
		go func() {
			e, ok := p.endpointStore.Data.Load(endpoint.UUID)
			if !ok || // proxy endpoint not in memory store
				e.Name != endpoint.Name || // proxy endpoint name changed
				(e.clusterID == "" && idPrefix) || // clusters added, prefix proxy with cluster id
				(e.clusterID != "" && !idPrefix) { // clusters deleted, remove cluster id prefix from proxy
				// store and create proxy middleware
				var clusterID string
				if idPrefix {
					clusterID = endpoint.ParentUUID
				} else {
					clusterID = ""
				}
				p.endpointStore.Data.Store(endpoint.UUID, &endpoint)
				p.createProxyMiddleware(clusterID, endpoint.UUID, endpoint.Name)
			} else if e.privateURL != endpoint.privateURL ||
				e.publicURL != endpoint.publicURL {
				// update the endpoint store
				p.endpointStore.Data.Store(endpoint.UUID, &endpoint)
			}
		}()
		// delete stale proxy middleware
		checkDeleteProxyMiddlewares(&endpoints)
	}
}

func (p *proxyServer) serve() {
	// serve forever
	for {
		// poll db for the endpoint resource
		endpoints := p.readEndpoints()
		clusterIDS := map[string]bool{}
		// find if the cluster id should be a prefix to proxy
		idPrefixRequired := p.checkIDPrefix(&endpoints, &clusterIDS)
		// create/delete proxy middleware
		p.manageProxyMiddlewares(&endpoints, idPrefixRequired)
	}
}
