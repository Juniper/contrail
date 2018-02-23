package cluster

import (
	"errors"

	pkglog "github.com/Juniper/contrail/pkg/log"
	"github.com/sirupsen/logrus"
)

type provisioner interface {
	provison()
}

type ansibleProvisioner struct {
	builder     *Builder
	clusterInfo map[interface{}]interface{}
	log         *logrus.Entry
}

type helmProvisioner struct {
	builder     *Builder
	clusterInfo map[interface{}]interface{}
	log         *logrus.Entry
}

func (a *ansibleProvisioner) provison() error {
	return nil
}

func (h *helmProvisioner) provison() error {
	return nil
}

func newAnsibleProvisioner(builder *Builder, clusterInfo map[interface{}]interface{}) (*ansibleProvisioner, error) {
	return &ansibleProvisioner{
		builder:     builder,
		clusterInfo: clusterInfo,
		log:         pkglog.NewLogger("ansible-provisioner"),
	}, nil
}

func newHelmProvisioner(builder *Builder, clusterInfo map[interface{}]interface{}) (*helmProvisioner, error) {
	return &helmProvisioner{
		builder:     builder,
		clusterInfo: clusterInfo,
		log:         pkglog.NewLogger("helm-provisioner"),
	}, nil
}

// Creates new provisioner based on the type
func newProvisioner(builder *Builder) (provisioner, error) {
	schemaId := "contrail_cluster"
	//clusterInfo, err := builder.Agent.Show(schemaId, builder.clusterID)
	_, err := builder.Agent.Show(schemaId, builder.clusterID)
	if err != nil {
		return nil, err
	}

	/*
		provisionerType, err = clusterInfo["provisioner_type"]
		if err != nil {
			return nil, err
		}

		switch provisionerType {
		case "ansible":
			return newAnsibleProvisioner(builder, clusterInfo)
		case "helm":
			return newHelmProvisioner(builder, clusterInfo)
		}
		return nil, errors.New("unsupported provisioner type")
	*/
	return nil, errors.New("unsupported provisioner type")
}
