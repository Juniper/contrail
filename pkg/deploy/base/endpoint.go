package base

import (
	"strconv"
	"strings"

	"github.com/sirupsen/logrus"

	"github.com/Juniper/contrail/pkg/format"
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
	xflow:     "8090",
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

func (e *EndpointData) getOpenstackPublicVip() (vip string) {
	vip = ""
	o := e.ClusterData.getOpenstackClusterData()
	if o.ClusterInfo.OpenstackExternalVip != "" {
		vip = o.ClusterInfo.OpenstackExternalVip
	} else if o.ClusterInfo.OpenstackInternalVip != "" {
		vip = o.ClusterInfo.OpenstackInternalVip
	}

	return vip
}

func (e *EndpointData) getkeystoneAdminCredential() (adminUser, adminPassword string) {
	var k []*models.KeyValuePair
	if o := e.ClusterData.GetOpenstackClusterInfo(); o != nil {
		if g := o.GetKollaPasswords(); g != nil {
			k = g.GetKeyValuePair()
			for _, keyValuePair := range k {
				switch keyValuePair.Key {
				case "keystone_admin_user":
					adminUser = keyValuePair.Value
				case "keystone_admin_password":
					adminPassword = keyValuePair.Value
				}
			}
		}
	}

	if adminUser == "" {
		adminUser = defaultAdminUser
	}
	if adminPassword == "" {
		adminPassword = defaultAdminPassword
	}

	return adminUser, adminPassword
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
		vip := e.getOpenstackPublicVip()
		if vip != "" {
			openstackControlNodes = []string{vip}
		} else {
			o := e.ClusterData.getOpenstackClusterData()
			openstackControlNodes = o.getControlNodeIPs()
		}
		endpointNodes[identity] = openstackControlNodes
		endpointNodes[nova] = openstackControlNodes
		endpointNodes[ironic] = openstackControlNodes
		endpointNodes[glance] = openstackControlNodes
	}
	if _, ok := endpointNodes[swift]; !ok {
		var openstackStorageNodes []string
		vip := e.getOpenstackPublicVip()
		if vip != "" {
			openstackStorageNodes = []string{vip}
		} else {
			o := e.ClusterData.getOpenstackClusterData()
			openstackStorageNodes = o.getStorageNodeIPs()
		}
		endpointNodes[swift] = openstackStorageNodes
	}
	return endpointNodes
}

func (e *EndpointData) isSSLEnabled() bool {
	if c := e.ClusterData.ClusterInfo.GetContrailConfiguration(); c != nil {
		for _, keyValuePair := range c.GetKeyValuePair() {
			switch keyValuePair.Key {
			case "SSL_ENABLE":
				yamlBoolTrue := []string{"yes", "true", "y", "on"}
				if format.ContainsString(
					yamlBoolTrue,
					strings.ToLower(keyValuePair.Value)) {
					return true
				}
			}
		}
	}
	return false
}

// nolint: gocyclo
func (e *EndpointData) getContrailEndpointNodes() (endpointNodes map[string][]string) {
	endpointNodes = make(map[string][]string)
	if c := e.ClusterData.ClusterInfo.GetContrailConfiguration(); c != nil {
		for _, keyValuePair := range c.GetKeyValuePair() {
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
	if a := e.ClusterData.ClusterInfo.GetAnnotations(); a != nil {
		for _, keyValuePair := range a.GetKeyValuePair() {
			switch keyValuePair.Key {
			case "contrail_external_vip":
				endpointNodes[config] = []string{keyValuePair.Value}
				endpointNodes[analytics] = []string{keyValuePair.Value}
				endpointNodes[webui] = []string{keyValuePair.Value}
			}
		}
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

func (e *EndpointData) getXflowEndpointNodes() (endpointNodes map[string][]string) {
	endpointNodes = make(map[string][]string)
	xflowData := e.ClusterData.GetXflowData()
	if xflowData != nil && xflowData.ClusterInfo != nil {
		endpointNodes[xflow] = []string{xflowData.ClusterInfo.KeepalivedSharedIP}
	}
	return endpointNodes
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
				if e.isSSLEnabled() {
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
		for service, endpointIPs := range openstackEndpoints {
			e.Log.Infof("Creating %s endpoints", service)
			for _, endpointIP := range endpointIPs {
				publicURL := e.endpointToURL(
					protocol, endpointIP, e.getPort(endpointIP, service))
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
	endpoints := format.MergeMultimap(e.getAppformixEndpointNodes(), e.getXflowEndpointNodes())
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
	err := e.Remove()
	if err != nil {
		return err
	}
	err = e.Create()
	return err
}

// Remove endpoint
func (e *EndpointData) Remove() error {
	e.Log.Infof("Deleting service endpoints for cluster: %s", e.ClusterID)
	endpointIDs, err := e.ResManager.getEndpoints([]string{e.ClusterID})
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
