// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package daxiface provides an interface to enable mocking the Amazon DynamoDB Accelerator (DAX) service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package daxiface

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/dax"
)

// DAXAPI provides an interface to enable mocking the
// dax.DAX service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // Amazon DynamoDB Accelerator (DAX).
//    func myFunc(svc daxiface.DAXAPI) bool {
//        // Make svc.CreateCluster request
//    }
//
//    func main() {
//        sess := session.New()
//        svc := dax.New(sess)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockDAXClient struct {
//        daxiface.DAXAPI
//    }
//    func (m *mockDAXClient) CreateCluster(input *dax.CreateClusterInput) (*dax.CreateClusterOutput, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockDAXClient{}
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
type DAXAPI interface {
	CreateCluster(*dax.CreateClusterInput) (*dax.CreateClusterOutput, error)
	CreateClusterWithContext(aws.Context, *dax.CreateClusterInput, ...request.Option) (*dax.CreateClusterOutput, error)
	CreateClusterRequest(*dax.CreateClusterInput) (*request.Request, *dax.CreateClusterOutput)

	CreateParameterGroup(*dax.CreateParameterGroupInput) (*dax.CreateParameterGroupOutput, error)
	CreateParameterGroupWithContext(aws.Context, *dax.CreateParameterGroupInput, ...request.Option) (*dax.CreateParameterGroupOutput, error)
	CreateParameterGroupRequest(*dax.CreateParameterGroupInput) (*request.Request, *dax.CreateParameterGroupOutput)

	CreateSubnetGroup(*dax.CreateSubnetGroupInput) (*dax.CreateSubnetGroupOutput, error)
	CreateSubnetGroupWithContext(aws.Context, *dax.CreateSubnetGroupInput, ...request.Option) (*dax.CreateSubnetGroupOutput, error)
	CreateSubnetGroupRequest(*dax.CreateSubnetGroupInput) (*request.Request, *dax.CreateSubnetGroupOutput)

	DecreaseReplicationFactor(*dax.DecreaseReplicationFactorInput) (*dax.DecreaseReplicationFactorOutput, error)
	DecreaseReplicationFactorWithContext(aws.Context, *dax.DecreaseReplicationFactorInput, ...request.Option) (*dax.DecreaseReplicationFactorOutput, error)
	DecreaseReplicationFactorRequest(*dax.DecreaseReplicationFactorInput) (*request.Request, *dax.DecreaseReplicationFactorOutput)

	DeleteCluster(*dax.DeleteClusterInput) (*dax.DeleteClusterOutput, error)
	DeleteClusterWithContext(aws.Context, *dax.DeleteClusterInput, ...request.Option) (*dax.DeleteClusterOutput, error)
	DeleteClusterRequest(*dax.DeleteClusterInput) (*request.Request, *dax.DeleteClusterOutput)

	DeleteParameterGroup(*dax.DeleteParameterGroupInput) (*dax.DeleteParameterGroupOutput, error)
	DeleteParameterGroupWithContext(aws.Context, *dax.DeleteParameterGroupInput, ...request.Option) (*dax.DeleteParameterGroupOutput, error)
	DeleteParameterGroupRequest(*dax.DeleteParameterGroupInput) (*request.Request, *dax.DeleteParameterGroupOutput)

	DeleteSubnetGroup(*dax.DeleteSubnetGroupInput) (*dax.DeleteSubnetGroupOutput, error)
	DeleteSubnetGroupWithContext(aws.Context, *dax.DeleteSubnetGroupInput, ...request.Option) (*dax.DeleteSubnetGroupOutput, error)
	DeleteSubnetGroupRequest(*dax.DeleteSubnetGroupInput) (*request.Request, *dax.DeleteSubnetGroupOutput)

	DescribeClusters(*dax.DescribeClustersInput) (*dax.DescribeClustersOutput, error)
	DescribeClustersWithContext(aws.Context, *dax.DescribeClustersInput, ...request.Option) (*dax.DescribeClustersOutput, error)
	DescribeClustersRequest(*dax.DescribeClustersInput) (*request.Request, *dax.DescribeClustersOutput)

	DescribeDefaultParameters(*dax.DescribeDefaultParametersInput) (*dax.DescribeDefaultParametersOutput, error)
	DescribeDefaultParametersWithContext(aws.Context, *dax.DescribeDefaultParametersInput, ...request.Option) (*dax.DescribeDefaultParametersOutput, error)
	DescribeDefaultParametersRequest(*dax.DescribeDefaultParametersInput) (*request.Request, *dax.DescribeDefaultParametersOutput)

	DescribeEvents(*dax.DescribeEventsInput) (*dax.DescribeEventsOutput, error)
	DescribeEventsWithContext(aws.Context, *dax.DescribeEventsInput, ...request.Option) (*dax.DescribeEventsOutput, error)
	DescribeEventsRequest(*dax.DescribeEventsInput) (*request.Request, *dax.DescribeEventsOutput)

	DescribeParameterGroups(*dax.DescribeParameterGroupsInput) (*dax.DescribeParameterGroupsOutput, error)
	DescribeParameterGroupsWithContext(aws.Context, *dax.DescribeParameterGroupsInput, ...request.Option) (*dax.DescribeParameterGroupsOutput, error)
	DescribeParameterGroupsRequest(*dax.DescribeParameterGroupsInput) (*request.Request, *dax.DescribeParameterGroupsOutput)

	DescribeParameters(*dax.DescribeParametersInput) (*dax.DescribeParametersOutput, error)
	DescribeParametersWithContext(aws.Context, *dax.DescribeParametersInput, ...request.Option) (*dax.DescribeParametersOutput, error)
	DescribeParametersRequest(*dax.DescribeParametersInput) (*request.Request, *dax.DescribeParametersOutput)

	DescribeSubnetGroups(*dax.DescribeSubnetGroupsInput) (*dax.DescribeSubnetGroupsOutput, error)
	DescribeSubnetGroupsWithContext(aws.Context, *dax.DescribeSubnetGroupsInput, ...request.Option) (*dax.DescribeSubnetGroupsOutput, error)
	DescribeSubnetGroupsRequest(*dax.DescribeSubnetGroupsInput) (*request.Request, *dax.DescribeSubnetGroupsOutput)

	IncreaseReplicationFactor(*dax.IncreaseReplicationFactorInput) (*dax.IncreaseReplicationFactorOutput, error)
	IncreaseReplicationFactorWithContext(aws.Context, *dax.IncreaseReplicationFactorInput, ...request.Option) (*dax.IncreaseReplicationFactorOutput, error)
	IncreaseReplicationFactorRequest(*dax.IncreaseReplicationFactorInput) (*request.Request, *dax.IncreaseReplicationFactorOutput)

	ListTags(*dax.ListTagsInput) (*dax.ListTagsOutput, error)
	ListTagsWithContext(aws.Context, *dax.ListTagsInput, ...request.Option) (*dax.ListTagsOutput, error)
	ListTagsRequest(*dax.ListTagsInput) (*request.Request, *dax.ListTagsOutput)

	RebootNode(*dax.RebootNodeInput) (*dax.RebootNodeOutput, error)
	RebootNodeWithContext(aws.Context, *dax.RebootNodeInput, ...request.Option) (*dax.RebootNodeOutput, error)
	RebootNodeRequest(*dax.RebootNodeInput) (*request.Request, *dax.RebootNodeOutput)

	TagResource(*dax.TagResourceInput) (*dax.TagResourceOutput, error)
	TagResourceWithContext(aws.Context, *dax.TagResourceInput, ...request.Option) (*dax.TagResourceOutput, error)
	TagResourceRequest(*dax.TagResourceInput) (*request.Request, *dax.TagResourceOutput)

	UntagResource(*dax.UntagResourceInput) (*dax.UntagResourceOutput, error)
	UntagResourceWithContext(aws.Context, *dax.UntagResourceInput, ...request.Option) (*dax.UntagResourceOutput, error)
	UntagResourceRequest(*dax.UntagResourceInput) (*request.Request, *dax.UntagResourceOutput)

	UpdateCluster(*dax.UpdateClusterInput) (*dax.UpdateClusterOutput, error)
	UpdateClusterWithContext(aws.Context, *dax.UpdateClusterInput, ...request.Option) (*dax.UpdateClusterOutput, error)
	UpdateClusterRequest(*dax.UpdateClusterInput) (*request.Request, *dax.UpdateClusterOutput)

	UpdateParameterGroup(*dax.UpdateParameterGroupInput) (*dax.UpdateParameterGroupOutput, error)
	UpdateParameterGroupWithContext(aws.Context, *dax.UpdateParameterGroupInput, ...request.Option) (*dax.UpdateParameterGroupOutput, error)
	UpdateParameterGroupRequest(*dax.UpdateParameterGroupInput) (*request.Request, *dax.UpdateParameterGroupOutput)

	UpdateSubnetGroup(*dax.UpdateSubnetGroupInput) (*dax.UpdateSubnetGroupOutput, error)
	UpdateSubnetGroupWithContext(aws.Context, *dax.UpdateSubnetGroupInput, ...request.Option) (*dax.UpdateSubnetGroupOutput, error)
	UpdateSubnetGroupRequest(*dax.UpdateSubnetGroupInput) (*request.Request, *dax.UpdateSubnetGroupOutput)
}

var _ DAXAPI = (*dax.DAX)(nil)
