package keystone

import (
	"context"
	"database/sql"
	"fmt"
	"path/filepath"
	"strings"
	"sync"

	"github.com/labstack/echo"

	"github.com/Juniper/contrail/pkg/models"
)

type proxyEndpoint struct {
	name       string
	privateURL string
	publicURL  string
}

type proxyServer struct {
	echoServer    *echo.Echo
	dbConn        *sql.DB
	endpointStore *EndpointStore
}

func (p *proxyServer) readEndpoints() (*map[string]models.Endpoint, error) {
	return nil, nil
}

func (p *proxyServer) createProxyMiddleware(id string, name string, target string) error {
	prefix := filepath.Join("/", id, name)
	g := p.echoServer.Group(prefix)
	g.Use(removePathPrefixMiddleware(prefix))
	g.Use(proxyMiddleware(target, true)) //TODO:(ijohnson) add insecure to endpoint schema
}

func (p *proxyServer) checkIDPrefix(endpoints *map[string]models.Endpoint) bool {
	// assume one cluster, no need to prefix proxy with  cluster id
	idPrefix := false
	for _, endpoint := range endpoints {
		if _, ok := clusterIDS[endpoint.ParentUUID]; ok {
			continue
		} else {
			clusterIDS[endpoint.ParentUUID] = true
		}
		if len(clusterIDS) > 1 {
			// more than one cluster, need to prefix proxy with cluster id
			idPrefix = true
			break
		}
	}
	return idPrefix
}

func (p *proxyServer) checkDeleteProxyMiddlewares(endpoints *map[string]models.Endpoint) bool {
	p.endpointStore.store.Range(func(id, endpoint interface{}) bool {
		// call N go routines to delete proxy middleware
		go func() {
			_, ok := endpoints[id]
			if !ok {
				p.endpointStore.store.Delete(id)
				// TODO:(ijohnson) Find a way to remove the middleware from echo server
			}
		}()
		return true
	})

}

func (p *proxyServer) manageProxyMiddlewares(endpoints *map[string]models.Endpoint, idPrefixRequired bool) bool {
	for _, endpoint := range endpoints {
		// call N go routines to create/update proxy middleware
		go func() {
			e, ok := p.endpointStore.store.Load(id)
			if !ok || // proxy endpoint not in memory store
				e.name != endpoint.Name || // proxy endpoint name changed
				(e.clusterID == "" && idPrefix) || // clusters added, prefix proxy with cluster id
				(e.clusterID != "" && !idPrefix) { // clusters deleted, remove cluster id prefix from proxy
				// store and create proxy middleware
				if idPrefixRequired {
					clusterID := endpoint.ParentUUID
				} else {
					clusterID := ""
				}
				ep := &proxyEndpoint{
					clusterID:  clusterID,
					name:       endpoint.Name,
					privateURL: endpoint.PrivateURL,
					publicURL:  endpoint.PublicURL,
				}
				p.endpointStore.store.Store(id, ep)
				p.createProxyMiddleware(ep.clusterID, ep.name, ep.privateURL)
				p.createProxyMiddleware(ep.clusterID, ep.name, ep.publicURL)
			} else if e.privateURL != endpoint.privateURL {
				p.createProxyMiddleware(ep.clusterID, ep.name, ep.privateURL)
			} else if e.publicURL != endpoint.publicURL {
				p.createProxyMiddleware(ep.clusterID, ep.name, ep.publicURL)
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
		idPrefixRequired := p.checkIDPrefix(&endpoints)
		// create/delete proxy middleware
		p.manageProxyMiddlewares(&endpoints, idPrefixRequired)
	}
}
