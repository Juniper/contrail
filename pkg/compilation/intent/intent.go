package intent

import (
	"context"

	"github.com/Juniper/contrail/pkg/services"
)

// EvaluateContext contains context information for Resource to handle CRUD
type EvaluateContext struct {
	WriteService services.WriteService
	IntentLoader Loader
}

// Intent contains Intent Compiler state for a resource.
type Intent interface {
	services.Resource
	Evaluate(ctx context.Context, evaluateCtx *EvaluateContext) error
}

// Loader provides access to existing intents (e.g. using a cache)
type Loader interface {
	LoadByFQName(typeName string, fqName []string) (Intent, bool)
	Load(typeName, uuid string) (Intent, bool)
}

// BaseIntent implements the default Evaluate interface
type BaseIntent struct {
}

// Evaluate creates/updates/deletes lower-level resources when needed.
func (b *BaseIntent) Evaluate(ctx context.Context, evaluateCtx *EvaluateContext) error {
	return nil
}
