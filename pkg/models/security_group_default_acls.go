package models

import (
	"strconv"

	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"

	"github.com/Juniper/contrail/pkg/models/basemodels"
)

type securityGroupLoader interface {
	// LoadByFQName loads a Security Group (e.g. from cache) given its FQName.
	LoadByFQName([]string) *SecurityGroup
}

// DefaultACLs returns default ACLs corresponding to the policy rules in a SecurityGroup.
func (m *SecurityGroup) DefaultACLs(l securityGroupLoader) (
	ingressACL *AccessControlList, egressACL *AccessControlList) {

	ingressRules, egressRules := m.toACLRules(l)

	ingressACL = m.makeChildACL("ingress-access-control-list", ingressRules)
	egressACL = m.makeChildACL("egress-access-control-list", egressRules)
	return ingressACL, egressACL
}

func (m *SecurityGroup) toACLRules(l securityGroupLoader) (ingressRules, egressRules []*AclRuleType) {
	for _, pair := range m.allAddressCombinations() {
		rule, err := m.makeACLRule(pair, l)
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

func (m *SecurityGroup) allAddressCombinations() (pairs []policyAddressPair) {
	for _, policyRule := range m.GetSecurityGroupEntries().GetPolicyRule() {
		pairs = append(pairs, policyRule.allAddressCombinations()...)
	}
	return pairs
}

func (m *SecurityGroup) makeACLRule(pair policyAddressPair, l securityGroupLoader) (*AclRuleType, error) {
	protocol, err := pair.policyRule.ACLProtocol()
	if err != nil {
		return nil, err
	}

	sourceAddress, err := m.policyAddressToACLAddress(pair.sourceAddress, l)
	if err != nil {
		return nil, errors.Wrap(err, "failed to convert source address for an ACL")
	}
	destinationAddress, err := m.policyAddressToACLAddress(pair.destinationAddress, l)
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

func (m *SecurityGroup) policyAddressToACLAddress(
	policyAddress *policyAddress, l securityGroupLoader) (*AddressType, error) {

	numericSecurityGroup, err := m.securityGroupNameToID(policyAddress.SecurityGroup, l)
	if err != nil {
		return nil, errors.Wrap(err, "failed to convert security group name for an ACL")
	}

	aclAddress := AddressType(*policyAddress)
	aclAddress.SecurityGroup = numericSecurityGroup
	return &aclAddress, nil
}

func (m *SecurityGroup) securityGroupNameToID(name string, l securityGroupLoader) (string, error) {
	switch {
	case name == "local" || name == "":
		return "", nil
	case name == "any":
		return "-1", nil
	case basemodels.FQNameToString(m.GetFQName()) == name:
		return strconv.FormatInt(m.GetSecurityGroupID(), 10), nil
	default:
		fqName := basemodels.ParseFQName(name)
		sg := l.LoadByFQName(fqName)
		if sg == nil {
			return "", errors.Errorf("unknown security group name %q", name)
		}
		return strconv.FormatInt(sg.GetSecurityGroupID(), 10), nil
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
