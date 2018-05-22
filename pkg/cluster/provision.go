package cluster

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/flosch/pongo2"
	"github.com/sirupsen/logrus"

	pkglog "github.com/Juniper/contrail/pkg/log"
	"github.com/Juniper/contrail/pkg/models"
)

const (
	pathSep   = ":"
	webSep    = "//"
	protocol  = "http"
	config    = "config"
	analytics = "telemetry"
	webui     = "nodejs"
	identity  = "keystone"
	nova      = "compute"
	ironic    = "baremetal"
	glance    = "glance"
	swift     = "swift"
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

type provisioner interface {
	provision() error
}

type provisionCommon struct {
	cluster     *Cluster
	clusterID   string
	action      string
	clusterData *Data
	log         *logrus.Entry
	reporter    *Reporter
}

func (p *provisionCommon) getK8sClusterData() *KubernetesData {
	// One kubernetes cluster is the supported topology
	if len(p.clusterData.kubernetesClusterData) > 0 {
		return p.clusterData.kubernetesClusterData[0]
	}
	return nil
}

func (p *provisionCommon) getK8sClusterInfo() *models.KubernetesCluster {
	if p.getK8sClusterData() != nil {
		return p.getK8sClusterData().clusterInfo
	}
	return nil
}

func (p *provisionCommon) getOpenstackClusterData() *OpenstackData {
	// One openstack cluster is the supported topology
	if len(p.clusterData.openstackClusterData) > 0 {
		return p.clusterData.openstackClusterData[0]
	}
	return nil
}

func (p *provisionCommon) getOpenstackClusterInfo() *models.OpenstackCluster {
	if p.getOpenstackClusterData() != nil {
		return p.getOpenstackClusterData().clusterInfo
	}
	return nil
}

func (p *provisionCommon) getAllNodesInfo() []*models.Node {
	nodes := p.clusterData.nodesInfo
	if p.getOpenstackClusterData() != nil {
		nodes = append(nodes, p.getOpenstackClusterData().nodesInfo...)
	}
	if p.getK8sClusterData() != nil {
		nodes = append(nodes, p.getK8sClusterData().nodesInfo...)
	}

	var uniqueNodes []*models.Node
	m := make(map[string]bool)

	for _, node := range nodes {
		if _, ok := m[node.UUID]; !ok {
			m[node.UUID] = true
			uniqueNodes = append(uniqueNodes, node)
		}
	}

	return uniqueNodes
}

func (p *provisionCommon) isCreated() bool {
	state := p.clusterData.clusterInfo.ProvisioningState
	if p.action == "create" && (state == "NOSTATE" || state == "") {
		return false
	}
	p.log.Infof("Cluster %s already provisioned, STATE: %s", p.clusterID, state)
	return true
}
func (p *provisionCommon) getTemplateRoot() string {
	templateRoot := p.cluster.config.TemplateRoot
	if templateRoot == "" {
		templateRoot = defaultTemplateRoot
	}
	return templateRoot
}

func (p *provisionCommon) applyTemplate(templateSrc string, context map[string]interface{}) ([]byte, error) {
	template, err := pongo2.FromFile(templateSrc)
	if err != nil {
		return nil, err
	}
	output, err := template.ExecuteBytes(context)
	if err != nil {
		return nil, err
	}
	return output, nil
}

func (p *provisionCommon) appendToFile(path string, content []byte) error {
	err := os.MkdirAll(filepath.Dir(path), os.ModePerm)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(path, content, os.FileMode(0600))
	return err
}

func (p *provisionCommon) getClusterHomeDir() string {
	dir := filepath.Join(defaultWorkRoot, p.clusterID)
	return dir
}

func (p *provisionCommon) getWorkingDir() string {
	dir := filepath.Join(p.getClusterHomeDir())
	return dir
}

func (p *provisionCommon) createWorkingDir() error {
	return os.MkdirAll(p.getWorkingDir(), os.ModePerm)
}

func (p *provisionCommon) deleteWorkingDir() error {
	return os.RemoveAll(p.getClusterHomeDir())
}

func (p *provisionCommon) endpointToURL(protocol, ip, port string) (endpointURL string) {
	return strings.Join([]string{protocol, webSep + ip, port}, pathSep)
}

func (p *provisionCommon) createEndpoints() error {
	p.log.Infof("Creating service endpoints for cluster: %s", p.clusterID)
	contrailCluster := p.clusterData.clusterInfo
	openstackCluster := p.getOpenstackClusterData().clusterInfo
	for _, configNode := range contrailCluster.ContrailConfigNodes {
		for _, nodeRef := range configNode.NodeRefs {
			for _, node := range p.clusterData.nodesInfo {
				if nodeRef.UUID == node.UUID {
					publicURL := p.endpointToURL(protocol, node.IPAddress, portMap[config])
					privateURL := publicURL
					err := p.cluster.createEndpoint(p.clusterID, config, publicURL, privateURL)
					if err != nil {
						return err
					}
				}
			}
		}
	}
	for _, analyticsNode := range contrailCluster.ContrailAnalyticsNodes {
		for _, nodeRef := range analyticsNode.NodeRefs {
			for _, node := range p.clusterData.nodesInfo {
				if nodeRef.UUID == node.UUID {
					publicURL := p.endpointToURL(
						protocol, node.IPAddress, portMap[analytics])
					privateURL := publicURL
					err := p.cluster.createEndpoint(p.clusterID, analytics, publicURL, privateURL)
					if err != nil {
						return err
					}
				}
			}
		}
	}
	for _, webuiNode := range contrailCluster.ContrailWebuiNodes {
		for _, nodeRef := range webuiNode.NodeRefs {
			for _, node := range p.clusterData.nodesInfo {
				if nodeRef.UUID == node.UUID {
					publicURL := p.endpointToURL(
						protocol, node.IPAddress, portMap[webui])
					privateURL := publicURL
					err := p.cluster.createEndpoint(p.clusterID, webui, publicURL, privateURL)
					if err != nil {
						return err
					}
				}
			}
		}
	}
	for _, openstackStorageNode := range openstackCluster.OpenstackStorageNodes {
		for _, nodeRef := range openstackStorageNode.NodeRefs {
			for _, node := range p.clusterData.nodesInfo {
				if nodeRef.UUID == node.UUID {
					publicURL := p.endpointToURL(
						protocol, node.IPAddress, portMap[swift])
					privateURL := publicURL
					err := p.cluster.createEndpoint(p.clusterID, swift, publicURL, privateURL)
					if err != nil {
						return err
					}
				}
			}
		}
	}
	for _, openstackControlNode := range openstackCluster.OpenstackControlNodes {
		for _, nodeRef := range openstackControlNode.NodeRefs {
			for _, node := range p.clusterData.nodesInfo {
				if nodeRef.UUID == node.UUID {
					publicURL := p.endpointToURL(
						protocol, node.IPAddress, portMap[nova])
					privateURL := publicURL
					err := p.cluster.createEndpoint(p.clusterID, nova, publicURL, privateURL)
					if err != nil {
						return err
					}
					publicURL = p.endpointToURL(
						protocol, node.IPAddress, portMap[ironic])
					privateURL = publicURL
					err = p.cluster.createEndpoint(p.clusterID, ironic, publicURL, privateURL)
					if err != nil {
						return err
					}
					publicURL = p.endpointToURL(
						protocol, node.IPAddress, portMap[glance])
					privateURL = publicURL
					err = p.cluster.createEndpoint(p.clusterID, glance, publicURL, privateURL)
					if err != nil {
						return err
					}
					publicURL = p.endpointToURL(
						protocol, node.IPAddress, portMap[identity])
					privateURL = publicURL
					err = p.cluster.createEndpoint(p.clusterID, identity, publicURL, privateURL)
					if err != nil {
						return err
					}
				}
			}
		}
	}
	return nil
}

func (p *provisionCommon) updateEndpoints() error {
	p.log.Infof("Updating service endpoints for cluster: %s", p.clusterID)
	endpointIDs, err := p.cluster.getEndpoints([]string{p.clusterID})
	if err != nil {
		return err
	}
	for _, endpointID := range endpointIDs {
		err = p.cluster.deleteEndpoint(endpointID)
		if err != nil {
			return err
		}
	}
	err = p.createEndpoints()
	return err
}

func (p *provisionCommon) deleteEndpoints() error {
	p.log.Infof("Deleting service endpoints for cluster: %s", p.clusterID)
	endpointIDs, err := p.cluster.getEndpoints([]string{p.clusterID})
	if err != nil {
		return err
	}
	for _, endpointID := range endpointIDs {
		err = p.cluster.deleteEndpoint(endpointID)
		if err != nil {
			return err
		}
	}
	return nil
}

func (p *provisionCommon) execCmd(cmd string, args []string, dir string) error {
	cmdline := exec.Command(cmd, args...)
	if dir != "" {
		cmdline.Dir = dir
	}
	stdout, err := cmdline.StdoutPipe()
	if err != nil {
		return err
	}
	stderr, err := cmdline.StderrPipe()
	if err != nil {
		return err
	}
	if err := cmdline.Start(); err != nil {
		return err
	}
	// Report progress log periodically to stdout/db
	go p.reporter.reportLog(stdout)
	go p.reporter.reportLog(stderr)
	return cmdline.Wait()
}

func newAnsibleProvisioner(cluster *Cluster, cData *Data, clusterID string, action string) (provisioner, error) {
	// create logger for reporter
	logger := pkglog.NewLogger("reporter")
	pkglog.SetLogLevel(logger, cluster.config.LogLevel)

	r := &Reporter{
		api:      cluster.APIServer,
		resource: fmt.Sprintf("%s/%s", defaultResourcePath, clusterID),
		log:      logger,
	}

	// create logger for ansible provisioner
	logger = pkglog.NewLogger("ansible-provisioner")
	pkglog.SetLogLevel(logger, cluster.config.LogLevel)

	return &ansibleProvisioner{provisionCommon{
		cluster:     cluster,
		clusterID:   clusterID,
		action:      action,
		clusterData: cData,
		reporter:    r,
		log:         logger,
	}}, nil
}

func newHelmProvisioner(cluster *Cluster, cData *Data, clusterID string, action string) (provisioner, error) {
	// create logger for reporter
	logger := pkglog.NewLogger("reporter")
	pkglog.SetLogLevel(logger, cluster.config.LogLevel)

	r := &Reporter{
		api:      cluster.APIServer,
		resource: fmt.Sprintf("%s/%s", defaultResourcePath, clusterID),
		log:      logger,
	}

	// create logger for Helm provisioner
	logger = pkglog.NewLogger("helm-provisioner")
	pkglog.SetLogLevel(logger, cluster.config.LogLevel)

	return &helmProvisioner{provisionCommon{
		cluster:     cluster,
		clusterID:   clusterID,
		action:      action,
		clusterData: cData,
		reporter:    r,
		log:         logger,
	}}, nil
}

// Creates new provisioner based on the type
func newProvisioner(cluster *Cluster) (provisioner, error) {
	return newProvisionerByID(cluster, cluster.config.ClusterID, cluster.config.Action)
}

func newProvisionerByID(cluster *Cluster, clusterID string, action string) (provisioner, error) {
	var cData *Data
	var err error
	if action == "delete" {
		cData = &Data{}
	} else {
		cData, err = cluster.getClusterDetails(clusterID)
	}
	if err != nil {
		return nil, err
	}
	provisionerType := cluster.config.ProvisionerType
	if provisionerType == "" {
		provisionerType = defaultProvisioner
	}

	switch provisionerType {
	case "ansible":
		return newAnsibleProvisioner(cluster, cData, clusterID, action)
	case "helm":
		return newHelmProvisioner(cluster, cData, clusterID, action)
	}
	return nil, errors.New("unsupported provisioner type")
}
