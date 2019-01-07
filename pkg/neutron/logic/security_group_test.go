package logic

import (
	"github.com/Juniper/contrail/pkg/services/mock"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

type mock struct {
	mockController *gomock.Controller
	mockService    *servicesmock.MockService
}

func TestSecurityGroupCreate(t *testing.T) {
	tests := []struct {
		name             string
		inputData        *SecurityGroup
		expectedResponse *SecurityGroupResponse
		expectedErr      bool
		expectedErrName  string
		mockData         *mock
	}{
		// Happy path is tested in integration tests.
		// TODO: test securityGroupDefault
		// TODO: test conflict of security group rules (give on input the same as the default ones)
		// TODO: test OverQouta for security group rules (give on input too many of them)
		// TODO: test emptySecurityGroupRules (security_group_rule.go:212)
		// TODO:
		{
			// TODO: test empty getProject()
			name:             "Security group create - error on empty parent project",
			inputData:        &SecurityGroup{},
			expectedResponse: nil,
			expectedErr:      true,
			mockData:         mockSgReadEmptyProject(t),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rp := RequestParameters{
				ReadService:  tt.mockData.mockService,
				WriteService: tt.mockData.mockService,
			}
			result, err := tt.inputData.Create(nil, rp)

			if tt.expectedErr {
				if tt.expectedErrName != "" {
					err, ok := err.(*Error)
					if !ok {
						t.Errorf("expected error name %s but returned error has no name.", tt.expectedErrName)
					}
					assert.Equal(t, err.fields["exception"], tt.expectedErrName)
				} else {
					assert.Error(t, err)
				}
			} else {
				assert.EqualValues(t, tt.expectedResponse, result)
			}

			defer tt.mockData.mockController.Finish()
		})
	}

}

func mockSgReadEmptyProject(t *testing.T) *mock {
	mockCtrl := gomock.NewController(t)
	mockServ := servicesmock.NewMockService(mockCtrl)

	// TODO write this mock...

	return &mock{mockCtrl, mockServ}
}
