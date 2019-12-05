package cloud

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/Juniper/asf/pkg/fileutil"
	"github.com/Juniper/asf/pkg/fileutil/template"
	"github.com/Juniper/asf/pkg/logutil"
	"github.com/Juniper/asf/pkg/logutil/report"
	"github.com/flosch/pongo2"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"

	yaml "gopkg.in/yaml.v2"
)

type topology struct {
	cloudData *Data
	cloud     *Cloud
	log       *logrus.Entry
	reporter  *report.Reporter
	ctx       context.Context
}

func (t *topology) getTmplFilePath() string {
	return filepath.Join(t.cloud.getTemplateRoot(), t.getTopoTemplate())
}

func (t *topology) getTopoTemplate() string {

	if t.cloudData.isCloudPrivate() {
		return defaultOnPremTopoTemplate
	}
	return defaultPublicCloudTopoTemplate
}

func newTopology(c *Cloud, data *Data) *topology {
	return &topology{
		cloudData: data,
		cloud:     c,
		log:       logutil.NewFileLogger("topology", c.config.LogFile),
		reporter: report.NewReporter(
			c.APIServer,
			fmt.Sprintf("%s/%s", defaultCloudResourcePath, c.config.CloudID),
			logutil.NewFileLogger("reporter", c.config.LogFile),
		),
		ctx: c.ctx,
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

func (t *topology) marshalAndSave(topoFile string, p *pubCloud) error {

	marshaled, err := yaml.Marshal(p.Providers)
	if err != nil {
		return errors.Wrapf(err, "cannot marshal topology")
	}
	err = fileutil.WriteToFile(topoFile, marshaled, defaultRWOnlyPerm)
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

func (t *topology) isUpToDate(resource string) (bool, error) {
	status := map[string]interface{}{}
	if _, err := os.Stat(GetTopoFile(t.cloud.config.CloudID)); err == nil {
		ok, err := t.compareTopoFile()
		if err != nil {
			status[statusField] = statusUpdateFailed
			t.reporter.ReportStatus(t.ctx, status, resource)
			return true, err
		}
		if ok {
			t.log.Infof("%s topology file is already up-to-date", resource)
			return true, nil
		}
	}
	return false, nil
}
