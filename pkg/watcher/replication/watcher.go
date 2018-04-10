package replication

import (
	"context"
	"fmt"
	"io"
	"time"

	pkglog "github.com/Juniper/contrail/pkg/log"
	"github.com/jackc/pgx"
	"github.com/kyleconroy/pgoutput"
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

// PostgresSubscriptionConfig stores configuration for logical replication connection used for Subsctiption object.
type PostgresSubscriptionConfig struct {
	pgx.ConnConfig

	Name          string // TODO(Michal): change to slotName
	Publication   string
	StatusTimeout time.Duration
}

type postgresWatcherConnection interface {
	io.Closer
	GetReplicationSlot(name string) (lastLSN uint64, snapshotName string, err error)
	RenewPublication(name string) error
	StartReplication(slot, publication string, startLSN uint64) error
	WaitForReplicationMessage(ctx context.Context) (*pgx.ReplicationMessage, error)
	SendStatus(lastLSN uint64) error
}

// PostgresWatcher allows subscribing to Postgresql logical replication messages.
type PostgresWatcher struct {
	PostgresSubscriptionConfig

	lastLSN uint64 // lastLSN holds max WAL LSN value processed by subscription.

	cancel  context.CancelFunc
	conn    postgresWatcherConnection
	handler pgoutput.Handler
	log     *logrus.Entry
}

// NewPostgresWatcher creates new watcher and initialises its connections.
func NewPostgresWatcher(
	config PostgresSubscriptionConfig,
	handler pgoutput.Handler,
) (*PostgresWatcher, error) {
	conn, err := newPostgresReplicationConnection(config.ConnConfig)
	if err != nil {
		return nil, err
	}

	return &PostgresWatcher{
		PostgresSubscriptionConfig: config,
		conn:    conn,
		handler: handler,
		log:     pkglog.NewLogger("postgres-watcher"),
	}, nil
}

// Watch starts subscription and store context cancel function.
func (w *PostgresWatcher) Watch(ctx context.Context) error {
	w.log.Debug("Watching events on PostgreSQL replication slot")
	ctx, cancel := context.WithCancel(ctx)
	w.cancel = cancel

	var snapshotName string
	slotLSN, snapshotName, err := w.conn.GetReplicationSlot(w.Name)
	if err != nil {
		return fmt.Errorf("error getting replication slot: %v", err)
	}
	w.log.Debug("consistentPoint: ", pgx.FormatLSN(slotLSN))
	w.log.Debug("snapshotName: ", snapshotName)

	w.lastLSN = slotLSN // TODO(Michal): get lastLSN from etcd

	err = w.conn.RenewPublication(w.Publication)
	if err != nil {
		return fmt.Errorf("failed to create publication: %s", err)
	}
	w.log.Debug("Created publication: ", w.Publication)

	_ = snapshotName // TODO(Michal): Use snapshotName to load initial database state

	err = w.conn.StartReplication(w.Name, w.Publication, 0)
	if err != nil {
		return fmt.Errorf("failed to start replication: %s", err)
	}

	return w.loop(ctx)
}

func (w *PostgresWatcher) loop(ctx context.Context) error {
	tick := time.NewTicker(w.StatusTimeout).C
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
			wctx, cancel := context.WithTimeout(ctx, w.StatusTimeout)
			message, err := w.conn.WaitForReplicationMessage(wctx)
			cancel()
			if err == context.DeadlineExceeded {
				continue
			}
			if err != nil {
				return fmt.Errorf("replication failed: %s", err)
			}
			if err = w.handleMessage(message); err != nil {
				return err
			}
		}
	}
}

func (w *PostgresWatcher) handleMessage(msg *pgx.ReplicationMessage) error {
	if msg.WalMessage != nil {
		if msg.WalMessage.WalStart > w.lastLSN {
			w.lastLSN = msg.WalMessage.WalStart
		}
		logmsg, err := pgoutput.Parse(msg.WalMessage.WalData)
		if err != nil {
			return fmt.Errorf("invalid pgoutput message: %s", err)
		}
		if err := w.handler(logmsg); err != nil {
			return fmt.Errorf("error handling waldata: %s", err)
		}
	}

	if msg.ServerHeartbeat != nil {
		if msg.ServerHeartbeat.ReplyRequested == 1 {
			w.log.Info("Server requested reply, sending standby status with position: ", pgx.FormatLSN)
			if err := w.conn.SendStatus(w.lastLSN); err != nil {
				return err
			}
		}
	}

	return nil
}

// Close stops subscription by calling cancel function of context passed in Watch.
func (w *PostgresWatcher) Close() {
	w.log.Debug("Stopping watching events on PostgreSQL replication slot")
	w.cancel()
}
