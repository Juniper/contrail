package sync_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/services"
)

func TestCycle(t *testing.T) {
	projectEvent := &services.Event{
		Request: &services.Event_CreateProjectRequest{
			CreateProjectRequest: &services.CreateProjectRequest{
				Project: &models.Project{
					UUID: "Project",
				},
			},
		},
	}

	projectWithRefEvent := &services.Event{
		Request: &services.Event_CreateProjectRequest{
			CreateProjectRequest: &services.CreateProjectRequest{
				Project: &models.Project{
					UUID: "Project",
					FloatingIPPoolRefs: []*models.ProjectFloatingIPPoolRef{
						{UUID: "FloatingIPPool"},
					},
				},
			},
		},
	}

	vnEvent := &services.Event{
		Request: &services.Event_CreateVirtualNetworkRequest{
			CreateVirtualNetworkRequest: &services.CreateVirtualNetworkRequest{
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

	fippEvent := &services.Event{
		Request: &services.Event_CreateFloatingIPPoolRequest{
			CreateFloatingIPPoolRequest: &services.CreateFloatingIPPoolRequest{
				FloatingIPPool: &models.FloatingIPPool{
					UUID:       "FloatingIPPool",
					ParentUUID: "VirtualNetwork",
				},
			},
		},
	}
	var tests = []struct {
		name   string
		events []*services.Event
		cycle  bool
	}{
		{
			name:  "no cycle",
			cycle: false,
			events: []*services.Event{
				projectEvent,
				vnEvent,
				fippEvent,
			},
		},
		{
			name:  "basic cycle",
			cycle: true,
			events: []*services.Event{
				projectWithRefEvent,
				vnEvent,
				fippEvent,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			graph, err := services.NewEventGraph(tt.events)
			assert.NoError(t, err)
			assert.Equal(t, graph.CheckCycle(), tt.cycle)
		})
	}
}

func TestSort(t *testing.T) {
	projectEvent := &services.Event{
		Request: &services.Event_CreateProjectRequest{
			CreateProjectRequest: &services.CreateProjectRequest{
				Project: &models.Project{
					UUID: "Project",
				},
			},
		},
	}

	vnEvent := &services.Event{
		Request: &services.Event_CreateVirtualNetworkRequest{
			CreateVirtualNetworkRequest: &services.CreateVirtualNetworkRequest{
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

	ipamEvent := &services.Event{
		Request: &services.Event_CreateNetworkIpamRequest{
			CreateNetworkIpamRequest: &services.CreateNetworkIpamRequest{
				NetworkIpam: &models.NetworkIpam{
					UUID:       "NetworkIpam",
					ParentUUID: "Project",
				},
			},
		},
	}

	fippEvent := &services.Event{
		Request: &services.Event_CreateFloatingIPPoolRequest{
			CreateFloatingIPPoolRequest: &services.CreateFloatingIPPoolRequest{
				FloatingIPPool: &models.FloatingIPPool{
					UUID:       "FloatingIPPool",
					ParentUUID: "VirtualNetwork",
				},
			},
		},
	}

	var tests = []struct {
		name     string
		events   []*services.Event
		expected services.EventList
	}{
		{
			name: "sort basic",
			events: []*services.Event{
				projectEvent,
				vnEvent,
				ipamEvent,
				fippEvent,
			},
			expected: services.EventList{
				Events: []*services.Event{
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
			graph, err := services.NewEventGraph(tt.events)
			assert.NoError(t, err)
			sorted := graph.SortEvents()
			assert.Equal(t, tt.expected, sorted)
		})
	}
}

func TestOperationKind(t *testing.T) {
	createEvent, err := services.NewEvent(services.EventOption{
		Operation: services.OperationCreate,
		Kind:      models.KindProject,
	})
	assert.NoError(t, err)

	updateEvent, err := services.NewEvent(services.EventOption{
		Operation: services.OperationUpdate,
		Kind:      models.KindProject,
	})
	assert.NoError(t, err)

	deleteEvent, err := services.NewEvent(services.EventOption{
		Operation: services.OperationDelete,
		Kind:      models.KindProject,
	})
	assert.NoError(t, err)

	var tests = []struct {
		name     string
		expected string
		events   *services.EventList
	}{
		{
			name:     "Multiple create events",
			expected: services.OperationCreate,
			events: &services.EventList{
				Events: []*services.Event{
					createEvent,
					createEvent,
					createEvent,
				},
			},
		},
		{
			name:     "Multiple update events",
			expected: services.OperationUpdate,
			events: &services.EventList{
				Events: []*services.Event{
					updateEvent,
					updateEvent,
					updateEvent,
				},
			},
		},
		{
			name:     "Multiple delete events",
			expected: services.OperationDelete,
			events: &services.EventList{
				Events: []*services.Event{
					deleteEvent,
					deleteEvent,
					deleteEvent,
				},
			},
		},
		{
			name:     "Mixed events",
			expected: "MIXED",
			events: &services.EventList{
				Events: []*services.Event{
					createEvent,
					updateEvent,
					deleteEvent,
				},
			},
		},
		{
			name:     "Empty event list",
			expected: "EMPTY",
			events: &services.EventList{
				Events: []*services.Event{},
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
		events      []*services.Event
		requireSort bool
		include     bool
	}{
		{
			name:   "No events",
			events: []*services.Event{},
		},
		{
			name: "One event create without references",
			events: []*services.Event{
				{
					Request: &services.Event_CreateProjectRequest{
						CreateProjectRequest: &services.CreateProjectRequest{
							Project: projectWithNoRefs,
						},
					},
				},
			},
		},
		{
			name: "One event create with already existing parent",
			events: []*services.Event{
				{
					Request: &services.Event_CreateProjectRequest{
						CreateProjectRequest: &services.CreateProjectRequest{
							Project: projectWithParent,
						},
					},
				},
			},
		},
		{
			name: "One event create with already existing reference",
			events: []*services.Event{
				{
					Request: &services.Event_CreateProjectRequest{
						CreateProjectRequest: &services.CreateProjectRequest{
							Project: projectWithExistingRef,
						},
					},
				},
			},
		},
		{
			name: "Two independent create events",
			events: []*services.Event{
				{
					Request: &services.Event_CreateProjectRequest{
						CreateProjectRequest: &services.CreateProjectRequest{
							Project: projectWithExistingRef,
						},
					},
				},
				{
					Request: &services.Event_CreateVirtualNetworkRequest{
						CreateVirtualNetworkRequest: &services.CreateVirtualNetworkRequest{
							VirtualNetwork: virtualNetworkWithExistingParent,
						},
					},
				},
			},
		},
		{
			name: "Two parent-child dependent create events in right order",
			events: []*services.Event{
				{
					Request: &services.Event_CreateProjectRequest{
						CreateProjectRequest: &services.CreateProjectRequest{
							Project: projectWithExistingRef,
						},
					},
				},
				{
					Request: &services.Event_CreateVirtualNetworkRequest{
						CreateVirtualNetworkRequest: &services.CreateVirtualNetworkRequest{
							VirtualNetwork: virtualNetwork,
						},
					},
				},
			},
		},
		{
			name: "Two parent-child dependent create events in wrong order",
			events: []*services.Event{
				{
					Request: &services.Event_CreateVirtualNetworkRequest{
						CreateVirtualNetworkRequest: &services.CreateVirtualNetworkRequest{
							VirtualNetwork: virtualNetwork,
						},
					},
				},
				{
					Request: &services.Event_CreateProjectRequest{
						CreateProjectRequest: &services.CreateProjectRequest{
							Project: projectWithExistingRef,
						},
					},
				},
			},
			requireSort: true,
		},
		{
			name: "Three reference dependent (with refs only to themselves) create events in right order",
			events: []*services.Event{
				{
					Request: &services.Event_CreateProjectRequest{
						CreateProjectRequest: &services.CreateProjectRequest{
							Project: projectWithNoRefs,
						},
					},
				},
				{
					Request: &services.Event_CreateNetworkIpamRequest{
						CreateNetworkIpamRequest: &services.CreateNetworkIpamRequest{
							NetworkIpam: networkIpam,
						},
					},
				},
				{
					Request: &services.Event_CreateVirtualNetworkRequest{
						CreateVirtualNetworkRequest: &services.CreateVirtualNetworkRequest{
							VirtualNetwork: virtualNetworkWithReference,
						},
					},
				},
			},
		},
		{
			name: "Three reference dependent (with refs only to themselves) create events in wrong order",
			events: []*services.Event{
				{
					Request: &services.Event_CreateVirtualNetworkRequest{
						CreateVirtualNetworkRequest: &services.CreateVirtualNetworkRequest{
							VirtualNetwork: virtualNetworkWithReference,
						},
					},
				},
				{
					Request: &services.Event_CreateProjectRequest{
						CreateProjectRequest: &services.CreateProjectRequest{
							Project: projectWithNoRefs,
						},
					},
				},
				{
					Request: &services.Event_CreateNetworkIpamRequest{
						CreateNetworkIpamRequest: &services.CreateNetworkIpamRequest{
							NetworkIpam: networkIpam,
						},
					},
				},
			},
			requireSort: true,
		},
		{
			name: "Three reference dependent (mixed) create events in right order",
			events: []*services.Event{
				{
					Request: &services.Event_CreateProjectRequest{
						CreateProjectRequest: &services.CreateProjectRequest{
							Project: projectWithParent,
						},
					},
				},
				{
					Request: &services.Event_CreateNetworkIpamRequest{
						CreateNetworkIpamRequest: &services.CreateNetworkIpamRequest{
							NetworkIpam: networkIpam,
						},
					},
				},
				{
					Request: &services.Event_CreateVirtualNetworkRequest{
						CreateVirtualNetworkRequest: &services.CreateVirtualNetworkRequest{
							VirtualNetwork: virtualNetworkWithReference,
						},
					},
				},
			},
		},
		{
			name: "Three reference dependent (mixed) create events in wrong order",
			events: []*services.Event{
				{
					Request: &services.Event_CreateVirtualNetworkRequest{
						CreateVirtualNetworkRequest: &services.CreateVirtualNetworkRequest{
							VirtualNetwork: virtualNetworkWithReference,
						},
					},
				},
				{
					Request: &services.Event_CreateProjectRequest{
						CreateProjectRequest: &services.CreateProjectRequest{
							Project: projectWithParent,
						},
					},
				},
				{
					Request: &services.Event_CreateNetworkIpamRequest{
						CreateNetworkIpamRequest: &services.CreateNetworkIpamRequest{
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
			list := services.EventList{
				Events: tt.events,
			}
			assert.Equal(t, tt.requireSort, list.IsSortRequired())
		})
	}
}

func TestSortCreateNoCycle(t *testing.T) {
	projectEvent := &services.Event{
		Request: &services.Event_CreateProjectRequest{
			CreateProjectRequest: &services.CreateProjectRequest{
				Project: &models.Project{
					UUID: "Project",
				},
			},
		},
	}

	projectWithRefEvent := &services.Event{
		Request: &services.Event_CreateProjectRequest{
			CreateProjectRequest: &services.CreateProjectRequest{
				Project: &models.Project{
					UUID: "Project",
					FloatingIPPoolRefs: []*models.ProjectFloatingIPPoolRef{
						{UUID: "FloatingIPPool"},
					},
				},
			},
		},
	}

	vnEvent := &services.Event{
		Request: &services.Event_CreateVirtualNetworkRequest{
			CreateVirtualNetworkRequest: &services.CreateVirtualNetworkRequest{
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

	ipamEvent := &services.Event{
		Request: &services.Event_CreateNetworkIpamRequest{
			CreateNetworkIpamRequest: &services.CreateNetworkIpamRequest{
				NetworkIpam: &models.NetworkIpam{
					UUID:       "NetworkIpam",
					ParentUUID: "Project",
				},
			},
		},
	}

	fippEvent := &services.Event{
		Request: &services.Event_CreateFloatingIPPoolRequest{
			CreateFloatingIPPoolRequest: &services.CreateFloatingIPPoolRequest{
				FloatingIPPool: &models.FloatingIPPool{
					UUID:       "FloatingIPPool",
					ParentUUID: "VirtualNetwork",
				},
			},
		},
	}
	var tests = []struct {
		name     string
		events   []*services.Event
		expected services.EventList
		fails    bool
	}{
		{
			name: "happy scenario",
			events: []*services.Event{
				projectEvent,
				vnEvent,
				ipamEvent,
				fippEvent,
			},
			expected: services.EventList{
				Events: []*services.Event{
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
			events: []*services.Event{
				projectWithRefEvent,
				vnEvent,
				ipamEvent,
				fippEvent,
			},
			expected: services.EventList{
				Events: []*services.Event{
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
			list := services.EventList{
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
