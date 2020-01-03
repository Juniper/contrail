package cloud

import (
	"github.com/Juniper/asf/pkg/fileutil"
	"github.com/Juniper/asf/pkg/logutil/report"
	"github.com/sirupsen/logrus"
)

const deployerCmd = "deployer"

// Providers is a list of a providers which should be used in Deployer operation.
type Providers []string

// DeployerOption is function which set optional params in Deployer
type DeployerOption func(*Deployer)

// DeployerWithTfStateFile is factory function for option which set TfStateFile.
func DeployerWithTfStateFile(file string) DeployerOption {
	return func(d *Deployer) {
		d.tfStateFile = file
	}
}

// DeployerWithStateFile is factory function for option which set StateFile.
func DeployerWithStateFile(file string) DeployerOption {
	return func(d *Deployer) {
		d.stateFile = file
	}
}

// DeployerWithDryRun is factory function for option which dry run deployer.
// Dry run will log deployer command instead of executing it.
func DeployerWithDryRun(log *logrus.Entry) DeployerOption {
	return func(d *Deployer) {
		d.commandExecutor = dryRunExecutor{log}
	}
}

type dryRunExecutor struct {
	log *logrus.Entry
}

func (d dryRunExecutor) ExecCmdAndWait(
	r *report.Reporter, deployerCmd string, args []string, dir string, envVars ...string,
) error {
	d.log.Infof("Dry runing %s command with [%s] args in %s directory with [%s] env", deployerCmd, args, dir, envVars)
	return nil
}

// NewDeployer creates new Deployer
func NewDeployer(
	config DeployerConfig, executor CommandExecutor, reporter *report.Reporter, opts ...DeployerOption,
) Deployer {
	deployer := Deployer{config: config, commandExecutor: executor, reporter: reporter}

	for _, o := range opts {
		o(&deployer)
	}
	return deployer
}

// DeployerConfig struct holds all required file and directory paths by Deployer.
type DeployerConfig struct {
	WorkDir      string
	TopologyFile string
	SecretFile   string
}

// Deployer is a multi cloud deployer. It wraps deployer cli execution.
type Deployer struct {
	config          DeployerConfig
	tfStateFile     string
	stateFile       string
	commandExecutor CommandExecutor
	reporter        *report.Reporter
}

// DestroyTopology method destroys topology.
func (d Deployer) DestroyTopology(providers Providers) error {
	if err := fileutil.WriteToFile(d.config.TopologyFile, []byte{}, defaultRWOnlyPerm); err != nil {
		return err
	}
	return d.UpdateTopology(providers)
}

// UpdateTopology method updates topology.
func (d Deployer) UpdateTopology(providers Providers) error {
	args := append([]string{"all", "topology", "--skip_validation"}, d.commonArgs()...)
	for _, p := range providers {
		args = append(args, "--limit", p)
	}
	return d.commandExecutor.ExecCmdAndWait(d.reporter, deployerCmd, args, d.config.WorkDir)
}

// Provision method provision multi cloud functionality.
func (d Deployer) Provision(isUpdate bool, vars []string) error {
	args := append([]string{"all", "provision"}, d.provisionArgs()...)
	if isUpdate {
		args = append(args, "--update")
	}
	return d.commandExecutor.ExecCmdAndWait(
		d.reporter, deployerCmd, args, d.config.WorkDir, vars...,
	)
}

// CleanupProvisioning method removes multi cloud functionality.
func (d Deployer) CleanupProvisioning(vars []string) error {
	args := append([]string{"all", "clean"}, d.provisionArgs()...)
	return d.commandExecutor.ExecCmdAndWait(d.reporter, deployerCmd, args, d.config.WorkDir, vars...)
}

func (d Deployer) commonArgs() []string {
	return []string{
		"--topology", d.config.TopologyFile,
		"--secret", d.config.SecretFile,
	}
}

func (d Deployer) provisionArgs() []string {
	return append(d.commonArgs(),
		"--tf_state", d.tfStateFile,
		"--state", d.stateFile,
	)
}
