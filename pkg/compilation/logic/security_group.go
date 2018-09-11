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
	i := &SecurityGroupIntent{
		SecurityGroup: request.GetSecurityGroup(),
	}

	err := s.handleCreate(ctx, i, nil, i.SecurityGroup)
	if err != nil {
		return nil, err
	}

	return s.BaseService.CreateSecurityGroup(ctx, request)
}

// DeleteSecurityGroup evaluates SecurityGroup dependencies.
func (s *Service) DeleteSecurityGroup(
	ctx context.Context,
	request *services.DeleteSecurityGroupRequest,
) (*services.DeleteSecurityGroupResponse, error) {

	i, ok := loadSecurityGroupIntent(s.cache, request.ID)
	if !ok {
		return nil, errors.New("failed to process SecurityGroup deletion: SecurityGroupIntent not found in cache")
	}

	// TODO Check if i.ingressACL exists
	err := deleteACL(ctx, s.WriteService, i.ingressACL.GetUUID())
	if err != nil {
		return nil, errors.Wrap(err, "failed to delete ingress access control list")
	}

	err = deleteACL(ctx, s.WriteService, i.egressACL.GetUUID())
	if err != nil {
		return nil, errors.Wrap(err, "failed to delete egress access control list")
	}

	// TODO Delete it even if there were errors, like the Python code does?
	s.cache.Delete(models.KindSecurityGroup, i.GetUUID())

	return s.BaseService.DeleteSecurityGroup(ctx, request)
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

func deleteACL(ctx context.Context, writeService services.WriteService, uuid string) error {
	_, err := writeService.DeleteAccessControlList(
		ctx, &services.DeleteAccessControlListRequest{
			ID: uuid,
		})
	return err
}
