package cluster

import (
	"errors"
	"fmt"
	"os/exec"
	"path/filepath"

	pkglog "github.com/Juniper/contrail/pkg/log"
	"github.com/sirupsen/logrus"
)

const (
	resourcePath = "contrail-cluster"
	workHome     = "/opt"
)

type provisioner interface {
	provision() error
}

type provisionCommon struct {
	cluster     *Cluster
	clusterID   string
	action      string
	clusterInfo map[string]interface{}
	log         *logrus.Entry
	reporter    *Reporter
}

func (p *provisionCommon) getClusterHomeDir() string {
	dir := filepath.Join(workHome, p.clusterID)
	return dir
}

func (p *provisionCommon) getWorkingDir() string {
	dir := filepath.Join(p.getClusterHomeDir(), p.action)
	return dir
}

func (p *provisionCommon) createWorkingDir() error {
	cmd := exec.Command("mkdir", "-p", p.getWorkingDir())
	err := cmd.Run()
	if err != nil {
		return err
	}
	return nil
}

func (p *provisionCommon) deleteWorkingDir() error {
	cmd := exec.Command("rm", "-r", p.getClusterHomeDir())
	err := cmd.Run()
	if err != nil {
		return err
	}
	return nil
}

func newAnsibleProvisioner(cluster *Cluster, clusterInfo map[string]interface{}, clusterID string, action string) (provisioner, error) {
	r := &Reporter{
		api:      cluster.APIServer,
		resource: fmt.Sprintf("%s?%s", resourcePath, clusterID),
		log:      pkglog.NewLogger("reporter"),
	}

	var p provisioner
	p = &ansibleProvisioner{provisionCommon{
		cluster:     cluster,
		clusterID:   clusterID,
		action:      action,
		clusterInfo: clusterInfo,
		reporter:    r,
		log:         pkglog.NewLogger("ansible-provisioner"),
	}}
	return p, nil
}

func newHelmProvisioner(cluster *Cluster, clusterInfo map[string]interface{}, clusterID string, action string) (provisioner, error) {
	r := &Reporter{
		api:      cluster.APIServer,
		resource: fmt.Sprintf("%s?%s", resourcePath, clusterID),
		log:      pkglog.NewLogger("reporter"),
	}

	var p provisioner
	p = &helmProvisioner{provisionCommon{
		cluster:     cluster,
		clusterID:   clusterID,
		action:      action,
		clusterInfo: clusterInfo,
		reporter:    r,
		log:         pkglog.NewLogger("helm-provisioner"),
	}}
	return p, nil
}

// Creates new provisioner based on the type
func newProvisioner(cluster *Cluster) (provisioner, error) {
	return newProvisionerByID(cluster, cluster.config.ClusterID, cluster.config.Action)
}

func newProvisionerByID(cluster *Cluster, clusterID string, action string) (provisioner, error) {
	var clusterInfo map[string]interface{}
	_, err := cluster.APIServer.Read(fmt.Sprintf("%s?%s", resourcePath, clusterID), &clusterInfo)
	if err != nil {
		return nil, err
	}

	provisionerType, present := clusterInfo["provisioner_type"]
	if !present {
		return nil, errors.New("Provisioner type not specified in the cluster")
	}

	switch provisionerType {
	case "ansible":
		return newAnsibleProvisioner(cluster, clusterInfo, clusterID, action)
	case "helm":
		return newHelmProvisioner(cluster, clusterInfo, clusterID, action)
	}
	return nil, errors.New("unsupported provisioner type")
}
