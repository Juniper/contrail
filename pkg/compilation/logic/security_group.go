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

// Evaluate Creates default AccessControlList's for the already created SecurityGroup.
func (i *SecurityGroupIntent) Evaluate(ctx context.Context, ec *intent.EvaluateContext) error {
	ingressACL, egressACL := i.DefaultACLs(ec)

	// TODO: Use batch create so that either both ACLs are created or none.
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
	// TODO (Kamil) use loadSecurityGroupIntent()
	i := ec.IntentLoader.Load(models.TypeNameSecurityGroup,
		intent.ByFQName(basemodels.ParseFQName(addr.SecurityGroup)))
	sg, _ := i.(*SecurityGroupIntent)
	rs.FQNameToSG[addr.SecurityGroup] = sg.SecurityGroup
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
