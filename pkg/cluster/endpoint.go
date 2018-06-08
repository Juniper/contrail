package cluster

import (
	"strings"

	"github.com/sirupsen/logrus"

	"github.com/Juniper/contrail/pkg/models"
)

const (
	pathSep        = ":"
	webSep         = "//"
	protocol       = "http"
	secureProtocol = "https"
	config         = "config"
	analytics      = "telemetry"
	webui          = "nodejs"
	identity       = "keystone"
	nova           = "compute"
	ironic         = "baremetal"
	glance         = "glance"
	swift          = "swift"
)

var portMap = map[string]string{
	config:    "8082",
	analytics: "8081",
	webui:     "8143",
	identity:  "5000",
	nova:      "8774",
	ironic:    "6385",
	glance:    "9292",
	swift:     "8080",
}

// EndpointData is the representation of cluster endpoints.
type EndpointData struct {
	clusterID   string
	cluster     *Cluster
	clusterData *Data
	log         *logrus.Entry
}

func (e *EndpointData) endpointToURL(protocol, ip, port string) (endpointURL string) {
	return strings.Join([]string{protocol, webSep + ip, port}, pathSep)
}

func (e *EndpointData) getOpenstackEndpointNodes() (endpointNodes map[string][]string) {
	var k []*models.KeyValuePair
	if o := e.clusterData.getOpenstackClusterInfo(); o == nil {
		if g := o.GetKollaGlobals(); g != nil {
			k = g.GetKeyValuePair()
		}
	}
	endpointNodes = make(map[string][]string)
	if k != nil {
		for _, keyValuePair := range k {
			switch keyValuePair.Key {
			case "contrail_external_vip":
				endpointNodes[identity] = []string{keyValuePair.Value}
				endpointNodes[nova] = []string{keyValuePair.Value}
				endpointNodes[ironic] = []string{keyValuePair.Value}
				endpointNodes[glance] = []string{keyValuePair.Value}
				endpointNodes[swift] = []string{keyValuePair.Value}
				break
			case "contrail_internal_vip":
				endpointNodes[identity] = []string{keyValuePair.Value}
				endpointNodes[nova] = []string{keyValuePair.Value}
				endpointNodes[ironic] = []string{keyValuePair.Value}
				endpointNodes[glance] = []string{keyValuePair.Value}
				endpointNodes[swift] = []string{keyValuePair.Value}
				break
			}
		}
	}
	if _, ok := endpointNodes[identity]; !ok {
		openstackControlNodes := e.clusterData.getOpenstackClusterData().getControlNodeIPs()
		endpointNodes[identity] = openstackControlNodes
		endpointNodes[nova] = openstackControlNodes
		endpointNodes[ironic] = openstackControlNodes
		endpointNodes[glance] = openstackControlNodes
	}
	if _, ok := endpointNodes[swift]; !ok {
		endpointNodes[swift] = e.clusterData.getOpenstackClusterData().getStorageNodeIPs()
	}
	return endpointNodes
}

func (e *EndpointData) getContrailEndpointNodes() (endpointNodes map[string][]string) {
	endpointNodes = make(map[string][]string)
	if c := e.clusterData.clusterInfo.GetContrailConfiguration(); c != nil {
		if k := c.GetKeyValuePair(); k != nil {
			for _, keyValuePair := range k {
				IPAddresses := strings.Split(keyValuePair.Value, ",")
				switch keyValuePair.Key {
				case "CONTROLLER_NODES":
					endpointNodes[config] = IPAddresses
					endpointNodes[analytics] = IPAddresses
					endpointNodes[webui] = IPAddresses
				case "CONFIG_NODES":
					endpointNodes[config] = IPAddresses
				case "ANALYTICS_NODES":
					endpointNodes[analytics] = IPAddresses
				case "WEBUI_NODES":
					endpointNodes[webui] = IPAddresses
				}
			}
		}
	}
	if _, ok := endpointNodes[config]; !ok {
		endpointNodes[config] = e.clusterData.getConfigNodeIPs()
	}
	if _, ok := endpointNodes[analytics]; !ok {
		endpointNodes[analytics] = e.clusterData.getAnalyticsNodeIPs()
	}
	if _, ok := endpointNodes[webui]; !ok {
		endpointNodes[webui] = e.clusterData.getWebuiNodeIPs()
	}
	return endpointNodes
}

func (e *EndpointData) create() error {
	e.log.Infof("Creating service endpoints for cluster: %s", e.clusterID)
	contrailEndpoints := e.getContrailEndpointNodes()
	for service, endpointIPs := range contrailEndpoints {
		e.log.Infof("Creating %s endpoints", service)
		for _, endpointIP := range endpointIPs {
			endpointProtocol := protocol
			if service == webui {
				endpointProtocol = secureProtocol
			}
			publicURL := e.endpointToURL(endpointProtocol, endpointIP, portMap[service])
			privateURL := publicURL
			err := e.cluster.createEndpoint(e.clusterID, service, publicURL, privateURL)
			if err != nil {
				return err
			}
		}
	}

	// openstack endpoints
	openstackEndpoints := e.getOpenstackEndpointNodes()
	for service, endpointIPs := range openstackEndpoints {
		e.log.Infof("Creating %s endpoints", service)
		for _, endpointIP := range endpointIPs {
			publicURL := e.endpointToURL(protocol, endpointIP, portMap[service])
			privateURL := publicURL
			err := e.cluster.createEndpoint(e.clusterID, service, publicURL, privateURL)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (e *EndpointData) update() error {
	e.log.Infof("Updating service endpoints for cluster: %s", e.clusterID)
	endpointIDs, err := e.cluster.getEndpoints([]string{e.clusterID})
	if err != nil {
		return err
	}
	for _, endpointID := range endpointIDs {
		err = e.cluster.deleteEndpoint(endpointID)
		if err != nil {
			return err
		}
	}
	err = e.create()
	return err
}

func (e *EndpointData) remove() error {
	e.log.Infof("Deleting service endpoints for cluster: %s", e.clusterID)
	endpointIDs, err := e.cluster.getEndpoints([]string{e.clusterID})
	if err != nil {
		return err
	}
	for _, endpointID := range endpointIDs {
		err = e.cluster.deleteEndpoint(endpointID)
		if err != nil {
			return err
		}
	}
	return nil
}
