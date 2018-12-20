package services

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	yaml "gopkg.in/yaml.v2"

	"github.com/Juniper/contrail/pkg/models"
)

type testDataList struct {
	structure ListFloatingIPPoolResponse
	bytes     []byte
}

var dataJSON = []testDataList{
	{
		ListFloatingIPPoolResponse{
			FloatingIPPools:     []*models.FloatingIPPool{{UUID: "vn_uuid"}},
			FloatingIPPoolCount: 1,
		},
		[]byte(`{"floating-ip-pools": [{"uuid": "vn_uuid"}]}`),
	}, {
		ListFloatingIPPoolResponse{
			FloatingIPPools:     make([]*models.FloatingIPPool, 0),
			FloatingIPPoolCount: 0,
		},
		[]byte(`{"floating-ip-pools": []}`),
	},
}

var detailedJSON = []testDataList{
	{
		ListFloatingIPPoolResponse{
			FloatingIPPools:     []*models.FloatingIPPool{{UUID: "vn_uuid"}},
			FloatingIPPoolCount: 1,
		},
		[]byte(`{"floating-ip-pools": [{"floating-ip-pool": {"uuid": "vn_uuid"}}]}`),
	}, {
		ListFloatingIPPoolResponse{
			FloatingIPPools:     make([]*models.FloatingIPPool, 0),
			FloatingIPPoolCount: 0,
		},
		[]byte(`{"floating-ip-pools": []}`),
	}, {
		ListFloatingIPPoolResponse{
			FloatingIPPools:     nil,
			FloatingIPPoolCount: 0,
		},
		[]byte(`{"floating-ip-pools": []}`),
	},
}

var dataCountJSON = []testDataList{
	{
		ListFloatingIPPoolResponse{
			FloatingIPPools:     []*models.FloatingIPPool{{UUID: "vn_uuid"}},
			FloatingIPPoolCount: 1,
		},
		[]byte(`{"floating-ip-pools": {"count": 1}}`),
	}, {
		ListFloatingIPPoolResponse{
			FloatingIPPools:     make([]*models.FloatingIPPool, 0),
			FloatingIPPoolCount: 0,
		},
		[]byte(`{"floating-ip-pools": {"count": 0}}`),
	}, {
		ListFloatingIPPoolResponse{
			FloatingIPPools:     nil,
			FloatingIPPoolCount: 0,
		},
		[]byte(`{"floating-ip-pools": {"count": 0}}`),
	},
}

var dataYAML = []testDataList{
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
  href: ""
  floating_ip_pool_subnets: null
  tag_refs: []
  project_back_refs: []
  floating_ips: []
`),
	},
	{
		ListFloatingIPPoolResponse{
			FloatingIPPools:     make([]*models.FloatingIPPool, 0),
			FloatingIPPoolCount: 0,
		},
		[]byte(`floating-ip-pools: []
`),
	},
}

func TestListResponseJSONMarshaling(t *testing.T) {
	for _, data := range dataJSON {
		dataBytes, err := json.Marshal(data.structure.Data())
		assert.NoError(t, err, "marshaling ListResponse.Data() failed")
		assert.JSONEq(t, string(data.bytes), string(dataBytes))
	}
}

func TestListDetailedResponseJSONMarshaling(t *testing.T) {
	for _, data := range detailedJSON {
		dataBytes, err := json.Marshal(data.structure.Detailed())
		assert.NoError(t, err, "marshaling ListResponse.Detailed() failed")
		fmt.Println("data:", string(dataBytes))
		assert.JSONEq(t, string(data.bytes), string(dataBytes))
	}
}

func TestListCountResponseJSONMarshaling(t *testing.T) {
	for _, data := range dataCountJSON {
		dataBytes, err := json.Marshal(data.structure.Count())
		assert.NoError(t, err, "marshaling ListResponse.Count() failed")
		fmt.Println("data:", string(dataBytes))
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
		dataBytes, err := yaml.Marshal(data.structure.Data())
		assert.NoError(t, err, "marshaling ListResponse failed")
		assert.Equal(t, data.bytes, dataBytes)
	}
}

func TestListResponseYAMLUnmarshaling(t *testing.T) {
	for _, data := range dataYAML {
		var dataStruct ListFloatingIPPoolResponse
		err := yaml.UnmarshalStrict(data.bytes, &dataStruct)
		assert.NoError(t, err, "unmarshaling ListResponse failed")
		assert.EqualValues(t, len(data.structure.FloatingIPPools), len(dataStruct.FloatingIPPools))
		for i := range data.structure.FloatingIPPools {
			assert.EqualValues(t, data.structure.FloatingIPPools[i], dataStruct.FloatingIPPools[i])

		}
	}
}
