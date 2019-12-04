package analytics

import (
	"context"
	"database/sql/driver"
	"time"

	"github.com/Juniper/contrail/pkg/collector"
)

const (
	CommitOperation = "COMMIT"
	SQLApplication  = "SQL"
)

func WithCommitLatencyReporting(c collector.Collector) func(driver.Driver) driver.Driver {
	if c == nil {
		return func(d driver.Driver) driver.Driver { return d }
	}
	return func(d driver.Driver) driver.Driver {
		return &latencyReportingDriver{Driver: d, c: c}
	}
}

type latencyReportingDriver struct {
	driver.Driver

	c collector.Collector
}

// Open returns a new connection to the database.
// The name is a string in a driver-specific format.
//
// Open may return a cached connection (one previously
// closed), but doing so is unnecessary; the sql package
// maintains a pool of idle connections for efficient re-use.
//
// The returned connection is only used by one goroutine at a
// time.
func (l *latencyReportingDriver) Open(name string) (driver.Conn, error) {
	conn, err := l.Driver.Open(name)
	return &latencyReportingConn{Conn: conn, c: l.c}, err
}

type latencyReportingConn struct {
	driver.Conn

	c collector.Collector
}

// BeginTx starts and returns a new transaction.
// If the context is canceled by the user the sql package will
// call Tx.Rollback before discarding and closing the connection.
//
// This must check opts.Isolation to determine if there is a set
// isolation level. If the driver does not support a non-default
// level and one is set or if there is a non-default isolation level
// that is not supported, an error must be returned.
//
// This must also check opts.ReadOnly to determine if the read-only
// value is true to either set the read-only transaction property if supported
// or return an error if it is not supported.
func (l *latencyReportingConn) BeginTx(ctx context.Context, opts driver.TxOptions) (driver.Tx, error) {
	t, err := l.Conn.Begin()
	return &latencyReportingTx{Tx: t, ctx: ctx, c: l.c}, err
}

type latencyReportingTx struct {
	driver.Tx

	ctx context.Context
	c   collector.Collector
}

func (l *latencyReportingTx) Commit() error {
	start := time.Now()
	if err := l.Tx.Commit(); err != nil {
		return err
	}

	l.c.Send(
		VncAPILatencyStatsLog(l.ctx, CommitOperation, SQLApplication, int64(time.Since(start).Microseconds())),
	)
	return nil
}
