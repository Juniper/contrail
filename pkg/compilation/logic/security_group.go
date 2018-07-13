package logic

import (
	"context"

	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/services"
	"github.com/pkg/errors"
)

// CreateSecurityGroup creates default AccessControlList's for the already created SecurityGroup.
func (s *Service) CreateSecurityGroup(
	ctx context.Context, request *services.CreateSecurityGroupRequest,
) (*services.CreateSecurityGroupResponse, error) {
	ingressACL, egressACL := defaultSecurityGroupACLs(request.SecurityGroup)

	_, err := s.api.CreateAccessControlList(ctx, &services.CreateAccessControlListRequest{
		AccessControlList: ingressACL,
	})
	if err != nil {
		return nil, errors.Wrap(err, "Failed to create ingress access control list")
	}

	_, err = s.api.CreateAccessControlList(ctx, &services.CreateAccessControlListRequest{
		AccessControlList: egressACL,
	})
	if err != nil {
		return nil, errors.Wrap(err, "Failed to create egress access control list")
	}

	// TODO The response isn't needed.
	return &services.CreateSecurityGroupResponse{SecurityGroup: request.SecurityGroup}, nil
}

func defaultSecurityGroupACLs(
	securityGroup *models.SecurityGroup,
) (ingressACL *models.AccessControlList, egressACL *models.AccessControlList) {
	aclRules := policyRulesToACLRules(securityGroup.SecurityGroupEntries.PolicyRule)
	ingressRules, egressRules := splitACLRules(aclRules)

	ingressACL = aclFromRules("ingress-access-control-list", ingressRules)
	egressACL = aclFromRules("egress-access-control-list", egressRules)
	return ingressACL, egressACL
}

func splitACLRules(rules []*models.AclRuleType) (ingressRules []*models.AclRuleType, egressRules []*models.AclRuleType) {
	for _, rule := range rules {
		switch {
		case rule.MatchCondition.SRCAddress.SecurityGroup == "local":
			// add to egress
		case rule.MatchCondition.DSTAddress.SecurityGroup == "local":
			// add to ingress
		default:
			// error out
		}
	}
	return nil, nil
}

func policyRulesToACLRules(policyRules []*models.PolicyRuleType) []*models.AclRuleType {
	// TODO Implement
	return nil
}

func aclFromRules(name string, rules []*models.AclRuleType) *models.AccessControlList {
	return &models.AccessControlList{
		Name: name,
		AccessControlListEntries: &models.AclEntriesType{
			ACLRule: rules,
		},
	}
}
