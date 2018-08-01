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

// ServerIntent
//   A struct to store attributes related to Server
//   needed by Intent Compiler
type ServerIntent struct {
	Uuid string
}

// EvaluateServer - evaluates the Server
func EvaluateServer(obj interface{}) {
	resourceObj := obj.(ServerIntent)
	log.Println("EvaluateServer Called ", resourceObj)
}

// CreateServer handles create request
func (service *PluginService) CreateServer(ctx context.Context, request *services.CreateServerRequest) (*services.CreateServerResponse, error) {
	log.Println(" CreateServer Entered")

	obj := request.GetServer()

	intentObj := ServerIntent{
		Uuid: obj.GetUUID(),
	}

	if _, ok := compilationif.ObjsCache.Load("ServerIntent"); !ok {
		compilationif.ObjsCache.Store("ServerIntent", &sync.Map{})
	}
	objMap, ok := compilationif.ObjsCache.Load("ServerIntent")
	if ok {
		objMap.(*sync.Map).Store(obj.GetUUID(), intentObj)
	}

	service.Debug("CreateServer", objMap.(*sync.Map))

	EvaluateDependencies(obj, "Server")

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().CreateServer(ctx, request)
}

// UpdateServer handles update request
func (service *PluginService) UpdateServer(ctx context.Context, request *services.UpdateServerRequest) (*services.UpdateServerResponse, error) {
	log.Println(" UpdateServer ENTERED")

	obj := request.GetServer()

	intentObj := ServerIntent{
		Uuid: obj.GetUUID(),
	}

	if _, ok := compilationif.ObjsCache.Load("ServerIntent"); !ok {
		compilationif.ObjsCache.Store("ServerIntent", &sync.Map{})
	}

	EvaluateDependencies(obj, "Server")

	objMap, ok := compilationif.ObjsCache.Load("Intent")
	if ok {
		objMap.(*sync.Map).Store(obj.GetUUID(), intentObj)
	}

	service.Debug("UpdateServer", objMap.(*sync.Map))

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().UpdateServer(ctx, request)
}

// DeleteServer handles delete request
func (service *PluginService) DeleteServer(ctx context.Context, request *services.DeleteServerRequest) (*services.DeleteServerResponse, error) {
	log.Println(" DeleteServer ENTERED")

	objUUID := request.GetID()

	//intentObj := ServerIntent {
	//Server: *obj,
	//}

	//EvaluateDependencies(intentObj, "Server")

	objMap, ok := compilationif.ObjsCache.Load("ServerIntent")
	if ok {
		objMap.(*sync.Map).Delete(objUUID)
	}

	service.Debug("DeleteServer", objMap.(*sync.Map))

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().DeleteServer(ctx, request)
}

// GetServer handles get request
func (service *PluginService) GetServer(ctx context.Context, request *services.GetServerRequest) (*services.GetServerResponse, error) {
	objMap, ok := compilationif.ObjsCache.Load("Server")
	if !ok {
		return nil, errors.New("Server get failed ")
	}

	obj, ok := objMap.(*sync.Map).Load(request.GetID())
	if !ok {
		return nil, errors.New("Server get failed ")
	}

	response := &services.GetServerResponse{
		Server: obj.(*models.Server),
	}
	return response, nil
}
