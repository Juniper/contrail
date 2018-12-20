package replication

import (
	"fmt"
	"sync"

	"github.com/jackc/pgx"
)

type lsnCounter struct {
	// receivedLSN holds max WAL LSN value received from master.
	receivedLSN uint64

	// savedLSN holds max WAL LSN value saved in etcd.
	// We can also assume that savedLSN <= receivedLSN.
	savedLSN uint64

	txnsInProgress int

	m sync.Mutex
}

func (c *lsnCounter) UpdateReceivedLSN(lsn uint64) {
	c.m.Lock()
	defer c.m.Unlock()

	if c.receivedLSN < lsn {
		c.receivedLSN = lsn
		fmt.Println("Updated lsn", pgx.FormatLSN(c.receivedLSN), pgx.FormatLSN(c.savedLSN))
	}
}

func (c *lsnCounter) LSNValues() (receivedLSN, savedLSN uint64) {
	c.m.Lock()
	defer c.m.Unlock()

	if c.txnsInProgress == 0 {
		c.savedLSN = c.receivedLSN
	}

	return c.receivedLSN, c.savedLSN
}

func (c *lsnCounter) TxnStarted() {
	c.m.Lock()
	defer c.m.Unlock()

	c.txnsInProgress++
}

func (c *lsnCounter) TxnFinished(processedLSN uint64) {
	c.m.Lock()
	defer c.m.Unlock()

	if c.savedLSN < processedLSN {
		c.savedLSN = processedLSN
	}

	if c.txnsInProgress > 0 {
		c.txnsInProgress--
	}
}
