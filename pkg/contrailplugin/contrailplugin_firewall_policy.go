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

// FirewallPolicyIntent
//   A struct to store attributes related to FirewallPolicy
//   needed by Intent Compiler
type FirewallPolicyIntent struct {
	Uuid string
}

// EvaluateFirewallPolicy - evaluates the FirewallPolicy
func EvaluateFirewallPolicy(obj interface{}) {
	resourceObj := obj.(FirewallPolicyIntent)
	log.Println("EvaluateFirewallPolicy Called ", resourceObj)
}

// CreateFirewallPolicy handles create request
func (service *PluginService) CreateFirewallPolicy(ctx context.Context, request *services.CreateFirewallPolicyRequest) (*services.CreateFirewallPolicyResponse, error) {
	log.Println(" CreateFirewallPolicy Entered")

	obj := request.GetFirewallPolicy()

	intentObj := FirewallPolicyIntent{
		Uuid: obj.GetUUID(),
	}

	if _, ok := compilationif.ObjsCache.Load("FirewallPolicyIntent"); !ok {
		compilationif.ObjsCache.Store("FirewallPolicyIntent", &sync.Map{})
	}
	objMap, ok := compilationif.ObjsCache.Load("FirewallPolicyIntent")
	if ok {
		objMap.(*sync.Map).Store(obj.GetUUID(), intentObj)
	}

	service.Debug("CreateFirewallPolicy", objMap.(*sync.Map))

	EvaluateDependencies(obj, "FirewallPolicy")

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().CreateFirewallPolicy(ctx, request)
}

// UpdateFirewallPolicy handles update request
func (service *PluginService) UpdateFirewallPolicy(ctx context.Context, request *services.UpdateFirewallPolicyRequest) (*services.UpdateFirewallPolicyResponse, error) {
	log.Println(" UpdateFirewallPolicy ENTERED")

	obj := request.GetFirewallPolicy()

	intentObj := FirewallPolicyIntent{
		Uuid: obj.GetUUID(),
	}

	if _, ok := compilationif.ObjsCache.Load("FirewallPolicyIntent"); !ok {
		compilationif.ObjsCache.Store("FirewallPolicyIntent", &sync.Map{})
	}

	EvaluateDependencies(obj, "FirewallPolicy")

	objMap, ok := compilationif.ObjsCache.Load("Intent")
	if ok {
		objMap.(*sync.Map).Store(obj.GetUUID(), intentObj)
	}

	service.Debug("UpdateFirewallPolicy", objMap.(*sync.Map))

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().UpdateFirewallPolicy(ctx, request)
}

// DeleteFirewallPolicy handles delete request
func (service *PluginService) DeleteFirewallPolicy(ctx context.Context, request *services.DeleteFirewallPolicyRequest) (*services.DeleteFirewallPolicyResponse, error) {
	log.Println(" DeleteFirewallPolicy ENTERED")

	objUUID := request.GetID()

	//intentObj := FirewallPolicyIntent {
	//FirewallPolicy: *obj,
	//}

	//EvaluateDependencies(intentObj, "FirewallPolicy")

	objMap, ok := compilationif.ObjsCache.Load("FirewallPolicyIntent")
	if ok {
		objMap.(*sync.Map).Delete(objUUID)
	}

	service.Debug("DeleteFirewallPolicy", objMap.(*sync.Map))

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().DeleteFirewallPolicy(ctx, request)
}

// GetFirewallPolicy handles get request
func (service *PluginService) GetFirewallPolicy(ctx context.Context, request *services.GetFirewallPolicyRequest) (*services.GetFirewallPolicyResponse, error) {
	objMap, ok := compilationif.ObjsCache.Load("FirewallPolicy")
	if !ok {
		return nil, errors.New("FirewallPolicy get failed ")
	}

	obj, ok := objMap.(*sync.Map).Load(request.GetID())
	if !ok {
		return nil, errors.New("FirewallPolicy get failed ")
	}

	response := &services.GetFirewallPolicyResponse{
		FirewallPolicy: obj.(*models.FirewallPolicy),
	}
	return response, nil
}
