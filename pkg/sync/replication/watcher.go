package replication

import (
	"context"
	"io"
	"time"

	"github.com/Juniper/contrail/pkg/db"
	pkglog "github.com/Juniper/contrail/pkg/log"
	"github.com/jackc/pgx"
	"github.com/kyleconroy/pgoutput"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

type abstractCanal interface {
	Run() error
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
	RenewPublication(ctx context.Context, name string) error
	StartReplication(slot, publication string, startLSN uint64) error
	WaitForReplicationMessage(ctx context.Context) (*pgx.ReplicationMessage, error)
	SendStatus(lastLSN uint64) error

	DumpSnapshot(context.Context, db.ObjectWriter, string) error
}

// PostgresWatcher allows subscribing to PostgreSQL logical replication messages.
type PostgresWatcher struct {
	conf PostgresSubscriptionConfig

	lastLSN uint64 // lastLSN holds max WAL LSN value processed by subscription.

	cancel  context.CancelFunc
	conn    postgresWatcherConnection
	handler pgoutput.Handler

	dumpWriter db.ObjectWriter

	log *logrus.Entry
}

// NewPostgresWatcher creates new watcher and initialises its connections.
func NewPostgresWatcher(
	config PostgresSubscriptionConfig,
	dbs *db.Service, replConn pgxReplicationConn,
	handler pgoutput.Handler,
	dumpWriter db.ObjectWriter,
) (*PostgresWatcher, error) {
	conn, err := newPostgresReplicationConnection(dbs, replConn)
	if err != nil {
		return nil, err
	}

	return &PostgresWatcher{
		conf:       config,
		conn:       conn,
		handler:    handler,
		dumpWriter: dumpWriter,
		log:        pkglog.NewLogger("postgres-watcher"),
	}, nil
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

	if err := w.conn.RenewPublication(ctx, w.conf.Publication); err != nil {
		return errors.Wrap(err, "failed to create publication")
	}
	w.log.Debug("Created publication: ", w.conf.Publication)

	w.log.Debug("Starting dump phase")
	dumpStart := time.Now()
	if err := w.conn.DumpSnapshot(ctx, w.dumpWriter, snapshotName); err != nil {
		return errors.Wrap(err, "dumping snapshot failed")
	}
	w.log.WithField("dumpTime", time.Since(dumpStart)).Debugf("Dump phase finished - starting replication")

	if err := w.conn.StartReplication(w.conf.Slot, w.conf.Publication, 0); err != nil {
		return errors.Wrap(err, "failed to start replication")
	}

	feed := make(chan *pgx.ReplicationMessage)
	go w.runMessageConsumer(feed)
	return w.runMessageProducer(ctx, feed)
}

func (w *PostgresWatcher) runMessageConsumer(feed <-chan *pgx.ReplicationMessage) {
	func() {
		for msg := range feed {
			if err := w.handleMessage(msg); err != nil {
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
			return w.conn.Close()
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

func (w *PostgresWatcher) handleMessage(msg *pgx.ReplicationMessage) error {
	if msg.WalMessage != nil {
		if err := w.handleWalMessage(msg.WalMessage); err != nil {
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

func (w *PostgresWatcher) handleWalMessage(msg *pgx.WalMessage) error {
	if msg.WalStart > w.lastLSN {
		w.lastLSN = msg.WalStart
	}

	logmsg, err := pgoutput.Parse(msg.WalData)
	if err != nil {
		return errors.Wrap(err, "invalid pgoutput message")
	}

	if err := w.handler(logmsg); err != nil {
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
