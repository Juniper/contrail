package agent

import (
	"context"
	"testing"

	"github.com/Juniper/asf/pkg/logutil"
	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/services"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
)

func TestProcess(t *testing.T) {
	processCases := map[string]struct {
		input *services.Event
		fail  bool
	}{
		"Unexpected Event": {
			input: &services.Event{
				Request: &services.Event_CreateAccessControlListRequest{},
			},
		},
		"Contrail Creation Event": {
			input: &services.Event{
				Request: &services.Event_CreateContrailClusterRequest{
					CreateContrailClusterRequest: &services.CreateContrailClusterRequest{
						ContrailCluster: &models.ContrailCluster{},
					},
				},
			},
		},
		"Contrail Update Event": {
			input: &services.Event{
				Request: &services.Event_UpdateContrailClusterRequest{
					UpdateContrailClusterRequest: &services.UpdateContrailClusterRequest{
						ContrailCluster: &models.ContrailCluster{},
					},
				},
			},
		},
		"Contrail Deletion Event": {
			input: &services.Event{
				Request: &services.Event_DeleteContrailClusterRequest{
					DeleteContrailClusterRequest: &services.DeleteContrailClusterRequest{},
				},
			},
		},
		"Rhospd Cloud Manager Creation Event": {
			input: &services.Event{
				Request: &services.Event_CreateRhospdCloudManagerRequest{
					CreateRhospdCloudManagerRequest: &services.CreateRhospdCloudManagerRequest{
						RhospdCloudManager: &models.RhospdCloudManager{},
					},
				},
			},
		},
		"Rhospd Cloud Manager Update Event": {
			input: &services.Event{
				Request: &services.Event_UpdateRhospdCloudManagerRequest{
					UpdateRhospdCloudManagerRequest: &services.UpdateRhospdCloudManagerRequest{
						RhospdCloudManager: &models.RhospdCloudManager{},
					},
				},
			},
		},
		"Rhospd Cloud Manager Deletion Event": {
			input: &services.Event{
				Request: &services.Event_DeleteRhospdCloudManagerRequest{
					DeleteRhospdCloudManagerRequest: &services.DeleteRhospdCloudManagerRequest{},
				},
			},
		},
		"Cloud Creation Event": {
			input: &services.Event{
				Request: &services.Event_CreateCloudRequest{
					CreateCloudRequest: &services.CreateCloudRequest{
						Cloud: &models.Cloud{},
					},
				},
			},
		},
		"Cloud Update Event": {
			input: &services.Event{
				Request: &services.Event_UpdateCloudRequest{
					UpdateCloudRequest: &services.UpdateCloudRequest{
						Cloud: &models.Cloud{},
					},
				},
			},
		},
		"Cloud Deletion Event": {
			input: &services.Event{
				Request: &services.Event_DeleteCloudRequest{
					DeleteCloudRequest: &services.DeleteCloudRequest{},
				},
			},
		},
	}
	for name, processCase := range processCases {
		t.Run(
			name,
			func(t *testing.T) {
				h := &handlerMock{}
				r := &Agent{handler: h, log: logutil.NewLogger("agent")}

				_, err := r.Process(context.Background(), processCase.input)

				assert.NoError(t, err)
			},
		)
	}
}

type handlerMock struct{}

func (h *handlerMock) processCluster(e *services.Event, c *Config) error {
	if e.Kind() == "contrail-cluster" || e.Kind() == "rhospd-cloud-manager" {
		return nil
	}

	return errors.New("event type unmatched for cluster event")
}
func (h *handlerMock) processCloud(e *services.Event, c *Config) error {
	if e.Kind() == "cloud" {
		return nil
	}

	return errors.New("event type unmatched for cloud event")
}
