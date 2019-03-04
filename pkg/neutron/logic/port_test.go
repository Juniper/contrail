package logic

import (
	"context"
	"testing"

	"github.com/gogo/protobuf/types"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/Juniper/contrail/pkg/logutil"
	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/services"
	servicesmock "github.com/Juniper/contrail/pkg/services/mock"
)

func TestPortRead(t *testing.T) {
	type readData struct {
		vnReq  *services.GetVirtualNetworkResponse
		vmiReq *services.GetVirtualMachineInterfaceResponse
	}

	tests := []struct {
		name     string
		port     *Port
		expected *PortResponse
		wantErr  bool
		id       string
		readData *readData
	}{
		{
			name: "Read port",
			port: &Port{},
			expected: &PortResponse{
				Status:    "ACTIVE",
				CreatedAt: "2018-12-06T11:30:11.787306",
				UpdatedAt: "2018-12-06T11:30:11.877835",
				BindingVifDetails: BindingVifDetails{
					PortFilter: true,
				},
				BindingVnicType: "normal",
				BindingVifType:  "vrouter",
				Name:            "default-vmi",
				DeviceOwner:     "compute:nova",
				NetworkID:       "623666e6-3929-4cb9-bedb-1dd98f63c569",
				TenantID:        "9933f4ed73f742f9a2bfe4bf4dd5f4df",
				MacAddress:      "00:0A:E6:3E:FD:EF",
				FQName:          []string{"default-project", "default-vmi"},
				FixedIps: []*FixedIp{
					{SubnetID: "a46ff943-72cd-41dc-b92b-a997c1287856", IPAddress: "10.0.1.3"},
				},
				ID:             "b6283c9b-07ec-4061-941e-3f392844059f",
				SecurityGroups: []string{"79ce33bf-1bac-48d5-8bbb-5782e26b3db8"},
				DeviceID:       "d51094ce-3f69-11e9-8faf-d70c75a42a13",
			},
			readData: &readData{
				vnReq: &services.GetVirtualNetworkResponse{
					VirtualNetwork: &models.VirtualNetwork{
						UUID: "623666e6-3929-4cb9-bedb-1dd98f63c569",
					},
				},
				vmiReq: &services.GetVirtualMachineInterfaceResponse{
					VirtualMachineInterface: &models.VirtualMachineInterface{
						Name:                               "hoge-hoge",
						UUID:                               "b6283c9b-07ec-4061-941e-3f392844059f",
						FQName:                             []string{"default-project", "default-vmi"},
						ParentType:                         models.KindProject,
						VirtualMachineInterfaceDeviceOwner: "compute:nova",
						VirtualMachineInterfaceMacAddresses: &models.MacAddressesType{
							MacAddress: []string{"00:0A:E6:3E:FD:EF"},
						},
						Perms2: &models.PermType2{
							Owner: "9933f4ed73f742f9a2bfe4bf4dd5f4df",
						},
						ParentUUID: "9933f4ed-73f7-42f9-a2bf-e4bf4dd5f4df",
						IDPerms: &models.IdPermsType{
							Enable:       true,
							Created:      "2018-12-06T11:30:11.787306",
							LastModified: "2018-12-06T11:30:11.877835",
						},
						VirtualNetworkRefs: []*models.VirtualMachineInterfaceVirtualNetworkRef{
							{UUID: "623666e6-3929-4cb9-bedb-1dd98f63c569"},
						},
						VirtualMachineRefs: []*models.VirtualMachineInterfaceVirtualMachineRef{
							{UUID: "12321431-3242-1234-bedb-4dd38f63c569", To: []string{"default", "d51094ce-3f69-11e9-8faf-d70c75a42a13"}},
						},
						VirtualMachineInterfaceBindings: &models.KeyValuePairs{
							KeyValuePair: []*models.KeyValuePair{
								{Key: "vnic_type", Value: "normal"},
								{Key: "vif_type", Value: "vrouter"},
								{Key: "hoge", Value: "hoge"},
							},
						},
						SecurityGroupRefs: []*models.VirtualMachineInterfaceSecurityGroupRef{
							{UUID: "79ce33bf-1bac-48d5-8bbb-5782e26b3db8"},
						},
						InstanceIPBackRefs: []*models.InstanceIP{
							{InstanceIPAddress: "10.0.1.3", SubnetUUID: "a46ff943-72cd-41dc-b92b-a997c1287856"},
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

			if tt.readData.vnReq != nil {
				mockServ.EXPECT().GetVirtualNetwork(gomock.Any(), gomock.Any()).Return(
					tt.readData.vnReq, nil,
				)
			}
			if tt.readData.vmiReq != nil {
				mockServ.EXPECT().GetVirtualMachineInterface(gomock.Any(), gomock.Any()).Return(
					tt.readData.vmiReq, nil,
				)
			}

			rp := RequestParameters{
				ReadService:  mockServ,
				WriteService: mockServ,
			}

			readRes, err := tt.port.Read(nil, rp, tt.id)

			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.Equal(t, tt.expected, readRes)
			}
		})
	}
}

func TestPortUpdate(t *testing.T) {
	type readData struct {
		vnResp  *services.GetVirtualNetworkResponse
		vmiResp *services.GetVirtualMachineInterfaceResponse
	}

	type updateData struct {
		vmiReq *services.UpdateVirtualMachineInterfaceRequest
	}

	type writeData struct {
		vmReq  *services.CreateVirtualMachineRequest
		vmResp *services.CreateVirtualMachineResponse
	}

	const portID = "b6283c9b-07ec-4061-941e-3f392844059f"
	const deviceID = "d51094ce-3f69-11e9-8faf-d70c75a42a13"
	const networkID = "623666e6-3929-4cb9-bedb-1dd98f63c569"
	const tenantID = "9933f4ed73f742f9a2bfe4bf4dd5f4df"

	tests := []struct {
		name       string
		port       *Port
		mask       types.FieldMask
		expected   *PortResponse
		id         string
		readData   *readData
		updateData *updateData
		writeData  *writeData
	}{
		{
			name: "Update name, device ID, binding host IP and device owner",
			port: &Port{
				Name:          "hoge-hoge",
				DeviceID:      deviceID,
				BindingHostID: "ignacy.osetek-spike.novalocal",
				DeviceOwner:   "compute:nova",
			},
			mask: types.FieldMask{
				Paths: []string{
					"data.resource." + PortFieldName,
					"data.resource." + PortFieldID,
					"data.resource." + PortFieldDeviceID,
					"data.resource." + PortFieldBindingHostID,
					"data.resource." + PortFieldDeviceOwner,
				},
			},
			id: portID,
			expected: &PortResponse{
				Status: "ACTIVE",
				BindingVifDetails: BindingVifDetails{
					PortFilter: true,
				},
				DeviceOwner:     "compute:nova",
				BindingHostID:   "ignacy.osetek-spike.novalocal",
				BindingVnicType: "normal",
				BindingVifType:  "vrouter",
				Name:            "hoge-hoge",
				NetworkID:       networkID,
				TenantID:        tenantID,
				MacAddress:      "00:0A:E6:3E:FD:EF",
				FQName:          []string{"default-project", "default-vmi"},
				FixedIps: []*FixedIp{
					{SubnetID: "a46ff943-72cd-41dc-b92b-a997c1287856", IPAddress: "10.0.1.3"},
				},
				ID:             portID,
				SecurityGroups: []string{"79ce33bf-1bac-48d5-8bbb-5782e26b3db8"},
				DeviceID:       deviceID,
			},
			updateData: &updateData{
				vmiReq: &services.UpdateVirtualMachineInterfaceRequest{
					FieldMask: types.FieldMask{
						Paths: []string{
							models.VirtualMachineInterfaceFieldDisplayName,
							models.VirtualMachineInterfaceFieldVirtualMachineRefs,
							models.VirtualMachineInterfaceFieldVirtualMachineInterfaceDeviceOwner,
							models.VirtualMachineInterfaceFieldVirtualMachineInterfaceBindings,
						},
					},
				},
			},
			writeData: &writeData{
				vmReq: &services.CreateVirtualMachineRequest{
					VirtualMachine: &models.VirtualMachine{
						UUID:       deviceID,
						Name:       deviceID,
						ServerType: "virtual-server",
						Perms2: &models.PermType2{
							Owner: tenantID,
						},
					},
				},
				vmResp: &services.CreateVirtualMachineResponse{
					VirtualMachine: &models.VirtualMachine{
						Name:       deviceID,
						ServerType: "virtual-server",
						Perms2: &models.PermType2{
							Owner: tenantID,
						},
						FQName: []string{"default", deviceID},
					},
				},
			},
			readData: &readData{
				vnResp: &services.GetVirtualNetworkResponse{
					VirtualNetwork: &models.VirtualNetwork{
						UUID:                networkID,
						PortSecurityEnabled: true,
					},
				},
				vmiResp: &services.GetVirtualMachineInterfaceResponse{
					VirtualMachineInterface: &models.VirtualMachineInterface{
						Name:       "hoge-hoge",
						UUID:       portID,
						FQName:     []string{"default-project", "default-vmi"},
						ParentType: models.KindProject,
						Perms2: &models.PermType2{
							Owner: tenantID,
						},
						VirtualMachineInterfaceMacAddresses: &models.MacAddressesType{
							MacAddress: []string{"00:0A:E6:3E:FD:EF"},
						},
						ParentUUID: "9933f4ed-73f7-42f9-a2bf-e4bf4dd5f4df",
						IDPerms: &models.IdPermsType{
							Enable: true,
						},
						VirtualNetworkRefs: []*models.VirtualMachineInterfaceVirtualNetworkRef{
							{UUID: networkID},
						},
						VirtualMachineInterfaceBindings: &models.KeyValuePairs{
							KeyValuePair: []*models.KeyValuePair{
								{Key: "vnic_type", Value: "normal"},
								{Key: "vif_type", Value: "vrouter"},
							},
						},
						SecurityGroupRefs: []*models.VirtualMachineInterfaceSecurityGroupRef{
							{UUID: "79ce33bf-1bac-48d5-8bbb-5782e26b3db8"},
						},
						InstanceIPBackRefs: []*models.InstanceIP{
							{InstanceIPAddress: "10.0.1.3", SubnetUUID: "a46ff943-72cd-41dc-b92b-a997c1287856"},
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
			service := servicesmock.NewMockService(mockCtrl)

			var call *gomock.Call
			if tt.readData.vmiResp != nil {
				call = service.EXPECT().GetVirtualMachineInterface(gomock.Any(), gomock.Any()).Return(
					tt.readData.vmiResp, nil,
				).Times(1)
			}
			if tt.readData.vnResp != nil {
				call = service.EXPECT().GetVirtualNetwork(gomock.Any(), gomock.Any()).Return(
					tt.readData.vnResp, nil,
				).Times(1).After(call)
			}
			call = service.EXPECT().GetVirtualMachine(
				gomock.Any(),
				gomock.Eq(&services.GetVirtualMachineRequest{ID: deviceID}),
			).Return(
				nil, status.Error(codes.NotFound, "not found"),
			).Times(1).After(call)
			if tt.writeData.vmReq != nil {
				call = service.EXPECT().CreateVirtualMachine(gomock.Any(), tt.writeData.vmReq).Return(
					tt.writeData.vmResp, nil,
				).Times(1).After(call)
			}
			if tt.updateData.vmiReq != nil {
				service.EXPECT().UpdateVirtualMachineInterface(gomock.Any(), gomock.Any()).DoAndReturn(
					func(
						_ context.Context,
						vmiReq *services.UpdateVirtualMachineInterfaceRequest,
					) (*services.UpdateVirtualMachineInterfaceResponse, error) {
						assert.NotNil(t, vmiReq.GetVirtualMachineInterface())
						assert.Equal(t, tt.updateData.vmiReq.GetFieldMask(), vmiReq.GetFieldMask())
						return &services.UpdateVirtualMachineInterfaceResponse{
							VirtualMachineInterface: vmiReq.GetVirtualMachineInterface(),
						}, nil
					},
				).Times(1).After(call)
			}

			r, err := tt.port.Update(
				nil,
				RequestParameters{
					ReadService:  service,
					WriteService: service,
					FieldMask:    tt.mask,
					Log:          logutil.NewLogger("test"),
				},
				tt.id,
			)

			assert.NoError(t, err)
			assert.Equal(t, tt.expected, r)
		})
	}
}

func TestPortDelete(t *testing.T) {
	type readData struct {
		vnReq  *services.GetVirtualNetworkResponse
		vmiReq *services.GetVirtualMachineInterfaceResponse
	}

	type deleteData struct {
		vmiReq *services.DeleteVirtualMachineInterfaceRequest
		iipReq *services.DeleteInstanceIPRequest
		vmReq  *services.DeleteVirtualMachineRequest
	}

	tests := []struct {
		name       string
		port       *Port
		expected   *PortResponse
		wantErr    bool
		id         string
		readData   *readData
		deleteData *deleteData
	}{
		{
			name:     "Delete port",
			port:     &Port{},
			expected: &PortResponse{},
			id:       "b6283c9b-07ec-4061-941e-3f392844059f",
			deleteData: &deleteData{
				vmiReq: &services.DeleteVirtualMachineInterfaceRequest{ID: "b6283c9b-07ec-4061-941e-3f392844059f"},
				iipReq: &services.DeleteInstanceIPRequest{ID: "a21186f0-d871-4ab4-b63c-cd8b27b556f0"},
				vmReq:  &services.DeleteVirtualMachineRequest{ID: "12321431-3242-1234-bedb-4dd38f63c569"},
			},
			readData: &readData{
				vnReq: &services.GetVirtualNetworkResponse{
					VirtualNetwork: &models.VirtualNetwork{
						UUID:                "623666e6-3929-4cb9-bedb-1dd98f63c569",
						PortSecurityEnabled: true,
					},
				},
				vmiReq: &services.GetVirtualMachineInterfaceResponse{
					VirtualMachineInterface: &models.VirtualMachineInterface{
						Name:       "hoge-hoge",
						UUID:       "b6283c9b-07ec-4061-941e-3f392844059f",
						FQName:     []string{"default-project", "default-vmi"},
						ParentType: models.KindProject,
						ParentUUID: "9933f4ed-73f7-42f9-a2bf-e4bf4dd5f4df",
						IDPerms: &models.IdPermsType{
							Enable: true,
						},
						VirtualNetworkRefs: []*models.VirtualMachineInterfaceVirtualNetworkRef{
							{UUID: "623666e6-3929-4cb9-bedb-1dd98f63c569"},
						},
						VirtualMachineRefs: []*models.VirtualMachineInterfaceVirtualMachineRef{
							{UUID: "12321431-3242-1234-bedb-4dd38f63c569", To: []string{"default", "d51094ce-3f69-11e9-8faf-d70c75a42a13"}},
						},
						InstanceIPBackRefs: []*models.InstanceIP{
							{
								UUID:              "a21186f0-d871-4ab4-b63c-cd8b27b556f0",
								InstanceIPAddress: "10.0.1.3",
								SubnetUUID:        "a46ff943-72cd-41dc-b92b-a997c1287856",
							},
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

			if tt.readData.vnReq != nil {
				mockServ.EXPECT().GetVirtualNetwork(gomock.Any(), gomock.Any()).Return(
					tt.readData.vnReq, nil,
				)
			}
			if tt.readData.vmiReq != nil {
				mockServ.EXPECT().GetVirtualMachineInterface(gomock.Any(), gomock.Any()).Return(
					tt.readData.vmiReq, nil,
				)
			}
			if tt.deleteData.vmiReq != nil {
				mockServ.EXPECT().DeleteVirtualMachineInterface(gomock.Any(), tt.deleteData.vmiReq)
			}
			if tt.deleteData.iipReq != nil {
				mockServ.EXPECT().DeleteInstanceIP(gomock.Any(), tt.deleteData.iipReq)
			}
			if tt.deleteData.vmReq != nil {
				mockServ.EXPECT().DeleteVirtualMachine(gomock.Any(), tt.deleteData.vmReq)
			}

			rp := RequestParameters{
				ReadService:  mockServ,
				WriteService: mockServ,
			}

			readRes, err := tt.port.Delete(nil, rp, tt.id)

			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.Equal(t, tt.expected, readRes)
			}
		})
	}
}
