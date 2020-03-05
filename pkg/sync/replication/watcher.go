package replication

import (
	"context"
	"fmt"
	"io"
	"time"

	"github.com/jackc/pgx"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	"github.com/Juniper/asf/pkg/db/basedb"
	"github.com/Juniper/asf/pkg/logutil"
)

// PostgresWatcher allows subscribing to PostgreSQL logical replication messages.
type PostgresWatcher struct {
	conf WatcherOptions

	lsnCounter lsnCounter

	conn postgresWatcherConnection

	consumer ChangeHandler
	decoder  *pgoutputDecoder

	log *logrus.Entry

	dumpDoneCh chan struct{}
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

// NewPostgresWatcher creates new watcher and initializes its connections.
func NewPostgresWatcher(consumer ChangeHandler, db DB, options ...WatcherOption) (*PostgresWatcher, error) {
	conf := defaultWatcherOptions()
	for _, o := range options {
		o(&conf)
	}
	conn, err := NewPostgresConnection(db)
	if err != nil {
		return nil, err
	}

	log := logutil.NewLogger("postgres-watcher")
	log.WithField("config", fmt.Sprintf("%+v", conf)).Debug("Got pgx config")

	return &PostgresWatcher{
		conn:       conn,
		conf:       conf,
		consumer:   consumer,
		decoder:    newPgoutputDecoder(),
		dumpDoneCh: make(chan struct{}),
		log:        log,
	}, nil
}

// WatcherOptions stores configuration for watcher object.
type WatcherOptions struct {
	Slot          string
	Publication   string
	StatusTimeout time.Duration
	NoDump        bool
}

func defaultWatcherOptions() WatcherOptions {
	return WatcherOptions{
		Slot:          SlotName("postgres-watcher"),
		Publication:   PostgreSQLPublicationName,
		StatusTimeout: viper.GetDuration("database.replication_status_timeout"),
	}
}

// WatcherOption is a function that sets some value in WatherOptions.
type WatcherOption func(*WatcherOptions)

// Slot sets replication slot name.
func Slot(s string) WatcherOption {
	return func(o *WatcherOptions) {
		o.Slot = SlotName(s)
	}
}

// Publication sets publication name.
func Publication(p string) WatcherOption {
	return func(o *WatcherOptions) {
		o.Publication = p
	}
}

// StatusTimeout sets duration between status calls.
func StatusTimeout(t time.Duration) WatcherOption {
	return func(o *WatcherOptions) {
		o.StatusTimeout = t
	}
}

// NoDump turns off the default DB dump that is done before starting the replication.
func NoDump() WatcherOption {
	return func(o *WatcherOptions) {
		o.NoDump = true
	}
}

// DumpDone returns a channel that is closed when dump is done.
func (w *PostgresWatcher) DumpDone() <-chan struct{} {
	return w.dumpDoneCh
}

// Start starts running the subscription. It can be stopped with context cancellation.
func (w *PostgresWatcher) Start(ctx context.Context) error {
	w.log.Debug("Starting Watch")

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
		w.conf.NoDump = true
		close(w.dumpDoneCh)
	}()
	if w.conf.NoDump {
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

			if err := w.handleMessage(ctx, msg); err != nil {
				w.log.Error("Error while handling replication message: ", err)
			}
		}
	}
}

func (w *PostgresWatcher) handleMessage(ctx context.Context, msg *pgx.ReplicationMessage) error {
	if msg == nil {
		return nil
	}

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
