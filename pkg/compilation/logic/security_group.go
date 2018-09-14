package logic

import (
	"context"
	"fmt"

	"github.com/gogo/protobuf/types"
	"github.com/pkg/errors"

	"github.com/Juniper/contrail/pkg/common"
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

// UpdateSecurityGroup evaluates SecurityGroup dependencies.
func (s *Service) UpdateSecurityGroup(
	ctx context.Context,
	request *services.UpdateSecurityGroupRequest,
) (*services.UpdateSecurityGroupResponse, error) {
	var sg *models.SecurityGroup
	if sg = request.GetSecurityGroup(); sg == nil {
		return nil, errors.New("failed to update Security Group." +
			" Security Group Request needs to contain resource!")
	}

	i, ok := loadSecurityGroupIntent(s.cache, sg.GetUUID())
	if !ok {
		return nil, errors.Errorf("cannot load intent from security group %v", request.SecurityGroup.GetUUID())
	}

	i.SecurityGroup = sg

	ec := &intent.EvaluateContext{
		WriteService: s.WriteService,
	}
	err := s.EvaluateDependencies(ctx, ec, sg, "SecurityGroup")
	if err != nil {
		return nil, errors.Wrap(err, "failed to evaluate Security Group dependencies")
	}

	return s.BaseService.UpdateSecurityGroup(ctx, request)
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

	err := deleteDefaultACLs(ctx, s.WriteService, i)
	if err != nil {
		return nil, errors.Wrap(err, "failed to process SecurityGroup deletion")
	}

	s.cache.Delete(models.KindSecurityGroup, i.GetUUID())

	return s.BaseService.DeleteSecurityGroup(ctx, request)
}

func deleteDefaultACLs(ctx context.Context, writeService services.WriteService, i *SecurityGroupIntent) error {
	var multiError common.MultiError

	err := deleteAndUnsetACL(ctx, writeService, &i.ingressACL, "ingress")
	if err != nil {
		multiError = append(multiError, err)
	}

	err = deleteAndUnsetACL(ctx, writeService, &i.egressACL, "egress")
	if err != nil {
		multiError = append(multiError, err)
	}

	if multiError != nil {
		return multiError
	}
	return nil
}

func deleteAndUnsetACL(
	ctx context.Context, writeService services.WriteService, acl **models.AccessControlList, aclType string,
) error {
	if *acl == nil {
		return nil
	}

	err := deleteACL(ctx, writeService, (*acl).GetUUID())
	return errors.Wrap(err, fmt.Sprintf("failed to delete %s access control list", aclType))
}

// Evaluate Creates default AccessControlList's for the already created SecurityGroup.
func (i *SecurityGroupIntent) Evaluate(ctx context.Context, ec *intent.EvaluateContext) error {
	ingressACL, egressACL := i.DefaultACLs()

	// TODO: Use batch create so that either both ACLs are created or none.
	var err error
	if i.ingressACL == nil {
		i.ingressACL, err = createACL(ctx, ec.WriteService, ingressACL)
		if err != nil {
			return errors.Wrap(err, "failed to create egress access control list")
		}
	} else {
		keepUUID := i.ingressACL.GetUUID()
		i.ingressACL, err = updateACL(ctx, ec.WriteService, ingressACL)
		if err != nil {
			return errors.Wrap(err, "failed to create egress access control list")
		}
		i.ingressACL.UUID = keepUUID
	}

	if i.egressACL == nil {
		i.egressACL, err = createACL(ctx, ec.WriteService, egressACL)
		if err != nil {
			return errors.Wrap(err, "failed to create egress access control list")
		}
	} else {
		keepUUID := i.egressACL.GetUUID()
		i.egressACL, err = updateACL(ctx, ec.WriteService, egressACL)
		if err != nil {
			return errors.Wrap(err, "failed to create egress access control list")
		}
		i.egressACL.UUID = keepUUID
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

func updateACL(
	ctx context.Context, writeService services.WriteService, acl *models.AccessControlList,
) (*models.AccessControlList, error) {
	response, err := writeService.UpdateAccessControlList(
		ctx, &services.UpdateAccessControlListRequest{
			AccessControlList: acl,
			FieldMask:         types.FieldMask{Paths: []string{models.AccessControlListFieldAccessControlListEntries}},
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
