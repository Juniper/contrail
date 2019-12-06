package base

import (
	"strconv"
	"strings"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"

	"github.com/Juniper/asf/pkg/format"
	"github.com/Juniper/contrail/pkg/models"
)

const (
	pathSep              = ":"
	webSep               = "//"
	protocol             = "http"
	secureProtocol       = "https"
	config               = "config"
	analytics            = "telemetry"
	webui                = "nodejs"
	identity             = "keystone"
	nova                 = "compute"
	ironic               = "baremetal"
	glance               = "glance"
	swift                = "swift"
	appformix            = "appformix"
	xflow                = "xflow"
	defaultAdminUser     = "admin"
	defaultAdminPassword = "contrail123"
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
	appformix: "9000",
	xflow:     "7090",
}

// EndpointData is the representation of cluster endpoints.
type EndpointData struct {
	ClusterID   string
	ResManager  *ResourceManager
	ClusterData *Data
	Log         *logrus.Entry
}

func (e *EndpointData) endpointToURL(protocol, ip, port string) (endpointURL string) {
	return strings.Join([]string{protocol, webSep + ip, port}, pathSep)
}

// nolint: gocyclo
func (e *EndpointData) getPort(nodeIP, service string) string {
	c := e.ClusterData
	var nodePortMap map[string]interface{}
	switch service {
	case config:
		nodePortMap = c.getConfigNodePorts()
	case analytics:
		nodePortMap = c.getAnalyticsNodePorts()
	case webui:
		nodePortMap = c.getWebuiNodePorts()
	case appformix:
		nodePortMap = c.getAppformixControllerNodePorts()
	case identity, nova, glance, ironic:
		o := c.getOpenstackClusterData()
		nodePortMap = o.getOpenstackControlPorts()
	case swift:
		o := c.getOpenstackClusterData()
		nodePortMap = o.getOpenstackStoragePorts()
	}
	if nodePortMap != nil {
		if portConfigured, ok := nodePortMap[nodeIP]; ok {
			if port, ok := format.InterfaceToInt64Map(portConfigured)[service]; ok {
				if port != 0 {
					return strconv.FormatInt(port, 10)
				}
			}
		}
	}
	return portMap[service]

}

func (e *EndpointData) getkeystoneAdminCredential() (adminUser, adminPassword string) {
	return e.ClusterData.KeystoneAdminCredential()
}

func (e *EndpointData) getOpenstackEndpointNodes() (endpointNodes map[string][]string) {
	var k []*models.KeyValuePair
	if o := e.ClusterData.GetOpenstackClusterInfo(); o != nil {
		if g := o.GetKollaGlobals(); g != nil {
			k = g.GetKeyValuePair()
		}
	}
	endpointNodes = make(map[string][]string)
	for _, keyValuePair := range k {
		switch keyValuePair.Key {
		case "openstack_external_vip":
			endpointNodes[identity] = []string{keyValuePair.Value}
			endpointNodes[nova] = []string{keyValuePair.Value}
			endpointNodes[ironic] = []string{keyValuePair.Value}
			endpointNodes[glance] = []string{keyValuePair.Value}
			endpointNodes[swift] = []string{keyValuePair.Value}
		case "openstack_internal_vip":
			endpointNodes[identity] = []string{keyValuePair.Value}
			endpointNodes[nova] = []string{keyValuePair.Value}
			endpointNodes[ironic] = []string{keyValuePair.Value}
			endpointNodes[glance] = []string{keyValuePair.Value}
			endpointNodes[swift] = []string{keyValuePair.Value}
		}
	}
	if _, ok := endpointNodes[identity]; !ok {
		var openstackControlNodes []string
		o := e.ClusterData.getOpenstackClusterData()
		vip := o.getOpenstackPublicVip()
		if vip != "" {
			openstackControlNodes = []string{vip}
		} else {
			openstackControlNodes = o.getControlNodeIPs()
		}
		endpointNodes[identity] = openstackControlNodes
		endpointNodes[nova] = openstackControlNodes
		endpointNodes[ironic] = openstackControlNodes
		endpointNodes[glance] = openstackControlNodes
	}
	if _, ok := endpointNodes[swift]; !ok {
		var openstackStorageNodes []string
		o := e.ClusterData.getOpenstackClusterData()
		vip := o.getOpenstackPublicVip()
		if vip != "" {
			openstackStorageNodes = []string{vip}
		} else {
			openstackStorageNodes = o.getStorageNodeIPs()
		}
		endpointNodes[swift] = openstackStorageNodes
	}
	return endpointNodes
}

// nolint: gocyclo
func (e *EndpointData) getContrailEndpointNodes() (endpointNodes map[string][]string) {
	endpointNodes = make(map[string][]string)
	if c := e.ClusterData.ClusterInfo.GetContrailConfiguration(); c != nil {
		for _, keyValuePair := range c.GetKeyValuePair() {
			IPAddresses := strings.Split(keyValuePair.Value, ",")
			switch keyValuePair.Key {
			case "CONTROLLER_NODES":
				if _, ok := endpointNodes[config]; !ok {
					endpointNodes[config] = IPAddresses
				}
				if _, ok := endpointNodes[analytics]; !ok {
					endpointNodes[analytics] = IPAddresses
				}
				if _, ok := endpointNodes[webui]; !ok {
					endpointNodes[webui] = IPAddresses
				}
			case "CONFIG_NODES":
				endpointNodes[config] = IPAddresses
			case "ANALYTICS_NODES":
				endpointNodes[analytics] = IPAddresses
			case "WEBUI_NODES":
				endpointNodes[webui] = IPAddresses
			}
		}
	}
	vip := e.ClusterData.getContrailExternalVip()
	if vip != "" {
		endpointNodes[config] = []string{vip}
		endpointNodes[analytics] = []string{vip}
		endpointNodes[webui] = []string{vip}
	}
	if _, ok := endpointNodes[config]; !ok {
		endpointNodes[config] = e.ClusterData.getConfigNodeIPs()
	}
	if _, ok := endpointNodes[analytics]; !ok {
		endpointNodes[analytics] = e.ClusterData.getAnalyticsNodeIPs()
	}
	if _, ok := endpointNodes[webui]; !ok {
		endpointNodes[webui] = e.ClusterData.getWebuiNodeIPs()
	}
	return endpointNodes
}

func (e *EndpointData) getAppformixEndpointNodes() (endpointNodes map[string][]string) {
	endpointNodes = make(map[string][]string)
	endpointNodes[appformix] = e.ClusterData.getAppformixControllerNodeIPs()
	return endpointNodes
}

func getXflowEndpointAddress(xflowData *XflowData) (string, error) {
	if xflowData.ClusterInfo.TelemetryInBandManagementVip != "" {
		return xflowData.ClusterInfo.TelemetryInBandManagementVip, nil
	}
	if xflowData.ClusterInfo.KeepalivedSharedIP != "" {
		return xflowData.ClusterInfo.KeepalivedSharedIP, nil
	}
	return "", errors.New("failed to find xflow node IP Address")
}

func (e *EndpointData) getXflowEndpointNodes() (map[string][]string, error) {
	endpointNodes := make(map[string][]string)
	xflowData := e.ClusterData.GetXflowData()
	if xflowData != nil && xflowData.ClusterInfo != nil && len(xflowData.ClusterInfo.AppformixFlowsNodes) > 0 {
		xflowAddress, err := getXflowEndpointAddress(xflowData)
		if err != nil {
			return nil, err
		}
		endpointNodes[xflow] = []string{xflowAddress}
	}
	return endpointNodes, nil
}

// Create endpoint
func (e *EndpointData) Create() error { //nolint: gocyclo
	e.Log.Infof("Creating service endpoints for cluster: %s", e.ClusterID)
	contrailEndpoints := e.getContrailEndpointNodes()
	for service, endpointIPs := range contrailEndpoints {
		e.Log.Infof("Creating %s endpoints", service)
		for _, endpointIP := range endpointIPs {
			endpointProtocol := protocol
			switch service {
			case webui:
				endpointProtocol = secureProtocol
			case config:
				if e.ClusterData.isSSLEnabled() {
					endpointProtocol = secureProtocol
				}
			case analytics:
				if e.ClusterData.isSSLEnabled() {
					endpointProtocol = secureProtocol
				}
			}
			publicURL := e.endpointToURL(
				endpointProtocol, endpointIP, e.getPort(endpointIP, service))
			privateURL := publicURL
			endpointData := map[string]string{
				"parent_uuid": e.ClusterID,
				"name":        service,
				"public_url":  publicURL,
				"private_url": privateURL,
			}
			err := e.ResManager.createEndpoint(endpointData)
			if err != nil {
				return err
			}
		}
	}

	// openstack endpoints
	if e.ClusterData.ClusterInfo.Orchestrator == "openstack" {
		openstackEndpoints := e.getOpenstackEndpointNodes()
		endpointProtocol := protocol
		if e.ClusterData.getOpenstackClusterData().isSSLEnabled() {
			endpointProtocol = secureProtocol
		}
		for service, endpointIPs := range openstackEndpoints {
			e.Log.Infof("Creating %s endpoints", service)
			for _, endpointIP := range endpointIPs {
				publicURL := e.endpointToURL(
					endpointProtocol, endpointIP, e.getPort(endpointIP, service))
				privateURL := publicURL
				endpointData := map[string]string{
					"parent_uuid": e.ClusterID,
					"name":        service,
					"public_url":  publicURL,
					"private_url": privateURL,
				}
				if service == identity {
					adminUser, adminPassword := e.getkeystoneAdminCredential()
					endpointData["username"] = adminUser
					endpointData["password"] = adminPassword
				}
				err := e.ResManager.createEndpoint(endpointData)
				if err != nil {
					return err
				}
			}
		}
	}

	// appformix and xflow endpoints
	xflowEndpoints, err := e.getXflowEndpointNodes()
	if err != nil {
		return err
	}
	endpoints := format.MergeMultimap(e.getAppformixEndpointNodes(), xflowEndpoints)
	for service, endpointIPs := range endpoints {
		e.Log.Infof("Creating %s endpoints:%s", service, endpointIPs)
		for _, endpointIP := range endpointIPs {
			endpointProtocol := protocol
			publicURL := e.endpointToURL(
				endpointProtocol, endpointIP, e.getPort(endpointIP, service))
			privateURL := publicURL
			endpointData := map[string]string{
				"parent_uuid": e.ClusterID,
				"name":        service,
				"public_url":  publicURL,
				"private_url": privateURL,
			}
			err := e.ResManager.createEndpoint(endpointData)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

// Update endpoint
func (e *EndpointData) Update() error {
	e.Log.Infof("Updating service endpoints for cluster: %s", e.ClusterID)
	err := e.remove(e.getEndpointPrefixes())
	if err != nil {
		return err
	}
	err = e.Create()
	return err
}

func (e *EndpointData) getEndpointPrefixes() []string {
	var prefixes []string
	for k := range e.getContrailEndpointNodes() {
		prefixes = append(prefixes, k)
	}
	if e.ClusterData.ClusterInfo.Orchestrator == "openstack" {
		for k := range e.getOpenstackEndpointNodes() {
			prefixes = append(prefixes, k)
		}
	}
	// appformix and xflow endpoints
	xflowEndpoints, err := e.getXflowEndpointNodes()
	if err != nil {
		for k := range xflowEndpoints {
			prefixes = append(prefixes, k)
		}
	}
	for k, v := range e.getAppformixEndpointNodes() {
		if v != nil {
			prefixes = append(prefixes, k)
		}
	}
	return prefixes
}

// Remove endpoint
func (e *EndpointData) Remove() error {
	return e.remove(format.GetKeys(portMap))
}

func (e *EndpointData) remove(prefixes []string) error {
	e.Log.WithFields(logrus.Fields{
		"prefixes":     prefixes,
		"cluster-uuid": e.ClusterID,
	}).Info("Deleting service endpoints for cluster")
	endpointIDs, err := e.ResManager.getEndpoints([]string{e.ClusterID}, prefixes)
	if err != nil {
		return err
	}
	for _, endpointID := range endpointIDs {
		err = e.ResManager.deleteEndpoint(endpointID)
		if err != nil {
			return err
		}
	}
	return nil
}
