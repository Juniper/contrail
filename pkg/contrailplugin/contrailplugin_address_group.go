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

// AddressGroupIntent
//   A struct to store attributes related to AddressGroup
//   needed by Intent Compiler
type AddressGroupIntent struct {
	Uuid string
}

// EvaluateAddressGroup - evaluates the AddressGroup
func EvaluateAddressGroup(obj interface{}) {
	resourceObj := obj.(AddressGroupIntent)
	log.Println("EvaluateAddressGroup Called ", resourceObj)
}

// CreateAddressGroup handles create request
func (service *PluginService) CreateAddressGroup(ctx context.Context, request *services.CreateAddressGroupRequest) (*services.CreateAddressGroupResponse, error) {
	log.Println(" CreateAddressGroup Entered")

	obj := request.GetAddressGroup()

	intentObj := AddressGroupIntent{
		Uuid: obj.GetUUID(),
	}

	if _, ok := compilationif.ObjsCache.Load("AddressGroupIntent"); !ok {
		compilationif.ObjsCache.Store("AddressGroupIntent", &sync.Map{})
	}
	objMap, ok := compilationif.ObjsCache.Load("AddressGroupIntent")
	if ok {
		objMap.(*sync.Map).Store(obj.GetUUID(), intentObj)
	}

	service.Debug("CreateAddressGroup", objMap.(*sync.Map))

	EvaluateDependencies(obj, "AddressGroup")

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().CreateAddressGroup(ctx, request)
}

// UpdateAddressGroup handles update request
func (service *PluginService) UpdateAddressGroup(ctx context.Context, request *services.UpdateAddressGroupRequest) (*services.UpdateAddressGroupResponse, error) {
	log.Println(" UpdateAddressGroup ENTERED")

	obj := request.GetAddressGroup()

	intentObj := AddressGroupIntent{
		Uuid: obj.GetUUID(),
	}

	if _, ok := compilationif.ObjsCache.Load("AddressGroupIntent"); !ok {
		compilationif.ObjsCache.Store("AddressGroupIntent", &sync.Map{})
	}

	EvaluateDependencies(obj, "AddressGroup")

	objMap, ok := compilationif.ObjsCache.Load("Intent")
	if ok {
		objMap.(*sync.Map).Store(obj.GetUUID(), intentObj)
	}

	service.Debug("UpdateAddressGroup", objMap.(*sync.Map))

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().UpdateAddressGroup(ctx, request)
}

// DeleteAddressGroup handles delete request
func (service *PluginService) DeleteAddressGroup(ctx context.Context, request *services.DeleteAddressGroupRequest) (*services.DeleteAddressGroupResponse, error) {
	log.Println(" DeleteAddressGroup ENTERED")

	objUUID := request.GetID()

	//intentObj := AddressGroupIntent {
	//AddressGroup: *obj,
	//}

	//EvaluateDependencies(intentObj, "AddressGroup")

	objMap, ok := compilationif.ObjsCache.Load("AddressGroupIntent")
	if ok {
		objMap.(*sync.Map).Delete(objUUID)
	}

	service.Debug("DeleteAddressGroup", objMap.(*sync.Map))

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().DeleteAddressGroup(ctx, request)
}

// GetAddressGroup handles get request
func (service *PluginService) GetAddressGroup(ctx context.Context, request *services.GetAddressGroupRequest) (*services.GetAddressGroupResponse, error) {
	objMap, ok := compilationif.ObjsCache.Load("AddressGroup")
	if !ok {
		return nil, errors.New("AddressGroup get failed ")
	}

	obj, ok := objMap.(*sync.Map).Load(request.GetID())
	if !ok {
		return nil, errors.New("AddressGroup get failed ")
	}

	response := &services.GetAddressGroupResponse{
		AddressGroup: obj.(*models.AddressGroup),
	}
	return response, nil
}
