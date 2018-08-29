package models

import (
	"strconv"

	"github.com/Juniper/contrail/pkg/models/basemodels"
	"github.com/sirupsen/logrus"
)

// DefaultACLs returns default ACLs corresponding to the policy rules in a SecurityGroup.
func (m *SecurityGroup) DefaultACLs() (ingressACL *AccessControlList, egressACL *AccessControlList) {
	ingressRules, egressRules := m.toACLRules()

	ingressACL = m.makeChildACL("ingress-access-control-list", ingressRules)
	egressACL = m.makeChildACL("egress-access-control-list", egressRules)
	return ingressACL, egressACL
}

func (m *SecurityGroup) toACLRules() (ingressRules, egressRules []*AclRuleType) {
	for _, pair := range m.allAddressCombinations() {
		rule, err := m.makeACLRule(pair)
		if err != nil {
			logrus.WithError(err).Error("Ignoring ACL rule")
			continue
		}

		isIngress, err := pair.isIngress()
		if err != nil {
			logrus.WithError(err).Error("Ignoring ACL rule")
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

func (m *SecurityGroup) allAddressCombinations() (pairs []policyAddressPair) {
	for _, policyRule := range m.GetSecurityGroupEntries().GetPolicyRule() {
		pairs = append(pairs, policyRule.allAddressCombinations()...)
	}
	return pairs
}

func (m *SecurityGroup) makeACLRule(pair policyAddressPair) (*AclRuleType, error) {
	protocol, err := pair.policyRule.ACLProtocol()
	if err != nil {
		return nil, err
	}

	return &AclRuleType{
		RuleUUID: pair.policyRule.GetRuleUUID(),
		MatchCondition: &MatchConditionType{
			Ethertype:  pair.policyRule.GetEthertype(),
			Protocol:   protocol,
			SRCAddress: m.policyAddressToACLAddress(pair.sourceAddress),
			DSTAddress: m.policyAddressToACLAddress(pair.destinationAddress),
			SRCPort:    pair.sourcePort,
			DSTPort:    pair.destinationPort,
		},
		ActionList: &ActionListType{
			SimpleAction: "pass",
		},
	}, nil
}

func (m *SecurityGroup) policyAddressToACLAddress(policyAddress *policyAddress) *AddressType {
	aclAddress := AddressType(*policyAddress)
	aclAddress.SecurityGroup = m.securityGroupNameToID(policyAddress.SecurityGroup)
	return &aclAddress
}

func (m *SecurityGroup) securityGroupNameToID(name string) string {
	switch {
	case name == "local" || name == "":
		return ""
	case basemodels.FQNameToString(m.GetFQName()) == name:
		return strconv.FormatInt(m.GetSecurityGroupID(), 10)
	default:
		// TODO: Handle name == "any".
		// TODO: If there is a security group in cache with FQName == name, take its SecurityGroupID.
		// TODO: Handle the "skip this rule" case.
		return ""
	}
}

func (m *SecurityGroup) makeChildACL(name string, rules []*AclRuleType) *AccessControlList {
	return &AccessControlList{
		Name:       name,
		ParentType: m.Kind(),
		ParentUUID: m.GetUUID(),
		AccessControlListEntries: &AclEntriesType{
			ACLRule: rules,
		},
	}
}
