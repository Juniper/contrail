package replication

import (
	"context"

	pkglog "github.com/Juniper/contrail/pkg/log"
	"github.com/jackc/pgx"
	"github.com/kyleconroy/pgoutput"
	"github.com/sirupsen/logrus"
)

type abstractCanal interface {
	Run() error
	Close()
}

type noopCanal struct{}

func (c *noopCanal) Run() error { return nil }

func (c *noopCanal) Close() {}

// BinlogWatcher uses canal to read MySQL binlog events.
type BinlogWatcher struct {
	canal abstractCanal
	log   *logrus.Entry
}

// NewBinlogWatcher creates new BinlogWatcher listening on provided canal.
func NewBinlogWatcher(c abstractCanal) *BinlogWatcher {
	if c == nil {
		c = &noopCanal{}
	}

	return &BinlogWatcher{
		canal: c,
		log:   pkglog.NewLogger("binlog-watcher"),
	}
}

// Watch starts listening on a canal.
func (w *BinlogWatcher) Watch(context.Context) error {
	w.log.Debug("Watching events on MySQL binlog")
	return w.canal.Run()
}

// Close closes canal.
func (w *BinlogWatcher) Close() {
	w.log.Debug("Stopping watching events on MySQL binlog")
	w.canal.Close()
}

// MessageHandler handles pgoutput logical replication messages.
type MessageHandler interface {
	Handle(pgoutput.Message) error
}

type subscriptionStarter interface {
	Start(context.Context, *pgx.ReplicationConn, pgoutput.Handler) error
}

// SubscriptionWatcher uses subscription to read PostgreSQL logical replciation messages.
type SubscriptionWatcher struct {
	conn    *pgx.ReplicationConn // TODO(Michal): Find a way to change this into interface
	starter subscriptionStarter

	handler MessageHandler
	cancel  context.CancelFunc

	log *logrus.Entry
}

// NewSubscriptionWatcher returns new SubscriptionWatcher.
func NewSubscriptionWatcher(
	replicationConn *pgx.ReplicationConn,
	starter subscriptionStarter,
	handler MessageHandler,
) *SubscriptionWatcher {
	return &SubscriptionWatcher{
		conn:    replicationConn,
		starter: starter,
		handler: handler,
		log:     pkglog.NewLogger("subscription-watcher"),
	}
}

// Watch starts subscription and store context cancel function.
func (w *SubscriptionWatcher) Watch(ctx context.Context) error {
	w.log.Debug("Watching events on PostgreSQL replication slot")
	ctx, cancel := context.WithCancel(ctx)
	w.cancel = cancel

	return w.starter.Start(ctx, w.conn, w.handler.Handle)
}

// Close stops subscription by calling cancel function.
func (w *SubscriptionWatcher) Close() {
	w.log.Debug("Stopping watching events on PostgreSQL replication slot")
	w.cancel()
}
