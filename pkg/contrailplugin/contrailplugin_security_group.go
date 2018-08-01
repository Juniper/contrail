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

// SecurityGroupIntent
//   A struct to store attributes related to SecurityGroup
//   needed by Intent Compiler
type SecurityGroupIntent struct {
	Uuid string
}

// EvaluateSecurityGroup - evaluates the SecurityGroup
func EvaluateSecurityGroup(obj interface{}) {
	resourceObj := obj.(SecurityGroupIntent)
	log.Println("EvaluateSecurityGroup Called ", resourceObj)
}

// CreateSecurityGroup handles create request
func (service *PluginService) CreateSecurityGroup(ctx context.Context, request *services.CreateSecurityGroupRequest) (*services.CreateSecurityGroupResponse, error) {
	log.Println(" CreateSecurityGroup Entered")

	obj := request.GetSecurityGroup()

	intentObj := SecurityGroupIntent{
		Uuid: obj.GetUUID(),
	}

	if _, ok := compilationif.ObjsCache.Load("SecurityGroupIntent"); !ok {
		compilationif.ObjsCache.Store("SecurityGroupIntent", &sync.Map{})
	}
	objMap, ok := compilationif.ObjsCache.Load("SecurityGroupIntent")
	if ok {
		objMap.(*sync.Map).Store(obj.GetUUID(), intentObj)
	}

	service.Debug("CreateSecurityGroup", objMap.(*sync.Map))

	EvaluateDependencies(obj, "SecurityGroup")

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().CreateSecurityGroup(ctx, request)
}

// UpdateSecurityGroup handles update request
func (service *PluginService) UpdateSecurityGroup(ctx context.Context, request *services.UpdateSecurityGroupRequest) (*services.UpdateSecurityGroupResponse, error) {
	log.Println(" UpdateSecurityGroup ENTERED")

	obj := request.GetSecurityGroup()

	intentObj := SecurityGroupIntent{
		Uuid: obj.GetUUID(),
	}

	if _, ok := compilationif.ObjsCache.Load("SecurityGroupIntent"); !ok {
		compilationif.ObjsCache.Store("SecurityGroupIntent", &sync.Map{})
	}

	EvaluateDependencies(obj, "SecurityGroup")

	objMap, ok := compilationif.ObjsCache.Load("Intent")
	if ok {
		objMap.(*sync.Map).Store(obj.GetUUID(), intentObj)
	}

	service.Debug("UpdateSecurityGroup", objMap.(*sync.Map))

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().UpdateSecurityGroup(ctx, request)
}

// DeleteSecurityGroup handles delete request
func (service *PluginService) DeleteSecurityGroup(ctx context.Context, request *services.DeleteSecurityGroupRequest) (*services.DeleteSecurityGroupResponse, error) {
	log.Println(" DeleteSecurityGroup ENTERED")

	objUUID := request.GetID()

	//intentObj := SecurityGroupIntent {
	//SecurityGroup: *obj,
	//}

	//EvaluateDependencies(intentObj, "SecurityGroup")

	objMap, ok := compilationif.ObjsCache.Load("SecurityGroupIntent")
	if ok {
		objMap.(*sync.Map).Delete(objUUID)
	}

	service.Debug("DeleteSecurityGroup", objMap.(*sync.Map))

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().DeleteSecurityGroup(ctx, request)
}

// GetSecurityGroup handles get request
func (service *PluginService) GetSecurityGroup(ctx context.Context, request *services.GetSecurityGroupRequest) (*services.GetSecurityGroupResponse, error) {
	objMap, ok := compilationif.ObjsCache.Load("SecurityGroup")
	if !ok {
		return nil, errors.New("SecurityGroup get failed ")
	}

	obj, ok := objMap.(*sync.Map).Load(request.GetID())
	if !ok {
		return nil, errors.New("SecurityGroup get failed ")
	}

	response := &services.GetSecurityGroupResponse{
		SecurityGroup: obj.(*models.SecurityGroup),
	}
	return response, nil
}
