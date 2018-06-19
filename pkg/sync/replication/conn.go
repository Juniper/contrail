package replication

import (
	"context"
	"database/sql"
	"fmt"
	"io"
	"strings"

	"github.com/Juniper/contrail/pkg/db"
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

type dbService interface {
	DB() *sql.DB
	DoInTransaction(ctx context.Context, do func(context.Context) error) error
	Dump(context.Context, db.ObjectWriter) error
}

type postgresReplicationConnection struct {
	replConn pgxReplicationConn
	db       dbService
}

func newPostgresReplicationConnection(
	db dbService, replConn pgxReplicationConn,
) (*postgresReplicationConnection, error) {
	return &postgresReplicationConnection{db: db, replConn: replConn}, nil
}

// GetReplicationSlot gets replication slot for replication.
func (c *postgresReplicationConnection) GetReplicationSlot(
	name string,
) (maxWal uint64, snapshotName string, err error) {
	_ = c.replConn.DropReplicationSlot(name) // nolint: gas

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
func (c *postgresReplicationConnection) RenewPublication(ctx context.Context, name string) error {
	return c.db.DoInTransaction(
		ctx,
		func(ctx context.Context) error {
			_, err := c.db.DB().ExecContext(ctx, fmt.Sprintf("DROP PUBLICATION IF EXISTS %s", name))
			if err != nil {
				return fmt.Errorf("failed to drop publication: %s", err)
			}
			_, err = c.db.DB().ExecContext(ctx, fmt.Sprintf("CREATE PUBLICATION %s FOR ALL TABLES", name))
			if err != nil {
				return fmt.Errorf("failed to create publication: %s", err)
			}
			return err
		},
	)
}

func (c *postgresReplicationConnection) DumpSnapshot(
	ctx context.Context, ow db.ObjectWriter, snapshotName string,
) error {
	return c.db.DoInTransaction(
		ctx,
		func(ctx context.Context) error {
			_, err := c.db.DB().ExecContext(ctx, "SET TRANSACTION ISOLATION LEVEL REPEATABLE READ")
			if err != nil {
				return err
			}

			return c.db.Dump(ctx, ow)
		},
	)
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
	var errs []string
	if dbErr := c.db.DB().Close(); dbErr != nil {
		errs = append(errs, dbErr.Error())
	}
	if replConnErr := c.replConn.Close(); replConnErr != nil {
		errs = append(errs, replConnErr.Error())
	}
	if len(errs) > 0 {
		return fmt.Errorf("errors while closing: %s", strings.Join(errs, "\n"))
	}
	return nil
}
