package types

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/services"
	"github.com/Juniper/contrail/pkg/services/mock"
)

func firewallPolicySetupMocks(s *ContrailTypeLogicService, databaseFP *models.FirewallPolicy) {
	nextService := s.Next().(*servicesmock.MockService)          //nolint: errcheck
	readService := s.ReadService.(*servicesmock.MockReadService) //nolint: errcheck

	nextService.EXPECT().CreateFirewallPolicy(
		gomock.Not(gomock.Nil()), gomock.Not(gomock.Nil()),
	).DoAndReturn(
		func(_ context.Context, request *services.CreateFirewallPolicyRequest) (
			response *services.CreateFirewallPolicyResponse, err error,
		) {
			return &services.CreateFirewallPolicyResponse{FirewallPolicy: request.FirewallPolicy}, nil
		},
	).AnyTimes()

	nextService.EXPECT().UpdateFirewallPolicy(
		gomock.Not(gomock.Nil()), gomock.Not(gomock.Nil()),
	).DoAndReturn(
		func(_ context.Context, request *services.UpdateFirewallPolicyRequest) (
			response *services.UpdateFirewallPolicyResponse, err error,
		) {
			return &services.UpdateFirewallPolicyResponse{FirewallPolicy: request.FirewallPolicy}, nil
		},
	).AnyTimes()

	if databaseFP != nil {
		readService.EXPECT().GetFirewallPolicy(gomock.Not(gomock.Nil()), gomock.Not(gomock.Nil())).Return(
			&services.GetFirewallPolicyResponse{
				FirewallPolicy: databaseFP,
			},
			nil,
		).AnyTimes()
	}
}

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
			name: "Try to create when firewall-rule refs are scoped",
			testFirewallPolicy: models.FirewallPolicy{
				UUID:   "test-firewall-policy",
				FQName: []string{"default-policy-management", "test-firewall-policy"},
				FirewallRuleRefs: []*models.FirewallPolicyFirewallRuleRef{
					{
						UUID: "firewall-rule-ref-uuid",
						To:   []string{"firewall-rule-ref-uuid"},
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
			firewallPolicySetupMocks(service, nil)

			ctx := context.Background()
			if tt.IsInternalRequest {
				ctx = WithInternalRequest(ctx)
			}

			paramRequest := &services.CreateFirewallPolicyRequest{FirewallPolicy: &tt.testFirewallPolicy}
			expectedResponse := &services.CreateFirewallPolicyResponse{
				FirewallPolicy: &tt.testFirewallPolicy,
			}

			createFirewallPolicyResponse, err := service.CreateFirewallPolicy(ctx, paramRequest)
			if tt.errorCode != codes.OK {
				assert.Error(t, err)
				status, ok := status.FromError(err)
				assert.True(t, ok)
				assert.Equal(t, tt.errorCode, status.Code())
			} else {
				assert.NoError(t, err)
				assert.Equal(t, expectedResponse, createFirewallPolicyResponse)
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
						UUID: "firewall-rule-ref-uuid",
						To:   []string{"firewall-rule-ref-uuid"},
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
			firewallPolicySetupMocks(service, &tt.dbFirewallPolicy)

			ctx := context.Background()
			if tt.IsInternalRequest {
				ctx = WithInternalRequest(ctx)
			}

			paramRequest := &services.UpdateFirewallPolicyRequest{FirewallPolicy: &tt.testFirewallPolicy}
			expectedResponse := &services.UpdateFirewallPolicyResponse{
				FirewallPolicy: &tt.testFirewallPolicy,
			}

			updateFirewallPolicyResponse, err := service.UpdateFirewallPolicy(ctx, paramRequest)
			if tt.errorCode != codes.OK {
				assert.Error(t, err)
				status, ok := status.FromError(err)
				assert.True(t, ok)
				assert.Equal(t, tt.errorCode, status.Code())
			} else {
				assert.NoError(t, err)
				assert.Equal(t, expectedResponse, updateFirewallPolicyResponse)
			}
		})
	}
}
