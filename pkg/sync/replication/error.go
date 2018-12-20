package replication

import (
	"context"
	"io"
	"strings"

	"github.com/jackc/pgx"
	"github.com/lib/pq"
	"github.com/pkg/errors"
)

const pgQueryCanceledErrorCode = "57014"

func isContextCancellationError(err error) bool {
	if pqErr, ok := errors.Cause(err).(*pq.Error); ok {
		if pqErr.Code == pgQueryCanceledErrorCode {
			return true
		}
	}
	if errors.Cause(err) == context.Canceled {
		return true
	}
	return false
}

type syncError struct {
	error
}

func wrapError(err error) error {
	if err != nil {
		return syncError{error: err}
	}
	return nil
}

func (e syncError) Cause() error {
	return errors.Cause(e.error)
}

func (e syncError) Temporary() bool {
	c := errors.Cause(e.error)
	if e, ok := c.(pgx.PgError); ok {
		if e.Code == "0A000" && strings.Contains(e.Message, "logical") {
			return true
		}
	}
	if c == io.EOF || c == pgx.ErrConnBusy {
		return true
	}
	return false
}
