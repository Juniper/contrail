package logic

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"

	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/services"
	servicesmock "github.com/Juniper/contrail/pkg/services/mock"
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
		{
			name: "Security group create - error on empty parent project",
			inputData: &SecurityGroup{
				Name:      "new security group",
				TenantID:  "92882ca8f99342f286430c05c96e12dd",
				ProjectID: "92882ca8f99342f286430c05c96e12dd",
			},
			expectedResponse: nil,
			expectedErr:      true,
			mockData:         mockSgReadEmptyProject(t),
		},
		{
			name: "Security group create - error on manual creation of default security group",
			inputData: &SecurityGroup{
				Name:      "default",
				TenantID:  "92882ca8f99342f286430c05c96e12dd",
				ProjectID: "92882ca8f99342f286430c05c96e12dd",
			},
			expectedResponse: nil,
			expectedErr:      true,
			mockData:         mockSgReadDefaultSgError(t),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rp := RequestParameters{
				ReadService:  tt.mockData.mockService,
				WriteService: tt.mockData.mockService,
			}
			defer tt.mockData.mockController.Finish()

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

		})
	}

}

func mockSgReadEmptyProject(t *testing.T) *mock {
	mockCtrl := gomock.NewController(t)
	mockServ := servicesmock.NewMockService(mockCtrl)

	mockServ.EXPECT().GetProject(gomock.Any(), gomock.Any()).Return(
		nil, errors.New("can't read project from the database"),
	)

	return &mock{mockCtrl, mockServ}
}

func mockSgReadDefaultSgError(t *testing.T) *mock {
	mockCtrl := gomock.NewController(t)
	mockServ := servicesmock.NewMockService(mockCtrl)

	mockServ.EXPECT().GetProject(gomock.Any(), gomock.Any()).Return(
		&services.GetProjectResponse{
			Project: &models.Project{
				UUID: "92882ca8-f993-42f2-8643-0c05c96e12dd",
				Name: "project mock",
				SecurityGroups: []*models.SecurityGroup{
					{
						UUID:   "92882ca8-f993-42f2-8643-0c05c96e12dd",
						FQName: []string{"domain_mock", "project_mock", "default"},
					},
				},
			},
		},
		nil,
	).AnyTimes()

	return &mock{mockCtrl, mockServ}
}
