package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToACLRules(t *testing.T) {
	testCases := []struct {
		name          string
		securityGroup *SecurityGroup

		expectedIngressACLRules []*AclRuleType
		expectedEgressACLRules  []*AclRuleType
	}{
		{
			// Behave properly, unlike
			// https://github.com/Juniper/contrail-controller/blob/be4053c84/src/config/schema-transformer/config_db.py#L2030
			name: "Non-local destination address after a local destination address",
			securityGroup: &SecurityGroup{
				FQName:          []string{"default-domain", "project-blue", "default"},
				SecurityGroupID: 8000002,
				SecurityGroupEntries: &PolicyEntriesType{PolicyRule: []*PolicyRuleType{
					{
						Direction: ">",
						Protocol:  "any",
						RuleUUID:  "rule1",
						Ethertype: "IPv4",
						SRCAddresses: []*AddressType{
							{
								SecurityGroup: "local",
							},
						},
						DSTAddresses: []*AddressType{
							AllIPv4Addresses(),
							{
								SecurityGroup: "local",
							},
							AllIPv4Addresses(),
						},
						SRCPorts: []*PortType{AllPorts()},
						DSTPorts: []*PortType{AllPorts()},
					},
				}},
			},

			expectedIngressACLRules: []*AclRuleType{
				{
					RuleUUID: "rule1",
					MatchCondition: &MatchConditionType{
						SRCPort:    AllPorts(),
						DSTPort:    AllPorts(),
						Protocol:   "any",
						Ethertype:  "IPv4",
						SRCAddress: &AddressType{},
						DSTAddress: &AddressType{},
					},
					ActionList: &ActionListType{
						SimpleAction: "pass",
					},
				},
			},

			expectedEgressACLRules: []*AclRuleType{
				{
					RuleUUID: "rule1",
					MatchCondition: &MatchConditionType{
						SRCPort:    AllPorts(),
						DSTPort:    AllPorts(),
						Protocol:   "any",
						Ethertype:  "IPv4",
						SRCAddress: &AddressType{},
						DSTAddress: AllIPv4Addresses(),
					},
					ActionList: &ActionListType{
						SimpleAction: "pass",
					},
				},
				{
					RuleUUID: "rule1",
					MatchCondition: &MatchConditionType{
						SRCPort:    AllPorts(),
						DSTPort:    AllPorts(),
						Protocol:   "any",
						Ethertype:  "IPv4",
						SRCAddress: &AddressType{},
						DSTAddress: AllIPv4Addresses(),
					},
					ActionList: &ActionListType{
						SimpleAction: "pass",
					},
				},
			},
		},

		{
			// Behave properly, unlike
			// https://github.com/Juniper/contrail-controller/blob/be4053c84/src/config/schema-transformer/config_db.py#L2030
			name: "Non-local source & destination addresses after a local source address",
			securityGroup: &SecurityGroup{
				FQName:          []string{"default-domain", "project-blue", "default"},
				SecurityGroupID: 8000002,
				SecurityGroupEntries: &PolicyEntriesType{PolicyRule: []*PolicyRuleType{
					{
						Direction: ">",
						Protocol:  "any",
						RuleUUID:  "rule1",
						Ethertype: "IPv4",
						SRCAddresses: []*AddressType{
							{
								SecurityGroup: "local",
							},
							AllIPv4Addresses(),
						},
						DSTAddresses: []*AddressType{
							AllIPv4Addresses(),
						},
						SRCPorts: []*PortType{AllPorts()},
						DSTPorts: []*PortType{AllPorts()},
					},
				}},
			},

			expectedIngressACLRules: nil,

			expectedEgressACLRules: []*AclRuleType{
				{
					RuleUUID: "rule1",
					MatchCondition: &MatchConditionType{
						SRCPort:    AllPorts(),
						DSTPort:    AllPorts(),
						Protocol:   "any",
						Ethertype:  "IPv4",
						SRCAddress: &AddressType{},
						DSTAddress: AllIPv4Addresses(),
					},
					ActionList: &ActionListType{
						SimpleAction: "pass",
					},
				},
			},
		},

		{
			// Behave properly, unlike
			// https://github.com/Juniper/contrail-controller/blob/be4053c84/src/config/schema-transformer/config_db.py#L2030
			name: "Non-local source & destination addresses after a local destination address",
			securityGroup: &SecurityGroup{
				FQName:          []string{"default-domain", "project-blue", "default"},
				SecurityGroupID: 8000002,
				SecurityGroupEntries: &PolicyEntriesType{PolicyRule: []*PolicyRuleType{
					{
						Direction: ">",
						Protocol:  "any",
						RuleUUID:  "rule1",
						Ethertype: "IPv4",
						SRCAddresses: []*AddressType{
							{
								SecurityGroup: "local",
							},
							AllIPv4Addresses(),
						},
						DSTAddresses: []*AddressType{
							AllIPv4Addresses(),
						},
						SRCPorts: []*PortType{AllPorts()},
						DSTPorts: []*PortType{AllPorts()},
					},
				}},
			},

			expectedIngressACLRules: nil,

			expectedEgressACLRules: []*AclRuleType{
				{
					RuleUUID: "rule1",
					MatchCondition: &MatchConditionType{
						SRCPort:    AllPorts(),
						DSTPort:    AllPorts(),
						Protocol:   "any",
						Ethertype:  "IPv4",
						SRCAddress: &AddressType{},
						DSTAddress: AllIPv4Addresses(),
					},
					ActionList: &ActionListType{
						SimpleAction: "pass",
					},
				},
			},
		},

		{
			name: "Unknown IPv4 protocol in the only rule",
			securityGroup: &SecurityGroup{
				FQName:          []string{"default-domain", "project-blue", "default"},
				SecurityGroupID: 8000002,
				SecurityGroupEntries: &PolicyEntriesType{PolicyRule: []*PolicyRuleType{
					{
						Direction: ">",
						Protocol:  "some unknown protocol",
						RuleUUID:  "rule1",
						Ethertype: "IPv4",
						SRCAddresses: []*AddressType{
							{
								SecurityGroup: "local",
							},
						},
						DSTAddresses: []*AddressType{
							AllIPv4Addresses(),
							{
								SecurityGroup: "local",
							},
						},
						SRCPorts: []*PortType{AllPorts()},
						DSTPorts: []*PortType{AllPorts()},
					},
				}},
			},

			expectedIngressACLRules: nil,
			expectedEgressACLRules:  nil,
		},

		{
			name: "Unknown IPv4 protocol in one of the rules",
			securityGroup: &SecurityGroup{
				FQName:          []string{"default-domain", "project-blue", "default"},
				SecurityGroupID: 8000002,
				SecurityGroupEntries: &PolicyEntriesType{PolicyRule: []*PolicyRuleType{
					{
						Protocol:  "unknown protocol 1",
						RuleUUID:  "rule1",
						Ethertype: "IPv4",
						SRCAddresses: []*AddressType{
							{
								SecurityGroup: "local",
							},
						},
						DSTAddresses: []*AddressType{
							AllIPv4Addresses(),
							{
								SecurityGroup: "local",
							},
						},
						SRCPorts: []*PortType{AllPorts()},
						DSTPorts: []*PortType{AllPorts()},
					},
					{
						Direction: ">",
						Protocol:  "any",
						RuleUUID:  "rule2",
						Ethertype: "IPv6",
						SRCAddresses: []*AddressType{
							{
								SecurityGroup: "local",
							},
						},
						DSTAddresses: []*AddressType{
							AllIPv6Addresses(),
							{
								SecurityGroup: "local",
							},
						},
						SRCPorts: []*PortType{AllPorts()},
						DSTPorts: []*PortType{AllPorts()},
					},
				}},
			},

			expectedIngressACLRules: []*AclRuleType{
				{
					RuleUUID: "rule2",
					MatchCondition: &MatchConditionType{
						SRCPort:    AllPorts(),
						DSTPort:    AllPorts(),
						Protocol:   "any",
						Ethertype:  "IPv6",
						SRCAddress: &AddressType{},
						DSTAddress: &AddressType{},
					},
					ActionList: &ActionListType{
						SimpleAction: "pass",
					},
				},
			},
			expectedEgressACLRules: []*AclRuleType{
				{
					RuleUUID: "rule2",
					MatchCondition: &MatchConditionType{
						SRCPort:    AllPorts(),
						DSTPort:    AllPorts(),
						Protocol:   "any",
						Ethertype:  "IPv6",
						SRCAddress: &AddressType{},
						DSTAddress: AllIPv6Addresses(),
					},
					ActionList: &ActionListType{
						SimpleAction: "pass",
					},
				},
			},
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			ingressACLRules, egressACLRules := tt.securityGroup.toACLRules()
			assert.Equal(t, tt.expectedIngressACLRules, ingressACLRules,
				"ingress ACL rules don't match the expected ones")
			assert.Equal(t, tt.expectedEgressACLRules, egressACLRules,
				"egress ACL rules don't match the expected ones")
		})
	}
}

func TestMakeACLRule(t *testing.T) {
	testCases := []struct {
		name          string
		securityGroup *SecurityGroup
		policyAddressPair

		expectedACLRule *AclRuleType
		expectedError   error
	}{
		{
			name: "IPv4, specified security group to local security group",
			securityGroup: &SecurityGroup{
				FQName:          []string{"default-domain", "project-blue", "default"},
				SecurityGroupID: 8000002,
			},
			policyAddressPair: policyAddressPair{
				policyRule: &PolicyRuleType{
					RuleUUID:  "bdf042c0-d2c2-4241-ba15-1c702c896e03",
					Direction: ">",
					Protocol:  "any",
					Ethertype: "IPv4",
				},
				sourceAddress: &policyAddress{
					SecurityGroup: "default-domain:project-blue:default",
				},
				destinationAddress: &policyAddress{
					SecurityGroup: "local",
				},
				sourcePort:      AllPorts(),
				destinationPort: AllPorts(),
			},

			expectedACLRule: &AclRuleType{
				RuleUUID: "bdf042c0-d2c2-4241-ba15-1c702c896e03",
				MatchCondition: &MatchConditionType{
					SRCPort:   AllPorts(),
					DSTPort:   AllPorts(),
					Protocol:  "any",
					Ethertype: "IPv4",
					SRCAddress: &AddressType{
						SecurityGroup: "8000002",
					},
					DSTAddress: &AddressType{},
				},
				ActionList: &ActionListType{
					SimpleAction: "pass",
				},
			},
		},

		{
			name: "IPv6, specified security group to local security group",
			securityGroup: &SecurityGroup{
				FQName:          []string{"default-domain", "project-blue", "default"},
				SecurityGroupID: 8000002,
			},
			policyAddressPair: policyAddressPair{
				policyRule: &PolicyRuleType{
					RuleUUID:  "1f77914a-0863-4f0d-888a-aee6a1988f6a",
					Direction: ">",
					Protocol:  "any",
					Ethertype: "IPv6",
				},
				sourceAddress: &policyAddress{
					SecurityGroup: "default-domain:project-blue:default",
				},
				destinationAddress: &policyAddress{
					SecurityGroup: "local",
				},
				sourcePort:      AllPorts(),
				destinationPort: AllPorts(),
			},

			expectedACLRule: &AclRuleType{
				RuleUUID: "1f77914a-0863-4f0d-888a-aee6a1988f6a",
				MatchCondition: &MatchConditionType{
					SRCPort:   AllPorts(),
					DSTPort:   AllPorts(),
					Protocol:  "any",
					Ethertype: "IPv6",
					SRCAddress: &AddressType{
						SecurityGroup: "8000002",
					},
					DSTAddress: &AddressType{},
				},
				ActionList: &ActionListType{
					SimpleAction: "pass",
				},
			},
		},

		{
			name: "IPv4, local security group to all addresses",
			securityGroup: &SecurityGroup{
				FQName:          []string{"default-domain", "project-blue", "default"},
				SecurityGroupID: 8000002,
			},
			policyAddressPair: policyAddressPair{
				policyRule: &PolicyRuleType{
					RuleUUID:  "b7c07625-e03e-43b9-a9fc-d11a6c863cc6",
					Direction: ">",
					Protocol:  "any",
					Ethertype: "IPv4",
				},
				sourceAddress: &policyAddress{
					SecurityGroup: "local",
				},
				destinationAddress: (*policyAddress)(AllIPv4Addresses()),
				sourcePort:         AllPorts(),
				destinationPort:    AllPorts(),
			},

			expectedACLRule: &AclRuleType{
				RuleUUID: "b7c07625-e03e-43b9-a9fc-d11a6c863cc6",
				MatchCondition: &MatchConditionType{
					SRCPort:    AllPorts(),
					DSTPort:    AllPorts(),
					Protocol:   "any",
					Ethertype:  "IPv4",
					SRCAddress: &AddressType{},
					DSTAddress: AllIPv4Addresses(),
				},
				ActionList: &ActionListType{
					SimpleAction: "pass",
				},
			},
		},

		{
			name: "IPv6, local security group to all addresses",
			securityGroup: &SecurityGroup{
				FQName:          []string{"default-domain", "project-blue", "default"},
				SecurityGroupID: 8000002,
			},
			policyAddressPair: policyAddressPair{
				policyRule: &PolicyRuleType{
					RuleUUID:  "6a5f3026-02bc-4ba1-abde-39abafd21f47",
					Direction: ">",
					Protocol:  "any",
					Ethertype: "IPv6",
				},
				sourceAddress: &policyAddress{
					SecurityGroup: "local",
				},
				destinationAddress: (*policyAddress)(AllIPv6Addresses()),
				sourcePort:         AllPorts(),
				destinationPort:    AllPorts(),
			},

			expectedACLRule: &AclRuleType{
				RuleUUID: "6a5f3026-02bc-4ba1-abde-39abafd21f47",
				MatchCondition: &MatchConditionType{
					SRCPort:    AllPorts(),
					DSTPort:    AllPorts(),
					Protocol:   "any",
					Ethertype:  "IPv6",
					SRCAddress: &AddressType{},
					DSTAddress: AllIPv6Addresses(),
				},
				ActionList: &ActionListType{
					SimpleAction: "pass",
				},
			},
		},

		{
			// Replicates the logic in
			// https://github.com/Juniper/contrail-controller/blob/474731ce0/src/config/schema-transformer/config_db.py#L2039
			name: "ActionList with a deny action (should be ignored)",
			securityGroup: &SecurityGroup{
				FQName:          []string{"default-domain", "project-blue", "default"},
				SecurityGroupID: 8000002,
			},
			policyAddressPair: policyAddressPair{
				policyRule: &PolicyRuleType{
					RuleUUID:  "rule2",
					Direction: ">",
					Protocol:  "any",
					Ethertype: "IPv4",
					ActionList: &ActionListType{
						SimpleAction: "deny",
					},
				},
				sourceAddress: &policyAddress{
					SecurityGroup: "local",
				},
				destinationAddress: (*policyAddress)(AllIPv4Addresses()),
				sourcePort:         AllPorts(),
				destinationPort:    AllPorts(),
			},

			expectedACLRule: &AclRuleType{
				RuleUUID: "rule2",
				MatchCondition: &MatchConditionType{
					SRCPort:    AllPorts(),
					DSTPort:    AllPorts(),
					Protocol:   "any",
					Ethertype:  "IPv4",
					SRCAddress: &AddressType{},
					DSTAddress: AllIPv4Addresses(),
				},
				ActionList: &ActionListType{
					SimpleAction: "pass",
				},
			},
		},

		{
			name: "IPv4, unknown protocol",
			policyAddressPair: policyAddressPair{
				policyRule: &PolicyRuleType{
					Protocol:  "some unknown protocol",
					Ethertype: "IPv4",
				},
			},

			expectedError: unknownProtocol{
				protocol:  "some unknown protocol",
				ethertype: "IPv4",
			},
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			aclRule, err := tt.securityGroup.makeACLRule(tt.policyAddressPair)
			assert.Equal(t, tt.expectedError, err)
			assert.Equal(t, tt.expectedACLRule, aclRule)
		})
	}
}

func TestPolicyProtocolToACLProtocol(t *testing.T) {
	testCases := []struct {
		name                string
		policyRule          *PolicyRuleType
		expectedACLProtocol string
		expectedError       error
	}{
		{
			name: "any",
			policyRule: &PolicyRuleType{
				Protocol: "any",
			},
			expectedACLProtocol: "any",
		},

		{
			name: "not specified",
			policyRule: &PolicyRuleType{
				Protocol: "",
			},
			expectedACLProtocol: "",
		},

		{
			name: "already a number",
			policyRule: &PolicyRuleType{
				Protocol: "58",
			},
			expectedACLProtocol: "58",
		},

		{
			name: "unknown IPv6 protocol",
			policyRule: &PolicyRuleType{
				Protocol:  "some unknown protocol",
				Ethertype: "IPv6",
			},
			expectedError: unknownProtocol{
				protocol:  "some unknown protocol",
				ethertype: "IPv6",
			},
		},

		{
			name: "unknown IPv4 protocol",
			policyRule: &PolicyRuleType{
				Protocol:  "some unknown protocol",
				Ethertype: "IPv4",
			},
			expectedError: unknownProtocol{
				protocol:  "some unknown protocol",
				ethertype: "IPv4",
			},
		},

		{
			name: "unknown ethertype and protocol",
			policyRule: &PolicyRuleType{
				Protocol:  "some unknown protocol",
				Ethertype: "some unknown ethertype",
			},
			expectedError: unknownProtocol{
				protocol:  "some unknown protocol",
				ethertype: "some unknown ethertype",
			},
		},

		{
			name: "icmp ipv6",
			policyRule: &PolicyRuleType{
				Protocol:  "icmp",
				Ethertype: "IPv6",
			},
			expectedACLProtocol: "58",
		},

		{
			name: "icmp ipv4",
			policyRule: &PolicyRuleType{
				Protocol:  "icmp",
				Ethertype: "IPv4",
			},
			expectedACLProtocol: "1",
		},

		// The rest of the tests are the same for IPv6 and IPv4
		{
			name: "icmp6 ipv6",
			policyRule: &PolicyRuleType{
				Protocol:  "icmp6",
				Ethertype: "IPv6",
			},
			expectedACLProtocol: "58",
		},

		{
			name: "icmp6 ipv4",
			policyRule: &PolicyRuleType{
				Protocol:  "icmp6",
				Ethertype: "IPv4",
			},
			expectedACLProtocol: "58",
		},

		{
			name: "tcp ipv6",
			policyRule: &PolicyRuleType{
				Protocol:  "tcp",
				Ethertype: "IPv6",
			},
			expectedACLProtocol: "6",
		},

		{
			name: "tcp ipv4",
			policyRule: &PolicyRuleType{
				Protocol:  "tcp",
				Ethertype: "IPv4",
			},
			expectedACLProtocol: "6",
		},

		{
			name: "udp ipv6",
			policyRule: &PolicyRuleType{
				Protocol:  "udp",
				Ethertype: "IPv6",
			},
			expectedACLProtocol: "17",
		},

		{
			name: "udp ipv4",
			policyRule: &PolicyRuleType{
				Protocol:  "udp",
				Ethertype: "IPv4",
			},
			expectedACLProtocol: "17",
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			aclProtocol, err := tt.policyRule.ACLProtocol()
			assert.Equal(t, tt.expectedError, err)
			assert.Equal(t, tt.expectedACLProtocol, aclProtocol)
		})
	}
}
