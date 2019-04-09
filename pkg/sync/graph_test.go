package sync_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/services"
)

func TestCycle(t *testing.T) {
	var tests = []struct {
		name   string
		events []*services.Event
		cycle  bool
	}{
		{
			name:  "no cycle",
			cycle: false,
			events: []*services.Event{
				{
					Request: &services.Event_CreateProjectRequest{
						CreateProjectRequest: &services.CreateProjectRequest{
							Project: &models.Project{
								UUID: "Project",
							},
						},
					},
				},
				{
					Request: &services.Event_CreateVirtualNetworkRequest{
						CreateVirtualNetworkRequest: &services.CreateVirtualNetworkRequest{
							VirtualNetwork: &models.VirtualNetwork{
								UUID:       "VirtualNetwork",
								ParentUUID: "Project",
							},
						},
					},
				},
				{
					Request: &services.Event_CreateFloatingIPPoolRequest{
						CreateFloatingIPPoolRequest: &services.CreateFloatingIPPoolRequest{
							FloatingIPPool: &models.FloatingIPPool{
								UUID:       "FloatingIPPool",
								ParentUUID: "VirtualNetwork",
							},
						},
					},
				},
			},
		},
		{
			name:  "basic cycle",
			cycle: true,
			events: []*services.Event{
				{
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
				},
				{
					Request: &services.Event_CreateVirtualNetworkRequest{
						CreateVirtualNetworkRequest: &services.CreateVirtualNetworkRequest{
							VirtualNetwork: &models.VirtualNetwork{
								UUID:       "VirtualNetwork",
								ParentUUID: "Project",
							},
						},
					},
				},
				{
					Request: &services.Event_CreateFloatingIPPoolRequest{
						CreateFloatingIPPoolRequest: &services.CreateFloatingIPPoolRequest{
							FloatingIPPool: &models.FloatingIPPool{
								UUID:       "FloatingIPPool",
								ParentUUID: "VirtualNetwork",
							},
						},
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nodes, err := services.EventsToEventNodes(tt.events)
			assert.NoError(t, err)
			graph := services.NewEventGraph(nodes)
			assert.Equal(t, graph.CheckCycle(), tt.cycle)
		})
	}
}

func TestSort(t *testing.T) {
	var tests = []struct {
		name   string
		events []*services.Event
	}{
		{
			name: "sort basic",
			events: []*services.Event{
				{
					Request: &services.Event_CreateProjectRequest{
						CreateProjectRequest: &services.CreateProjectRequest{
							Project: &models.Project{
								UUID: "Project",
							},
						},
					},
				},
				{
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
				},
				{
					Request: &services.Event_CreateNetworkIpamRequest{
						CreateNetworkIpamRequest: &services.CreateNetworkIpamRequest{
							NetworkIpam: &models.NetworkIpam{
								UUID:       "NetworkIpam",
								ParentUUID: "Project",
							},
						},
					},
				},
				{
					Request: &services.Event_CreateFloatingIPPoolRequest{
						CreateFloatingIPPoolRequest: &services.CreateFloatingIPPoolRequest{
							FloatingIPPool: &models.FloatingIPPool{
								UUID:       "FloatingIPPool",
								ParentUUID: "VirtualNetwork",
							},
						},
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nodes, err := services.EventsToEventNodes(tt.events)
			assert.NoError(t, err)
			graph := services.NewEventGraph(nodes)
			sorted := graph.SortEvents()
			fmt.Println("Sorted events")
			for _, s := range sorted.Events {
				fmt.Println(s.GetUUID())
			}
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
		result := tt.events.CheckOperationType()
		assert.NoError(t, err)
		assert.Equal(t, result, tt.expected)
	}
}
