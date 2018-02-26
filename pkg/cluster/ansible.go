package cluster

import (
	"errors"

	pkglog "github.com/Juniper/contrail/pkg/log"
	"github.com/sirupsen/logrus"
)

type ansibleProvisioner struct {
	provisionCommon
}

func (a *ansibleProvisioner) cloneAnsibleDeployer() {
	return nil
}

func (a *ansibleProvisioner) createInventory() {
	return nil
}

func (a *ansibleProvisioner) playBook() {
	return nil
}

func (a *ansibleProvisioner) createCluster() {
	a.log.Info("Starting %s of contrail cluster: %s", action, a.clusterInfo.FQName)
	a.reporter.reportStatus("Intializing")

	err := a.createWorkingDir()
	if err != nil {
		a.reporter.reportStatus("Failed")
		return err
	}

	err = a.cloneAnsibleDeployer()
	if err != nil {
		a.reporter.reportStatus("Failed")
		return err
	}
	err = a.createInventory()
	if err != nil {
		a.reporter.reportStatus("Failed")
		return err
	}

	a.reporter.reportStatus("create_progress")
	err = a.playBook()
	if err != nil {
		a.reporter.reportStatus("Failed")
		return err
	}

	a.reporter.reportStatus("Created")
	return nil
}

func (a *ansibleProvisioner) updateCluster() {
	a.log.Info("Starting %s of contrail cluster: %s", action, a.clusterInfo.FQName)
	a.reporter.reportStatus("update_progress")
	err := a.playBook()
	if err != nil {
		a.reporter.reportStatus("Failed")
		return err
	}

	a.reporter.reportStatus("Updated")
	return nil
}

func (a *ansibleProvisioner) deleteCluster() {
	a.log.Info("Starting %s of contrail cluster: %s", action, a.clusterInfo.FQName)
	err := a.playBook()
	if err != nil {
		return err
	}

	err = a.deleteWorkingDir()
	if err != nil {
		a.reporter.reportStatus("Failed")
		return err
	}
	return nil
}

func (a *ansibleProvisioner) provison() error {
	switch a.action {
	case "create":
		return a.createCluster()
	case "update":
		return a.updateCluster()
	case "delete":
		return a.deleteCluster()
	}
	return nil, errors.New("unsupported action on cluster.")
}
