package services

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/models/basemodels"
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
			refMap := make(map[*Event]basemodels.References)
			var err error
			for _, ev := range tt.events {
				refMap[ev], err = ev.getReferences()
				assert.NoError(t, err)
			}
			graph := NewEventGraph(tt.events, refMap)
			assert.Equal(t, tt.isCycle, graph.CheckCycle())
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
			refMap := make(map[*Event]basemodels.References)
			var err error
			for _, ev := range tt.events {
				refMap[ev], err = ev.getReferences()
				assert.NoError(t, err)
			}
			graph := NewEventGraph(tt.events, refMap)
			assert.Equal(t, tt.expected, graph.SortEvents())
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

func TestIsSortRequired(t *testing.T) {
	projectWithNoRefs := &models.Project{
		UUID: "Project",
	}
	projectWithParent := &models.Project{
		UUID:       "Project",
		ParentUUID: "beefbeef-beef-beef-beef-beefbeef0002",
	}
	projectWithExistingRef := &models.Project{
		UUID: "Project",
		ApplicationPolicySetRefs: []*models.ProjectApplicationPolicySetRef{
			{
				UUID: "8a05c096-09ed-4c4b-a763-cf1d5ba92a27",
			},
		},
	}
	virtualNetworkWithExistingParent := &models.VirtualNetwork{
		UUID:       "VirtualNetwork",
		ParentUUID: "beefbeef-beef-beef-beef-beefbeef0003",
	}
	virtualNetwork := &models.VirtualNetwork{
		UUID:       "VirtualNetwork",
		ParentUUID: "Project",
	}
	virtualNetworkWithReference := &models.VirtualNetwork{
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

	tests := []struct {
		name        string
		events      []*Event
		requireSort bool
		include     bool
	}{
		{
			name:   "No events",
			events: []*Event{},
		},
		{
			name: "One event create without references",
			events: []*Event{
				{
					Request: &Event_CreateProjectRequest{
						CreateProjectRequest: &CreateProjectRequest{
							Project: projectWithNoRefs,
						},
					},
				},
			},
		},
		{
			name: "One event create with already existing parent",
			events: []*Event{
				{
					Request: &Event_CreateProjectRequest{
						CreateProjectRequest: &CreateProjectRequest{
							Project: projectWithParent,
						},
					},
				},
			},
		},
		{
			name: "One event create with already existing reference",
			events: []*Event{
				{
					Request: &Event_CreateProjectRequest{
						CreateProjectRequest: &CreateProjectRequest{
							Project: projectWithExistingRef,
						},
					},
				},
			},
		},
		{
			name: "Two independent create events",
			events: []*Event{
				{
					Request: &Event_CreateProjectRequest{
						CreateProjectRequest: &CreateProjectRequest{
							Project: projectWithExistingRef,
						},
					},
				},
				{
					Request: &Event_CreateVirtualNetworkRequest{
						CreateVirtualNetworkRequest: &CreateVirtualNetworkRequest{
							VirtualNetwork: virtualNetworkWithExistingParent,
						},
					},
				},
			},
		},
		{
			name: "Two parent-child dependent create events in right order",
			events: []*Event{
				{
					Request: &Event_CreateProjectRequest{
						CreateProjectRequest: &CreateProjectRequest{
							Project: projectWithExistingRef,
						},
					},
				},
				{
					Request: &Event_CreateVirtualNetworkRequest{
						CreateVirtualNetworkRequest: &CreateVirtualNetworkRequest{
							VirtualNetwork: virtualNetwork,
						},
					},
				},
			},
		},
		{
			name: "Two parent-child dependent create events in wrong order",
			events: []*Event{
				{
					Request: &Event_CreateVirtualNetworkRequest{
						CreateVirtualNetworkRequest: &CreateVirtualNetworkRequest{
							VirtualNetwork: virtualNetwork,
						},
					},
				},
				{
					Request: &Event_CreateProjectRequest{
						CreateProjectRequest: &CreateProjectRequest{
							Project: projectWithExistingRef,
						},
					},
				},
			},
			requireSort: true,
		},
		{
			name: "Three reference dependent (with refs only to themselves) create events in right order",
			events: []*Event{
				{
					Request: &Event_CreateProjectRequest{
						CreateProjectRequest: &CreateProjectRequest{
							Project: projectWithNoRefs,
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
					Request: &Event_CreateVirtualNetworkRequest{
						CreateVirtualNetworkRequest: &CreateVirtualNetworkRequest{
							VirtualNetwork: virtualNetworkWithReference,
						},
					},
				},
			},
		},
		{
			name: "Three reference dependent (with refs only to themselves) create events in wrong order",
			events: []*Event{
				{
					Request: &Event_CreateVirtualNetworkRequest{
						CreateVirtualNetworkRequest: &CreateVirtualNetworkRequest{
							VirtualNetwork: virtualNetworkWithReference,
						},
					},
				},
				{
					Request: &Event_CreateProjectRequest{
						CreateProjectRequest: &CreateProjectRequest{
							Project: projectWithNoRefs,
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
			},
			requireSort: true,
		},
		{
			name: "Three reference dependent (mixed) create events in right order",
			events: []*Event{
				{
					Request: &Event_CreateProjectRequest{
						CreateProjectRequest: &CreateProjectRequest{
							Project: projectWithParent,
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
					Request: &Event_CreateVirtualNetworkRequest{
						CreateVirtualNetworkRequest: &CreateVirtualNetworkRequest{
							VirtualNetwork: virtualNetworkWithReference,
						},
					},
				},
			},
		},
		{
			name: "Three reference dependent (mixed) create events in wrong order",
			events: []*Event{
				{
					Request: &Event_CreateVirtualNetworkRequest{
						CreateVirtualNetworkRequest: &CreateVirtualNetworkRequest{
							VirtualNetwork: virtualNetworkWithReference,
						},
					},
				},
				{
					Request: &Event_CreateProjectRequest{
						CreateProjectRequest: &CreateProjectRequest{
							Project: projectWithParent,
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
			},
			requireSort: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			list := EventList{
				Events: tt.events,
			}
			assert.Equal(t, tt.requireSort, list.isSortRequired())
		})
	}
}

func TestSortCreateNoCycle(t *testing.T) {
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
		fails    bool
	}{
		{
			name: "happy scenario",
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
		{
			name:  "ref cycle",
			fails: true,
			events: []*Event{
				projectWithRefEvent,
				vnEvent,
				ipamEvent,
				fippEvent,
			},
			expected: EventList{
				Events: []*Event{
					projectWithRefEvent,
					ipamEvent,
					vnEvent,
					fippEvent,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			list := EventList{
				Events: tt.events,
			}
			err := list.SortCreateNoCycle()
			if tt.fails {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expected, list)
			}
		})
	}
}
