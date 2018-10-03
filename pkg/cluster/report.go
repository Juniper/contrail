package cluster

import (
	"bufio"
	"bytes"
	"context"
	"io"

	"github.com/sirupsen/logrus"

	"github.com/Juniper/contrail/pkg/apisrv/client"
)

// Reporter reports provisioing status and log
type Reporter struct {
	API      *client.HTTP
	Resource string
	Log      *logrus.Entry
}

//ReportStatus reports status to APIServer
func (r *Reporter) ReportStatus(status map[string]interface{}, resource string) {
	data := map[string]map[string]interface{}{resource: status}
	var response interface{}
	//TODO(nati) fixed context
	_, err := r.API.Update(context.Background(), r.Resource, data, &response)
	if err != nil {
		r.Log.Infof("update cluster status failed: %s", err)
	}
	r.Log.Infof("cluster status updated: %s", status)
}

//ReportLog reports log
func (r *Reporter) ReportLog(stdio io.Reader) {
	var output bytes.Buffer
	scanner := bufio.NewScanner(stdio)
	for scanner.Scan() {
		m := scanner.Text()
		output.WriteString(m)
		r.Log.Info(m)
	}
	//TODO(ijohnson) Implement status update to db
}
