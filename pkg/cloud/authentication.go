package cloud

import (
	"fmt"
	"os"
	"strings"

	"github.com/sirupsen/logrus"

	"github.com/Juniper/contrail/pkg/logutil"
	"github.com/Juniper/contrail/pkg/logutil/report"
	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/osutil"
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
	err = osutil.ExecCmdAndWait(u.reporter, cmd, args, workDir)
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

	err = user.login()
	if err != nil {
		return err
	}

	//TODO(madhukar) to select the right subscription

	return nil
}

func (c *Cloud) newAzureUser(d *Data) (*azureUser, error) {
	r := report.NewReporter(
		c.APIServer,
		fmt.Sprintf("%s/%s", defaultCloudResourcePath, c.config.CloudID),
		logutil.NewFileLogger("reporter", c.config.LogFile),
	)

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
		log:      logutil.NewFileLogger("topology", c.config.LogFile),
		test:     c.config.Test,
	}, nil
}
