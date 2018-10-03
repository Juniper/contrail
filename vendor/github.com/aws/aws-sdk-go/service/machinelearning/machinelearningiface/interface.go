// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package machinelearningiface provides an interface to enable mocking the Amazon Machine Learning service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package machinelearningiface

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/machinelearning"
)

// MachineLearningAPI provides an interface to enable mocking the
// machinelearning.MachineLearning service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // Amazon Machine Learning.
//    func myFunc(svc machinelearningiface.MachineLearningAPI) bool {
//        // Make svc.AddTags request
//    }
//
//    func main() {
//        sess := session.New()
//        svc := machinelearning.New(sess)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockMachineLearningClient struct {
//        machinelearningiface.MachineLearningAPI
//    }
//    func (m *mockMachineLearningClient) AddTags(input *machinelearning.AddTagsInput) (*machinelearning.AddTagsOutput, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockMachineLearningClient{}
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
type MachineLearningAPI interface {
	AddTags(*machinelearning.AddTagsInput) (*machinelearning.AddTagsOutput, error)
	AddTagsWithContext(aws.Context, *machinelearning.AddTagsInput, ...request.Option) (*machinelearning.AddTagsOutput, error)
	AddTagsRequest(*machinelearning.AddTagsInput) (*request.Request, *machinelearning.AddTagsOutput)

	CreateBatchPrediction(*machinelearning.CreateBatchPredictionInput) (*machinelearning.CreateBatchPredictionOutput, error)
	CreateBatchPredictionWithContext(aws.Context, *machinelearning.CreateBatchPredictionInput, ...request.Option) (*machinelearning.CreateBatchPredictionOutput, error)
	CreateBatchPredictionRequest(*machinelearning.CreateBatchPredictionInput) (*request.Request, *machinelearning.CreateBatchPredictionOutput)

	CreateDataSourceFromRDS(*machinelearning.CreateDataSourceFromRDSInput) (*machinelearning.CreateDataSourceFromRDSOutput, error)
	CreateDataSourceFromRDSWithContext(aws.Context, *machinelearning.CreateDataSourceFromRDSInput, ...request.Option) (*machinelearning.CreateDataSourceFromRDSOutput, error)
	CreateDataSourceFromRDSRequest(*machinelearning.CreateDataSourceFromRDSInput) (*request.Request, *machinelearning.CreateDataSourceFromRDSOutput)

	CreateDataSourceFromRedshift(*machinelearning.CreateDataSourceFromRedshiftInput) (*machinelearning.CreateDataSourceFromRedshiftOutput, error)
	CreateDataSourceFromRedshiftWithContext(aws.Context, *machinelearning.CreateDataSourceFromRedshiftInput, ...request.Option) (*machinelearning.CreateDataSourceFromRedshiftOutput, error)
	CreateDataSourceFromRedshiftRequest(*machinelearning.CreateDataSourceFromRedshiftInput) (*request.Request, *machinelearning.CreateDataSourceFromRedshiftOutput)

	CreateDataSourceFromS3(*machinelearning.CreateDataSourceFromS3Input) (*machinelearning.CreateDataSourceFromS3Output, error)
	CreateDataSourceFromS3WithContext(aws.Context, *machinelearning.CreateDataSourceFromS3Input, ...request.Option) (*machinelearning.CreateDataSourceFromS3Output, error)
	CreateDataSourceFromS3Request(*machinelearning.CreateDataSourceFromS3Input) (*request.Request, *machinelearning.CreateDataSourceFromS3Output)

	CreateEvaluation(*machinelearning.CreateEvaluationInput) (*machinelearning.CreateEvaluationOutput, error)
	CreateEvaluationWithContext(aws.Context, *machinelearning.CreateEvaluationInput, ...request.Option) (*machinelearning.CreateEvaluationOutput, error)
	CreateEvaluationRequest(*machinelearning.CreateEvaluationInput) (*request.Request, *machinelearning.CreateEvaluationOutput)

	CreateMLModel(*machinelearning.CreateMLModelInput) (*machinelearning.CreateMLModelOutput, error)
	CreateMLModelWithContext(aws.Context, *machinelearning.CreateMLModelInput, ...request.Option) (*machinelearning.CreateMLModelOutput, error)
	CreateMLModelRequest(*machinelearning.CreateMLModelInput) (*request.Request, *machinelearning.CreateMLModelOutput)

	CreateRealtimeEndpoint(*machinelearning.CreateRealtimeEndpointInput) (*machinelearning.CreateRealtimeEndpointOutput, error)
	CreateRealtimeEndpointWithContext(aws.Context, *machinelearning.CreateRealtimeEndpointInput, ...request.Option) (*machinelearning.CreateRealtimeEndpointOutput, error)
	CreateRealtimeEndpointRequest(*machinelearning.CreateRealtimeEndpointInput) (*request.Request, *machinelearning.CreateRealtimeEndpointOutput)

	DeleteBatchPrediction(*machinelearning.DeleteBatchPredictionInput) (*machinelearning.DeleteBatchPredictionOutput, error)
	DeleteBatchPredictionWithContext(aws.Context, *machinelearning.DeleteBatchPredictionInput, ...request.Option) (*machinelearning.DeleteBatchPredictionOutput, error)
	DeleteBatchPredictionRequest(*machinelearning.DeleteBatchPredictionInput) (*request.Request, *machinelearning.DeleteBatchPredictionOutput)

	DeleteDataSource(*machinelearning.DeleteDataSourceInput) (*machinelearning.DeleteDataSourceOutput, error)
	DeleteDataSourceWithContext(aws.Context, *machinelearning.DeleteDataSourceInput, ...request.Option) (*machinelearning.DeleteDataSourceOutput, error)
	DeleteDataSourceRequest(*machinelearning.DeleteDataSourceInput) (*request.Request, *machinelearning.DeleteDataSourceOutput)

	DeleteEvaluation(*machinelearning.DeleteEvaluationInput) (*machinelearning.DeleteEvaluationOutput, error)
	DeleteEvaluationWithContext(aws.Context, *machinelearning.DeleteEvaluationInput, ...request.Option) (*machinelearning.DeleteEvaluationOutput, error)
	DeleteEvaluationRequest(*machinelearning.DeleteEvaluationInput) (*request.Request, *machinelearning.DeleteEvaluationOutput)

	DeleteMLModel(*machinelearning.DeleteMLModelInput) (*machinelearning.DeleteMLModelOutput, error)
	DeleteMLModelWithContext(aws.Context, *machinelearning.DeleteMLModelInput, ...request.Option) (*machinelearning.DeleteMLModelOutput, error)
	DeleteMLModelRequest(*machinelearning.DeleteMLModelInput) (*request.Request, *machinelearning.DeleteMLModelOutput)

	DeleteRealtimeEndpoint(*machinelearning.DeleteRealtimeEndpointInput) (*machinelearning.DeleteRealtimeEndpointOutput, error)
	DeleteRealtimeEndpointWithContext(aws.Context, *machinelearning.DeleteRealtimeEndpointInput, ...request.Option) (*machinelearning.DeleteRealtimeEndpointOutput, error)
	DeleteRealtimeEndpointRequest(*machinelearning.DeleteRealtimeEndpointInput) (*request.Request, *machinelearning.DeleteRealtimeEndpointOutput)

	DeleteTags(*machinelearning.DeleteTagsInput) (*machinelearning.DeleteTagsOutput, error)
	DeleteTagsWithContext(aws.Context, *machinelearning.DeleteTagsInput, ...request.Option) (*machinelearning.DeleteTagsOutput, error)
	DeleteTagsRequest(*machinelearning.DeleteTagsInput) (*request.Request, *machinelearning.DeleteTagsOutput)

	DescribeBatchPredictions(*machinelearning.DescribeBatchPredictionsInput) (*machinelearning.DescribeBatchPredictionsOutput, error)
	DescribeBatchPredictionsWithContext(aws.Context, *machinelearning.DescribeBatchPredictionsInput, ...request.Option) (*machinelearning.DescribeBatchPredictionsOutput, error)
	DescribeBatchPredictionsRequest(*machinelearning.DescribeBatchPredictionsInput) (*request.Request, *machinelearning.DescribeBatchPredictionsOutput)

	DescribeBatchPredictionsPages(*machinelearning.DescribeBatchPredictionsInput, func(*machinelearning.DescribeBatchPredictionsOutput, bool) bool) error
	DescribeBatchPredictionsPagesWithContext(aws.Context, *machinelearning.DescribeBatchPredictionsInput, func(*machinelearning.DescribeBatchPredictionsOutput, bool) bool, ...request.Option) error

	DescribeDataSources(*machinelearning.DescribeDataSourcesInput) (*machinelearning.DescribeDataSourcesOutput, error)
	DescribeDataSourcesWithContext(aws.Context, *machinelearning.DescribeDataSourcesInput, ...request.Option) (*machinelearning.DescribeDataSourcesOutput, error)
	DescribeDataSourcesRequest(*machinelearning.DescribeDataSourcesInput) (*request.Request, *machinelearning.DescribeDataSourcesOutput)

	DescribeDataSourcesPages(*machinelearning.DescribeDataSourcesInput, func(*machinelearning.DescribeDataSourcesOutput, bool) bool) error
	DescribeDataSourcesPagesWithContext(aws.Context, *machinelearning.DescribeDataSourcesInput, func(*machinelearning.DescribeDataSourcesOutput, bool) bool, ...request.Option) error

	DescribeEvaluations(*machinelearning.DescribeEvaluationsInput) (*machinelearning.DescribeEvaluationsOutput, error)
	DescribeEvaluationsWithContext(aws.Context, *machinelearning.DescribeEvaluationsInput, ...request.Option) (*machinelearning.DescribeEvaluationsOutput, error)
	DescribeEvaluationsRequest(*machinelearning.DescribeEvaluationsInput) (*request.Request, *machinelearning.DescribeEvaluationsOutput)

	DescribeEvaluationsPages(*machinelearning.DescribeEvaluationsInput, func(*machinelearning.DescribeEvaluationsOutput, bool) bool) error
	DescribeEvaluationsPagesWithContext(aws.Context, *machinelearning.DescribeEvaluationsInput, func(*machinelearning.DescribeEvaluationsOutput, bool) bool, ...request.Option) error

	DescribeMLModels(*machinelearning.DescribeMLModelsInput) (*machinelearning.DescribeMLModelsOutput, error)
	DescribeMLModelsWithContext(aws.Context, *machinelearning.DescribeMLModelsInput, ...request.Option) (*machinelearning.DescribeMLModelsOutput, error)
	DescribeMLModelsRequest(*machinelearning.DescribeMLModelsInput) (*request.Request, *machinelearning.DescribeMLModelsOutput)

	DescribeMLModelsPages(*machinelearning.DescribeMLModelsInput, func(*machinelearning.DescribeMLModelsOutput, bool) bool) error
	DescribeMLModelsPagesWithContext(aws.Context, *machinelearning.DescribeMLModelsInput, func(*machinelearning.DescribeMLModelsOutput, bool) bool, ...request.Option) error

	DescribeTags(*machinelearning.DescribeTagsInput) (*machinelearning.DescribeTagsOutput, error)
	DescribeTagsWithContext(aws.Context, *machinelearning.DescribeTagsInput, ...request.Option) (*machinelearning.DescribeTagsOutput, error)
	DescribeTagsRequest(*machinelearning.DescribeTagsInput) (*request.Request, *machinelearning.DescribeTagsOutput)

	GetBatchPrediction(*machinelearning.GetBatchPredictionInput) (*machinelearning.GetBatchPredictionOutput, error)
	GetBatchPredictionWithContext(aws.Context, *machinelearning.GetBatchPredictionInput, ...request.Option) (*machinelearning.GetBatchPredictionOutput, error)
	GetBatchPredictionRequest(*machinelearning.GetBatchPredictionInput) (*request.Request, *machinelearning.GetBatchPredictionOutput)

	GetDataSource(*machinelearning.GetDataSourceInput) (*machinelearning.GetDataSourceOutput, error)
	GetDataSourceWithContext(aws.Context, *machinelearning.GetDataSourceInput, ...request.Option) (*machinelearning.GetDataSourceOutput, error)
	GetDataSourceRequest(*machinelearning.GetDataSourceInput) (*request.Request, *machinelearning.GetDataSourceOutput)

	GetEvaluation(*machinelearning.GetEvaluationInput) (*machinelearning.GetEvaluationOutput, error)
	GetEvaluationWithContext(aws.Context, *machinelearning.GetEvaluationInput, ...request.Option) (*machinelearning.GetEvaluationOutput, error)
	GetEvaluationRequest(*machinelearning.GetEvaluationInput) (*request.Request, *machinelearning.GetEvaluationOutput)

	GetMLModel(*machinelearning.GetMLModelInput) (*machinelearning.GetMLModelOutput, error)
	GetMLModelWithContext(aws.Context, *machinelearning.GetMLModelInput, ...request.Option) (*machinelearning.GetMLModelOutput, error)
	GetMLModelRequest(*machinelearning.GetMLModelInput) (*request.Request, *machinelearning.GetMLModelOutput)

	Predict(*machinelearning.PredictInput) (*machinelearning.PredictOutput, error)
	PredictWithContext(aws.Context, *machinelearning.PredictInput, ...request.Option) (*machinelearning.PredictOutput, error)
	PredictRequest(*machinelearning.PredictInput) (*request.Request, *machinelearning.PredictOutput)

	UpdateBatchPrediction(*machinelearning.UpdateBatchPredictionInput) (*machinelearning.UpdateBatchPredictionOutput, error)
	UpdateBatchPredictionWithContext(aws.Context, *machinelearning.UpdateBatchPredictionInput, ...request.Option) (*machinelearning.UpdateBatchPredictionOutput, error)
	UpdateBatchPredictionRequest(*machinelearning.UpdateBatchPredictionInput) (*request.Request, *machinelearning.UpdateBatchPredictionOutput)

	UpdateDataSource(*machinelearning.UpdateDataSourceInput) (*machinelearning.UpdateDataSourceOutput, error)
	UpdateDataSourceWithContext(aws.Context, *machinelearning.UpdateDataSourceInput, ...request.Option) (*machinelearning.UpdateDataSourceOutput, error)
	UpdateDataSourceRequest(*machinelearning.UpdateDataSourceInput) (*request.Request, *machinelearning.UpdateDataSourceOutput)

	UpdateEvaluation(*machinelearning.UpdateEvaluationInput) (*machinelearning.UpdateEvaluationOutput, error)
	UpdateEvaluationWithContext(aws.Context, *machinelearning.UpdateEvaluationInput, ...request.Option) (*machinelearning.UpdateEvaluationOutput, error)
	UpdateEvaluationRequest(*machinelearning.UpdateEvaluationInput) (*request.Request, *machinelearning.UpdateEvaluationOutput)

	UpdateMLModel(*machinelearning.UpdateMLModelInput) (*machinelearning.UpdateMLModelOutput, error)
	UpdateMLModelWithContext(aws.Context, *machinelearning.UpdateMLModelInput, ...request.Option) (*machinelearning.UpdateMLModelOutput, error)
	UpdateMLModelRequest(*machinelearning.UpdateMLModelInput) (*request.Request, *machinelearning.UpdateMLModelOutput)

	WaitUntilBatchPredictionAvailable(*machinelearning.DescribeBatchPredictionsInput) error
	WaitUntilBatchPredictionAvailableWithContext(aws.Context, *machinelearning.DescribeBatchPredictionsInput, ...request.WaiterOption) error

	WaitUntilDataSourceAvailable(*machinelearning.DescribeDataSourcesInput) error
	WaitUntilDataSourceAvailableWithContext(aws.Context, *machinelearning.DescribeDataSourcesInput, ...request.WaiterOption) error

	WaitUntilEvaluationAvailable(*machinelearning.DescribeEvaluationsInput) error
	WaitUntilEvaluationAvailableWithContext(aws.Context, *machinelearning.DescribeEvaluationsInput, ...request.WaiterOption) error

	WaitUntilMLModelAvailable(*machinelearning.DescribeMLModelsInput) error
	WaitUntilMLModelAvailableWithContext(aws.Context, *machinelearning.DescribeMLModelsInput, ...request.WaiterOption) error
}

var _ MachineLearningAPI = (*machinelearning.MachineLearning)(nil)
