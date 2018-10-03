// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package sfniface provides an interface to enable mocking the AWS Step Functions service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package sfniface

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/sfn"
)

// SFNAPI provides an interface to enable mocking the
// sfn.SFN service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // AWS Step Functions.
//    func myFunc(svc sfniface.SFNAPI) bool {
//        // Make svc.CreateActivity request
//    }
//
//    func main() {
//        sess := session.New()
//        svc := sfn.New(sess)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockSFNClient struct {
//        sfniface.SFNAPI
//    }
//    func (m *mockSFNClient) CreateActivity(input *sfn.CreateActivityInput) (*sfn.CreateActivityOutput, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockSFNClient{}
//
//        myfunc(mockSvc)
//
//        // Verify myFunc's functionality
//    }
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters. Its suggested to use the pattern above for testing, or using
// tooling to generate mocks to satisfy the interfaces.
type SFNAPI interface {
	CreateActivity(*sfn.CreateActivityInput) (*sfn.CreateActivityOutput, error)
	CreateActivityWithContext(aws.Context, *sfn.CreateActivityInput, ...request.Option) (*sfn.CreateActivityOutput, error)
	CreateActivityRequest(*sfn.CreateActivityInput) (*request.Request, *sfn.CreateActivityOutput)

	CreateStateMachine(*sfn.CreateStateMachineInput) (*sfn.CreateStateMachineOutput, error)
	CreateStateMachineWithContext(aws.Context, *sfn.CreateStateMachineInput, ...request.Option) (*sfn.CreateStateMachineOutput, error)
	CreateStateMachineRequest(*sfn.CreateStateMachineInput) (*request.Request, *sfn.CreateStateMachineOutput)

	DeleteActivity(*sfn.DeleteActivityInput) (*sfn.DeleteActivityOutput, error)
	DeleteActivityWithContext(aws.Context, *sfn.DeleteActivityInput, ...request.Option) (*sfn.DeleteActivityOutput, error)
	DeleteActivityRequest(*sfn.DeleteActivityInput) (*request.Request, *sfn.DeleteActivityOutput)

	DeleteStateMachine(*sfn.DeleteStateMachineInput) (*sfn.DeleteStateMachineOutput, error)
	DeleteStateMachineWithContext(aws.Context, *sfn.DeleteStateMachineInput, ...request.Option) (*sfn.DeleteStateMachineOutput, error)
	DeleteStateMachineRequest(*sfn.DeleteStateMachineInput) (*request.Request, *sfn.DeleteStateMachineOutput)

	DescribeActivity(*sfn.DescribeActivityInput) (*sfn.DescribeActivityOutput, error)
	DescribeActivityWithContext(aws.Context, *sfn.DescribeActivityInput, ...request.Option) (*sfn.DescribeActivityOutput, error)
	DescribeActivityRequest(*sfn.DescribeActivityInput) (*request.Request, *sfn.DescribeActivityOutput)

	DescribeExecution(*sfn.DescribeExecutionInput) (*sfn.DescribeExecutionOutput, error)
	DescribeExecutionWithContext(aws.Context, *sfn.DescribeExecutionInput, ...request.Option) (*sfn.DescribeExecutionOutput, error)
	DescribeExecutionRequest(*sfn.DescribeExecutionInput) (*request.Request, *sfn.DescribeExecutionOutput)

	DescribeStateMachine(*sfn.DescribeStateMachineInput) (*sfn.DescribeStateMachineOutput, error)
	DescribeStateMachineWithContext(aws.Context, *sfn.DescribeStateMachineInput, ...request.Option) (*sfn.DescribeStateMachineOutput, error)
	DescribeStateMachineRequest(*sfn.DescribeStateMachineInput) (*request.Request, *sfn.DescribeStateMachineOutput)

	DescribeStateMachineForExecution(*sfn.DescribeStateMachineForExecutionInput) (*sfn.DescribeStateMachineForExecutionOutput, error)
	DescribeStateMachineForExecutionWithContext(aws.Context, *sfn.DescribeStateMachineForExecutionInput, ...request.Option) (*sfn.DescribeStateMachineForExecutionOutput, error)
	DescribeStateMachineForExecutionRequest(*sfn.DescribeStateMachineForExecutionInput) (*request.Request, *sfn.DescribeStateMachineForExecutionOutput)

	GetActivityTask(*sfn.GetActivityTaskInput) (*sfn.GetActivityTaskOutput, error)
	GetActivityTaskWithContext(aws.Context, *sfn.GetActivityTaskInput, ...request.Option) (*sfn.GetActivityTaskOutput, error)
	GetActivityTaskRequest(*sfn.GetActivityTaskInput) (*request.Request, *sfn.GetActivityTaskOutput)

	GetExecutionHistory(*sfn.GetExecutionHistoryInput) (*sfn.GetExecutionHistoryOutput, error)
	GetExecutionHistoryWithContext(aws.Context, *sfn.GetExecutionHistoryInput, ...request.Option) (*sfn.GetExecutionHistoryOutput, error)
	GetExecutionHistoryRequest(*sfn.GetExecutionHistoryInput) (*request.Request, *sfn.GetExecutionHistoryOutput)

	GetExecutionHistoryPages(*sfn.GetExecutionHistoryInput, func(*sfn.GetExecutionHistoryOutput, bool) bool) error
	GetExecutionHistoryPagesWithContext(aws.Context, *sfn.GetExecutionHistoryInput, func(*sfn.GetExecutionHistoryOutput, bool) bool, ...request.Option) error

	ListActivities(*sfn.ListActivitiesInput) (*sfn.ListActivitiesOutput, error)
	ListActivitiesWithContext(aws.Context, *sfn.ListActivitiesInput, ...request.Option) (*sfn.ListActivitiesOutput, error)
	ListActivitiesRequest(*sfn.ListActivitiesInput) (*request.Request, *sfn.ListActivitiesOutput)

	ListActivitiesPages(*sfn.ListActivitiesInput, func(*sfn.ListActivitiesOutput, bool) bool) error
	ListActivitiesPagesWithContext(aws.Context, *sfn.ListActivitiesInput, func(*sfn.ListActivitiesOutput, bool) bool, ...request.Option) error

	ListExecutions(*sfn.ListExecutionsInput) (*sfn.ListExecutionsOutput, error)
	ListExecutionsWithContext(aws.Context, *sfn.ListExecutionsInput, ...request.Option) (*sfn.ListExecutionsOutput, error)
	ListExecutionsRequest(*sfn.ListExecutionsInput) (*request.Request, *sfn.ListExecutionsOutput)

	ListExecutionsPages(*sfn.ListExecutionsInput, func(*sfn.ListExecutionsOutput, bool) bool) error
	ListExecutionsPagesWithContext(aws.Context, *sfn.ListExecutionsInput, func(*sfn.ListExecutionsOutput, bool) bool, ...request.Option) error

	ListStateMachines(*sfn.ListStateMachinesInput) (*sfn.ListStateMachinesOutput, error)
	ListStateMachinesWithContext(aws.Context, *sfn.ListStateMachinesInput, ...request.Option) (*sfn.ListStateMachinesOutput, error)
	ListStateMachinesRequest(*sfn.ListStateMachinesInput) (*request.Request, *sfn.ListStateMachinesOutput)

	ListStateMachinesPages(*sfn.ListStateMachinesInput, func(*sfn.ListStateMachinesOutput, bool) bool) error
	ListStateMachinesPagesWithContext(aws.Context, *sfn.ListStateMachinesInput, func(*sfn.ListStateMachinesOutput, bool) bool, ...request.Option) error

	SendTaskFailure(*sfn.SendTaskFailureInput) (*sfn.SendTaskFailureOutput, error)
	SendTaskFailureWithContext(aws.Context, *sfn.SendTaskFailureInput, ...request.Option) (*sfn.SendTaskFailureOutput, error)
	SendTaskFailureRequest(*sfn.SendTaskFailureInput) (*request.Request, *sfn.SendTaskFailureOutput)

	SendTaskHeartbeat(*sfn.SendTaskHeartbeatInput) (*sfn.SendTaskHeartbeatOutput, error)
	SendTaskHeartbeatWithContext(aws.Context, *sfn.SendTaskHeartbeatInput, ...request.Option) (*sfn.SendTaskHeartbeatOutput, error)
	SendTaskHeartbeatRequest(*sfn.SendTaskHeartbeatInput) (*request.Request, *sfn.SendTaskHeartbeatOutput)

	SendTaskSuccess(*sfn.SendTaskSuccessInput) (*sfn.SendTaskSuccessOutput, error)
	SendTaskSuccessWithContext(aws.Context, *sfn.SendTaskSuccessInput, ...request.Option) (*sfn.SendTaskSuccessOutput, error)
	SendTaskSuccessRequest(*sfn.SendTaskSuccessInput) (*request.Request, *sfn.SendTaskSuccessOutput)

	StartExecution(*sfn.StartExecutionInput) (*sfn.StartExecutionOutput, error)
	StartExecutionWithContext(aws.Context, *sfn.StartExecutionInput, ...request.Option) (*sfn.StartExecutionOutput, error)
	StartExecutionRequest(*sfn.StartExecutionInput) (*request.Request, *sfn.StartExecutionOutput)

	StopExecution(*sfn.StopExecutionInput) (*sfn.StopExecutionOutput, error)
	StopExecutionWithContext(aws.Context, *sfn.StopExecutionInput, ...request.Option) (*sfn.StopExecutionOutput, error)
	StopExecutionRequest(*sfn.StopExecutionInput) (*request.Request, *sfn.StopExecutionOutput)

	UpdateStateMachine(*sfn.UpdateStateMachineInput) (*sfn.UpdateStateMachineOutput, error)
	UpdateStateMachineWithContext(aws.Context, *sfn.UpdateStateMachineInput, ...request.Option) (*sfn.UpdateStateMachineOutput, error)
	UpdateStateMachineRequest(*sfn.UpdateStateMachineInput) (*request.Request, *sfn.UpdateStateMachineOutput)
}

var _ SFNAPI = (*sfn.SFN)(nil)
