package apisrv

// import (
// 	"encoding/base64"
// 	"fmt"
// 	"net/http"

// 	"github.com/Juniper/contrail/pkg/cloud"

// 	"github.com/labstack/echo"
// )

// //API Path definitions.
// const (
// 	MultiCloudLogUploadPath = "multicloud-logs-upload"
// )

// type MultiCloudLogFiles struct {
// 	Cloud   string
// 	Cluster string
// }

// // RESTUploadMultiCloudLogs handles an /upload-cloud-keys REST request.
// func RESTUploadMultiCloudLogs(c echo.Context) error {
// 	response, err := UploadMultiCloudLogs()
// 	if err != nil {
// 		return echo.NewHTTPError(http.StatusBadRequest,
// 			fmt.Sprintf("cannot load log files: %v", err))
// 	}
// 	return c.JSON(http.StatusOK, response)
// }

// // UploadMultiCloudLogs sends back log files.
// func UploadMultiCloudLogs() (*MultiCloudLogFiles, error) {
// 	clusterLog, err := cloud.ClusterLog()
// 	if err != nil {
// 		return nil, err
// 	}

// 	cloudLog, err := cloud.CloudLog()
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &MultiCloudLogFiles{
// 		Cloud:   base64.StdEncoding.EncodeToString(clusterLog),
// 		Cluster: base64.StdEncoding.EncodeToString(cloudLog),
// 	}, nil
// }
