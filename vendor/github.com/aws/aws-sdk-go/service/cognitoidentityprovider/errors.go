// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cognitoidentityprovider

const (

	// ErrCodeAliasExistsException for service response error code
	// "AliasExistsException".
	//
	// This exception is thrown when a user tries to confirm the account with an
	// email or phone number that has already been supplied as an alias from a different
	// account. This exception tells user that an account with this email or phone
	// already exists.
	ErrCodeAliasExistsException = "AliasExistsException"

	// ErrCodeCodeDeliveryFailureException for service response error code
	// "CodeDeliveryFailureException".
	//
	// This exception is thrown when a verification code fails to deliver successfully.
	ErrCodeCodeDeliveryFailureException = "CodeDeliveryFailureException"

	// ErrCodeCodeMismatchException for service response error code
	// "CodeMismatchException".
	//
	// This exception is thrown if the provided code does not match what the server
	// was expecting.
	ErrCodeCodeMismatchException = "CodeMismatchException"

	// ErrCodeConcurrentModificationException for service response error code
	// "ConcurrentModificationException".
	//
	// This exception is thrown if two or more modifications are happening concurrently.
	ErrCodeConcurrentModificationException = "ConcurrentModificationException"

	// ErrCodeDuplicateProviderException for service response error code
	// "DuplicateProviderException".
	//
	// This exception is thrown when the provider is already supported by the user
	// pool.
	ErrCodeDuplicateProviderException = "DuplicateProviderException"

	// ErrCodeEnableSoftwareTokenMFAException for service response error code
	// "EnableSoftwareTokenMFAException".
	//
	// This exception is thrown when there is a code mismatch and the service fails
	// to configure the software token TOTP multi-factor authentication (MFA).
	ErrCodeEnableSoftwareTokenMFAException = "EnableSoftwareTokenMFAException"

	// ErrCodeExpiredCodeException for service response error code
	// "ExpiredCodeException".
	//
	// This exception is thrown if a code has expired.
	ErrCodeExpiredCodeException = "ExpiredCodeException"

	// ErrCodeGroupExistsException for service response error code
	// "GroupExistsException".
	//
	// This exception is thrown when Amazon Cognito encounters a group that already
	// exists in the user pool.
	ErrCodeGroupExistsException = "GroupExistsException"

	// ErrCodeInternalErrorException for service response error code
	// "InternalErrorException".
	//
	// This exception is thrown when Amazon Cognito encounters an internal error.
	ErrCodeInternalErrorException = "InternalErrorException"

	// ErrCodeInvalidEmailRoleAccessPolicyException for service response error code
	// "InvalidEmailRoleAccessPolicyException".
	//
	// This exception is thrown when Amazon Cognito is not allowed to use your email
	// identity. HTTP status code: 400.
	ErrCodeInvalidEmailRoleAccessPolicyException = "InvalidEmailRoleAccessPolicyException"

	// ErrCodeInvalidLambdaResponseException for service response error code
	// "InvalidLambdaResponseException".
	//
	// This exception is thrown when the Amazon Cognito service encounters an invalid
	// AWS Lambda response.
	ErrCodeInvalidLambdaResponseException = "InvalidLambdaResponseException"

	// ErrCodeInvalidOAuthFlowException for service response error code
	// "InvalidOAuthFlowException".
	//
	// This exception is thrown when the specified OAuth flow is invalid.
	ErrCodeInvalidOAuthFlowException = "InvalidOAuthFlowException"

	// ErrCodeInvalidParameterException for service response error code
	// "InvalidParameterException".
	//
	// This exception is thrown when the Amazon Cognito service encounters an invalid
	// parameter.
	ErrCodeInvalidParameterException = "InvalidParameterException"

	// ErrCodeInvalidPasswordException for service response error code
	// "InvalidPasswordException".
	//
	// This exception is thrown when the Amazon Cognito service encounters an invalid
	// password.
	ErrCodeInvalidPasswordException = "InvalidPasswordException"

	// ErrCodeInvalidSmsRoleAccessPolicyException for service response error code
	// "InvalidSmsRoleAccessPolicyException".
	//
	// This exception is returned when the role provided for SMS configuration does
	// not have permission to publish using Amazon SNS.
	ErrCodeInvalidSmsRoleAccessPolicyException = "InvalidSmsRoleAccessPolicyException"

	// ErrCodeInvalidSmsRoleTrustRelationshipException for service response error code
	// "InvalidSmsRoleTrustRelationshipException".
	//
	// This exception is thrown when the trust relationship is invalid for the role
	// provided for SMS configuration. This can happen if you do not trust cognito-idp.amazonaws.com
	// or the external ID provided in the role does not match what is provided in
	// the SMS configuration for the user pool.
	ErrCodeInvalidSmsRoleTrustRelationshipException = "InvalidSmsRoleTrustRelationshipException"

	// ErrCodeInvalidUserPoolConfigurationException for service response error code
	// "InvalidUserPoolConfigurationException".
	//
	// This exception is thrown when the user pool configuration is invalid.
	ErrCodeInvalidUserPoolConfigurationException = "InvalidUserPoolConfigurationException"

	// ErrCodeLimitExceededException for service response error code
	// "LimitExceededException".
	//
	// This exception is thrown when a user exceeds the limit for a requested AWS
	// resource.
	ErrCodeLimitExceededException = "LimitExceededException"

	// ErrCodeMFAMethodNotFoundException for service response error code
	// "MFAMethodNotFoundException".
	//
	// This exception is thrown when Amazon Cognito cannot find a multi-factor authentication
	// (MFA) method.
	ErrCodeMFAMethodNotFoundException = "MFAMethodNotFoundException"

	// ErrCodeNotAuthorizedException for service response error code
	// "NotAuthorizedException".
	//
	// This exception is thrown when a user is not authorized.
	ErrCodeNotAuthorizedException = "NotAuthorizedException"

	// ErrCodePasswordResetRequiredException for service response error code
	// "PasswordResetRequiredException".
	//
	// This exception is thrown when a password reset is required.
	ErrCodePasswordResetRequiredException = "PasswordResetRequiredException"

	// ErrCodePreconditionNotMetException for service response error code
	// "PreconditionNotMetException".
	//
	// This exception is thrown when a precondition is not met.
	ErrCodePreconditionNotMetException = "PreconditionNotMetException"

	// ErrCodeResourceNotFoundException for service response error code
	// "ResourceNotFoundException".
	//
	// This exception is thrown when the Amazon Cognito service cannot find the
	// requested resource.
	ErrCodeResourceNotFoundException = "ResourceNotFoundException"

	// ErrCodeScopeDoesNotExistException for service response error code
	// "ScopeDoesNotExistException".
	//
	// This exception is thrown when the specified scope does not exist.
	ErrCodeScopeDoesNotExistException = "ScopeDoesNotExistException"

	// ErrCodeSoftwareTokenMFANotFoundException for service response error code
	// "SoftwareTokenMFANotFoundException".
	//
	// This exception is thrown when the software token TOTP multi-factor authentication
	// (MFA) is not enabled for the user pool.
	ErrCodeSoftwareTokenMFANotFoundException = "SoftwareTokenMFANotFoundException"

	// ErrCodeTooManyFailedAttemptsException for service response error code
	// "TooManyFailedAttemptsException".
	//
	// This exception is thrown when the user has made too many failed attempts
	// for a given action (e.g., sign in).
	ErrCodeTooManyFailedAttemptsException = "TooManyFailedAttemptsException"

	// ErrCodeTooManyRequestsException for service response error code
	// "TooManyRequestsException".
	//
	// This exception is thrown when the user has made too many requests for a given
	// operation.
	ErrCodeTooManyRequestsException = "TooManyRequestsException"

	// ErrCodeUnexpectedLambdaException for service response error code
	// "UnexpectedLambdaException".
	//
	// This exception is thrown when the Amazon Cognito service encounters an unexpected
	// exception with the AWS Lambda service.
	ErrCodeUnexpectedLambdaException = "UnexpectedLambdaException"

	// ErrCodeUnsupportedIdentityProviderException for service response error code
	// "UnsupportedIdentityProviderException".
	//
	// This exception is thrown when the specified identifier is not supported.
	ErrCodeUnsupportedIdentityProviderException = "UnsupportedIdentityProviderException"

	// ErrCodeUnsupportedUserStateException for service response error code
	// "UnsupportedUserStateException".
	//
	// The request failed because the user is in an unsupported state.
	ErrCodeUnsupportedUserStateException = "UnsupportedUserStateException"

	// ErrCodeUserImportInProgressException for service response error code
	// "UserImportInProgressException".
	//
	// This exception is thrown when you are trying to modify a user pool while
	// a user import job is in progress for that pool.
	ErrCodeUserImportInProgressException = "UserImportInProgressException"

	// ErrCodeUserLambdaValidationException for service response error code
	// "UserLambdaValidationException".
	//
	// This exception is thrown when the Amazon Cognito service encounters a user
	// validation exception with the AWS Lambda service.
	ErrCodeUserLambdaValidationException = "UserLambdaValidationException"

	// ErrCodeUserNotConfirmedException for service response error code
	// "UserNotConfirmedException".
	//
	// This exception is thrown when a user is not confirmed successfully.
	ErrCodeUserNotConfirmedException = "UserNotConfirmedException"

	// ErrCodeUserNotFoundException for service response error code
	// "UserNotFoundException".
	//
	// This exception is thrown when a user is not found.
	ErrCodeUserNotFoundException = "UserNotFoundException"

	// ErrCodeUserPoolAddOnNotEnabledException for service response error code
	// "UserPoolAddOnNotEnabledException".
	//
	// This exception is thrown when user pool add-ons are not enabled.
	ErrCodeUserPoolAddOnNotEnabledException = "UserPoolAddOnNotEnabledException"

	// ErrCodeUserPoolTaggingException for service response error code
	// "UserPoolTaggingException".
	//
	// This exception is thrown when a user pool tag cannot be set or updated.
	ErrCodeUserPoolTaggingException = "UserPoolTaggingException"

	// ErrCodeUsernameExistsException for service response error code
	// "UsernameExistsException".
	//
	// This exception is thrown when Amazon Cognito encounters a user name that
	// already exists in the user pool.
	ErrCodeUsernameExistsException = "UsernameExistsException"
)
