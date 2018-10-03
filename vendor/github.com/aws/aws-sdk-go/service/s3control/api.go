// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package s3control

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awsutil"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/private/protocol"
	"github.com/aws/aws-sdk-go/private/protocol/restxml"
)

const opDeletePublicAccessBlock = "DeletePublicAccessBlock"

// DeletePublicAccessBlockRequest generates a "aws/request.Request" representing the
// client's request for the DeletePublicAccessBlock operation. The "output" return
// value will be populated with the request's response once the request completes
// successfully.
//
// Use "Send" method on the returned Request to send the API call to the service.
// the "output" return value is not valid until after Send returns without error.
//
// See DeletePublicAccessBlock for more information on using the DeletePublicAccessBlock
// API call, and error handling.
//
// This method is useful when you want to inject custom logic or configuration
// into the SDK's request lifecycle. Such as custom headers, or retry logic.
//
//
//    // Example sending a request using the DeletePublicAccessBlockRequest method.
//    req, resp := client.DeletePublicAccessBlockRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
//
// See also, https://docs.aws.amazon.com/goto/WebAPI/s3control-2018-08-20/DeletePublicAccessBlock
func (c *S3Control) DeletePublicAccessBlockRequest(input *DeletePublicAccessBlockInput) (req *request.Request, output *DeletePublicAccessBlockOutput) {
	op := &request.Operation{
		Name:       opDeletePublicAccessBlock,
		HTTPMethod: "DELETE",
		HTTPPath:   "/v20180820/configuration/publicAccessBlock",
	}

	if input == nil {
		input = &DeletePublicAccessBlockInput{}
	}

	output = &DeletePublicAccessBlockOutput{}
	req = c.newRequest(op, input, output)
	req.Handlers.Unmarshal.Swap(restxml.UnmarshalHandler.Name, protocol.UnmarshalDiscardBodyHandler)
	req.Handlers.Build.PushBackNamed(buildPrefixHostHandler("AccountID", aws.StringValue(input.AccountId)))
	req.Handlers.Build.PushBackNamed(buildRemoveHeaderHandler("X-Amz-Account-Id"))
	return
}

// DeletePublicAccessBlock API operation for AWS S3 Control.
//
// Removes the Public Access Block configuration for an Amazon Web Services
// account.
//
// Returns awserr.Error for service API and SDK errors. Use runtime type assertions
// with awserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the AWS API reference guide for AWS S3 Control's
// API operation DeletePublicAccessBlock for usage and error information.
// See also, https://docs.aws.amazon.com/goto/WebAPI/s3control-2018-08-20/DeletePublicAccessBlock
func (c *S3Control) DeletePublicAccessBlock(input *DeletePublicAccessBlockInput) (*DeletePublicAccessBlockOutput, error) {
	req, out := c.DeletePublicAccessBlockRequest(input)
	return out, req.Send()
}

// DeletePublicAccessBlockWithContext is the same as DeletePublicAccessBlock with the addition of
// the ability to pass a context and additional request options.
//
// See DeletePublicAccessBlock for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *S3Control) DeletePublicAccessBlockWithContext(ctx aws.Context, input *DeletePublicAccessBlockInput, opts ...request.Option) (*DeletePublicAccessBlockOutput, error) {
	req, out := c.DeletePublicAccessBlockRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opGetPublicAccessBlock = "GetPublicAccessBlock"

// GetPublicAccessBlockRequest generates a "aws/request.Request" representing the
// client's request for the GetPublicAccessBlock operation. The "output" return
// value will be populated with the request's response once the request completes
// successfully.
//
// Use "Send" method on the returned Request to send the API call to the service.
// the "output" return value is not valid until after Send returns without error.
//
// See GetPublicAccessBlock for more information on using the GetPublicAccessBlock
// API call, and error handling.
//
// This method is useful when you want to inject custom logic or configuration
// into the SDK's request lifecycle. Such as custom headers, or retry logic.
//
//
//    // Example sending a request using the GetPublicAccessBlockRequest method.
//    req, resp := client.GetPublicAccessBlockRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
//
// See also, https://docs.aws.amazon.com/goto/WebAPI/s3control-2018-08-20/GetPublicAccessBlock
func (c *S3Control) GetPublicAccessBlockRequest(input *GetPublicAccessBlockInput) (req *request.Request, output *GetPublicAccessBlockOutput) {
	op := &request.Operation{
		Name:       opGetPublicAccessBlock,
		HTTPMethod: "GET",
		HTTPPath:   "/v20180820/configuration/publicAccessBlock",
	}

	if input == nil {
		input = &GetPublicAccessBlockInput{}
	}

	output = &GetPublicAccessBlockOutput{}
	req = c.newRequest(op, input, output)
	req.Handlers.Build.PushBackNamed(buildPrefixHostHandler("AccountID", aws.StringValue(input.AccountId)))
	req.Handlers.Build.PushBackNamed(buildRemoveHeaderHandler("X-Amz-Account-Id"))
	return
}

// GetPublicAccessBlock API operation for AWS S3 Control.
//
// Retrieves the Public Access Block configuration for an Amazon Web Services
// account.
//
// Returns awserr.Error for service API and SDK errors. Use runtime type assertions
// with awserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the AWS API reference guide for AWS S3 Control's
// API operation GetPublicAccessBlock for usage and error information.
//
// Returned Error Codes:
//   * ErrCodeNoSuchPublicAccessBlockConfiguration "NoSuchPublicAccessBlockConfiguration"
//   This exception is thrown if a GetPublicAccessBlock request is made against
//   an account that does not have a PublicAccessBlockConfiguration set.
//
// See also, https://docs.aws.amazon.com/goto/WebAPI/s3control-2018-08-20/GetPublicAccessBlock
func (c *S3Control) GetPublicAccessBlock(input *GetPublicAccessBlockInput) (*GetPublicAccessBlockOutput, error) {
	req, out := c.GetPublicAccessBlockRequest(input)
	return out, req.Send()
}

// GetPublicAccessBlockWithContext is the same as GetPublicAccessBlock with the addition of
// the ability to pass a context and additional request options.
//
// See GetPublicAccessBlock for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *S3Control) GetPublicAccessBlockWithContext(ctx aws.Context, input *GetPublicAccessBlockInput, opts ...request.Option) (*GetPublicAccessBlockOutput, error) {
	req, out := c.GetPublicAccessBlockRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opPutPublicAccessBlock = "PutPublicAccessBlock"

// PutPublicAccessBlockRequest generates a "aws/request.Request" representing the
// client's request for the PutPublicAccessBlock operation. The "output" return
// value will be populated with the request's response once the request completes
// successfully.
//
// Use "Send" method on the returned Request to send the API call to the service.
// the "output" return value is not valid until after Send returns without error.
//
// See PutPublicAccessBlock for more information on using the PutPublicAccessBlock
// API call, and error handling.
//
// This method is useful when you want to inject custom logic or configuration
// into the SDK's request lifecycle. Such as custom headers, or retry logic.
//
//
//    // Example sending a request using the PutPublicAccessBlockRequest method.
//    req, resp := client.PutPublicAccessBlockRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
//
// See also, https://docs.aws.amazon.com/goto/WebAPI/s3control-2018-08-20/PutPublicAccessBlock
func (c *S3Control) PutPublicAccessBlockRequest(input *PutPublicAccessBlockInput) (req *request.Request, output *PutPublicAccessBlockOutput) {
	op := &request.Operation{
		Name:       opPutPublicAccessBlock,
		HTTPMethod: "PUT",
		HTTPPath:   "/v20180820/configuration/publicAccessBlock",
	}

	if input == nil {
		input = &PutPublicAccessBlockInput{}
	}

	output = &PutPublicAccessBlockOutput{}
	req = c.newRequest(op, input, output)
	req.Handlers.Unmarshal.Swap(restxml.UnmarshalHandler.Name, protocol.UnmarshalDiscardBodyHandler)
	req.Handlers.Build.PushBackNamed(buildPrefixHostHandler("AccountID", aws.StringValue(input.AccountId)))
	req.Handlers.Build.PushBackNamed(buildRemoveHeaderHandler("X-Amz-Account-Id"))
	return
}

// PutPublicAccessBlock API operation for AWS S3 Control.
//
// Creates or modifies the Public Access Block configuration for an Amazon Web
// Services account.
//
// Returns awserr.Error for service API and SDK errors. Use runtime type assertions
// with awserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the AWS API reference guide for AWS S3 Control's
// API operation PutPublicAccessBlock for usage and error information.
// See also, https://docs.aws.amazon.com/goto/WebAPI/s3control-2018-08-20/PutPublicAccessBlock
func (c *S3Control) PutPublicAccessBlock(input *PutPublicAccessBlockInput) (*PutPublicAccessBlockOutput, error) {
	req, out := c.PutPublicAccessBlockRequest(input)
	return out, req.Send()
}

// PutPublicAccessBlockWithContext is the same as PutPublicAccessBlock with the addition of
// the ability to pass a context and additional request options.
//
// See PutPublicAccessBlock for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *S3Control) PutPublicAccessBlockWithContext(ctx aws.Context, input *PutPublicAccessBlockInput, opts ...request.Option) (*PutPublicAccessBlockOutput, error) {
	req, out := c.PutPublicAccessBlockRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type DeletePublicAccessBlockInput struct {
	_ struct{} `type:"structure"`

	// The Account ID for the Amazon Web Services account whose Public Access Block
	// configuration you want to remove.
	//
	// AccountId is a required field
	AccountId *string `location:"header" locationName:"x-amz-account-id" type:"string" required:"true"`
}

// String returns the string representation
func (s DeletePublicAccessBlockInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s DeletePublicAccessBlockInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeletePublicAccessBlockInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "DeletePublicAccessBlockInput"}
	if s.AccountId == nil {
		invalidParams.Add(request.NewErrParamRequired("AccountId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetAccountId sets the AccountId field's value.
func (s *DeletePublicAccessBlockInput) SetAccountId(v string) *DeletePublicAccessBlockInput {
	s.AccountId = &v
	return s
}

type DeletePublicAccessBlockOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeletePublicAccessBlockOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s DeletePublicAccessBlockOutput) GoString() string {
	return s.String()
}

type GetPublicAccessBlockInput struct {
	_ struct{} `type:"structure"`

	// The Account ID for the Amazon Web Services account whose Public Access Block
	// configuration you want to retrieve.
	//
	// AccountId is a required field
	AccountId *string `location:"header" locationName:"x-amz-account-id" type:"string" required:"true"`
}

// String returns the string representation
func (s GetPublicAccessBlockInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s GetPublicAccessBlockInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetPublicAccessBlockInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "GetPublicAccessBlockInput"}
	if s.AccountId == nil {
		invalidParams.Add(request.NewErrParamRequired("AccountId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetAccountId sets the AccountId field's value.
func (s *GetPublicAccessBlockInput) SetAccountId(v string) *GetPublicAccessBlockInput {
	s.AccountId = &v
	return s
}

type GetPublicAccessBlockOutput struct {
	_ struct{} `type:"structure" payload:"PublicAccessBlockConfiguration"`

	// The Public Access Block configuration currently in effect for this Amazon
	// Web Services account.
	PublicAccessBlockConfiguration *PublicAccessBlockConfiguration `type:"structure"`
}

// String returns the string representation
func (s GetPublicAccessBlockOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s GetPublicAccessBlockOutput) GoString() string {
	return s.String()
}

// SetPublicAccessBlockConfiguration sets the PublicAccessBlockConfiguration field's value.
func (s *GetPublicAccessBlockOutput) SetPublicAccessBlockConfiguration(v *PublicAccessBlockConfiguration) *GetPublicAccessBlockOutput {
	s.PublicAccessBlockConfiguration = v
	return s
}

// The container element for all Public Access Block configuration options.
// You can enable the configuration options in any combination.
//
// Amazon S3 considers a bucket policy public unless at least one of the following
// conditions is true:
//
// The policy limits access to a set of CIDRs using aws:SourceIp. For more information
// on CIDR, see http://www.rfc-editor.org/rfc/rfc4632.txt (http://www.rfc-editor.org/rfc/rfc4632.txt)
//
// The policy grants permissions, not including any "bad actions," to one of
// the following:
//
// A fixed AWS principal, user, role, or service principal
//
// A fixed aws:SourceArn
//
// A fixed aws:SourceVpc
//
// A fixed aws:SourceVpce
//
// A fixed aws:SourceOwner
//
// A fixed aws:SourceAccount
//
// A fixed value of s3:x-amz-server-side-encryption-aws-kms-key-id
//
// A fixed value of aws:userid outside the pattern "AROLEID:*"
//
// "Bad actions" are those that could expose the data inside a bucket to reads
// or writes by the public. These actions are s3:Get*, s3:List*, s3:AbortMultipartUpload,
// s3:Delete*, s3:Put*, and s3:RestoreObject.
//
// The star notation for bad actions indicates that all matching operations
// are considered bad actions. For example, because s3:Get* is a bad action,
// s3:GetObject, s3:GetObjectVersion, and s3:GetObjectAcl are all bad actions.
type PublicAccessBlockConfiguration struct {
	_ struct{} `type:"structure"`

	// Specifies whether Amazon S3 should block public ACLs for buckets in this
	// account. Setting this element to TRUE causes the following behavior:
	//
	//    * PUT Bucket acl and PUT Object acl calls will fail if the specified ACL
	//    allows public access.
	//
	//    * PUT Object calls will fail if the request includes an object ACL.
	//
	// Note that enabling this setting doesn't affect existing policies or ACLs.
	BlockPublicAcls *bool `locationName:"BlockPublicAcls" type:"boolean"`

	// Specifies whether Amazon S3 should block public bucket policies for buckets
	// in this account. Setting this element to TRUE causes Amazon S3 to reject
	// calls to PUT Bucket policy if the specified bucket policy allows public access.
	//
	// Note that enabling this setting doesn't affect existing bucket policies.
	BlockPublicPolicy *bool `locationName:"BlockPublicPolicy" type:"boolean"`

	// Specifies whether Amazon S3 should ignore public ACLs for buckets in this
	// account. Setting this element to TRUE causes Amazon S3 to ignore all public
	// ACLs on buckets in this account and any objects that they contain.
	//
	// Note that enabling this setting doesn't affect the persistence of any existing
	// ACLs and doesn't prevent new public ACLs from being set.
	IgnorePublicAcls *bool `locationName:"IgnorePublicAcls" type:"boolean"`

	// Specifies whether Amazon S3 should restrict public bucket policies for buckets
	// in this account. If this element is set to TRUE, then only the bucket owner
	// and AWS Services can access buckets with public policies.
	//
	// Note that enabling this setting doesn't affect previously stored bucket policies,
	// except that public and cross-account access within any public bucket policy,
	// including non-public delegation to specific accounts, is blocked.
	RestrictPublicBuckets *bool `locationName:"RestrictPublicBuckets" type:"boolean"`
}

// String returns the string representation
func (s PublicAccessBlockConfiguration) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s PublicAccessBlockConfiguration) GoString() string {
	return s.String()
}

// SetBlockPublicAcls sets the BlockPublicAcls field's value.
func (s *PublicAccessBlockConfiguration) SetBlockPublicAcls(v bool) *PublicAccessBlockConfiguration {
	s.BlockPublicAcls = &v
	return s
}

// SetBlockPublicPolicy sets the BlockPublicPolicy field's value.
func (s *PublicAccessBlockConfiguration) SetBlockPublicPolicy(v bool) *PublicAccessBlockConfiguration {
	s.BlockPublicPolicy = &v
	return s
}

// SetIgnorePublicAcls sets the IgnorePublicAcls field's value.
func (s *PublicAccessBlockConfiguration) SetIgnorePublicAcls(v bool) *PublicAccessBlockConfiguration {
	s.IgnorePublicAcls = &v
	return s
}

// SetRestrictPublicBuckets sets the RestrictPublicBuckets field's value.
func (s *PublicAccessBlockConfiguration) SetRestrictPublicBuckets(v bool) *PublicAccessBlockConfiguration {
	s.RestrictPublicBuckets = &v
	return s
}

type PutPublicAccessBlockInput struct {
	_ struct{} `type:"structure" payload:"PublicAccessBlockConfiguration"`

	// The Account ID for the Amazon Web Services account whose Public Access Block
	// configuration you want to set.
	//
	// AccountId is a required field
	AccountId *string `location:"header" locationName:"x-amz-account-id" type:"string" required:"true"`

	// The Public Access Block configuration that you want to apply to this Amazon
	// Web Services account.
	//
	// PublicAccessBlockConfiguration is a required field
	PublicAccessBlockConfiguration *PublicAccessBlockConfiguration `locationName:"PublicAccessBlockConfiguration" type:"structure" required:"true" xmlURI:"http://awss3control.amazonaws.com/doc/2018-08-20/"`
}

// String returns the string representation
func (s PutPublicAccessBlockInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s PutPublicAccessBlockInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *PutPublicAccessBlockInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "PutPublicAccessBlockInput"}
	if s.AccountId == nil {
		invalidParams.Add(request.NewErrParamRequired("AccountId"))
	}
	if s.PublicAccessBlockConfiguration == nil {
		invalidParams.Add(request.NewErrParamRequired("PublicAccessBlockConfiguration"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetAccountId sets the AccountId field's value.
func (s *PutPublicAccessBlockInput) SetAccountId(v string) *PutPublicAccessBlockInput {
	s.AccountId = &v
	return s
}

// SetPublicAccessBlockConfiguration sets the PublicAccessBlockConfiguration field's value.
func (s *PutPublicAccessBlockInput) SetPublicAccessBlockConfiguration(v *PublicAccessBlockConfiguration) *PutPublicAccessBlockInput {
	s.PublicAccessBlockConfiguration = v
	return s
}

type PutPublicAccessBlockOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s PutPublicAccessBlockOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s PutPublicAccessBlockOutput) GoString() string {
	return s.String()
}
