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

// SubConfig stores configuration for logical replication connection used for Subsctiption object.
type SubConfig struct {
	Name          string
	Publication   string
	StatusTimeout time.Duration
}

type conn interface {
	io.Closer
	GetReplicationSlot(name string) (maxWal uint64, snapshotName string, err error)
	RenewPublication(name string) error
	StartReplication(slot, publication string, startLSN uint64) error
	WaitForReplicationMessage(ctx context.Context) (*pgx.ReplicationMessage, error)
	SendStatus(maxWal uint64) error
}

// Subscription allows subscribing to Postgresql logical replication messages.
type Subscription struct {
	SubConfig

	maxWal uint64

	conn    conn
	handler pgoutput.Handler
	log     *logrus.Entry
}

// NewSubscription creates new subscription and initialises its connections.
func NewSubscription(config SubConfig, conn conn, handler pgoutput.Handler) (*Subscription, error) {
	return &Subscription{
		SubConfig: config,
		conn:      conn,
		handler:   handler,
		log:       pkglog.NewLogger("subscription"),
	}, nil
}

// Start starts replication process.
func (s *Subscription) Start(ctx context.Context) (err error) {
	var snapshotName string
	s.maxWal, snapshotName, err = s.conn.GetReplicationSlot(s.Name)
	if err != nil {
		return fmt.Errorf("error getting replication slot: %v", err)
	}

	s.log.Debug("consistentPoint: ", pgx.FormatLSN(s.maxWal))
	s.log.Debug("snapshotName: ", snapshotName)

	err = s.conn.RenewPublication(s.Publication)
	if err != nil {
		return fmt.Errorf("failed to create publication: %s", err)
	}
	s.log.Debug("Created publication: ", s.Publication)

	_ = snapshotName // TODO(Michal): Use snapshotName to load initial database state

	err = s.conn.StartReplication(s.Name, s.Publication, 0)
	if err != nil {
		return fmt.Errorf("failed to start replication: %s", err)
	}

	return s.loop(ctx)
}

func (s *Subscription) loop(ctx context.Context) error {
	tick := time.NewTicker(s.StatusTimeout).C
	for {
		select {
		case <-ctx.Done():
			return s.conn.Close()
		case <-tick:
			s.log.Debug("Sending standby status with position: ", pgx.FormatLSN(s.maxWal))
			if err := s.conn.SendStatus(s.maxWal); err != nil {
				return err
			}
		default:
			wctx, cancel := context.WithTimeout(ctx, s.StatusTimeout)
			message, err := s.conn.WaitForReplicationMessage(wctx)
			cancel()
			if err == context.DeadlineExceeded {
				continue
			}
			if err != nil {
				return fmt.Errorf("replication failed: %s", err)
			}

			if err = s.handleMessage(message); err != nil {
				return err
			}
		}
	}
}

func (s *Subscription) handleMessage(msg *pgx.ReplicationMessage) error {
	if msg.WalMessage != nil {
		if msg.WalMessage.WalStart > s.maxWal {
			s.maxWal = msg.WalMessage.WalStart
		}
		logmsg, err := pgoutput.Parse(msg.WalMessage.WalData)
		if err != nil {
			return fmt.Errorf("invalid pgoutput message: %s", err)
		}
		if err := s.handler(logmsg); err != nil {
			return fmt.Errorf("error handling waldata: %s", err)
		}
	}

	if msg.ServerHeartbeat != nil {
		if msg.ServerHeartbeat.ReplyRequested == 1 {
			s.log.Info("Server requested reply, sending standby status with position: ", pgx.FormatLSN)
			if err := s.conn.SendStatus(s.maxWal); err != nil {
				return err
			}
		}
	}

	return nil
}
