package intent

import (
	"context"

	"github.com/Juniper/contrail/pkg/models/basemodels"
	"github.com/Juniper/contrail/pkg/services"
)

// EvaluateContext contains context information for Resource to handle CRUD
type EvaluateContext struct {
	WriteService     services.WriteService
	ReadService      services.ReadService
	IntPoolAllocator services.IntPoolAllocator
	IntentLoader     Loader
}

type uuidSet map[string]struct{}

// Intent contains Intent Compiler state for a resource.
type Intent interface {
	basemodels.Object
	Evaluate(ctx context.Context, evaluateCtx *EvaluateContext) error
	GetObject() basemodels.Object
	GetDependencies() map[string]uuidSet
	AddDependentIntent(i Intent)
	RemoveDependentIntent(i Intent)
}

// Loader provides access to existing intents (e.g. using a cache)
type Loader interface {
	Load(typeName string, id Query) Intent
}

// BaseIntent implements the default Evaluate interface
type BaseIntent struct {
	// Dependencies maps type of dependent intents to set of theirs uuids
	Dependencies map[string]uuidSet
}

// GetDependencies returns collection of references and backrefs' uuids
func (b *BaseIntent) GetDependencies() map[string]uuidSet {
	return b.Dependencies
}

// Evaluate creates/updates/deletes lower-level resources when needed.
func (b *BaseIntent) Evaluate(ctx context.Context, evaluateCtx *EvaluateContext) error {
	return nil
}

// AddDependentIntent adds a intent to dependencies.
// This method can be used to add custom intents which are not refs or backrefs
func (b *BaseIntent) AddDependentIntent(i Intent) {
	if b.Dependencies == nil {
		b.Dependencies = map[string]uuidSet{}
	}
	kindMap := b.Dependencies[i.Kind()]
	if kindMap == nil {
		kindMap = map[string]struct{}{}
		b.Dependencies[i.Kind()] = kindMap
	}
	kindMap[i.GetUUID()] = struct{}{}
}

// RemoveDependentIntent removes a intent from dependencies.
// This method can be used to add custom logic and reactions.
func (b *BaseIntent) RemoveDependentIntent(i Intent) {
	if b.Dependencies == nil {
		return
	}
	kindMap := b.Dependencies[i.Kind()]
	if kindMap != nil {
		delete(kindMap, i.GetUUID())
	}
}
