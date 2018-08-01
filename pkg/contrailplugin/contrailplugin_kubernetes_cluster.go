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

// KubernetesClusterIntent
//   A struct to store attributes related to KubernetesCluster
//   needed by Intent Compiler
type KubernetesClusterIntent struct {
	Uuid string
}

// EvaluateKubernetesCluster - evaluates the KubernetesCluster
func EvaluateKubernetesCluster(obj interface{}) {
	resourceObj := obj.(KubernetesClusterIntent)
	log.Println("EvaluateKubernetesCluster Called ", resourceObj)
}

// CreateKubernetesCluster handles create request
func (service *PluginService) CreateKubernetesCluster(ctx context.Context, request *services.CreateKubernetesClusterRequest) (*services.CreateKubernetesClusterResponse, error) {
	log.Println(" CreateKubernetesCluster Entered")

	obj := request.GetKubernetesCluster()

	intentObj := KubernetesClusterIntent{
		Uuid: obj.GetUUID(),
	}

	if _, ok := compilationif.ObjsCache.Load("KubernetesClusterIntent"); !ok {
		compilationif.ObjsCache.Store("KubernetesClusterIntent", &sync.Map{})
	}
	objMap, ok := compilationif.ObjsCache.Load("KubernetesClusterIntent")
	if ok {
		objMap.(*sync.Map).Store(obj.GetUUID(), intentObj)
	}

	service.Debug("CreateKubernetesCluster", objMap.(*sync.Map))

	EvaluateDependencies(obj, "KubernetesCluster")

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().CreateKubernetesCluster(ctx, request)
}

// UpdateKubernetesCluster handles update request
func (service *PluginService) UpdateKubernetesCluster(ctx context.Context, request *services.UpdateKubernetesClusterRequest) (*services.UpdateKubernetesClusterResponse, error) {
	log.Println(" UpdateKubernetesCluster ENTERED")

	obj := request.GetKubernetesCluster()

	intentObj := KubernetesClusterIntent{
		Uuid: obj.GetUUID(),
	}

	if _, ok := compilationif.ObjsCache.Load("KubernetesClusterIntent"); !ok {
		compilationif.ObjsCache.Store("KubernetesClusterIntent", &sync.Map{})
	}

	EvaluateDependencies(obj, "KubernetesCluster")

	objMap, ok := compilationif.ObjsCache.Load("Intent")
	if ok {
		objMap.(*sync.Map).Store(obj.GetUUID(), intentObj)
	}

	service.Debug("UpdateKubernetesCluster", objMap.(*sync.Map))

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().UpdateKubernetesCluster(ctx, request)
}

// DeleteKubernetesCluster handles delete request
func (service *PluginService) DeleteKubernetesCluster(ctx context.Context, request *services.DeleteKubernetesClusterRequest) (*services.DeleteKubernetesClusterResponse, error) {
	log.Println(" DeleteKubernetesCluster ENTERED")

	objUUID := request.GetID()

	//intentObj := KubernetesClusterIntent {
	//KubernetesCluster: *obj,
	//}

	//EvaluateDependencies(intentObj, "KubernetesCluster")

	objMap, ok := compilationif.ObjsCache.Load("KubernetesClusterIntent")
	if ok {
		objMap.(*sync.Map).Delete(objUUID)
	}

	service.Debug("DeleteKubernetesCluster", objMap.(*sync.Map))

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().DeleteKubernetesCluster(ctx, request)
}

// GetKubernetesCluster handles get request
func (service *PluginService) GetKubernetesCluster(ctx context.Context, request *services.GetKubernetesClusterRequest) (*services.GetKubernetesClusterResponse, error) {
	objMap, ok := compilationif.ObjsCache.Load("KubernetesCluster")
	if !ok {
		return nil, errors.New("KubernetesCluster get failed ")
	}

	obj, ok := objMap.(*sync.Map).Load(request.GetID())
	if !ok {
		return nil, errors.New("KubernetesCluster get failed ")
	}

	response := &services.GetKubernetesClusterResponse{
		KubernetesCluster: obj.(*models.KubernetesCluster),
	}
	return response, nil
}
