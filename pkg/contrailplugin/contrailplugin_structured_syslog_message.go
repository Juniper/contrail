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

// StructuredSyslogMessageIntent
//   A struct to store attributes related to StructuredSyslogMessage
//   needed by Intent Compiler
type StructuredSyslogMessageIntent struct {
	Uuid string
}

// EvaluateStructuredSyslogMessage - evaluates the StructuredSyslogMessage
func EvaluateStructuredSyslogMessage(obj interface{}) {
	resourceObj := obj.(StructuredSyslogMessageIntent)
	log.Println("EvaluateStructuredSyslogMessage Called ", resourceObj)
}

// CreateStructuredSyslogMessage handles create request
func (service *PluginService) CreateStructuredSyslogMessage(ctx context.Context, request *services.CreateStructuredSyslogMessageRequest) (*services.CreateStructuredSyslogMessageResponse, error) {
	log.Println(" CreateStructuredSyslogMessage Entered")

	obj := request.GetStructuredSyslogMessage()

	intentObj := StructuredSyslogMessageIntent{
		Uuid: obj.GetUUID(),
	}

	if _, ok := compilationif.ObjsCache.Load("StructuredSyslogMessageIntent"); !ok {
		compilationif.ObjsCache.Store("StructuredSyslogMessageIntent", &sync.Map{})
	}
	objMap, ok := compilationif.ObjsCache.Load("StructuredSyslogMessageIntent")
	if ok {
		objMap.(*sync.Map).Store(obj.GetUUID(), intentObj)
	}

	service.Debug("CreateStructuredSyslogMessage", objMap.(*sync.Map))

	EvaluateDependencies(obj, "StructuredSyslogMessage")

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().CreateStructuredSyslogMessage(ctx, request)
}

// UpdateStructuredSyslogMessage handles update request
func (service *PluginService) UpdateStructuredSyslogMessage(ctx context.Context, request *services.UpdateStructuredSyslogMessageRequest) (*services.UpdateStructuredSyslogMessageResponse, error) {
	log.Println(" UpdateStructuredSyslogMessage ENTERED")

	obj := request.GetStructuredSyslogMessage()

	intentObj := StructuredSyslogMessageIntent{
		Uuid: obj.GetUUID(),
	}

	if _, ok := compilationif.ObjsCache.Load("StructuredSyslogMessageIntent"); !ok {
		compilationif.ObjsCache.Store("StructuredSyslogMessageIntent", &sync.Map{})
	}

	EvaluateDependencies(obj, "StructuredSyslogMessage")

	objMap, ok := compilationif.ObjsCache.Load("Intent")
	if ok {
		objMap.(*sync.Map).Store(obj.GetUUID(), intentObj)
	}

	service.Debug("UpdateStructuredSyslogMessage", objMap.(*sync.Map))

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().UpdateStructuredSyslogMessage(ctx, request)
}

// DeleteStructuredSyslogMessage handles delete request
func (service *PluginService) DeleteStructuredSyslogMessage(ctx context.Context, request *services.DeleteStructuredSyslogMessageRequest) (*services.DeleteStructuredSyslogMessageResponse, error) {
	log.Println(" DeleteStructuredSyslogMessage ENTERED")

	objUUID := request.GetID()

	//intentObj := StructuredSyslogMessageIntent {
	//StructuredSyslogMessage: *obj,
	//}

	//EvaluateDependencies(intentObj, "StructuredSyslogMessage")

	objMap, ok := compilationif.ObjsCache.Load("StructuredSyslogMessageIntent")
	if ok {
		objMap.(*sync.Map).Delete(objUUID)
	}

	service.Debug("DeleteStructuredSyslogMessage", objMap.(*sync.Map))

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().DeleteStructuredSyslogMessage(ctx, request)
}

// GetStructuredSyslogMessage handles get request
func (service *PluginService) GetStructuredSyslogMessage(ctx context.Context, request *services.GetStructuredSyslogMessageRequest) (*services.GetStructuredSyslogMessageResponse, error) {
	objMap, ok := compilationif.ObjsCache.Load("StructuredSyslogMessage")
	if !ok {
		return nil, errors.New("StructuredSyslogMessage get failed ")
	}

	obj, ok := objMap.(*sync.Map).Load(request.GetID())
	if !ok {
		return nil, errors.New("StructuredSyslogMessage get failed ")
	}

	response := &services.GetStructuredSyslogMessageResponse{
		StructuredSyslogMessage: obj.(*models.StructuredSyslogMessage),
	}
	return response, nil
}
