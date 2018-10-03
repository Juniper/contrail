// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package applicationautoscaling

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/client"
	"github.com/aws/aws-sdk-go/aws/client/metadata"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/signer/v4"
	"github.com/aws/aws-sdk-go/private/protocol/jsonrpc"
)

// ApplicationAutoScaling provides the API operation methods for making requests to
// Application Auto Scaling. See this package's package overview docs
// for details on the service.
//
// ApplicationAutoScaling methods are safe to use concurrently. It is not safe to
// modify mutate any of the struct's properties though.
type ApplicationAutoScaling struct {
	*client.Client
}

// Used for custom client initialization logic
var initClient func(*client.Client)

// Used for custom request initialization logic
var initRequest func(*request.Request)

// Service information constants
const (
	ServiceName = "autoscaling"              // Name of service.
	EndpointsID = "application-autoscaling"  // ID to lookup a service endpoint with.
	ServiceID   = "Application Auto Scaling" // ServiceID is a unique identifer of a specific service.
)

// New creates a new instance of the ApplicationAutoScaling client with a session.
// If additional configuration is needed for the client instance use the optional
// aws.Config parameter to add your extra config.
//
// Example:
//     // Create a ApplicationAutoScaling client from just a session.
//     svc := applicationautoscaling.New(mySession)
//
//     // Create a ApplicationAutoScaling client with additional configuration
//     svc := applicationautoscaling.New(mySession, aws.NewConfig().WithRegion("us-west-2"))
func New(p client.ConfigProvider, cfgs ...*aws.Config) *ApplicationAutoScaling {
	c := p.ClientConfig(EndpointsID, cfgs...)
	if c.SigningNameDerived || len(c.SigningName) == 0 {
		c.SigningName = "application-autoscaling"
	}
	return newClient(*c.Config, c.Handlers, c.Endpoint, c.SigningRegion, c.SigningName)
}

// newClient creates, initializes and returns a new service client instance.
func newClient(cfg aws.Config, handlers request.Handlers, endpoint, signingRegion, signingName string) *ApplicationAutoScaling {
	svc := &ApplicationAutoScaling{
		Client: client.New(
			cfg,
			metadata.ClientInfo{
				ServiceName:   ServiceName,
				ServiceID:     ServiceID,
				SigningName:   signingName,
				SigningRegion: signingRegion,
				Endpoint:      endpoint,
				APIVersion:    "2016-02-06",
				JSONVersion:   "1.1",
				TargetPrefix:  "AnyScaleFrontendService",
			},
			handlers,
		),
	}

	// Handlers
	svc.Handlers.Sign.PushBackNamed(v4.SignRequestHandler)
	svc.Handlers.Build.PushBackNamed(jsonrpc.BuildHandler)
	svc.Handlers.Unmarshal.PushBackNamed(jsonrpc.UnmarshalHandler)
	svc.Handlers.UnmarshalMeta.PushBackNamed(jsonrpc.UnmarshalMetaHandler)
	svc.Handlers.UnmarshalError.PushBackNamed(jsonrpc.UnmarshalErrorHandler)

	// Run custom client initialization if present
	if initClient != nil {
		initClient(svc.Client)
	}

	return svc
}

// newRequest creates a new request for a ApplicationAutoScaling operation and runs any
// custom request initialization.
func (c *ApplicationAutoScaling) newRequest(op *request.Operation, params, data interface{}) *request.Request {
	req := c.NewRequest(op, params, data)

	// Run custom request initialization if present
	if initRequest != nil {
		initRequest(req)
	}

	return req
}
