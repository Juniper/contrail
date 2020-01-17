package cloud

import (
	"context"
	"fmt"
	"path/filepath"

	"github.com/flosch/pongo2"
	"github.com/sirupsen/logrus"

	"github.com/Juniper/asf/pkg/fileutil"
	"github.com/Juniper/asf/pkg/fileutil/template"
	"github.com/Juniper/asf/pkg/logutil"
	"github.com/Juniper/asf/pkg/logutil/report"
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
