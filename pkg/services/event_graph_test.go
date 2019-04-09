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

func TestCycle(t *testing.T) {
	projectEvent := &Event{
		Request: &Event_CreateProjectRequest{
			CreateProjectRequest: &CreateProjectRequest{
				Project: &models.Project{
					UUID: "Project",
				},
			},
		},
	}

	projectWithRefEvent := &Event{
		Request: &Event_CreateProjectRequest{
			CreateProjectRequest: &CreateProjectRequest{
				Project: &models.Project{
					UUID: "Project",
					FloatingIPPoolRefs: []*models.ProjectFloatingIPPoolRef{
						{UUID: "FloatingIPPool"},
					},
				},
			},
		},
	}

	vnEvent := &Event{
		Request: &Event_CreateVirtualNetworkRequest{
			CreateVirtualNetworkRequest: &CreateVirtualNetworkRequest{
				VirtualNetwork: &models.VirtualNetwork{
					UUID:       "VirtualNetwork",
					ParentUUID: "Project",
					NetworkIpamRefs: []*models.VirtualNetworkNetworkIpamRef{
						{UUID: "NetworkIpam"},
					},
				},
			},
		},
	}

	fippEvent := &Event{
		Request: &Event_CreateFloatingIPPoolRequest{
			CreateFloatingIPPoolRequest: &CreateFloatingIPPoolRequest{
				FloatingIPPool: &models.FloatingIPPool{
					UUID:       "FloatingIPPool",
					ParentUUID: "VirtualNetwork",
				},
			},
		},
	}
	var tests = []struct {
		name    string
		events  []*Event
		isCycle bool
	}{
		{
			name: "no cycle",
			events: []*Event{
				projectEvent,
				vnEvent,
				fippEvent,
			},
		},
		{
			name:    "basic cycle",
			isCycle: true,
			events: []*Event{
				projectWithRefEvent,
				vnEvent,
				fippEvent,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			graph, err := newEventGraph(tt.events)
			assert.NoError(t, err)
			assert.Equal(t, tt.isCycle, graph.checkCycle())
		})
	}
}

func TestSort(t *testing.T) {
	projectEvent := &Event{
		Request: &Event_CreateProjectRequest{
			CreateProjectRequest: &CreateProjectRequest{
				Project: &models.Project{
					UUID: "Project",
				},
			},
		},
	}

	vnEvent := &Event{
		Request: &Event_CreateVirtualNetworkRequest{
			CreateVirtualNetworkRequest: &CreateVirtualNetworkRequest{
				VirtualNetwork: &models.VirtualNetwork{
					UUID:       "VirtualNetwork",
					ParentUUID: "Project",
					NetworkIpamRefs: []*models.VirtualNetworkNetworkIpamRef{
						{UUID: "NetworkIpam"},
					},
				},
			},
		},
	}

	ipamEvent := &Event{
		Request: &Event_CreateNetworkIpamRequest{
			CreateNetworkIpamRequest: &CreateNetworkIpamRequest{
				NetworkIpam: &models.NetworkIpam{
					UUID:       "NetworkIpam",
					ParentUUID: "Project",
				},
			},
		},
	}

	fippEvent := &Event{
		Request: &Event_CreateFloatingIPPoolRequest{
			CreateFloatingIPPoolRequest: &CreateFloatingIPPoolRequest{
				FloatingIPPool: &models.FloatingIPPool{
					UUID:       "FloatingIPPool",
					ParentUUID: "VirtualNetwork",
				},
			},
		},
	}

	var tests = []struct {
		name     string
		events   []*Event
		expected EventList
	}{
		{
			name: "sort basic",
			events: []*Event{
				projectEvent,
				vnEvent,
				ipamEvent,
				fippEvent,
			},
			expected: EventList{
				Events: []*Event{
					projectEvent,
					ipamEvent,
					vnEvent,
					fippEvent,
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			graph, err := newEventGraph(tt.events)
			assert.NoError(t, err)
			assert.Equal(t, tt.expected, graph.sortEvents())
		})
	}
}

func TestOperationKind(t *testing.T) {
	createEvent, err := NewEvent(EventOption{
		Operation: OperationCreate,
		Kind:      models.KindProject,
	})
	assert.NoError(t, err)

	updateEvent, err := NewEvent(EventOption{
		Operation: OperationUpdate,
		Kind:      models.KindProject,
	})
	assert.NoError(t, err)

	deleteEvent, err := NewEvent(EventOption{
		Operation: OperationDelete,
		Kind:      models.KindProject,
	})
	assert.NoError(t, err)

	var tests = []struct {
		name     string
		expected string
		events   *EventList
	}{
		{
			name:     "Multiple create events",
			expected: OperationCreate,
			events: &EventList{
				Events: []*Event{
					createEvent,
					createEvent,
					createEvent,
				},
			},
		},
		{
			name:     "Multiple update events",
			expected: OperationUpdate,
			events: &EventList{
				Events: []*Event{
					updateEvent,
					updateEvent,
					updateEvent,
				},
			},
		},
		{
			name:     "Multiple delete events",
			expected: OperationDelete,
			events: &EventList{
				Events: []*Event{
					deleteEvent,
					deleteEvent,
					deleteEvent,
				},
			},
		},
		{
			name:     "Mixed events",
			expected: "MIXED",
			events: &EventList{
				Events: []*Event{
					createEvent,
					updateEvent,
					deleteEvent,
				},
			},
		},
		{
			name:     "Empty event list",
			expected: "EMPTY",
			events: &EventList{
				Events: []*Event{},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.events.CheckOperationType(), tt.expected)
		})
	}
}
