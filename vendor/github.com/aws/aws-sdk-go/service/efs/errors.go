// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package efs

const (

	// ErrCodeBadRequest for service response error code
	// "BadRequest".
	//
	// Returned if the request is malformed or contains an error such as an invalid
	// parameter value or a missing required parameter.
	ErrCodeBadRequest = "BadRequest"

	// ErrCodeDependencyTimeout for service response error code
	// "DependencyTimeout".
	//
	// The service timed out trying to fulfill the request, and the client should
	// try the call again.
	ErrCodeDependencyTimeout = "DependencyTimeout"

	// ErrCodeFileSystemAlreadyExists for service response error code
	// "FileSystemAlreadyExists".
	//
	// Returned if the file system you are trying to create already exists, with
	// the creation token you provided.
	ErrCodeFileSystemAlreadyExists = "FileSystemAlreadyExists"

	// ErrCodeFileSystemInUse for service response error code
	// "FileSystemInUse".
	//
	// Returned if a file system has mount targets.
	ErrCodeFileSystemInUse = "FileSystemInUse"

	// ErrCodeFileSystemLimitExceeded for service response error code
	// "FileSystemLimitExceeded".
	//
	// Returned if the AWS account has already created the maximum number of file
	// systems allowed per account.
	ErrCodeFileSystemLimitExceeded = "FileSystemLimitExceeded"

	// ErrCodeFileSystemNotFound for service response error code
	// "FileSystemNotFound".
	//
	// Returned if the specified FileSystemId value doesn't exist in the requester's
	// AWS account.
	ErrCodeFileSystemNotFound = "FileSystemNotFound"

	// ErrCodeIncorrectFileSystemLifeCycleState for service response error code
	// "IncorrectFileSystemLifeCycleState".
	//
	// Returned if the file system's lifecycle state is not "available".
	ErrCodeIncorrectFileSystemLifeCycleState = "IncorrectFileSystemLifeCycleState"

	// ErrCodeIncorrectMountTargetState for service response error code
	// "IncorrectMountTargetState".
	//
	// Returned if the mount target is not in the correct state for the operation.
	ErrCodeIncorrectMountTargetState = "IncorrectMountTargetState"

	// ErrCodeInsufficientThroughputCapacity for service response error code
	// "InsufficientThroughputCapacity".
	//
	// Returned if there's not enough capacity to provision additional throughput.
	// This value might be returned when you try to create a file system in provisioned
	// throughput mode, when you attempt to increase the provisioned throughput
	// of an existing file system, or when you attempt to change an existing file
	// system from bursting to provisioned throughput mode.
	ErrCodeInsufficientThroughputCapacity = "InsufficientThroughputCapacity"

	// ErrCodeInternalServerError for service response error code
	// "InternalServerError".
	//
	// Returned if an error occurred on the server side.
	ErrCodeInternalServerError = "InternalServerError"

	// ErrCodeIpAddressInUse for service response error code
	// "IpAddressInUse".
	//
	// Returned if the request specified an IpAddress that is already in use in
	// the subnet.
	ErrCodeIpAddressInUse = "IpAddressInUse"

	// ErrCodeMountTargetConflict for service response error code
	// "MountTargetConflict".
	//
	// Returned if the mount target would violate one of the specified restrictions
	// based on the file system's existing mount targets.
	ErrCodeMountTargetConflict = "MountTargetConflict"

	// ErrCodeMountTargetNotFound for service response error code
	// "MountTargetNotFound".
	//
	// Returned if there is no mount target with the specified ID found in the caller's
	// account.
	ErrCodeMountTargetNotFound = "MountTargetNotFound"

	// ErrCodeNetworkInterfaceLimitExceeded for service response error code
	// "NetworkInterfaceLimitExceeded".
	//
	// The calling account has reached the limit for elastic network interfaces
	// for the specific AWS Region. The client should try to delete some elastic
	// network interfaces or get the account limit raised. For more information,
	// see Amazon VPC Limits (http://docs.aws.amazon.com/AmazonVPC/latest/UserGuide/VPC_Appendix_Limits.html)
	// in the Amazon VPC User Guide  (see the Network interfaces per VPC entry in
	// the table).
	ErrCodeNetworkInterfaceLimitExceeded = "NetworkInterfaceLimitExceeded"

	// ErrCodeNoFreeAddressesInSubnet for service response error code
	// "NoFreeAddressesInSubnet".
	//
	// Returned if IpAddress was not specified in the request and there are no free
	// IP addresses in the subnet.
	ErrCodeNoFreeAddressesInSubnet = "NoFreeAddressesInSubnet"

	// ErrCodeSecurityGroupLimitExceeded for service response error code
	// "SecurityGroupLimitExceeded".
	//
	// Returned if the size of SecurityGroups specified in the request is greater
	// than five.
	ErrCodeSecurityGroupLimitExceeded = "SecurityGroupLimitExceeded"

	// ErrCodeSecurityGroupNotFound for service response error code
	// "SecurityGroupNotFound".
	//
	// Returned if one of the specified security groups doesn't exist in the subnet's
	// VPC.
	ErrCodeSecurityGroupNotFound = "SecurityGroupNotFound"

	// ErrCodeSubnetNotFound for service response error code
	// "SubnetNotFound".
	//
	// Returned if there is no subnet with ID SubnetId provided in the request.
	ErrCodeSubnetNotFound = "SubnetNotFound"

	// ErrCodeThroughputLimitExceeded for service response error code
	// "ThroughputLimitExceeded".
	//
	// Returned if the throughput mode or amount of provisioned throughput can't
	// be changed because the throughput limit of 1024 MiB/s has been reached.
	ErrCodeThroughputLimitExceeded = "ThroughputLimitExceeded"

	// ErrCodeTooManyRequests for service response error code
	// "TooManyRequests".
	//
	// Returned if you don’t wait at least 24 hours before changing the throughput
	// mode, or decreasing the Provisioned Throughput value.
	ErrCodeTooManyRequests = "TooManyRequests"

	// ErrCodeUnsupportedAvailabilityZone for service response error code
	// "UnsupportedAvailabilityZone".
	ErrCodeUnsupportedAvailabilityZone = "UnsupportedAvailabilityZone"
)
