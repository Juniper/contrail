package db

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestIntPool(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	DoInTransaction(
		ctx,
		db.DB,
		func(ctx context.Context) error {
			poolKey := "testPool"
			err := db.CreateIntPool(ctx, poolKey, 0, 2)
			assert.NoError(t, err, "create pool failed")
			err = db.CreateIntPool(ctx, poolKey, 3, 5)
			assert.NoError(t, err, "create pool failed")
			pools, err := db.GetIntPools(ctx, poolKey)
			assert.Equal(t, 2, len(pools), "get pool failed")

			i, err := db.AllocateInt(ctx, poolKey)
			assert.NoError(t, err, "allocate failed")
			assert.Equal(t, 0, i, "allocate failed")

			i, err = db.AllocateInt(ctx, poolKey)
			assert.NoError(t, err, "allocate failed")
			assert.Equal(t, 1, i, "allocate failed")

			i, err = db.AllocateInt(ctx, poolKey)
			assert.NoError(t, err, "allocate failed")
			assert.Equal(t, 3, i, "allocate failed")

			pools, err = db.GetIntPools(ctx, poolKey)
			assert.Equal(t, 1, len(pools), "get pool failed")

			err = db.DeleteIntPools(ctx, poolKey)
			assert.NoError(t, err, "delete pool failed")

			pools, err = db.GetIntPools(ctx, poolKey)
			assert.Equal(t, 0, len(pools), "get pool failed")
			return nil
		})
}
