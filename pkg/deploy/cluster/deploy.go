package cluster

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/pkg/errors"

	"github.com/Juniper/contrail/pkg/deploy/base"
	"github.com/Juniper/contrail/pkg/deploy/rhospd/overcloud"
	"github.com/Juniper/contrail/pkg/logutil"
	"github.com/Juniper/contrail/pkg/logutil/report"
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
	clusterData *base.Data
}

func (p *deployCluster) isCreated() bool {
	state := p.clusterData.ClusterInfo.ProvisioningState
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

func (p *deployCluster) getWorkRoot() string {
	workRoot := p.cluster.config.WorkRoot
	if workRoot == "" {
		workRoot = defaultWorkRoot
	}
	return workRoot
}

func (p *deployCluster) getClusterHomeDir() string {
	dir := filepath.Join(p.getWorkRoot(), p.clusterID)
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
	e := &base.EndpointData{
		ClusterID:   p.clusterID,
		ResManager:  base.NewResourceManager(p.cluster.APIServer, p.cluster.config.LogFile),
		ClusterData: p.clusterData,
		Log:         p.Log,
	}

	return e.Create()
}

func (p *deployCluster) updateEndpoints() error {
	e := &base.EndpointData{
		ClusterID:   p.clusterID,
		ResManager:  base.NewResourceManager(p.cluster.APIServer, p.cluster.config.LogFile),
		ClusterData: p.clusterData,
		Log:         p.Log,
	}

	return e.Update()
}

func (p *deployCluster) deleteEndpoints() error {
	e := &base.EndpointData{
		ClusterID:  p.clusterID,
		ResManager: base.NewResourceManager(p.cluster.APIServer, p.cluster.config.LogFile),
		Log:        p.Log,
	}

	return e.Remove()
}

func newAnsibleDeployer(cluster *Cluster, cData *base.Data) (base.Deployer, error) {
	return &contrailAnsibleDeployer{deployCluster{
		cluster:     cluster,
		clusterID:   cluster.config.ClusterID,
		action:      cluster.config.Action,
		clusterData: cData,
		Deploy: base.Deploy{
			Reporter: report.NewReporter(
				cluster.APIServer,
				fmt.Sprintf("%s/%s", defaultResourcePath, cluster.config.ClusterID),
				logutil.NewFileLogger("reporter", cluster.config.LogFile),
			),
			Log: logutil.NewFileLogger("contrail-ansible-deployer", cluster.config.LogFile),
		},
	}}, nil
}

func newMCProvisioner(cluster *Cluster, cData *base.Data, clusterID string, action string) (base.Deployer, error) {
	return &multiCloudProvisioner{contrailAnsibleDeployer{deployCluster{
		cluster:     cluster,
		clusterID:   clusterID,
		action:      action,
		clusterData: cData,
		Deploy: base.Deploy{
			Reporter: report.NewReporter(
				cluster.APIServer,
				fmt.Sprintf("%s/%s", defaultResourcePath, clusterID),
				logutil.NewFileLogger("reporter", cluster.config.LogFile),
			),
			Log: logutil.NewFileLogger("multi-cloud-provisioner", cluster.config.LogFile),
		},
	}}, ""}, nil
}

func newHelmDeployer(cluster *Cluster, cData *base.Data) (base.Deployer, error) {
	return &helmDeployer{deployCluster{
		cluster:     cluster,
		clusterID:   cluster.config.ClusterID,
		action:      cluster.config.Action,
		clusterData: cData,
		Deploy: base.Deploy{
			Reporter: report.NewReporter(
				cluster.APIServer,
				fmt.Sprintf("%s/%s", defaultResourcePath, cluster.config.ClusterID),
				logutil.NewFileLogger("reporter", cluster.config.LogFile),
			),
			Log: logutil.NewFileLogger("helm-deployer", cluster.config.LogFile),
		},
	}}, nil
}

// nolint: gocyclo
func newDeployerByID(cluster *Cluster) (base.Deployer, error) {
	var cData *base.Data
	var err error
	if cluster.config.Action == "delete" {
		cData = &base.Data{Reader: cluster.APIServer}
	} else {
		r := base.NewResourceManager(cluster.APIServer, cluster.config.LogFile)
		cData, err = r.GetClusterDetails(cluster.config.ClusterID)
	}
	if err != nil {
		return nil, err
	}
	deployerType := defaultDeployer
	if cData.ClusterInfo != nil && cData.ClusterInfo.ProvisionerType != "" {
		deployerType = cData.ClusterInfo.ProvisionerType
	}

	// Check if cloudbackrefs are present
	if cluster.config.Action != deleteAction {
		if isMCProvisioner(cData) {
			deployerType = mCProvisioner
		}
	}

	switch deployerType {
	case "rhospd":
		c := &overcloud.Config{
			APIServer:    cluster.APIServer,
			ResourceID:   cluster.config.ClusterID,
			Action:       cluster.config.Action,
			TemplateRoot: cluster.config.TemplateRoot,
			LogLevel:     cluster.config.LogLevel,
			LogFile:      cluster.config.LogFile,
		}
		overcloud, err := overcloud.NewOverCloud(c)
		if err != nil {
			return nil, err
		}
		return overcloud.GetDeployer()
	case "ansible":
		return newAnsibleDeployer(cluster, cData)
	case "helm":
		return newHelmDeployer(cluster, cData)
	case mCProvisioner:
		return newMCProvisioner(cluster, cData, cluster.config.ClusterID, cluster.config.Action)
	}
	return nil, errors.New("unsupported deployer type")
}

func hasCloudRefs(cData *base.Data) bool {

	if cData.CloudInfo != nil {
		return true
	}
	return false

}

func hasMCGWNodes(ClusterInfo *models.ContrailCluster) bool {

	if ClusterInfo.ContrailMulticloudGWNodes != nil {
		return true
	}
	return false
}

func isMCProvisioner(cData *base.Data) bool {

	if hasCloudRefs(cData) && hasMCGWNodes(cData.ClusterInfo) {
		switch cData.ClusterInfo.ProvisioningAction {
		case addCloud:
			return true
		case updateCloud:
			return true
		case deleteCloud:
			return true
		}
	}
	return false
}
