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

// KubernetesMasterNodeIntent
//   A struct to store attributes related to KubernetesMasterNode
//   needed by Intent Compiler
type KubernetesMasterNodeIntent struct {
	Uuid string
}

// EvaluateKubernetesMasterNode - evaluates the KubernetesMasterNode
func EvaluateKubernetesMasterNode(obj interface{}) {
	resourceObj := obj.(KubernetesMasterNodeIntent)
	log.Println("EvaluateKubernetesMasterNode Called ", resourceObj)
}

// CreateKubernetesMasterNode handles create request
func (service *PluginService) CreateKubernetesMasterNode(ctx context.Context, request *services.CreateKubernetesMasterNodeRequest) (*services.CreateKubernetesMasterNodeResponse, error) {
	log.Println(" CreateKubernetesMasterNode Entered")

	obj := request.GetKubernetesMasterNode()

	intentObj := KubernetesMasterNodeIntent{
		Uuid: obj.GetUUID(),
	}

	if _, ok := compilationif.ObjsCache.Load("KubernetesMasterNodeIntent"); !ok {
		compilationif.ObjsCache.Store("KubernetesMasterNodeIntent", &sync.Map{})
	}
	objMap, ok := compilationif.ObjsCache.Load("KubernetesMasterNodeIntent")
	if ok {
		objMap.(*sync.Map).Store(obj.GetUUID(), intentObj)
	}

	service.Debug("CreateKubernetesMasterNode", objMap.(*sync.Map))

	EvaluateDependencies(obj, "KubernetesMasterNode")

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().CreateKubernetesMasterNode(ctx, request)
}

// UpdateKubernetesMasterNode handles update request
func (service *PluginService) UpdateKubernetesMasterNode(ctx context.Context, request *services.UpdateKubernetesMasterNodeRequest) (*services.UpdateKubernetesMasterNodeResponse, error) {
	log.Println(" UpdateKubernetesMasterNode ENTERED")

	obj := request.GetKubernetesMasterNode()

	intentObj := KubernetesMasterNodeIntent{
		Uuid: obj.GetUUID(),
	}

	if _, ok := compilationif.ObjsCache.Load("KubernetesMasterNodeIntent"); !ok {
		compilationif.ObjsCache.Store("KubernetesMasterNodeIntent", &sync.Map{})
	}

	EvaluateDependencies(obj, "KubernetesMasterNode")

	objMap, ok := compilationif.ObjsCache.Load("Intent")
	if ok {
		objMap.(*sync.Map).Store(obj.GetUUID(), intentObj)
	}

	service.Debug("UpdateKubernetesMasterNode", objMap.(*sync.Map))

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().UpdateKubernetesMasterNode(ctx, request)
}

// DeleteKubernetesMasterNode handles delete request
func (service *PluginService) DeleteKubernetesMasterNode(ctx context.Context, request *services.DeleteKubernetesMasterNodeRequest) (*services.DeleteKubernetesMasterNodeResponse, error) {
	log.Println(" DeleteKubernetesMasterNode ENTERED")

	objUUID := request.GetID()

	//intentObj := KubernetesMasterNodeIntent {
	//KubernetesMasterNode: *obj,
	//}

	//EvaluateDependencies(intentObj, "KubernetesMasterNode")

	objMap, ok := compilationif.ObjsCache.Load("KubernetesMasterNodeIntent")
	if ok {
		objMap.(*sync.Map).Delete(objUUID)
	}

	service.Debug("DeleteKubernetesMasterNode", objMap.(*sync.Map))

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().DeleteKubernetesMasterNode(ctx, request)
}

// GetKubernetesMasterNode handles get request
func (service *PluginService) GetKubernetesMasterNode(ctx context.Context, request *services.GetKubernetesMasterNodeRequest) (*services.GetKubernetesMasterNodeResponse, error) {
	objMap, ok := compilationif.ObjsCache.Load("KubernetesMasterNode")
	if !ok {
		return nil, errors.New("KubernetesMasterNode get failed ")
	}

	obj, ok := objMap.(*sync.Map).Load(request.GetID())
	if !ok {
		return nil, errors.New("KubernetesMasterNode get failed ")
	}

	response := &services.GetKubernetesMasterNodeResponse{
		KubernetesMasterNode: obj.(*models.KubernetesMasterNode),
	}
	return response, nil
}
