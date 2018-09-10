package logic

import (
	"context"

	"github.com/pkg/errors"

	"github.com/Juniper/contrail/pkg/compilation/intent"
	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/models/basemodels"
	"github.com/Juniper/contrail/pkg/services"
)

// SecurityGroupIntent contains Intent Compiler state for SecurityGroup
type SecurityGroupIntent struct {
	intent.BaseIntent
	*models.SecurityGroup
}

func (i *SecurityGroupIntent) GetObject() basemodels.Object {
	return i.SecurityGroup
}

// CreateSecurityGroup evaluates SecurityGroup dependencies.
func (s *Service) CreateSecurityGroup(
	ctx context.Context,
	request *services.CreateSecurityGroupRequest,
) (*services.CreateSecurityGroupResponse, error) {

	obj := request.GetSecurityGroup()

	i := &SecurityGroupIntent{
		SecurityGroup: obj,
	}

	s.cache.Store(i)

	ec := &intent.EvaluateContext{
		WriteService: s.WriteService,
	}
	err := s.EvaluateDependencies(ctx, ec, obj, "SecurityGroup")
	if err != nil {
		return nil, errors.Wrap(err, "failed to evaluate Security Group dependencies")
	}

	return s.BaseService.CreateSecurityGroup(ctx, request)
}

// Evaluate Creates default AccessControlList's for the already created SecurityGroup.
func (i *SecurityGroupIntent) Evaluate(
	ctx context.Context,
	evaluateContext *intent.EvaluateContext,
) error {
	ingressACL, egressACL := i.DefaultACLs()

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
