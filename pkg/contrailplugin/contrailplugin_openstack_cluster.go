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

// OpenstackClusterIntent
//   A struct to store attributes related to OpenstackCluster
//   needed by Intent Compiler
type OpenstackClusterIntent struct {
	Uuid string
}

// EvaluateOpenstackCluster - evaluates the OpenstackCluster
func EvaluateOpenstackCluster(obj interface{}) {
	resourceObj := obj.(OpenstackClusterIntent)
	log.Println("EvaluateOpenstackCluster Called ", resourceObj)
}

// CreateOpenstackCluster handles create request
func (service *PluginService) CreateOpenstackCluster(ctx context.Context, request *services.CreateOpenstackClusterRequest) (*services.CreateOpenstackClusterResponse, error) {
	log.Println(" CreateOpenstackCluster Entered")

	obj := request.GetOpenstackCluster()

	intentObj := OpenstackClusterIntent{
		Uuid: obj.GetUUID(),
	}

	if _, ok := compilationif.ObjsCache.Load("OpenstackClusterIntent"); !ok {
		compilationif.ObjsCache.Store("OpenstackClusterIntent", &sync.Map{})
	}
	objMap, ok := compilationif.ObjsCache.Load("OpenstackClusterIntent")
	if ok {
		objMap.(*sync.Map).Store(obj.GetUUID(), intentObj)
	}

	service.Debug("CreateOpenstackCluster", objMap.(*sync.Map))

	EvaluateDependencies(obj, "OpenstackCluster")

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().CreateOpenstackCluster(ctx, request)
}

// UpdateOpenstackCluster handles update request
func (service *PluginService) UpdateOpenstackCluster(ctx context.Context, request *services.UpdateOpenstackClusterRequest) (*services.UpdateOpenstackClusterResponse, error) {
	log.Println(" UpdateOpenstackCluster ENTERED")

	obj := request.GetOpenstackCluster()

	intentObj := OpenstackClusterIntent{
		Uuid: obj.GetUUID(),
	}

	if _, ok := compilationif.ObjsCache.Load("OpenstackClusterIntent"); !ok {
		compilationif.ObjsCache.Store("OpenstackClusterIntent", &sync.Map{})
	}

	EvaluateDependencies(obj, "OpenstackCluster")

	objMap, ok := compilationif.ObjsCache.Load("Intent")
	if ok {
		objMap.(*sync.Map).Store(obj.GetUUID(), intentObj)
	}

	service.Debug("UpdateOpenstackCluster", objMap.(*sync.Map))

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().UpdateOpenstackCluster(ctx, request)
}

// DeleteOpenstackCluster handles delete request
func (service *PluginService) DeleteOpenstackCluster(ctx context.Context, request *services.DeleteOpenstackClusterRequest) (*services.DeleteOpenstackClusterResponse, error) {
	log.Println(" DeleteOpenstackCluster ENTERED")

	objUUID := request.GetID()

	//intentObj := OpenstackClusterIntent {
	//OpenstackCluster: *obj,
	//}

	//EvaluateDependencies(intentObj, "OpenstackCluster")

	objMap, ok := compilationif.ObjsCache.Load("OpenstackClusterIntent")
	if ok {
		objMap.(*sync.Map).Delete(objUUID)
	}

	service.Debug("DeleteOpenstackCluster", objMap.(*sync.Map))

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().DeleteOpenstackCluster(ctx, request)
}

// GetOpenstackCluster handles get request
func (service *PluginService) GetOpenstackCluster(ctx context.Context, request *services.GetOpenstackClusterRequest) (*services.GetOpenstackClusterResponse, error) {
	objMap, ok := compilationif.ObjsCache.Load("OpenstackCluster")
	if !ok {
		return nil, errors.New("OpenstackCluster get failed ")
	}

	obj, ok := objMap.(*sync.Map).Load(request.GetID())
	if !ok {
		return nil, errors.New("OpenstackCluster get failed ")
	}

	response := &services.GetOpenstackClusterResponse{
		OpenstackCluster: obj.(*models.OpenstackCluster),
	}
	return response, nil
}
