package types

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/models/basemodels"
	"github.com/Juniper/contrail/pkg/services"
	"github.com/Juniper/contrail/pkg/services/mock"
	"github.com/Juniper/contrail/pkg/types/mock"
)

//Structure expectedParams is used to store some expected values
type expectedParams struct {
	MatchTags     *models.FirewallRuleMatchTagsType
	MatchTagTypes *models.FirewallRuleMatchTagsTypeIdList
	Service       *models.FirewallServiceType
}

func updateExpectedFirewallRule(fr models.FirewallRule, params expectedParams) *models.FirewallRule {
	fr.MatchTags = params.MatchTags
	fr.MatchTagTypes = params.MatchTagTypes
	fr.Service = params.Service
	return &fr
}

func firewallRuleSetupMocks(s *ContrailTypeLogicService) {
	nextService := s.Next().(*servicesmock.MockService)
	readService := s.ReadService.(*servicesmock.MockReadService)

	nextService.EXPECT().CreateFirewallRule(
		gomock.Not(gomock.Nil()), gomock.Not(gomock.Nil()),
	).DoAndReturn(
		func(_ context.Context, request *services.CreateFirewallRuleRequest) (
			response *services.CreateFirewallRuleResponse, err error,
		) {
			return &services.CreateFirewallRuleResponse{FirewallRule: request.FirewallRule}, nil
		},
	).AnyTimes()

	s.MetadataGetter.(*typesmock.MockMetadataGetter).EXPECT().GetMetadata(
		gomock.Not(gomock.Nil()), gomock.Not(gomock.Nil()),
	).Return(
		&basemodels.Metadata{
			UUID: "tag-type-uuid",
		},
		nil,
	).AnyTimes()

	readService.EXPECT().GetTagType(gomock.Not(gomock.Nil()), gomock.Not(gomock.Nil())).Return(
		&services.GetTagTypeResponse{
			TagType: &models.TagType{
				TagTypeID: "121",
			},
		},
		nil,
	).AnyTimes()
}

func TestCreateFirewallRule(t *testing.T) {
	tests := []struct {
		name             string
		testFirewallRule models.FirewallRule
		expectedParams   expectedParams
		errorCode        codes.Code
	}{
		{
			name: "Try to create when address-group refs are scoped",
			testFirewallRule: models.FirewallRule{
				UUID:   "test-firewall-rule",
				FQName: []string{"default-policy-management", "test-firewall-rule"},
				AddressGroupRefs: []*models.FirewallRuleAddressGroupRef{
					{
						UUID: "address-group-ref-uuid",
						To:   []string{"address-group-ref-uuid"},
					},
				},
			},
			errorCode: codes.InvalidArgument,
		},
		{
			name: "Try to create when service-group refs are scoped",
			testFirewallRule: models.FirewallRule{
				UUID:   "test-firewall-rule",
				FQName: []string{"default-policy-management", "test-firewall-rule"},
				ServiceGroupRefs: []*models.FirewallRuleServiceGroupRef{
					{
						UUID: "service-group-ref-uuid",
						To:   []string{"service-group-ref-uuid"},
					},
				},
			},
			errorCode: codes.InvalidArgument,
		},
		{
			name: "Try to create when virtual-network refs are scoped",
			testFirewallRule: models.FirewallRule{
				UUID:   "test-firewall-rule",
				FQName: []string{"default-policy-management", "test-firewall-rule"},
				VirtualNetworkRefs: []*models.FirewallRuleVirtualNetworkRef{
					{
						UUID: "virtual-network-ref-uuid",
						To:   []string{"virtual-network-ref-uuid"},
					},
				},
			},
			errorCode: codes.InvalidArgument,
		},
		{
			name: "Try to create when service and service-group-ref properties are undefined",
			testFirewallRule: models.FirewallRule{
				UUID:   "test-firewall-rule",
				FQName: []string{"default-policy-management", "test-firewall-rule"},
				VirtualNetworkRefs: []*models.FirewallRuleVirtualNetworkRef{
					{
						UUID: "virtual-network-ref-uuid",
						To:   []string{"default-project-uuid", "virtual-network-ref-uuid"},
					},
				},
			},
			errorCode: codes.InvalidArgument,
		},
		{
			name: "Try to create when service and service-group-ref properties are defined",
			testFirewallRule: models.FirewallRule{
				Service:          &models.FirewallServiceType{},
				ServiceGroupRefs: []*models.FirewallRuleServiceGroupRef{{}},
			},
			errorCode: codes.InvalidArgument,
		},
		{
			name: "Try to create with invalid int protocol name ",
			testFirewallRule: models.FirewallRule{
				Service: &models.FirewallServiceType{
					Protocol: "-1",
				},
			},
			errorCode: codes.InvalidArgument,
		},
		{
			name: "Try to create with invalid string protocol name",
			testFirewallRule: models.FirewallRule{
				Service: &models.FirewallServiceType{
					Protocol: "none",
				},
			},
			errorCode: codes.InvalidArgument,
		},
		{
			name: "Try to create when multiple endpoint types enabled",
			testFirewallRule: models.FirewallRule{
				Service: &models.FirewallServiceType{
					Protocol: "icmp",
				},
				Endpoint1: &models.FirewallRuleEndpointType{
					Any:            true,
					VirtualNetwork: "virtual-network-uuid",
					Tags:           []string{"namespace=default"},
				},
			},
			errorCode: codes.InvalidArgument,
		},
		{
			name: "Try to create with label match-tag",
			testFirewallRule: models.FirewallRule{
				Service: &models.FirewallServiceType{
					Protocol: "tcp",
				},
				MatchTags: &models.FirewallRuleMatchTagsType{
					TagList: []string{"Label"},
				},
			},
			errorCode: codes.InvalidArgument,
		},
		{
			name: "Create firewall-rule with default match-tag propetry",
			testFirewallRule: models.FirewallRule{
				Service: &models.FirewallServiceType{
					Protocol: "tcp",
				},
			},
			expectedParams: expectedParams{
				MatchTags: &models.FirewallRuleMatchTagsType{
					TagList: []string{"application"},
				},
				MatchTagTypes: &models.FirewallRuleMatchTagsTypeIdList{
					TagType: []int64{1},
				},
				Service: &models.FirewallServiceType{
					Protocol:   "tcp",
					ProtocolID: 6,
				},
			},
		},
		{
			name: "Create firewall-rule match-tag propetry defined",
			testFirewallRule: models.FirewallRule{
				Service: &models.FirewallServiceType{
					Protocol: "240",
				},
				MatchTags: &models.FirewallRuleMatchTagsType{
					TagList: []string{"Application", "Tier", "OtherTag"},
				},
			},
			expectedParams: expectedParams{
				MatchTags: &models.FirewallRuleMatchTagsType{
					TagList: []string{"Application", "Tier", "OtherTag"},
				},
				MatchTagTypes: &models.FirewallRuleMatchTagsTypeIdList{
					TagType: []int64{1, 2, 121},
				},
				Service: &models.FirewallServiceType{
					Protocol:   "240",
					ProtocolID: 240,
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()
			service := makeMockedContrailTypeLogicService(mockCtrl)
			firewallRuleSetupMocks(service)

			ctx := context.Background()

			paramRequest := &services.CreateFirewallRuleRequest{FirewallRule: &tt.testFirewallRule}
			expectedResponse := &services.CreateFirewallRuleResponse{
				FirewallRule: updateExpectedFirewallRule(tt.testFirewallRule, tt.expectedParams),
			}
			createFirewallRuleResponse, err := service.CreateFirewallRule(ctx, paramRequest)
			if tt.errorCode != codes.OK {
				assert.Error(t, err)
				status, ok := status.FromError(err)
				assert.True(t, ok)
				assert.Equal(t, tt.errorCode, status.Code())
			} else {
				assert.NoError(t, err)
				assert.Equal(t, expectedResponse, createFirewallRuleResponse)
			}
		})
	}
}
