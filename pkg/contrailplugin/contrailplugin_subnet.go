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

// SubnetIntent
//   A struct to store attributes related to Subnet
//   needed by Intent Compiler
type SubnetIntent struct {
	Uuid string
}

// EvaluateSubnet - evaluates the Subnet
func EvaluateSubnet(obj interface{}) {
	resourceObj := obj.(SubnetIntent)
	log.Println("EvaluateSubnet Called ", resourceObj)
}

// CreateSubnet handles create request
func (service *PluginService) CreateSubnet(ctx context.Context, request *services.CreateSubnetRequest) (*services.CreateSubnetResponse, error) {
	log.Println(" CreateSubnet Entered")

	obj := request.GetSubnet()

	intentObj := SubnetIntent{
		Uuid: obj.GetUUID(),
	}

	if _, ok := compilationif.ObjsCache.Load("SubnetIntent"); !ok {
		compilationif.ObjsCache.Store("SubnetIntent", &sync.Map{})
	}
	objMap, ok := compilationif.ObjsCache.Load("SubnetIntent")
	if ok {
		objMap.(*sync.Map).Store(obj.GetUUID(), intentObj)
	}

	service.Debug("CreateSubnet", objMap.(*sync.Map))

	EvaluateDependencies(obj, "Subnet")

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().CreateSubnet(ctx, request)
}

// UpdateSubnet handles update request
func (service *PluginService) UpdateSubnet(ctx context.Context, request *services.UpdateSubnetRequest) (*services.UpdateSubnetResponse, error) {
	log.Println(" UpdateSubnet ENTERED")

	obj := request.GetSubnet()

	intentObj := SubnetIntent{
		Uuid: obj.GetUUID(),
	}

	if _, ok := compilationif.ObjsCache.Load("SubnetIntent"); !ok {
		compilationif.ObjsCache.Store("SubnetIntent", &sync.Map{})
	}

	EvaluateDependencies(obj, "Subnet")

	objMap, ok := compilationif.ObjsCache.Load("Intent")
	if ok {
		objMap.(*sync.Map).Store(obj.GetUUID(), intentObj)
	}

	service.Debug("UpdateSubnet", objMap.(*sync.Map))

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().UpdateSubnet(ctx, request)
}

// DeleteSubnet handles delete request
func (service *PluginService) DeleteSubnet(ctx context.Context, request *services.DeleteSubnetRequest) (*services.DeleteSubnetResponse, error) {
	log.Println(" DeleteSubnet ENTERED")

	objUUID := request.GetID()

	//intentObj := SubnetIntent {
	//Subnet: *obj,
	//}

	//EvaluateDependencies(intentObj, "Subnet")

	objMap, ok := compilationif.ObjsCache.Load("SubnetIntent")
	if ok {
		objMap.(*sync.Map).Delete(objUUID)
	}

	service.Debug("DeleteSubnet", objMap.(*sync.Map))

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().DeleteSubnet(ctx, request)
}

// GetSubnet handles get request
func (service *PluginService) GetSubnet(ctx context.Context, request *services.GetSubnetRequest) (*services.GetSubnetResponse, error) {
	objMap, ok := compilationif.ObjsCache.Load("Subnet")
	if !ok {
		return nil, errors.New("Subnet get failed ")
	}

	obj, ok := objMap.(*sync.Map).Load(request.GetID())
	if !ok {
		return nil, errors.New("Subnet get failed ")
	}

	response := &services.GetSubnetResponse{
		Subnet: obj.(*models.Subnet),
	}
	return response, nil
}
