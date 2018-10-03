// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package rds

import (
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
)

// WaitUntilDBInstanceAvailable uses the Amazon RDS API operation
// DescribeDBInstances to wait for a condition to be met before returning.
// If the condition is not met within the max attempt window, an error will
// be returned.
func (c *RDS) WaitUntilDBInstanceAvailable(input *DescribeDBInstancesInput) error {
	return c.WaitUntilDBInstanceAvailableWithContext(aws.BackgroundContext(), input)
}

// WaitUntilDBInstanceAvailableWithContext is an extended version of WaitUntilDBInstanceAvailable.
// With the support for passing in a context and options to configure the
// Waiter and the underlying request options.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *RDS) WaitUntilDBInstanceAvailableWithContext(ctx aws.Context, input *DescribeDBInstancesInput, opts ...request.WaiterOption) error {
	w := request.Waiter{
		Name:        "WaitUntilDBInstanceAvailable",
		MaxAttempts: 60,
		Delay:       request.ConstantWaiterDelay(30 * time.Second),
		Acceptors: []request.WaiterAcceptor{
			{
				State:   request.SuccessWaiterState,
				Matcher: request.PathAllWaiterMatch, Argument: "DBInstances[].DBInstanceStatus",
				Expected: "available",
			},
			{
				State:   request.FailureWaiterState,
				Matcher: request.PathAnyWaiterMatch, Argument: "DBInstances[].DBInstanceStatus",
				Expected: "deleted",
			},
			{
				State:   request.FailureWaiterState,
				Matcher: request.PathAnyWaiterMatch, Argument: "DBInstances[].DBInstanceStatus",
				Expected: "deleting",
			},
			{
				State:   request.FailureWaiterState,
				Matcher: request.PathAnyWaiterMatch, Argument: "DBInstances[].DBInstanceStatus",
				Expected: "failed",
			},
			{
				State:   request.FailureWaiterState,
				Matcher: request.PathAnyWaiterMatch, Argument: "DBInstances[].DBInstanceStatus",
				Expected: "incompatible-restore",
			},
			{
				State:   request.FailureWaiterState,
				Matcher: request.PathAnyWaiterMatch, Argument: "DBInstances[].DBInstanceStatus",
				Expected: "incompatible-parameters",
			},
		},
		Logger: c.Config.Logger,
		NewRequest: func(opts []request.Option) (*request.Request, error) {
			var inCpy *DescribeDBInstancesInput
			if input != nil {
				tmp := *input
				inCpy = &tmp
			}
			req, _ := c.DescribeDBInstancesRequest(inCpy)
			req.SetContext(ctx)
			req.ApplyOptions(opts...)
			return req, nil
		},
	}
	w.ApplyOptions(opts...)

	return w.WaitWithContext(ctx)
}

// WaitUntilDBInstanceDeleted uses the Amazon RDS API operation
// DescribeDBInstances to wait for a condition to be met before returning.
// If the condition is not met within the max attempt window, an error will
// be returned.
func (c *RDS) WaitUntilDBInstanceDeleted(input *DescribeDBInstancesInput) error {
	return c.WaitUntilDBInstanceDeletedWithContext(aws.BackgroundContext(), input)
}

// WaitUntilDBInstanceDeletedWithContext is an extended version of WaitUntilDBInstanceDeleted.
// With the support for passing in a context and options to configure the
// Waiter and the underlying request options.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *RDS) WaitUntilDBInstanceDeletedWithContext(ctx aws.Context, input *DescribeDBInstancesInput, opts ...request.WaiterOption) error {
	w := request.Waiter{
		Name:        "WaitUntilDBInstanceDeleted",
		MaxAttempts: 60,
		Delay:       request.ConstantWaiterDelay(30 * time.Second),
		Acceptors: []request.WaiterAcceptor{
			{
				State:   request.SuccessWaiterState,
				Matcher: request.PathAllWaiterMatch, Argument: "DBInstances[].DBInstanceStatus",
				Expected: "deleted",
			},
			{
				State:    request.SuccessWaiterState,
				Matcher:  request.ErrorWaiterMatch,
				Expected: "DBInstanceNotFound",
			},
			{
				State:   request.FailureWaiterState,
				Matcher: request.PathAnyWaiterMatch, Argument: "DBInstances[].DBInstanceStatus",
				Expected: "creating",
			},
			{
				State:   request.FailureWaiterState,
				Matcher: request.PathAnyWaiterMatch, Argument: "DBInstances[].DBInstanceStatus",
				Expected: "modifying",
			},
			{
				State:   request.FailureWaiterState,
				Matcher: request.PathAnyWaiterMatch, Argument: "DBInstances[].DBInstanceStatus",
				Expected: "rebooting",
			},
			{
				State:   request.FailureWaiterState,
				Matcher: request.PathAnyWaiterMatch, Argument: "DBInstances[].DBInstanceStatus",
				Expected: "resetting-master-credentials",
			},
		},
		Logger: c.Config.Logger,
		NewRequest: func(opts []request.Option) (*request.Request, error) {
			var inCpy *DescribeDBInstancesInput
			if input != nil {
				tmp := *input
				inCpy = &tmp
			}
			req, _ := c.DescribeDBInstancesRequest(inCpy)
			req.SetContext(ctx)
			req.ApplyOptions(opts...)
			return req, nil
		},
	}
	w.ApplyOptions(opts...)

	return w.WaitWithContext(ctx)
}

// WaitUntilDBSnapshotAvailable uses the Amazon RDS API operation
// DescribeDBSnapshots to wait for a condition to be met before returning.
// If the condition is not met within the max attempt window, an error will
// be returned.
func (c *RDS) WaitUntilDBSnapshotAvailable(input *DescribeDBSnapshotsInput) error {
	return c.WaitUntilDBSnapshotAvailableWithContext(aws.BackgroundContext(), input)
}

// WaitUntilDBSnapshotAvailableWithContext is an extended version of WaitUntilDBSnapshotAvailable.
// With the support for passing in a context and options to configure the
// Waiter and the underlying request options.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *RDS) WaitUntilDBSnapshotAvailableWithContext(ctx aws.Context, input *DescribeDBSnapshotsInput, opts ...request.WaiterOption) error {
	w := request.Waiter{
		Name:        "WaitUntilDBSnapshotAvailable",
		MaxAttempts: 60,
		Delay:       request.ConstantWaiterDelay(30 * time.Second),
		Acceptors: []request.WaiterAcceptor{
			{
				State:   request.SuccessWaiterState,
				Matcher: request.PathAllWaiterMatch, Argument: "DBSnapshots[].Status",
				Expected: "available",
			},
			{
				State:   request.FailureWaiterState,
				Matcher: request.PathAnyWaiterMatch, Argument: "DBSnapshots[].Status",
				Expected: "deleted",
			},
			{
				State:   request.FailureWaiterState,
				Matcher: request.PathAnyWaiterMatch, Argument: "DBSnapshots[].Status",
				Expected: "deleting",
			},
			{
				State:   request.FailureWaiterState,
				Matcher: request.PathAnyWaiterMatch, Argument: "DBSnapshots[].Status",
				Expected: "failed",
			},
			{
				State:   request.FailureWaiterState,
				Matcher: request.PathAnyWaiterMatch, Argument: "DBSnapshots[].Status",
				Expected: "incompatible-restore",
			},
			{
				State:   request.FailureWaiterState,
				Matcher: request.PathAnyWaiterMatch, Argument: "DBSnapshots[].Status",
				Expected: "incompatible-parameters",
			},
		},
		Logger: c.Config.Logger,
		NewRequest: func(opts []request.Option) (*request.Request, error) {
			var inCpy *DescribeDBSnapshotsInput
			if input != nil {
				tmp := *input
				inCpy = &tmp
			}
			req, _ := c.DescribeDBSnapshotsRequest(inCpy)
			req.SetContext(ctx)
			req.ApplyOptions(opts...)
			return req, nil
		},
	}
	w.ApplyOptions(opts...)

	return w.WaitWithContext(ctx)
}

// WaitUntilDBSnapshotDeleted uses the Amazon RDS API operation
// DescribeDBSnapshots to wait for a condition to be met before returning.
// If the condition is not met within the max attempt window, an error will
// be returned.
func (c *RDS) WaitUntilDBSnapshotDeleted(input *DescribeDBSnapshotsInput) error {
	return c.WaitUntilDBSnapshotDeletedWithContext(aws.BackgroundContext(), input)
}

// WaitUntilDBSnapshotDeletedWithContext is an extended version of WaitUntilDBSnapshotDeleted.
// With the support for passing in a context and options to configure the
// Waiter and the underlying request options.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *RDS) WaitUntilDBSnapshotDeletedWithContext(ctx aws.Context, input *DescribeDBSnapshotsInput, opts ...request.WaiterOption) error {
	w := request.Waiter{
		Name:        "WaitUntilDBSnapshotDeleted",
		MaxAttempts: 60,
		Delay:       request.ConstantWaiterDelay(30 * time.Second),
		Acceptors: []request.WaiterAcceptor{
			{
				State:   request.SuccessWaiterState,
				Matcher: request.PathAllWaiterMatch, Argument: "DBSnapshots[].Status",
				Expected: "deleted",
			},
			{
				State:    request.SuccessWaiterState,
				Matcher:  request.ErrorWaiterMatch,
				Expected: "DBSnapshotNotFound",
			},
			{
				State:   request.FailureWaiterState,
				Matcher: request.PathAnyWaiterMatch, Argument: "DBSnapshots[].Status",
				Expected: "creating",
			},
			{
				State:   request.FailureWaiterState,
				Matcher: request.PathAnyWaiterMatch, Argument: "DBSnapshots[].Status",
				Expected: "modifying",
			},
			{
				State:   request.FailureWaiterState,
				Matcher: request.PathAnyWaiterMatch, Argument: "DBSnapshots[].Status",
				Expected: "rebooting",
			},
			{
				State:   request.FailureWaiterState,
				Matcher: request.PathAnyWaiterMatch, Argument: "DBSnapshots[].Status",
				Expected: "resetting-master-credentials",
			},
		},
		Logger: c.Config.Logger,
		NewRequest: func(opts []request.Option) (*request.Request, error) {
			var inCpy *DescribeDBSnapshotsInput
			if input != nil {
				tmp := *input
				inCpy = &tmp
			}
			req, _ := c.DescribeDBSnapshotsRequest(inCpy)
			req.SetContext(ctx)
			req.ApplyOptions(opts...)
			return req, nil
		},
	}
	w.ApplyOptions(opts...)

	return w.WaitWithContext(ctx)
}
