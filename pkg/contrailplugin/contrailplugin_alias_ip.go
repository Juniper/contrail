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

// AliasIPIntent
//   A struct to store attributes related to AliasIP
//   needed by Intent Compiler
type AliasIPIntent struct {
	Uuid string
}

// EvaluateAliasIP - evaluates the AliasIP
func EvaluateAliasIP(obj interface{}) {
	resourceObj := obj.(AliasIPIntent)
	log.Println("EvaluateAliasIP Called ", resourceObj)
}

// CreateAliasIP handles create request
func (service *PluginService) CreateAliasIP(ctx context.Context, request *services.CreateAliasIPRequest) (*services.CreateAliasIPResponse, error) {
	log.Println(" CreateAliasIP Entered")

	obj := request.GetAliasIP()

	intentObj := AliasIPIntent{
		Uuid: obj.GetUUID(),
	}

	if _, ok := compilationif.ObjsCache.Load("AliasIPIntent"); !ok {
		compilationif.ObjsCache.Store("AliasIPIntent", &sync.Map{})
	}
	objMap, ok := compilationif.ObjsCache.Load("AliasIPIntent")
	if ok {
		objMap.(*sync.Map).Store(obj.GetUUID(), intentObj)
	}

	service.Debug("CreateAliasIP", objMap.(*sync.Map))

	EvaluateDependencies(obj, "AliasIP")

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().CreateAliasIP(ctx, request)
}

// UpdateAliasIP handles update request
func (service *PluginService) UpdateAliasIP(ctx context.Context, request *services.UpdateAliasIPRequest) (*services.UpdateAliasIPResponse, error) {
	log.Println(" UpdateAliasIP ENTERED")

	obj := request.GetAliasIP()

	intentObj := AliasIPIntent{
		Uuid: obj.GetUUID(),
	}

	if _, ok := compilationif.ObjsCache.Load("AliasIPIntent"); !ok {
		compilationif.ObjsCache.Store("AliasIPIntent", &sync.Map{})
	}

	EvaluateDependencies(obj, "AliasIP")

	objMap, ok := compilationif.ObjsCache.Load("Intent")
	if ok {
		objMap.(*sync.Map).Store(obj.GetUUID(), intentObj)
	}

	service.Debug("UpdateAliasIP", objMap.(*sync.Map))

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().UpdateAliasIP(ctx, request)
}

// DeleteAliasIP handles delete request
func (service *PluginService) DeleteAliasIP(ctx context.Context, request *services.DeleteAliasIPRequest) (*services.DeleteAliasIPResponse, error) {
	log.Println(" DeleteAliasIP ENTERED")

	objUUID := request.GetID()

	//intentObj := AliasIPIntent {
	//AliasIP: *obj,
	//}

	//EvaluateDependencies(intentObj, "AliasIP")

	objMap, ok := compilationif.ObjsCache.Load("AliasIPIntent")
	if ok {
		objMap.(*sync.Map).Delete(objUUID)
	}

	service.Debug("DeleteAliasIP", objMap.(*sync.Map))

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().DeleteAliasIP(ctx, request)
}

// GetAliasIP handles get request
func (service *PluginService) GetAliasIP(ctx context.Context, request *services.GetAliasIPRequest) (*services.GetAliasIPResponse, error) {
	objMap, ok := compilationif.ObjsCache.Load("AliasIP")
	if !ok {
		return nil, errors.New("AliasIP get failed ")
	}

	obj, ok := objMap.(*sync.Map).Load(request.GetID())
	if !ok {
		return nil, errors.New("AliasIP get failed ")
	}

	response := &services.GetAliasIPResponse{
		AliasIP: obj.(*models.AliasIP),
	}
	return response, nil
}
