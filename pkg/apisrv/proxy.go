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

type proxyCreateMessage struct {
	clusterID    string
	endpointID   string
	endpointName string
}

// ProxyServer holds info about the proxy server and endpoint store
type ProxyServer struct {
	echoServer         *echo.Echo
	dbService          serviceif.Service
	endpointStore      *common.EndpointStore
	multiCluster       bool
	proxyCreateChannel chan proxyCreateMessage
	// context to stop producing proxy messages
	serverContext     context.Context
	stopServerContext context.CancelFunc
	serverWaitGroup   *sync.WaitGroup
	// context to stop consuming proxy messages
	createContext     context.Context
	stopCreateContext context.CancelFunc
	createWaitGroup   *sync.WaitGroup
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
	spec := models.ListSpec{Limit: 100}
	request := &models.ListEndpointRequest{Spec: &spec}
	response, err := p.dbService.ListEndpoint(ctx, request)
	if err != nil {
		return nil, err
	}
	endpoints := make(map[string]*models.Endpoint)
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
	if !ok {
		log.Fatalf("Unable to serve proxy for endpoint(%s) as not able to read from in-memory store", id)
	}
	var target string
	if urlType == private {
		target = ep.PrivateURL
	} else {
		target = ep.PublicURL
	}
	log.Debugf("with target url: %s for the proxy", target)
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

func (p *ProxyServer) createProxyMiddleware() {
	for {
		select {
		case <-p.createContext.Done():
			log.Info("stopping dynamic proxy creator")
			return
		default:
			msg := <-p.proxyCreateChannel
			if msg.endpointName != "" {
				// proxy public url
				publicPrefix := filepath.Join("/", msg.clusterID, msg.endpointName)
				publicGroup := p.echoServer.Group(publicPrefix)
				publicGroup.Use(removePathPrefixMiddleware(publicPrefix))
				log.Debugf("creating dynamic proxy for public url: %s", publicPrefix)
				publicGroup.Use(p.dynamicProxyMiddleware(msg.endpointID, public))

				// proxy private url
				privatePrefix := filepath.Join("/", msg.clusterID, msg.endpointName, private)
				privateGroup := p.echoServer.Group(privatePrefix)
				privateGroup.Use(removePathPrefixMiddleware(privatePrefix))
				log.Debugf("creating dynamic proxy for private url: %s", privatePrefix)
				privateGroup.Use(p.dynamicProxyMiddleware(msg.endpointID, private))
			}
		}
	}
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
	var multiCluster bool
	clusterIDS := make(map[string]bool)
	p.endpointStore.Data.Range(func(id, endpoint interface{}) bool {
		e, _ := endpoint.(*models.Endpoint)
		if _, ok := clusterIDS[e.ParentUUID]; !ok {
			clusterIDS[e.ParentUUID] = true
		}
		if len(clusterIDS) > 1 {
			// more than one cluster, need to prefix proxy with cluster id
			multiCluster = true
			// return false to stop range function iteration
			return false
		}
		// return false to continue range function iteration
		multiCluster = false
		return true
	})
	return multiCluster
}

func (p *ProxyServer) checkDeleteProxyMiddlewares(endpoints map[string]*models.Endpoint) {
	wg := &sync.WaitGroup{}
	p.endpointStore.Data.Range(func(id, endpoint interface{}) bool {
		wg.Add(1)
		// call N go routines to delete proxy middleware
		go func() {
			defer wg.Done()
			_, ok := endpoints[id.(string)]
			if !ok {
				p.endpointStore.Data.Delete(id)
				// TODO:(ijohnson) Find a way to remove the middleware from echo server
				log.Debugf("deleting dynamic proxy for id: %s", id)
			}
		}()
		wg.Wait()
		return true
	})
}

func (p *ProxyServer) getClusterIDPrefix(endpoint *models.Endpoint) string {
	var clusterID string
	if p.multiCluster {
		clusterID = endpoint.ParentUUID
	} else {
		clusterID = ""
	}
	return clusterID
}

func (p *ProxyServer) storeAndSendCreateMessage(
	clusterID string, endpoint *models.Endpoint) {
	p.endpointStore.Data.Store(endpoint.UUID, endpoint)
	p.proxyCreateChannel <- proxyCreateMessage{
		clusterID:    clusterID,
		endpointID:   endpoint.UUID,
		endpointName: endpoint.Name,
	}
}

func (p *ProxyServer) endpointNameChanged(name string, newName string) bool {
	if name != newName { // proxy endpoint name changed
		log.Debugf("endpoint name changed from %s to %s", name, newName)
		return true
	}
	return false
}

func (p *ProxyServer) endpointURLChanged(url string, newURL string) bool {
	if url != newURL { // proxy endpoint url changed
		log.Debugf("endpoint url changed from %s to %s", url, newURL)
		return true
	}
	return false
}

func (p *ProxyServer) clusterChanged() bool {
	alreadyMultiCluster := p.isMultiClusterNow()
	if !alreadyMultiCluster && p.multiCluster { // multi cluster proxies
		log.Debug("multiple clusters found, add cluster-id to prefix")
		return true
	} else if alreadyMultiCluster && !p.multiCluster { // single cluster proxies
		log.Debugf("one cluster exist, remove cluster-id from prefix")
		return true
	}
	return false
}

func (p *ProxyServer) manageProxyMiddlewares(endpoints map[string]*models.Endpoint) {
	// delete stale proxy middleware
	p.checkDeleteProxyMiddlewares(endpoints)
	// create/update proxy middleware
	wg := &sync.WaitGroup{}
	for _, endpoint := range endpoints {
		wg.Add(1)
		// call N go routines to create/update proxy middleware
		go func(endpoint *models.Endpoint) {
			defer wg.Done()
			clusterID := p.getClusterIDPrefix(endpoint)
			ep, ok := p.endpointStore.Data.Load(endpoint.UUID)
			if ok { // proxy endpoint in memory store
				// check whether multi cluster proxy served already
				e, ok := ep.(*models.Endpoint)
				if !ok {
					log.Fatalf("Unable to load endpoint(%s) data from in-memory store",
						endpoint.UUID)
					return
				}
				if p.endpointNameChanged(e.Name, endpoint.Name) ||
					p.clusterChanged() {
					// store and modify proxy middleware
					p.storeAndSendCreateMessage(clusterID, endpoint)
				} else if p.endpointURLChanged(e.PrivateURL, endpoint.PrivateURL) ||
					p.endpointURLChanged(e.PublicURL, endpoint.PublicURL) {
					// update the endpoint store
					p.endpointStore.Data.Store(endpoint.UUID, endpoint)
				}
			} else { // proxy endpoint not in memory store
				// store and create proxy middleware
				log.Debugf("new endpoint(%s, %s) created",
					endpoint.Name, endpoint.UUID)
				p.storeAndSendCreateMessage(clusterID, endpoint)
			}
		}(endpoint)
	}
	wg.Wait()
}

func (p *ProxyServer) serve() {
	// create channel to send/receive proxy create messages
	p.createContext, p.stopCreateContext = context.WithCancel(context.Background())
	p.proxyCreateChannel = make(chan proxyCreateMessage)
	p.createWaitGroup = &sync.WaitGroup{}
	p.createWaitGroup.Add(1)
	go func() {
		p.createProxyMiddleware()
		p.createWaitGroup.Done()
	}()

	p.serverContext, p.stopServerContext = context.WithCancel(context.Background())
	p.serverWaitGroup = &sync.WaitGroup{}
	p.serverWaitGroup.Add(1)
	go func() {
		// serve forever
		defer p.serverWaitGroup.Done()
		var err error
		var clusterIDS map[string]bool
		var endpoints map[string]*models.Endpoint
		for {
			select {
			case <-p.serverContext.Done():
				log.Info("stopping dynamic proxy server")
				return
			default:
				// poll db for the endpoint resource
				endpoints, err = p.readEndpoints()
				if err != nil {
					log.Error("not able to read endpoints")
					log.Fatal(err)
				}
				if endpoints != nil {
					clusterIDS = map[string]bool{}
					// find if the cluster id should be a prefix to proxy
					p.multiCluster = p.isMultiCluster(endpoints, clusterIDS)
					// create/delete proxy middleware
					p.manageProxyMiddlewares(endpoints)
				}
			}
		}
	}()
}

func (p *ProxyServer) stop() {
	// stop producing proxy create messages
	p.stopServerContext()
	// wait for the producers complete
	p.serverWaitGroup.Wait()
	// close message channel
	close(p.proxyCreateChannel)
	// stop consuming proxy create messages
	p.stopCreateContext()
	// wait for the consumer to complete
	p.createWaitGroup.Wait()
}
