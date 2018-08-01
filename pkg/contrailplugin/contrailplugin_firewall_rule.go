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

// FirewallRuleIntent
//   A struct to store attributes related to FirewallRule
//   needed by Intent Compiler
type FirewallRuleIntent struct {
	Uuid string
}

// EvaluateFirewallRule - evaluates the FirewallRule
func EvaluateFirewallRule(obj interface{}) {
	resourceObj := obj.(FirewallRuleIntent)
	log.Println("EvaluateFirewallRule Called ", resourceObj)
}

// CreateFirewallRule handles create request
func (service *PluginService) CreateFirewallRule(ctx context.Context, request *services.CreateFirewallRuleRequest) (*services.CreateFirewallRuleResponse, error) {
	log.Println(" CreateFirewallRule Entered")

	obj := request.GetFirewallRule()

	intentObj := FirewallRuleIntent{
		Uuid: obj.GetUUID(),
	}

	if _, ok := compilationif.ObjsCache.Load("FirewallRuleIntent"); !ok {
		compilationif.ObjsCache.Store("FirewallRuleIntent", &sync.Map{})
	}
	objMap, ok := compilationif.ObjsCache.Load("FirewallRuleIntent")
	if ok {
		objMap.(*sync.Map).Store(obj.GetUUID(), intentObj)
	}

	service.Debug("CreateFirewallRule", objMap.(*sync.Map))

	EvaluateDependencies(obj, "FirewallRule")

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().CreateFirewallRule(ctx, request)
}

// UpdateFirewallRule handles update request
func (service *PluginService) UpdateFirewallRule(ctx context.Context, request *services.UpdateFirewallRuleRequest) (*services.UpdateFirewallRuleResponse, error) {
	log.Println(" UpdateFirewallRule ENTERED")

	obj := request.GetFirewallRule()

	intentObj := FirewallRuleIntent{
		Uuid: obj.GetUUID(),
	}

	if _, ok := compilationif.ObjsCache.Load("FirewallRuleIntent"); !ok {
		compilationif.ObjsCache.Store("FirewallRuleIntent", &sync.Map{})
	}

	EvaluateDependencies(obj, "FirewallRule")

	objMap, ok := compilationif.ObjsCache.Load("Intent")
	if ok {
		objMap.(*sync.Map).Store(obj.GetUUID(), intentObj)
	}

	service.Debug("UpdateFirewallRule", objMap.(*sync.Map))

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().UpdateFirewallRule(ctx, request)
}

// DeleteFirewallRule handles delete request
func (service *PluginService) DeleteFirewallRule(ctx context.Context, request *services.DeleteFirewallRuleRequest) (*services.DeleteFirewallRuleResponse, error) {
	log.Println(" DeleteFirewallRule ENTERED")

	objUUID := request.GetID()

	//intentObj := FirewallRuleIntent {
	//FirewallRule: *obj,
	//}

	//EvaluateDependencies(intentObj, "FirewallRule")

	objMap, ok := compilationif.ObjsCache.Load("FirewallRuleIntent")
	if ok {
		objMap.(*sync.Map).Delete(objUUID)
	}

	service.Debug("DeleteFirewallRule", objMap.(*sync.Map))

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().DeleteFirewallRule(ctx, request)
}

// GetFirewallRule handles get request
func (service *PluginService) GetFirewallRule(ctx context.Context, request *services.GetFirewallRuleRequest) (*services.GetFirewallRuleResponse, error) {
	objMap, ok := compilationif.ObjsCache.Load("FirewallRule")
	if !ok {
		return nil, errors.New("FirewallRule get failed ")
	}

	obj, ok := objMap.(*sync.Map).Load(request.GetID())
	if !ok {
		return nil, errors.New("FirewallRule get failed ")
	}

	response := &services.GetFirewallRuleResponse{
		FirewallRule: obj.(*models.FirewallRule),
	}
	return response, nil
}
