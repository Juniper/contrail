package report

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
	APIServer    *client.HTTP
	Resource     string
	ResourcePath string
	Log          *logrus.Entry
}

func (r *Reporter) ReportStatus(status map[string]interface{}) {
	data := map[string]map[string]interface{}{r.Resource: status}
	var response interface{}
	//TODO(nati) fixed context
	_, err := r.APIServer.Update(context.Background(), r.ResourcePath, data, &response)
	if err != nil {
		r.Log.Infof("update cluster status failed: %s", err)
	}
	r.Log.Infof("cluster status updated: %s", status)
}

func (r *Reporter) ReportLog(stdio io.Reader) {
	var output bytes.Buffer
	scanner := bufio.NewScanner(stdio)
	for scanner.Scan() {
		m := scanner.Text()
		output.WriteString(m)
		r.Log.Info(m)
	}
}
