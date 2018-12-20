package replication

import (
	"fmt"
	"io"
	"sync"

	"github.com/jackc/pgx"
	"github.com/pkg/errors"
)

type syncError struct {
	error
}

func (e syncError) Cause() error {
	return errors.Cause(e.error)
}

func (e syncError) Temporary() bool {
	c := errors.Cause(e.error)
	if e, ok := c.(pgx.PgError); ok {
		if e.Code == "0A000" {
			return true
		}
	}
	if c == io.EOF || c == pgx.ErrConnBusy {
		return true
	}
	return false
}

func wrapError(err error) error {
	if err != nil {
		return syncError{error: err}
	}
	return nil
}

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
