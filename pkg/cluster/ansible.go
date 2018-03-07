package cluster

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/flosch/pongo2"
)

type ansibleProvisioner struct {
	provisionCommon
}

func (a *ansibleProvisioner) getInstanceTemplate() string {
	return filepath.Join(a.getTemplateRoot(), defaultInstanceTemplate)
}

func (a *ansibleProvisioner) getInstanceFile() string {
	return filepath.Join(a.getWorkingDir(), defaultInstanceFile)
}

func (a *ansibleProvisioner) getAnsibleRepoDir() string {
	return filepath.Join(a.getWorkingDir(), defaultAnsibleRepo)
}

func (a *ansibleProvisioner) cloneAnsibleDeployer() error {
	a.log.Infof("Clean working dir to clone %s", defaultAnsibleRepo)
	repoDir := a.getAnsibleRepoDir()
	err := os.RemoveAll(repoDir)
	if err != nil {
		return err
	}
	a.log.Infof("Cloning repo:%s into %s", defaultAnsibleRepoURL, repoDir)
	args := []string{"clone", defaultAnsibleRepoURL, repoDir}
	cmd := exec.Command("git", args...)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return err
	}
	if err := cmd.Start(); err != nil {
		return err
	}
	// Report progress log periodically to stdout/db
	go a.reporter.reportLog(stdout)
	if err := cmd.Wait(); err != nil {
		return err
	}
	a.log.Info("Cloning completed")

	return nil
}

func (a *ansibleProvisioner) createInventory() error {
	a.log.Info("Creating instance.yml input file for ansible deployer")
	context := pongo2.Context{
		"cluster": a.clusterData.clusterInfo,
		"nodes":   a.clusterData.nodesInfo,
	}
	content, err := a.applyTemplate(a.getInstanceTemplate(), context)
	if err != nil {
		return err
	}
	err = a.appendToFile(a.getInstanceFile(), content)
	if err != nil {
		return err
	}
	a.log.Info("Created instance.yml input file for ansible deployer")
	return nil
}

func (a *ansibleProvisioner) playBook() error {
	repoDir := a.getAnsibleRepoDir()
	cmdline := "ansible-playbook"
	args := []string{"-i", "inventory/", "-e",
		"config_file=" + a.getInstanceFile(),
		defaultInstanceProvPlay}

	a.log.Infof("Playing instance provisioning playbook: %s %s",
		cmdline, strings.Join(args, " "))
	cmd := exec.Command(cmdline, args...)
	cmd.Dir = repoDir
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return err
	}
	if err = cmd.Start(); err != nil {
		return err
	}

	// Report progress log periodically to stdout/db
	go a.reporter.reportLog(stdout)

	if err = cmd.Wait(); err != nil {
		return err
	}
	a.log.Info("Instance provisioning play completed")

	args = args[:len(args)-1]
	args = append(args, defaultInstanceConfPlay)
	a.log.Infof("Playing instance configuration playbook: %s %s",
		cmdline, strings.Join(args, " "))
	cmd = exec.Command(cmdline, args...)
	cmd.Dir = repoDir
	stdout, err = cmd.StdoutPipe()
	if err != nil {
		return err
	}
	if err = cmd.Start(); err != nil {
		return err
	}

	// Report progress log periodically to stdout/db
	go a.reporter.reportLog(stdout)

	if err = cmd.Wait(); err != nil {
		return err
	}
	a.log.Info("Instance configuration play completed")

	args = args[:len(args)-1]
	args = append(args, defaultClusterProvPlay)
	a.log.Infof("Playing contrail cluster provisioning playbook: %s %s",
		cmdline, strings.Join(args, " "))
	cmd = exec.Command(cmdline, args...)
	cmd.Dir = repoDir
	stdout, err = cmd.StdoutPipe()
	if err != nil {
		return err
	}
	if err = cmd.Start(); err != nil {
		return err
	}

	// Report progress log periodically to stdout/db
	go a.reporter.reportLog(stdout)

	if err = cmd.Wait(); err != nil {
		return err
	}
	a.log.Info("Instance configuration play completed")
	return nil
}

func (a *ansibleProvisioner) createCluster() error {
	a.log.Infof("Starting %s of contrail cluster: %s", a.action, a.clusterData.clusterInfo.FQName)
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
	a.log.Infof("Starting %s of contrail cluster: %s", a.action, a.clusterData.clusterInfo.FQName)
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
	a.log.Infof("Starting %s of contrail cluster: %s", a.action, a.clusterData.clusterInfo.FQName)
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
