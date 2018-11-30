package neutron

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/Juniper/contrail/pkg/fileutil"
	"github.com/Juniper/contrail/pkg/openstack/models"
	"github.com/Juniper/contrail/pkg/testutil/integration"
)

var server *integration.APIServer

func TestMain(m *testing.M) {
	integration.TestMain(m, &server)
}

func loadRequestFromJSONFile(path string) (*Request, error) {
	var rawJSON map[string]json.RawMessage
	fileutil.LoadFile(path, rawJSON)
	request := NewRequest()
	if err := parseField(rawJSON, "context", &request.Context); err != nil {
		return nil, err
	}
	r, err := models.GetResource(request.Context.Type)
	if err != nil {
		return nil, err
	}
	request.Data.Resource = r
	if err := parseField(rawJSON, "data", &request.Data); err != nil {
		return nil, err
	}
	return request, nil
}

func parseField(rawJSON map[string]json.RawMessage, key string, dst interface{}) error {
	if val, ok := rawJSON[key]; ok {
		if err := json.Unmarshal(val, dst); err != nil {
			return fmt.Errorf("invalid '%s' format: %v", key, err)
		}
		delete(rawJSON, key)
	}
	return nil
}
