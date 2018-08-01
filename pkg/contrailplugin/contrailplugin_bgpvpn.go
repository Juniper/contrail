// nolint
package contrailplugin

import (
	"context"
	"errors"
	"sync"

	"github.com/Juniper/contrail/pkg/compilationif"
	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/services"
	log "github.com/sirupsen/logrus"
)

// BGPVPNIntent
//   A struct to store attributes related to BGPVPN
//   needed by Intent Compiler
type BGPVPNIntent struct {
	Uuid string
}

// EvaluateBGPVPN - evaluates the BGPVPN
func EvaluateBGPVPN(obj interface{}) {
	resourceObj := obj.(BGPVPNIntent)
	log.Println("EvaluateBGPVPN Called ", resourceObj)
}

// CreateBGPVPN handles create request
func (service *PluginService) CreateBGPVPN(ctx context.Context, request *services.CreateBGPVPNRequest) (*services.CreateBGPVPNResponse, error) {
	log.Println(" CreateBGPVPN Entered")

	obj := request.GetBGPVPN()

	intentObj := BGPVPNIntent{
		Uuid: obj.GetUUID(),
	}

	if _, ok := compilationif.ObjsCache.Load("BGPVPNIntent"); !ok {
		compilationif.ObjsCache.Store("BGPVPNIntent", &sync.Map{})
	}
	objMap, ok := compilationif.ObjsCache.Load("BGPVPNIntent")
	if ok {
		objMap.(*sync.Map).Store(obj.GetUUID(), intentObj)
	}

	service.Debug("CreateBGPVPN", objMap.(*sync.Map))

	EvaluateDependencies(obj, "BGPVPN")

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().CreateBGPVPN(ctx, request)
}

// UpdateBGPVPN handles update request
func (service *PluginService) UpdateBGPVPN(ctx context.Context, request *services.UpdateBGPVPNRequest) (*services.UpdateBGPVPNResponse, error) {
	log.Println(" UpdateBGPVPN ENTERED")

	obj := request.GetBGPVPN()

	intentObj := BGPVPNIntent{
		Uuid: obj.GetUUID(),
	}

	if _, ok := compilationif.ObjsCache.Load("BGPVPNIntent"); !ok {
		compilationif.ObjsCache.Store("BGPVPNIntent", &sync.Map{})
	}

	EvaluateDependencies(obj, "BGPVPN")

	objMap, ok := compilationif.ObjsCache.Load("Intent")
	if ok {
		objMap.(*sync.Map).Store(obj.GetUUID(), intentObj)
	}

	service.Debug("UpdateBGPVPN", objMap.(*sync.Map))

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().UpdateBGPVPN(ctx, request)
}

// DeleteBGPVPN handles delete request
func (service *PluginService) DeleteBGPVPN(ctx context.Context, request *services.DeleteBGPVPNRequest) (*services.DeleteBGPVPNResponse, error) {
	log.Println(" DeleteBGPVPN ENTERED")

	objUUID := request.GetID()

	//intentObj := BGPVPNIntent {
	//BGPVPN: *obj,
	//}

	//EvaluateDependencies(intentObj, "BGPVPN")

	objMap, ok := compilationif.ObjsCache.Load("BGPVPNIntent")
	if ok {
		objMap.(*sync.Map).Delete(objUUID)
	}

	service.Debug("DeleteBGPVPN", objMap.(*sync.Map))

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().DeleteBGPVPN(ctx, request)
}

// GetBGPVPN handles get request
func (service *PluginService) GetBGPVPN(ctx context.Context, request *services.GetBGPVPNRequest) (*services.GetBGPVPNResponse, error) {
	objMap, ok := compilationif.ObjsCache.Load("BGPVPN")
	if !ok {
		return nil, errors.New("BGPVPN get failed ")
	}

	obj, ok := objMap.(*sync.Map).Load(request.GetID())
	if !ok {
		return nil, errors.New("BGPVPN get failed ")
	}

	response := &services.GetBGPVPNResponse{
		BGPVPN: obj.(*models.BGPVPN),
	}
	return response, nil
}
