package replication

import (
	"context"
	"fmt"
	"io"
	"time"

	"github.com/jackc/pgx"
	"github.com/kyleconroy/pgoutput"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"

	"github.com/Juniper/asf/pkg/db/basedb"
	"github.com/Juniper/asf/pkg/logutil"
)

// PostgresSubscriptionConfig stores configuration for logical replication connection used for Subscription object.
type PostgresSubscriptionConfig struct {
	Slot          string
	Publication   string
	StatusTimeout time.Duration
}

type postgresWatcherConnection interface {
	io.Closer
	GetReplicationSlot(name string) (receivedLSN LSN, snapshotName string, err error)
	StartReplication(slot, publication string, start LSN) error
	WaitForReplicationMessage(ctx context.Context) (*pgx.ReplicationMessage, error)
	SendStatus(received, saved LSN) error
	IsInRecovery(context.Context) (bool, error)

	DumpSnapshot(ctx context.Context, snapshotName string) (basedb.DatabaseData, error)
}

// Handler handles pgoutput.Message with context.
type Handler func(context.Context, pgoutput.Message) error

// PostgresWatcher allows subscribing to PostgreSQL logical replication messages.
type PostgresWatcher struct {
	conf PostgresSubscriptionConfig

	lsnCounter lsnCounter

	cancel context.CancelFunc
	conn   postgresWatcherConnection

	consumer ChangeHandler
	decoder  *pgoutputDecoder

	log *logrus.Entry

	dumpDoneCh chan struct{}
	shouldDump bool
}

// NewPostgresWatcher creates new watcher and initializes its connections.
func NewPostgresWatcher(
	config PostgresSubscriptionConfig,
	conn postgresWatcherConnection,
	consumer ChangeHandler,
	shouldDump bool,
) *PostgresWatcher {
	log := logutil.NewLogger("postgres-watcher")
	log.WithField("config", fmt.Sprintf("%+v", config)).Debug("Got pgx config")

	return &PostgresWatcher{
		conf:       config,
		conn:       conn,
		consumer:   consumer,
		decoder:    newPgoutputDecoder(),
		shouldDump: shouldDump,
		dumpDoneCh: make(chan struct{}),
		log:        log,
	}
}

// DumpDone returns a channel that is closed when dump is done.
func (w *PostgresWatcher) DumpDone() <-chan struct{} {
	return w.dumpDoneCh
}

// Watch starts subscription and store context cancel function.
func (w *PostgresWatcher) Watch(ctx context.Context) error {
	w.log.Debug("Starting Watch")
	ctx, cancel := context.WithCancel(ctx)
	w.cancel = cancel

	// Logical replication cannot be run in recovery mode.
	isInRecovery, err := w.conn.IsInRecovery(ctx)
	if isInRecovery {
		return markTemporaryError(errors.New("database is in recovery mode"))
	}
	if err != nil {
		return wrapError(err)
	}

	var snapshotName string
	slotLSN, snapshotName, err := w.conn.GetReplicationSlot(w.conf.Slot)
	if err != nil {
		return wrapError(errors.Wrap(err, "error getting replication slot"))
	}
	w.log.WithFields(logrus.Fields{
		"consistentPoint": slotLSN,
		"snapshotName: ":  snapshotName,
	}).Debug("Got replication slot")

	w.lsnCounter.updateReceivedLSN(slotLSN) // TODO(Michal): get receivedLSN from etcd

	if err := w.dumpIfShould(ctx, snapshotName); err != nil {
		return wrapError(err)
	}

	w.lsnCounter.txnFinished(slotLSN)

	if err := w.conn.StartReplication(w.conf.Slot, w.conf.Publication, 0); err != nil {
		return wrapError(errors.Wrap(err, "failed to start replication"))
	}

	return wrapError(w.runMessageProducer(ctx))
}

func (w *PostgresWatcher) dumpIfShould(ctx context.Context, snapshotName string) error {
	defer func() {
		w.shouldDump = false
		close(w.dumpDoneCh)
	}()
	if !w.shouldDump {
		return nil
	}

	return w.Dump(ctx, snapshotName)
}

// Dump dumps whole db state using provided snapshot name.
func (w *PostgresWatcher) Dump(ctx context.Context, snapshotName string) error {
	w.log.Debug("Starting dump phase")
	dumpStart := time.Now()

	dumpData, err := w.conn.DumpSnapshot(ctx, snapshotName)
	if err != nil {
		return errors.Wrap(w.muteCancellationError(err), "dumping snapshot failed")
	}

	if err := dumpData.ForEachRow(func(schemaID string, row basedb.RowData) error {
		return w.consumer.Handle(
			ctx, []Change{change{kind: schemaID, data: row, pk: row.PK(), operation: CreateOperation}},
		)
	}); err != nil {
		return errors.Wrap(w.muteCancellationError(err), "storing dump data failed")
	}

	w.log.WithField("dumpTime", time.Since(dumpStart)).Debug("Dump phase finished - starting replication")

	return nil
}

func (w *PostgresWatcher) muteCancellationError(err error) error {
	if isContextCancellationError(err) {
		w.log.Infof("Watcher exited with cancellation error: %v", err)
		return nil
	}
	return err
}

func (w *PostgresWatcher) runMessageProducer(ctx context.Context) error {
	ticker := time.NewTicker(w.conf.StatusTimeout)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			w.log.Debug("Stopping watching events on PostgreSQL replication slot")
			return w.muteCancellationError(w.conn.Close())
		case <-ticker.C:
			if err := w.sendStatus(); err != nil {
				return err
			}
		default:
			wctx, cancel := context.WithTimeout(ctx, w.conf.StatusTimeout)
			defer cancel()

			msg, err := w.conn.WaitForReplicationMessage(wctx)
			if err != nil && err != context.DeadlineExceeded && err != context.Canceled {
				return errors.Wrap(err, "replication failed")
			}

			if msg != nil {
				if err := w.handleMessage(ctx, msg); err != nil {
					w.log.Error("Error while handling replication message: ", err)
				}
			}
		}
	}
}

func (w *PostgresWatcher) handleMessage(ctx context.Context, msg *pgx.ReplicationMessage) error {
	if msg.WalMessage != nil {
		if err := w.handleWalMessage(ctx, msg.WalMessage); err != nil {
			return err
		}
	}

	if msg.ServerHeartbeat != nil {
		if err := w.handleServerHeartbeat(msg.ServerHeartbeat); err != nil {
			return err
		}
	}

	return nil
}

func (w *PostgresWatcher) handleWalMessage(ctx context.Context, msg *pgx.WalMessage) error {
	w.lsnCounter.txnStarted()
	defer w.lsnCounter.txnFinished(LSN(msg.WalStart))

	w.lsnCounter.updateReceivedLSN(MessageLSN(msg))

	changes, err := w.decoder.DecodeChanges(msg.WalData)
	if err != nil {
		return errors.Wrap(err, "error decoding pgoutput changes")
	}

	if len(changes) == 0 {
		return nil
	}
	return w.consumer.Handle(ctx, changes)
}

func (w *PostgresWatcher) handleServerHeartbeat(shb *pgx.ServerHeartbeat) error {
	w.lsnCounter.updateReceivedLSN(LSN(shb.ServerWalEnd))
	if shb.ReplyRequested == 1 {
		return w.sendStatus()
	}
	return nil
}

func (w *PostgresWatcher) sendStatus() error {
	r, s := w.lsnCounter.lsnValues()
	return w.conn.SendStatus(r, s)
}

// Close stops subscription by calling cancel function of context passed in Watch.
func (w *PostgresWatcher) Close() {
	w.cancel()
}
