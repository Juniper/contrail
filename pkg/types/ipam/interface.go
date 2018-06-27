package ipam

import (
	"context"

	"github.com/Juniper/contrail/pkg/db"
)

// IntPoolAllocator integer pool allocator
type IntPoolAllocator interface {
	AllocateInt(context.Context, string) (int64, error)
	DeallocateInt(context.Context, string, int64) error
	CreateIntPool(context.Context, *db.IntPool) error
	DeleteIntPools(context.Context, *db.IntPool) error
}
