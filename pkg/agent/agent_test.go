package agent

import (
	"context"
	"testing"

	"github.com/Juniper/asf/pkg/logutil"
	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/services"
	"github.com/stretchr/testify/assert"
)

func TestProcess(t *testing.T) {
	processCases := map[string]struct {
		inputEvent    *services.Event
		expectedEvent *services.Event
		expectedKind  string
	}{
		"Unexpected Event": {
			inputEvent: &services.Event{
				Request: &services.Event_CreateAccessControlListRequest{},
			},
			expectedEvent: nil,
			expectedKind:  "",
		},
		"Contrail Creation Event": {
			inputEvent: &services.Event{
				Request: &services.Event_CreateContrailClusterRequest{
					CreateContrailClusterRequest: &services.CreateContrailClusterRequest{
						ContrailCluster: &models.ContrailCluster{},
					},
				},
			},
			expectedEvent: &services.Event{
				Request: &services.Event_CreateContrailClusterRequest{
					CreateContrailClusterRequest: &services.CreateContrailClusterRequest{
						ContrailCluster: &models.ContrailCluster{},
					},
				},
			},
			expectedKind: "cluster",
		},
		"Contrail Update Event": {
			inputEvent: &services.Event{
				Request: &services.Event_UpdateContrailClusterRequest{
					UpdateContrailClusterRequest: &services.UpdateContrailClusterRequest{
						ContrailCluster: &models.ContrailCluster{},
					},
				},
			},
			expectedEvent: &services.Event{
				Request: &services.Event_UpdateContrailClusterRequest{
					UpdateContrailClusterRequest: &services.UpdateContrailClusterRequest{
						ContrailCluster: &models.ContrailCluster{},
					},
				},
			},
			expectedKind: "cluster",
		},
		"Contrail Deletion Event": {
			inputEvent: &services.Event{
				Request: &services.Event_DeleteContrailClusterRequest{
					DeleteContrailClusterRequest: &services.DeleteContrailClusterRequest{},
				},
			},
			expectedEvent: &services.Event{
				Request: &services.Event_DeleteContrailClusterRequest{
					DeleteContrailClusterRequest: &services.DeleteContrailClusterRequest{},
				},
			},
			expectedKind: "cluster",
		},
		"Rhospd Cloud Manager Creation Event": {
			inputEvent: &services.Event{
				Request: &services.Event_CreateRhospdCloudManagerRequest{
					CreateRhospdCloudManagerRequest: &services.CreateRhospdCloudManagerRequest{
						RhospdCloudManager: &models.RhospdCloudManager{},
					},
				},
			},
			expectedEvent: &services.Event{
				Request: &services.Event_CreateRhospdCloudManagerRequest{
					CreateRhospdCloudManagerRequest: &services.CreateRhospdCloudManagerRequest{
						RhospdCloudManager: &models.RhospdCloudManager{},
					},
				},
			},
			expectedKind: "cluster",
		},
		"Rhospd Cloud Manager Update Event": {
			inputEvent: &services.Event{
				Request: &services.Event_UpdateRhospdCloudManagerRequest{
					UpdateRhospdCloudManagerRequest: &services.UpdateRhospdCloudManagerRequest{
						RhospdCloudManager: &models.RhospdCloudManager{},
					},
				},
			},
			expectedEvent: &services.Event{
				Request: &services.Event_UpdateRhospdCloudManagerRequest{
					UpdateRhospdCloudManagerRequest: &services.UpdateRhospdCloudManagerRequest{
						RhospdCloudManager: &models.RhospdCloudManager{},
					},
				},
			},
			expectedKind: "cluster",
		},
		"Rhospd Cloud Manager Deletion Event": {
			inputEvent: &services.Event{
				Request: &services.Event_DeleteRhospdCloudManagerRequest{
					DeleteRhospdCloudManagerRequest: &services.DeleteRhospdCloudManagerRequest{},
				},
			},
			expectedEvent: &services.Event{
				Request: &services.Event_DeleteRhospdCloudManagerRequest{
					DeleteRhospdCloudManagerRequest: &services.DeleteRhospdCloudManagerRequest{},
				},
			},
			expectedKind: "cluster",
		},
		"Cloud Creation Event": {
			inputEvent: &services.Event{
				Request: &services.Event_CreateCloudRequest{
					CreateCloudRequest: &services.CreateCloudRequest{
						Cloud: &models.Cloud{},
					},
				},
			},
			expectedEvent: &services.Event{
				Request: &services.Event_CreateCloudRequest{
					CreateCloudRequest: &services.CreateCloudRequest{
						Cloud: &models.Cloud{},
					},
				},
			},
			expectedKind: "cloud",
		},
		"Cloud Update Event": {
			inputEvent: &services.Event{
				Request: &services.Event_UpdateCloudRequest{
					UpdateCloudRequest: &services.UpdateCloudRequest{
						Cloud: &models.Cloud{},
					},
				},
			},
			expectedEvent: &services.Event{
				Request: &services.Event_UpdateCloudRequest{
					UpdateCloudRequest: &services.UpdateCloudRequest{
						Cloud: &models.Cloud{},
					},
				},
			},
			expectedKind: "cloud",
		},
		"Cloud Deletion Event": {
			inputEvent: &services.Event{
				Request: &services.Event_DeleteCloudRequest{
					DeleteCloudRequest: &services.DeleteCloudRequest{},
				},
			},
			expectedEvent: &services.Event{
				Request: &services.Event_DeleteCloudRequest{
					DeleteCloudRequest: &services.DeleteCloudRequest{},
				},
			},
			expectedKind: "cloud",
		},
	}
	for name, processCase := range processCases {
		t.Run(
			name,
			func(t *testing.T) {
				h := &mockHandler{}
				r := &Agent{handler: h, log: logutil.NewLogger("agent")}

				_, err := r.Process(context.Background(), processCase.inputEvent)

				assert.NoError(t, err)
				assert.Equal(t, processCase.expectedEvent, h.event)
				assert.Equal(t, processCase.expectedKind, h.kind)
			},
		)
	}
}

type mockHandler struct {
	event *services.Event
	kind  string
}

func (h *mockHandler) handleCluster(e *services.Event, c *Config) error {
	h.event = e
	h.kind = "cluster"

	return nil
}

func (h *mockHandler) handleCloud(e *services.Event, c *Config) error {
	h.event = e
	h.kind = "cloud"

	return nil
}
