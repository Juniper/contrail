package analytics

import (
	"context"
	"database/sql/driver"
	"time"

	"github.com/Juniper/contrail/pkg/collector"
)

const (
	commitOperation = "COMMIT"
	sqlApplication  = "SQL"
)

// WithCommitLatencyReporting returns database driver wrapper reporting database commit latency.
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

// Open opens database connection using underlying database driver and returns wrapped connection.
func (l *latencyReportingDriver) Open(name string) (driver.Conn, error) {
	conn, err := l.Driver.Open(name)
	return &latencyReportingConn{Conn: conn, c: l.c}, err
}

type latencyReportingConn struct {
	driver.Conn

	c collector.Collector
}

// BeginTx begins transaction using underlying database driver and wraps wrapped transaction.
func (l *latencyReportingConn) BeginTx(ctx context.Context, opts driver.TxOptions) (driver.Tx, error) {
	t, err := l.Conn.Begin()
	return &latencyReportingTx{Tx: t, ctx: ctx, c: l.c}, err
}

type latencyReportingTx struct {
	driver.Tx

	ctx context.Context
	c   collector.Collector
}

// Commit commits transaction using underlying database driver and reports database commit latency.
func (l *latencyReportingTx) Commit() error {
	start := time.Now()
	if err := l.Tx.Commit(); err != nil {
		return err
	}

	l.c.Send(
		VncAPILatencyStatsLog(l.ctx, commitOperation, sqlApplication, int64(time.Since(start)/time.Microsecond)),
	)
	return nil
}
