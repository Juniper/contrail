package errutil

import (
	"net/http"
	"strings"

	"github.com/labstack/echo"
	"github.com/pkg/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// MultiError implements errors with multiple causes.
type MultiError []error

// Error implements default errors interface for MultiError.
func (m MultiError) Error() string {
	var msgs []string
	for _, e := range m {
		msgs = append(msgs, e.Error())
	}
	return strings.Join(msgs, "\n")
}

// Cause returns the first error.
func (m MultiError) Cause() error {
	if len(m) == 0 {
		return nil
	}
	return m[0]
}

// WithErrors returns new MultiError with appended errors which are non-nil.
func (m MultiError) WithErrors(errors ...error) (mErr MultiError) {
	mErr = m
	for _, err := range errors {
		if err != nil {
			mErr = append(mErr, err)
		}
	}
	return mErr
}

// CauseCode returns wrapped grpc error code
func CauseCode(err error) codes.Code {
	return status.Code(errors.Cause(err))
}

// IsNotFound returns true if error is of NotFound type.
func IsNotFound(err error) bool {
	return status.Code(errors.Cause(err)) == codes.NotFound
}

// IsConflict returns true if error is of Conflict type.
func IsConflict(err error) bool {
	return status.Code(errors.Cause(err)) == codes.AlreadyExists
}

// IsBadRequest returns true if error is of BadRequest type.
func IsBadRequest(err error) bool {
	return status.Code(errors.Cause(err)) == codes.InvalidArgument
}

// IsQuotaExceeded returns true if error is of QuotaExceeded type.
func IsQuotaExceeded(err error) bool {
	return status.Code(errors.Cause(err)) == codes.FailedPrecondition
}

// IsInternal returns true if error is of Internal type.
func IsInternal(err error) bool {
	return status.Code(errors.Cause(err)) == codes.Internal
}

// IsUnauthenticated returns true if error is of Unauthenticated type.
func IsUnauthenticated(err error) bool {
	return status.Code(errors.Cause(err)) == codes.Unauthenticated
}

// IsForbidden returns true if error is of Forbidden type.
func IsForbidden(err error) bool {
	return status.Code(errors.Cause(err)) == codes.PermissionDenied
}

// ErrorForbiddenf makes forbidden error with format.
func ErrorForbiddenf(format string, a ...interface{}) error {
	return status.Errorf(codes.PermissionDenied, format, a...)
}

// ErrorForbidden makes forbidden error.
func ErrorForbidden(msgs ...string) error {
	return status.Error(codes.PermissionDenied, errorMessage(msgs, "permission denied"))
}

// ErrorBadRequestf makes bad request error with format.
func ErrorBadRequestf(format string, a ...interface{}) error {
	return status.Errorf(codes.InvalidArgument, format, a...)
}

// ErrorBadRequest makes bad request error.
func ErrorBadRequest(msgs ...string) error {
	return status.Error(codes.InvalidArgument, errorMessage(msgs, "bad request"))
}

// ErrorNotFoundf makes not found error with format.
func ErrorNotFoundf(format string, a ...interface{}) error {
	return status.Errorf(codes.NotFound, format, a...)
}

// ErrorNotFound makes not found error.
func ErrorNotFound(msgs ...string) error {
	return status.Error(codes.NotFound, errorMessage(msgs, "not found"))
}

// ErrorConflictf makes already exists error with format.
func ErrorConflictf(format string, a ...interface{}) error {
	return status.Errorf(codes.AlreadyExists, format, a...)
}

// ErrorConflict makes already exists error.
func ErrorConflict(msgs ...string) error {
	return status.Error(codes.AlreadyExists, errorMessage(msgs, "resource conflict"))
}

// ErrorInternalf makes unknown error with format.
func ErrorInternalf(format string, a ...interface{}) error {
	return status.Errorf(codes.Internal, format, a...)
}

// ErrorInternal makes unknown error.
func ErrorInternal(msgs ...string) error {
	return status.Error(codes.Internal, errorMessage(msgs, "internal server error"))
}

// ErrorQuotaExceededf makes quota exceed error with format.
func ErrorQuotaExceededf(format string, a ...interface{}) error {
	return status.Errorf(codes.FailedPrecondition, format, a...)
}

// ErrorQuotaExceeded makes quota exceed error.
func ErrorQuotaExceeded(msgs ...string) error {
	return status.Error(codes.FailedPrecondition, errorMessage(msgs, "quota exceeded"))
}

// ErrorUnauthenticatedf makes unauthenticated error with format.
func ErrorUnauthenticatedf(format string, a ...interface{}) error {
	return status.Errorf(codes.Unauthenticated, format, a...)
}

// ErrorUnauthenticated makes unauthenticated error.
func ErrorUnauthenticated(msgs ...string) error {
	return status.Error(codes.Unauthenticated, errorMessage(msgs, "unauthenticated"))
}

func errorMessage(msgs []string, fallback string) string {
	if len(msgs) == 0 {
		return fallback
	}
	return strings.Join(msgs, " ")
}

func getErrorMessage(err error) string {
	if err == nil {
		return ""
	}
	return err.Error()
}

// ToHTTPError translates grpc error to error.
func ToHTTPError(err error) error {
	cause := errors.Cause(err)
	return echo.NewHTTPError(
		httpStatusFromCode(status.Code(cause)),
		getErrorMessage(err),
	)
}

// httpStatusFromCode converts a gRPC error code into the corresponding HTTP response status.
// nolint: gocyclo
func httpStatusFromCode(code codes.Code) int {
	switch code {
	case codes.OK:
		return http.StatusOK
	case codes.Canceled:
		return http.StatusRequestTimeout
	case codes.Unknown:
		return http.StatusInternalServerError
	case codes.InvalidArgument:
		return http.StatusBadRequest
	case codes.DeadlineExceeded:
		return http.StatusRequestTimeout
	case codes.NotFound:
		return http.StatusNotFound
	case codes.AlreadyExists:
		return http.StatusConflict
	case codes.PermissionDenied:
		return http.StatusForbidden
	case codes.Unauthenticated:
		return http.StatusUnauthorized
	case codes.ResourceExhausted:
		return http.StatusForbidden
	case codes.FailedPrecondition:
		return http.StatusPreconditionFailed
	case codes.Aborted:
		return http.StatusConflict
	case codes.OutOfRange:
		return http.StatusBadRequest
	case codes.Unimplemented:
		return http.StatusNotImplemented
	case codes.Internal:
		return http.StatusInternalServerError
	case codes.Unavailable:
		return http.StatusServiceUnavailable
	case codes.DataLoss:
		return http.StatusInternalServerError
	}
	return http.StatusInternalServerError
}

//StatusFromError returns HTTP status based on echo.HTTPError.
func StatusFromError(err error) int {
	if err == nil {
		return http.StatusOK
	}
	if e, ok := err.(*echo.HTTPError); ok {
		return e.Code
	}
	return http.StatusInternalServerError
}
