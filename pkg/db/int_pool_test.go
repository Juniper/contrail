package db

import (
	"context"
	"testing"
	"time"

	"github.com/Juniper/contrail/pkg/types/ipam"
	"github.com/stretchr/testify/assert"
)

func TestIntPool(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	err := db.DoInTransaction(
		ctx,
		func(ctx context.Context) error {
			poolKey := "testPool"

			err := db.DeleteIntPools(ctx, &ipam.IntPool{
				Key: poolKey,
			})
			assert.NoError(t, err, "clear pool failed")

			err = db.CreateIntPool(ctx, &ipam.IntPool{Key: poolKey, Start: 0, End: 2})
			assert.NoError(t, err, "create pool failed")

			err = db.CreateIntPool(ctx, &ipam.IntPool{Key: poolKey, Start: 3, End: 5})
			assert.NoError(t, err, "create pool failed")

			pools, err := db.GetIntPools(ctx, &ipam.IntPool{Key: poolKey})
			assert.NoError(t, err)
			assert.Equal(t, 2, len(pools), "get pool failed")

			size, err := db.SizeIntPool(ctx, poolKey)
			assert.NoError(t, err)
			assert.Equal(t, 4, size, "size pool failed")

			i, err := db.AllocateInt(ctx, poolKey)
			assert.NoError(t, err, "allocate failed")
			assert.Equal(t, int64(0), i, "allocate failed")

			i, err = db.AllocateInt(ctx, poolKey)
			assert.NoError(t, err, "allocate failed")
			assert.Equal(t, int64(1), i, "allocate failed")

			i, err = db.AllocateInt(ctx, poolKey)
			assert.NoError(t, err, "allocate failed")
			assert.Equal(t, int64(3), i, "allocate failed")

			size, err = db.SizeIntPool(ctx, poolKey)
			assert.NoError(t, err)
			assert.Equal(t, 1, size, "size pool failed")

			pools, err = db.GetIntPools(ctx, &ipam.IntPool{Key: poolKey})
			assert.NoError(t, err)
			assert.Equal(t, 1, len(pools), "get pool failed")

			err = db.DeallocateInt(ctx, poolKey, 0)
			assert.NoError(t, err, "deallocate failed")

			err = db.DeallocateInt(ctx, poolKey, 3)
			assert.NoError(t, err, "deallocate failed")

			pools, err = db.GetIntPools(ctx, &ipam.IntPool{Key: poolKey})
			assert.NoError(t, err)
			assert.Equal(t, 2, len(pools), "get pool failed")

			size, err = db.SizeIntPool(ctx, poolKey)
			assert.NoError(t, err)
			assert.Equal(t, 3, size, "size pool failed")

			err = db.SetInt(ctx, poolKey, 4)
			assert.NoError(t, err, "set failed")

			pools, err = db.GetIntPools(ctx, &ipam.IntPool{Key: poolKey})
			assert.NoError(t, err)
			assert.Equal(t, 2, len(pools), "get pool failed")

			size, err = db.SizeIntPool(ctx, poolKey)
			assert.NoError(t, err)
			assert.Equal(t, 2, size, "size pool failed")

			err = db.DeleteIntPools(ctx, &ipam.IntPool{Key: poolKey})
			assert.NoError(t, err, "delete pool failed")

			pools, err = db.GetIntPools(ctx, &ipam.IntPool{Key: poolKey})
			assert.NoError(t, err)
			assert.Equal(t, 0, len(pools), "get pool failed")
			return nil
		})
	assert.NoError(t, err)
}
