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

// WidgetIntent
//   A struct to store attributes related to Widget
//   needed by Intent Compiler
type WidgetIntent struct {
	Uuid string
}

// EvaluateWidget - evaluates the Widget
func EvaluateWidget(obj interface{}) {
	resourceObj := obj.(WidgetIntent)
	log.Println("EvaluateWidget Called ", resourceObj)
}

// CreateWidget handles create request
func (service *PluginService) CreateWidget(ctx context.Context, request *services.CreateWidgetRequest) (*services.CreateWidgetResponse, error) {
	log.Println(" CreateWidget Entered")

	obj := request.GetWidget()

	intentObj := WidgetIntent{
		Uuid: obj.GetUUID(),
	}

	if _, ok := compilationif.ObjsCache.Load("WidgetIntent"); !ok {
		compilationif.ObjsCache.Store("WidgetIntent", &sync.Map{})
	}
	objMap, ok := compilationif.ObjsCache.Load("WidgetIntent")
	if ok {
		objMap.(*sync.Map).Store(obj.GetUUID(), intentObj)
	}

	service.Debug("CreateWidget", objMap.(*sync.Map))

	EvaluateDependencies(obj, "Widget")

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().CreateWidget(ctx, request)
}

// UpdateWidget handles update request
func (service *PluginService) UpdateWidget(ctx context.Context, request *services.UpdateWidgetRequest) (*services.UpdateWidgetResponse, error) {
	log.Println(" UpdateWidget ENTERED")

	obj := request.GetWidget()

	intentObj := WidgetIntent{
		Uuid: obj.GetUUID(),
	}

	if _, ok := compilationif.ObjsCache.Load("WidgetIntent"); !ok {
		compilationif.ObjsCache.Store("WidgetIntent", &sync.Map{})
	}

	EvaluateDependencies(obj, "Widget")

	objMap, ok := compilationif.ObjsCache.Load("Intent")
	if ok {
		objMap.(*sync.Map).Store(obj.GetUUID(), intentObj)
	}

	service.Debug("UpdateWidget", objMap.(*sync.Map))

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().UpdateWidget(ctx, request)
}

// DeleteWidget handles delete request
func (service *PluginService) DeleteWidget(ctx context.Context, request *services.DeleteWidgetRequest) (*services.DeleteWidgetResponse, error) {
	log.Println(" DeleteWidget ENTERED")

	objUUID := request.GetID()

	//intentObj := WidgetIntent {
	//Widget: *obj,
	//}

	//EvaluateDependencies(intentObj, "Widget")

	objMap, ok := compilationif.ObjsCache.Load("WidgetIntent")
	if ok {
		objMap.(*sync.Map).Delete(objUUID)
	}

	service.Debug("DeleteWidget", objMap.(*sync.Map))

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().DeleteWidget(ctx, request)
}

// GetWidget handles get request
func (service *PluginService) GetWidget(ctx context.Context, request *services.GetWidgetRequest) (*services.GetWidgetResponse, error) {
	objMap, ok := compilationif.ObjsCache.Load("Widget")
	if !ok {
		return nil, errors.New("Widget get failed ")
	}

	obj, ok := objMap.(*sync.Map).Load(request.GetID())
	if !ok {
		return nil, errors.New("Widget get failed ")
	}

	response := &services.GetWidgetResponse{
		Widget: obj.(*models.Widget),
	}
	return response, nil
}
