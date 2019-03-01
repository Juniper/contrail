package logic

import (
	"context"
	"reflect"
	"testing"

	"github.com/Juniper/contrail/pkg/format"
	"github.com/Juniper/contrail/pkg/models/basemodels"
	"github.com/Juniper/contrail/pkg/services/mock"
	"github.com/golang/mock/gomock"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/Juniper/contrail/pkg/fileutil"
	"github.com/Juniper/contrail/pkg/services"
	"github.com/Juniper/contrail/pkg/services/baseservices"
)

const (
	testDataDirectory = "test_data/floatingip/"
	createDirectory   = testDataDirectory + "create/"
	readDirectory     = testDataDirectory + "read/"
	updateDirectory   = testDataDirectory + "update/"
)

type neutronTestCase struct {
	name          string
	request       Request
	expected      Response
	expectedCalls func()
	fails         bool
}

func TestFloatingip_Create(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	read := servicesmock.NewMockReadService(mockCtrl)
	write := servicesmock.NewMockWriteService(mockCtrl)
	fqNameToID := servicesmock.NewMockFQNameToIDService(mockCtrl)
	idToFQName := servicesmock.NewMockIDToFQNameService(mockCtrl)

	tests := []*neutronTestCase{
		{
			name:  "empty request",
			fails: true,
		},
		{
			name:    "create",
			request: loadNeutronRequest(t, createDirectory+"create_fip.json"),
			expectedCalls: func() {
				read.EXPECT().ListFloatingIPPool(gomock.Any(), &services.ListFloatingIPPoolRequest{
					Spec: &baseservices.ListSpec{
						ParentUUIDs: []string{
							"0a673570-47eb-4b88-b648-5de06c65a37e",
						},
					},
				}).Return(loadListFloatingIPPoolResponse(t, createDirectory+"list_floating-ip-pool_response.json"), nil)

				read.EXPECT().GetProject(gomock.Any(), &services.GetProjectRequest{
					ID: "8a5e9e61-0938-4238-a4a5-fd5f23bebea4",
				}).Return(loadProjectResponse(t, createDirectory+"get_project_response.json"), nil)

				write.EXPECT().CreateFloatingIP(gomock.Any(),
					&createFloatingIPRequestMatcher{
						expected: loadCreateFloatingIPRequest(t, createDirectory+"create_floating-ip_request.json"),
					}).Return(loadCreateFloatingIPResponse(t, createDirectory+"create_floating-ip_response.json"), nil)

				fqNameToID.EXPECT().FQNameToID(gomock.Any(),
					loadFQNameToIDRequest(t, createDirectory+"network_fqname-to-id_request.json"),
				).Return(loadFQNameToIDResponse(t, createDirectory+"network_fqname-to-id_response.json"), nil)
			},
			expected: loadFloatingipResponse(t, createDirectory+"create_fip_response.json"),
		},
		{
			name:    "read",
			request: loadNeutronRequest(t, readDirectory+"read_fip.json"),
			expectedCalls: func() {
				read.EXPECT().GetFloatingIP(gomock.Any(),
					&services.GetFloatingIPRequest{
						ID: "f4d63b5a-22e6-4aad-8b83-624b75a82e45",
					}).Return(loadGetFloatingIPResponse(t, readDirectory+"get_floating-ip_response.json"), nil)

				fqNameToID.EXPECT().FQNameToID(gomock.Any(),
					loadFQNameToIDRequest(t, readDirectory+"network_fqname-to-id_request.json"),
				).Return(loadFQNameToIDResponse(t, readDirectory+"network_fqname-to-id_response.json"), nil)
			},
			expected: loadFloatingipResponse(t, readDirectory+"read_fip_response.json"),
		},
		{
			name:    "update",
			request: loadNeutronRequest(t, updateDirectory+"update_fip.json"),
			expectedCalls: func() {
				idToFQName.EXPECT().IDToFQName(gomock.Any(), gomock.Any()).Return(
					&services.IDToFQNameResponse{
						FQName: []string{
							"default-domain",
							"ctest-FloatingipBasicTestSanity-99684235",
							"5c5829af-8331-4e19-b3c3-d307ec619e95",
						},
					}, nil)

				read.EXPECT().GetVirtualMachineInterface(gomock.Any(), &services.GetVirtualMachineInterfaceRequest{
					ID:     "5c5829af-8331-4e19-b3c3-d307ec619e95",
					Fields: []string{"instance_ip_back_refs", "floating_ip_back_refs"},
				}).Return(loadVirtualMachineInterfaceResponse(t, updateDirectory+"get_vmi_response.json"), nil)

				write.EXPECT().UpdateFloatingIP(gomock.Any(),
					loadUpdateFloatingIPRequest(t, updateDirectory+"update_floating-ip_request.json"))

				read.EXPECT().GetFloatingIP(gomock.Any(),
					&services.GetFloatingIPRequest{
						ID: "f4d63b5a-22e6-4aad-8b83-624b75a82e45",
					}).Return(loadGetFloatingIPResponse(t, updateDirectory+"get_floating-ip_response.json"), nil)

				read.EXPECT().GetVirtualMachineInterface(gomock.Any(), &services.GetVirtualMachineInterfaceRequest{
					ID: "5c5829af-8331-4e19-b3c3-d307ec619e95",
				}).Return(loadVirtualMachineInterfaceResponse(t, updateDirectory+"get_vmi_response.json"), nil)

				fqNameToID.EXPECT().FQNameToID(gomock.Any(),
					loadFQNameToIDRequest(t, updateDirectory+"network_fqname-to-id_request.json"),
				).Return(loadFQNameToIDResponse(t, updateDirectory+"network_fqname-to-id_response.json"), nil)
			},
			expected: loadFloatingipResponse(t, updateDirectory+"update_fip_response.json"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()

			rp := RequestParameters{
				ReadService:       read,
				WriteService:      write,
				UserAgentKV:       nil,
				IDToFQNameService: idToFQName,
				FQNameToIDService: fqNameToID,
				RequestContext:    tt.request.Context,
				FieldMask:         tt.request.Data.FieldMask,
			}

			tt.run(t, ctx, rp)
		})
	}
}

//nolint: golint
func (tt *neutronTestCase) run(t *testing.T, ctx context.Context, rp RequestParameters) {
	var got Response
	var err error
	if tt.expectedCalls != nil {
		tt.expectedCalls()
	}

	if tt.request.Data.Resource == nil && !tt.fails {
		t.Error("invalid test scenario, got nil Resource")
		return
	}
	got, err = processRequest(t, &tt.request, ctx, rp)
	if tt.fails {
		assert.Error(t, err)
	} else {
		assert.NoError(t, err)
		assert.Equal(t, tt.expected, got)
	}
}

//nolint: golint
func processRequest(t *testing.T, r *Request, ctx context.Context, rp RequestParameters) (got Response, err error) {
	switch r.Context.Operation {
	case OperationCreate:
		got, err = r.Data.Resource.Create(ctx, rp)
	case OperationUpdate:
		got, err = r.Data.Resource.Update(ctx, rp, r.Data.ID)
	case OperationDelete:
		got, err = r.Data.Resource.Delete(ctx, rp, r.Data.ID)
	case OperationRead:
		got, err = r.Data.Resource.Read(ctx, rp, r.Data.ID)
	case OperationReadAll:
		got, err = r.Data.Resource.ReadAll(ctx, rp, r.Data.Filters, r.Data.Fields)
	case OperationReadCount:
		got, err = r.Data.Resource.ReadCount(ctx, rp, r.Data.Filters)
	case OperationAddInterface:
		got, err = r.Data.Resource.AddInterface(ctx, rp)
	case OperationDelInterface:
		got, err = r.Data.Resource.DeleteInterface(ctx, rp)
	default:
		err = errors.Errorf("invalid request operation: '%s'", r.Context.Operation)
	}
	return got, err
}

type createFloatingIPRequestMatcher struct {
	expected *services.CreateFloatingIPRequest
}

func (m *createFloatingIPRequestMatcher) Matches(x interface{}) bool {
	m.expected.FloatingIP.SetUUID("")
	got, ok := x.(*services.CreateFloatingIPRequest)
	if !ok {
		return false
	}
	got.FloatingIP.SetUUID("")
	reflect.DeepEqual(m.expected, got)
	return true
}

func (m *createFloatingIPRequestMatcher) String() string {
	return "this matchers matches requests ignoring resources' uuids"
}

func loadNeutronRequest(t *testing.T, path string) (r Request) {
	var requestMap map[string]interface{}
	require.NoError(t, fileutil.LoadFile(path, &requestMap))
	require.NoError(t, format.ApplyMap(requestMap, &r))
	r.Data.FieldMask = basemodels.MapToFieldMask(requestMap)
	require.NoError(t, fileutil.LoadFile(path, &r))
	return r
}

func loadFloatingipResponse(t *testing.T, path string) (r *FloatingipResponse) {
	require.NoError(t, fileutil.LoadFile(path, &r))
	return r
}

func loadListFloatingIPPoolResponse(t *testing.T, path string) (r *services.ListFloatingIPPoolResponse) {
	require.NoError(t, fileutil.LoadFile(path, &r))
	return r
}

func loadGetFloatingIPResponse(t *testing.T, path string) (r *services.GetFloatingIPResponse) {
	require.NoError(t, fileutil.LoadFile(path, &r))
	return r
}

func loadProjectResponse(t *testing.T, path string) (r *services.GetProjectResponse) {
	require.NoError(t, fileutil.LoadFile(path, &r))
	return r
}

func loadVirtualMachineInterfaceResponse(t *testing.T, path string) (r *services.GetVirtualMachineInterfaceResponse) {
	require.NoError(t, fileutil.LoadFile(path, &r))
	return r
}

func loadCreateFloatingIPRequest(t *testing.T, path string) (r *services.CreateFloatingIPRequest) {
	require.NoError(t, fileutil.LoadFile(path, &r))
	return r
}
func loadUpdateFloatingIPRequest(t *testing.T, path string) (r *services.UpdateFloatingIPRequest) {
	require.NoError(t, fileutil.LoadFile(path, &r))
	return r
}

func loadCreateFloatingIPResponse(t *testing.T, path string) (r *services.CreateFloatingIPResponse) {
	require.NoError(t, fileutil.LoadFile(path, &r))
	return r
}

func loadFQNameToIDRequest(t *testing.T, path string) (r *services.FQNameToIDRequest) {
	require.NoError(t, fileutil.LoadFile(path, &r))
	return r
}

func loadFQNameToIDResponse(t *testing.T, path string) (r *services.FQNameToIDResponse) {
	require.NoError(t, fileutil.LoadFile(path, &r))
	return r
}
