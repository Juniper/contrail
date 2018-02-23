package cluster

import (
	"fmt"

	"github.com/Juniper/contrail/pkg/agent"
	pkglog "github.com/Juniper/contrail/pkg/log"
	"github.com/sirupsen/logrus"
)

// Builder represents cluster builder
type Builder struct {
	clusterID string
	Agent     *agent.Agent
	log       *logrus.Entry
}

// Creates new cluster builder
func NewClusterBuilder(configPath string, clusterID string) (*Builder, error) {
	a, err := agent.NewAgentByFile(configPath)
	if err != nil {
		return nil, fmt.Errorf("agent creation failed: %s", err)
	}

	return &Builder{
		clusterID: clusterID,
		Agent:     a,
		log:       pkglog.NewLogger("cluster"),
	}, nil
}

// build starts building the cluster.
func (b *Builder) Build() error {
	b.log.Info("Start building cluster")
	err := b.Agent.APIServer.Login()
	if err != nil {
		return fmt.Errorf("login to API Server failed: %s", err)
	}

	//provisioner, err = newProvisioner(b)
	_, err = newProvisioner(b)
	if err != nil {
		return err
	}
	/*
		err = provisioner.provision()
		if err != nil {
			return err
		}
	*/

	return nil
}
