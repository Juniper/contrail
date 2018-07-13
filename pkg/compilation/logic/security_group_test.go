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
	allPorts := []*models.PortType{
		{
			EndPort:   65535,
			StartPort: 0,
		},
	}

	allV4Addresses := &models.AddressType{
		Subnet: &models.SubnetType{
			IPPrefix:    "0.0.0.0",
			IPPrefixLen: 0,
		},
	}
	allV6Addresses := &models.AddressType{
		Subnet: &models.SubnetType{
			IPPrefix:    "::",
			IPPrefixLen: 0,
		},
	}
	noLocalAddresses := []*models.AddressType{
		{
			SecurityGroup: "local",
			SubnetList:    []*models.SubnetType{},
		},
	}
	noSecurityGroupAddresses := []*models.AddressType{
		{
			SecurityGroup: "default-domain:project-blue:default",
			SubnetList:    []*models.SubnetType{},
		},
	}

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
				SRCAddresses: noSecurityGroupAddresses,
				SRCPorts:     allPorts,
				DSTAddresses: noLocalAddresses,
				DSTPorts:     allPorts,
			},
			{
				Direction:    ">",
				Protocol:     "any",
				RuleUUID:     "1f77914a-0863-4f0d-888a-aee6a1988f6a",
				Ethertype:    "IPv6",
				SRCAddresses: noSecurityGroupAddresses,
				SRCPorts:     allPorts,
				DSTAddresses: noLocalAddresses,
				DSTPorts:     allPorts,
			},
			{
				Direction:    ">",
				Protocol:     "any",
				RuleUUID:     "b7c07625-e03e-43b9-a9fc-d11a6c863cc6",
				Ethertype:    "IPv4",
				SRCAddresses: noLocalAddresses,
				SRCPorts:     allPorts,
				DSTAddresses: []*models.AddressType{allV4Addresses},
				DSTPorts:     allPorts,
			},
			{
				Direction:    ">",
				Protocol:     "any",
				RuleUUID:     "6a5f3026-02bc-4ba1-abde-39abafd21f47",
				Ethertype:    "IPv6",
				SRCAddresses: noLocalAddresses,
				SRCPorts:     allPorts,
				DSTAddresses: []*models.AddressType{allV6Addresses},
				DSTPorts:     allPorts,
			},
		}},
	}

	expectedIngressACL := &models.AccessControlList{
		DisplayName: "ingress-access-control-list",
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
		AccessControlListEntries: &models.AclEntriesType{
			Dynamic: false,
			ACLRule: []*models.AclRuleType{
				{
					RuleUUID: "bdf042c0-d2c2-4241-ba15-1c702c896e03",
					MatchCondition: &models.MatchConditionType{
						SRCPort:    allPorts[0],
						DSTPort:    allPorts[0],
						Protocol:   "any",
						Ethertype:  "IPv4",
						SRCAddress: &models.AddressType{
							SecurityGroup: "8000002",
						},
						DSTAddress: &models.AddressType{},
					},
					ActionList: &models.ActionListType{
						SimpleAction: "pass",
					},
				},
				{
					RuleUUID: "1f77914a-0863-4f0d-888a-aee6a1988f6a",
					MatchCondition: &models.MatchConditionType{
						SRCPort:    allPorts[0],
						DSTPort:    allPorts[0],
						Protocol:   "any",
						Ethertype:  "IPv6",
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
		},
	}

	expectedEgressACL := &models.AccessControlList{
		DisplayName: "egress-access-control-list",
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
		AccessControlListEntries: &models.AclEntriesType{
			Dynamic: false,
			ACLRule: []*models.AclRuleType{
				{
					RuleUUID: "b7c07625-e03e-43b9-a9fc-d11a6c863cc6",
					MatchCondition: &models.MatchConditionType{
						SRCPort:    allPorts[0],
						DSTPort:    allPorts[0],
						Protocol:   "any",
						Ethertype:  "IPv4",
						SRCAddress: &models.AddressType{},
						DSTAddress: allV4Addresses,
					},
					ActionList: &models.ActionListType{
						SimpleAction: "pass",
					},
				},
				{
					RuleUUID: "6a5f3026-02bc-4ba1-abde-39abafd21f47",
					MatchCondition: &models.MatchConditionType{
						SRCPort:    allPorts[0],
						DSTPort:    allPorts[0],
						Protocol:   "any",
						Ethertype:  "IPv6",
						SRCAddress: &models.AddressType{},
						DSTAddress: allV6Addresses,
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

func expectCreateACL(mockAPIService *servicesmock.MockService, expectedACL *models.AccessControlList) {
	mockAPIService.EXPECT().CreateAccessControlList(notNil(), &services.CreateAccessControlListRequest{
		AccessControlList: expectedACL,
	}).Return(&services.CreateAccessControlListResponse{
		AccessControlList: expectedACL,
	}, nil)
}

func notNil() gomock.Matcher {
	return gomock.Not(gomock.Nil())
}
