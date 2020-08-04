package sync

import (
	"context"
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/Juniper/asf/pkg/db"
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

func newReplConn() (*pgx.ReplicationConn, error) {
	c := db.ConnectionConfigFromViper()
	databasePort, err := strconv.ParseUint(c.Port, 0, 16)
	if err != nil {
		return nil, errors.Wrap(err, "parse database prot")
	}
	replConn, err := pgx.ReplicationConnect(
		pgx.ConnConfig{
			Host:     c.Host,
			Database: c.Name,
			User:     c.User,
			Password: c.Password,
			Port:     uint16(databasePort),
		},
	)
	if err != nil {
		return nil, errors.Wrap(err, "create pgx replication connection")
	}

	return replConn, nil
}

// PostgresConnection is a connection that uses replication connection to manage logical subscription.
type PostgresConnection struct {
	replConn pgxReplicationConn
	log      *logrus.Entry
}

// NewPostgresConnection returns a postgres Connection.
func NewPostgresConnection() (*PostgresConnection, error) {
	replConn, err := newReplConn()
	if err != nil {
		errors.Wrap(err, "create replConn")
	}

	return &PostgresConnection{
		replConn: replConn,
		log:      logutil.NewLogger("postgres-replication-connection"),
	}, nil
}

// Reconnect replace old replConn with a new one, this function is only desired to be used when old connection is dead
func (c *PostgresConnection) Reconnect() error {
	replConn, err := newReplConn()
	if err != nil {
		return errors.Wrap(err, "create replConn for reconnection")
	}

	c.replConn = replConn

	return nil
}

// GetReplicationSlot gets replication slot for replication.
func (c *PostgresConnection) GetReplicationSlot(name string) (maxWal LSN, snapshotName string, err error) {
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

func pluginArgs(publication string) string {
	return fmt.Sprintf(`("proto_version" '1', "publication_names" '%s')`, publication)
}

// StartReplication sends start replication message to server.
func (c *PostgresConnection) StartReplication(slot, publication string, start LSN) error {
	// timeline argument should be -1 otherwise postgres reutrns error - pgx library bug
	return c.replConn.StartReplication(slot, uint64(start), -1, pluginArgs(publication))
}

// WaitForReplicationMessage blocks until message arrives on replication connection.
func (c *PostgresConnection) WaitForReplicationMessage(ctx context.Context) (*pgx.ReplicationMessage, error) {
	return c.replConn.WaitForReplicationMessage(ctx)
}

// SendStatus sends standby status to server connected with replication connection.
func (c *PostgresConnection) SendStatus(received, saved LSN) error {
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
func (c *PostgresConnection) Close() error {
	var errs []string
	if replConnErr := c.replConn.Close(); replConnErr != nil {
		errs = append(errs, replConnErr.Error())
	}
	if len(errs) > 0 {
		return fmt.Errorf("errors while closing: %s", strings.Join(errs, "\n"))
	}
	return nil
}
