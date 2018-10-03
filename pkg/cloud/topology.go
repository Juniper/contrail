package cloud

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/flosch/pongo2"
	"github.com/sirupsen/logrus"

	"github.com/Juniper/contrail/pkg/common"
	"github.com/Juniper/contrail/pkg/fileutil"
	"github.com/Juniper/contrail/pkg/fileutil/template"
	pkglog "github.com/Juniper/contrail/pkg/log"
)

const (
	defaultTopologyFile = "topology.yml"
	//defaultTopoTemplate = "topology.tmpl"
	defaultCloudTopoTemplate  = "public_cloud_topology.tmpl"
	defaultOnPremTopoTemplate = "onprem_cloud_topology.tmpl"
)

type topology struct {
	cloudData *Data
	cloud     *Cloud
	log       *logrus.Entry
	reporter  *common.Reporter
}

func (t *topology) getTmplFilePath() string {
	return filepath.Join(t.cloud.getTemplateRoot(), t.getTopoTemplate())
}

func (t *topology) getTopoTemplate() string {

	if t.cloudData.isCloudPrivate() {
		return defaultOnPremTopoTemplate
	}
	return defaultCloudTopoTemplate
}

func (c *Cloud) newTopology(data *Data) *topology {

	// create reporter for topology
	logger := pkglog.NewFileLogger("reporter", c.config.LogFile)
	pkglog.SetLogLevel(logger, c.config.LogLevel)

	r := &common.Reporter{
		API:      c.APIServer,
		Resource: fmt.Sprintf("%s/%s", defaultCloudResourcePath, c.config.CloudID),
		Log:      logger,
	}

	// create logger for topology
	logger = pkglog.NewFileLogger("topology", c.config.LogFile)
	pkglog.SetLogLevel(logger, c.config.LogLevel)

	return &topology{
		cloudData: data,
		cloud:     c,
		log:       logger,
		reporter:  r,
	}
}

func (t *topology) createTopologyFile(topoFile string) error {

	context := pongo2.Context{
		"cloud": t.cloudData,
	}
	content, err := template.Apply(t.getTmplFilePath(), context)
	if err != nil {
		return err
	}

	err = fileutil.WriteToFile(topoFile, content, defaultRWOnlyPerm)
	if err != nil {
		return err
	}
	t.log.Infof("Created topology file for cloud with uuid: %s ", t.cloudData.info.UUID)
	return nil
}

//TODO(madhukar) common func logic in cluster and cloud pkg
func (t *topology) compareTopoFile() (bool, error) {

	tmpfile, err := ioutil.TempFile("", "topology")
	if err != nil {
		return false, err
	}
	tmpFileName := tmpfile.Name()
	defer func() {
		if err = os.Remove(tmpFileName); err != nil {
			t.log.Errorf("error while deleting tmpfile: %s", err)
		}
	}()

	t.log.Debugf("Creating temporary topology %s", tmpFileName)
	err = t.createTopologyFile(tmpFileName)
	if err != nil {
		return false, err
	}

	newTopoFile, err := ioutil.ReadFile(tmpFileName)
	if err != nil {
		return false, err
	}
	oldTopoFile, err := ioutil.ReadFile(GetTopoFile(t.cloud.config.CloudID))
	if err != nil {
		return false, err
	}

	return bytes.Equal(oldTopoFile, newTopoFile), nil
}

//TODO(madhukar) common func logic in cluster and cloud pkg
func (t *topology) isUpdated(resource string) (bool, error) {

	status := map[string]interface{}{}
	if _, err := os.Stat(GetTopoFile(t.cloud.config.CloudID)); err == nil {
		ok, err := t.compareTopoFile()
		if err != nil {
			status[statusField] = statusUpdateFailed
			t.reporter.ReportStatus(status, resource)
			return true, err
		}
		if ok {
			t.log.Infof("%s topology file is already up-to-date", resource)
			return true, nil
		}
	}
	return false, nil
}
