package cloud

import (
	"fmt"
	"strings"

	"github.com/sirupsen/logrus"

	report "github.com/Juniper/contrail/pkg/cluster"
	pkglog "github.com/Juniper/contrail/pkg/log"
)

const (
	tfCommand = "terraform"
)

type terraform struct {
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
		log:        logger,
		reporter:   r,
		action:     c.config.Action,
		topoFile:   topoFile,
		secretFile: secretFile,
		mcDir:      mcDir,
		test:       c.config.Test,
	}, nil
}

func (c *Cloud) manageTerraform(t *topology, s *secret) error {

	tf, err := c.newTF(t, s)
	if err != nil {
		return err
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

	tf.log.Infof("Generating input file (*.tf.json)")

	tf.log.Debugf("Command executed : %s %s", cmd,
		strings.Join(args, " "))

	if !tf.test {
		err := execCmd(tf.reporter, cmd, args, tf.mcDir)
		if err != nil {
			return err
		}
	}

	tf.log.Infof("Successfully generated infut file")

	return nil
}

func (tf *terraform) initalize() error {

	args := []string{"init"}

	tf.log.Debugf("Command executed : %s %s", tfCommand,
		strings.Join(args, " "))

	if !tf.test {
		err := execCmd(tf.reporter, tfCommand, args, tf.mcDir)
		if err != nil {
			return err
		}
	}

	tf.log.Infof("Terraform initialized successfully")

	return nil
}

func (tf *terraform) apply() error {
	args := []string{"apply", "-auto-approve"}

	tf.log.Debugf("Command executed : %s %s", tfCommand,
		strings.Join(args, " "))

	if !tf.test {
		err := execCmd(tf.reporter, tfCommand, args, tf.mcDir)
		if err != nil {
			return err
		}
	}

	tf.log.Infof("Terraform created all resources successfully")

	return nil

}
