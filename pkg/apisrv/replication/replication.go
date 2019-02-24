package replication

import (
	"context"
	"sync"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	apicommon "github.com/Juniper/contrail/pkg/apisrv/common"
	"github.com/Juniper/contrail/pkg/db/cache"
	"github.com/Juniper/contrail/pkg/logutil"
	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/services"
)

const (
	createTagURL            = "tags"
	createVirtualNetworkURL = "virtual-networks"
	createNetworkIpamURL    = "network-ipams"
	createPortURL           = "ports"
	createNodeProfileURL    = "node-profiles"
	createEndSystemURL      = "end-systems"
	createHardwareURL       = "hardwares"
	createCardURL           = "cards"
	updateTagURL            = "tag/"
	updateVirtualNetworkURL = "virtual-network/"
	updateNetworkIpamURL    = "network-ipam/"
	updatePortURL           = "port/"
	updateNodeProfileURL    = "node-profile/"
	updateEndSystemURL      = "end-system/"
	updateHardwareURL       = "hardware/"
	updateCardURL           = "card/"
	vnIDKey                 = "virtual_network_network_id"

	deleteAction = "delete"
	createAction = "create"
	updateAction = "update"
)

// Replicator is an implementation to replicate objects to python API
type Replicator struct {
	serviceWaitGroup   *sync.WaitGroup
	stopServiceContext context.CancelFunc
	producer           *cache.DB
	vncAPIHandle       *vncAPIHandle
	log                *logrus.Entry
}

// New initializes replication data
func New(cacheDB *cache.DB,
	epStore *apicommon.EndpointStore) (*Replicator, error) {

	if err := logutil.Configure(viper.GetString("log_level")); err != nil {
		return nil, err
	}
	return &Replicator{
		producer:         cacheDB,
		serviceWaitGroup: &sync.WaitGroup{},
		vncAPIHandle:     newVncAPIHandle(epStore),
		log:              logutil.NewLogger("vnc_replication"),
	}, nil
}

// Start replication service
func (r *Replicator) Start() error {

	var serviceContext context.Context
	serviceContext, r.stopServiceContext = context.WithCancel(context.Background())
	watcher, err := r.producer.AddWatcher(serviceContext, 0)
	if err != nil {
		return err
	}
	r.serviceWaitGroup.Add(1)
	go func() {
		defer r.serviceWaitGroup.Done()
		for {
			select {
			case <-serviceContext.Done():
				r.log.Info("Stopping VNC API replication service")
				return
			case e := <-watcher.Chan():
				r.process(e)
			}
		}
	}()
	return nil
}

func (r *Replicator) nodeToVNCEndSystem(node *models.Node) interface{} {
	endSystem := node.ToMap()
	if hostname, ok := endSystem["hostname"]; ok {
		endSystem["end_system_hostname"] = hostname
	}
	output := map[string]map[string]interface{}{"end-system": endSystem}
	return output
}

func (r *Replicator) portToVNCPort(port *models.Port) interface{} {
	portVNC := port.ToMap()
	portVNC["parent_type"] = "end-system"
	if bmsPortInfo, ok := portVNC["bms_port_info"]; ok {
		portVNC["port_bms_port_info"] = bmsPortInfo
	}
	output := map[string]map[string]interface{}{"port": portVNC}
	return output
}

func (r *Replicator) vnToVNCVirtualNetwork(vn *models.VirtualNetwork) interface{} {
	vnVNC := vn.ToMap()
	if _, ok := vnVNC[vnIDKey]; ok {
		delete(vnVNC, vnIDKey)
	}
	output := map[string]map[string]interface{}{"virtual-network": vnVNC}
	return output
}

func (r *Replicator) process(e *services.Event) { //nolint: gocyclo
	switch event := e.Request.(type) {
	// watch endpoint event and prepare clients
	case *services.Event_CreateEndpointRequest:
		ep := event.CreateEndpointRequest.Endpoint
		r.vncAPIHandle.createClient(ep)
	case *services.Event_UpdateEndpointRequest:
		ep := event.UpdateEndpointRequest.Endpoint
		r.vncAPIHandle.updateClient(ep)
	case *services.Event_DeleteEndpointRequest:
		id := event.DeleteEndpointRequest.ID
		r.vncAPIHandle.deleteClient(id)
		// handle tag
	case *services.Event_CreateTagRequest:
		event.CreateTagRequest.Tag.TagID = ""
		r.vncAPIHandle.replicate(createAction, createTagURL,
			event.CreateTagRequest, &services.CreateTagResponse{})
	case *services.Event_UpdateTagRequest:
		event.UpdateTagRequest.Tag.TagID = ""
		objID := event.UpdateTagRequest.Tag.UUID
		r.vncAPIHandle.replicate(updateAction, updateTagURL+objID,
			event.UpdateTagRequest, &services.UpdateTagResponse{})
	case *services.Event_DeleteTagRequest:
		objID := event.DeleteTagRequest.ID
		r.vncAPIHandle.replicate(deleteAction, updateTagURL+objID,
			event.DeleteTagRequest, &services.DeleteTagResponse{})
	// handle virtual-network
	case *services.Event_CreateVirtualNetworkRequest:
		event.CreateVirtualNetworkRequest.VirtualNetwork.VirtualNetworkNetworkID = 0
		r.vncAPIHandle.replicate(createAction, createVirtualNetworkURL,
			r.vnToVNCVirtualNetwork(event.CreateVirtualNetworkRequest.VirtualNetwork),
			&services.CreateVirtualNetworkResponse{})
	case *services.Event_UpdateVirtualNetworkRequest:
		event.UpdateVirtualNetworkRequest.VirtualNetwork.VirtualNetworkNetworkID = 0
		objID := event.UpdateVirtualNetworkRequest.VirtualNetwork.UUID
		r.vncAPIHandle.replicate(updateAction, updateVirtualNetworkURL+objID,
			r.vnToVNCVirtualNetwork(event.UpdateVirtualNetworkRequest.VirtualNetwork),
			&services.UpdateVirtualNetworkResponse{})
	case *services.Event_DeleteVirtualNetworkRequest:
		objID := event.DeleteVirtualNetworkRequest.ID
		r.vncAPIHandle.replicate(deleteAction, updateVirtualNetworkURL+objID,
			event.DeleteVirtualNetworkRequest, &services.DeleteVirtualNetworkResponse{})
		// handle network-ipam
	case *services.Event_CreateNetworkIpamRequest:
		r.vncAPIHandle.replicate(createAction, createNetworkIpamURL,
			event.CreateNetworkIpamRequest, &services.CreateNetworkIpamResponse{})
	case *services.Event_UpdateNetworkIpamRequest:
		objID := event.UpdateNetworkIpamRequest.NetworkIpam.UUID
		r.vncAPIHandle.replicate(updateAction, updateNetworkIpamURL+objID,
			event.UpdateNetworkIpamRequest, &services.UpdateNetworkIpamResponse{})
	case *services.Event_DeleteNetworkIpamRequest:
		objID := event.DeleteNetworkIpamRequest.ID
		r.vncAPIHandle.replicate(deleteAction, updateNetworkIpamURL+objID,
			event.DeleteNetworkIpamRequest, &services.DeleteNetworkIpamResponse{})
	// handle node-profile
	case *services.Event_CreateNodeProfileRequest:
		r.vncAPIHandle.replicate(createAction, createNodeProfileURL,
			event.CreateNodeProfileRequest, &services.CreateNodeProfileResponse{})
	case *services.Event_UpdateNodeProfileRequest:
		objID := event.UpdateNodeProfileRequest.NodeProfile.UUID
		r.vncAPIHandle.replicate(updateAction, updateNodeProfileURL+objID,
			event.UpdateNodeProfileRequest, &services.UpdateNodeProfileResponse{})
	case *services.Event_DeleteNodeProfileRequest:
		objID := event.DeleteNodeProfileRequest.ID
		r.vncAPIHandle.replicate(deleteAction, updateNodeProfileURL+objID,
			event.DeleteNodeProfileRequest, &services.DeleteNodeProfileResponse{})
	// handle nodes(end-systems)
	case *services.Event_CreateNodeRequest:
		r.vncAPIHandle.replicate(createAction, createEndSystemURL,
			r.nodeToVNCEndSystem(event.CreateNodeRequest.Node),
			&services.CreateNodeResponse{})
	case *services.Event_UpdateNodeRequest:
		objID := event.UpdateNodeRequest.Node.UUID
		r.vncAPIHandle.replicate(updateAction, updateEndSystemURL+objID,
			r.nodeToVNCEndSystem(event.UpdateNodeRequest.Node),
			&services.UpdateNodeResponse{})
	case *services.Event_DeleteNodeRequest:
		objID := event.DeleteNodeRequest.ID
		r.vncAPIHandle.replicate(deleteAction, updateEndSystemURL+objID,
			event.DeleteNodeRequest, &services.DeleteNodeResponse{})
	// handle ports
	case *services.Event_CreatePortRequest:
		r.vncAPIHandle.replicate(createAction, createPortURL,
			r.portToVNCPort(event.CreatePortRequest.Port),
			&services.CreatePortResponse{})
	case *services.Event_UpdatePortRequest:
		objID := event.UpdatePortRequest.Port.UUID
		r.vncAPIHandle.replicate(updateAction, updatePortURL+objID,
			r.portToVNCPort(event.UpdatePortRequest.Port),
			&services.UpdatePortResponse{})
	case *services.Event_DeletePortRequest:
		objID := event.DeletePortRequest.ID
		r.vncAPIHandle.replicate(deleteAction, updatePortURL+objID,
			event.DeletePortRequest, &services.DeletePortResponse{})
	// handle hardware
	case *services.Event_CreateHardwareRequest:
		r.vncAPIHandle.replicate(createAction, createHardwareURL,
			event.CreateHardwareRequest, &services.CreateHardwareResponse{})
	case *services.Event_UpdateHardwareRequest:
		objID := event.UpdateHardwareRequest.Hardware.UUID
		r.vncAPIHandle.replicate(updateAction, updateHardwareURL+objID,
			event.UpdateHardwareRequest, &services.UpdateHardwareResponse{})
	case *services.Event_DeleteHardwareRequest:
		objID := event.DeleteHardwareRequest.ID
		r.vncAPIHandle.replicate(deleteAction, updateHardwareURL+objID,
			event.DeleteHardwareRequest, &services.DeleteHardwareResponse{})
		// handle card
	case *services.Event_CreateCardRequest:
		r.vncAPIHandle.replicate(createAction, createCardURL,
			event.CreateCardRequest, &services.CreateCardResponse{})
	case *services.Event_UpdateCardRequest:
		objID := event.UpdateCardRequest.Card.UUID
		r.vncAPIHandle.replicate(updateAction, updateCardURL+objID,
			event.UpdateCardRequest, &services.UpdateCardResponse{})
	case *services.Event_DeleteCardRequest:
		objID := event.DeleteCardRequest.ID
		r.vncAPIHandle.replicate(deleteAction, updateCardURL+objID,
			event.DeleteCardRequest, &services.DeleteCardResponse{})
	}
}

// Stop replication routine
func (r *Replicator) Stop() {
	r.stopServiceContext()
	r.serviceWaitGroup.Wait()
}
