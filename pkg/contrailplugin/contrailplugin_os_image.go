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

// OsImageIntent
//   A struct to store attributes related to OsImage
//   needed by Intent Compiler
type OsImageIntent struct {
	Uuid string
}

// EvaluateOsImage - evaluates the OsImage
func EvaluateOsImage(obj interface{}) {
	resourceObj := obj.(OsImageIntent)
	log.Println("EvaluateOsImage Called ", resourceObj)
}

// CreateOsImage handles create request
func (service *PluginService) CreateOsImage(ctx context.Context, request *services.CreateOsImageRequest) (*services.CreateOsImageResponse, error) {
	log.Println(" CreateOsImage Entered")

	obj := request.GetOsImage()

	intentObj := OsImageIntent{
		Uuid: obj.GetUUID(),
	}

	if _, ok := compilationif.ObjsCache.Load("OsImageIntent"); !ok {
		compilationif.ObjsCache.Store("OsImageIntent", &sync.Map{})
	}
	objMap, ok := compilationif.ObjsCache.Load("OsImageIntent")
	if ok {
		objMap.(*sync.Map).Store(obj.GetUUID(), intentObj)
	}

	service.Debug("CreateOsImage", objMap.(*sync.Map))

	EvaluateDependencies(obj, "OsImage")

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().CreateOsImage(ctx, request)
}

// UpdateOsImage handles update request
func (service *PluginService) UpdateOsImage(ctx context.Context, request *services.UpdateOsImageRequest) (*services.UpdateOsImageResponse, error) {
	log.Println(" UpdateOsImage ENTERED")

	obj := request.GetOsImage()

	intentObj := OsImageIntent{
		Uuid: obj.GetUUID(),
	}

	if _, ok := compilationif.ObjsCache.Load("OsImageIntent"); !ok {
		compilationif.ObjsCache.Store("OsImageIntent", &sync.Map{})
	}

	EvaluateDependencies(obj, "OsImage")

	objMap, ok := compilationif.ObjsCache.Load("Intent")
	if ok {
		objMap.(*sync.Map).Store(obj.GetUUID(), intentObj)
	}

	service.Debug("UpdateOsImage", objMap.(*sync.Map))

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().UpdateOsImage(ctx, request)
}

// DeleteOsImage handles delete request
func (service *PluginService) DeleteOsImage(ctx context.Context, request *services.DeleteOsImageRequest) (*services.DeleteOsImageResponse, error) {
	log.Println(" DeleteOsImage ENTERED")

	objUUID := request.GetID()

	//intentObj := OsImageIntent {
	//OsImage: *obj,
	//}

	//EvaluateDependencies(intentObj, "OsImage")

	objMap, ok := compilationif.ObjsCache.Load("OsImageIntent")
	if ok {
		objMap.(*sync.Map).Delete(objUUID)
	}

	service.Debug("DeleteOsImage", objMap.(*sync.Map))

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().DeleteOsImage(ctx, request)
}

// GetOsImage handles get request
func (service *PluginService) GetOsImage(ctx context.Context, request *services.GetOsImageRequest) (*services.GetOsImageResponse, error) {
	objMap, ok := compilationif.ObjsCache.Load("OsImage")
	if !ok {
		return nil, errors.New("OsImage get failed ")
	}

	obj, ok := objMap.(*sync.Map).Load(request.GetID())
	if !ok {
		return nil, errors.New("OsImage get failed ")
	}

	response := &services.GetOsImageResponse{
		OsImage: obj.(*models.OsImage),
	}
	return response, nil
}
