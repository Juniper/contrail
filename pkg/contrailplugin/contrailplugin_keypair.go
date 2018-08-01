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

// KeypairIntent
//   A struct to store attributes related to Keypair
//   needed by Intent Compiler
type KeypairIntent struct {
	Uuid string
}

// EvaluateKeypair - evaluates the Keypair
func EvaluateKeypair(obj interface{}) {
	resourceObj := obj.(KeypairIntent)
	log.Println("EvaluateKeypair Called ", resourceObj)
}

// CreateKeypair handles create request
func (service *PluginService) CreateKeypair(ctx context.Context, request *services.CreateKeypairRequest) (*services.CreateKeypairResponse, error) {
	log.Println(" CreateKeypair Entered")

	obj := request.GetKeypair()

	intentObj := KeypairIntent{
		Uuid: obj.GetUUID(),
	}

	if _, ok := compilationif.ObjsCache.Load("KeypairIntent"); !ok {
		compilationif.ObjsCache.Store("KeypairIntent", &sync.Map{})
	}
	objMap, ok := compilationif.ObjsCache.Load("KeypairIntent")
	if ok {
		objMap.(*sync.Map).Store(obj.GetUUID(), intentObj)
	}

	service.Debug("CreateKeypair", objMap.(*sync.Map))

	EvaluateDependencies(obj, "Keypair")

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().CreateKeypair(ctx, request)
}

// UpdateKeypair handles update request
func (service *PluginService) UpdateKeypair(ctx context.Context, request *services.UpdateKeypairRequest) (*services.UpdateKeypairResponse, error) {
	log.Println(" UpdateKeypair ENTERED")

	obj := request.GetKeypair()

	intentObj := KeypairIntent{
		Uuid: obj.GetUUID(),
	}

	if _, ok := compilationif.ObjsCache.Load("KeypairIntent"); !ok {
		compilationif.ObjsCache.Store("KeypairIntent", &sync.Map{})
	}

	EvaluateDependencies(obj, "Keypair")

	objMap, ok := compilationif.ObjsCache.Load("Intent")
	if ok {
		objMap.(*sync.Map).Store(obj.GetUUID(), intentObj)
	}

	service.Debug("UpdateKeypair", objMap.(*sync.Map))

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().UpdateKeypair(ctx, request)
}

// DeleteKeypair handles delete request
func (service *PluginService) DeleteKeypair(ctx context.Context, request *services.DeleteKeypairRequest) (*services.DeleteKeypairResponse, error) {
	log.Println(" DeleteKeypair ENTERED")

	objUUID := request.GetID()

	//intentObj := KeypairIntent {
	//Keypair: *obj,
	//}

	//EvaluateDependencies(intentObj, "Keypair")

	objMap, ok := compilationif.ObjsCache.Load("KeypairIntent")
	if ok {
		objMap.(*sync.Map).Delete(objUUID)
	}

	service.Debug("DeleteKeypair", objMap.(*sync.Map))

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().DeleteKeypair(ctx, request)
}

// GetKeypair handles get request
func (service *PluginService) GetKeypair(ctx context.Context, request *services.GetKeypairRequest) (*services.GetKeypairResponse, error) {
	objMap, ok := compilationif.ObjsCache.Load("Keypair")
	if !ok {
		return nil, errors.New("Keypair get failed ")
	}

	obj, ok := objMap.(*sync.Map).Load(request.GetID())
	if !ok {
		return nil, errors.New("Keypair get failed ")
	}

	response := &services.GetKeypairResponse{
		Keypair: obj.(*models.Keypair),
	}
	return response, nil
}
