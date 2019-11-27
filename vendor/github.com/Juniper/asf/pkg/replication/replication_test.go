package replication

import (
	"context"
	"testing"

	"github.com/Juniper/asf/pkg/logutil"
	"github.com/Juniper/asf/pkg/models"
	"github.com/Juniper/asf/pkg/services"
	"github.com/stretchr/testify/assert"
)

func TestReplicator_Process(t *testing.T) {
	const vnID = "vn-uuid"

	tests := []struct {
		name     string
		input    *services.Event
		wantCall *replicateCall
		fails    bool
	}{{
		name: "nil",
	}, {
		name: "create ACL is omitted",
		input: &services.Event{Request: &services.Event_CreateAccessControlListRequest{
			CreateAccessControlListRequest: &services.CreateAccessControlListRequest{
				AccessControlList: &models.AccessControlList{},
			},
		}},
	}, {
		name: "create VN is handled",
		input: &services.Event{Request: &services.Event_CreateVirtualNetworkRequest{
			CreateVirtualNetworkRequest: &services.CreateVirtualNetworkRequest{
				VirtualNetwork: &models.VirtualNetwork{},
			},
		}},
		wantCall: &replicateCall{
			action:   createAction,
			url:      createVirtualNetworkURL,
			data:     &services.CreateVirtualNetworkRequest{VirtualNetwork: &models.VirtualNetwork{}},
			response: &services.CreateVirtualNetworkResponse{},
		},
	}, {
		name: "create VN-IPAM ref is handled",
		input: &services.Event{Request: &services.Event_CreateVirtualNetworkNetworkIpamRefRequest{
			CreateVirtualNetworkNetworkIpamRefRequest: &services.CreateVirtualNetworkNetworkIpamRefRequest{ID: vnID},
		}},
		wantCall: &replicateCall{
			action: refUpdateAction,
			url:    services.RefUpdatePath,
			data: services.RefUpdate{
				Operation: services.RefOperationAdd,
				Type:      models.KindVirtualNetwork,
				RefType:   models.KindNetworkIpam,
				UUID:      vnID,
			},
			response: map[string]interface{}{},
		},
	}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &handlerMock{}
			r := &Replicator{handler: h, log: logutil.NewLogger("vnc_replication")}

			_, err := r.Process(context.Background(), tt.input)

			if tt.fails {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.wantCall, h.last)
			}
		})
	}
}

type handlerMock struct {
	last *replicateCall
}

type replicateCall struct {
	action, url    string
	data, response interface{}
}

func (h *handlerMock) CreateClient(ep *models.Endpoint) {}
func (h *handlerMock) UpdateClient(ep *models.Endpoint) {}
func (h *handlerMock) DeleteClient(endpointID string)   {}

func (h *handlerMock) Replicate(action, url string, data interface{}, response interface{}) {
	h.last = &replicateCall{action: action, url: url, data: data, response: response}
}
