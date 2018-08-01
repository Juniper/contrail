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

// DsaRuleIntent
//   A struct to store attributes related to DsaRule
//   needed by Intent Compiler
type DsaRuleIntent struct {
	Uuid string
}

// EvaluateDsaRule - evaluates the DsaRule
func EvaluateDsaRule(obj interface{}) {
	resourceObj := obj.(DsaRuleIntent)
	log.Println("EvaluateDsaRule Called ", resourceObj)
}

// CreateDsaRule handles create request
func (service *PluginService) CreateDsaRule(ctx context.Context, request *services.CreateDsaRuleRequest) (*services.CreateDsaRuleResponse, error) {
	log.Println(" CreateDsaRule Entered")

	obj := request.GetDsaRule()

	intentObj := DsaRuleIntent{
		Uuid: obj.GetUUID(),
	}

	if _, ok := compilationif.ObjsCache.Load("DsaRuleIntent"); !ok {
		compilationif.ObjsCache.Store("DsaRuleIntent", &sync.Map{})
	}
	objMap, ok := compilationif.ObjsCache.Load("DsaRuleIntent")
	if ok {
		objMap.(*sync.Map).Store(obj.GetUUID(), intentObj)
	}

	service.Debug("CreateDsaRule", objMap.(*sync.Map))

	EvaluateDependencies(obj, "DsaRule")

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().CreateDsaRule(ctx, request)
}

// UpdateDsaRule handles update request
func (service *PluginService) UpdateDsaRule(ctx context.Context, request *services.UpdateDsaRuleRequest) (*services.UpdateDsaRuleResponse, error) {
	log.Println(" UpdateDsaRule ENTERED")

	obj := request.GetDsaRule()

	intentObj := DsaRuleIntent{
		Uuid: obj.GetUUID(),
	}

	if _, ok := compilationif.ObjsCache.Load("DsaRuleIntent"); !ok {
		compilationif.ObjsCache.Store("DsaRuleIntent", &sync.Map{})
	}

	EvaluateDependencies(obj, "DsaRule")

	objMap, ok := compilationif.ObjsCache.Load("Intent")
	if ok {
		objMap.(*sync.Map).Store(obj.GetUUID(), intentObj)
	}

	service.Debug("UpdateDsaRule", objMap.(*sync.Map))

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().UpdateDsaRule(ctx, request)
}

// DeleteDsaRule handles delete request
func (service *PluginService) DeleteDsaRule(ctx context.Context, request *services.DeleteDsaRuleRequest) (*services.DeleteDsaRuleResponse, error) {
	log.Println(" DeleteDsaRule ENTERED")

	objUUID := request.GetID()

	//intentObj := DsaRuleIntent {
	//DsaRule: *obj,
	//}

	//EvaluateDependencies(intentObj, "DsaRule")

	objMap, ok := compilationif.ObjsCache.Load("DsaRuleIntent")
	if ok {
		objMap.(*sync.Map).Delete(objUUID)
	}

	service.Debug("DeleteDsaRule", objMap.(*sync.Map))

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().DeleteDsaRule(ctx, request)
}

// GetDsaRule handles get request
func (service *PluginService) GetDsaRule(ctx context.Context, request *services.GetDsaRuleRequest) (*services.GetDsaRuleResponse, error) {
	objMap, ok := compilationif.ObjsCache.Load("DsaRule")
	if !ok {
		return nil, errors.New("DsaRule get failed ")
	}

	obj, ok := objMap.(*sync.Map).Load(request.GetID())
	if !ok {
		return nil, errors.New("DsaRule get failed ")
	}

	response := &services.GetDsaRuleResponse{
		DsaRule: obj.(*models.DsaRule),
	}
	return response, nil
}
