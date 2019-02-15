package models

import (
	"strings"

	"github.com/Juniper/contrail/pkg/models/basemodels"
	"github.com/twinj/uuid"
)

// DefaultSecurityGroupName is the Name of a project's default SecurityGroup.
const (
	DefaultSecurityGroupName = "default"
)

// DefaultSecurityGroup returns the default SecurityGroup for the project.
func (m *Project) DefaultSecurityGroup() *SecurityGroup {
	fqNameString := basemodels.FQNameToString(m.GetFQName())
	return &SecurityGroup{
		Name:       DefaultSecurityGroupName,
		ParentUUID: m.GetUUID(),
		ParentType: KindProject,
		IDPerms: &IdPermsType{
			Enable:      true,
			Description: "Default security group",
		},
		SecurityGroupEntries: &PolicyEntriesType{
			PolicyRule: []*PolicyRuleType{
				MakeSecurityGroupPolicyRule(true, DefaultSecurityGroupName, "", IPv4Ethertype, fqNameString),
				MakeSecurityGroupPolicyRule(true, DefaultSecurityGroupName, "", IPv6Ethertype, fqNameString),
				MakeSecurityGroupPolicyRule(false, UnspecifiedSecurityGroup, IPv4ZeroValue, IPv4Ethertype, fqNameString),
				MakeSecurityGroupPolicyRule(false, UnspecifiedSecurityGroup, IPv6ZeroValue, IPv6Ethertype, fqNameString),
			},
		},
	}
}

// MakeSecurityGroupPolicyRule makes a policy rule for a SecurityGroup.
func MakeSecurityGroupPolicyRule(
	ingress bool,
	securityGroup string,
	prefix string,
	ethertype string,
	projectFQNameString string,
) *PolicyRuleType {

	uuid := uuid.NewV4().String()
	localAddr := AddressType{
		SecurityGroup: LocalSecurityGroup,
	}

	var addr AddressType
	if securityGroup != "" {
		addr = AddressType{
			SecurityGroup: strings.Join([]string{projectFQNameString, securityGroup}, ":"),
		}
	}

	if prefix != "" {
		addr = AddressType{
			Subnet: &SubnetType{
				IPPrefix:    prefix,
				IPPrefixLen: 0,
			},
		}
	}

	rule := PolicyRuleType{
		RuleUUID:  uuid,
		Direction: SRCToDSTDirection,
		Protocol:  AnyProtocol,
		SRCPorts:  []*PortType{AllPorts()},
		DSTPorts:  []*PortType{AllPorts()},
		Ethertype: ethertype,
	}
	if ingress {
		rule.SRCAddresses = []*AddressType{&addr}
		rule.DSTAddresses = []*AddressType{&localAddr}
		return &rule
	}

	rule.SRCAddresses = []*AddressType{&localAddr}
	rule.DSTAddresses = []*AddressType{&addr}
	return &rule
}
