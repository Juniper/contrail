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

// SubClusterIntent
//   A struct to store attributes related to SubCluster
//   needed by Intent Compiler
type SubClusterIntent struct {
	Uuid string
}

// EvaluateSubCluster - evaluates the SubCluster
func EvaluateSubCluster(obj interface{}) {
	resourceObj := obj.(SubClusterIntent)
	log.Println("EvaluateSubCluster Called ", resourceObj)
}

// CreateSubCluster handles create request
func (service *PluginService) CreateSubCluster(ctx context.Context, request *services.CreateSubClusterRequest) (*services.CreateSubClusterResponse, error) {
	log.Println(" CreateSubCluster Entered")

	obj := request.GetSubCluster()

	intentObj := SubClusterIntent{
		Uuid: obj.GetUUID(),
	}

	if _, ok := compilationif.ObjsCache.Load("SubClusterIntent"); !ok {
		compilationif.ObjsCache.Store("SubClusterIntent", &sync.Map{})
	}
	objMap, ok := compilationif.ObjsCache.Load("SubClusterIntent")
	if ok {
		objMap.(*sync.Map).Store(obj.GetUUID(), intentObj)
	}

	service.Debug("CreateSubCluster", objMap.(*sync.Map))

	EvaluateDependencies(obj, "SubCluster")

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().CreateSubCluster(ctx, request)
}

// UpdateSubCluster handles update request
func (service *PluginService) UpdateSubCluster(ctx context.Context, request *services.UpdateSubClusterRequest) (*services.UpdateSubClusterResponse, error) {
	log.Println(" UpdateSubCluster ENTERED")

	obj := request.GetSubCluster()

	intentObj := SubClusterIntent{
		Uuid: obj.GetUUID(),
	}

	if _, ok := compilationif.ObjsCache.Load("SubClusterIntent"); !ok {
		compilationif.ObjsCache.Store("SubClusterIntent", &sync.Map{})
	}

	EvaluateDependencies(obj, "SubCluster")

	objMap, ok := compilationif.ObjsCache.Load("Intent")
	if ok {
		objMap.(*sync.Map).Store(obj.GetUUID(), intentObj)
	}

	service.Debug("UpdateSubCluster", objMap.(*sync.Map))

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().UpdateSubCluster(ctx, request)
}

// DeleteSubCluster handles delete request
func (service *PluginService) DeleteSubCluster(ctx context.Context, request *services.DeleteSubClusterRequest) (*services.DeleteSubClusterResponse, error) {
	log.Println(" DeleteSubCluster ENTERED")

	objUUID := request.GetID()

	//intentObj := SubClusterIntent {
	//SubCluster: *obj,
	//}

	//EvaluateDependencies(intentObj, "SubCluster")

	objMap, ok := compilationif.ObjsCache.Load("SubClusterIntent")
	if ok {
		objMap.(*sync.Map).Delete(objUUID)
	}

	service.Debug("DeleteSubCluster", objMap.(*sync.Map))

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().DeleteSubCluster(ctx, request)
}

// GetSubCluster handles get request
func (service *PluginService) GetSubCluster(ctx context.Context, request *services.GetSubClusterRequest) (*services.GetSubClusterResponse, error) {
	objMap, ok := compilationif.ObjsCache.Load("SubCluster")
	if !ok {
		return nil, errors.New("SubCluster get failed ")
	}

	obj, ok := objMap.(*sync.Map).Load(request.GetID())
	if !ok {
		return nil, errors.New("SubCluster get failed ")
	}

	response := &services.GetSubClusterResponse{
		SubCluster: obj.(*models.SubCluster),
	}
	return response, nil
}
