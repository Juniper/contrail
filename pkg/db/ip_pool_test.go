package db

import (
	"context"
	"net"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestIpPool(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)

	defer cancel()
	// Cleanup

	tests := []struct {
		name         string
		poolKey      string
		ipPools      []ipPool
		poolSize     int
		rangesNumber int

		allocateIPs []net.IP
		fails       bool
	}{
		{
			name: "Create IP Pool",
			ipPools: []ipPool{
				{
					key:   "subnet-uuid-1",
					start: net.ParseIP("10.0.0.1"),
					end:   net.ParseIP("10.0.0.10"),
				},
			},
			poolKey:  "subnet-uuid-1",
			poolSize: 1,
		},
	}
	for _, tt := range tests {
		dbTransaction := func(ctx context.Context) error {
			for _, pool := range tt.ipPools {
				err := db.createIPPool(ctx, &pool)
				assert.NoError(t, err, "create pool failed")
			}

			pools, err := db.getIPPools(ctx, &ipPool{key: tt.poolKey})
			assert.NoError(t, err)
			assert.Equal(t, 1, len(pools))

			GetTransaction(ctx).ExecContext(ctx, "delete from ipaddress_pool")
			return nil
		}
		t.Run(tt.name, func(t *testing.T) {
			DoInTransaction(
				ctx,
				db.DB(),
				dbTransaction)
		})
	}
}

func TestAllocateIp(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)

	defer cancel()
	// Cleanup

	tests := []struct {
		name    string
		poolKey string
		ipPools []ipPool

		expectedIp net.IP
		fails      bool
	}{
		{
			name: "Simple example, one pool",
			ipPools: []ipPool{
				{
					key:   "subnet-uuid-1",
					start: net.ParseIP("10.0.0.1"),
					end:   net.ParseIP("10.0.0.10"),
				},
			},
			poolKey:    "subnet-uuid-1",
			expectedIp: net.ParseIP("10.0.0.1"),
			fails:      false,
		},
		{
			name: "Several pools",
			ipPools: []ipPool{
				{
					key:   "subnet-uuid-1",
					start: net.ParseIP("10.0.0.1"),
					end:   net.ParseIP("10.0.0.10"),
				},
				{
					key:   "subnet-uuid-1",
					start: net.ParseIP("20.0.0.1"),
					end:   net.ParseIP("20.0.0.10"),
				},
			},
			poolKey:    "subnet-uuid-1",
			expectedIp: net.ParseIP("10.0.0.1"),
			fails:      false,
		},
		{
			name:    "Empty pool",
			poolKey: "subnet-uuid-1",
			fails:   true,
		},
	}
	for _, tt := range tests {
		dbTransaction := func(ctx context.Context) error {
			for _, pool := range tt.ipPools {
				err := db.createIPPool(ctx, &pool)
				assert.NoError(t, err, "create pool failed")
			}

			ipReceived, err := db.allocateIP(ctx, tt.poolKey)
			if tt.fails {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedIp.To4(), ipReceived.To4())
			}

			GetTransaction(ctx).ExecContext(ctx, "delete from ipaddress_pool")
			return nil
		}
		t.Run(tt.name, func(t *testing.T) {
			DoInTransaction(
				ctx,
				db.DB(),
				dbTransaction)
		})
	}
}

func TestDeleteIPPools(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)

	defer cancel()
	// Cleanup

	tests := []struct {
		name    string
		poolKey string
		ipPools []ipPool

		deletePool    ipPool
		expectedCount int
	}{
		{
			name: "Remove all pools",
			ipPools: []ipPool{
				{
					key:   "subnet-uuid-1",
					start: net.ParseIP("10.0.0.1"),
					end:   net.ParseIP("10.0.0.10"),
				},
				{
					key:   "subnet-uuid-1",
					start: net.ParseIP("11.0.0.1"),
					end:   net.ParseIP("11.0.0.10"),
				},
				{
					key:   "subnet-uuid-1",
					start: net.ParseIP("12.0.0.1"),
					end:   net.ParseIP("12.0.0.10"),
				},
			},
			deletePool: ipPool{
				key:   "subnet-uuid-1",
				start: net.ParseIP("12.0.0.1"),
			},
			poolKey:       "subnet-uuid-1",
			expectedCount: 0,
		},
		{
			name: "No overlapping pools",
			ipPools: []ipPool{
				{
					key:   "subnet-uuid-1",
					start: net.ParseIP("10.0.0.1"),
					end:   net.ParseIP("10.0.0.10"),
				},
				{
					key:   "subnet-uuid-1",
					start: net.ParseIP("11.0.0.1"),
					end:   net.ParseIP("11.0.0.10"),
				},
				{
					key:   "subnet-uuid-1",
					start: net.ParseIP("12.0.0.1"),
					end:   net.ParseIP("12.0.0.10"),
				},
			},
			deletePool: ipPool{
				key:   "subnet-uuid-1",
				start: net.ParseIP("13.0.0.1"),
				end:   net.ParseIP("13.0.0.10"),
			},
			poolKey:       "subnet-uuid-1",
			expectedCount: 3,
		},
		{
			name: "Two overlapping pools",
			ipPools: []ipPool{
				{
					key:   "subnet-uuid-1",
					start: net.ParseIP("10.0.0.1"),
					end:   net.ParseIP("10.0.0.10"),
				},
				{
					key:   "subnet-uuid-1",
					start: net.ParseIP("11.0.0.1"),
					end:   net.ParseIP("11.0.0.9"),
				},
				{
					key:   "subnet-uuid-1",
					start: net.ParseIP("12.0.0.1"),
					end:   net.ParseIP("12.0.0.10"),
				},
			},
			deletePool: ipPool{
				key:   "subnet-uuid-1",
				start: net.ParseIP("10.0.0.8"),
				end:   net.ParseIP("11.0.0.2"),
			},
			poolKey:       "subnet-uuid-1",
			expectedCount: 1,
		},
	}
	for _, tt := range tests {
		dbTransaction := func(ctx context.Context) error {
			for _, pool := range tt.ipPools {
				err := db.createIPPool(ctx, &pool)
				assert.NoError(t, err, "create pool failed")
			}

			pools, err := db.getIPPools(ctx, &ipPool{key: tt.poolKey})
			assert.NoError(t, err)

			err = db.deleteIPPools(ctx, &tt.deletePool)
			assert.NoError(t, err)

			pools, err = db.getIPPools(ctx, &ipPool{key: tt.poolKey})
			assert.NoError(t, err)
			assert.Equal(t, tt.expectedCount, len(pools))

			GetTransaction(ctx).ExecContext(ctx, "delete from ipaddress_pool")
			return nil
		}
		t.Run(tt.name, func(t *testing.T) {
			DoInTransaction(
				ctx,
				db.DB(),
				dbTransaction)
		})
	}
}
