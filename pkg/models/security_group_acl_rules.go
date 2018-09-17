package models

import (
	"strconv"

	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

// PolicyRulesWithRefs is a list of policy rules together with resolved references
// to security groups whose names are mentioned in the rules.
type PolicyRulesWithRefs struct {
	Rules      []*PolicyRuleType
	FQNameToSG map[string]*SecurityGroup
}

// ToACLRules translates policy rules to ACL rules.
func (rs *PolicyRulesWithRefs) ToACLRules() (ingressRules, egressRules []*AclRuleType) {
	for _, pair := range allAddressCombinations(rs.Rules) {
		rule, err := makeACLRule(pair, rs.FQNameToSG)
		if err != nil {
			log.WithError(err).Error("Ignoring ACL rule")
			continue
		}

		isIngress, err := pair.isIngress()
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

// MakeChildACL returns a child ACL for a security group with the given ACL rules.
func (m *SecurityGroup) MakeChildACL(name string, rules []*AclRuleType) *AccessControlList {
	return &AccessControlList{
		Name:       name,
		ParentType: m.Kind(),
		ParentUUID: m.GetUUID(),
		AccessControlListEntries: &AclEntriesType{
			ACLRule: rules,
		},
	}
}

func allAddressCombinations(rs []*PolicyRuleType) (pairs []policyAddressPair) {
	for _, r := range rs {
		pairs = append(pairs, r.allAddressCombinations()...)
	}
	return pairs
}

func makeACLRule(pair policyAddressPair, fqNameToSG map[string]*SecurityGroup) (*AclRuleType, error) {
	protocol, err := pair.policyRule.ACLProtocol()
	if err != nil {
		return nil, err
	}

	sourceAddress, err := policyAddressToACLAddress(pair.sourceAddress, fqNameToSG)
	if err != nil {
		return nil, errors.Wrap(err, "failed to convert source address for an ACL")
	}
	destinationAddress, err := policyAddressToACLAddress(pair.destinationAddress, fqNameToSG)
	if err != nil {
		return nil, errors.Wrap(err, "failed to convert destination address for an ACL")
	}

	return &AclRuleType{
		RuleUUID: pair.policyRule.GetRuleUUID(),
		MatchCondition: &MatchConditionType{
			Ethertype:  pair.policyRule.GetEthertype(),
			Protocol:   protocol,
			SRCAddress: sourceAddress,
			DSTAddress: destinationAddress,
			SRCPort:    pair.sourcePort,
			DSTPort:    pair.destinationPort,
		},
		ActionList: &ActionListType{
			SimpleAction: "pass",
		},
	}, nil
}

func policyAddressToACLAddress(
	policyAddress *policyAddress, fqNameToSG map[string]*SecurityGroup) (*AddressType, error) {

	numericSecurityGroup, err := securityGroupNameToID(policyAddress.SecurityGroup, fqNameToSG)
	if err != nil {
		return nil, errors.Wrap(err, "failed to convert security group name for an ACL")
	}

	aclAddress := AddressType(*policyAddress)
	aclAddress.SecurityGroup = numericSecurityGroup
	return &aclAddress, nil
}

func securityGroupNameToID(name string, fqNameToSG map[string]*SecurityGroup) (string, error) {
	switch {
	case name == LocalSecurityGroup || name == UnspecifiedSecurityGroup:
		return "", nil
	case name == AnySecurityGroup:
		return "-1", nil
	default:
		sg := fqNameToSG[name]
		if sg == nil {
			return "", errors.Errorf("unknown security group name %q", name)
		}
		return strconv.FormatInt(sg.GetSecurityGroupID(), 10), nil
	}
}
