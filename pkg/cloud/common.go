package cloud

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"

	report "github.com/Juniper/contrail/pkg/cluster"
	"github.com/Juniper/contrail/pkg/models"
	"github.com/flosch/pongo2"
	"github.com/labstack/echo"
)

const (
	defaultWorkRoot      = "/var/tmp/cloud"
	defaultTemplateRoot  = "./pkg/cloud/configs"
	defaultGenTopoScript = "transform/generate_topology.py"
)

func writeToFile(path string, content []byte) error {
	err := os.MkdirAll(filepath.Dir(path), os.ModePerm)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(path, content, os.FileMode(0600))
	return err
}

func applyTemplate(templateSrc string, context map[string]interface{}) ([]byte, error) {
	template, err := pongo2.FromFile(templateSrc)
	if err != nil {
		return nil, err
	}
	output, err := template.ExecuteBytes(context)
	if err != nil {
		return nil, err
	}
	// strip empty lines in output content
	regex, _ := regexp.Compile("\n[ \r\n\t]*\n") // nolint: errcheck
	outputString := regex.ReplaceAllString(string(output), "\n")
	return []byte(outputString), nil
}

func returnEchoContext(c echo.Context) context.Context {
	return c.Request().Context()
}

func (c *Cloud) getHomeDir() string {
	dir := filepath.Join(c.getTemplateRoot(), c.config.CloudID)
	return dir
}

func (c *Cloud) getWorkingDir() string {
	dir := filepath.Join(c.getHomeDir())
	return dir
}

func (c *Cloud) getTemplateRoot() string {
	templateRoot := c.config.TemplateRoot
	if templateRoot == "" {
		templateRoot = defaultTemplateRoot
	}
	return templateRoot
}

func (c *Cloud) getCloudType() (string, error) {

	cloud, err := c.getCloudObject()
	if err != nil {
		return "", err
	}

	return cloud.Type, nil
}

func (c *Cloud) getCloudUser(cloud *models.Cloud) (*models.CloudUser, error) {

	for _, account := range cloud.CloudAccounts {
		for _, project := range account.CloudProjects {
			for _, userRef := range project.CloudUserRefs {
				userObject, err := getUserObject(userRef.UUID, c)
				if err != nil {
					return nil, err
				}
				return userObject, nil
			}
		}
	}
	return nil, fmt.Errorf("Cloud User not found")
}

func getUserCred(user *models.CloudUser, cloudType string) (username string, password string, err error) {

	if cloudType == azureType {
		username = user.AzureCredential.UsernameCred.Username
		password = user.AzureCredential.UsernameCred.Password

		if username == "" || password == "" {
			return username, password, fmt.Errorf("Username or Password not found for user uuid: %s", user.UUID)
		}

		return username, password, nil
	}

	return username, password, fmt.Errorf("Cannot get user credential for cloud type: %s", cloudType)
}

func getMultiCloudRepodir() (string, error) {
	return filepath.Join(defaultMultiCloudDir, defaultMultiCloudRepo), nil
}

func execCmd(r *report.Reporter, cmd string, args []string, dir string) error {
	cmdline := exec.Command(cmd, args...)
	if dir != "" {
		cmdline.Dir = dir
	}
	stdout, err := cmdline.StdoutPipe()
	if err != nil {
		return err
	}
	stderr, err := cmdline.StderrPipe()
	if err != nil {
		return err
	}
	if err := cmdline.Start(); err != nil {
		return err
	}
	// Report progress log periodically to stdout/db
	go r.ReportLog(stdout)
	go r.ReportLog(stderr)
	return cmdline.Wait()
}

func getGenTopoFile(mcDir string) string {
	return filepath.Join(mcDir, defaultGenTopoScript)
}

func returnContext() context.Context {
	return context.Background()
}
