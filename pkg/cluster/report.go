package cluster

import (
	"bufio"
	"bytes"
	"io"

	"github.com/Juniper/contrail/pkg/apisrv"
	"github.com/sirupsen/logrus"
)

// Reporter reports provisioing status and log
type Reporter struct {
	api      *apisrv.Client
	resource string
	log      *logrus.Entry
}

func (r *Reporter) reportStatus(status string) {
	//TODO(ijohnson) Implement status update
}

func (r *Reporter) reportLog(stdout io.ReadCloser) {
	var output bytes.Buffer
	scanner := bufio.NewScanner(stdout)
	for scanner.Scan() {
		m := scanner.Text()
		output.WriteString(m) // nolint
		r.log.Info(m)
	}
	//TODO(ijohnson) Implement status update to db
}
