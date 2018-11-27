package cluster

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/sirupsen/logrus"

	"github.com/Juniper/contrail/pkg/common"
	pkglog "github.com/Juniper/contrail/pkg/log"
)

type provisioner interface {
	provision() error
}

type provisionCommon struct {
	cluster     *Cluster
	clusterID   string
	action      string
	clusterData *Data
	log         *logrus.Entry
	reporter    *common.Reporter
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

func (p *provisionCommon) createEndpoints() error {
	e := &EndpointData{
		clusterID:   p.clusterID,
		cluster:     p.cluster,
		clusterData: p.clusterData,
		log:         p.log,
	}

	return e.create()
}

func (p *provisionCommon) updateEndpoints() error {
	e := &EndpointData{
		clusterID:   p.clusterID,
		cluster:     p.cluster,
		clusterData: p.clusterData,
		log:         p.log,
	}

	return e.update()
}

func (p *provisionCommon) deleteEndpoints() error {
	e := &EndpointData{
		clusterID: p.clusterID,
		cluster:   p.cluster,
		log:       p.log,
	}

	return e.remove()
}

func newAnsibleProvisioner(cluster *Cluster, cData *Data, clusterID string, action string) (provisioner, error) {
	// create logger for reporter
	logger := pkglog.NewLogger("reporter")
	pkglog.SetLogLevel(logger, cluster.config.LogLevel)

	r := common.NewReporter(cluster.APIServer,
		fmt.Sprintf("%s/%s", defaultResourcePath, clusterID), logger)

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

	r := common.NewReporter(cluster.APIServer,
		fmt.Sprintf("%s/%s", defaultResourcePath, clusterID), logger)

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
