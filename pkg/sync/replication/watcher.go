package replication

import (
	"context"
	"io"
	"time"

	"github.com/jackc/pgx"
	"github.com/kyleconroy/pgoutput"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"

	"github.com/Juniper/contrail/pkg/db"
	pkglog "github.com/Juniper/contrail/pkg/log"
	"github.com/Juniper/contrail/pkg/services"
)

type abstractCanal interface {
	Run() error
	WaitDumpDone() <-chan struct{}
	Close()
}

// MySQLWatcher uses canal to read MySQL binlog events.
type MySQLWatcher struct {
	canal abstractCanal
	log   *logrus.Entry
}

// NewMySQLWatcher creates new BinlogWatcher listening on provided canal.
func NewMySQLWatcher(c abstractCanal) *MySQLWatcher {
	return &MySQLWatcher{
		canal: c,
		log:   pkglog.NewLogger("mysql-watcher"),
	}
}

// Watch starts listening on a canal.
func (w *MySQLWatcher) Watch(context.Context) error {
	w.log.Debug("Watching events on MySQL binlog")
	return w.canal.Run()
}

// DumpDone returns a channel that is closed when dump is done.
func (w *MySQLWatcher) DumpDone() <-chan struct{} {
	return w.canal.WaitDumpDone()
}

// Close closes canal.
func (w *MySQLWatcher) Close() {
	w.log.Debug("Stopping watching events on MySQL binlog")
	w.canal.Close()
}

// PostgresSubscriptionConfig stores configuration for logical replication connection used for Subscription object.
type PostgresSubscriptionConfig struct {
	Slot          string
	Publication   string
	StatusTimeout time.Duration
}

type postgresWatcherConnection interface {
	io.Closer
	GetReplicationSlot(name string) (lastLSN uint64, snapshotName string, err error)
	StartReplication(slot, publication string, startLSN uint64) error
	WaitForReplicationMessage(ctx context.Context) (*pgx.ReplicationMessage, error)
	SendStatus(lastLSN uint64) error

	DoInTransactionSnapshot(ctx context.Context, snapshotName string, do func(context.Context) error) error
}

// Handler handles pgoutput.Message with context.
type Handler func(context.Context, pgoutput.Message) error

// PostgresWatcher allows subscribing to PostgreSQL logical replication messages.
type PostgresWatcher struct {
	conf PostgresSubscriptionConfig

	lastLSN uint64 // lastLSN holds max WAL LSN value processed by subscription.

	cancel  context.CancelFunc
	conn    postgresWatcherConnection
	handler Handler

	db        services.Service
	processor services.EventProcessor

	log *logrus.Entry

	dumpDoneCh chan struct{}
	shouldDump bool
}

// NewPostgresWatcher creates new watcher and initializes its connections.
func NewPostgresWatcher(
	config PostgresSubscriptionConfig,
	dbs *db.Service, replConn pgxReplicationConn,
	handler Handler,
	processor services.EventProcessor,
	shouldDump bool,
) (*PostgresWatcher, error) {
	conn, err := newPostgresReplicationConnection(dbs, replConn)
	if err != nil {
		return nil, err
	}

	w := &PostgresWatcher{
		conf:       config,
		conn:       conn,
		handler:    handler,
		db:         dbs,
		processor:  processor,
		shouldDump: shouldDump,
		dumpDoneCh: make(chan struct{}),
		log:        pkglog.NewLogger("postgres-watcher"),
	}
	if !w.shouldDump {
		close(w.dumpDoneCh)
	}
	return w, nil
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

	var snapshotName string
	slotLSN, snapshotName, err := w.conn.GetReplicationSlot(w.conf.Slot)
	if err != nil {
		return errors.Wrap(err, "error getting replication slot")
	}
	w.log.Debug("consistentPoint: ", pgx.FormatLSN(slotLSN))
	w.log.Debug("snapshotName: ", snapshotName)

	w.lastLSN = slotLSN // TODO(Michal): get lastLSN from etcd

	if w.shouldDump {
		if err := w.Dump(ctx, snapshotName); err != nil {
			return nil
		}
	}

	if err := w.conn.StartReplication(w.conf.Slot, w.conf.Publication, 0); err != nil {
		return errors.Wrap(err, "failed to start replication")
	}

	feed := make(chan *pgx.ReplicationMessage)
	go w.runMessageConsumer(ctx, feed)
	return w.runMessageProducer(ctx, feed)
}

// Dump dumps whole db state using provided snapshot name.
func (w *PostgresWatcher) Dump(ctx context.Context, snapshotName string) error {
	w.log.Debug("Starting dump phase")
	defer func() {
		w.shouldDump = false
		close(w.dumpDoneCh)
	}()
	dumpStart := time.Now()

	if err := w.conn.DoInTransactionSnapshot(ctx, snapshotName, func(ctx context.Context) error {
		es, err := services.Dump(ctx, w.db)
		if err != nil {
			return err
		}
		for _, e := range es.Events {
			_, err = w.processor.Process(ctx, e)
			if err != nil {
				return errors.Wrapf(err, "error processing event: %v", e)
			}
		}

		return nil
	}); err != nil {
		return errors.Wrap(w.muteCancellationError(err), "dumping snapshot failed")
	}
	w.log.WithField("dumpTime", time.Since(dumpStart)).Debugf("Dump phase finished - starting replication")

	return nil
}

func (w *PostgresWatcher) muteCancellationError(err error) error {
	if isContextCancellationError(err) {
		w.log.Infof("Watcher exited with cancellation error: %v", err)
		return nil
	}
	return err
}

func (w *PostgresWatcher) runMessageConsumer(ctx context.Context, feed <-chan *pgx.ReplicationMessage) {
	func() {
		for msg := range feed {
			if err := w.handleMessage(ctx, msg); err != nil {
				w.log.Error("Error while handling replication message: ", err)
			}
		}
	}()
}

func (w *PostgresWatcher) runMessageProducer(ctx context.Context, feed chan<- *pgx.ReplicationMessage) error {
	tick := time.NewTicker(w.conf.StatusTimeout).C
	for {
		select {
		case <-ctx.Done():
			return w.muteCancellationError(w.conn.Close())
		case <-tick:
			w.log.Debug("Sending standby status with position: ", pgx.FormatLSN(w.lastLSN))
			if err := w.conn.SendStatus(w.lastLSN); err != nil {
				return err
			}
		default:
			msg, err := w.waitForMessageWithTimeout(ctx)
			if err != nil {
				return err
			}

			if msg != nil {
				feed <- msg
			}
		}
	}
}

func (w *PostgresWatcher) waitForMessageWithTimeout(ctx context.Context) (*pgx.ReplicationMessage, error) {
	wctx, cancel := context.WithTimeout(ctx, w.conf.StatusTimeout)
	defer cancel()

	msg, err := w.conn.WaitForReplicationMessage(wctx)

	if err == context.DeadlineExceeded || err == context.Canceled {
		return nil, nil
	}
	if err != nil {
		return nil, errors.Wrap(err, "replication failed")
	}

	return msg, nil
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
	if msg.WalStart > w.lastLSN {
		w.lastLSN = msg.WalStart
	}

	logmsg, err := pgoutput.Parse(msg.WalData)
	if err != nil {
		return errors.Wrap(err, "invalid pgoutput message")
	}

	if err := w.handler(ctx, logmsg); err != nil {
		return errors.Wrap(err, "error handling waldata")
	}

	return nil
}

func (w *PostgresWatcher) handleServerHeartbeat(shb *pgx.ServerHeartbeat) error {
	if shb.ReplyRequested == 1 {
		w.log.Info("Server requested reply, sending standby status with position: ", pgx.FormatLSN)
		return w.conn.SendStatus(w.lastLSN)
	}
	return nil
}

// Close stops subscription by calling cancel function of context passed in Watch.
func (w *PostgresWatcher) Close() {
	w.log.Debug("Stopping watching events on PostgreSQL replication slot")
	w.cancel()
}
