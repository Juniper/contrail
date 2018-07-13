package logic

import (
	"fmt"
	"strconv"

	log "github.com/sirupsen/logrus"

	"github.com/Juniper/contrail/pkg/models"
)

func defaultSecurityGroupACLs(
	securityGroup *models.SecurityGroup,
) (ingressACL *models.AccessControlList, egressACL *models.AccessControlList) {
	ingressRules, egressRules := securityGroupToACLRules(securityGroup)

	ingressACL = makeACL("ingress-access-control-list", securityGroup, ingressRules)
	egressACL = makeACL("egress-access-control-list", securityGroup, egressRules)
	return ingressACL, egressACL
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
			Protocol:   policyProtocolToACLProtocol(policyRule.GetProtocol()),
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
func policyProtocolToACLProtocol(policyProtocol string) (aclProtocol string) {
	// TODO(Witaut): Make this work for policyProtocol != "any".
	return policyProtocol
}

func policyAddressToACLAddress(
	policyAddress *models.AddressType, securityGroup *models.SecurityGroup,
) *models.AddressType {
	aclAddress := *policyAddress
	aclAddress.SecurityGroup = securityGroupNameToID(policyAddress.SecurityGroup, securityGroup)
	return &aclAddress
}

func securityGroupNameToID(name string, thisSecurityGroup *models.SecurityGroup) string {
	if name == "local" || name == "" {
		return ""
	} else if fqNameEqualsString(thisSecurityGroup.GetFQName(), name) {
		return strconv.FormatInt(thisSecurityGroup.GetSecurityGroupID(), 10)
	} else {
		// TODO(Witaut): Handle name == "any".
		// TODO(Witaut): If there is a security group in cache with FQName == name,
		// take its SecurityGroupID.
		// TODO(Witaut): Handle the "skip this rule" case.
		return ""
	}
}

func fqNameEqualsString(first []string, second string) bool {
	return models.FQNameToString(first) == second
}

func isIngressRule(sourceAddress *models.AddressType, destinationAddress *models.AddressType) (bool, error) {
	switch {
	case isLocal(sourceAddress) && isLocal(destinationAddress):
		return true, nil
	case isLocal(destinationAddress):
		return true, nil
	case isLocal(sourceAddress):
		return false, nil
	default:
		return false, neitherAddressIsLocal{
			sourceAddress:      sourceAddress,
			destinationAddress: destinationAddress,
		}
	}
}

func isLocal(policyAddress *models.AddressType) bool {
	return policyAddress.SecurityGroup == "local"
}

type neitherAddressIsLocal struct {
	sourceAddress, destinationAddress *models.AddressType
}

func (err neitherAddressIsLocal) Error() string {
	return fmt.Sprintf("neither source nor destination address is local. Source address: %v. Destination address: %v",
		err.sourceAddress, err.destinationAddress)
}

func makeACL(name string, parent *models.SecurityGroup, rules []*models.AclRuleType) *models.AccessControlList {
	return &models.AccessControlList{
		DisplayName: name,
		ParentType:  parent.Kind(),
		ParentUUID:  parent.GetUUID(),
		AccessControlListEntries: &models.AclEntriesType{
			ACLRule: rules,
		},
	}
}
