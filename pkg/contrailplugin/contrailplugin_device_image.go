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

// DeviceImageIntent
//   A struct to store attributes related to DeviceImage
//   needed by Intent Compiler
type DeviceImageIntent struct {
	Uuid string
}

// EvaluateDeviceImage - evaluates the DeviceImage
func EvaluateDeviceImage(obj interface{}) {
	resourceObj := obj.(DeviceImageIntent)
	log.Println("EvaluateDeviceImage Called ", resourceObj)
}

// CreateDeviceImage handles create request
func (service *PluginService) CreateDeviceImage(ctx context.Context, request *services.CreateDeviceImageRequest) (*services.CreateDeviceImageResponse, error) {
	log.Println(" CreateDeviceImage Entered")

	obj := request.GetDeviceImage()

	intentObj := DeviceImageIntent{
		Uuid: obj.GetUUID(),
	}

	if _, ok := compilationif.ObjsCache.Load("DeviceImageIntent"); !ok {
		compilationif.ObjsCache.Store("DeviceImageIntent", &sync.Map{})
	}
	objMap, ok := compilationif.ObjsCache.Load("DeviceImageIntent")
	if ok {
		objMap.(*sync.Map).Store(obj.GetUUID(), intentObj)
	}

	service.Debug("CreateDeviceImage", objMap.(*sync.Map))

	EvaluateDependencies(obj, "DeviceImage")

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().CreateDeviceImage(ctx, request)
}

// UpdateDeviceImage handles update request
func (service *PluginService) UpdateDeviceImage(ctx context.Context, request *services.UpdateDeviceImageRequest) (*services.UpdateDeviceImageResponse, error) {
	log.Println(" UpdateDeviceImage ENTERED")

	obj := request.GetDeviceImage()

	intentObj := DeviceImageIntent{
		Uuid: obj.GetUUID(),
	}

	if _, ok := compilationif.ObjsCache.Load("DeviceImageIntent"); !ok {
		compilationif.ObjsCache.Store("DeviceImageIntent", &sync.Map{})
	}

	EvaluateDependencies(obj, "DeviceImage")

	objMap, ok := compilationif.ObjsCache.Load("Intent")
	if ok {
		objMap.(*sync.Map).Store(obj.GetUUID(), intentObj)
	}

	service.Debug("UpdateDeviceImage", objMap.(*sync.Map))

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().UpdateDeviceImage(ctx, request)
}

// DeleteDeviceImage handles delete request
func (service *PluginService) DeleteDeviceImage(ctx context.Context, request *services.DeleteDeviceImageRequest) (*services.DeleteDeviceImageResponse, error) {
	log.Println(" DeleteDeviceImage ENTERED")

	objUUID := request.GetID()

	//intentObj := DeviceImageIntent {
	//DeviceImage: *obj,
	//}

	//EvaluateDependencies(intentObj, "DeviceImage")

	objMap, ok := compilationif.ObjsCache.Load("DeviceImageIntent")
	if ok {
		objMap.(*sync.Map).Delete(objUUID)
	}

	service.Debug("DeleteDeviceImage", objMap.(*sync.Map))

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().DeleteDeviceImage(ctx, request)
}

// GetDeviceImage handles get request
func (service *PluginService) GetDeviceImage(ctx context.Context, request *services.GetDeviceImageRequest) (*services.GetDeviceImageResponse, error) {
	objMap, ok := compilationif.ObjsCache.Load("DeviceImage")
	if !ok {
		return nil, errors.New("DeviceImage get failed ")
	}

	obj, ok := objMap.(*sync.Map).Load(request.GetID())
	if !ok {
		return nil, errors.New("DeviceImage get failed ")
	}

	response := &services.GetDeviceImageResponse{
		DeviceImage: obj.(*models.DeviceImage),
	}
	return response, nil
}
