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
	for _, pair := range allAddressCombinations(policyRules) {
		rule := makeACLRule(securityGroup, pair)

		isIngress, err := isIngressRule(pair.sourceAddress, pair.destinationAddress)
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
	return ingressRules, egressRules
}

type policyAddressPair struct {
	policyRule                        *models.PolicyRuleType
	sourceAddress, destinationAddress *models.AddressType
	sourcePort, destinationPort       *models.PortType
}

func allAddressCombinations(policyRules []*models.PolicyRuleType) (pairs []policyAddressPair) {
	for _, policyRule := range policyRules {
		pairs = append(pairs, allRuleAddressCombinations(policyRule)...)
	}
	return pairs
}

func allRuleAddressCombinations(policyRule *models.PolicyRuleType) (pairs []policyAddressPair) {
	for _, sourceAddress := range policyRule.SRCAddresses {
		for _, sourcePort := range policyRule.SRCPorts {
			for _, destinationAddress := range policyRule.DSTAddresses {
				for _, destinationPort := range policyRule.DSTPorts {
					pairs = append(pairs, policyAddressPair{
						policyRule: policyRule,

						sourceAddress:      sourceAddress,
						sourcePort:         sourcePort,
						destinationAddress: destinationAddress,
						destinationPort:    destinationPort,
					})
				}
			}
		}
	}
	return pairs
}

func makeACLRule(securityGroup *models.SecurityGroup, pair policyAddressPair) *models.AclRuleType {
	return &models.AclRuleType{
		RuleUUID: pair.policyRule.GetRuleUUID(),
		MatchCondition: &models.MatchConditionType{
			Ethertype:  pair.policyRule.GetEthertype(),
			Protocol:   policyProtocolToACLProtocol(pair.policyRule.GetProtocol()),
			SRCAddress: policyAddressToACLAddress(pair.sourceAddress, securityGroup),
			DSTAddress: policyAddressToACLAddress(pair.destinationAddress, securityGroup),
			SRCPort:    pair.sourcePort,
			DSTPort:    pair.destinationPort,
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
		Name:       name,
		ParentType: parent.Kind(),
		ParentUUID: parent.GetUUID(),
		AccessControlListEntries: &models.AclEntriesType{
			ACLRule: rules,
		},
	}
}
