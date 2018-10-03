// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package gamelift

const (

	// ErrCodeConflictException for service response error code
	// "ConflictException".
	//
	// The requested operation would cause a conflict with the current state of
	// a service resource associated with the request. Resolve the conflict before
	// retrying this request.
	ErrCodeConflictException = "ConflictException"

	// ErrCodeFleetCapacityExceededException for service response error code
	// "FleetCapacityExceededException".
	//
	// The specified fleet has no available instances to fulfill a CreateGameSession
	// request. Clients can retry such requests immediately or after a waiting period.
	ErrCodeFleetCapacityExceededException = "FleetCapacityExceededException"

	// ErrCodeGameSessionFullException for service response error code
	// "GameSessionFullException".
	//
	// The game instance is currently full and cannot allow the requested player(s)
	// to join. Clients can retry such requests immediately or after a waiting period.
	ErrCodeGameSessionFullException = "GameSessionFullException"

	// ErrCodeIdempotentParameterMismatchException for service response error code
	// "IdempotentParameterMismatchException".
	//
	// A game session with this custom ID string already exists in this fleet. Resolve
	// this conflict before retrying this request.
	ErrCodeIdempotentParameterMismatchException = "IdempotentParameterMismatchException"

	// ErrCodeInternalServiceException for service response error code
	// "InternalServiceException".
	//
	// The service encountered an unrecoverable internal failure while processing
	// the request. Clients can retry such requests immediately or after a waiting
	// period.
	ErrCodeInternalServiceException = "InternalServiceException"

	// ErrCodeInvalidFleetStatusException for service response error code
	// "InvalidFleetStatusException".
	//
	// The requested operation would cause a conflict with the current state of
	// a resource associated with the request and/or the fleet. Resolve the conflict
	// before retrying.
	ErrCodeInvalidFleetStatusException = "InvalidFleetStatusException"

	// ErrCodeInvalidGameSessionStatusException for service response error code
	// "InvalidGameSessionStatusException".
	//
	// The requested operation would cause a conflict with the current state of
	// a resource associated with the request and/or the game instance. Resolve
	// the conflict before retrying.
	ErrCodeInvalidGameSessionStatusException = "InvalidGameSessionStatusException"

	// ErrCodeInvalidRequestException for service response error code
	// "InvalidRequestException".
	//
	// One or more parameter values in the request are invalid. Correct the invalid
	// parameter values before retrying.
	ErrCodeInvalidRequestException = "InvalidRequestException"

	// ErrCodeLimitExceededException for service response error code
	// "LimitExceededException".
	//
	// The requested operation would cause the resource to exceed the allowed service
	// limit. Resolve the issue before retrying.
	ErrCodeLimitExceededException = "LimitExceededException"

	// ErrCodeNotFoundException for service response error code
	// "NotFoundException".
	//
	// A service resource associated with the request could not be found. Clients
	// should not retry such requests.
	ErrCodeNotFoundException = "NotFoundException"

	// ErrCodeTerminalRoutingStrategyException for service response error code
	// "TerminalRoutingStrategyException".
	//
	// The service is unable to resolve the routing for a particular alias because
	// it has a terminal RoutingStrategy associated with it. The message returned
	// in this exception is the message defined in the routing strategy itself.
	// Such requests should only be retried if the routing strategy for the specified
	// alias is modified.
	ErrCodeTerminalRoutingStrategyException = "TerminalRoutingStrategyException"

	// ErrCodeUnauthorizedException for service response error code
	// "UnauthorizedException".
	//
	// The client failed authentication. Clients should not retry such requests.
	ErrCodeUnauthorizedException = "UnauthorizedException"

	// ErrCodeUnsupportedRegionException for service response error code
	// "UnsupportedRegionException".
	//
	// The requested operation is not supported in the region specified.
	ErrCodeUnsupportedRegionException = "UnsupportedRegionException"
)
