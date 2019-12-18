package types

import (
	"context"
	"testing"

	"github.com/gogo/protobuf/types"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/Juniper/asf/pkg/errutil"
	"github.com/Juniper/asf/pkg/models/basemodels"
	"github.com/Juniper/contrail/pkg/auth"
	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/services"
	servicesmock "github.com/Juniper/contrail/pkg/services/mock"
	typesmock "github.com/Juniper/contrail/pkg/types/mock"
)

//Structure expectedParams is used to store some expected values
type expectedParams struct {
	MatchTags          *models.FirewallRuleMatchTagsType
	MatchTagTypes      *models.FirewallRuleMatchTagsTypeIdList
	Service            *models.FirewallServiceType
	Endpoint1          *models.FirewallRuleEndpointType
	Endpoint2          *models.FirewallRuleEndpointType
	TagRefs            []*models.FirewallRuleTagRef
	AddressGroupRefs   []*models.FirewallRuleAddressGroupRef
	VirtualNetworkRefs []*models.FirewallRuleVirtualNetworkRef
}

func updateExpectedFirewallRule(fr models.FirewallRule, params expectedParams) *models.FirewallRule {
	fr.MatchTags = params.MatchTags
	fr.MatchTagTypes = params.MatchTagTypes
	fr.Service = params.Service
	fr.Endpoint1 = params.Endpoint1
	fr.Endpoint2 = params.Endpoint2
	fr.TagRefs = params.TagRefs
	fr.AddressGroupRefs = params.AddressGroupRefs
	fr.VirtualNetworkRefs = params.VirtualNetworkRefs
	return &fr
}
func firewallRuleSetupMocks(s *ContrailTypeLogicService, databaseFR *models.FirewallRule) {
	nextService := s.Next().(*servicesmock.MockService) //nolint: errcheck

	nextService.EXPECT().CreateFirewallRule(
		gomock.Not(gomock.Nil()), gomock.Not(gomock.Nil()),
	).DoAndReturn(
		func(_ context.Context, request *services.CreateFirewallRuleRequest) (
			response *services.CreateFirewallRuleResponse, err error,
		) {
			return &services.CreateFirewallRuleResponse{FirewallRule: request.FirewallRule}, nil
		},
	).AnyTimes()

	nextService.EXPECT().UpdateFirewallRule(
		gomock.Not(gomock.Nil()), gomock.Not(gomock.Nil()),
	).DoAndReturn(
		func(_ context.Context, request *services.UpdateFirewallRuleRequest) (
			response *services.UpdateFirewallRuleResponse, err error,
		) {
			return &services.UpdateFirewallRuleResponse{FirewallRule: request.FirewallRule}, nil
		},
	).AnyTimes()

	firewallRuleSetupMetadataMocks(s)
	firewallRuleSetupGettersMocks(s, databaseFR)
}

func firewallRuleSetupMetadataMocks(s *ContrailTypeLogicService) {
	s.MetadataGetter.(*typesmock.MockMetadataGetter).EXPECT().GetMetadata(
		gomock.Not(gomock.Nil()), gomock.Not(gomock.Nil()),
	).DoAndReturn(
		func(_ context.Context, requested basemodels.Metadata) (
			response *basemodels.Metadata, err error,
		) {
			return firewallRuleSetupMetadata(requested)
		},
	).AnyTimes()
}

func firewallRuleSetupMetadata(requested basemodels.Metadata) (
	*basemodels.Metadata, error) {
	if requested.Type == models.KindVirtualNetwork {
		if requested.UUID == "virtual-network-uuid-1" {
			return nil, errutil.ErrorNotFound
		}

		return &basemodels.Metadata{
			UUID:   "virtual-network-uuid-2",
			FQName: []string{"virtual-network-uuid-2"},
		}, nil
	}

	if requested.Type == models.KindAddressGroup {
		if requested.FQName[0] == "address-group-uuid-1" {
			return nil, errutil.ErrorNotFound
		}

		return &basemodels.Metadata{
			UUID:   "address-group-uuid-2",
			FQName: []string{"address-group-uuid-2"},
		}, nil
	}

	if requested.Type == models.KindTagType {
		return &basemodels.Metadata{
			UUID: "tag-type-uuid",
		}, nil
	}

	if requested.FQName[0] == "namespace=default" {
		return &basemodels.Metadata{
			UUID:   "tag-uuid-1",
			FQName: []string{"namespace=default"},
		}, nil
	}

	if requested.FQName[0] == "domain-uuid" {
		return &basemodels.Metadata{
			UUID:   "tag-uuid-2",
			FQName: []string{"domain-uuid", "project-uuid", "namespace=default"},
		}, nil
	}

	return &basemodels.Metadata{
		UUID:   "tag-uuid-3",
		FQName: []string{"namespace=contrail"},
	}, nil
}

func firewallRuleSetupGettersMocks(s *ContrailTypeLogicService, databaseFR *models.FirewallRule) {
	readService := s.ReadService.(*servicesmock.MockReadService) //nolint: errcheck

	readService.EXPECT().GetTagType(gomock.Not(gomock.Nil()), gomock.Not(gomock.Nil())).Return(
		&services.GetTagTypeResponse{
			TagType: &models.TagType{
				TagTypeID: "0xff",
			},
		},
		nil,
	).AnyTimes()

	readService.EXPECT().GetTag(
		gomock.Not(gomock.Nil()), gomock.Not(gomock.Nil()),
	).DoAndReturn(
		func(_ context.Context, request *services.GetTagRequest) (
			response *services.GetTagResponse, err error,
		) {
			if request.GetID() == "tag-uuid-1" {
				return &services.GetTagResponse{
						Tag: &models.Tag{
							UUID:   "tag-uuid-1",
							FQName: []string{"namespace=default"},
							TagID:  "12",
						},
					},
					nil
			}

			if request.GetID() == "tag-uuid-2" {
				return nil, errutil.ErrorNotFound
			}

			return &services.GetTagResponse{
					Tag: &models.Tag{
						UUID:   "tag-uuid-3",
						FQName: []string{"namespace=contrail"},
						TagID:  "0x00ff0002",
					},
				},
				nil
		},
	).AnyTimes()

	if databaseFR != nil {
		readService.EXPECT().GetFirewallRule(gomock.Not(gomock.Nil()), gomock.Not(gomock.Nil())).Return(
			&services.GetFirewallRuleResponse{
				FirewallRule: databaseFR,
			},
			nil,
		).AnyTimes()
	}
}

func TestCreateFirewallRule(t *testing.T) {
	tests := []struct {
		name              string
		testFirewallRule  models.FirewallRule
		expectedParams    expectedParams
		IsInternalRequest bool
		errorCode         codes.Code
	}{
		{
			name: "Try to create with read-only draft-mode-state property",
			testFirewallRule: models.FirewallRule{
				UUID:           "test-firewall-rule",
				DraftModeState: "draft_mode_state",
			},
			errorCode: codes.InvalidArgument,
		},
		{
			name: "Try to create when cannot find refs in database",
			testFirewallRule: models.FirewallRule{
				UUID:   "test-firewall-rule",
				FQName: []string{"default-policy-management", "test-firewall-rule"},
				VirtualNetworkRefs: []*models.FirewallRuleVirtualNetworkRef{
					{
						UUID: "virtual-network-uuid-1",
					},
				},
			},
			errorCode: codes.NotFound,
		},
		{
			name: "Try to create when filled virtual-network refs are of different scope",
			testFirewallRule: models.FirewallRule{
				UUID:   "test-firewall-rule",
				FQName: []string{"default-policy-management", "test-firewall-rule"},
				VirtualNetworkRefs: []*models.FirewallRuleVirtualNetworkRef{
					{
						UUID: "virtual-network-uuid-2",
					},
				},
			},
			errorCode: codes.InvalidArgument,
		},
		{
			name: "Try to create when address-group refs are scoped",
			testFirewallRule: models.FirewallRule{
				UUID:           "test-firewall-rule",
				FQName:         []string{"default-policy-management", "test-firewall-rule"},
				DraftModeState: "draft_mode_state",
				AddressGroupRefs: []*models.FirewallRuleAddressGroupRef{
					{
						UUID: "address-group-ref-uuid",
						To:   []string{"address-group-ref-uuid"},
					},
				},
			},
			IsInternalRequest: true,
			errorCode:         codes.InvalidArgument,
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
			errorCode:         codes.InvalidArgument,
			IsInternalRequest: true,
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
				Service: &models.FirewallServiceType{},
				ServiceGroupRefs: []*models.FirewallRuleServiceGroupRef{
					{
						UUID: "service-group-ref-uuid",
						To:   []string{"service-group-ref-uuid"},
					},
				},
			},
			errorCode:         codes.InvalidArgument,
			IsInternalRequest: true,
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
					Protocol: models.ICMPProtocol,
				},
				Endpoint1: &models.FirewallRuleEndpointType{
					Any:            true,
					VirtualNetwork: "virtual-network-uuid",
					Tags:           []string{"namespace=default"},
				},
			},
			errorCode:         codes.InvalidArgument,
			IsInternalRequest: true,
		},
		{
			name: "Try to create with label match-tag",
			testFirewallRule: models.FirewallRule{
				Service: &models.FirewallServiceType{
					Protocol: models.TCPProtocol,
				},
				MatchTags: &models.FirewallRuleMatchTagsType{
					TagList: []string{"Label"},
				},
			},
			errorCode: codes.InvalidArgument,
		},
		{
			name: "Try to create with tag references definied",
			testFirewallRule: models.FirewallRule{
				Service: &models.FirewallServiceType{
					Protocol: models.TCPProtocol,
				},
				TagRefs: []*models.FirewallRuleTagRef{
					{
						UUID: "tag-uuid",
						To:   []string{"default-project", "tag-uuid"},
					},
				},
			},
			errorCode: codes.InvalidArgument,
		},
		{
			name: "Try to create with improper tag name",
			testFirewallRule: models.FirewallRule{
				Service: &models.FirewallServiceType{
					Protocol: models.TCPProtocol,
				},
				Endpoint1: &models.FirewallRuleEndpointType{
					Tags: []string{"namespace"},
				},
			},
			errorCode:         codes.NotFound,
			IsInternalRequest: true,
		},
		{
			name: "Try to create with improper parent type",
			testFirewallRule: models.FirewallRule{
				ParentType: models.KindVirtualNetwork,
				Service: &models.FirewallServiceType{
					Protocol: models.TCPProtocol,
				},
				Endpoint1: &models.FirewallRuleEndpointType{
					Tags: []string{"namespace=default"},
				},
			},
			errorCode: codes.InvalidArgument,
		},
		{
			name: "Try to create with non existing tag",
			testFirewallRule: models.FirewallRule{
				FQName:     []string{"domain-uuid", "project-uuid", "firewall-rule-uuid"},
				ParentType: models.KindProject,
				Service: &models.FirewallServiceType{
					Protocol: models.TCPProtocol,
				},
				Endpoint1: &models.FirewallRuleEndpointType{
					Tags: []string{"namespace=default"},
				},
			},
			errorCode:         codes.NotFound,
			IsInternalRequest: true,
		},
		{
			name: "Try to create with address group references definied",
			testFirewallRule: models.FirewallRule{
				Service: &models.FirewallServiceType{
					Protocol: models.TCPProtocol,
				},
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
			name: "Try to create with non existing address-group",
			testFirewallRule: models.FirewallRule{
				Service: &models.FirewallServiceType{
					Protocol: models.TCPProtocol,
				},
				Endpoint1: &models.FirewallRuleEndpointType{
					AddressGroup: "address-group-uuid-1",
				},
			},
			errorCode:         codes.NotFound,
			IsInternalRequest: true,
		},
		{
			name: "Create firewall-rule with default match-tag propetry",
			testFirewallRule: models.FirewallRule{
				FQName:     []string{"firewall-rule-uuid"},
				ParentType: models.KindProject,
				Service: &models.FirewallServiceType{
					Protocol: models.TCPProtocol,
				},
				Endpoint1: &models.FirewallRuleEndpointType{
					Tags: []string{"global:namespace=contrail"},
				},
				Endpoint2: &models.FirewallRuleEndpointType{
					Tags: []string{"namespace=default"},
				},
				VirtualNetworkRefs: []*models.FirewallRuleVirtualNetworkRef{
					{
						UUID: "virtual-network-uuid-2",
					},
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
					Protocol:   models.TCPProtocol,
					ProtocolID: 6,
				},
				Endpoint1: &models.FirewallRuleEndpointType{
					Tags:   []string{"global:namespace=contrail"},
					TagIds: []int64{0x00ff0002},
				},
				Endpoint2: &models.FirewallRuleEndpointType{
					Tags:   []string{"namespace=default"},
					TagIds: []int64{12},
				},
				TagRefs: []*models.FirewallRuleTagRef{
					{
						UUID: "tag-uuid-3",
						To:   []string{"namespace=contrail"},
					},
					{
						UUID: "tag-uuid-1",
						To:   []string{"namespace=default"},
					},
				},
				AddressGroupRefs: []*models.FirewallRuleAddressGroupRef{},
				VirtualNetworkRefs: []*models.FirewallRuleVirtualNetworkRef{
					{
						UUID: "virtual-network-uuid-2",
						To:   []string{"virtual-network-uuid-2"},
					},
				},
			},
			IsInternalRequest: true,
		},
		{
			name: "Create firewall-rule with match-tag propetry defined",
			testFirewallRule: models.FirewallRule{
				Service: &models.FirewallServiceType{
					Protocol: "240",
				},
				MatchTags: &models.FirewallRuleMatchTagsType{
					TagList: []string{"Application", "Tier", "OtherTag"},
				},
				Endpoint1: &models.FirewallRuleEndpointType{
					AddressGroup: "address-group-uuid-2",
				},
			},
			expectedParams: expectedParams{
				MatchTags: &models.FirewallRuleMatchTagsType{
					TagList: []string{"Application", "Tier", "OtherTag"},
				},
				MatchTagTypes: &models.FirewallRuleMatchTagsTypeIdList{
					TagType: []int64{1, 2, 255},
				},
				Endpoint1: &models.FirewallRuleEndpointType{
					AddressGroup: "address-group-uuid-2",
				},
				Service: &models.FirewallServiceType{
					Protocol:   "240",
					ProtocolID: 240,
				},
				TagRefs: []*models.FirewallRuleTagRef{},
				AddressGroupRefs: []*models.FirewallRuleAddressGroupRef{
					{
						UUID: "address-group-uuid-2",
						To:   []string{"address-group-uuid-2"},
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()
			service := makeMockedContrailTypeLogicService(mockCtrl)
			firewallRuleSetupMocks(service, nil)

			ctx := context.Background()
			if tt.IsInternalRequest {
				ctx = auth.WithInternalRequest(ctx)
			}

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

func TestUpdateFirewallRule(t *testing.T) {
	tests := []struct {
		name              string
		request           services.UpdateFirewallRuleRequest
		databaseFR        *models.FirewallRule
		expected          *models.FirewallRule
		IsInternalRequest bool
		errorCode         codes.Code
	}{
		{
			name: "Try to update read-only draft-mode-state property",
			request: services.UpdateFirewallRuleRequest{
				FirewallRule: &models.FirewallRule{
					DraftModeState: "draft_mode_state",
				},
				FieldMask: types.FieldMask{
					Paths: []string{
						models.FirewallRuleFieldDraftModeState,
					},
				},
			},
			databaseFR: &models.FirewallRule{},
			errorCode:  codes.InvalidArgument,
		},
		{
			name: "Try to update with not existing virtual-network ref",
			request: services.UpdateFirewallRuleRequest{
				FirewallRule: &models.FirewallRule{
					VirtualNetworkRefs: []*models.FirewallRuleVirtualNetworkRef{
						{
							UUID: "virtual-network-uuid-1",
						},
					},
				},
				FieldMask: types.FieldMask{
					Paths: []string{
						models.FirewallRuleFieldVirtualNetworkRefs,
					},
				},
			},
			databaseFR: &models.FirewallRule{
				FQName: []string{"default-policy-management", "test-firewall-rule"},
			},
			errorCode: codes.NotFound,
		},
		{
			name: "Try to update virtual-network refs with different scope",
			request: services.UpdateFirewallRuleRequest{
				FirewallRule: &models.FirewallRule{
					VirtualNetworkRefs: []*models.FirewallRuleVirtualNetworkRef{
						{
							UUID: "virtual-network-uuid-2",
						},
					},
				},
				FieldMask: types.FieldMask{
					Paths: []string{
						models.FirewallRuleFieldVirtualNetworkRefs,
					},
				},
			},
			databaseFR: &models.FirewallRule{
				FQName: []string{"default-policy-management", "test-firewall-rule"},
			},
			IsInternalRequest: true,
			errorCode:         codes.InvalidArgument,
		},
		{
			name: "Try to define service and service-group refs",
			request: services.UpdateFirewallRuleRequest{
				FirewallRule: &models.FirewallRule{
					Service: &models.FirewallServiceType{
						Protocol: "240",
					},
				},
				FieldMask: types.FieldMask{
					Paths: []string{
						basemodels.JoinPath(
							models.FirewallRuleFieldService,
							models.FirewallServiceTypeFieldProtocol,
						),
					},
				},
			},
			IsInternalRequest: true,
			databaseFR: &models.FirewallRule{
				ServiceGroupRefs: []*models.FirewallRuleServiceGroupRef{{}},
			},
			errorCode: codes.InvalidArgument,
		},
		{
			name: "Try to remove service ref",
			request: services.UpdateFirewallRuleRequest{
				FirewallRule: &models.FirewallRule{},
				FieldMask: types.FieldMask{
					Paths: []string{
						models.FirewallRuleFieldService,
					},
				},
			},
			databaseFR: &models.FirewallRule{
				Service: &models.FirewallServiceType{
					Protocol:   "240",
					ProtocolID: 240,
				},
			},
			errorCode: codes.InvalidArgument,
		},
		{
			name: "Update firewall rule to default match tag",
			request: services.UpdateFirewallRuleRequest{
				FirewallRule: &models.FirewallRule{
					MatchTags: &models.FirewallRuleMatchTagsType{},
				},
				FieldMask: types.FieldMask{
					Paths: []string{
						models.FirewallRuleFieldMatchTags,
					},
				},
			},
			expected: &models.FirewallRule{
				MatchTags: &models.FirewallRuleMatchTagsType{
					TagList: []string{"application"},
				},
				MatchTagTypes: &models.FirewallRuleMatchTagsTypeIdList{
					TagType: []int64{1},
				},
				TagRefs:          []*models.FirewallRuleTagRef{},
				AddressGroupRefs: []*models.FirewallRuleAddressGroupRef{},
			},
			databaseFR: &models.FirewallRule{
				Service: &models.FirewallServiceType{
					Protocol:   "240",
					ProtocolID: 240,
				},
				MatchTags: &models.FirewallRuleMatchTagsType{
					TagList: []string{"tier"},
				},
				MatchTagTypes: &models.FirewallRuleMatchTagsTypeIdList{
					TagType: []int64{2},
				},
				TagRefs: []*models.FirewallRuleTagRef{},
			},
			IsInternalRequest: true,
		},
		{
			name: "Update firewall rule service property",
			request: services.UpdateFirewallRuleRequest{
				FirewallRule: &models.FirewallRule{
					Service: &models.FirewallServiceType{
						Protocol: models.TCPProtocol,
					},
				},
				FieldMask: types.FieldMask{
					Paths: []string{
						basemodels.JoinPath(
							models.FirewallRuleFieldService,
							models.FirewallServiceTypeFieldProtocol,
						),
					},
				},
			},
			expected: &models.FirewallRule{
				Service: &models.FirewallServiceType{
					Protocol:   models.TCPProtocol,
					ProtocolID: 6,
				},
				MatchTagTypes: &models.FirewallRuleMatchTagsTypeIdList{
					TagType: []int64{},
				},
				TagRefs:          []*models.FirewallRuleTagRef{},
				AddressGroupRefs: []*models.FirewallRuleAddressGroupRef{},
			},
			databaseFR: &models.FirewallRule{
				Service: &models.FirewallServiceType{
					Protocol:   "240",
					ProtocolID: 240,
				},
			},
		},
		{
			name: "Try to update with improper endpoint type",
			request: services.UpdateFirewallRuleRequest{
				FirewallRule: &models.FirewallRule{
					Endpoint1: &models.FirewallRuleEndpointType{
						AddressGroup: "address-group-uuid-1",
						Any:          true,
					},
				},
				FieldMask: types.FieldMask{
					Paths: []string{
						basemodels.JoinPath(
							models.FirewallRuleFieldEndpoint1,
							models.FirewallRuleEndpointTypeFieldAddressGroup,
						),
						basemodels.JoinPath(
							models.FirewallRuleFieldEndpoint1,
							models.FirewallRuleEndpointTypeFieldAny,
						),
					},
				},
			},
			databaseFR: &models.FirewallRule{
				Service: &models.FirewallServiceType{
					Protocol:   models.TCPProtocol,
					ProtocolID: 6,
				},
			},
			IsInternalRequest: true,
			errorCode:         codes.InvalidArgument,
		},
		{
			name: "Update firewall rule tags",
			request: services.UpdateFirewallRuleRequest{
				FirewallRule: &models.FirewallRule{
					Endpoint1: &models.FirewallRuleEndpointType{
						Tags: []string{"global:namespace=default"},
					},
				},
				FieldMask: types.FieldMask{
					Paths: []string{
						basemodels.JoinPath(
							models.FirewallRuleFieldEndpoint1,
							models.FirewallRuleEndpointTypeFieldTags,
						),
					},
				},
			},
			expected: &models.FirewallRule{
				Endpoint1: &models.FirewallRuleEndpointType{
					Tags:   []string{"global:namespace=default"},
					TagIds: []int64{12},
				},
				MatchTagTypes: &models.FirewallRuleMatchTagsTypeIdList{
					TagType: []int64{},
				},
				TagRefs: []*models.FirewallRuleTagRef{
					{
						UUID: "tag-uuid-1",
						To:   []string{"namespace=default"},
					},
				},
				AddressGroupRefs: []*models.FirewallRuleAddressGroupRef{},
			},
			databaseFR: &models.FirewallRule{
				Endpoint1: &models.FirewallRuleEndpointType{
					Tags:   []string{"global:namespace=contrail"},
					TagIds: []int64{0x00ff0002},
				},
				Service: &models.FirewallServiceType{
					Protocol:   models.TCPProtocol,
					ProtocolID: 6,
				},
			},
			IsInternalRequest: true,
		},
		{
			name: "Update firewall rule address group refs",
			request: services.UpdateFirewallRuleRequest{
				FirewallRule: &models.FirewallRule{
					Endpoint1: &models.FirewallRuleEndpointType{
						AddressGroup: "address-group-uuid-2",
					},
				},
				FieldMask: types.FieldMask{
					Paths: []string{
						basemodels.JoinPath(
							models.FirewallRuleFieldEndpoint1,
							models.FirewallRuleEndpointTypeFieldAddressGroup,
						),
					},
				},
			},
			expected: &models.FirewallRule{
				Endpoint1: &models.FirewallRuleEndpointType{
					AddressGroup: "address-group-uuid-2",
				},
				MatchTagTypes: &models.FirewallRuleMatchTagsTypeIdList{
					TagType: []int64{},
				},
				TagRefs: []*models.FirewallRuleTagRef{},
				AddressGroupRefs: []*models.FirewallRuleAddressGroupRef{
					{
						UUID: "address-group-uuid-2",
						To:   []string{"address-group-uuid-2"},
					},
				},
			},
			databaseFR: &models.FirewallRule{
				Endpoint1: &models.FirewallRuleEndpointType{
					Tags:   []string{"global:namespace=contrail"},
					TagIds: []int64{0x00ff0002},
				},
				Service: &models.FirewallServiceType{
					Protocol:   models.TCPProtocol,
					ProtocolID: 6,
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()
			service := makeMockedContrailTypeLogicService(mockCtrl)
			firewallRuleSetupMocks(service, tt.databaseFR)

			ctx := context.Background()
			if tt.IsInternalRequest {
				ctx = auth.WithInternalRequest(ctx)
			}

			expectedResponse := &services.UpdateFirewallRuleResponse{
				FirewallRule: tt.expected,
			}

			updateFirewallRuleResponse, err := service.UpdateFirewallRule(ctx, &tt.request)
			if tt.errorCode != codes.OK {
				assert.Error(t, err)
				status, ok := status.FromError(err)
				assert.True(t, ok)
				assert.Equal(t, tt.errorCode, status.Code())
			} else {
				assert.NoError(t, err)
				assert.Equal(t, expectedResponse, updateFirewallRuleResponse)
			}
		})
	}
}
