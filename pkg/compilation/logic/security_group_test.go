package logic

import (
	"context"
	"testing"

	"github.com/gogo/protobuf/types"
	"github.com/golang/mock/gomock"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"

	"github.com/Juniper/contrail/pkg/compilation/intent"
	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/services"
	"github.com/Juniper/contrail/pkg/services/mock"
	"github.com/Juniper/contrail/pkg/testutil"
	"github.com/Juniper/contrail/pkg/types/mock"
)

func TestCreateSecurityGroupCreatesACLs(t *testing.T) {
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
			Permissions: allPermissions(),
			UUID: &models.UuidType{
				UUIDMslong: 4466401091260269000,
				// Should really be -9223372036854776000, but that overflows int64
				UUIDLslong: 0,
			},
		},
		Perms2: ownerOnlyPerms2(),
		SecurityGroupEntries: &models.PolicyEntriesType{PolicyRule: []*models.PolicyRuleType{
			{
				Direction:    ">",
				Protocol:     "any",
				RuleUUID:     "bdf042c0-d2c2-4241-ba15-1c702c896e03",
				Ethertype:    "IPv4",
				SRCAddresses: []*models.AddressType{securityGroupAddresses()},
				DSTAddresses: []*models.AddressType{localAddresses()},
				SRCPorts:     []*models.PortType{models.AllPorts()},
				DSTPorts:     []*models.PortType{models.AllPorts()},
				ActionList: &models.ActionListType{
					SimpleAction: "pass",
				},
			},
			{
				Direction:    ">",
				Protocol:     "any",
				RuleUUID:     "1f77914a-0863-4f0d-888a-aee6a1988f6a",
				Ethertype:    "IPv6",
				SRCAddresses: []*models.AddressType{securityGroupAddresses()},
				DSTAddresses: []*models.AddressType{localAddresses()},
				SRCPorts:     []*models.PortType{models.AllPorts()},
				DSTPorts:     []*models.PortType{models.AllPorts()},
				ActionList: &models.ActionListType{
					SimpleAction: "pass",
				},
			},
			{
				Direction:    ">",
				Protocol:     "any",
				RuleUUID:     "b7c07625-e03e-43b9-a9fc-d11a6c863cc6",
				Ethertype:    "IPv4",
				SRCAddresses: []*models.AddressType{localAddresses()},
				DSTAddresses: []*models.AddressType{models.AllIPv4Addresses()},
				SRCPorts:     []*models.PortType{models.AllPorts()},
				DSTPorts:     []*models.PortType{models.AllPorts()},
				ActionList: &models.ActionListType{
					SimpleAction: "pass",
				},
			},
			{
				Direction:    ">",
				Protocol:     "any",
				RuleUUID:     "6a5f3026-02bc-4ba1-abde-39abafd21f47",
				Ethertype:    "IPv6",
				SRCAddresses: []*models.AddressType{localAddresses()},
				DSTAddresses: []*models.AddressType{models.AllIPv6Addresses()},
				SRCPorts:     []*models.PortType{models.AllPorts()},
				DSTPorts:     []*models.PortType{models.AllPorts()},
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
						SRCPort:    models.AllPorts(),
						DSTPort:    models.AllPorts(),
						Protocol:   "any",
						Ethertype:  "IPv4",
						SRCAddress: securityGroupIDAddresses("8000002"),
						DSTAddress: noAddresses(),
					},
					ActionList: &models.ActionListType{
						SimpleAction: "pass",
					},
				},
				{
					RuleUUID: "1f77914a-0863-4f0d-888a-aee6a1988f6a",
					MatchCondition: &models.MatchConditionType{
						SRCPort:    models.AllPorts(),
						DSTPort:    models.AllPorts(),
						Protocol:   "any",
						Ethertype:  "IPv6",
						SRCAddress: securityGroupIDAddresses("8000002"),
						DSTAddress: noAddresses(),
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
						SRCPort:    models.AllPorts(),
						DSTPort:    models.AllPorts(),
						Protocol:   "any",
						Ethertype:  "IPv4",
						SRCAddress: noAddresses(),
						DSTAddress: models.AllIPv4Addresses(),
					},
					ActionList: &models.ActionListType{
						SimpleAction: "pass",
					},
				},
				{
					RuleUUID: "6a5f3026-02bc-4ba1-abde-39abafd21f47",
					MatchCondition: &models.MatchConditionType{
						SRCPort:    models.AllPorts(),
						DSTPort:    models.AllPorts(),
						Protocol:   "any",
						Ethertype:  "IPv6",
						SRCAddress: noAddresses(),
						DSTAddress: models.AllIPv6Addresses(),
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

	mockAPIClient := servicesmock.NewMockWriteService(mockCtrl)
	mockIntPoolAllocator := typesmock.NewMockIntPoolAllocator(mockCtrl)
	cache := intent.NewCache()

	service := NewService(mockAPIClient, mockIntPoolAllocator, cache)

	expectCreateACL(mockAPIClient, expectedIngressACL)
	expectCreateACL(mockAPIClient, expectedEgressACL)

	_, err := service.CreateSecurityGroup(context.Background(), &services.CreateSecurityGroupRequest{
		SecurityGroup: securityGroup,
	})
	assert.NoError(t, err)

	assert.Equal(t,
		&SecurityGroupIntent{
			SecurityGroup: securityGroup,
			ingressACL:    expectedIngressACL,
			egressACL:     expectedEgressACL,
		},
		loadSecurityGroupIntent(cache, intent.ByUUID(securityGroup.GetUUID())))
}

func TestSecurityGroupUpdate(t *testing.T) {
	ingressACLAfterCreate := &models.AccessControlList{
		Name:       "ingress-access-control-list",
		ParentType: "security-group",
		ParentUUID: "3dfbd820-e4fc-414f-b1d9-d720ebe93cd8",
		AccessControlListEntries: &models.AclEntriesType{
			Dynamic: false,
			ACLRule: []*models.AclRuleType{
				{
					RuleUUID: "bdf042c0-d2c2-4241-ba15-1c702c896e03",
					MatchCondition: &models.MatchConditionType{
						SRCPort:    models.AllPorts(),
						DSTPort:    models.AllPorts(),
						Protocol:   "any",
						Ethertype:  "IPv4",
						SRCAddress: securityGroupIDAddresses("8000002"),
						DSTAddress: noAddresses(),
					},
					ActionList: &models.ActionListType{
						SimpleAction: "pass",
					},
				},
			},
		},
	}

	egressACLAfterCreate := &models.AccessControlList{
		Name:       "egress-access-control-list",
		ParentType: "security-group",
		ParentUUID: "3dfbd820-e4fc-414f-b1d9-d720ebe93cd8",
		AccessControlListEntries: &models.AclEntriesType{
			Dynamic: false,
		},
	}

	ingressACLAfterUpdate := &models.AccessControlList{
		Name:       "ingress-access-control-list",
		ParentType: "security-group",
		ParentUUID: "3dfbd820-e4fc-414f-b1d9-d720ebe93cd8",
		AccessControlListEntries: &models.AclEntriesType{
			Dynamic: false,
		},
	}

	egressACLAfterUpdate := &models.AccessControlList{
		Name:       "egress-access-control-list",
		ParentType: "security-group",
		ParentUUID: "3dfbd820-e4fc-414f-b1d9-d720ebe93cd8",
		AccessControlListEntries: &models.AclEntriesType{
			Dynamic: false,
			ACLRule: []*models.AclRuleType{
				{
					RuleUUID: "b7c07625-e03e-43b9-a9fc-d11a6c863cc6",
					MatchCondition: &models.MatchConditionType{
						SRCPort:    models.AllPorts(),
						DSTPort:    models.AllPorts(),
						Protocol:   "any",
						Ethertype:  "IPv4",
						SRCAddress: noAddresses(),
						DSTAddress: models.AllIPv4Addresses(),
					},
					ActionList: &models.ActionListType{
						SimpleAction: "pass",
					},
				},
			},
		},
	}

	creSecurityGroup := &models.SecurityGroup{
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
			Permissions: allPermissions(),
			UUID: &models.UuidType{
				UUIDMslong: 4466401091260269000,
				// Should really be -9223372036854776000, but that overflows int64
				UUIDLslong: 0,
			},
		},
		Perms2: ownerOnlyPerms2(),
		SecurityGroupEntries: &models.PolicyEntriesType{PolicyRule: []*models.PolicyRuleType{
			{
				Direction:    ">",
				Protocol:     "any",
				RuleUUID:     "bdf042c0-d2c2-4241-ba15-1c702c896e03",
				Ethertype:    "IPv4",
				SRCAddresses: []*models.AddressType{securityGroupAddresses()},
				DSTAddresses: []*models.AddressType{localAddresses()},
				SRCPorts:     []*models.PortType{models.AllPorts()},
				DSTPorts:     []*models.PortType{models.AllPorts()},
				ActionList: &models.ActionListType{
					SimpleAction: "pass",
				},
			},
		}},
	}

	updSecurityGroup := &models.SecurityGroup{
		UUID: "3dfbd820-e4fc-414f-b1d9-d720ebe93cd8",
		SecurityGroupEntries: &models.PolicyEntriesType{PolicyRule: []*models.PolicyRuleType{
			{
				Direction:    ">",
				Protocol:     "any",
				RuleUUID:     "b7c07625-e03e-43b9-a9fc-d11a6c863cc6",
				Ethertype:    "IPv4",
				SRCAddresses: []*models.AddressType{localAddresses()},
				DSTAddresses: []*models.AddressType{models.AllIPv4Addresses()},
				SRCPorts:     []*models.PortType{models.AllPorts()},
				DSTPorts:     []*models.PortType{models.AllPorts()},
				ActionList: &models.ActionListType{
					SimpleAction: "pass",
				},
			},
		}},
	}

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockAPIClient := servicesmock.NewMockWriteService(mockCtrl)
	mockIntPoolAllocator := typesmock.NewMockIntPoolAllocator(mockCtrl)
	cache := intent.NewCache()

	service := NewService(mockAPIClient, mockIntPoolAllocator, cache)

	expectCreateACL(mockAPIClient, ingressACLAfterCreate)
	expectCreateACL(mockAPIClient, egressACLAfterCreate)
	expectUpdateACL(mockAPIClient, ingressACLAfterUpdate)
	expectUpdateACL(mockAPIClient, egressACLAfterUpdate)

	_, err := service.CreateSecurityGroup(context.Background(), &services.CreateSecurityGroupRequest{
		SecurityGroup: creSecurityGroup,
	})

	assert.Equal(t,
		&SecurityGroupIntent{
			SecurityGroup: creSecurityGroup,
			ingressACL:    ingressACLAfterCreate,
			egressACL:     egressACLAfterCreate,
		},
		loadSecurityGroupIntent(cache, intent.ByUUID(creSecurityGroup.GetUUID())),
	)
	_, err = service.UpdateSecurityGroup(context.Background(), &services.UpdateSecurityGroupRequest{
		SecurityGroup: updSecurityGroup,
	})

	assert.Equal(t,
		&SecurityGroupIntent{
			SecurityGroup: updSecurityGroup,
			ingressACL:    ingressACLAfterUpdate,
			egressACL:     egressACLAfterUpdate,
		},
		loadSecurityGroupIntent(cache, intent.ByUUID(updSecurityGroup.GetUUID())),
	)

	assert.NoError(t, err)
}

func TestSecurityGroupDelete(t *testing.T) {
	testCases := []struct {
		name           string
		intent         *SecurityGroupIntent
		mock           func(*servicesmock.MockWriteService)
		fails          bool
		expectedIntent *SecurityGroupIntent
	}{
		{
			name: "deleting ACLs succeeds",
			intent: &SecurityGroupIntent{
				SecurityGroup: &models.SecurityGroup{UUID: "sg_uuid"},
				ingressACL:    &models.AccessControlList{UUID: "ingress_uuid"},
				egressACL:     &models.AccessControlList{UUID: "egress_uuid"},
			},
			mock: func(mockAPIClient *servicesmock.MockWriteService) {
				expectDeleteACL(mockAPIClient, "ingress_uuid")
				expectDeleteACL(mockAPIClient, "egress_uuid")
			},
			expectedIntent: nil,
		},

		{
			name: "deleting ingress ACL fails",
			intent: &SecurityGroupIntent{
				SecurityGroup: &models.SecurityGroup{UUID: "sg_uuid"},
				ingressACL:    &models.AccessControlList{UUID: "ingress_uuid"},
				egressACL:     &models.AccessControlList{UUID: "egress_uuid"},
			},
			mock: func(mockAPIClient *servicesmock.MockWriteService) {
				expectDeleteACLFailure(mockAPIClient, "ingress_uuid")
				expectDeleteACL(mockAPIClient, "egress_uuid")
			},
			fails: true,
			expectedIntent: &SecurityGroupIntent{
				SecurityGroup: &models.SecurityGroup{UUID: "sg_uuid"},
				ingressACL:    &models.AccessControlList{UUID: "ingress_uuid"},
				egressACL:     nil,
			},
		},

		{
			name: "deleting egress ACL fails",
			intent: &SecurityGroupIntent{
				SecurityGroup: &models.SecurityGroup{UUID: "sg_uuid"},
				ingressACL:    &models.AccessControlList{UUID: "ingress_uuid"},
				egressACL:     &models.AccessControlList{UUID: "egress_uuid"},
			},
			mock: func(mockAPIClient *servicesmock.MockWriteService) {
				expectDeleteACL(mockAPIClient, "ingress_uuid")
				expectDeleteACLFailure(mockAPIClient, "egress_uuid")
			},
			fails: true,
			expectedIntent: &SecurityGroupIntent{
				SecurityGroup: &models.SecurityGroup{UUID: "sg_uuid"},
				ingressACL:    nil,
				egressACL:     &models.AccessControlList{UUID: "egress_uuid"},
			},
		},

		{
			name: "deleting both ACLs fails",
			intent: &SecurityGroupIntent{
				SecurityGroup: &models.SecurityGroup{UUID: "sg_uuid"},
				ingressACL:    &models.AccessControlList{UUID: "ingress_uuid"},
				egressACL:     &models.AccessControlList{UUID: "egress_uuid"},
			},
			mock: func(mockAPIClient *servicesmock.MockWriteService) {
				expectDeleteACLFailure(mockAPIClient, "ingress_uuid")
				expectDeleteACLFailure(mockAPIClient, "egress_uuid")
			},
			fails: true,
			expectedIntent: &SecurityGroupIntent{
				SecurityGroup: &models.SecurityGroup{UUID: "sg_uuid"},
				ingressACL:    &models.AccessControlList{UUID: "ingress_uuid"},
				egressACL:     &models.AccessControlList{UUID: "egress_uuid"},
			},
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()

			mockAPIClient := servicesmock.NewMockWriteService(mockCtrl)
			mockIntPoolAllocator := typesmock.NewMockIntPoolAllocator(mockCtrl)
			cache := intent.NewCache()

			cache.Store(tt.intent)

			service := NewService(mockAPIClient, mockIntPoolAllocator, cache)

			tt.mock(mockAPIClient)

			_, err := service.DeleteSecurityGroup(context.Background(), &services.DeleteSecurityGroupRequest{
				ID: tt.intent.GetUUID(),
			})
			if tt.fails {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}

			assert.Equal(t,
				tt.expectedIntent,
				loadSecurityGroupIntent(cache, intent.ByUUID(tt.intent.GetUUID())))
		})
	}
}

func TestLoadSecurityGroupIntent(t *testing.T) {
	expectedIntent := &SecurityGroupIntent{
		SecurityGroup: &models.SecurityGroup{UUID: "a"},
	}

	cache := intent.NewCache()
	cache.Store(expectedIntent)

	actualIntent := loadSecurityGroupIntent(cache, intent.ByUUID(expectedIntent.UUID))
	assert.Equal(t, expectedIntent, actualIntent)
}

func TestUpdateSecurityGroupIntent(t *testing.T) {
	newIntent := &SecurityGroupIntent{
		SecurityGroup: &models.SecurityGroup{UUID: "a", DisplayName: "before"},
	}

	cache := intent.NewCache()
	cache.Store(newIntent)

	loadedIntent := loadSecurityGroupIntent(cache, intent.ByUUID(newIntent.GetUUID()))
	assert.NotNil(t, loadedIntent)
	loadedIntent.DisplayName = "after"

	assert.Equal(t, "after", newIntent.DisplayName)

	expectedUpdatedIntent := &SecurityGroupIntent{
		SecurityGroup: &models.SecurityGroup{UUID: "a", DisplayName: "after"},
	}

	actualIntent := loadSecurityGroupIntent(cache, intent.ByUUID(expectedUpdatedIntent.GetUUID()))
	assert.NotNil(t, actualIntent)
	assert.Equal(t, expectedUpdatedIntent, actualIntent)
}

func expectCreateACL(mockAPIClient *servicesmock.MockWriteService, expectedACL *models.AccessControlList) {
	mockAPIClient.EXPECT().CreateAccessControlList(testutil.NotNil(), &services.CreateAccessControlListRequest{
		AccessControlList: expectedACL,
	}).Return(&services.CreateAccessControlListResponse{
		AccessControlList: expectedACL,
	}, nil).Times(1)
}

func expectUpdateACL(mockAPIClient *servicesmock.MockWriteService, expectedACL *models.AccessControlList) {
	mockAPIClient.EXPECT().UpdateAccessControlList(testutil.NotNil(), &services.UpdateAccessControlListRequest{
		AccessControlList: expectedACL,
		FieldMask:         types.FieldMask{Paths: []string{models.AccessControlListFieldAccessControlListEntries}},
	}).Return(&services.UpdateAccessControlListResponse{
		AccessControlList: expectedACL,
	}, nil).Times(1)
}

func expectDeleteACL(mockAPIClient *servicesmock.MockWriteService, expectedUUID string) {
	mockAPIClient.EXPECT().DeleteAccessControlList(testutil.NotNil(), &services.DeleteAccessControlListRequest{
		ID: expectedUUID,
	}).Return(&services.DeleteAccessControlListResponse{
		ID: expectedUUID,
	}, nil).Times(1)
}

func expectDeleteACLFailure(mockAPIClient *servicesmock.MockWriteService, expectedUUID string) {
	mockAPIClient.EXPECT().DeleteAccessControlList(testutil.NotNil(), &services.DeleteAccessControlListRequest{
		ID: expectedUUID,
	}).Return(nil, errors.New("failed to delete the ACL for some reason")).Times(1)
}

func allPermissions() *models.PermType {
	return &models.PermType{
		Owner:       "cloud-admin",
		Group:       "cloud-admin-group",
		OwnerAccess: 7,
		OtherAccess: 7,
		GroupAccess: 7,
	}
}

func ownerOnlyPerms2() *models.PermType2 {
	return &models.PermType2{
		Owner:        "950b2912-a742-47c8-acdb-ab361f173867",
		OwnerAccess:  7,
		GlobalAccess: 0,
	}
}

func localAddresses() *models.AddressType {
	return &models.AddressType{
		SecurityGroup: "local",
	}
}

func securityGroupAddresses() *models.AddressType {
	return &models.AddressType{
		SecurityGroup: "default-domain:project-blue:default",
	}
}

func securityGroupIDAddresses(id string) *models.AddressType {
	return &models.AddressType{
		SecurityGroup: id,
	}
}

func noAddresses() *models.AddressType {
	return &models.AddressType{}
}
