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

// FloatingIPIntent
//   A struct to store attributes related to FloatingIP
//   needed by Intent Compiler
type FloatingIPIntent struct {
	Uuid string
}

// EvaluateFloatingIP - evaluates the FloatingIP
func EvaluateFloatingIP(obj interface{}) {
	resourceObj := obj.(FloatingIPIntent)
	log.Println("EvaluateFloatingIP Called ", resourceObj)
}

// CreateFloatingIP handles create request
func (service *PluginService) CreateFloatingIP(ctx context.Context, request *services.CreateFloatingIPRequest) (*services.CreateFloatingIPResponse, error) {
	log.Println(" CreateFloatingIP Entered")

	obj := request.GetFloatingIP()

	intentObj := FloatingIPIntent{
		Uuid: obj.GetUUID(),
	}

	if _, ok := compilationif.ObjsCache.Load("FloatingIPIntent"); !ok {
		compilationif.ObjsCache.Store("FloatingIPIntent", &sync.Map{})
	}
	objMap, ok := compilationif.ObjsCache.Load("FloatingIPIntent")
	if ok {
		objMap.(*sync.Map).Store(obj.GetUUID(), intentObj)
	}

	service.Debug("CreateFloatingIP", objMap.(*sync.Map))

	EvaluateDependencies(obj, "FloatingIP")

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().CreateFloatingIP(ctx, request)
}

// UpdateFloatingIP handles update request
func (service *PluginService) UpdateFloatingIP(ctx context.Context, request *services.UpdateFloatingIPRequest) (*services.UpdateFloatingIPResponse, error) {
	log.Println(" UpdateFloatingIP ENTERED")

	obj := request.GetFloatingIP()

	intentObj := FloatingIPIntent{
		Uuid: obj.GetUUID(),
	}

	if _, ok := compilationif.ObjsCache.Load("FloatingIPIntent"); !ok {
		compilationif.ObjsCache.Store("FloatingIPIntent", &sync.Map{})
	}

	EvaluateDependencies(obj, "FloatingIP")

	objMap, ok := compilationif.ObjsCache.Load("Intent")
	if ok {
		objMap.(*sync.Map).Store(obj.GetUUID(), intentObj)
	}

	service.Debug("UpdateFloatingIP", objMap.(*sync.Map))

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().UpdateFloatingIP(ctx, request)
}

// DeleteFloatingIP handles delete request
func (service *PluginService) DeleteFloatingIP(ctx context.Context, request *services.DeleteFloatingIPRequest) (*services.DeleteFloatingIPResponse, error) {
	log.Println(" DeleteFloatingIP ENTERED")

	objUUID := request.GetID()

	//intentObj := FloatingIPIntent {
	//FloatingIP: *obj,
	//}

	//EvaluateDependencies(intentObj, "FloatingIP")

	objMap, ok := compilationif.ObjsCache.Load("FloatingIPIntent")
	if ok {
		objMap.(*sync.Map).Delete(objUUID)
	}

	service.Debug("DeleteFloatingIP", objMap.(*sync.Map))

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().DeleteFloatingIP(ctx, request)
}

// GetFloatingIP handles get request
func (service *PluginService) GetFloatingIP(ctx context.Context, request *services.GetFloatingIPRequest) (*services.GetFloatingIPResponse, error) {
	objMap, ok := compilationif.ObjsCache.Load("FloatingIP")
	if !ok {
		return nil, errors.New("FloatingIP get failed ")
	}

	obj, ok := objMap.(*sync.Map).Load(request.GetID())
	if !ok {
		return nil, errors.New("FloatingIP get failed ")
	}

	response := &services.GetFloatingIPResponse{
		FloatingIP: obj.(*models.FloatingIP),
	}
	return response, nil
}
