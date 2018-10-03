package cloud

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/sirupsen/logrus"

	report "github.com/Juniper/contrail/pkg/cluster"
	pkglog "github.com/Juniper/contrail/pkg/log"
)

const (
	tfCommand   = "terraform"
	tfStateFile = "terraform.tfstate"
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

func (c *Cloud) newTF(t *topology, s *secret) (*terraform, error) {

	topoFile := t.getTopoFile()
	secretFile := s.getSecretFile()
	mcDir, err := getMultiCloudRepodir()
	if err != nil {
		return nil, err
	}

	logger := pkglog.NewFileLogger("reporter", c.config.LogFile)
	pkglog.SetLogLevel(logger, c.config.LogLevel)

	r := &report.Reporter{
		API:      c.APIServer,
		Resource: fmt.Sprintf("%s/%s", defaultCloudResourcePath, c.config.CloudID),
		Log:      logger,
	}

	// create logger for secret
	logger = pkglog.NewFileLogger("topology", c.config.LogFile)
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

func (c *Cloud) manageTerraform(t *topology, s *secret, action string) error {

	tf, err := c.newTF(t, s)
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

	cmd := getGenTopoFile(tf.mcDir)
	args := strings.Split(fmt.Sprintf("-t %s -s %s",
		tf.topoFile, tf.secretFile), " ")

	tf.log.Infof("Generating input file (cloudType.tf.json)")

	tf.log.Debugf("Command executed: %s %s", cmd,
		strings.Join(args, " "))

	workDir := tf.cloud.getHomeDir()
	if tf.test {
		return testHelper(cmd, args, workDir)
	}

	err := execCmd(tf.reporter, cmd, args, workDir)
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

	workDir := tf.cloud.getHomeDir()
	if tf.test {
		return testHelper(tfCommand, args, workDir)
	}

	err := execCmd(tf.reporter, tfCommand, args, workDir)
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

	workDir := tf.cloud.getHomeDir()
	if tf.test {
		return testHelper(tfCommand, args, workDir)
	}

	err := execCmd(tf.reporter, tfCommand, args, workDir)
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

	workDir := tf.cloud.getHomeDir()
	if tf.test {
		return testHelper(tfCommand, args, workDir)
	}

	err := execCmd(tf.reporter, tfCommand, args, workDir)
	if err != nil {
		return err
	}
	tf.log.Infof("Terraform destroyed all resources successfully")

	// Delete the tfStateFile created by terraform
	args = []string{filepath.Join(workDir, tfStateFile)}
	if !tf.test {
		err := deleteFiles(tf.reporter, args)
		if err != nil {
			return err
		}
	}

	return nil
}
