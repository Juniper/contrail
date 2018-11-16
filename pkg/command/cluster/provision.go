package cluster

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/Juniper/contrail/pkg/command/provision"
	pkglog "github.com/Juniper/contrail/pkg/log"
	"github.com/Juniper/contrail/pkg/log/report"
)

const (
	defaultTemplateRoot = "./pkg/cluster/configs"
)

type provisionCluster struct {
	provision.Provision
	cluster     *Cluster
	clusterID   string
	action      string
	clusterData *Data
}

func (p *provisionCluster) isCreated() bool {
	state := p.clusterData.clusterInfo.ProvisioningState
	if p.action == "create" && (state == "NOSTATE" || state == "") {
		return false
	}
	p.Log.Infof("Cluster %s already provisioned, STATE: %s", p.clusterID, state)
	return true
}

func (p *provisionCluster) getTemplateRoot() string {
	templateRoot := p.cluster.config.TemplateRoot
	if templateRoot == "" {
		templateRoot = defaultTemplateRoot
	}
	return templateRoot
}

func (p *provisionCluster) getClusterHomeDir() string {
	dir := filepath.Join(defaultWorkRoot, p.clusterID)
	return dir
}

func (p *provisionCluster) getWorkingDir() string {
	dir := filepath.Join(p.getClusterHomeDir())
	return dir
}

func (p *provisionCluster) createWorkingDir() error {
	return os.MkdirAll(p.getWorkingDir(), os.ModePerm)
}

func (p *provisionCluster) deleteWorkingDir() error {
	return os.RemoveAll(p.getClusterHomeDir())
}

func (p *provisionCluster) createEndpoints() error {
	e := &EndpointData{
		clusterID:   p.clusterID,
		cluster:     p.cluster,
		clusterData: p.clusterData,
		log:         p.Log,
	}

	return e.create()
}

func (p *provisionCluster) updateEndpoints() error {
	e := &EndpointData{
		clusterID:   p.clusterID,
		cluster:     p.cluster,
		clusterData: p.clusterData,
		log:         p.Log,
	}

	return e.update()
}

func (p *provisionCluster) deleteEndpoints() error {
	e := &EndpointData{
		clusterID: p.clusterID,
		cluster:   p.cluster,
		log:       p.Log,
	}

	return e.remove()
}

func newAnsibleProvisioner(cluster *Cluster, cData *Data) (provision.Provisioner, error) {
	clusterID := cluster.config.ResourceID
	// create logger for reporter
	logger := pkglog.NewFileLogger("reporter", cluster.config.LogFile)
	pkglog.SetLogLevel(logger, cluster.config.LogLevel)

	r := report.NewReporter(cluster.APIServer,
		fmt.Sprintf("%s/%s", defaultResourcePath, clusterID), logger)

	// create logger for ansible provisioner
	logger = pkglog.NewFileLogger("ansible-provisioner", cluster.config.LogFile)
	pkglog.SetLogLevel(logger, cluster.config.LogLevel)

	return &ansibleProvisioner{provisionCluster{
		cluster:     cluster,
		clusterID:   clusterID,
		action:      cluster.config.Action,
		clusterData: cData,
		Provision: provision.Provision{
			Reporter: r,
			Log:      logger,
		},
	}}, nil
}

func newHelmProvisioner(cluster *Cluster, cData *Data) (provision.Provisioner, error) {
	clusterID := cluster.config.ResourceID
	// create logger for reporter
	logger := pkglog.NewFileLogger("reporter", cluster.config.LogFile)
	pkglog.SetLogLevel(logger, cluster.config.LogLevel)

	r := report.NewReporter(cluster.APIServer,
		fmt.Sprintf("%s/%s", defaultResourcePath, clusterID), logger)

	// create logger for Helm provisioner
	logger = pkglog.NewFileLogger("helm-provisioner", cluster.config.LogFile)
	pkglog.SetLogLevel(logger, cluster.config.LogLevel)

	return &helmProvisioner{provisionCluster{
		cluster:     cluster,
		clusterID:   clusterID,
		action:      cluster.config.Action,
		clusterData: cData,
		Provision: provision.Provision{
			Reporter: r,
			Log:      logger,
		},
	}}, nil
}

func newProvisionerByID(cluster *Cluster) (provision.Provisioner, error) {
	var cData *Data
	var err error
	if cluster.config.Action == "delete" {
		cData = &Data{}
	} else {
		cData, err = cluster.getClusterDetails(cluster.config.ResourceID)
	}
	if err != nil {
		return nil, err
	}
	provisionerType := defaultProvisioner
	if cData.clusterInfo != nil && cData.clusterInfo.ProvisionerType != "" {
		provisionerType = cData.clusterInfo.ProvisionerType
	}

	switch provisionerType {
	case "ansible":
		return newAnsibleProvisioner(cluster, cData)
	case "helm":
		return newHelmProvisioner(cluster, cData)
	}
	return nil, errors.New("unsupported provisioner type")
}
