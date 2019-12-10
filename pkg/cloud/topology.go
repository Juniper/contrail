package cloud

import (
	"context"
	"fmt"
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

func (t *topology) getOnPremTemplatePath() string {
	return filepath.Join(t.cloud.getTemplateRoot(), defaultOnPremTopoTemplate)
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

func (t *topology) createOnPremTopologyFile(topoFile string) error {
	context := pongo2.Context{
		"cloud": t.cloudData,
	}
	content, err := template.Apply(t.getOnPremTemplatePath(), context)
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

func (t *topology) marshalAndSave(topoFile string, p *publicCloud) error {

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

func (t *topology) marshalOnPremAndSave(topoFile string, p *onPremCloud) error {

	marshaled, err := yaml.Marshal([]*onPremCloud{p})
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
