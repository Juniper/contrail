package cluster

import (
	"bufio"
	"bytes"
	"io"

	"github.com/sirupsen/logrus"

	"github.com/Juniper/contrail/pkg/apisrv/client"
)

// Reporter reports provisioing status and log
type Reporter struct {
	api      *client.HTTP
	resource string
	log      *logrus.Entry
}

func (r *Reporter) reportStatus(status map[string]interface{}) {
	data := map[string]map[string]interface{}{defaultResource: status}
	var response interface{}
	_, err := r.api.Update(r.resource, data, &response)
	if err != nil {
		r.log.Infof("update cluster status failed: %s", err)
	}
	r.log.Infof("cluster status updated: %s", status)
}

func (r *Reporter) reportLog(stdio io.Reader) {
	var output bytes.Buffer
	scanner := bufio.NewScanner(stdio)
	for scanner.Scan() {
		m := scanner.Text()
		output.WriteString(m) // nolint
		r.log.Info(m)
	}
	//TODO(ijohnson) Implement status update to db
}
