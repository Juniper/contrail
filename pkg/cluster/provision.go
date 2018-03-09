package cluster

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	pkglog "github.com/Juniper/contrail/pkg/log"
	"github.com/flosch/pongo2"
	"github.com/sirupsen/logrus"
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
	reporter    *Reporter
}

func (p *provisionCommon) getSudoUserPassword() string {
	return p.clusterData.clusterInfo.SudoPassword
}
func (p *provisionCommon) getTemplateRoot() string {
	return defaultTemplateRoot
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
	dir := filepath.Join(p.getClusterHomeDir(), p.action)
	return dir
}

func (p *provisionCommon) createWorkingDir() error {
	err := os.MkdirAll(p.getWorkingDir(), os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}

func (p *provisionCommon) deleteWorkingDir() error {
	err := os.RemoveAll(p.getClusterHomeDir())
	if err != nil {
		return err
	}
	return nil
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

	var p provisioner
	p = &ansibleProvisioner{provisionCommon{
		cluster:     cluster,
		clusterID:   clusterID,
		action:      action,
		clusterData: cData,
		reporter:    r,
		log:         logger,
	}}
	return p, nil
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

	var p provisioner
	p = &helmProvisioner{provisionCommon{
		cluster:     cluster,
		clusterID:   clusterID,
		action:      action,
		clusterData: cData,
		reporter:    r,
		log:         logger,
	}}
	return p, nil
}

// Creates new provisioner based on the type
func newProvisioner(cluster *Cluster) (provisioner, error) {
	return newProvisionerByID(cluster, cluster.config.ClusterID, cluster.config.Action)
}

func newProvisionerByID(cluster *Cluster, clusterID string, action string) (provisioner, error) {
	cData, err := cluster.getClusterDetails(clusterID)
	if err != nil {
		return nil, err
	}
	provisionerType := cData.clusterInfo.ProvisionerType
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
