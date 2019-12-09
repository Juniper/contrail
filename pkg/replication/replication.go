package replication

import (
	"context"
	"sync"

	"github.com/Juniper/asf/pkg/logutil"
	"github.com/Juniper/contrail/pkg/endpoint"
	"github.com/Juniper/contrail/pkg/keystone"
	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/services"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	syncp "github.com/Juniper/contrail/pkg/sync"
)

const (
	createTagURL            = "tags"
	createVirtualNetworkURL = "virtual-networks"
	createNetworkIpamURL    = "network-ipams"
	createPortURL           = "ports"
	createNodeProfileURL    = "node-profiles"
	createNodeURL           = "nodes"
	createHardwareURL       = "hardwares"
	createCardURL           = "cards"
	updateTagURL            = "tag/"
	updateVirtualNetworkURL = "virtual-network/"
	updateNetworkIpamURL    = "network-ipam/"
	updatePortURL           = "port/"
	updateNodeProfileURL    = "node-profile/"
	updateNodeURL           = "node/"
	updateHardwareURL       = "hardware/"
	updateCardURL           = "card/"

	deleteAction    = "delete"
	createAction    = "create"
	updateAction    = "update"
	refUpdateAction = "ref-update"
)

type handler interface {
	Replicate(action, url string, data interface{}, response interface{})
	CreateClient(e *models.Endpoint)
	UpdateClient(e *models.Endpoint)
	DeleteClient(id string)
}

// Replicator is an implementation to replicate objects to python API
type Replicator struct {
	serviceWaitGroup   *sync.WaitGroup
	stopServiceContext context.CancelFunc
	handler            handler
	log                *logrus.Entry
}

// New initializes replication data
func New(epStore *sync.Map, localKeystone *keystone.Keystone) (*Replicator, error) {

	if err := logutil.Configure(viper.GetString("log_level")); err != nil {
		return nil, err
	}

	return &Replicator{
		serviceWaitGroup: &sync.WaitGroup{},
		handler:          newVncAPIHandle(&endpoint.Store{epStore}, localKeystone),
		log:              logutil.NewLogger("vnc_replication"),
	}, nil
}

// Start replication service
func (r *Replicator) Start() error {
	processor := &services.EventListProcessor{
		EventProcessor:    r,
		InTransactionDoer: services.NoTransaction,
	}
	producer, err := syncp.NewEventProducer("replicator-watcher", processor)
	if err != nil {
		return err
	}

	var ctx context.Context
	ctx, r.stopServiceContext = context.WithCancel(context.Background())

	r.serviceWaitGroup.Add(1)
	go func() {
		defer r.serviceWaitGroup.Done()
		defer producer.Close()

		err = producer.Start(ctx)
	}()
	<-producer.Watcher.DumpDone()

	return err
}

// Process processes event by sending requests to all registered clusters.
func (r *Replicator) Process(ctx context.Context, e *services.Event) (*services.Event, error) { //nolint: gocyclo
	r.log.Infof("Received event: %v", e)
	if e == nil {
		return nil, nil
	}

	if re, ok := e.Unwrap().(services.ReferenceEvent); ok {
		refUpdate := services.NewRefUpdateFromEvent(re)
		r.handler.Replicate(
			refUpdateAction,
			services.RefUpdatePath,
			refUpdate,
			map[string]interface{}{},
		)
		return e, nil
	}

	switch event := e.Request.(type) {
	// watch endpoint event and prepare clients
	case *services.Event_CreateEndpointRequest:
		ep := event.CreateEndpointRequest.Endpoint
		r.handler.CreateClient(ep)
	case *services.Event_UpdateEndpointRequest:
		ep := event.UpdateEndpointRequest.Endpoint
		r.handler.UpdateClient(ep)
	case *services.Event_DeleteEndpointRequest:
		id := event.DeleteEndpointRequest.ID
		r.handler.DeleteClient(id)
		// handle tag
	case *services.Event_CreateTagRequest:
		event.CreateTagRequest.Tag.TagID = ""
		r.handler.Replicate(createAction, createTagURL,
			event.CreateTagRequest, &services.CreateTagResponse{})
	case *services.Event_UpdateTagRequest:
		event.UpdateTagRequest.Tag.TagID = ""
		objID := event.UpdateTagRequest.Tag.UUID
		r.handler.Replicate(updateAction, updateTagURL+objID,
			event.UpdateTagRequest, &services.UpdateTagResponse{})
	case *services.Event_DeleteTagRequest:
		objID := event.DeleteTagRequest.ID
		r.handler.Replicate(deleteAction, updateTagURL+objID,
			event.DeleteTagRequest, &services.DeleteTagResponse{})
	// handle virtual-network
	case *services.Event_CreateVirtualNetworkRequest:
		event.CreateVirtualNetworkRequest.VirtualNetwork.VirtualNetworkNetworkID = 0
		r.handler.Replicate(createAction, createVirtualNetworkURL,
			event.CreateVirtualNetworkRequest,
			&services.CreateVirtualNetworkResponse{})
	case *services.Event_UpdateVirtualNetworkRequest:
		event.UpdateVirtualNetworkRequest.VirtualNetwork.VirtualNetworkNetworkID = 0
		objID := event.UpdateVirtualNetworkRequest.VirtualNetwork.UUID
		r.handler.Replicate(updateAction, updateVirtualNetworkURL+objID,
			event.UpdateVirtualNetworkRequest,
			&services.UpdateVirtualNetworkResponse{})
	case *services.Event_DeleteVirtualNetworkRequest:
		objID := event.DeleteVirtualNetworkRequest.ID
		r.handler.Replicate(deleteAction, updateVirtualNetworkURL+objID,
			event.DeleteVirtualNetworkRequest, &services.DeleteVirtualNetworkResponse{})
		// handle network-ipam
	case *services.Event_CreateNetworkIpamRequest:
		r.handler.Replicate(createAction, createNetworkIpamURL,
			event.CreateNetworkIpamRequest, &services.CreateNetworkIpamResponse{})
	case *services.Event_UpdateNetworkIpamRequest:
		objID := event.UpdateNetworkIpamRequest.NetworkIpam.UUID
		r.handler.Replicate(updateAction, updateNetworkIpamURL+objID,
			event.UpdateNetworkIpamRequest, &services.UpdateNetworkIpamResponse{})
	case *services.Event_DeleteNetworkIpamRequest:
		objID := event.DeleteNetworkIpamRequest.ID
		r.handler.Replicate(deleteAction, updateNetworkIpamURL+objID,
			event.DeleteNetworkIpamRequest, &services.DeleteNetworkIpamResponse{})
	// handle node-profile
	case *services.Event_CreateNodeProfileRequest:
		r.handler.Replicate(createAction, createNodeProfileURL,
			event.CreateNodeProfileRequest, &services.CreateNodeProfileResponse{})
	case *services.Event_UpdateNodeProfileRequest:
		objID := event.UpdateNodeProfileRequest.NodeProfile.UUID
		r.handler.Replicate(updateAction, updateNodeProfileURL+objID,
			event.UpdateNodeProfileRequest, &services.UpdateNodeProfileResponse{})
	case *services.Event_DeleteNodeProfileRequest:
		objID := event.DeleteNodeProfileRequest.ID
		r.handler.Replicate(deleteAction, updateNodeProfileURL+objID,
			event.DeleteNodeProfileRequest, &services.DeleteNodeProfileResponse{})
	// handle nodes
	case *services.Event_CreateNodeRequest:
		r.handler.Replicate(createAction, createNodeURL,
			event.CreateNodeRequest, &services.CreateNodeResponse{})
	case *services.Event_UpdateNodeRequest:
		objID := event.UpdateNodeRequest.Node.UUID
		r.handler.Replicate(updateAction, updateNodeURL+objID,
			event.UpdateNodeRequest, &services.UpdateNodeResponse{})
	case *services.Event_DeleteNodeRequest:
		objID := event.DeleteNodeRequest.ID
		r.handler.Replicate(deleteAction, updateNodeURL+objID,
			event.DeleteNodeRequest, &services.DeleteNodeResponse{})
	// handle ports
	case *services.Event_CreatePortRequest:
		r.handler.Replicate(createAction, createPortURL,
			event.CreatePortRequest, &services.CreatePortResponse{})
	case *services.Event_UpdatePortRequest:
		objID := event.UpdatePortRequest.Port.UUID
		r.handler.Replicate(updateAction, updatePortURL+objID,
			event.UpdatePortRequest, &services.UpdatePortResponse{})
	case *services.Event_DeletePortRequest:
		objID := event.DeletePortRequest.ID
		r.handler.Replicate(deleteAction, updatePortURL+objID,
			event.DeletePortRequest, &services.DeletePortResponse{})
	// handle hardware
	case *services.Event_CreateHardwareRequest:
		r.handler.Replicate(createAction, createHardwareURL,
			event.CreateHardwareRequest, &services.CreateHardwareResponse{})
	case *services.Event_UpdateHardwareRequest:
		objID := event.UpdateHardwareRequest.Hardware.UUID
		r.handler.Replicate(updateAction, updateHardwareURL+objID,
			event.UpdateHardwareRequest, &services.UpdateHardwareResponse{})
	case *services.Event_DeleteHardwareRequest:
		objID := event.DeleteHardwareRequest.ID
		r.handler.Replicate(deleteAction, updateHardwareURL+objID,
			event.DeleteHardwareRequest, &services.DeleteHardwareResponse{})
		// handle card
	case *services.Event_CreateCardRequest:
		r.handler.Replicate(createAction, createCardURL,
			event.CreateCardRequest, &services.CreateCardResponse{})
	case *services.Event_UpdateCardRequest:
		objID := event.UpdateCardRequest.Card.UUID
		r.handler.Replicate(updateAction, updateCardURL+objID,
			event.UpdateCardRequest, &services.UpdateCardResponse{})
	case *services.Event_DeleteCardRequest:
		objID := event.DeleteCardRequest.ID
		r.handler.Replicate(deleteAction, updateCardURL+objID,
			event.DeleteCardRequest, &services.DeleteCardResponse{})
	}

	return e, nil
}

// Stop replication routine
func (r *Replicator) Stop() {
	r.stopServiceContext()
	r.serviceWaitGroup.Wait()
}
