package dependencies

import (
	"errors"
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

func ParseReactions(path, prefix string) (Reactions, error) {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	raw := map[string]map[string]map[string][]string{}
	err = yaml.Unmarshal(bytes, raw)
	if err != nil {
		return nil, err
	}

	rawReactions, ok := raw[prefix]
	if !ok {
		return nil, errors.New(
			fmt.Sprintf(
				"failed to parse reactions. specified key (%s) not found in input file (%s)",
				prefix,
				path,
			))
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
