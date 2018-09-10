package logic

import (
	"context"

	"github.com/pkg/errors"

	"github.com/Juniper/contrail/pkg/compilation/intent"
	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/services"
)

// SecurityGroupIntent contains Intent Compiler state for SecurityGroup
type SecurityGroupIntent struct {
	intent.BaseIntent
	*models.SecurityGroup
}

// CreateSecurityGroup evaluates SecurityGroup dependencies.
func (s *Service) CreateSecurityGroup(
	ctx context.Context,
	request *services.CreateSecurityGroupRequest,
) (*services.CreateSecurityGroupResponse, error) {
	i := &SecurityGroupIntent{
		SecurityGroup: request.GetSecurityGroup(),
	}

	err := s.handleCreate(ctx, i, nil, i.SecurityGroup)
	if err != nil {
		return nil, err
	}

	return s.BaseService.CreateSecurityGroup(ctx, request)
}

// Evaluate Creates default AccessControlList's for the already created SecurityGroup.
func (i *SecurityGroupIntent) Evaluate(
	ctx context.Context,
	evaluateContext *intent.EvaluateContext,
) error {
	ingressACL, egressACL := i.DefaultACLs(makeSGLoader(evaluateContext))

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

type securityGroupLoader struct {
	loader intent.Loader
}

func (l *securityGroupLoader) LoadByFQName(fqName []string) *models.SecurityGroup {
	i, ok := l.loader.LoadByFQName("SecurityGroup", fqName)
	if !ok {
		return nil
	}
	sgi, ok := i.(*SecurityGroupIntent)
	if !ok {
		return nil
	}
	return sgi.SecurityGroup
}

func makeSGLoader(ctx *intent.EvaluateContext) *securityGroupLoader {
	return &securityGroupLoader{
		loader: ctx.IntentLoader,
	}
}
