package logic

import (
	"context"

	"github.com/gogo/protobuf/types"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"

	"github.com/Juniper/asf/pkg/errutil"
	"github.com/Juniper/asf/pkg/models/basemodels"
	"github.com/Juniper/contrail/pkg/compilation/intent"
	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/services"
)

// SecurityGroupIntent contains Intent Compiler state for SecurityGroup
type SecurityGroupIntent struct {
	intent.BaseIntent
	*models.SecurityGroup

	referredSGs           map[string]*SecurityGroupIntent
	ingressACL, egressACL *models.AccessControlList
}

// LoadSecurityGroupIntent loads a security group intent from cache.
func LoadSecurityGroupIntent(loader intent.Loader, query intent.Query) *SecurityGroupIntent {
	intent := loader.Load(models.KindSecurityGroup, query)
	sgIntent, ok := intent.(*SecurityGroupIntent)
	if ok == false {
		logrus.Warning("Cannot cast intent to Security Group Intent")
	}
	return sgIntent
}

// GetObject returns embedded resource object
func (i *SecurityGroupIntent) GetObject() basemodels.Object {
	return i.SecurityGroup
}

// CreateSecurityGroup evaluates SecurityGroup dependencies.
func (s *Service) CreateSecurityGroup(
	ctx context.Context,
	request *services.CreateSecurityGroupRequest,
) (*services.CreateSecurityGroupResponse, error) {
	i := &SecurityGroupIntent{
		SecurityGroup: request.GetSecurityGroup(),
		referredSGs:   map[string]*SecurityGroupIntent{},
	}

	if err := s.storeAndEvaluate(ctx, i); err != nil {
		return nil, err
	}
	i.processReferredSecurityGroups(s.evaluateContext())

	return s.BaseService.CreateSecurityGroup(ctx, request)
}

// UpdateSecurityGroup evaluates SecurityGroup dependencies.
func (s *Service) UpdateSecurityGroup(
	ctx context.Context,
	request *services.UpdateSecurityGroupRequest,
) (*services.UpdateSecurityGroupResponse, error) {
	sg := request.GetSecurityGroup()
	if sg == nil {
		return nil, errors.New("failed to update Security Group." +
			" Security Group Request needs to contain resource!")
	}

	i := LoadSecurityGroupIntent(s.cache, intent.ByUUID(sg.GetUUID()))
	if i == nil {
		return nil, errors.Errorf("cannot load intent for security group %v", sg.GetUUID())
	}

	i.SecurityGroup = sg
	if err := s.storeAndEvaluate(ctx, i); err != nil {
		return nil, err
	}
	i.processReferredSecurityGroups(s.evaluateContext())

	return s.BaseService.UpdateSecurityGroup(ctx, request)
}

// DeleteSecurityGroup evaluates SecurityGroup dependencies.
func (s *Service) DeleteSecurityGroup(
	ctx context.Context,
	request *services.DeleteSecurityGroupRequest,
) (*services.DeleteSecurityGroupResponse, error) {

	i := LoadSecurityGroupIntent(s.cache, intent.ByUUID(request.GetID()))
	if i == nil {
		return nil, errors.New("failed to process SecurityGroup deletion: SecurityGroupIntent not found in cache")
	}

	if err := i.deleteDefaultACLs(ctx, s.WriteService); err != nil {
		return nil, errors.Wrap(err, "failed to process SecurityGroup deletion")
	}
	i.SecurityGroupEntries = nil

	ec := s.evaluateContext()

	s.cache.Delete(models.KindSecurityGroup, intent.ByUUID(i.GetUUID()))

	err := s.EvaluateDependencies(ctx, ec, i)
	if err != nil {
		return nil, errors.Wrap(err, "failed to evaluate Security Group dependencies")
	}
	i.processReferredSecurityGroups(ec)

	return s.BaseService.DeleteSecurityGroup(ctx, request)
}

func (i *SecurityGroupIntent) deleteDefaultACLs(ctx context.Context, writeService services.WriteService) error {
	var multiError errutil.MultiError

	if err := deleteACLIfNeeded(ctx, writeService, i.ingressACL); err != nil {
		multiError = append(multiError, errors.Wrap(err, "failed to delete ingress access control list"))
	} else {
		i.ingressACL = nil
	}

	if err := deleteACLIfNeeded(ctx, writeService, i.egressACL); err != nil {
		multiError = append(multiError, errors.Wrap(err, "failed to delete egress access control list"))
	} else {
		i.egressACL = nil
	}

	if len(multiError) > 0 {
		return multiError
	}
	return nil
}

func deleteACLIfNeeded(
	ctx context.Context, writeService services.WriteService, acl *models.AccessControlList,
) error {
	if acl == nil {
		return nil
	}
	return deleteACL(ctx, writeService, acl.GetUUID())
}

// Evaluate Creates default AccessControlList's for the already created SecurityGroup.
func (i *SecurityGroupIntent) Evaluate(ctx context.Context, ec *intent.EvaluateContext) error {
	newIngressACL, newEgressACL := i.DefaultACLs(ec)
	var err error

	// TODO: Use batch create so that either both ACLs are created or none.
	if i.ingressACL, err = createOrUpdateDefaultACL(ctx, ec, i.ingressACL, newIngressACL); err != nil {
		return errors.Wrap(err, "failed to create ingress access control list")
	}
	if i.egressACL, err = createOrUpdateDefaultACL(ctx, ec, i.egressACL, newEgressACL); err != nil {
		return errors.Wrap(err, "failed to create egress access control list")
	}
	return nil
}

// DefaultACLs returns default ACLs corresponding to the security group's policy rules.
func (i *SecurityGroupIntent) DefaultACLs(ec *intent.EvaluateContext) (
	ingressACL *models.AccessControlList, egressACL *models.AccessControlList) {

	rs := &models.PolicyRulesWithRefs{
		Rules:      i.GetSecurityGroupEntries().GetPolicyRule(),
		FQNameToSG: make(map[string]*models.SecurityGroup),
	}
	resolveSGRefs(rs, ec)

	ingressRules, egressRules := rs.ToACLRules()

	ingressACL = i.MakeChildACL("ingress-access-control-list", ingressRules)
	egressACL = i.MakeChildACL("egress-access-control-list", egressRules)
	return ingressACL, egressACL
}

func (i *SecurityGroupIntent) processReferredSecurityGroups(ec *intent.EvaluateContext) {
	var rules []*models.PolicyRuleType
	if sges := i.GetSecurityGroupEntries(); sges != nil {
		rules = sges.GetPolicyRule()
	}
	referredSGs := map[string]*SecurityGroupIntent{}
	for _, rule := range rules {
		for _, addr := range rule.GetSRCAddresses() {
			if ok, sg := checkAddressIsReferred(ec, addr); ok {
				referredSGs[sg.GetUUID()] = sg
			}
		}
		for _, addr := range rule.GetDSTAddresses() {
			if ok, sg := checkAddressIsReferred(ec, addr); ok {
				referredSGs[sg.GetUUID()] = sg
			}
		}
	}
	for _, sg := range subtractSGSets(i.referredSGs, referredSGs) {
		logrus.Debugf("removing dependent intent: %s-%s %s-%s", sg.Kind(), i.Kind(), sg.GetUUID(), i.GetUUID())
		sg.RemoveDependentIntent(i)
	}
	for _, sg := range subtractSGSets(referredSGs, i.referredSGs) {
		logrus.Debugf("adding dependent intent: %s-%s %s-%s", sg.Kind(), i.Kind(), sg.GetUUID(), i.GetUUID())
		sg.AddDependentIntent(i)
	}
	i.referredSGs = referredSGs
}

func subtractSGSets(old, new map[string]*SecurityGroupIntent) map[string]*SecurityGroupIntent {
	result := map[string]*SecurityGroupIntent{}
	for uuid, sg := range old {
		_, ok := new[uuid]
		if !ok {
			result[uuid] = sg
		}
	}
	return result
}

func checkAddressIsReferred(
	ec *intent.EvaluateContext,
	address *models.AddressType,
) (bool, *SecurityGroupIntent) {
	if !address.IsSecurityGroupNameAReference() {
		return false, nil
	}
	fqName := basemodels.ParseFQName(address.GetSecurityGroup())
	sg := LoadSecurityGroupIntent(ec.IntentLoader, intent.ByFQName(fqName))
	return sg != nil, sg
}

func resolveSGRefs(rs *models.PolicyRulesWithRefs, ec *intent.EvaluateContext) {
	for _, r := range rs.Rules {
		for _, addr := range r.SRCAddresses {
			resolveSGRef(rs, addr, ec)
		}
		for _, addr := range r.DSTAddresses {
			resolveSGRef(rs, addr, ec)
		}
	}
}

func resolveSGRef(rs *models.PolicyRulesWithRefs, addr *models.AddressType, ec *intent.EvaluateContext) {
	if !addr.IsSecurityGroupNameAReference() {
		return
	}
	i := LoadSecurityGroupIntent(
		ec.IntentLoader,
		intent.ByFQName(basemodels.ParseFQName(addr.SecurityGroup)))
	if i == nil {
		return
	}
	rs.FQNameToSG[addr.SecurityGroup] = i.SecurityGroup
}

func createOrUpdateDefaultACL(
	ctx context.Context,
	ec *intent.EvaluateContext,
	oldACL *models.AccessControlList,
	newACL *models.AccessControlList,
) (*models.AccessControlList, error) {
	if oldACL == nil {
		return createACL(ctx, ec.WriteService, newACL)
	}
	newACL.UUID = oldACL.GetUUID()
	updatedACL, err := updateACL(ctx, ec.WriteService, newACL)
	if err != nil {
		return oldACL, err
	}
	return updatedACL, nil
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
			FieldMask: types.FieldMask{
				Paths: []string{models.AccessControlListFieldAccessControlListEntries},
			},
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
