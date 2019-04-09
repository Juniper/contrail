package services

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/Juniper/contrail/pkg/models"
)

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
