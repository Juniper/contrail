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

// ContrailConfigNodeIntent
//   A struct to store attributes related to ContrailConfigNode
//   needed by Intent Compiler
type ContrailConfigNodeIntent struct {
	Uuid string
}

// EvaluateContrailConfigNode - evaluates the ContrailConfigNode
func EvaluateContrailConfigNode(obj interface{}) {
	resourceObj := obj.(ContrailConfigNodeIntent)
	log.Println("EvaluateContrailConfigNode Called ", resourceObj)
}

// CreateContrailConfigNode handles create request
func (service *PluginService) CreateContrailConfigNode(ctx context.Context, request *services.CreateContrailConfigNodeRequest) (*services.CreateContrailConfigNodeResponse, error) {
	log.Println(" CreateContrailConfigNode Entered")

	obj := request.GetContrailConfigNode()

	intentObj := ContrailConfigNodeIntent{
		Uuid: obj.GetUUID(),
	}

	if _, ok := compilationif.ObjsCache.Load("ContrailConfigNodeIntent"); !ok {
		compilationif.ObjsCache.Store("ContrailConfigNodeIntent", &sync.Map{})
	}
	objMap, ok := compilationif.ObjsCache.Load("ContrailConfigNodeIntent")
	if ok {
		objMap.(*sync.Map).Store(obj.GetUUID(), intentObj)
	}

	service.Debug("CreateContrailConfigNode", objMap.(*sync.Map))

	EvaluateDependencies(obj, "ContrailConfigNode")

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().CreateContrailConfigNode(ctx, request)
}

// UpdateContrailConfigNode handles update request
func (service *PluginService) UpdateContrailConfigNode(ctx context.Context, request *services.UpdateContrailConfigNodeRequest) (*services.UpdateContrailConfigNodeResponse, error) {
	log.Println(" UpdateContrailConfigNode ENTERED")

	obj := request.GetContrailConfigNode()

	intentObj := ContrailConfigNodeIntent{
		Uuid: obj.GetUUID(),
	}

	if _, ok := compilationif.ObjsCache.Load("ContrailConfigNodeIntent"); !ok {
		compilationif.ObjsCache.Store("ContrailConfigNodeIntent", &sync.Map{})
	}

	EvaluateDependencies(obj, "ContrailConfigNode")

	objMap, ok := compilationif.ObjsCache.Load("Intent")
	if ok {
		objMap.(*sync.Map).Store(obj.GetUUID(), intentObj)
	}

	service.Debug("UpdateContrailConfigNode", objMap.(*sync.Map))

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().UpdateContrailConfigNode(ctx, request)
}

// DeleteContrailConfigNode handles delete request
func (service *PluginService) DeleteContrailConfigNode(ctx context.Context, request *services.DeleteContrailConfigNodeRequest) (*services.DeleteContrailConfigNodeResponse, error) {
	log.Println(" DeleteContrailConfigNode ENTERED")

	objUUID := request.GetID()

	//intentObj := ContrailConfigNodeIntent {
	//ContrailConfigNode: *obj,
	//}

	//EvaluateDependencies(intentObj, "ContrailConfigNode")

	objMap, ok := compilationif.ObjsCache.Load("ContrailConfigNodeIntent")
	if ok {
		objMap.(*sync.Map).Delete(objUUID)
	}

	service.Debug("DeleteContrailConfigNode", objMap.(*sync.Map))

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().DeleteContrailConfigNode(ctx, request)
}

// GetContrailConfigNode handles get request
func (service *PluginService) GetContrailConfigNode(ctx context.Context, request *services.GetContrailConfigNodeRequest) (*services.GetContrailConfigNodeResponse, error) {
	objMap, ok := compilationif.ObjsCache.Load("ContrailConfigNode")
	if !ok {
		return nil, errors.New("ContrailConfigNode get failed ")
	}

	obj, ok := objMap.(*sync.Map).Load(request.GetID())
	if !ok {
		return nil, errors.New("ContrailConfigNode get failed ")
	}

	response := &services.GetContrailConfigNodeResponse{
		ContrailConfigNode: obj.(*models.ContrailConfigNode),
	}
	return response, nil
}
