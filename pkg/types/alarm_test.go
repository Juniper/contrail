package types

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/gogo/protobuf/types"
	"github.com/stretchr/testify/assert"

	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/services"
	servicesmock "github.com/Juniper/contrail/pkg/services/mock"
)

func TestCreateAlarm(t *testing.T) {
	tests := []struct {
		name      string
		testAlarm *models.Alarm
		errorCode codes.Code
	}{
		{
			name:      "Without parameters",
			testAlarm: &models.Alarm{},
			errorCode: codes.InvalidArgument,
		},
		{
			name: "Empty Rules",
			testAlarm: &models.Alarm{
				AlarmRules: &models.AlarmOrList{},
			},
			errorCode: codes.InvalidArgument,
		},
		{
			name: "Empty OrList",
			testAlarm: &models.Alarm{
				AlarmRules: &models.AlarmOrList{
					OrList: []*models.AlarmAndList{},
				},
			},
		},
		{
			name: "Empty AndList",
			testAlarm: &models.Alarm{
				AlarmRules: &models.AlarmOrList{
					OrList: []*models.AlarmAndList{
						{
							AndList: []*models.AlarmExpression{},
						},
					},
				},
			},
		},
		{
			name: "Empty Operand2",
			testAlarm: &models.Alarm{
				AlarmRules: &models.AlarmOrList{
					OrList: []*models.AlarmAndList{
						{
							AndList: []*models.AlarmExpression{
								{
									Operation: "==",
									Operand2:  &models.AlarmOperand2{},
								},
							},
						},
					},
				},
			},
			errorCode: codes.InvalidArgument,
		},
		{
			name: "Everything filled",
			testAlarm: &models.Alarm{
				AlarmRules: &models.AlarmOrList{
					OrList: []*models.AlarmAndList{
						{
							AndList: []*models.AlarmExpression{
								{
									Operation: "==",
									Operand2: &models.AlarmOperand2{
										JSONValue: "\"testvalue\"",
									},
								},
							},
						},
					},
				},
			},
		},
		{
			name: "Both Operand2 fields filled",
			testAlarm: &models.Alarm{
				AlarmRules: &models.AlarmOrList{
					OrList: []*models.AlarmAndList{
						{
							AndList: []*models.AlarmExpression{
								{
									Operation: "==",
									Operand2: &models.AlarmOperand2{
										JSONValue:    "testvalue",
										UveAttribute: "test.value",
									},
								},
							},
						},
					},
				},
			},
			errorCode: codes.InvalidArgument,
		},
		{
			name: "Range operation, correct JSON",
			testAlarm: &models.Alarm{
				AlarmRules: &models.AlarmOrList{
					OrList: []*models.AlarmAndList{
						{
							AndList: []*models.AlarmExpression{
								{
									Operation: "range",
									Operand2: &models.AlarmOperand2{
										JSONValue: "[0, 1]",
									},
								},
							},
						},
					},
				},
			},
		},
		{
			name: "Range operation, incorrect JSON",
			testAlarm: &models.Alarm{
				AlarmRules: &models.AlarmOrList{
					OrList: []*models.AlarmAndList{
						{
							AndList: []*models.AlarmExpression{
								{
									Operation: "range",
									Operand2: &models.AlarmOperand2{
										JSONValue: "[1, 0]",
									},
								},
							},
						},
					},
				},
			},
			errorCode: codes.InvalidArgument,
		},
		{
			name: "Range operation, invalid JSON",
			testAlarm: &models.Alarm{
				AlarmRules: &models.AlarmOrList{
					OrList: []*models.AlarmAndList{
						{
							AndList: []*models.AlarmExpression{
								{
									Operation: "range",
									Operand2: &models.AlarmOperand2{
										JSONValue: "/0, 1]",
									},
								},
							},
						},
					},
				},
			},
			errorCode: codes.InvalidArgument,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()
			service := makeMockedContrailTypeLogicService(mockCtrl)

			ctx := context.Background()
			paramRequest := services.CreateAlarmRequest{Alarm: tt.testAlarm}
			expectedResponse := services.CreateAlarmResponse{Alarm: tt.testAlarm}

			createAlarmCall := service.Next().(*servicesmock.MockService).EXPECT().CreateAlarm(
				gomock.Not(gomock.Nil()), gomock.Not(gomock.Nil()),
			).DoAndReturn(
				func(_ context.Context, _ *services.CreateAlarmRequest,
				) (response *services.CreateAlarmResponse, err error) {
					return &services.CreateAlarmResponse{Alarm: tt.testAlarm}, nil
				},
			)

			if tt.errorCode != codes.OK {
				createAlarmCall.MaxTimes(1)
			} else {
				createAlarmCall.Times(1)
			}

			createAlarmResponse, err := service.CreateAlarm(ctx, &paramRequest)
			if tt.errorCode != codes.OK {
				assert.Error(t, err)
				status, ok := status.FromError(err)
				assert.True(t, ok)
				assert.Equal(t, tt.errorCode, status.Code())
			} else {
				assert.NoError(t, err)
				assert.Equal(t, &expectedResponse, createAlarmResponse)
			}
		})
	}
}

func TestUpdateAlarm(t *testing.T) {
	rulesMask := types.FieldMask{Paths: []string{models.AlarmFieldAlarmRules}}
	tests := []struct {
		name      string
		request   services.UpdateAlarmRequest
		errorCode codes.Code
	}{
		{
			name: "Without parameters",
			request: services.UpdateAlarmRequest{
				Alarm: &models.Alarm{},
			},
		},
		{
			name: "Empty Rules",
			request: services.UpdateAlarmRequest{
				Alarm: &models.Alarm{
					AlarmRules: &models.AlarmOrList{},
				},
				FieldMask: rulesMask,
			},
			errorCode: codes.InvalidArgument,
		},
		{
			name: "Empty OrList",
			request: services.UpdateAlarmRequest{
				Alarm: &models.Alarm{
					AlarmRules: &models.AlarmOrList{
						OrList: []*models.AlarmAndList{},
					},
				},
				FieldMask: rulesMask,
			},
		},
		{
			name: "Empty AndList",
			request: services.UpdateAlarmRequest{
				Alarm: &models.Alarm{
					AlarmRules: &models.AlarmOrList{
						OrList: []*models.AlarmAndList{
							{
								AndList: []*models.AlarmExpression{},
							},
						},
					},
				},
				FieldMask: rulesMask,
			},
		},
		{
			name: "Empty Operand2",
			request: services.UpdateAlarmRequest{
				Alarm: &models.Alarm{
					AlarmRules: &models.AlarmOrList{
						OrList: []*models.AlarmAndList{
							{
								AndList: []*models.AlarmExpression{
									{
										Operation: "==",
										Operand2:  &models.AlarmOperand2{},
									},
								},
							},
						},
					},
				},
				FieldMask: rulesMask,
			},
			errorCode: codes.InvalidArgument,
		},
		{
			name: "Everything filled",
			request: services.UpdateAlarmRequest{
				Alarm: &models.Alarm{
					AlarmRules: &models.AlarmOrList{
						OrList: []*models.AlarmAndList{
							{
								AndList: []*models.AlarmExpression{
									{
										Operation: "==",
										Operand2: &models.AlarmOperand2{
											JSONValue: "\"testvalue\"",
										},
									},
								},
							},
						},
					},
				},
				FieldMask: rulesMask,
			},
		},
		{
			name: "Both Operand2 fields filled",
			request: services.UpdateAlarmRequest{
				Alarm: &models.Alarm{
					AlarmRules: &models.AlarmOrList{
						OrList: []*models.AlarmAndList{
							{
								AndList: []*models.AlarmExpression{
									{
										Operation: "==",
										Operand2: &models.AlarmOperand2{
											JSONValue:    "\"testvalue\"",
											UveAttribute: "\"testvalue\"",
										},
									},
								},
							},
						},
					},
				},
				FieldMask: rulesMask,
			},
			errorCode: codes.InvalidArgument,
		},
		{
			name: "Range operation, correct JSON",
			request: services.UpdateAlarmRequest{
				Alarm: &models.Alarm{
					AlarmRules: &models.AlarmOrList{
						OrList: []*models.AlarmAndList{
							{
								AndList: []*models.AlarmExpression{
									{
										Operation: "range",
										Operand2: &models.AlarmOperand2{
											JSONValue: "[0, 1]",
										},
									},
								},
							},
						},
					},
				},
				FieldMask: rulesMask,
			},
		},
		{
			name: "Range operation, incorrect JSON",
			request: services.UpdateAlarmRequest{
				Alarm: &models.Alarm{
					AlarmRules: &models.AlarmOrList{
						OrList: []*models.AlarmAndList{
							{
								AndList: []*models.AlarmExpression{
									{
										Operation: "range",
										Operand2: &models.AlarmOperand2{
											JSONValue: "[1, 0]",
										},
									},
								},
							},
						},
					},
				},
				FieldMask: rulesMask,
			},
			errorCode: codes.InvalidArgument,
		},
		{
			name: "Range operation, invalid JSON",
			request: services.UpdateAlarmRequest{
				Alarm: &models.Alarm{
					AlarmRules: &models.AlarmOrList{
						OrList: []*models.AlarmAndList{
							{
								AndList: []*models.AlarmExpression{
									{
										Operation: "range",
										Operand2: &models.AlarmOperand2{
											JSONValue: "/0, 1]",
										},
									},
								},
							},
						},
					},
				},
				FieldMask: rulesMask,
			},
			errorCode: codes.InvalidArgument,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()
			service := makeMockedContrailTypeLogicService(mockCtrl)

			ctx := context.Background()
			expectedResponse := services.UpdateAlarmResponse{Alarm: tt.request.Alarm}

			updateAlarmCall := service.Next().(*servicesmock.MockService).EXPECT().UpdateAlarm(
				gomock.Not(gomock.Nil()), gomock.Not(gomock.Nil()),
			).DoAndReturn(
				func(_ context.Context, _ *services.UpdateAlarmRequest,
				) (response *services.UpdateAlarmResponse, err error) {
					return &services.UpdateAlarmResponse{Alarm: tt.request.Alarm}, nil
				},
			)

			if tt.errorCode != codes.OK {
				updateAlarmCall.MaxTimes(1)
			} else {
				updateAlarmCall.Times(1)
			}

			updateAlarmResponse, err := service.UpdateAlarm(ctx, &tt.request)
			if tt.errorCode != codes.OK {
				assert.Error(t, err)
				status, ok := status.FromError(err)
				assert.True(t, ok)
				assert.Equal(t, tt.errorCode, status.Code())
			} else {
				assert.NoError(t, err)
				assert.Equal(t, &expectedResponse, updateAlarmResponse)
			}
		})
	}
}
