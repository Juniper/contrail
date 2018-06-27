package db

import (
	"context"
	"testing"
	"time"

	"github.com/twinj/uuid"

	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/types/ipam"
	"github.com/stretchr/testify/assert"
)

func TestService_CreateIpamSubnet(t *testing.T) {

	tests := []struct {
		name    string
		request *ipam.CreateIpamSubnetRequest
		wantErr bool
	}{
		{
			name: "Create address manager with any subnetUUID",
			request: &ipam.CreateIpamSubnetRequest{
				IpamSubnet: &models.IpamSubnetType{},
			},
		},
		{
			name: "Create address manager with provided subnetUUID",
			request: &ipam.CreateIpamSubnetRequest{
				IpamSubnet: &models.IpamSubnetType{
					SubnetUUID: "uuid-1",
				},
			},
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
			defer cancel()
			err := db.DoInTransaction(ctx,
				func(ctx context.Context) error {
					gotSubnetUUID, err := db.CreateIpamSubnet(ctx, tt.request)
					assert.NoError(t, err)

					// Clean up after test
					_, err = GetTransaction(ctx).ExecContext(ctx, "delete from ipaddress_pool")
					assert.NoError(t, err)

					if tt.wantErr {
						assert.Error(t, err)
						return nil
					}

					if len(tt.request.IpamSubnet.SubnetUUID) > 0 {
						assert.Equal(t, tt.request.IpamSubnet.SubnetUUID, gotSubnetUUID)
					} else {
						_, err := uuid.Parse(gotSubnetUUID)
						assert.NoError(t, err)
					}
					return nil
				})
			assert.NoError(t, err)
		})
	}
}
