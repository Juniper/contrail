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

// CustomerAttachmentIntent
//   A struct to store attributes related to CustomerAttachment
//   needed by Intent Compiler
type CustomerAttachmentIntent struct {
	Uuid string
}

// EvaluateCustomerAttachment - evaluates the CustomerAttachment
func EvaluateCustomerAttachment(obj interface{}) {
	resourceObj := obj.(CustomerAttachmentIntent)
	log.Println("EvaluateCustomerAttachment Called ", resourceObj)
}

// CreateCustomerAttachment handles create request
func (service *PluginService) CreateCustomerAttachment(ctx context.Context, request *services.CreateCustomerAttachmentRequest) (*services.CreateCustomerAttachmentResponse, error) {
	log.Println(" CreateCustomerAttachment Entered")

	obj := request.GetCustomerAttachment()

	intentObj := CustomerAttachmentIntent{
		Uuid: obj.GetUUID(),
	}

	if _, ok := compilationif.ObjsCache.Load("CustomerAttachmentIntent"); !ok {
		compilationif.ObjsCache.Store("CustomerAttachmentIntent", &sync.Map{})
	}
	objMap, ok := compilationif.ObjsCache.Load("CustomerAttachmentIntent")
	if ok {
		objMap.(*sync.Map).Store(obj.GetUUID(), intentObj)
	}

	service.Debug("CreateCustomerAttachment", objMap.(*sync.Map))

	EvaluateDependencies(obj, "CustomerAttachment")

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().CreateCustomerAttachment(ctx, request)
}

// UpdateCustomerAttachment handles update request
func (service *PluginService) UpdateCustomerAttachment(ctx context.Context, request *services.UpdateCustomerAttachmentRequest) (*services.UpdateCustomerAttachmentResponse, error) {
	log.Println(" UpdateCustomerAttachment ENTERED")

	obj := request.GetCustomerAttachment()

	intentObj := CustomerAttachmentIntent{
		Uuid: obj.GetUUID(),
	}

	if _, ok := compilationif.ObjsCache.Load("CustomerAttachmentIntent"); !ok {
		compilationif.ObjsCache.Store("CustomerAttachmentIntent", &sync.Map{})
	}

	EvaluateDependencies(obj, "CustomerAttachment")

	objMap, ok := compilationif.ObjsCache.Load("Intent")
	if ok {
		objMap.(*sync.Map).Store(obj.GetUUID(), intentObj)
	}

	service.Debug("UpdateCustomerAttachment", objMap.(*sync.Map))

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().UpdateCustomerAttachment(ctx, request)
}

// DeleteCustomerAttachment handles delete request
func (service *PluginService) DeleteCustomerAttachment(ctx context.Context, request *services.DeleteCustomerAttachmentRequest) (*services.DeleteCustomerAttachmentResponse, error) {
	log.Println(" DeleteCustomerAttachment ENTERED")

	objUUID := request.GetID()

	//intentObj := CustomerAttachmentIntent {
	//CustomerAttachment: *obj,
	//}

	//EvaluateDependencies(intentObj, "CustomerAttachment")

	objMap, ok := compilationif.ObjsCache.Load("CustomerAttachmentIntent")
	if ok {
		objMap.(*sync.Map).Delete(objUUID)
	}

	service.Debug("DeleteCustomerAttachment", objMap.(*sync.Map))

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().DeleteCustomerAttachment(ctx, request)
}

// GetCustomerAttachment handles get request
func (service *PluginService) GetCustomerAttachment(ctx context.Context, request *services.GetCustomerAttachmentRequest) (*services.GetCustomerAttachmentResponse, error) {
	objMap, ok := compilationif.ObjsCache.Load("CustomerAttachment")
	if !ok {
		return nil, errors.New("CustomerAttachment get failed ")
	}

	obj, ok := objMap.(*sync.Map).Load(request.GetID())
	if !ok {
		return nil, errors.New("CustomerAttachment get failed ")
	}

	response := &services.GetCustomerAttachmentResponse{
		CustomerAttachment: obj.(*models.CustomerAttachment),
	}
	return response, nil
}
