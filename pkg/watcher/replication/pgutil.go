package replication

import (
	"fmt"

	"github.com/jackc/pgx"
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

type execer interface {
	Exec(string, ...interface{}) (pgx.CommandTag, error)
}

func createPublication(e execer, name, args string) error {
	_, err := e.Exec(fmt.Sprintf("CREATE PUBLICATION %s %s", name, args))
	return err
}

func createPublicationForAll(e execer, name string) error {
	return createPublication(e, name, "FOR ALL TABLES")
}

func dropPublication(e execer, name string) error {
	_, err := e.Exec(fmt.Sprintf("DROP PUBLICATION %s", name))
	return err
}

func setTransactionSnapshot(e execer, snapshotName string) error {
	_, err := e.Exec(fmt.Sprintf("SET TRANSACTION SNAPSHOT '%s'", snapshotName))
	return err
}
