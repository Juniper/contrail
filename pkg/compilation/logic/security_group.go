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

	ingressACL, egressACL *models.AccessControlList
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
func (i *SecurityGroupIntent) Evaluate(ctx context.Context, ec *intent.EvaluateContext) error {
	ingressACL, egressACL := i.DefaultACLs()

	var err error
	i.ingressACL, err = createACL(ctx, ec.WriteService, ingressACL)
	if err != nil {
		return errors.Wrap(err, "failed to create ingress access control list")
	}

	i.egressACL, err = createACL(ctx, ec.WriteService, egressACL)
	if err != nil {
		return errors.Wrap(err, "failed to create egress access control list")
	}

	return nil
}

func loadSecurityGroupIntent(cache *intent.Cache, uuid string) (intent *SecurityGroupIntent, ok bool) {
	i, ok := cache.Load(models.KindSecurityGroup, uuid)
	if !ok {
		return nil, false
	}
	intent, ok = i.(*SecurityGroupIntent)
	return intent, ok
}

func createACL(
	ctx context.Context, writeService services.WriteService, acl *models.AccessControlList,
) (*models.AccessControlList, error) {
	response, err := writeService.CreateAccessControlList(
		ctx, &services.CreateAccessControlListRequest{
			AccessControlList: acl,
		})
	return response.GetAccessControlList(), err
}
