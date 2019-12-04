package types

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/Juniper/asf/pkg/models/basemodels"
	"github.com/Juniper/asf/pkg/services/baseservices"
	"github.com/Juniper/contrail/pkg/errutil"
	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/services"
	servicesmock "github.com/Juniper/contrail/pkg/services/mock"
	typesmock "github.com/Juniper/contrail/pkg/types/mock"
)

func TestCreateFirewallPolicy(t *testing.T) {
	tests := []struct {
		name               string
		testFirewallPolicy models.FirewallPolicy
		IsInternalRequest  bool
		errorCode          codes.Code
	}{
		{
			name: "Try to create with read-only draft-mode-state property",
			testFirewallPolicy: models.FirewallPolicy{
				UUID:           "test-firewall-policy",
				DraftModeState: "draft_mode_state",
			},
			errorCode: codes.InvalidArgument,
		},
		{
			name: "Internal create request with draft-mode-state defined",
			testFirewallPolicy: models.FirewallPolicy{
				UUID:           "test-firewall-policy",
				DraftModeState: "draft_mode_state",
			},
			IsInternalRequest: true,
		},
		{
			name: "Try to create when cannot find reference in database",
			testFirewallPolicy: models.FirewallPolicy{
				UUID:   "test-firewall-policy",
				FQName: []string{"default-policy-management", "test-firewall-policy"},
				FirewallRuleRefs: []*models.FirewallPolicyFirewallRuleRef{
					{
						UUID: "firewall-rule-uuid-1",
					},
				},
			},
			errorCode: codes.NotFound,
		},
		{
			name: "Try to create when firewall-rule refs are scoped",
			testFirewallPolicy: models.FirewallPolicy{
				UUID:   "test-firewall-policy",
				FQName: []string{"default-policy-management", "test-firewall-policy"},
				FirewallRuleRefs: []*models.FirewallPolicyFirewallRuleRef{
					{
						UUID: "firewall-rule-uuid-2",
					},
				},
			},
			errorCode: codes.InvalidArgument,
		},
		{
			name: "Create firewall policy",
			testFirewallPolicy: models.FirewallPolicy{
				UUID:   "test-firewall-policy",
				FQName: []string{"default-domain", "default-project", "test-firewall-policy"},
				FirewallRuleRefs: []*models.FirewallPolicyFirewallRuleRef{
					{
						UUID: "firewall-rule-ref-uuid",
						To:   []string{"firewall-rule-ref-uuid"},
					},
				},
			},
			IsInternalRequest: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()
			service := makeMockedContrailTypeLogicService(mockCtrl)
			setupNextServiceMocks(service)
			setupReadServiceMock(service, nil)

			createFirewallPolicyResponse, err := service.CreateFirewallPolicy(
				getContext(tt.IsInternalRequest),
				&services.CreateFirewallPolicyRequest{FirewallPolicy: &tt.testFirewallPolicy},
			)
			if tt.errorCode != codes.OK {
				assert.Error(t, err)
				status, ok := status.FromError(err)
				assert.True(t, ok)
				assert.Equal(t, tt.errorCode, status.Code())
			} else {
				assert.NoError(t, err)
				assert.Equal(
					t,
					&services.CreateFirewallPolicyResponse{FirewallPolicy: &tt.testFirewallPolicy},
					createFirewallPolicyResponse,
				)
			}
		})
	}
}

func TestUpdateFirewallPolicy(t *testing.T) {
	tests := []struct {
		name               string
		testFirewallPolicy models.FirewallPolicy
		dbFirewallPolicy   models.FirewallPolicy
		IsInternalRequest  bool
		errorCode          codes.Code
	}{
		{
			name: "Try to update read-only draft-mode-state property",
			testFirewallPolicy: models.FirewallPolicy{
				UUID:           "test-firewall-policy",
				DraftModeState: "draft_mode_state",
			},
			errorCode: codes.InvalidArgument,
		},
		{
			name: "Try to update with non existing firewall-rule ref",
			testFirewallPolicy: models.FirewallPolicy{
				UUID:   "test-firewall-policy",
				FQName: []string{"default-policy-management", "test-firewall-policy"},
				FirewallRuleRefs: []*models.FirewallPolicyFirewallRuleRef{
					{
						UUID: "firewall-rule-uuid-1",
					},
				},
			},
			dbFirewallPolicy: models.FirewallPolicy{
				FQName: []string{"default-policy-management", "test-firewall-policy"},
			},
			errorCode: codes.NotFound,
		},
		{
			name: "Try to update firewall-rule refs",
			testFirewallPolicy: models.FirewallPolicy{
				UUID:   "test-firewall-policy",
				FQName: []string{"default-policy-management", "test-firewall-policy"},
				FirewallRuleRefs: []*models.FirewallPolicyFirewallRuleRef{
					{
						UUID: "firewall-rule-ref-uuid",
						To:   []string{"default-domain", "default-project", "firewall-rule-ref-uuid"},
					},
				},
			},
			dbFirewallPolicy: models.FirewallPolicy{
				FQName: []string{"default-policy-management", "test-firewall-policy"},
			},
			IsInternalRequest: true,
			errorCode:         codes.InvalidArgument,
		},
		{
			name: "Update firewall-rule refs",
			testFirewallPolicy: models.FirewallPolicy{
				UUID: "test-firewall-policy",
				FirewallRuleRefs: []*models.FirewallPolicyFirewallRuleRef{
					{
						UUID: "firewall-rule-uuid-2",
					},
				},
			},
			dbFirewallPolicy: models.FirewallPolicy{
				FQName: []string{"default-domain", "default-project", "test-firewall-policy"},
			},
			IsInternalRequest: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()
			service := makeMockedContrailTypeLogicService(mockCtrl)
			setupNextServiceMocks(service)
			setupReadServiceMock(service, &tt.dbFirewallPolicy)

			updateFirewallPolicyResponse, err := service.UpdateFirewallPolicy(
				getContext(tt.IsInternalRequest),
				&services.UpdateFirewallPolicyRequest{FirewallPolicy: &tt.testFirewallPolicy},
			)
			if tt.errorCode != codes.OK {
				assert.Error(t, err)
				status, ok := status.FromError(err)
				assert.True(t, ok)
				assert.Equal(t, tt.errorCode, status.Code())
			} else {
				assert.NoError(t, err)
				assert.Equal(
					t,
					&services.UpdateFirewallPolicyResponse{FirewallPolicy: &tt.testFirewallPolicy},
					updateFirewallPolicyResponse,
				)
			}
		})
	}
}

func setupNextServiceMocks(s *ContrailTypeLogicService) {
	nextService := s.Next().(*servicesmock.MockService) //nolint: errcheck
	nextService.EXPECT().CreateFirewallPolicy(
		gomock.Not(gomock.Nil()), gomock.Not(gomock.Nil()),
	).DoAndReturn(
		func(_ context.Context, request *services.CreateFirewallPolicyRequest) (
			response *services.CreateFirewallPolicyResponse, err error,
		) {
			return &services.CreateFirewallPolicyResponse{FirewallPolicy: request.FirewallPolicy}, nil
		},
	).MaxTimes(1)

	nextService.EXPECT().UpdateFirewallPolicy(
		gomock.Not(gomock.Nil()), gomock.Not(gomock.Nil()),
	).DoAndReturn(
		func(_ context.Context, request *services.UpdateFirewallPolicyRequest) (
			response *services.UpdateFirewallPolicyResponse, err error,
		) {
			return &services.UpdateFirewallPolicyResponse{FirewallPolicy: request.FirewallPolicy}, nil
		},
	).MaxTimes(1)
}

func setupReadServiceMock(s *ContrailTypeLogicService, databaseFP *models.FirewallPolicy) {
	readService := s.ReadService.(*servicesmock.MockReadService) //nolint: errcheck
	readService.EXPECT().GetFirewallPolicy(gomock.Not(gomock.Nil()), gomock.Not(gomock.Nil())).Return(
		&services.GetFirewallPolicyResponse{
			FirewallPolicy: databaseFP,
		},
		nil,
	).MaxTimes(1)

	s.MetadataGetter.(*typesmock.MockMetadataGetter).EXPECT().GetMetadata(
		gomock.Not(gomock.Nil()), gomock.Not(gomock.Nil()),
	).DoAndReturn(
		func(_ context.Context, requested basemodels.Metadata) (
			response *basemodels.Metadata, err error,
		) {
			if requested.UUID == "firewall-rule-uuid-1" {
				return nil, errutil.ErrorNotFound
			}

			return &basemodels.Metadata{
				UUID:   "firewall-rule-uuid-2",
				FQName: []string{"firewall-rule-uuid-2"},
			}, nil
		},
	).AnyTimes()
}

func getContext(isInternalRequest bool) context.Context {
	if isInternalRequest {
		return baseservices.WithInternalRequest(context.Background())
	}

	return context.Background()
}
