package cloud

import (
	"fmt"
	"strings"

	"github.com/sirupsen/logrus"

	report "github.com/Juniper/contrail/pkg/cluster"
	pkglog "github.com/Juniper/contrail/pkg/log"
	"github.com/Juniper/contrail/pkg/models"
)

const (
	azureType             = "azure"
	defaultMultiCloudDir  = "/usr/share/contrail/"
	defaultMultiCloudRepo = "contrail-multi-cloud"
)

type azureUser struct {
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

	mcDir, err := getMultiCloudRepodir()
	if err != nil {
		return err
	}

	u.log.Infof("Executing az login command")

	u.log.Debugf("Login command executed: %s %s", cmd,
		strings.Join(args, " "))

	if !u.test {
		err = execCmd(u.reporter, cmd, args, mcDir)
		if err != nil {
			return err
		}
	}

	u.log.Infof("az login command executed successfully")

	return nil

}

func (c *Cloud) authenticate() error {

	user, err := c.newAzureUser()
	if err != nil {
		return err
	}

	err = user.login()
	if err != nil {
		return err
	}

	return nil
}

func (c *Cloud) newAzureUser() (*azureUser, error) {

	logger := pkglog.NewFileLogger("reporter", c.config.LogFile)
	pkglog.SetLogLevel(logger, c.config.LogLevel)

	r := &report.Reporter{
		API:      c.APIServer,
		Resource: fmt.Sprintf("%s/%s", defaultCloudResourcePath, c.config.CloudID),
		Log:      logger,
	}

	// create logger for topology
	logger = pkglog.NewFileLogger("topology", c.config.LogFile)
	pkglog.SetLogLevel(logger, c.config.LogLevel)

	cloud, err := c.getCloudObject()
	if err != nil {
		return nil, err
	}

	user, err := c.getCloudUser(cloud)
	if err != nil {
		return nil, err
	}

	username, password, err := getUserCred(user, azureType)

	if err != nil {
		return nil, err
	}

	return &azureUser{
		user:     user,
		username: username,
		password: password,
		reporter: r,
		log:      logger,
		test:     c.config.Test,
	}, nil

}
