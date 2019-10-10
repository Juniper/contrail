package cloud

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"

	"github.com/labstack/echo"
	"github.com/pkg/errors"

	yaml "gopkg.in/yaml.v2"
)

const (
	// MultiCloudLogPathsFile contains paths of Cloud and Cluster log files.
	MultiCloudLogPathsFile = "/var/tmp/contrail/mc_log_paths.yml"
	// MultiCloudLogUploadPath describes path of multicloud-logs-upload endpoint
	MultiCloudLogUploadPath = "multicloud-logs-upload"
)

// MultiCloudLogFiles something
type MultiCloudLogFiles struct {
	Cloud   string
	Cluster string
}

// RESTUploadMultiCloudLogs handles an /upload-cloud-keys REST request.
func RESTUploadMultiCloudLogs(c echo.Context) error {
	response, err := uploadMultiCloudLogs()
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest,
			fmt.Sprintf("cannot load log files: %v", err))
	}
	return c.JSON(http.StatusOK, response)
}

func uploadMultiCloudLogs() (*MultiCloudLogFiles, error) {
	clusterLog, err := clusterLog()
	if err != nil {
		return nil, err
	}

	cloudLog, err := cloudLog()
	if err != nil {
		return nil, err
	}

	return &MultiCloudLogFiles{
		Cloud:   base64.StdEncoding.EncodeToString(cloudLog),
		Cluster: base64.StdEncoding.EncodeToString(clusterLog),
	}, nil
}

type multiCloudLogFilepaths struct {
	CloudPath   string `yaml:"cloud_path"`
	ClusterPath string `yaml:"cluster_path"`
}

// SetCloudLogFilepath sets path of active Cloud log file.
func SetCloudLogFilepath(path string) error {
	return saveMultiCloudLogFilepaths(path, "")
}

// SetClusterLogFilepath sets path of active Cluster log file.
func SetClusterLogFilepath(path string) error {
	return saveMultiCloudLogFilepaths("", path)
}

func saveMultiCloudLogFilepaths(cloudPath, clusterPath string) error {
	paths := multiCloudLogFilepaths{
		CloudPath:   cloudPath,
		ClusterPath: clusterPath,
	}
	if _, err := os.Stat(MultiCloudLogPathsFile); err != nil {
		return updateMultiCloudLogFilepath(paths)
	} else if os.IsNotExist(err) {
		return createMultiCloudLogFilepath(paths)
	} else {
		return errors.Wrap(err, "Cannot save MultiCloud log filepaths")
	}
}

func createMultiCloudLogFilepath(paths multiCloudLogFilepaths) error {
	if err := os.MkdirAll(filepath.Base(MultiCloudLogPathsFile), os.ModePerm); err != nil {
		return errors.Wrap(err, "Cannot ensure directory for MultiCloud log paths exist")
	}
	content, err := yaml.Marshal(paths)
	if err != nil {
		return errors.Wrap(err, "Cannot marshal MultiCloud log paths")
	}
	if err = ioutil.WriteFile(MultiCloudLogPathsFile, content, defaultRWOnlyPerm); err != nil {
		return errors.Wrap(err, "Cannot save MultiCloud log paths file")
	}
	return nil
}

func updateMultiCloudLogFilepath(paths multiCloudLogFilepaths) error {
	oldPaths, err := logFilepaths()
	if err != nil {
		errors.Wrap(err, "Couldn't resolve old MultiCloud logs filepath")
	}
	if paths.CloudPath != "" {
		oldPaths.CloudPath = paths.CloudPath
	}
	if paths.ClusterPath != "" {
		oldPaths.ClusterPath = paths.ClusterPath
	}
	if err = os.Remove(MultiCloudLogPathsFile); err != nil {
		return errors.Wrap(err, "Couldn't update MultiCloud logs filepath")
	}
	return createMultiCloudLogFilepath(oldPaths)
}

func logFilepaths() (multiCloudLogFilepaths, error) {
	content, err := ioutil.ReadFile(MultiCloudLogPathsFile)
	if err != nil {
		return multiCloudLogFilepaths{}, errors.Wrap(err, "Couldn't resolve Cloud log filepath")
	}

	var paths multiCloudLogFilepaths
	err = yaml.Unmarshal(content, &paths)

	return paths, err
}

func clusterLogFilepath() (string, error) {
	paths, err := logFilepaths()
	return paths.ClusterPath, err
}

func cloudLogFilepath() (string, error) {
	paths, err := logFilepaths()
	return paths.CloudPath, err
}

func cloudLog() ([]byte, error) {
	filepath, err := cloudLogFilepath()
	if err != nil {
		return nil, errors.Wrap(err, "Cannot load log file")
	}
	content, err := ioutil.ReadFile(filepath)
	if err != nil {
		return nil, errors.Wrap(err, "Cannot load log file")
	}
	return content, nil
}

func clusterLog() ([]byte, error) {
	filepath, err := clusterLogFilepath()
	if err != nil {
		return nil, errors.Wrap(err, "Cannot load log file")
	}
	content, err := ioutil.ReadFile(filepath)
	if err != nil {
		return nil, errors.Wrap(err, "Cannot load log file")
	}
	return content, nil
}
