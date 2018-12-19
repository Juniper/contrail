package logic

import (
	"testing"

	"github.com/Juniper/contrail/pkg/services"

	"github.com/Juniper/contrail/pkg/models"

	"github.com/stretchr/testify/assert"

	"github.com/Juniper/contrail/pkg/services/mock"
	"github.com/golang/mock/gomock"
)

func TestPortCreate(t *testing.T) {
	type readData struct {
		virtualNetworkReq *services.GetVirtualNetworkResponse
	}

	type writeData struct {
		vmiReq *services.CreateVirtualMachineInterfaceRequest
	}

	tests := []struct {
		name      string
		port      *Port
		wantErr   bool
		readData  *readData
		writeData *writeData
	}{
		{
			name: "Port simple create",
			port: &Port{
				Name: "hoge-hoge",
			},
			readData: &readData{
				virtualNetworkReq: &services.GetVirtualNetworkResponse{
					VirtualNetwork: &models.VirtualNetwork{},
				},
			},
			writeData: &writeData{
				vmiReq: &services.CreateVirtualMachineInterfaceRequest{
					VirtualMachineInterface: &models.VirtualMachineInterface{
						Name:       "hoge-hoge",
						ParentType: models.KindProject,
						IDPerms: &models.IdPermsType{
							Enable: true,
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()
			mockServ := servicesmock.NewMockService(mockCtrl)

			if tt.readData.virtualNetworkReq != nil {
				mockServ.EXPECT().GetVirtualNetwork(gomock.Any(), gomock.Any()).Return(tt.readData.virtualNetworkReq, nil)
			}

			if tt.writeData.vmiReq != nil {
				mockServ.EXPECT().CreateVirtualMachineInterface(gomock.Any(), tt.writeData.vmiReq)
			}

			rp := RequestParameters{
				ReadService:  mockServ,
				WriteService: mockServ,
			}

			_, err := tt.port.Create(nil, rp)

			if tt.wantErr {
				assert.Error(t, err)
			}
		})
	}
}
