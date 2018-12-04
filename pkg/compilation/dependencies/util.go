package dependencies

import (
	"fmt"

	yaml "gopkg.in/yaml.v2"
)

// ParseReactions parses reactions from file.
func ParseReactions(bytes []byte, prefix string) (Reactions, error) {
	raw := map[string]map[string]map[string][]string{}
	err := yaml.Unmarshal(bytes, raw)
	if err != nil {
		return nil, err
	}

	rawReactions, ok := raw[prefix]
	if !ok {
		return nil, fmt.Errorf(
			"failed to parse reactions. specified key (%s) not found in input file",
			prefix,
		)
	}
	reactions := Reactions{}
	for kind, reactionMap := range rawReactions {
		reactions[kind] = map[string]KindSet{}
		for fromKind, l := range reactionMap {
			reactions[kind][fromKind] = KindSet{}
			for _, reaction := range l {
				reactions[kind][fromKind][reaction] = struct{}{}
			}
		}
	}
	return reactions, nil
}
