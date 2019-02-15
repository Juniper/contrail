package models

import (
	"github.com/Juniper/contrail/pkg/models/basemodels"
	"github.com/twinj/uuid"
)

// DefaultSecurityGroupName is the Name of a project's default SecurityGroup.
const (
	DefaultSecurityGroupName = "default"
)

// DefaultSecurityGroup returns the default SecurityGroup for the project.
func (m *Project) DefaultSecurityGroup() *SecurityGroup {
	thisSecurityGroup := basemodels.FQNameToString(m.DefaultSecurityGroupFQName())
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
				MakeSecurityGroupPolicyRule(true, IPv4Ethertype, &AddressType{SecurityGroup: thisSecurityGroup}),
				MakeSecurityGroupPolicyRule(true, IPv6Ethertype, &AddressType{SecurityGroup: thisSecurityGroup}),
				MakeSecurityGroupPolicyRule(false, IPv4Ethertype, &AddressType{
					Subnet: &SubnetType{
						IPPrefix:    IPv4ZeroValue,
						IPPrefixLen: 0,
					},
				}),
				MakeSecurityGroupPolicyRule(false, IPv6Ethertype, &AddressType{
					Subnet: &SubnetType{
						IPPrefix:    IPv6ZeroValue,
						IPPrefixLen: 0,
					},
				}),
			},
		},
	}
}

// DefaultSecurityGroupFQName returns the FQName of the project's default SecurityGroup.
func (m *Project) DefaultSecurityGroupFQName() []string {
	return basemodels.ChildFQName(m.GetFQName(), DefaultSecurityGroupName)
}

// MakeSecurityGroupPolicyRule makes a policy rule for a SecurityGroup.
func MakeSecurityGroupPolicyRule(
	ingress bool,
	ethertype string,
	remoteAddr *AddressType,
) *PolicyRuleType {
	rule := &PolicyRuleType{
		RuleUUID:  uuid.NewV4().String(),
		Direction: SRCToDSTDirection,
		Ethertype: ethertype,
		Protocol:  AnyProtocol,
		SRCPorts:  []*PortType{AllPorts()},
		DSTPorts:  []*PortType{AllPorts()},
	}

	localAddr := &AddressType{
		SecurityGroup: LocalSecurityGroup,
	}

	if ingress {
		rule.SRCAddresses = []*AddressType{remoteAddr}
		rule.DSTAddresses = []*AddressType{localAddr}
	} else {
		rule.SRCAddresses = []*AddressType{localAddr}
		rule.DSTAddresses = []*AddressType{remoteAddr}
	}

	return rule
}
