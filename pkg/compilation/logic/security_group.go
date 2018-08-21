package logic

import (
	"context"
	"sync"

	"github.com/pkg/errors"

	"github.com/Juniper/contrail/pkg/compilationif"
	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/services"
)

// SecurityGroupIntent contains Intent Compiler state for SecurityGroup
type SecurityGroupIntent struct {
	BaseIntent
	*models.SecurityGroup
}

// CreateSecurityGroup evaluates SecurityGroup dependencies.
func (s *Service) CreateSecurityGroup(
	ctx context.Context, request *services.CreateSecurityGroupRequest,
) (*services.CreateSecurityGroupResponse, error) {

	obj := request.GetSecurityGroup()

	intent := &SecurityGroupIntent{
		SecurityGroup: obj,
	}

	if _, ok := compilationif.ObjsCache.Load("SecurityGroupIntent"); !ok {
		compilationif.ObjsCache.Store("SecurityGroupIntent", &sync.Map{})
	}

	objMap, ok := compilationif.ObjsCache.Load("SecurityGroupIntent")
	if ok {
		objMap.(*sync.Map).Store(obj.GetUUID(), intent)
	}

	ec := &EvaluateContext{
		WriteService: s.WriteService,
	}
	err := EvaluateDependencies(ctx, ec, obj, "SecurityGroup")
	if err != nil {
		return nil, errors.Wrap(err, "failed to evaluate Security Group dependencies")
	}

	return s.BaseService.CreateSecurityGroup(ctx, request)
}

// Evaluate Creates default AccessControlList's for the already created SecurityGroup.
func (s *SecurityGroupIntent) Evaluate(ctx context.Context, evaluateContext *EvaluateContext) error {
	ingressACL, egressACL := s.DefaultACLs()

	_, err := evaluateContext.WriteService.CreateAccessControlList(ctx, &services.CreateAccessControlListRequest{
		AccessControlList: ingressACL,
	})
	if err != nil {
		return errors.Wrap(err, "failed to create ingress access control list")
	}

	_, err = evaluateContext.WriteService.CreateAccessControlList(ctx, &services.CreateAccessControlListRequest{
		AccessControlList: egressACL,
	})
	if err != nil {
		return errors.Wrap(err, "failed to create egress access control list")
	}

	return nil
}
