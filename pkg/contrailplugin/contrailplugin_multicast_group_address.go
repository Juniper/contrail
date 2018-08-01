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

// MulticastGroupAddressIntent
//   A struct to store attributes related to MulticastGroupAddress
//   needed by Intent Compiler
type MulticastGroupAddressIntent struct {
	Uuid string
}

// EvaluateMulticastGroupAddress - evaluates the MulticastGroupAddress
func EvaluateMulticastGroupAddress(obj interface{}) {
	resourceObj := obj.(MulticastGroupAddressIntent)
	log.Println("EvaluateMulticastGroupAddress Called ", resourceObj)
}

// CreateMulticastGroupAddress handles create request
func (service *PluginService) CreateMulticastGroupAddress(ctx context.Context, request *services.CreateMulticastGroupAddressRequest) (*services.CreateMulticastGroupAddressResponse, error) {
	log.Println(" CreateMulticastGroupAddress Entered")

	obj := request.GetMulticastGroupAddress()

	intentObj := MulticastGroupAddressIntent{
		Uuid: obj.GetUUID(),
	}

	if _, ok := compilationif.ObjsCache.Load("MulticastGroupAddressIntent"); !ok {
		compilationif.ObjsCache.Store("MulticastGroupAddressIntent", &sync.Map{})
	}
	objMap, ok := compilationif.ObjsCache.Load("MulticastGroupAddressIntent")
	if ok {
		objMap.(*sync.Map).Store(obj.GetUUID(), intentObj)
	}

	service.Debug("CreateMulticastGroupAddress", objMap.(*sync.Map))

	EvaluateDependencies(obj, "MulticastGroupAddress")

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().CreateMulticastGroupAddress(ctx, request)
}

// UpdateMulticastGroupAddress handles update request
func (service *PluginService) UpdateMulticastGroupAddress(ctx context.Context, request *services.UpdateMulticastGroupAddressRequest) (*services.UpdateMulticastGroupAddressResponse, error) {
	log.Println(" UpdateMulticastGroupAddress ENTERED")

	obj := request.GetMulticastGroupAddress()

	intentObj := MulticastGroupAddressIntent{
		Uuid: obj.GetUUID(),
	}

	if _, ok := compilationif.ObjsCache.Load("MulticastGroupAddressIntent"); !ok {
		compilationif.ObjsCache.Store("MulticastGroupAddressIntent", &sync.Map{})
	}

	EvaluateDependencies(obj, "MulticastGroupAddress")

	objMap, ok := compilationif.ObjsCache.Load("Intent")
	if ok {
		objMap.(*sync.Map).Store(obj.GetUUID(), intentObj)
	}

	service.Debug("UpdateMulticastGroupAddress", objMap.(*sync.Map))

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().UpdateMulticastGroupAddress(ctx, request)
}

// DeleteMulticastGroupAddress handles delete request
func (service *PluginService) DeleteMulticastGroupAddress(ctx context.Context, request *services.DeleteMulticastGroupAddressRequest) (*services.DeleteMulticastGroupAddressResponse, error) {
	log.Println(" DeleteMulticastGroupAddress ENTERED")

	objUUID := request.GetID()

	//intentObj := MulticastGroupAddressIntent {
	//MulticastGroupAddress: *obj,
	//}

	//EvaluateDependencies(intentObj, "MulticastGroupAddress")

	objMap, ok := compilationif.ObjsCache.Load("MulticastGroupAddressIntent")
	if ok {
		objMap.(*sync.Map).Delete(objUUID)
	}

	service.Debug("DeleteMulticastGroupAddress", objMap.(*sync.Map))

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().DeleteMulticastGroupAddress(ctx, request)
}

// GetMulticastGroupAddress handles get request
func (service *PluginService) GetMulticastGroupAddress(ctx context.Context, request *services.GetMulticastGroupAddressRequest) (*services.GetMulticastGroupAddressResponse, error) {
	objMap, ok := compilationif.ObjsCache.Load("MulticastGroupAddress")
	if !ok {
		return nil, errors.New("MulticastGroupAddress get failed ")
	}

	obj, ok := objMap.(*sync.Map).Load(request.GetID())
	if !ok {
		return nil, errors.New("MulticastGroupAddress get failed ")
	}

	response := &services.GetMulticastGroupAddressResponse{
		MulticastGroupAddress: obj.(*models.MulticastGroupAddress),
	}
	return response, nil
}
