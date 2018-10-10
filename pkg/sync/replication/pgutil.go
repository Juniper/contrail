package replication

import (
	"fmt"

	"github.com/kyleconroy/pgoutput"
)

type relationSet map[uint32]pgoutput.Relation

func (rs relationSet) Add(r pgoutput.Relation) {
	rs[r.ID] = r
}

func (rs relationSet) Get(id uint32) (pgoutput.Relation, error) {
	rel, ok := rs[id]
	if !ok {
		return rel, fmt.Errorf("no relation for %d", id)
	}
	return rel, nil
}
