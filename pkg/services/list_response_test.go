package services

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v2"

	"github.com/Juniper/contrail/pkg/models"
)

var dataJSON = []struct {
	structure ListFloatingIPPoolResponse
	bytes     []byte
}{
	{
		ListFloatingIPPoolResponse{
			FloatingIPPools:     []*models.FloatingIPPool{{UUID: "vn_uuid"}},
			FloatingIPPoolCount: 1,
		},
		[]byte(`{"floating-ip-pools": [{"uuid": "vn_uuid"}]}`),
	}, {
		ListFloatingIPPoolResponse{
			FloatingIPPools:     nil,
			FloatingIPPoolCount: 1,
		},
		[]byte(`{"floating-ip-pools": {"count": 1}}`),
	},
}

var dataYAML = []struct {
	structure ListFloatingIPPoolResponse
	bytes     []byte
}{
	{
		ListFloatingIPPoolResponse{
			FloatingIPPools: []*models.FloatingIPPool{
				// We initialize the lists for comparing since yaml.v2 marshals null lists to empty lists
				{
					UUID:            "vn_uuid",
					FQName:          []string{},
					ProjectBackRefs: []*models.Project{},
					FloatingIPs:     []*models.FloatingIP{},
					TagRefs:         []*models.FloatingIPPoolTagRef{},
				},
			},
			FloatingIPPoolCount: 1,
		},
		[]byte(`floating-ip-pools:
- uuid: vn_uuid
  name: ""
  parent_uuid: ""
  parent_type: ""
  fq_name: []
  id_perms: null
  display_name: ""
  annotations: null
  perms2: null
  configuration_version: 0
  floating_ip_pool_subnets: null
  tag_refs: []
  project_backrefs: []
  floating_ips: []
`),
	}, {
		ListFloatingIPPoolResponse{
			FloatingIPPools:     nil,
			FloatingIPPoolCount: 1,
		},
		[]byte(`floating-ip-pools:
  count: 1
`),
	},
}

func TestListResponseJSONMarshaling(t *testing.T) {
	for _, data := range dataJSON {
		dataBytes, err := json.Marshal(data.structure)
		assert.NoError(t, err, "marshaling ListResponse failed")
		assert.JSONEq(t, string(data.bytes), string(dataBytes))
	}
}

func TestListResponseJSONUnmarshaling(t *testing.T) {
	for _, data := range dataJSON {
		var dataStruct ListFloatingIPPoolResponse
		err := json.Unmarshal(data.bytes, &dataStruct)
		assert.NoError(t, err, "unmarshaling ListResponse failed")
		assert.Equal(t, data.structure, dataStruct)
	}
}

func TestListResponseYAMLMarshaling(t *testing.T) {
	for _, data := range dataYAML {
		dataBytes, err := yaml.Marshal(data.structure)
		assert.NoError(t, err, "marshaling ListResponse failed")
		assert.Equal(t, data.bytes, dataBytes)
	}
}

func TestListResponseYAMLUnmarshaling(t *testing.T) {
	for _, data := range dataYAML {
		var dataStruct ListFloatingIPPoolResponse
		err := yaml.Unmarshal(data.bytes, &dataStruct)
		assert.NoError(t, err, "unmarshaling ListResponse failed")
		assert.EqualValues(t, len(data.structure.FloatingIPPools), len(dataStruct.FloatingIPPools))
		for i := range data.structure.FloatingIPPools {
			assert.EqualValues(t, data.structure.FloatingIPPools[i], dataStruct.FloatingIPPools[i])

		}
	}
}
