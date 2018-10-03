package cloud

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/flosch/pongo2"
	"github.com/sirupsen/logrus"

	report "github.com/Juniper/contrail/pkg/cluster"
	pkglog "github.com/Juniper/contrail/pkg/log"
)

const (
	defaultTopologyFile = "topology.yml"
	//defaultTopoTemplate = "topology.tmpl"
	defaultTopoTemplate = "topology.tmpl"
)

type topology struct {
	cloudData *Data
	cloud     *Cloud
	cloudID   string
	action    string
	log       *logrus.Entry
	reporter  *report.Reporter
}

func (t *topology) getTopoTemplate() string {
	return filepath.Join(t.cloud.getTemplateRoot(), defaultTopoTemplate)
}

func (t *topology) getTopoFile() string {
	return filepath.Join(t.cloud.getWorkingDir(), defaultTopologyFile)
}

func (c *Cloud) newTopology(d *Data) *topology {

	//create reporter for topology
	logger := pkglog.NewFileLogger("reporter", c.config.LogFile)
	pkglog.SetLogLevel(logger, c.config.LogLevel)

	r := &report.Reporter{
		API:      c.APIServer,
		Resource: fmt.Sprintf("%s/%s", defaultCloudResourcePath, c.config.CloudID),
		Log:      logger,
	}
	// create logger for topology
	logger = pkglog.NewFileLogger("topology", c.config.LogFile)
	pkglog.SetLogLevel(logger, c.config.LogLevel)

	return &topology{
		log:       logger,
		reporter:  r,
		cloudData: d,
		cloud:     c,
		cloudID:   c.config.CloudID,
		action:    c.config.Action,
	}
}

func (t *topology) createTopologyFile(topoFile string) error {

	context := pongo2.Context{
		"cloud": t.cloudData,
	}
	content, err := applyTemplate(t.getTopoTemplate(), context)
	if err != nil {
		return err
	}

	err = writeToFile(topoFile, content)
	if err != nil {
		return err
	}
	t.log.Infof("Created topology file for cloud with uuid: %s ", t.cloudData.cloudInfo.UUID)
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
			t.log.Errorf("Error while deleting tmpfile: %s", err)
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
	oldTopoFile, err := ioutil.ReadFile(t.getTopoFile())
	if err != nil {
		return false, err
	}

	return bytes.Equal(oldTopoFile, newTopoFile), nil
}

//TODO(madhukar) common func logic in cluster and cloud pkg
func (t *topology) isUpdated() (bool, error) {

	status := map[string]interface{}{}
	if _, err := os.Stat(t.getTopoFile()); err == nil {
		ok, err := t.compareTopoFile()
		if err != nil {
			status[statusField] = statusUpdateFailed
			t.reporter.ReportStatus(status)
			return false, err
		}
		if ok {
			t.log.Infof("Cloud topology file is already up-to-date")
			return true, nil
		}
	}
	return false, nil

}
