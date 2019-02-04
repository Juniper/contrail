package cloud

import (
	"fmt"
	"os"
	"strings"

	"github.com/sirupsen/logrus"

	"github.com/Juniper/contrail/pkg/common"
	pkglog "github.com/Juniper/contrail/pkg/log"
	"github.com/Juniper/contrail/pkg/log/report"
	"github.com/Juniper/contrail/pkg/models"
)

type azureUser struct {
	cloud    *Cloud
	user     *models.CloudUser
	username string
	password string
	reporter *report.Reporter
	log      *logrus.Entry
	test     bool
}

func (u *azureUser) login() error {
	cmd := "az"
	args := strings.Split(fmt.Sprintf("login -u %s -p %s",
		u.username, u.password), " ")

	u.log.Debugf("Executing az login command")

	workDir := GetCloudDir(u.cloud.config.CloudID)
	err := os.MkdirAll(workDir, os.ModePerm)
	if err != nil {
		return err
	}

	if u.test {
		return TestCmdHelper(cmd, args, workDir, testTemplate)
	}
	err = common.ExecCmdAndWait(u.reporter, cmd, args, workDir)
	if err != nil {
		return err
	}
	u.log.Debugf("az login command executed successfully")
	return nil
}

func (c *Cloud) authenticate(d *Data) error {

	user, err := c.newAzureUser(d)
	if err != nil {
		return err
	}

	return user.login()

	//TODO(madhukar) to select the right subscription
}

func (c *Cloud) newAzureUser(d *Data) (*azureUser, error) {

	logger := pkglog.NewFileLogger("cloud-topology", c.config.LogFile)
	pkglog.SetLogLevel(logger, c.config.LogLevel)

	r := report.NewReporter(c.APIServer,
		fmt.Sprintf("%s/%s", defaultCloudResourcePath, c.config.CloudID), logger)

	// create logger for topology
	logger = pkglog.NewFileLogger("cloud-topology", c.config.LogFile)
	pkglog.SetLogLevel(logger, c.config.LogLevel)

	user, err := getCloudUser(d)
	if err != nil {
		return nil, err
	}

	username, password, err := getUserCred(user)

	if err != nil {
		return nil, err
	}

	return &azureUser{
		cloud:    c,
		user:     user,
		username: username,
		password: password,
		reporter: r,
		log:      logger,
		test:     c.config.Test,
	}, nil

}
