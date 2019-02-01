package cluster

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/pkg/errors"

	"github.com/Juniper/contrail/pkg/deploy/base"
	"github.com/Juniper/contrail/pkg/log"
	"github.com/Juniper/contrail/pkg/log/report"
	"github.com/Juniper/contrail/pkg/models"
)

const (
	defaultTemplateRoot = "./pkg/cluster/configs"
)

type deployCluster struct {
	base.Deploy
	cluster     *Cluster
	clusterID   string
	action      string
	clusterData *Data
}

func (p *deployCluster) isCreated() bool {
	state := p.clusterData.clusterInfo.ProvisioningState
	if p.action == "create" && (state == statusNoState || state == "") {
		return false
	}
	p.Log.Infof("Cluster %s already deployed, STATE: %s", p.clusterID, state)
	return true
}

func (p *deployCluster) getTemplateRoot() string {
	templateRoot := p.cluster.config.TemplateRoot
	if templateRoot == "" {
		templateRoot = defaultTemplateRoot
	}
	return templateRoot
}

func (p *deployCluster) getClusterHomeDir() string {
	dir := filepath.Join(defaultWorkRoot, p.clusterID)
	return dir
}

func (p *deployCluster) getWorkingDir() string {
	dir := filepath.Join(p.getClusterHomeDir())
	return dir
}

func (p *deployCluster) createWorkingDir() error {
	return os.MkdirAll(p.getWorkingDir(), os.ModePerm)
}

func (p *deployCluster) deleteWorkingDir() error {
	return os.RemoveAll(p.getClusterHomeDir())
}

func (p *deployCluster) createEndpoints() error {
	e := &EndpointData{
		clusterID:   p.clusterID,
		cluster:     p.cluster,
		clusterData: p.clusterData,
		log:         p.Log,
	}

	return e.create()
}

func (p *deployCluster) updateEndpoints() error {
	e := &EndpointData{
		clusterID:   p.clusterID,
		cluster:     p.cluster,
		clusterData: p.clusterData,
		log:         p.Log,
	}

	return e.update()
}

func (p *deployCluster) deleteEndpoints() error {
	e := &EndpointData{
		clusterID: p.clusterID,
		cluster:   p.cluster,
		log:       p.Log,
	}

	return e.remove()
}

func newAnsibleDeployer(cluster *Cluster, cData *Data) (base.Deployer, error) {
	return &contrailAnsibleDeployer{deployCluster{
		cluster:     cluster,
		clusterID:   cluster.config.ClusterID,
		action:      cluster.config.Action,
		clusterData: cData,
		Deploy: base.Deploy{
			Reporter: report.NewReporter(
				cluster.APIServer,
				fmt.Sprintf("%s/%s", defaultResourcePath, cluster.config.ClusterID),
				log.NewFileLogger("reporter", cluster.config.LogFile),
			),
			Log: log.NewFileLogger("contrail-ansible-deployer", cluster.config.LogFile),
		},
	}}, nil
}

func newMCProvisioner(cluster *Cluster, cData *Data, clusterID string, action string) (base.Deployer, error) {
	return &multiCloudProvisioner{contrailAnsibleDeployer{deployCluster{
		cluster:     cluster,
		clusterID:   clusterID,
		action:      action,
		clusterData: cData,
		Deploy: base.Deploy{
			Reporter: report.NewReporter(
				cluster.APIServer,
				fmt.Sprintf("%s/%s", defaultResourcePath, clusterID),
				log.NewFileLogger("reporter", cluster.config.LogFile),
			),
			Log: log.NewFileLogger("multi-cloud-provisioner", cluster.config.LogFile),
		},
	}}, ""}, nil
}

func newHelmDeployer(cluster *Cluster, cData *Data) (base.Deployer, error) {
	return &helmDeployer{deployCluster{
		cluster:     cluster,
		clusterID:   cluster.config.ClusterID,
		action:      cluster.config.Action,
		clusterData: cData,
		Deploy: base.Deploy{
			Reporter: report.NewReporter(
				cluster.APIServer,
				fmt.Sprintf("%s/%s", defaultResourcePath, cluster.config.ClusterID),
				log.NewFileLogger("reporter", cluster.config.LogFile),
			),
			Log: log.NewFileLogger("helm-deployer", cluster.config.LogFile),
		},
	}}, nil
}

func newDeployerByID(cluster *Cluster) (base.Deployer, error) {
	var cData *Data
	var err error
	if cluster.config.Action == "delete" {
		cData = &Data{}
	} else {
		cData, err = cluster.getClusterDetails(cluster.config.ClusterID)
	}
	if err != nil {
		return nil, err
	}
	deployerType := defaultDeployer
	if cData.clusterInfo != nil && cData.clusterInfo.ProvisionerType != "" {
		deployerType = cData.clusterInfo.ProvisionerType
	}

	// Check if cloudbackrefs are present
	if cluster.config.Action != deleteAction {
		if isMCProvisioner(cData) {
			deployerType = mCProvisioner
		}
	}

	switch deployerType {
	case "ansible":
		return newAnsibleDeployer(cluster, cData)
	case "helm":
		return newHelmDeployer(cluster, cData)
	case mCProvisioner:
		return newMCProvisioner(cluster, cData, cluster.config.ClusterID, cluster.config.Action)
	}
	return nil, errors.New("unsupported deployer type")
}

func hasCloudRefs(cData *Data) bool {

	if cData.cloudInfo != nil {
		return true
	}
	return false

}

func hasMCGWNodes(clusterInfo *models.ContrailCluster) bool {

	if clusterInfo.ContrailMulticloudGWNodes != nil {
		return true
	}
	return false
}

func isMCProvisioner(cData *Data) bool {

	state := cData.clusterInfo.ProvisioningState
	if hasCloudRefs(cData) && hasMCGWNodes(cData.clusterInfo) && (state == "NOSTATE" || state == "") {
		switch cData.clusterInfo.ProvisioningAction {
		case "ADD_CLOUD":
			return true
		case "UPDATE_CLOUD":
			return true
		case "DELETE_CLOUD":
			return true
		}
	}
	return false
}
