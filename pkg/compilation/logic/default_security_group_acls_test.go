package logic

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/Juniper/contrail/pkg/models"
)

func TestMakeACLRule(t *testing.T) {
	testCases := []struct {
		name                              string
		securityGroup                     *models.SecurityGroup
		policyRule                        *models.PolicyRuleType
		sourceAddress, destinationAddress *models.AddressType
		sourcePort, destinationPort       *models.PortType

		expectedACLRule *models.AclRuleType
	}{
		{
			name: "IPv4, specified security group to local security group",
			securityGroup: &models.SecurityGroup{
				FQName:          []string{"default-domain", "project-blue", "default"},
				SecurityGroupID: 8000002,
			},
			policyRule: &models.PolicyRuleType{
				RuleUUID:  "bdf042c0-d2c2-4241-ba15-1c702c896e03",
				Direction: ">",
				Protocol:  "any",
				Ethertype: "IPv4",
			},
			sourceAddress: &models.AddressType{
				SecurityGroup: "default-domain:project-blue:default",
			},
			destinationAddress: &models.AddressType{
				SecurityGroup: "local",
			},
			sourcePort:      allPorts(),
			destinationPort: allPorts(),

			expectedACLRule: &models.AclRuleType{
				RuleUUID: "bdf042c0-d2c2-4241-ba15-1c702c896e03",
				MatchCondition: &models.MatchConditionType{
					SRCPort:   allPorts(),
					DSTPort:   allPorts(),
					Protocol:  "any",
					Ethertype: "IPv4",
					SRCAddress: &models.AddressType{
						SecurityGroup: "8000002",
					},
					DSTAddress: &models.AddressType{},
				},
				ActionList: &models.ActionListType{
					SimpleAction: "pass",
				},
			},
		},

		{
			name: "IPv6, specified security group to local security group",
			securityGroup: &models.SecurityGroup{
				FQName:          []string{"default-domain", "project-blue", "default"},
				SecurityGroupID: 8000002,
			},
			policyRule: &models.PolicyRuleType{
				RuleUUID:  "1f77914a-0863-4f0d-888a-aee6a1988f6a",
				Direction: ">",
				Protocol:  "any",
				Ethertype: "IPv6",
			},
			sourceAddress: &models.AddressType{
				SecurityGroup: "default-domain:project-blue:default",
			},
			destinationAddress: &models.AddressType{
				SecurityGroup: "local",
			},
			sourcePort:      allPorts(),
			destinationPort: allPorts(),

			expectedACLRule: &models.AclRuleType{
				RuleUUID: "1f77914a-0863-4f0d-888a-aee6a1988f6a",
				MatchCondition: &models.MatchConditionType{
					SRCPort:   allPorts(),
					DSTPort:   allPorts(),
					Protocol:  "any",
					Ethertype: "IPv6",
					SRCAddress: &models.AddressType{
						SecurityGroup: "8000002",
					},
					DSTAddress: &models.AddressType{},
				},
				ActionList: &models.ActionListType{
					SimpleAction: "pass",
				},
			},
		},

		{
			name: "IPv4, local security group to all addresses",
			securityGroup: &models.SecurityGroup{
				FQName:          []string{"default-domain", "project-blue", "default"},
				SecurityGroupID: 8000002,
			},
			policyRule: &models.PolicyRuleType{
				RuleUUID:  "b7c07625-e03e-43b9-a9fc-d11a6c863cc6",
				Direction: ">",
				Protocol:  "any",
				Ethertype: "IPv4",
			},
			sourceAddress: &models.AddressType{
				SecurityGroup: "local",
			},
			destinationAddress: allV4Addresses(),
			sourcePort:         allPorts(),
			destinationPort:    allPorts(),

			expectedACLRule: &models.AclRuleType{
				RuleUUID: "b7c07625-e03e-43b9-a9fc-d11a6c863cc6",
				MatchCondition: &models.MatchConditionType{
					SRCPort:    allPorts(),
					DSTPort:    allPorts(),
					Protocol:   "any",
					Ethertype:  "IPv4",
					SRCAddress: &models.AddressType{},
					DSTAddress: allV4Addresses(),
				},
				ActionList: &models.ActionListType{
					SimpleAction: "pass",
				},
			},
		},

		{
			name: "IPv6, local security group to all addresses",
			securityGroup: &models.SecurityGroup{
				FQName:          []string{"default-domain", "project-blue", "default"},
				SecurityGroupID: 8000002,
			},
			policyRule: &models.PolicyRuleType{
				RuleUUID:  "6a5f3026-02bc-4ba1-abde-39abafd21f47",
				Direction: ">",
				Protocol:  "any",
				Ethertype: "IPv6",
			},
			sourceAddress: &models.AddressType{
				SecurityGroup: "local",
			},
			destinationAddress: allV6Addresses(),
			sourcePort:         allPorts(),
			destinationPort:    allPorts(),

			expectedACLRule: &models.AclRuleType{
				RuleUUID: "6a5f3026-02bc-4ba1-abde-39abafd21f47",
				MatchCondition: &models.MatchConditionType{
					SRCPort:    allPorts(),
					DSTPort:    allPorts(),
					Protocol:   "any",
					Ethertype:  "IPv6",
					SRCAddress: &models.AddressType{},
					DSTAddress: allV6Addresses(),
				},
				ActionList: &models.ActionListType{
					SimpleAction: "pass",
				},
			},
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			aclRule := makeACLRule(
				tt.securityGroup, tt.policyRule,
				tt.sourceAddress, tt.sourcePort,
				tt.destinationAddress, tt.destinationPort)
			assert.Equal(t, tt.expectedACLRule, aclRule)
		})
	}
}

func TestIsIngress(t *testing.T) {
	testCases := []struct {
		name                              string
		sourceAddress, destinationAddress *models.AddressType
		isIngress                         bool
		err                               error
	}{
		{
			name: "specified security group to local security group",
			sourceAddress: &models.AddressType{
				SecurityGroup: "default-domain:project-blue:default",
			},
			destinationAddress: &models.AddressType{
				SecurityGroup: "local",
			},
			isIngress: true,
		},
		{
			name: "local security group to all IPv4 addresses",
			sourceAddress: &models.AddressType{
				SecurityGroup: "local",
			},
			destinationAddress: allV4Addresses(),
			isIngress:          false,
		},
		{
			name: "local security group to all IPv6 addresses",
			sourceAddress: &models.AddressType{
				SecurityGroup: "local",
			},
			destinationAddress: allV6Addresses(),
			isIngress:          false,
		},
		{
			name: "both with local security group",
			sourceAddress: &models.AddressType{
				SecurityGroup: "local",
			},
			destinationAddress: &models.AddressType{
				SecurityGroup: "local",
			},
			// https://github.com/Juniper/contrail-controller/blob/08f2b11d3/src/config/schema-transformer/config_db.py#L2030
			isIngress: true,
		},
		{
			name:               "neither with local security group",
			sourceAddress:      &models.AddressType{},
			destinationAddress: &models.AddressType{},
			err: neitherAddressIsLocal{
				sourceAddress:      &models.AddressType{},
				destinationAddress: &models.AddressType{},
			},
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			isIngress, err := isIngressRule(tt.sourceAddress, tt.destinationAddress)
			if tt.err != nil {
				assert.Equal(t, tt.err, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.isIngress, isIngress)
			}
		})
	}
}

func TestIsLocal(t *testing.T) {
	testCases := []struct {
		name    string
		address *models.AddressType
		is      bool
	}{
		{
			name: "local security group",
			address: &models.AddressType{
				SecurityGroup: "local",
			},
			is: true,
		},
		{
			name: "specified security group",
			address: &models.AddressType{
				SecurityGroup: "default-domain:project-blue:default",
			},
			is: false,
		},
		{
			name:    "all IPv4 addresses",
			address: allV4Addresses(),
			is:      false,
		},
		{
			name:    "all IPv6 addresses",
			address: allV6Addresses(),
			is:      false,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.is, isLocal(tt.address))
		})
	}
}

func allV6Addresses() *models.AddressType {
	return &models.AddressType{
		Subnet: &models.SubnetType{
			IPPrefix:    "::",
			IPPrefixLen: 0,
		},
	}
}

func allV4Addresses() *models.AddressType {
	return &models.AddressType{
		Subnet: &models.SubnetType{
			IPPrefix:    "0.0.0.0",
			IPPrefixLen: 0,
		},
	}
}

func allPorts() *models.PortType {
	return &models.PortType{
		StartPort: 0,
		EndPort:   65535,
	}
}
