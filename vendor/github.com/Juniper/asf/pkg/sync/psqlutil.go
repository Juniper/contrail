package sync

import (
	"strings"
	"sync"

	"github.com/jackc/pgx"
)

const (
	// PostgreSQLPublicationName contains name of publication created in database.
	PostgreSQLPublicationName = "syncpub"
)

// SlotName transforms watcher ID to replication slot name.
func SlotName(id string) string {
	return strings.Replace(id, "-", "_", -1)
}

type LSN uint64

func (l LSN) String() string {
	return pgx.FormatLSN(uint64(l))
}

func MessageLSN(msg *pgx.WalMessage) LSN {
	if msg.WalStart < msg.ServerWalEnd {
		return LSN(msg.ServerWalEnd)
	}
	return LSN(msg.WalStart)
}

func ParseLSN(s string) (LSN, error) {
	l, err := pgx.ParseLSN(s)
	return LSN(l), err
}

type lsnCounter struct {
	// receivedLSN holds max WAL LSN value received from master.
	received LSN

	// savedLSN holds max WAL LSN value saved in etcd.
	// We can also assume that savedLSN <= receivedLSN.
	saved LSN

	txnsInProgress int

	m sync.Mutex
}

func (c *lsnCounter) updateReceivedLSN(value LSN) {
	c.m.Lock()
	defer c.m.Unlock()

	if c.received < value {
		c.received = value
	}
}

func (c *lsnCounter) lsnValues() (received, saved LSN) {
	c.m.Lock()
	defer c.m.Unlock()

	if c.txnsInProgress == 0 {
		c.saved = c.received
	}

	return c.received, c.saved
}

func (c *lsnCounter) txnStarted() {
	c.m.Lock()
	defer c.m.Unlock()

	c.txnsInProgress++
}

func (c *lsnCounter) txnFinished(processed LSN) {
	c.m.Lock()
	defer c.m.Unlock()

	if c.saved < processed {
		c.saved = processed
	}

	if c.txnsInProgress > 0 {
		c.txnsInProgress--
	}
}
