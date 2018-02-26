package common

import (
	"net/http"

	"github.com/labstack/echo"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

//ErrorNotFound for not found error.
var ErrorNotFound = grpc.Errorf(codes.NotFound, "not found")

//ErrorUnauthenticated for unauthenticated error.
var ErrorUnauthenticated = grpc.Errorf(codes.Unauthenticated, "Unauthenticated")

//ErrorPermissionDenied for permission denied errror.
var ErrorPermissionDenied = grpc.Errorf(codes.PermissionDenied, "Permission Denied")

//ErrorInternal for Internal Server Error.
var ErrorInternal = grpc.Errorf(codes.Internal, "Internal Server Error")

//ErrorConflict is for resource conflict error.
var ErrorConflict = grpc.Errorf(codes.AlreadyExists, "Resource conflict")

//ErrorBadRequest makes bad request error.
func ErrorBadRequest(message string) error {
	if message == "" {
		message = "bad request error."
	}
	return grpc.Errorf(codes.InvalidArgument, message)
}

// HTTPStatusFromCode converts a gRPC error code into the corresponding HTTP response status.
func HTTPStatusFromCode(code codes.Code) int {
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

//ToHTTPError translates grpc error to error.
func ToHTTPError(err error) error {
	code := HTTPStatusFromCode(grpc.Code(err))
	return echo.NewHTTPError(code, grpc.ErrorDesc(err))
}
