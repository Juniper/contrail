package cloud

import (
	"path/filepath"

	"github.com/flosch/pongo2"
	"github.com/sirupsen/logrus"

	pkglog "github.com/Juniper/contrail/pkg/log"
)

const (
	defaultTopologyFile = "topology.yml"
	//defaultTopoTemplate = "topology.tmpl"
	defaultTopoTemplate = "topology_test.tmpl"
)

type topology struct {
	cloudData *Data
	cloud     *Cloud
	cloudID   string
	action    string
	log       *logrus.Entry
}

func (t *topology) getTopoTemplate() string {
	return filepath.Join(t.cloud.getTemplateRoot(), defaultTopoTemplate)
}

func (t *topology) getTopoFile() string {
	return filepath.Join(t.cloud.getWorkingDir(), defaultTopologyFile)
}

func (c *Cloud) newtopology(data *Data) *topology {

	// create logger for topology
	logger := pkglog.NewFileLogger("topology", c.config.LogFile)
	pkglog.SetLogLevel(logger, c.config.LogLevel)

	return &topology{
		log:       logger,
		cloudData: data,
		cloud:     c,
		cloudID:   c.config.CloudID,
		action:    c.config.Action,
	}
}

func (t *topology) createTopologyFile() error {
	topoFile := t.getTopoFile()

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

func (c *Cloud) getTopology() (*topology, error) {
	//TODO(madhukar) Support daemon manager with etcd

	cloudData, err := c.getCloudData()
	if err != nil {
		return nil, err
	}

	topo := c.newtopology(cloudData)

	err = topo.createTopologyFile()
	if err != nil {
		return nil, err
	}
	return topo, nil
}
