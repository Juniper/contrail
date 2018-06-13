package ipam

import (
	"context"
)

//IntPool for int pool.
type IntPool struct {
	Key   string
	Start int64
	End   int64
}

// IntPoolAllocator integer pool allocator
type IntPoolAllocator interface {
	AllocateInt(context.Context, string) (int64, error)
	DeallocateInt(context.Context, string, int64) error
	CreateIntPool(context.Context, *IntPool) error
	DeleteIntPools(context.Context, *IntPool) error
}
