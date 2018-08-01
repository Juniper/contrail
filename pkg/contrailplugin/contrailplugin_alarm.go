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

// AlarmIntent
//   A struct to store attributes related to Alarm
//   needed by Intent Compiler
type AlarmIntent struct {
	Uuid string
}

// EvaluateAlarm - evaluates the Alarm
func EvaluateAlarm(obj interface{}) {
	resourceObj := obj.(AlarmIntent)
	log.Println("EvaluateAlarm Called ", resourceObj)
}

// CreateAlarm handles create request
func (service *PluginService) CreateAlarm(ctx context.Context, request *services.CreateAlarmRequest) (*services.CreateAlarmResponse, error) {
	log.Println(" CreateAlarm Entered")

	obj := request.GetAlarm()

	intentObj := AlarmIntent{
		Uuid: obj.GetUUID(),
	}

	if _, ok := compilationif.ObjsCache.Load("AlarmIntent"); !ok {
		compilationif.ObjsCache.Store("AlarmIntent", &sync.Map{})
	}
	objMap, ok := compilationif.ObjsCache.Load("AlarmIntent")
	if ok {
		objMap.(*sync.Map).Store(obj.GetUUID(), intentObj)
	}

	service.Debug("CreateAlarm", objMap.(*sync.Map))

	EvaluateDependencies(obj, "Alarm")

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().CreateAlarm(ctx, request)
}

// UpdateAlarm handles update request
func (service *PluginService) UpdateAlarm(ctx context.Context, request *services.UpdateAlarmRequest) (*services.UpdateAlarmResponse, error) {
	log.Println(" UpdateAlarm ENTERED")

	obj := request.GetAlarm()

	intentObj := AlarmIntent{
		Uuid: obj.GetUUID(),
	}

	if _, ok := compilationif.ObjsCache.Load("AlarmIntent"); !ok {
		compilationif.ObjsCache.Store("AlarmIntent", &sync.Map{})
	}

	EvaluateDependencies(obj, "Alarm")

	objMap, ok := compilationif.ObjsCache.Load("Intent")
	if ok {
		objMap.(*sync.Map).Store(obj.GetUUID(), intentObj)
	}

	service.Debug("UpdateAlarm", objMap.(*sync.Map))

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().UpdateAlarm(ctx, request)
}

// DeleteAlarm handles delete request
func (service *PluginService) DeleteAlarm(ctx context.Context, request *services.DeleteAlarmRequest) (*services.DeleteAlarmResponse, error) {
	log.Println(" DeleteAlarm ENTERED")

	objUUID := request.GetID()

	//intentObj := AlarmIntent {
	//Alarm: *obj,
	//}

	//EvaluateDependencies(intentObj, "Alarm")

	objMap, ok := compilationif.ObjsCache.Load("AlarmIntent")
	if ok {
		objMap.(*sync.Map).Delete(objUUID)
	}

	service.Debug("DeleteAlarm", objMap.(*sync.Map))

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().DeleteAlarm(ctx, request)
}

// GetAlarm handles get request
func (service *PluginService) GetAlarm(ctx context.Context, request *services.GetAlarmRequest) (*services.GetAlarmResponse, error) {
	objMap, ok := compilationif.ObjsCache.Load("Alarm")
	if !ok {
		return nil, errors.New("Alarm get failed ")
	}

	obj, ok := objMap.(*sync.Map).Load(request.GetID())
	if !ok {
		return nil, errors.New("Alarm get failed ")
	}

	response := &services.GetAlarmResponse{
		Alarm: obj.(*models.Alarm),
	}
	return response, nil
}
