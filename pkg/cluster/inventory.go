package cluster

import (
	"fmt"

	"github.com/flosch/pongo2"
)

type ansibleInventory struct {
	ansible *ansibleProvisioner
}

func (a *ansibleInventory) create() error {
	fmt.Printf(a.ansible.clusterData.clusterInfo.UUID)
	context := pongo2.Context{
		"cluster": a.ansible.clusterData.clusterInfo,
		"nodes":   a.ansible.clusterData.nodesInfo,
	}
	content, err := a.ansible.applyTemplate(a.ansible.getInstanceTemplate(), context)
	if err != nil {
		return err
	}
	err = a.ansible.appendToFile(a.ansible.getInstanceFile(), content)
	if err != nil {
		return err
	}
	return nil
}
