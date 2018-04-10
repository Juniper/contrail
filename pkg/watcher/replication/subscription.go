package replication

import (
	"context"
	"fmt"
	"io"
	"strings"
	"time"

	pkglog "github.com/Juniper/contrail/pkg/log"
	"github.com/jackc/pgx"
	"github.com/kyleconroy/pgoutput"
	"github.com/sirupsen/logrus"
)

// SubConfig stores configuration for logical replication connection used for Subsctiption object.
type SubConfig struct {
	pgx.ConnConfig

	Name          string
	Publication   string
	StatusTimeout time.Duration
}

type replSlotManager interface {
	DropReplicationSlot(slotName string) (err error)
	CreateReplicationSlotEx(slotName, outputPlugin string) (consistentPoint string, snapshotName string, err error)
}

type statusSender interface {
	SendStandbyStatus(k *pgx.StandbyStatus) (err error)
}

type replConn interface {
	replSlotManager
	statusSender
	io.Closer

	StartReplication(slotName string, startLsn uint64, timeline int64, pluginArguments ...string) (err error)
	WaitForReplicationMessage(ctx context.Context) (*pgx.ReplicationMessage, error)
}

type pgConn interface {
	execer
	io.Closer

	BeginEx(ctx context.Context, txOptions *pgx.TxOptions) (*pgx.Tx, error)
}

// Subscription allows subscribing to Postgresql logical replication messages.
type Subscription struct {
	SubConfig

	replConn replConn
	conn     pgConn
	handler  pgoutput.Handler
	log      *logrus.Entry
}

// NewSubscription creates new subscription and initialises its connections.
func NewSubscription(config SubConfig, handler pgoutput.Handler) (*Subscription, error) {
	replConn, err := pgx.ReplicationConnect(config.ConnConfig)
	if err != nil {
		return nil, err
	}

	conn, err := pgx.Connect(config.ConnConfig)
	if err != nil {
		return nil, err
	}

	return &Subscription{
		SubConfig: config,
		replConn:  replConn,
		conn:      conn,
		handler:   handler,
		log:       pkglog.NewLogger("subscription"),
	}, nil
}

func pluginArgs(version, publication string) string {
	return fmt.Sprintf(`("proto_version" '%s', "publication_names" '%s')`, version, publication)
}

// Start starts replication process.
func (s *Subscription) Start(ctx context.Context) error {
	_ = s.replConn.DropReplicationSlot(s.Name)

	// If creating the replication slot fails with code 42710, this means
	// the replication slot already exists.
	consistentPoint, snapshotName, err := s.replConn.CreateReplicationSlotEx(s.Name, "pgoutput")
	if err != nil {
		if pgerr, ok := err.(pgx.PgError); !ok || pgerr.Code != "42710" {
			return fmt.Errorf("failed to create replication slot: %s", err)
		}
	}

	s.log.Debug("consistentPoint: ", consistentPoint)
	s.log.Debug("snapshotName: ", snapshotName)

	maxWal, err := pgx.ParseLSN(consistentPoint)
	if err != nil {
		return fmt.Errorf("error parsing received LSN: %v", err)
	}

	_ = dropPublication(s.conn, s.Publication)
	err = createPublicationForAll(s.conn, s.Publication)
	if err != nil {
		return fmt.Errorf("failed to create publication: %s", err)
	}
	s.log.Debug("Created publication: ", s.Publication)

	_ = snapshotName // TODO(Michal): Use snapshotName to load initial database state

	err = s.replConn.StartReplication(s.Name, 0, -1, pluginArgs("1", s.Publication))
	if err != nil {
		return fmt.Errorf("failed to start replication: %s", err)
	}

	tick := time.NewTicker(s.StatusTimeout).C
	for {
		select {
		case <-ctx.Done():
			errs := []string{}
			errs = append(errs, s.conn.Close().Error())
			errs = append(errs, s.replConn.Close().Error())
			if len(errs) > 0 {
				return fmt.Errorf("Errors while closing: %s", strings.Join(errs, "\n"))
			}
			return nil
		case <-tick:
			s.log.Debug("Sending standby status with position:  ", pgx.FormatLSN(maxWal))
			if err := sendStatus(s.replConn, maxWal); err != nil {
				return err
			}
		default:
			var message *pgx.ReplicationMessage
			wctx, cancel := context.WithTimeout(ctx, s.StatusTimeout)
			message, err = s.replConn.WaitForReplicationMessage(wctx)
			cancel()
			if err == context.DeadlineExceeded {
				continue
			}
			if err != nil {
				return fmt.Errorf("replication failed: %s", err)
			}

			if message.WalMessage != nil {
				if message.WalMessage.WalStart > maxWal {
					maxWal = message.WalMessage.WalStart
				}
				logmsg, err := pgoutput.Parse(message.WalMessage.WalData)
				if err != nil {
					return fmt.Errorf("invalid pgoutput message: %s", err)
				}
				if err := s.handler(logmsg); err != nil {
					return fmt.Errorf("error handling waldata: %s", err)
				}
			}

			if message.ServerHeartbeat != nil {
				if message.ServerHeartbeat.ReplyRequested == 1 {
					s.log.Info("Server requested reply")
					if err := sendStatus(s.replConn, maxWal); err != nil {
						return err
					}
				}
			}
		}
	}
}

func sendStatus(s statusSender, maxWal uint64) error {
	k, err := pgx.NewStandbyStatus(maxWal)
	if err != nil {
		return fmt.Errorf("error creating standby status: %s", err)
	}
	if err := s.SendStandbyStatus(k); err != nil {
		return fmt.Errorf("failed to send standy status: %s", err)
	}
	return nil
}
