package logic

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/services"
	"github.com/Juniper/contrail/pkg/services/mock"
	"github.com/Juniper/contrail/pkg/testutil"
)

func TestCreateSecurityGroupCreatesACLs(t *testing.T) {
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

	// https://review.opencontrail.org/#/c/44499/19/pkg/compilation/testdata/security_group.yml
	// with default values stripped.
	securityGroup := &models.SecurityGroup{
		UUID:            "3dfbd820-e4fc-414f-b1d9-d720ebe93cd8",
		ParentType:      "project",
		ParentUUID:      "950b2912-a742-47c8-acdb-ab361f173867",
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
				ActionList: &models.ActionListType{
					SimpleAction: "pass",
				},
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
				ActionList: &models.ActionListType{
					SimpleAction: "pass",
				},
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
				ActionList: &models.ActionListType{
					SimpleAction: "pass",
				},
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
				ActionList: &models.ActionListType{
					SimpleAction: "pass",
				},
			},
		}},
	}

	expectedIngressACL := &models.AccessControlList{
		Name:       "ingress-access-control-list",
		ParentType: "security-group",
		ParentUUID: "3dfbd820-e4fc-414f-b1d9-d720ebe93cd8",
		// FQName, IDPerms, Perms2 omitted,
		// as they should be filled by the API server.
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
		Name:       "egress-access-control-list",
		ParentType: "security-group",
		ParentUUID: "3dfbd820-e4fc-414f-b1d9-d720ebe93cd8",
		// FQName, IDPerms, Perms2 omitted,
		// as they should be filled by the API server.
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

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockAPIService := servicesmock.NewMockWriteService(mockCtrl)
	service := NewService(mockAPIService)

	expectCreateACL(mockAPIService, expectedIngressACL)
	expectCreateACL(mockAPIService, expectedEgressACL)

	_, err := service.CreateSecurityGroup(context.Background(), &services.CreateSecurityGroupRequest{
		SecurityGroup: securityGroup,
	})
	assert.NoError(t, err)
}

func expectCreateACL(mockAPIService *servicesmock.MockWriteService, expectedACL *models.AccessControlList) {
	mockAPIService.EXPECT().CreateAccessControlList(testutil.NotNil(), &services.CreateAccessControlListRequest{
		AccessControlList: expectedACL,
	}).Return(&services.CreateAccessControlListResponse{
		AccessControlList: expectedACL,
	}, nil).Times(1)
}
