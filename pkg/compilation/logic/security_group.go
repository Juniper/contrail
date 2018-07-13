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
	ingressRules, egressRules := securityGroupToACLRules(securityGroup)

	ingressACL = makeACL("ingress-access-control-list", ingressRules)
	egressACL = makeACL("egress-access-control-list", egressRules)
	return ingressACL, egressACL
}

func isLocal(policyAddress *models.AddressType) bool {
	return policyAddress.SecurityGroup == "local"
}

func securityGroupToACLRules(securityGroup *models.SecurityGroup) (ingressRules, egressRules []*models.AclRuleType) {
	policyRules := securityGroup.SecurityGroupEntries.PolicyRule
	for _, policyRule := range policyRules {
		for _, sourceAddress := range policyRule.SRCAddresses {
			for _, sourcePort := range policyRule.SRCPorts {
				for _, destinationAddress := range policyRule.DSTAddresses {
					for _, destinationPort := range policyRule.DSTPorts {
						rule := makeACLRule(
							securityGroup, policyRule,
							sourceAddress, sourcePort,
							destinationAddress, destinationPort)

						isIngress, err := isIngressRule(sourceAddress, destinationAddress)
						if err != nil {
							// TODO Add details
							log.WithError(err).Error("Ignoring ACL rule")
							continue
						}

						if isIngress {
							ingressRules = append(ingressRules, rule)
						} else {
							egressRules = append(egressRules, rule)
						}
					}
				}
			}
		}
	}
	return ingressRules, egressRules
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

func convertProtocol(policyProtocol string, ethertype string) (aclProtocol string) {
	// TODO Make this work for policyProtocol != "any"
	return policyProtocol
}

func policyAddressToACLAddress(
	policyAddress *models.AddressType, securityGroup *models.SecurityGroup,
) *models.AddressType {
	aclAddress := *policyAddress
	aclAddress.SecurityGroup = securityGroupFQNameToID(policyAddress.SecurityGroup, securityGroup)
	return &aclAddress
}

func securityGroupFQNameToID(fqName string, thisSecurityGroup *models.SecurityGroup) string {
	// TODO Get the ID of an arbitrary security group from cache
	if fqNameEqualsString(thisSecurityGroup.GetFQName(), fqName) {
		return strconv.FormatInt(thisSecurityGroup.GetSecurityGroupID(), 10)
	} else if fqName == "local" || fqName == "" {
		return ""
	} else {
		// TODO Implement the rest + !ok
		return ""
	}
}

func fqNameEqualsString(first []string, second string) bool {
	return models.FQNameToString(first) == second
}

func isIngressRule(sourceAddress *models.AddressType, destinationAddress *models.AddressType) (bool, error) {
	switch {
	case isLocal(sourceAddress) && isLocal(destinationAddress):
		// TODO Make a URL to the commit
		// https://github.com/Juniper/contrail-controller/blob/master/src/config/schema-transformer/config_db.py#L2030
		return true, nil
	case isLocal(destinationAddress):
		return true, nil
	case isLocal(sourceAddress):
		return false, nil
	default:
		return false, errors.New("neither source nor destination address is local")
	}
}

func makeACL(name string, rules []*models.AclRuleType) *models.AccessControlList {
	return &models.AccessControlList{
		DisplayName: name,
		AccessControlListEntries: &models.AclEntriesType{
			ACLRule: rules,
		},
	}
}
