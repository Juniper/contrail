package services

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gopkg.in/yaml.v2"

	"github.com/Juniper/contrail/pkg/models"
)

func TestListResponseEncoding(t *testing.T) {
	vn := &models.VirtualNetwork{UUID: "vn_uuid"}
	listResponse := &ListVirtualNetworkResponse{
		VirtualNetworks:     []*models.VirtualNetwork{vn},
		VirtualNetworkCount: 1,
	}
	data, err := json.Marshal(listResponse)
	require.NoError(t, err, "marshaling listResponse failed")

	var listRaw map[string]interface{}
	err = json.Unmarshal(data, &listRaw)
	require.NoError(t, err, "unmarshaling listResponse raw failed")

	expected := map[string]interface{}{
		"virtual-networks": []interface{}{
			map[string]interface{}{"uuid": "vn_uuid"},
		},
	}
	assert.Equal(t, expected, listRaw)

	var listStruct ListVirtualNetworkResponse
	err = json.Unmarshal(data, &listStruct)
	require.NoError(t, err, "unmarshaling listResponse struct failed")

	assert.Equal(t, 1, int(listStruct.VirtualNetworkCount))
	require.Equal(t, 1, len(listStruct.VirtualNetworks))
	assert.Equal(t, vn, listStruct.VirtualNetworks[0])
}

func TestCountResponseEncoding(t *testing.T) {
	listResponse := &ListVirtualNetworkResponse{
		VirtualNetworks:     nil,
		VirtualNetworkCount: 1,
	}

	data, err := json.Marshal(listResponse)
	require.NoError(t, err, "marshaling listResponse (count) failed")

	var countRaw map[string]interface{}
	err = json.Unmarshal(data, &countRaw)
	require.NoError(t, err, "unmarshaling listResponse (count) raw failed")

	expected := map[string]interface{}{
		"virtual-networks": map[string]interface{}{
			"count": 1.,
		},
	}
	assert.Equal(t, expected, countRaw)

	var countStruct ListVirtualNetworkResponse
	err = json.Unmarshal(data, &countStruct)
	require.NoError(t, err, "unmarshaling listResponse (count) struct failed")

	assert.Equal(t, listResponse, &countStruct)
}

func TestListResponseYAMLEncoding(t *testing.T) {
	vn := &models.VirtualNetwork{UUID: "vn_uuid"}
	listResponse := &ListVirtualNetworkResponse{
		VirtualNetworks:     []*models.VirtualNetwork{vn},
		VirtualNetworkCount: 1,
	}
	data, err := yaml.Marshal(listResponse)
	require.NoError(t, err, "marshaling listResponse failed")

	var listRaw struct {
		VirtualNetworks []*models.VirtualNetwork `yaml:"virtual-networks"`
	}
	err = yaml.Unmarshal(data, &listRaw)
	require.NoError(t, err, "unmarshaling listResponse raw failed")

	require.Equal(t, 1, len(listRaw.VirtualNetworks))
	assert.Equal(t, vn.UUID, listRaw.VirtualNetworks[0].UUID)

	var listStruct ListVirtualNetworkResponse
	err = yaml.Unmarshal(data, &listStruct)
	require.NoError(t, err, "unmarshaling listResponse struct failed")

	assert.Equal(t, 1, int(listStruct.VirtualNetworkCount))
	require.Equal(t, 1, len(listStruct.VirtualNetworks))
	assert.Equal(t, vn.UUID, listStruct.VirtualNetworks[0].UUID)
}

func TestCountResponseYAMLEncoding(t *testing.T) {
	listResponse := &ListVirtualNetworkResponse{
		VirtualNetworks:     nil,
		VirtualNetworkCount: 1,
	}

	data, err := yaml.Marshal(listResponse)
	require.NoError(t, err, "marshaling listResponse (count) failed")

	var countRaw map[interface{}]interface{}
	err = yaml.Unmarshal(data, &countRaw)
	require.NoError(t, err, "unmarshaling listResponse (count) raw failed")

	expected := map[interface{}]interface{}{
		"virtual-networks": map[interface{}]interface{}{
			"count": 1,
		},
	}
	assert.Equal(t, expected, countRaw)

	var countStruct ListVirtualNetworkResponse
	err = yaml.Unmarshal(data, &countStruct)
	require.NoError(t, err, "unmarshaling listResponse (count) struct failed")

	assert.Equal(t, listResponse, &countStruct)
}
