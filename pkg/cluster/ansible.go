package cluster

import ()

type ansibleProvisioner struct {
	provisionCommon
}

func (a *ansibleProvisioner) cloneAnsibleDeployer() error {
	return nil
}

func (a *ansibleProvisioner) createInventory() error {
	return nil
}

func (a *ansibleProvisioner) playBook() error {
	return nil
}

func (a *ansibleProvisioner) createCluster() error {
	a.log.Info("Starting %s of contrail cluster: %s", a.action, a.clusterInfo["fq_name"])
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

func (a *ansibleProvisioner) updateCluster() error {
	a.log.Info("Starting %s of contrail cluster: %s", a.action, a.clusterInfo["fq_name"])
	a.reporter.reportStatus("update_progress")
	err := a.playBook()
	if err != nil {
		a.reporter.reportStatus("Failed")
		return err
	}

	a.reporter.reportStatus("Updated")
	return nil
}

func (a *ansibleProvisioner) deleteCluster() error {
	a.log.Info("Starting %s of contrail cluster: %s", a.action, a.clusterInfo["fq_name"])
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

func (a *ansibleProvisioner) provision() error {
	switch a.action {
	case "create":
		return a.createCluster()
	case "update":
		return a.updateCluster()
	case "delete":
		return a.deleteCluster()
	}
	return nil
}
