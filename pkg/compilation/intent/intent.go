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

// Intent contains Intent Compiler state for a resource.
type Intent interface {
	basemodels.Object
	Evaluate(ctx context.Context, evaluateCtx *EvaluateContext) error
	GetObject() basemodels.Object
	GetDependencies() map[string]map[string]struct{}
}

// Loader provides access to existing intents (e.g. using a cache)
type Loader interface {
	Load(typeName string, id Query) Intent
}

// BaseIntent implements the default Evaluate interface
type BaseIntent struct {
	// Dependencies maps type of dependent intents to set of theirs uuids
	Dependencies map[string]map[string]struct{}
}

func (b *BaseIntent) GetDependencies() map[string]map[string]struct{} {
	return b.Dependencies
}

// Evaluate creates/updates/deletes lower-level resources when needed.
func (b *BaseIntent) Evaluate(ctx context.Context, evaluateCtx *EvaluateContext) error {
	return nil
}

func GetDependencies(
	i Intent,
	from string,
	loader Loader,
	reactionMap map[string]map[string]map[string]struct{},
) []Intent {
	intents := []Intent{}
	intentReactionMap, ok := reactionMap[i.Kind()]
	if !ok {
		return intents
	}
	dependentTypes, ok := intentReactionMap[from]
	if !ok {
		return intents
	}
	intents = append(intents, i)
	dependentIntents := i.GetDependencies()
	for t, uuids := range dependentIntents {
		_, ok := dependentTypes[t]
		if ok {
			for uuid := range uuids {
				dependentIntent := loader.Load(t, ByUUID(uuid))
				intents = append(intents, dependentIntent)
				for _, k := range GetDependencies(dependentIntent, t, loader, reactionMap) {
					intents = append(intents, k)
				}
			}
		}
	}
	return intents
}
