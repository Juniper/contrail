package apisrv_test

import (
	"context"
	"encoding/base64"
	"io/ioutil"
	"net/http"
	"os"
	"testing"

	"github.com/Juniper/contrail/pkg/cloud"
	"github.com/Juniper/contrail/pkg/testutil/integration"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMultiCloudLogUpload(t *testing.T) {
	hc, err := integration.NewHTTPClient(server.URL())
	require.NoError(t, err)

	logFiles := cloud.MultiCloudLogFiles{}

	for _, copy := range []struct {
		src  string
		dest string
	}{
		{
			src:  "./test_data/multicloud/test_multicloud_log_filepaths.yml",
			dest: cloud.MultiCloudLogPathsFile,
		},
		{
			src:  "./test_data/multicloud/test_cloud_log_file",
			dest: "/var/tmp/contrail/test_cloud.log",
		},
		{
			src:  "./test_data/multicloud/test_cluster_log_file",
			dest: "/var/tmp/contrail/test_cluster.log",
		},
	} {
		err = copyFile(copy.src, copy.dest)
		require.NoError(t, err)
	}

	_, err = hc.Do(
		context.Background(),
		echo.POST,
		"/"+cloud.MultiCloudLogUploadPath,
		nil,
		nil,
		&logFiles,
		[]int{http.StatusOK},
	)
	assert.NoError(t, err)

	encodedCloudLogFile, err := base64.StdEncoding.DecodeString(logFiles.Cloud)
	assert.NoError(t, err)
	assert.Equal(t, "cloud_log_test", string(encodedCloudLogFile))

	encodedClusterLogFile, err := base64.StdEncoding.DecodeString(logFiles.Cluster)
	assert.NoError(t, err)
	assert.Equal(t, "cluster_log_test", string(encodedClusterLogFile))
}

func copyFile(src, dest string) error {
	content, err := ioutil.ReadFile(src)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(dest, content, os.ModePerm)
}
