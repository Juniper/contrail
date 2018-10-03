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
func (r *Reporter) ReportStatus(status map[string]interface{}) {
	data := map[string]map[string]interface{}{defaultResource: status}
	var response interface{}
	//TODO(nati) fixed context
	_, err := r.api.Update(context.Background(), r.resource, data, &response)
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
		r.log.Info(m)
	}
	//TODO(ijohnson) Implement status update to db
}
