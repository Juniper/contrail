package logic

import (
	"context"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/Juniper/contrail/pkg/fileutil"
	"github.com/Juniper/contrail/pkg/services"
	"github.com/Juniper/contrail/pkg/services/baseservices"
	"github.com/Juniper/contrail/pkg/services/mock"
)

func TestFloatingip_Create(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	read := servicesmock.NewMockReadService(mockCtrl)
	write := servicesmock.NewMockWriteService(mockCtrl)
	fqNameToID := servicesmock.NewMockFQNameToIDService(mockCtrl)

	tests := []struct {
		name          string
		request       Request
		expected      Response
		expectedCalls func(read *servicesmock.MockReadService, write *servicesmock.MockWriteService)
		fails         bool
	}{
		{
			name:  "empty request",
			fails: true,
		},
		{
			name:    "create logic",
			request: loadNeutronRequest(t, "test_data/floatingip/create_fip.json"),
			expectedCalls: func(read *servicesmock.MockReadService, write *servicesmock.MockWriteService) {
				read.EXPECT().ListFloatingIPPool(gomock.Any(), &services.ListFloatingIPPoolRequest{
					Spec: &baseservices.ListSpec{
						ParentUUIDs: []string{
							"0a673570-47eb-4b88-b648-5de06c65a37e",
						},
					},
				}).Return(loadListFloatingIPPoolResponse(t, "test_data/floatingip/list_floating-ip-pool_response.json"), nil)

				read.EXPECT().GetFloatingIPPool(gomock.Any(), &services.GetFloatingIPPoolRequest{
					ID: "c894d6c0-d8d2-403f-97a0-7d3e8fea207b",
				}).Return(loadGetFloatingIPPoolResponse(t, "test_data/floatingip/get_floating-ip-pool_response.json"), nil)

				read.EXPECT().GetProject(gomock.Any(), &services.GetProjectRequest{
					ID: "8a5e9e61-0938-4238-a4a5-fd5f23bebea4",
				}).Return(loadProjectResponse(t, "test_data/floatingip/get_project_response.json"), nil)

				write.EXPECT().CreateFloatingIP(gomock.Any(),
					&createFloatingIPRequestMatcher{
						expected: loadCreateFloatingIPRequest(t, "test_data/floatingip/create_floating-ip_request.json"),
					}).Return(loadCreateFloatingIPResponse(t, "test_data/floatingip/create_floating-ip_response.json"), nil)

				read.EXPECT().GetFloatingIP(gomock.Any(),
					&services.GetFloatingIPRequest{
						ID: "f4d63b5a-22e6-4aad-8b83-624b75a82e45",
					}).Return(loadGetFloatingIPResponse(t, "test_data/floatingip/get_floating-ip_response.json"), nil)

				fqNameToID.EXPECT().FQNameToID(gomock.Any(),
					loadFQNameToIDRequest(t, "test_data/floatingip/network_fqname-to-id_request.json"),
				).Return(loadFQNameToIDResponse(t, "test_data/floatingip/network_fqname-to-id_response.json"), nil)
			},
			expected: loadFloatingipResponse(t, "test_data/floatingip/create_fip_response.json"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got Response
			var err error
			ctx := context.Background()

			rp := RequestParameters{
				ReadService:       read,
				WriteService:      write,
				UserAgentKV:       nil,
				IDToFQNameService: nil,
				FQNameToIDService: fqNameToID,
				RequestContext:    tt.request.Context,
				FieldMask:         tt.request.Data.FieldMask,
			}

			if tt.expectedCalls != nil {
				tt.expectedCalls(read, write)
			}

			if tt.request.Data.Resource == nil && !tt.fails {
				t.Error("invalid test scenario, got nil Resource")
				return
			}

			switch tt.request.Context.Operation {
			case OperationCreate:
				got, err = tt.request.Data.Resource.Create(ctx, rp)
			case OperationUpdate:
				got, err = tt.request.Data.Resource.Update(ctx, rp, tt.request.Data.ID)
			case OperationDelete:
				got, err = tt.request.Data.Resource.Delete(ctx, rp, tt.request.Data.ID)
			case OperationRead:
				got, err = tt.request.Data.Resource.Read(ctx, rp, tt.request.Data.ID)
			case OperationReadAll:
				got, err = tt.request.Data.Resource.ReadAll(ctx, rp, tt.request.Data.Filters, tt.request.Data.Fields)
			case OperationReadCount:
				got, err = tt.request.Data.Resource.ReadCount(ctx, rp, tt.request.Data.Filters)
			case OperationAddInterface:
				got, err = tt.request.Data.Resource.AddInterface(ctx, rp)
			case OperationDelInterface:
				got, err = tt.request.Data.Resource.DeleteInterface(ctx, rp)
			default:
				err = errors.Errorf("invalid request operation: '%s'", tt.request.Context.Operation)
			}
			if tt.fails {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expected, got)
			}
		})
	}
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

func loadGetFloatingIPPoolResponse(t *testing.T, path string) (r *services.GetFloatingIPPoolResponse) {
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

func loadCreateFloatingIPRequest(t *testing.T, path string) (r *services.CreateFloatingIPRequest) {
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
