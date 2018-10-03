package cluster

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"

	"github.com/flosch/pongo2"
	"github.com/sirupsen/logrus"

	"github.com/Juniper/contrail/pkg/common"
	pkglog "github.com/Juniper/contrail/pkg/log"
	"github.com/Juniper/contrail/pkg/models"
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

func (p *provisionCommon) applyTemplate(templateSrc string, context map[string]interface{}) ([]byte, error) {
	template, err := pongo2.FromFile(templateSrc)
	if err != nil {
		return nil, err
	}
	output, err := template.ExecuteBytes(context)
	if err != nil {
		return nil, err
	}
	// strip empty lines in output content
	regex, _ := regexp.Compile("\n[ \r\n\t]*\n") // nolint: errcheck
	outputString := regex.ReplaceAllString(string(output), "\n")
	return []byte(outputString), nil
}

func (p *provisionCommon) writeToFile(path string, content []byte) error {
	err := os.MkdirAll(filepath.Dir(path), os.ModePerm)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(path, content, os.FileMode(0600))
	return err
}

func (p *provisionCommon) appendToFile(path string, content []byte) error {
	err := os.MkdirAll(filepath.Dir(path), os.ModePerm)
	if err != nil {
		return err
	}

	f, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0600)
	if err != nil {
		return err
	}
	_, err = f.Write(content)
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
	go p.reporter.ReportLog(stdout)
	go p.reporter.ReportLog(stderr)
	return cmdline.Wait()
}

func newAnsibleProvisioner(cluster *Cluster, cData *Data, clusterID string, action string) (provisioner, error) {
	// create logger for reporter
	logger := pkglog.NewFileLogger("reporter", cluster.config.LogFile)
	pkglog.SetLogLevel(logger, cluster.config.LogLevel)

	r := &common.Reporter{
		API:      cluster.APIServer,
		Resource: fmt.Sprintf("%s/%s", defaultResourcePath, clusterID),
		Log:      logger,
	}

	// create logger for ansible provisioner
	logger = pkglog.NewFileLogger("ansible-provisioner", cluster.config.LogFile)
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

func newMCProvisioner(cluster *Cluster, cData *Data, clusterID string, action string) (provisioner, error) {
	// create logger for reporter
	logger := pkglog.NewFileLogger("reporter", cluster.config.LogFile)
	pkglog.SetLogLevel(logger, cluster.config.LogLevel)

	r := &common.Reporter{
		API:      cluster.APIServer,
		Resource: fmt.Sprintf("%s/%s", defaultResourcePath, clusterID),
		Log:      logger,
	}

	// create logger for multi-cloud provisioner
	logger = pkglog.NewFileLogger("multi-cloud-provisioner", cluster.config.LogFile)
	pkglog.SetLogLevel(logger, cluster.config.LogLevel)

	action, err := getMCAction(cData.clusterInfo)
	if err != nil {
		return nil, err
	}

	return &multiCloudProvisioner{ansibleProvisioner{provisionCommon{
		cluster:     cluster,
		clusterID:   clusterID,
		action:      action,
		clusterData: cData,
		reporter:    r,
		log:         logger,
	}}, ""}, nil
}

func newHelmProvisioner(cluster *Cluster, cData *Data, clusterID string, action string) (provisioner, error) {
	// create logger for reporter
	logger := pkglog.NewFileLogger("reporter", cluster.config.LogFile)
	pkglog.SetLogLevel(logger, cluster.config.LogLevel)

	r := &common.Reporter{
		API:      cluster.APIServer,
		Resource: fmt.Sprintf("%s/%s", defaultResourcePath, clusterID),
		Log:      logger,
	}

	// create logger for Helm provisioner
	logger = pkglog.NewFileLogger("helm-provisioner", cluster.config.LogFile)
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

	// Check if cloudbackrefs are present
	if isMCProvisioner(cData) {
		provisionerType = mCProvisioner
	}

	switch provisionerType {
	case "ansible":
		return newAnsibleProvisioner(cluster, cData, clusterID, action)
	case "helm":
		return newHelmProvisioner(cluster, cData, clusterID, action)
	case mCProvisioner:
		return newMCProvisioner(cluster, cData, clusterID, action)
	}
	return nil, errors.New("unsupported provisioner type")
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

func getMCAction(clusterInfo *models.ContrailCluster) (string, error) {

	switch clusterInfo.ProvisioningAction {
	case "ADD_CLOUD":
		return createAction, nil
	case "UPDATE_CLOUD":
		return updateAction, nil
	case "DELETE_CLOUD":
		return deleteAction, nil
	case "ADD_COMPUTE":
		return updateAction, nil
	}

	return "", fmt.Errorf("invalid provisioning action for cluster with cloud ref:  %s", clusterInfo.ProvisioningAction)

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
		case "ADD_COMPUTE":
			return true
			// implement delete compute
		}
	}
	return false
}
