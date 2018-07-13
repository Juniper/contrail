package logic

import (
	"context"
	"strconv"

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
	aclRules := securityGroupToACLRules(securityGroup)
	ingressRules, egressRules := splitACLRules(aclRules)

	ingressACL = makeACL("ingress-access-control-list", ingressRules)
	egressACL = makeACL("egress-access-control-list", egressRules)
	return ingressACL, egressACL
}

func splitACLRules(rules []*models.AclRuleType) (ingressRules []*models.AclRuleType, egressRules []*models.AclRuleType) {
	for _, rule := range rules {
		switch {
		case isLocal(cond.SRCAddress) && isLocal(cond.DSTAddress):
			// TODO Make a URL to the commit
			// https://github.com/Juniper/contrail-controller/blob/master/src/config/schema-transformer/config_db.py#L2030
			ingressRules = append(ingressRules, rule)
		case isLocal(cond.DSTAddress):
			ingressRules = append(ingressRules, rule)
		case isLocal(cond.SRCAddress):
			egressRules = append(egressRules, rule)
		default:
			log.WithField("acl_rule", rule).Error(
				"Ignoring ACL rule: neither source nor destination address is local")
		}
	}
	return ingressRules, egressRules
}

func isLocal(aclAddress *models.AddressType) bool {
	return aclAddress.SecurityGroup == ""
}

func securityGroupToACLRules(securityGroup *models.SecurityGroup) (aclRules []*models.AclRuleType) {
	policyRules := securityGroup.SecurityGroupEntries.PolicyRule
	for _, policyRule := range policyRules {
		for _, sourceAddress := range policyRule.SRCAddresses {
			for _, sourcePort := range policyRule.SRCPorts {
				for _, destinationAddress := range policyRule.DSTAddresses {
					for _, destinationPort := range policyRule.DSTPorts {
						aclRules = append(aclRules, makeACLRule(
							securityGroup, policyRule,
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
	securityGroup *models.SecurityGroup, policyRule *models.PolicyRuleType,
	sourceAddress *models.AddressType, sourcePort *models.PortType,
	destinationAddress *models.AddressType, destinationPort *models.PortType,
) *models.AclRuleType {
	return &models.AclRuleType{
		RuleUUID: policyRule.GetRuleUUID(),
		MatchCondition: &models.MatchConditionType{
			Ethertype:  policyRule.GetEthertype(),
			Protocol:   convertProtocol(policyRule.GetProtocol(), policyRule.GetEthertype()),
			SRCAddress: policyAddressToACLAddress(sourceAddress, securityGroup),
			DSTAddress: policyAddressToACLAddress(destinationAddress, securityGroup),
			SRCPort:    sourcePort,
			DSTPort:    destinationPort,
		},
		ActionList: &models.ActionListType{
			SimpleAction: "pass",
		},
	}
}

func policyAddressToACLAddress(address *models.AddressType, securityGroup *models.SecurityGroup) *models.AddressType {
	address = &*address
	address.SecurityGroup = securityGroupFQNameToID(address.SecurityGroup, securityGroup)
	return address
}

func securityGroupFQNameToID(fqName string, thisSecurityGroup *models.SecurityGroup) string {
	// TODO Get the ID of an arbitrary security group from cache
	if models.FQNameToString(thisSecurityGroup.GetFQName()) == fqName {
		return strconv.FormatInt(thisSecurityGroup.GetSecurityGroupID(), 10)
	} else {
		return ""
	}
}

func convertProtocol(policyProtocol string, ethertype string) (aclProtocol string) {
	// TODO Make this work for policyProtocol != "any"
	return policyProtocol
}

func makeACL(name string, rules []*models.AclRuleType) *models.AccessControlList {
	return &models.AccessControlList{
		DisplayName: name,
		AccessControlListEntries: &models.AclEntriesType{
			ACLRule: rules,
		},
	}
}
