package cloud

import (
	"github.com/Juniper/asf/pkg/fileutil"
	"github.com/Juniper/asf/pkg/logutil/report"
	"github.com/sirupsen/logrus"
)

const deployerCmd = "deployer"

type Providers []string

type dryRunExecutor struct {
	log *logrus.Entry
}

func (d dryRunExecutor) ExecCmdAndWait(r *report.Reporter, deployerCmd string, args []string, dir string, envVars ...string) error {
	d.log.Infof("Dry runing %s command with [%s] args in %s directory with [%s] env", deployerCmd, args, dir, envVars)
	return nil
}

type Option func(*Deployer)

func DeployerWithTfStateFile(file string) Option {
	return func(d *Deployer) {
		d.tfStateFile = file
	}
}

func DeployerWithStateFile(file string) Option {
	return func(d *Deployer) {
		d.stateFile = file
	}
}

func DeployerWithDryRun(log *logrus.Entry) Option {
	return func(d *Deployer) {
		d.commandExecutor = dryRunExecutor{log}
	}
}

func NewDeployer(config DeployerConfig, executor CommandExecutor, reporter *report.Reporter, opts ...Option) Deployer {
	d := Deployer{config: config, commandExecutor: executor, reporter: reporter}

	for _, o := range opts {
		o(&d)
	}
	return d
}

type DeployerConfig struct {
	WorkDir      string
	TopologyFile string
	SecretFile   string
}

type Deployer struct {
	config          DeployerConfig
	tfStateFile     string
	stateFile       string
	commandExecutor CommandExecutor
	reporter        *report.Reporter
}

func (d Deployer) DestroyTopology(providers Providers) error {
	if err := fileutil.WriteToFile(d.config.TopologyFile, []byte{}, defaultRWOnlyPerm); err != nil {
		return err
	}
	return d.UpdateTopology(providers)
}

func (d Deployer) UpdateTopology(providers Providers) error {
	args := append([]string{"all", "topology", "--skip_validation"}, d.commonArgs()...)
	for _, p := range providers {
		args = append(args, "--limit", p)
	}
	return d.commandExecutor.ExecCmdAndWait(d.reporter, deployerCmd, args, d.config.WorkDir)
}

func (d Deployer) Provision(isUpdate bool, vars []string) error {
	args := append([]string{"all", "provision"}, d.provisionArgs()...)
	if isUpdate {
		args = append(args, "--update")
	}
	return d.commandExecutor.ExecCmdAndWait(
		d.reporter, deployerCmd, args, d.config.WorkDir, vars...,
	)
}

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
