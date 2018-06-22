package services

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/Juniper/contrail/pkg/common"
	"github.com/Juniper/contrail/pkg/models"
	"github.com/stretchr/testify/assert"
	yaml "gopkg.in/yaml.v2"
)

func TestCreateEventJSONEncoding(t *testing.T) {
	e := &Event{
		Request: &Event_CreateVirtualNetworkRequest{
			CreateVirtualNetworkRequest: &CreateVirtualNetworkRequest{
				VirtualNetwork: &models.VirtualNetwork{
					UUID: "vn_uuid",
				},
			},
		},
	}
	m, err := json.Marshal(e)
	assert.NoError(t, err, "marhsal event failed")
	var i map[string]interface{}
	err = json.Unmarshal(m, &i)
	assert.NoError(t, err, "unmarhsal event failed")
	assert.Equal(t, "virtual_network", i["kind"])
	assert.Equal(t, "CREATE", i["operation"])

	var d Event
	err = json.Unmarshal(m, &d)
	assert.NoError(t, err, "unmarhsal event failed")
	request := d.GetCreateVirtualNetworkRequest()
	assert.Equal(t, "vn_uuid", request.GetVirtualNetwork().GetUUID())

	d2 := InterfaceToEvent(i)
	request = d2.GetCreateVirtualNetworkRequest()
	assert.Equal(t, "vn_uuid", request.GetVirtualNetwork().GetUUID())
}

func TestDeleteEventJSONEncoding(t *testing.T) {
	e := &Event{
		Request: &Event_DeleteVirtualNetworkRequest{
			DeleteVirtualNetworkRequest: &DeleteVirtualNetworkRequest{
				ID: "vn_uuid",
			},
		},
	}
	m, err := json.Marshal(e)
	assert.NoError(t, err, "marhsal event failed")
	fmt.Println(string(m))
	var i map[string]interface{}
	err = json.Unmarshal(m, &i)
	assert.NoError(t, err, "unmarhsal event failed")
	assert.Equal(t, "virtual_network", i["kind"])
	assert.Equal(t, "DELETE", i["operation"])
	assert.Equal(t, "vn_uuid", i["data"].(map[string]interface{})["uuid"])

	var d Event
	err = json.Unmarshal(m, &d)
	assert.NoError(t, err, "unmarhsal event failed")
	request := d.GetDeleteVirtualNetworkRequest()
	assert.Equal(t, "vn_uuid", request.ID)
}

func TestCreateEventYAMLEncoding(t *testing.T) {
	t.Skip("TODO: Fix me")
	e := &Event{
		Request: &Event_CreateVirtualNetworkRequest{
			CreateVirtualNetworkRequest: &CreateVirtualNetworkRequest{
				VirtualNetwork: &models.VirtualNetwork{
					UUID: "vn_uuid",
				},
			},
		},
	}
	m, err := yaml.Marshal(e)
	fmt.Println(string(m))
	assert.NoError(t, err, "marhsal event failed")
	var i map[string]interface{}
	err = yaml.Unmarshal(m, &i)
	assert.NoError(t, err, "unmarhsal event failed")
	assert.Equal(t, "virtual_network", i["kind"])
	assert.Equal(t, "CREATE", i["operation"])

	var d Event
	err = yaml.Unmarshal(m, &d)
	assert.NoError(t, err, "unmarhsal event failed")
	request := d.GetCreateVirtualNetworkRequest()
	assert.Equal(t, "vn_uuid", request.GetVirtualNetwork().GetUUID())
	d2 := InterfaceToEvent(common.YAMLtoJSONCompat(i))
	request = d2.GetCreateVirtualNetworkRequest()
	assert.Equal(t, "vn_uuid", request.GetVirtualNetwork().GetUUID())
}

func TestReorderEventList(t *testing.T) {
	eventList := &EventList{
		Events: []*Event{
			&Event{
				Request: &Event_CreateVirtualNetworkRequest{
					&CreateVirtualNetworkRequest{
						VirtualNetwork: &models.VirtualNetwork{
							UUID: "vn1",
							NetworkPolicyRefs: []*models.VirtualNetworkNetworkPolicyRef{
								&models.VirtualNetworkNetworkPolicyRef{
									UUID: "network_policy1",
								},
							},
						},
					},
				},
			},
			&Event{
				Request: &Event_CreateNetworkPolicyRequest{
					CreateNetworkPolicyRequest: &CreateNetworkPolicyRequest{
						NetworkPolicy: &models.NetworkPolicy{
							UUID: "network_policy1",
						},
					},
				},
			},
		},
	}

	err := eventList.Sort()
	assert.NoError(t, err)
	networkPolicy := eventList.Events[0].GetCreateNetworkPolicyRequest().GetNetworkPolicy()
	assert.Equal(t, "network_policy1", networkPolicy.GetUUID())
	virtualNetwork := eventList.Events[1].GetCreateVirtualNetworkRequest().GetVirtualNetwork()
	assert.Equal(t, "vn1", virtualNetwork.GetUUID())
}

func TestReorderLoopedList(t *testing.T) {
	eventList := &EventList{
		Events: []*Event{
			&Event{
				Request: &Event_CreateVirtualNetworkRequest{
					&CreateVirtualNetworkRequest{
						VirtualNetwork: &models.VirtualNetwork{
							UUID: "vn1",
							VirtualNetworkRefs: []*models.VirtualNetworkVirtualNetworkRef{
								&models.VirtualNetworkVirtualNetworkRef{
									UUID: "vn2",
								},
							},
						},
					},
				},
			},
			&Event{
				Request: &Event_CreateVirtualNetworkRequest{
					&CreateVirtualNetworkRequest{
						VirtualNetwork: &models.VirtualNetwork{
							UUID: "vn2",
							VirtualNetworkRefs: []*models.VirtualNetworkVirtualNetworkRef{
								&models.VirtualNetworkVirtualNetworkRef{
									UUID: "vn1",
								},
							},
						},
					},
				},
			},
		},
	}
	err := eventList.Sort()
	assert.Error(t, err)
}
