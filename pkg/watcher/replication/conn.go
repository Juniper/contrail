package replication

import (
	"context"
	"fmt"
	"io"
	"strings"

	"github.com/jackc/pgx"
)

type replSlotManager interface {
	DropReplicationSlot(slotName string) (err error)
	CreateReplicationSlotEx(slotName, outputPlugin string) (consistentPoint string, snapshotName string, err error)
}

type statusSender interface {
	SendStandbyStatus(k *pgx.StandbyStatus) (err error)
}

type replConn interface {
	replSlotManager
	statusSender
	io.Closer

	StartReplication(slotName string, startLsn uint64, timeline int64, pluginArguments ...string) (err error)
	WaitForReplicationMessage(ctx context.Context) (*pgx.ReplicationMessage, error)
}

type pgConn interface {
	execer
	io.Closer

	BeginEx(ctx context.Context, txOptions *pgx.TxOptions) (*pgx.Tx, error)
}

type Conn struct {
	replConn replConn
	conn     pgConn
}

func NewConn(conf pgx.ConnConfig) (*Conn, error) {
	replConn, err := pgx.ReplicationConnect(conf)
	if err != nil {
		return nil, err
	}

	conn, err := pgx.Connect(conf)
	if err != nil {
		return nil, err
	}

	return &Conn{replConn: replConn, conn: conn}, nil
}

func (c *Conn) GetReplicationSlot(name string) (maxWal uint64, snapshotName string, err error) {
	_ = c.replConn.DropReplicationSlot(name)

	// If creating the replication slot fails with code 42710, this means
	// the replication slot already exists.
	consistentPoint, snapshotName, err := c.replConn.CreateReplicationSlotEx(name, "pgoutput")
	if err != nil {
		if pgerr, ok := err.(pgx.PgError); !ok || pgerr.Code != "42710" {
			return 0, "", fmt.Errorf("failed to create replication slot: %s", err)
		}
	}

	maxWal, err = pgx.ParseLSN(consistentPoint)
	if err != nil {
		return 0, "", fmt.Errorf("error parsing received LSN: %v", err)
	}
	return maxWal, snapshotName, err
}

func (c *Conn) RenewPublication(name string) error {
	_ = dropPublication(c.conn, name)
	err := createPublicationForAll(c.conn, name)
	if err != nil {
		return fmt.Errorf("failed to create publication: %s", err)
	}
	return nil
}

func pluginArgs(version, publication string) string {
	return fmt.Sprintf(`("proto_version" '%s', "publication_names" '%s')`, version, publication)
}

func (c *Conn) StartReplication(slot, publication string, startLSN uint64) error {
	// timeline argument should be -1 otherwise postgres reutrns error - pgx library bug
	return c.replConn.StartReplication(slot, startLSN, -1, pluginArgs("1", publication))
}

func (c *Conn) WaitForReplicationMessage(ctx context.Context) (*pgx.ReplicationMessage, error) {
	return c.replConn.WaitForReplicationMessage(ctx)
}

func (c *Conn) SendStatus(maxWal uint64) error {
	k, err := pgx.NewStandbyStatus(maxWal)
	if err != nil {
		return fmt.Errorf("error creating standby status: %s", err)
	}
	if err = c.replConn.SendStandbyStatus(k); err != nil {
		return fmt.Errorf("failed to send standy status: %s", err)
	}
	return nil
}

func (c *Conn) Close() error {
	errs := []string{}
	errs = append(errs, c.conn.Close().Error())
	errs = append(errs, c.replConn.Close().Error())
	if len(errs) > 0 {
		return fmt.Errorf("errors while closing: %s", strings.Join(errs, "\n"))
	}
	return nil
}
