package logic

import (
	"context"
	"testing"

	"github.com/Juniper/contrail/pkg/services"
	"github.com/Juniper/contrail/pkg/services/baseservices"
	"github.com/golang/mock/gomock"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/Juniper/contrail/pkg/fileutil"
	"github.com/Juniper/contrail/pkg/services/mock"
)

func TestFloatingip_Create(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	read := servicesmock.NewMockReadService(mockCtrl)
	write := servicesmock.NewMockWriteService(mockCtrl)

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
			request: loadNeutronRequest(t, "test_data/create_fip.json"),
			expectedCalls: func(read *servicesmock.MockReadService, write *servicesmock.MockWriteService) {
				read.EXPECT().ListFloatingIPPool(gomock.Any(), &services.ListFloatingIPPoolRequest{
					Spec: &baseservices.ListSpec{
						ParentUUIDs: []string{
							"0a673570-47eb-4b88-b648-5de06c65a37e",
						},
					},
				}).Return(nil, errors.New("to be continued"))
			},
			fails: true,
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

func loadNeutronRequest(t *testing.T, path string) (r Request) {
	require.NoError(t, fileutil.LoadFile(path, &r))
	return r
}
