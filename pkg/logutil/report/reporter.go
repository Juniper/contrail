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
	api      *client.HTTP
	resource string
	log      *logrus.Entry
}

// NewReporter creates a reporter object
func NewReporter(apiServer *client.HTTP, resource string, logger *logrus.Entry) *Reporter {
	return &Reporter{
		api:      apiServer,
		resource: resource,
		log:      logger,
	}
}

// ReportStatus reports status to a particular resource
func (r *Reporter) ReportStatus(ctx context.Context, status map[string]interface{}, resource string) {
	data := map[string]map[string]interface{}{resource: status}
	var response interface{}
	//TODO(nati) fixed context
	_, err := r.api.Update(ctx, r.resource, data, &response)
	if err != nil {
		r.log.Infof("update %s status failed: %s", resource, err)
	}
	r.log.Infof("%s status updated: %s", resource, status)
}

// ReportLog reports log
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
