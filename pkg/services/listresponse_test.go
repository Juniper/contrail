package services

import (
	"encoding/json"
	"testing"

	"github.com/Juniper/contrail/pkg/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	yaml "gopkg.in/yaml.v2"
)

func TestListResponseEncoding(t *testing.T) {
	vn := &models.VirtualNetwork{UUID: "vn_uuid"}
	listResponse := &ListVirtualNetworkResponse{
		VirtualNetworks:     []*models.VirtualNetwork{vn},
		VirtualNetworkCount: 1,
	}
	data, err := json.Marshal(listResponse)
	require.NoError(t, err, "marshal listResponse failed")

	{
		var list_raw map[string]interface{}
		err = json.Unmarshal(data, &list_raw)
		require.NoError(t, err, "unmarshal listResponse raw failed")

		expected := map[string]interface{}{
			"virtual-networks": []interface{}{
				map[string]interface{}{"uuid": "vn_uuid"},
			},
		}
		assert.Equal(t, expected, list_raw)
	}

	var list_struct ListVirtualNetworkResponse
	err = json.Unmarshal(data, &list_struct)
	require.NoError(t, err, "unmarshal listResponse struct failed")

	assert.Equal(t, 1, int(list_struct.VirtualNetworkCount))
	require.Equal(t, 1, len(list_struct.VirtualNetworks))
	assert.Equal(t, vn, list_struct.VirtualNetworks[0])

	listResponse.VirtualNetworks = nil
	data, err = json.Marshal(listResponse)
	require.NoError(t, err, "marshal listResponse (count) failed")

	var count_raw map[string]interface{}
	err = json.Unmarshal(data, &count_raw)
	require.NoError(t, err, "unmarshal listResponse (count) raw failed")

	expected := map[string]interface{}{
		"virtual-networks": map[string]interface{}{
			"count": 1.,
		},
	}
	assert.Equal(t, expected, count_raw)

	var count_struct ListVirtualNetworkResponse
	err = json.Unmarshal(data, &count_struct)
	require.NoError(t, err, "unmarshal listResponse (count) struct failed")

	assert.Equal(t, listResponse, &count_struct)
}

func TestListResponseYAMLEncoding(t *testing.T) {
	vn := &models.VirtualNetwork{UUID: "vn_uuid"}
	listResponse := &ListVirtualNetworkResponse{
		VirtualNetworks:     []*models.VirtualNetwork{vn},
		VirtualNetworkCount: 1,
	}
	data, err := yaml.Marshal(listResponse)
	require.NoError(t, err, "marshal listResponse failed")

	var list_raw struct {
		VirtualNetworks []*models.VirtualNetwork `yaml:"virtual-networks"`
	}
	err = yaml.Unmarshal(data, &list_raw)
	require.NoError(t, err, "unmarshal listResponse raw failed")

	require.Equal(t, 1, len(list_raw.VirtualNetworks))
	assert.Equal(t, vn.UUID, list_raw.VirtualNetworks[0].UUID)

	var list_struct ListVirtualNetworkResponse
	err = yaml.Unmarshal(data, &list_struct)
	require.NoError(t, err, "unmarshal listResponse struct failed")

	assert.Equal(t, 1, int(list_struct.VirtualNetworkCount))
	require.Equal(t, 1, len(list_struct.VirtualNetworks))
	assert.Equal(t, vn.UUID, list_struct.VirtualNetworks[0].UUID)

	listResponse.VirtualNetworks = nil
	data, err = yaml.Marshal(listResponse)
	require.NoError(t, err, "marshal listResponse (count) failed")

	var count_raw map[interface{}]interface{}
	err = yaml.Unmarshal(data, &count_raw)
	require.NoError(t, err, "unmarshal listResponse (count) raw failed")

	expected := map[interface{}]interface{}{
		"virtual-networks": map[interface{}]interface{}{
			"count": 1,
		},
	}
	assert.Equal(t, expected, count_raw)

	var count_struct ListVirtualNetworkResponse
	err = yaml.Unmarshal(data, &count_struct)
	require.NoError(t, err, "unmarshal listResponse (count) struct failed")

	assert.Equal(t, listResponse, &count_struct)
}
