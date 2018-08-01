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

// BridgeDomainIntent
//   A struct to store attributes related to BridgeDomain
//   needed by Intent Compiler
type BridgeDomainIntent struct {
	Uuid string
}

// EvaluateBridgeDomain - evaluates the BridgeDomain
func EvaluateBridgeDomain(obj interface{}) {
	resourceObj := obj.(BridgeDomainIntent)
	log.Println("EvaluateBridgeDomain Called ", resourceObj)
}

// CreateBridgeDomain handles create request
func (service *PluginService) CreateBridgeDomain(ctx context.Context, request *services.CreateBridgeDomainRequest) (*services.CreateBridgeDomainResponse, error) {
	log.Println(" CreateBridgeDomain Entered")

	obj := request.GetBridgeDomain()

	intentObj := BridgeDomainIntent{
		Uuid: obj.GetUUID(),
	}

	if _, ok := compilationif.ObjsCache.Load("BridgeDomainIntent"); !ok {
		compilationif.ObjsCache.Store("BridgeDomainIntent", &sync.Map{})
	}
	objMap, ok := compilationif.ObjsCache.Load("BridgeDomainIntent")
	if ok {
		objMap.(*sync.Map).Store(obj.GetUUID(), intentObj)
	}

	service.Debug("CreateBridgeDomain", objMap.(*sync.Map))

	EvaluateDependencies(obj, "BridgeDomain")

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().CreateBridgeDomain(ctx, request)
}

// UpdateBridgeDomain handles update request
func (service *PluginService) UpdateBridgeDomain(ctx context.Context, request *services.UpdateBridgeDomainRequest) (*services.UpdateBridgeDomainResponse, error) {
	log.Println(" UpdateBridgeDomain ENTERED")

	obj := request.GetBridgeDomain()

	intentObj := BridgeDomainIntent{
		Uuid: obj.GetUUID(),
	}

	if _, ok := compilationif.ObjsCache.Load("BridgeDomainIntent"); !ok {
		compilationif.ObjsCache.Store("BridgeDomainIntent", &sync.Map{})
	}

	EvaluateDependencies(obj, "BridgeDomain")

	objMap, ok := compilationif.ObjsCache.Load("Intent")
	if ok {
		objMap.(*sync.Map).Store(obj.GetUUID(), intentObj)
	}

	service.Debug("UpdateBridgeDomain", objMap.(*sync.Map))

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().UpdateBridgeDomain(ctx, request)
}

// DeleteBridgeDomain handles delete request
func (service *PluginService) DeleteBridgeDomain(ctx context.Context, request *services.DeleteBridgeDomainRequest) (*services.DeleteBridgeDomainResponse, error) {
	log.Println(" DeleteBridgeDomain ENTERED")

	objUUID := request.GetID()

	//intentObj := BridgeDomainIntent {
	//BridgeDomain: *obj,
	//}

	//EvaluateDependencies(intentObj, "BridgeDomain")

	objMap, ok := compilationif.ObjsCache.Load("BridgeDomainIntent")
	if ok {
		objMap.(*sync.Map).Delete(objUUID)
	}

	service.Debug("DeleteBridgeDomain", objMap.(*sync.Map))

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().DeleteBridgeDomain(ctx, request)
}

// GetBridgeDomain handles get request
func (service *PluginService) GetBridgeDomain(ctx context.Context, request *services.GetBridgeDomainRequest) (*services.GetBridgeDomainResponse, error) {
	objMap, ok := compilationif.ObjsCache.Load("BridgeDomain")
	if !ok {
		return nil, errors.New("BridgeDomain get failed ")
	}

	obj, ok := objMap.(*sync.Map).Load(request.GetID())
	if !ok {
		return nil, errors.New("BridgeDomain get failed ")
	}

	response := &services.GetBridgeDomainResponse{
		BridgeDomain: obj.(*models.BridgeDomain),
	}
	return response, nil
}
