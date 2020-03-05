package replication

import (
	"context"
	"database/sql"
	"fmt"
	"io"
	"strings"

	"github.com/Juniper/asf/pkg/db/basedb"
	"github.com/Juniper/asf/pkg/logutil"
	"github.com/jackc/pgx"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

type pgxReplicationConn interface {
	io.Closer

	DropReplicationSlot(slotName string) (err error)
	CreateReplicationSlotEx(slotName, outputPlugin string) (consistentPoint string, snapshotName string, err error)
	SendStandbyStatus(k *pgx.StandbyStatus) (err error)
	StartReplication(slotName string, startLsn uint64, timeline int64, pluginArguments ...string) (err error)
	WaitForReplicationMessage(ctx context.Context) (*pgx.ReplicationMessage, error)
}

// DB is a DB access interface.
type DB interface {
	DB() *sql.DB
	DoInTransactionWithOpts(ctx context.Context, do func(context.Context) error, opts *sql.TxOptions) error
	Dump(context.Context) (basedb.DatabaseData, error)
}

type postgresConnection struct {
	replConn pgxReplicationConn
	db       DB
	log      *logrus.Entry
}

func NewPostgresConnection(db DB) (*postgresConnection, error) {
	c := basedb.ConnectionConfigFromViper()
	replConn, err := pgx.ReplicationConnect(
		pgx.ConnConfig{Host: c.Host, Database: c.Name, User: c.User, Password: c.Password},
	)
	if err != nil {
		return nil, err
	}

	return &postgresConnection{
		db:       db,
		replConn: replConn,
		log:      logutil.NewLogger("postgres-replication-connection"),
	}, nil
}

// GetReplicationSlot gets replication slot for replication.
func (c *postgresConnection) GetReplicationSlot(name string) (maxWal LSN, snapshotName string, err error) {
	if dropErr := c.replConn.DropReplicationSlot(name); err != nil {
		c.log.WithError(dropErr).Info("Could not drop replication slot just before getting new one - safely ignoring")
	}

	// If creating the replication slot fails with code 42710, this means
	// the replication slot already exists.
	consistentPoint, snapshotName, err := c.replConn.CreateReplicationSlotEx(name, "pgoutput")
	if err != nil {
		if pgerr, ok := err.(pgx.PgError); !ok || pgerr.Code != "42710" {
			return 0, "", errors.Wrap(err, "failed to create replication slot")
		}
	}

	maxWal, err = ParseLSN(consistentPoint)
	if err != nil {
		return 0, "", fmt.Errorf("error parsing received LSN: %v", err)
	}
	return maxWal, snapshotName, err
}

// RenewPublication ensures that publication exists for all tables.
func (c *postgresConnection) RenewPublication(ctx context.Context, name string) error {
	return c.db.DoInTransactionWithOpts(
		ctx,
		func(ctx context.Context) error {
			_, err := c.db.DB().ExecContext(ctx, fmt.Sprintf("DROP PUBLICATION IF EXISTS %s", name))
			if err != nil {
				return errors.Wrap(err, "failed to drop publication")
			}
			_, err = c.db.DB().ExecContext(ctx, fmt.Sprintf("CREATE PUBLICATION %s FOR ALL TABLES", name))
			if err != nil {
				return errors.Wrap(err, "failed to create publication")
			}
			return err
		},
		nil,
	)
}

// IsInRecovery checks is database server is in recovery mode.
func (c *postgresConnection) IsInRecovery(ctx context.Context) (isInRecovery bool, err error) {
	return isInRecovery, c.db.DoInTransactionWithOpts(
		ctx,
		func(ctx context.Context) error {
			r, err := c.db.DB().QueryContext(ctx, "SELECT pg_is_in_recovery()")
			if err != nil {
				return errors.Wrap(err, "failed to check recovery mode")
			}
			if !r.Next() {
				return errors.New("pg_is_in_recovery() returned zero rows")
			}
			if err := r.Scan(&isInRecovery); err != nil {
				return errors.Wrap(err, "error scanning recovery status")
			}
			return nil
		},
		&sql.TxOptions{ReadOnly: true},
	)
}

func (c *postgresConnection) DoInTransactionSnapshot(
	ctx context.Context,
	snapshotName string,
	do func(context.Context) error,
) error {
	return c.db.DoInTransactionWithOpts(
		ctx,
		func(ctx context.Context) error {
			tx := basedb.GetTransaction(ctx)
			_, err := tx.ExecContext(ctx, "SET TRANSACTION ISOLATION LEVEL REPEATABLE READ")
			if err != nil {
				return errors.Wrap(err, "error setting transaction isolation")
			}
			_, err = tx.ExecContext(ctx, fmt.Sprintf("SET TRANSACTION SNAPSHOT '%s'", snapshotName))
			if err != nil {
				return errors.Wrap(err, "error setting transaction snapshot")
			}

			return do(ctx)
		},
		&sql.TxOptions{ReadOnly: true},
	)
}

func (c *postgresConnection) DumpSnapshot(
	ctx context.Context, snapshotName string,
) (dump basedb.DatabaseData, err error) {
	err = c.DoInTransactionSnapshot(ctx, snapshotName, func(ctx context.Context) error {
		dump, err = c.db.Dump(ctx)
		if err != nil {
			return err
		}
		return nil
	})

	return dump, nil
}

func pluginArgs(publication string) string {
	return fmt.Sprintf(`("proto_version" '1', "publication_names" '%s')`, publication)
}

// StartReplication sends start replication message to server.
func (c *postgresConnection) StartReplication(slot, publication string, start LSN) error {
	// timeline argument should be -1 otherwise postgres reutrns error - pgx library bug
	return c.replConn.StartReplication(slot, uint64(start), -1, pluginArgs(publication))
}

// WaitForReplicationMessage blocks until message arrives on replication connection.
func (c *postgresConnection) WaitForReplicationMessage(ctx context.Context) (*pgx.ReplicationMessage, error) {
	return c.replConn.WaitForReplicationMessage(ctx)
}

// SendStatus sends standby status to server connected with replication connection.
func (c *postgresConnection) SendStatus(received, saved LSN) error {
	c.log.WithFields(logrus.Fields{
		"receivedLSN": received,
		"savedLSN":    saved,
	}).Info("Sending standby status")
	k, err := pgx.NewStandbyStatus(
		uint64(saved),    // flush - savedLSN is already stored in etcd so we can say that it's flushed
		uint64(saved),    // apply - savedLSN is stored and visible in etcd so it's also applied
		uint64(received), // write - receivedLSN is last wal segment that was received by watcher
	)
	if err != nil {
		return errors.Wrap(err, "error creating standby status")
	}
	if err = c.replConn.SendStandbyStatus(k); err != nil {
		return errors.Wrap(err, "failed to send standy status")
	}
	return nil
}

// Close closes underlying connections.
func (c *postgresConnection) Close() error {
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
