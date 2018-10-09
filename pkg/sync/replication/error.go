package replication

import (
	"context"

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
