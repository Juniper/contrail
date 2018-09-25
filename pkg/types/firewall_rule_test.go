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
)

func TestCreateFirewallRule(t *testing.T) {
	tests := []struct {
		name             string
		testFirewallRule models.FirewallRule
		errorCode        codes.Code
	}{
		{
			name: "TODO",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()
			service := makeMockedContrailTypeLogicService(mockCtrl)

			ctx := context.Background()

			paramRequest := &services.CreateFirewallRuleRequest{FirewallRule: &tt.testFirewallRule}
			expectedResponse := &services.CreateFirewallRuleResponse{FirewallRule: &tt.testFirewallRule}
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
