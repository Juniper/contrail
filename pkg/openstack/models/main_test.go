package models_test

import (
	"encoding/json"
	"fmt"
	"github.com/Juniper/contrail/pkg/format"
	"github.com/Juniper/contrail/pkg/testutil"
	"github.com/stretchr/testify/require"
	"testing"

	"github.com/Juniper/contrail/pkg/openstack/models"
	"github.com/Juniper/contrail/pkg/openstack/neutron"

	"github.com/Juniper/contrail/pkg/fileutil"
	"github.com/Juniper/contrail/pkg/testutil/integration"
)

var server *integration.APIServer

func TestMain(m *testing.M) {
	integration.TestMain(m, &server)
}

func loadRequestFromJSONFile(t *testing.T, path string) (*neutron.Request) {
	var rawJSON map[string]json.RawMessage
	require.NoError(t, fileutil.LoadFile(path, &rawJSON))
	request := neutron.NewRequest()
	err := parseField(rawJSON, "context", &request.Context)
	require.NoError(t, err, "failed to load request. invalid context")
	resource, err := models.GetResource(request.Context.Type)
	require.NoError(t, err)
	request.Data.Resource = resource
	err = parseField(rawJSON, "data", &request.Data)
	require.NoError(t, err, "failed to load request. invalid data")
	return request
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

func assertEqual(t *testing.T, expected, actual interface{}) {
	testutil.AssertEqual(t, format.MustYAML(expected), format.MustYAML(actual))
}
