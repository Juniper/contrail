package replication

import (
	"context"
	"fmt"
	"io"
	"strings"

	"github.com/jackc/pgx"
)

type pgxReplicationConn interface {
	io.Closer

	DropReplicationSlot(slotName string) (err error)
	CreateReplicationSlotEx(slotName, outputPlugin string) (consistentPoint string, snapshotName string, err error)
	SendStandbyStatus(k *pgx.StandbyStatus) (err error)
	StartReplication(slotName string, startLsn uint64, timeline int64, pluginArguments ...string) (err error)
	WaitForReplicationMessage(ctx context.Context) (*pgx.ReplicationMessage, error)
}

type pgxConn interface {
	io.Closer

	Exec(sql string, args ...interface{}) (pgx.CommandTag, error)
	BeginEx(ctx context.Context, txOptions *pgx.TxOptions) (*pgx.Tx, error)
}

type postgresReplicationConnection struct {
	replConn pgxReplicationConn
	conn     pgxConn
}

func newPostgresReplicationConnection(conf pgx.ConnConfig) (*postgresReplicationConnection, error) {
	replConn, err := pgx.ReplicationConnect(conf)
	if err != nil {
		return nil, err
	}

	conn, err := pgx.Connect(conf)
	if err != nil {
		return nil, err
	}

	return &postgresReplicationConnection{replConn: replConn, conn: conn}, nil
}

// GetReplicationSlot gets replication slot for replication.
func (c *postgresReplicationConnection) GetReplicationSlot(
	name string,
) (maxWal uint64, snapshotName string, err error) {
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

// RenewPublication ensures that publication exists for all tables.
func (c *postgresReplicationConnection) RenewPublication(name string) error {
	_, _ = c.conn.Exec(fmt.Sprintf("DROP PUBLICATION %s", name))
	_, err := c.conn.Exec(fmt.Sprintf("CREATE PUBLICATION %s FOR ALL TABLES", name))
	if err != nil {
		return fmt.Errorf("failed to create publication: %s", err)
	}
	return nil
}

func pluginArgs(publication string) string {
	return fmt.Sprintf(`("proto_version" '1', "publication_names" '%s')`, publication)
}

// StartReplication sends start replication message to server.
func (c *postgresReplicationConnection) StartReplication(slot, publication string, startLSN uint64) error {
	// timeline argument should be -1 otherwise postgres reutrns error - pgx library bug
	return c.replConn.StartReplication(slot, startLSN, -1, pluginArgs(publication))
}

// WaitForReplicationMessage blocks until message arrives on replication connection.
func (c *postgresReplicationConnection) WaitForReplicationMessage(
	ctx context.Context,
) (*pgx.ReplicationMessage, error) {
	return c.replConn.WaitForReplicationMessage(ctx)
}

// SendStatus sends standby status to server connected with replication connection.
func (c *postgresReplicationConnection) SendStatus(maxWal uint64) error {
	k, err := pgx.NewStandbyStatus(maxWal)
	if err != nil {
		return fmt.Errorf("error creating standby status: %s", err)
	}
	if err = c.replConn.SendStandbyStatus(k); err != nil {
		return fmt.Errorf("failed to send standy status: %s", err)
	}
	return nil
}

// Close closes underlying connections.
func (c *postgresReplicationConnection) Close() error {
	errs := []string{}
	errs = append(errs, c.conn.Close().Error())
	errs = append(errs, c.replConn.Close().Error())
	if len(errs) > 0 {
		return fmt.Errorf("errors while closing: %s", strings.Join(errs, "\n"))
	}
	return nil
}
