package services_test

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/Juniper/contrail/pkg/common"
	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/services"
	"github.com/Juniper/contrail/pkg/services/mock"
)

func TestChain(t *testing.T) {
	s := []services.Service{
		&services.BaseService{},
		&services.BaseService{},
		&services.BaseService{},
	}

	services.Chain(s...)

	assert.Equal(t, s[0].Next(), s[1])
	assert.Equal(t, s[1].Next(), s[2])
	assert.Equal(t, s[2].Next(), nil)
}

func TestGetObject(t *testing.T) {
	tests := []struct {
		name     string
		initMock func(services.ReadService)
		typename string
		uuid     string
		want     models.Object
		fails    bool
	}{
		{name: "empty", fails: true},
		{name: "unknown type", typename: "does_not_exist", uuid: "some-uuid", fails: true},
		{
			name:     "try to get non existing virtual network",
			typename: "virtual_network",
			uuid:     "some-uuid",
			initMock: func(r services.ReadService) {
				rsMock := r.(*servicesmock.MockReadService)
				rsMock.EXPECT().GetVirtualNetwork(
					gomock.Not(gomock.Nil()), gomock.Not(gomock.Nil()),
				).Return((*services.GetVirtualNetworkResponse)(nil), common.ErrorNotFound).Times(1)
			},
			fails: true,
		},
		{
			name:     "get existing virtual network",
			typename: "virtual_network",
			uuid:     "some-uuid",
			initMock: func(r services.ReadService) {
				rsMock := r.(*servicesmock.MockReadService)
				rsMock.EXPECT().GetVirtualNetwork(
					gomock.Not(gomock.Nil()), gomock.Not(gomock.Nil()),
				).Return(
					&services.GetVirtualNetworkResponse{VirtualNetwork: &models.VirtualNetwork{
						UUID: "some-uuid",
					}},
					nil,
				).Times(1)
			},
			want: &models.VirtualNetwork{
				UUID: "some-uuid",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()
			rs := servicesmock.NewMockReadService(mockCtrl)
			if tt.initMock != nil {
				tt.initMock(rs)
			}

			got, err := services.GetObject(context.Background(), rs, tt.typename, tt.uuid)
			if tt.fails {
				assert.Error(t, err)
				assert.Nil(t, got)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.want, got)
			}
		})
	}
}
