package logic

import (
	"context"

	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"

	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/services"
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

	ingressACL = makeACL("ingress-access-control-list", ingressRules)
	egressACL = makeACL("egress-access-control-list", egressRules)
	return ingressACL, egressACL
}

func splitACLRules(rules []*models.AclRuleType) (ingressRules []*models.AclRuleType, egressRules []*models.AclRuleType) {
	for _, rule := range rules {
		switch {
		case rule.MatchCondition.DSTAddress.SecurityGroup == "local":
			ingressRules = append(ingressRules, rule)
		case rule.MatchCondition.SRCAddress.SecurityGroup == "local":
			egressRules = append(egressRules, rule)
		default:
			log.WithField("acl_rule", rule).Error(
				"Ignoring ACL rule: neither source nor destination address is local")
		}
	}
	return ingressRules, egressRules
}

func policyRulesToACLRules(policyRules []*models.PolicyRuleType) (aclRules []*models.AclRuleType) {
	for _, policyRule := range policyRules {
		for _, sourceAddress := range policyRule.SRCAddresses {
			for _, sourcePort := range policyRule.SRCPorts {
				for _, destinationAddress := range policyRule.DSTAddresses {
					for _, destinationPort := range policyRule.DSTPorts {
						aclRules = append(aclRules, makeACLRule(
							policyRule,
							sourceAddress, sourcePort,
							destinationAddress, destinationPort))
					}
				}
			}
		}
	}
	return aclRules
}

func makeACLRule(
	policyRule *models.PolicyRuleType,
	sourceAddress *models.AddressType, sourcePort *models.PortType,
	destinationAddress *models.AddressType, destinationPort *models.PortType,
) *models.AclRuleType {
	return &models.AclRuleType{
		RuleUUID: policyRule.GetRuleUUID(),
		MatchCondition: &models.MatchConditionType{
			Ethertype: policyRule.GetEthertype(),
			Protocol: convertProtocol(policyRule.GetProtocol(), policyRule.GetEthertype()),
			SRCAddress: sourceAddress,
			SRCPort: sourcePort,
			DSTAddress: destinationAddress,
			DSTPort: destinationPort,
		},
		ActionList: &models.ActionListType{
			SimpleAction: "pass",
		},
	}
}

func convertProtocol(policyProtocol string, ethertype string) (aclProtocol string) {
	// TODO Make this work for policyProtocol != "any"
	return policyProtocol
}

func makeACL(name string, rules []*models.AclRuleType) *models.AccessControlList {
	return &models.AccessControlList{
		Name: name,
		AccessControlListEntries: &models.AclEntriesType{
			ACLRule: rules,
		},
	}
}
