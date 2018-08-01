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

// ProviderAttachmentIntent
//   A struct to store attributes related to ProviderAttachment
//   needed by Intent Compiler
type ProviderAttachmentIntent struct {
	Uuid string
}

// EvaluateProviderAttachment - evaluates the ProviderAttachment
func EvaluateProviderAttachment(obj interface{}) {
	resourceObj := obj.(ProviderAttachmentIntent)
	log.Println("EvaluateProviderAttachment Called ", resourceObj)
}

// CreateProviderAttachment handles create request
func (service *PluginService) CreateProviderAttachment(ctx context.Context, request *services.CreateProviderAttachmentRequest) (*services.CreateProviderAttachmentResponse, error) {
	log.Println(" CreateProviderAttachment Entered")

	obj := request.GetProviderAttachment()

	intentObj := ProviderAttachmentIntent{
		Uuid: obj.GetUUID(),
	}

	if _, ok := compilationif.ObjsCache.Load("ProviderAttachmentIntent"); !ok {
		compilationif.ObjsCache.Store("ProviderAttachmentIntent", &sync.Map{})
	}
	objMap, ok := compilationif.ObjsCache.Load("ProviderAttachmentIntent")
	if ok {
		objMap.(*sync.Map).Store(obj.GetUUID(), intentObj)
	}

	service.Debug("CreateProviderAttachment", objMap.(*sync.Map))

	EvaluateDependencies(obj, "ProviderAttachment")

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().CreateProviderAttachment(ctx, request)
}

// UpdateProviderAttachment handles update request
func (service *PluginService) UpdateProviderAttachment(ctx context.Context, request *services.UpdateProviderAttachmentRequest) (*services.UpdateProviderAttachmentResponse, error) {
	log.Println(" UpdateProviderAttachment ENTERED")

	obj := request.GetProviderAttachment()

	intentObj := ProviderAttachmentIntent{
		Uuid: obj.GetUUID(),
	}

	if _, ok := compilationif.ObjsCache.Load("ProviderAttachmentIntent"); !ok {
		compilationif.ObjsCache.Store("ProviderAttachmentIntent", &sync.Map{})
	}

	EvaluateDependencies(obj, "ProviderAttachment")

	objMap, ok := compilationif.ObjsCache.Load("Intent")
	if ok {
		objMap.(*sync.Map).Store(obj.GetUUID(), intentObj)
	}

	service.Debug("UpdateProviderAttachment", objMap.(*sync.Map))

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().UpdateProviderAttachment(ctx, request)
}

// DeleteProviderAttachment handles delete request
func (service *PluginService) DeleteProviderAttachment(ctx context.Context, request *services.DeleteProviderAttachmentRequest) (*services.DeleteProviderAttachmentResponse, error) {
	log.Println(" DeleteProviderAttachment ENTERED")

	objUUID := request.GetID()

	//intentObj := ProviderAttachmentIntent {
	//ProviderAttachment: *obj,
	//}

	//EvaluateDependencies(intentObj, "ProviderAttachment")

	objMap, ok := compilationif.ObjsCache.Load("ProviderAttachmentIntent")
	if ok {
		objMap.(*sync.Map).Delete(objUUID)
	}

	service.Debug("DeleteProviderAttachment", objMap.(*sync.Map))

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().DeleteProviderAttachment(ctx, request)
}

// GetProviderAttachment handles get request
func (service *PluginService) GetProviderAttachment(ctx context.Context, request *services.GetProviderAttachmentRequest) (*services.GetProviderAttachmentResponse, error) {
	objMap, ok := compilationif.ObjsCache.Load("ProviderAttachment")
	if !ok {
		return nil, errors.New("ProviderAttachment get failed ")
	}

	obj, ok := objMap.(*sync.Map).Load(request.GetID())
	if !ok {
		return nil, errors.New("ProviderAttachment get failed ")
	}

	response := &services.GetProviderAttachmentResponse{
		ProviderAttachment: obj.(*models.ProviderAttachment),
	}
	return response, nil
}
