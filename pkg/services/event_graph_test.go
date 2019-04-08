package services

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/Juniper/contrail/pkg/models"
)

func Test_eventsToEventNodes(t *testing.T) {
	project := &models.Project{
		UUID: "Project",
	}
	virtualNetwork := &models.VirtualNetwork{
		UUID:       "VirtualNetwork",
		ParentUUID: "Project",
		NetworkIpamRefs: []*models.VirtualNetworkNetworkIpamRef{
			{
				UUID: "NetworkIpam",
			},
		},
	}
	networkIpam := &models.NetworkIpam{
		UUID:       "NetworkIpam",
		ParentUUID: "Project",
	}

	events := []*Event{
		{
			Request: &Event_CreateVirtualNetworkRequest{
				CreateVirtualNetworkRequest: &CreateVirtualNetworkRequest{
					VirtualNetwork: virtualNetwork,
				},
			},
		},
		{
			Request: &Event_CreateNetworkIpamRequest{
				CreateNetworkIpamRequest: &CreateNetworkIpamRequest{
					NetworkIpam: networkIpam,
				},
			},
		},
		{
			Request: &Event_CreateProjectRequest{
				CreateProjectRequest: &CreateProjectRequest{
					Project: project,
				},
			},
		},
	}

	eventNodes, err := eventsToEventNodes(events)

	assert.NoError(t, err)

	assert.Equal(t, events[0], eventNodes[0].event)
	assert.Equal(t, events[1], eventNodes[1].event)
	assert.Equal(t, events[2], eventNodes[2].event)

	assert.True(t, containsRefsTo(eventNodes[0], []*eventNode{eventNodes[1], eventNodes[2]}))
	assert.True(t, containsRefsTo(eventNodes[1], []*eventNode{eventNodes[2]}))
	assert.False(t, containsRefsTo(eventNodes[1], []*eventNode{eventNodes[0]}))
	assert.False(t, containsRefsTo(eventNodes[2], []*eventNode{eventNodes[0], eventNodes[1]}))
}

func containsRefsTo(original *eventNode, referencedNodes []*eventNode) bool {
	dependencies := make(map[*eventNode]bool)
	for _, nodeDependency := range original.referencesAndParent {
		dependencies[nodeDependency] = true
	}

	for _, ref := range referencedNodes {
		if !dependencies[ref] {
			return false
		}
	}
	return true
}
