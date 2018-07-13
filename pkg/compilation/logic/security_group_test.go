package logic

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/services"
	"github.com/Juniper/contrail/pkg/services/mock"
)

func TestCreateSecurityGroupCreatesACLs(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockAPIService := servicesmock.NewMockService(mockCtrl)
	service := NewService(mockAPIService)

	allPermissions := &models.PermType{
		Owner:       "cloud-admin",
		Group:       "cloud-admin-group",
		OwnerAccess: 7,
		OtherAccess: 7,
		GroupAccess: 7,
	}
	ownerOnlyPerms2 := &models.PermType2{
		Owner:        "950b2912-a742-47c8-acdb-ab361f173867",
		OwnerAccess:  7,
		GlobalAccess: 0,
	}

	localAddresses := &models.AddressType{
		SecurityGroup: "local",
	}
	securityGroupAddresses := &models.AddressType{
		SecurityGroup: "default-domain:project-blue:default",
	}
	securityGroupIDAddresses := &models.AddressType{
		SecurityGroup: "8000002",
	}
	noAddresses := &models.AddressType{}

	securityGroup := &models.SecurityGroup{
		ParentUUID:      "950b2912-a742-47c8-acdb-ab361f173867",
		ParentType:      "project",
		FQName:          []string{"default-domain", "project-blue", "default"},
		DisplayName:     "default",
		SecurityGroupID: 8000002,
		IDPerms: &models.IdPermsType{
			Enable:      true,
			Description: "Default security group",
			UserVisible: true,
			Permissions: allPermissions,
			UUID: &models.UuidType{
				UUIDMslong: 4466401091260269000,
				// Should really be -9223372036854776000, but that overflows int64
				UUIDLslong: 0,
			},
		},
		Perms2: ownerOnlyPerms2,
		SecurityGroupEntries: &models.PolicyEntriesType{PolicyRule: []*models.PolicyRuleType{
			{
				Direction:    ">",
				Protocol:     "any",
				RuleUUID:     "bdf042c0-d2c2-4241-ba15-1c702c896e03",
				Ethertype:    "IPv4",
				SRCAddresses: []*models.AddressType{securityGroupAddresses},
				DSTAddresses: []*models.AddressType{localAddresses},
				SRCPorts:     []*models.PortType{allPorts()},
				DSTPorts:     []*models.PortType{allPorts()},
			},
			{
				Direction:    ">",
				Protocol:     "any",
				RuleUUID:     "1f77914a-0863-4f0d-888a-aee6a1988f6a",
				Ethertype:    "IPv6",
				SRCAddresses: []*models.AddressType{securityGroupAddresses},
				DSTAddresses: []*models.AddressType{localAddresses},
				SRCPorts:     []*models.PortType{allPorts()},
				DSTPorts:     []*models.PortType{allPorts()},
			},
			{
				Direction:    ">",
				Protocol:     "any",
				RuleUUID:     "b7c07625-e03e-43b9-a9fc-d11a6c863cc6",
				Ethertype:    "IPv4",
				SRCAddresses: []*models.AddressType{localAddresses},
				DSTAddresses: []*models.AddressType{allV4Addresses()},
				SRCPorts:     []*models.PortType{allPorts()},
				DSTPorts:     []*models.PortType{allPorts()},
			},
			{
				Direction:    ">",
				Protocol:     "any",
				RuleUUID:     "6a5f3026-02bc-4ba1-abde-39abafd21f47",
				Ethertype:    "IPv6",
				SRCAddresses: []*models.AddressType{localAddresses},
				DSTAddresses: []*models.AddressType{allV6Addresses()},
				SRCPorts:     []*models.PortType{allPorts()},
				DSTPorts:     []*models.PortType{allPorts()},
			},
		}},
	}

	expectedIngressACL := &models.AccessControlList{
		DisplayName: "ingress-access-control-list",
		/*
			ParentType:  "security-group",
			FQName:      []string{"default-domain", "project-blue", "default", "ingress-access-control-list"},
			IDPerms: &models.IdPermsType{
				Enable:      true,
				UserVisible: true,
				Permissions: allPermissions,
				UUID: &models.UuidType{
					UUIDMslong: 8323894507853989000,
					// Should really be -9223372036854776000, but that overflows int64
					UUIDLslong: 0,
				},
			},
			Perms2: ownerOnlyPerms2,
		*/
		AccessControlListEntries: &models.AclEntriesType{
			Dynamic: false,
			ACLRule: []*models.AclRuleType{
				{
					RuleUUID: "bdf042c0-d2c2-4241-ba15-1c702c896e03",
					MatchCondition: &models.MatchConditionType{
						SRCPort:    allPorts(),
						DSTPort:    allPorts(),
						Protocol:   "any",
						Ethertype:  "IPv4",
						SRCAddress: securityGroupIDAddresses,
						DSTAddress: noAddresses,
					},
					ActionList: &models.ActionListType{
						SimpleAction: "pass",
					},
				},
				{
					RuleUUID: "1f77914a-0863-4f0d-888a-aee6a1988f6a",
					MatchCondition: &models.MatchConditionType{
						SRCPort:    allPorts(),
						DSTPort:    allPorts(),
						Protocol:   "any",
						Ethertype:  "IPv6",
						SRCAddress: securityGroupIDAddresses,
						DSTAddress: noAddresses,
					},
					ActionList: &models.ActionListType{
						SimpleAction: "pass",
					},
				},
			},
		},
	}

	expectedEgressACL := &models.AccessControlList{
		DisplayName: "egress-access-control-list",
		/*
			ParentType:  "security-group",
			FQName:      []string{"default-domain", "project-blue", "default", "egress-access-control-list"},
			IDPerms: &models.IdPermsType{
				Enable:      true,
				UserVisible: true,
				Permissions: allPermissions,
				UUID: &models.UuidType{
					// Should really be -9223372036854776000, but that overflows int64
					UUIDMslong: 0,
					// Should really be -9223372036854776000, but that overflows int64
					UUIDLslong: 0,
				},
			},
			Perms2: ownerOnlyPerms2,
		*/
		AccessControlListEntries: &models.AclEntriesType{
			Dynamic: false,
			ACLRule: []*models.AclRuleType{
				{
					RuleUUID: "b7c07625-e03e-43b9-a9fc-d11a6c863cc6",
					MatchCondition: &models.MatchConditionType{
						SRCPort:    allPorts(),
						DSTPort:    allPorts(),
						Protocol:   "any",
						Ethertype:  "IPv4",
						SRCAddress: noAddresses,
						DSTAddress: allV4Addresses(),
					},
					ActionList: &models.ActionListType{
						SimpleAction: "pass",
					},
				},
				{
					RuleUUID: "6a5f3026-02bc-4ba1-abde-39abafd21f47",
					MatchCondition: &models.MatchConditionType{
						SRCPort:    allPorts(),
						DSTPort:    allPorts(),
						Protocol:   "any",
						Ethertype:  "IPv6",
						SRCAddress: noAddresses,
						DSTAddress: allV6Addresses(),
					},
					ActionList: &models.ActionListType{
						SimpleAction: "pass",
					},
				},
			},
		},
	}

	expectCreateACL(mockAPIService, expectedIngressACL)
	expectCreateACL(mockAPIService, expectedEgressACL)

	_, err := service.CreateSecurityGroup(context.Background(), &services.CreateSecurityGroupRequest{
		SecurityGroup: securityGroup,
	})
	assert.NoError(t, err)
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

func expectCreateACL(mockAPIService *servicesmock.MockService, expectedACL *models.AccessControlList) {
	mockAPIService.EXPECT().CreateAccessControlList(notNil(), &services.CreateAccessControlListRequest{
		AccessControlList: expectedACL,
	}).Return(&services.CreateAccessControlListResponse{
		AccessControlList: expectedACL,
	}, nil).Times(1)
}

func notNil() gomock.Matcher {
	return gomock.Not(gomock.Nil())
}

func allPorts() *models.PortType {
	return &models.PortType{
		StartPort: 0,
		EndPort:   65535,
	}
}

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
		fails                             bool
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
			isIngress: true,
		},
		{
			name:               "neither with local security group",
			sourceAddress:      &models.AddressType{},
			destinationAddress: &models.AddressType{},
			fails:              true,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			isIngress, err := isIngressRule(tt.sourceAddress, tt.destinationAddress)
			if tt.fails {
				assert.Error(t, err)
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
