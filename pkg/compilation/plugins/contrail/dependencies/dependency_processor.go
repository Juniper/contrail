package dependencies

import "github.com/Juniper/contrail/pkg/compilation/intent"

// KindSet Set of strings representing resource Kind
type KindSet map[string]struct{}

// ReactionMap describes reactions between intents on create/update/delete event.
type ReactionMap map[string]map[string]KindSet

// DependencyProcessor resolves relations between intents based on ReactionMap
type DependencyProcessor struct {
	reactionMap map[string]map[string]KindSet
}

// NewDependencyProcessor creates new DependencyProcessor
func NewDependencyProcessor(
	reactionMap map[string]map[string]KindSet,
) *DependencyProcessor {
	return &DependencyProcessor{
		reactionMap: reactionMap,
	}
}

// GetDependencies resolves dependent intents and returns them as a map uuid -> intent
func (d *DependencyProcessor) GetDependencies(
	loader intent.Loader,
	i intent.Intent,
	from string,
) map[string]intent.Intent {
	intents := map[string]intent.Intent{}
	intentReactionMap, ok := d.reactionMap[i.Kind()]
	if !ok {
		return intents
	}
	dependentTypes, ok := intentReactionMap[from]
	if !ok {
		return intents
	}
	intents[i.GetUUID()] = i
	for t, uuids := range i.GetDependencies() {
		_, ok := dependentTypes[t]
		if ok {
			for uuid := range uuids {
				dependentIntent := loader.Load(t, intent.ByUUID(uuid))
				intents[dependentIntent.GetUUID()] = dependentIntent
				for _, k := range d.GetDependencies(loader, dependentIntent, t) {
					intents[k.GetUUID()] = k
				}
			}
		}
	}
	return intents
}
