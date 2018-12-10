package cloud

import (
	"fmt"
	"strings"

	"github.com/sirupsen/logrus"

	pkglog "github.com/Juniper/contrail/pkg/log"
	"github.com/Juniper/contrail/pkg/log/report"
	"github.com/Juniper/contrail/pkg/osutil"
)

const (
	tfCommand = "terraform"
)

type terraform struct {
	cloud      *Cloud
	log        *logrus.Entry
	reporter   *report.Reporter
	action     string
	topoFile   string
	secretFile string
	mcDir      string
	test       bool
}

func (c *Cloud) newTF() (*terraform, error) {

	topoFile := GetTopoFile(c.config.CloudID)
	secretFile := GetSecretFile(c.config.CloudID)
	mcDir := GetMultiCloudRepodir()

	logger := pkglog.NewFileLogger("reporter", c.config.LogFile)
	pkglog.SetLogLevel(logger, c.config.LogLevel)

	r := report.NewReporter(c.APIServer,
		fmt.Sprintf("%s/%s", defaultCloudResourcePath, c.config.CloudID), logger)

	// create logger for secret
	logger = pkglog.NewFileLogger("terraform", c.config.LogFile)
	pkglog.SetLogLevel(logger, c.config.LogLevel)

	return &terraform{
		cloud:      c,
		log:        logger,
		reporter:   r,
		action:     c.config.Action,
		topoFile:   topoFile,
		secretFile: secretFile,
		mcDir:      mcDir,
		test:       c.config.Test,
	}, nil
}

func (c *Cloud) manageTerraform(action string) error {

	tf, err := c.newTF()
	if err != nil {
		return err
	}

	if action == deleteAction {
		err = tf.destroy()
		if err != nil {
			return err
		}
		return nil
	}

	err = tf.createInputFile()
	if err != nil {
		return err
	}

	err = tf.initalize()
	if err != nil {
		return err
	}

	err = tf.apply()
	if err != nil {
		return err
	}

	return nil

}

func (tf *terraform) createInputFile() error {

	cmd := getGenerateTopologyCmd(tf.mcDir)
	args := strings.Split(fmt.Sprintf("-t %s -s %s",
		tf.topoFile, tf.secretFile), " ")

	tf.log.Infof("Generating input file (cloudType.tf.json)")

	tf.log.Debugf("Command executed: %s %s", cmd,
		strings.Join(args, " "))

	workDir := GetCloudDir(tf.cloud.config.CloudID)
	if tf.test {
		return TestCmdHelper(cmd, args, workDir, testTemplate)
	}

	err := osutil.ExecCmdAndWait(tf.reporter, cmd, args, workDir)
	if err != nil {
		return err
	}
	tf.log.Infof("Successfully generated infut file")
	return nil
}

func (tf *terraform) initalize() error {

	args := []string{"init"}

	tf.log.Debugf("Command being executed : %s %s", tfCommand,
		strings.Join(args, " "))

	workDir := GetCloudDir(tf.cloud.config.CloudID)
	if tf.test {
		return TestCmdHelper(tfCommand, args, workDir, testTemplate)
	}

	err := osutil.ExecCmdAndWait(tf.reporter, tfCommand, args, workDir)
	if err != nil {
		return err
	}
	tf.log.Infof("Terraform initialized successfully")
	return nil
}

func (tf *terraform) apply() error {
	args := []string{"apply", "-auto-approve"}

	tf.log.Infof("Creating cloud infra using terraform")

	tf.log.Debugf("Command being executed : %s %s", tfCommand,
		strings.Join(args, " "))

	workDir := GetCloudDir(tf.cloud.config.CloudID)
	if tf.test {
		return TestCmdHelper(tfCommand, args, workDir, testTemplate)
	}

	err := osutil.ExecCmdAndWait(tf.reporter, tfCommand, args, workDir)
	if err != nil {
		return err
	}
	tf.log.Infof("Terraform created all resources successfully")
	return nil
}

func (tf *terraform) destroy() error {

	// teardown terraform

	args := []string{"destroy", "-force"}

	tf.log.Debugf("Command being executed : %s %s", tfCommand,
		strings.Join(args, " "))

	workDir := GetCloudDir(tf.cloud.config.CloudID)
	if tf.test {
		return TestCmdHelper(tfCommand, args, workDir, testTemplate)
	}

	err := osutil.ExecCmdAndWait(tf.reporter, tfCommand, args, workDir)
	if err != nil {
		return err
	}
	tf.log.Infof("Terraform destroyed all resources successfully")

	return nil
}
