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

// DashboardIntent
//   A struct to store attributes related to Dashboard
//   needed by Intent Compiler
type DashboardIntent struct {
	Uuid string
}

// EvaluateDashboard - evaluates the Dashboard
func EvaluateDashboard(obj interface{}) {
	resourceObj := obj.(DashboardIntent)
	log.Println("EvaluateDashboard Called ", resourceObj)
}

// CreateDashboard handles create request
func (service *PluginService) CreateDashboard(ctx context.Context, request *services.CreateDashboardRequest) (*services.CreateDashboardResponse, error) {
	log.Println(" CreateDashboard Entered")

	obj := request.GetDashboard()

	intentObj := DashboardIntent{
		Uuid: obj.GetUUID(),
	}

	if _, ok := compilationif.ObjsCache.Load("DashboardIntent"); !ok {
		compilationif.ObjsCache.Store("DashboardIntent", &sync.Map{})
	}
	objMap, ok := compilationif.ObjsCache.Load("DashboardIntent")
	if ok {
		objMap.(*sync.Map).Store(obj.GetUUID(), intentObj)
	}

	service.Debug("CreateDashboard", objMap.(*sync.Map))

	EvaluateDependencies(obj, "Dashboard")

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().CreateDashboard(ctx, request)
}

// UpdateDashboard handles update request
func (service *PluginService) UpdateDashboard(ctx context.Context, request *services.UpdateDashboardRequest) (*services.UpdateDashboardResponse, error) {
	log.Println(" UpdateDashboard ENTERED")

	obj := request.GetDashboard()

	intentObj := DashboardIntent{
		Uuid: obj.GetUUID(),
	}

	if _, ok := compilationif.ObjsCache.Load("DashboardIntent"); !ok {
		compilationif.ObjsCache.Store("DashboardIntent", &sync.Map{})
	}

	EvaluateDependencies(obj, "Dashboard")

	objMap, ok := compilationif.ObjsCache.Load("Intent")
	if ok {
		objMap.(*sync.Map).Store(obj.GetUUID(), intentObj)
	}

	service.Debug("UpdateDashboard", objMap.(*sync.Map))

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().UpdateDashboard(ctx, request)
}

// DeleteDashboard handles delete request
func (service *PluginService) DeleteDashboard(ctx context.Context, request *services.DeleteDashboardRequest) (*services.DeleteDashboardResponse, error) {
	log.Println(" DeleteDashboard ENTERED")

	objUUID := request.GetID()

	//intentObj := DashboardIntent {
	//Dashboard: *obj,
	//}

	//EvaluateDependencies(intentObj, "Dashboard")

	objMap, ok := compilationif.ObjsCache.Load("DashboardIntent")
	if ok {
		objMap.(*sync.Map).Delete(objUUID)
	}

	service.Debug("DeleteDashboard", objMap.(*sync.Map))

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().DeleteDashboard(ctx, request)
}

// GetDashboard handles get request
func (service *PluginService) GetDashboard(ctx context.Context, request *services.GetDashboardRequest) (*services.GetDashboardResponse, error) {
	objMap, ok := compilationif.ObjsCache.Load("Dashboard")
	if !ok {
		return nil, errors.New("Dashboard get failed ")
	}

	obj, ok := objMap.(*sync.Map).Load(request.GetID())
	if !ok {
		return nil, errors.New("Dashboard get failed ")
	}

	response := &services.GetDashboardResponse{
		Dashboard: obj.(*models.Dashboard),
	}
	return response, nil
}
