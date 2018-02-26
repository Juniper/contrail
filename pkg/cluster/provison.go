package cluster

import (
	"errors"
	"os"
	"path/filepath"

	pkglog "github.com/Juniper/contrail/pkg/log"
	"github.com/sirupsen/logrus"
)

const (
	resourcePath = "contrail-cluster"
	workHome     = "/opt"
)

type provisioner interface {
	provison()
}

type provisionCommon struct {
	manager     *Manager
	clusterID   string
	action      string
	clusterInfo map[interface{}]interface{}
	log         *logrus.Entry
	reporter    *Reporter
}

func (p *provisionCommon) getClusterHomeDir() (string, error) {
	dir, err := filepath.Join(workHome, p.clusterID)
	if err != nil {
		return nil, err
	}
	return dir, nil
}

func (p *provisionCommon) getWorkingDir() (string, error) {
	dir, err := filepath.Join(p.getClusterHomeDir(), p.action)
	if err != nil {
		return nil, err
	}
	return dir, nil
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

func newAnsibleProvisioner(manager *Manager, clusterInfo map[interface{}]interface{}, clusterID string, action string) (*provisioner, error) {
	r := &Reporter{
		api:      manager.cluster.APIServer,
		resource: fmt.Sprintf("%s?%s", resourcePath, clusterID),
		log:      pkglog.NewLogger("reporter"),
	}

	var p provisioner
	p = ansibleProvisioner{provisionCommon{
		manager:     manager,
		clusterID:   clusterID,
		action:      action,
		clusterInfo: clusterInfo,
		reporter:    r,
		log:         pkglog.NewLogger("ansible-provisioner"),
	}}
	return &p, nil
}

func newHelmProvisioner(manager *Manager, clusterInfo map[interface{}]interface{}, clusterID string, action string) (*provisioner, error) {
	r := &Reporter{
		api:      manager.cluster.APIServer,
		resource: fmt.Sprintf("%s?%s", resourcePath, clusterID),
		log:      pkglog.NewLogger("reporter"),
	}

	var p provisioner
	p = helmProvisioner{provisionCommon{
		manager:     manager,
		clusterID:   clusterID,
		action:      action,
		clusterInfo: clusterInfo,
		reporter:    r,
		log:         pkglog.NewLogger("helm-provisioner"),
	}}
	return &p, nil
}

// Creates new provisioner based on the type
func newProvisioner(manager *Manager) (provisioner, error) {
	return newProvisionerByID(manager, manager.cluster.config.clusterID, manager.cluster.config.Action)
}

func newProvisionerByID(manager *Manager, clusterID string, action string) (provisioner, error) {
	var clusterInfo map[string]interface{}
	manager.log.Debug("Polling data")
	_, err := manager.cluster.APIServer.Read(fmt.Sprintf("%s?%s", resourcePath, clusterID), &clusterInfo)
	if err != nil {
		return nil, err
	}

	provisionerType, err = clusterInfo["provisioner_type"]
	if err != nil {
		return nil, err
	}

	switch provisionerType {
	case "ansible":
		return newAnsibleProvisioner(manager, clusterInfo, clusterID, action)
	case "helm":
		return newHelmProvisioner(manager, clusterInfo, clusterID, action)
	}
	return nil, errors.New("unsupported provisioner type")
}
