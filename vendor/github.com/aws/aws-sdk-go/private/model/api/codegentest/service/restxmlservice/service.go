// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package restxmlservice

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/client"
	"github.com/aws/aws-sdk-go/aws/client/metadata"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/signer/v4"
	"github.com/aws/aws-sdk-go/private/protocol/restxml"
)

// RESTXMLService provides the API operation methods for making requests to
// REST XML Service. See this package's package overview docs
// for details on the service.
//
// RESTXMLService methods are safe to use concurrently. It is not safe to
// modify mutate any of the struct's properties though.
type RESTXMLService struct {
	*client.Client
}

// Used for custom client initialization logic
var initClient func(*client.Client)

// Used for custom request initialization logic
var initRequest func(*request.Request)

// Service information constants
const (
	ServiceName = "RESTXMLService" // Name of service.
	EndpointsID = "restxmlservice" // ID to lookup a service endpoint with.
	ServiceID   = "RESTXMLService" // ServiceID is a unique identifer of a specific service.
)

// New creates a new instance of the RESTXMLService client with a session.
// If additional configuration is needed for the client instance use the optional
// aws.Config parameter to add your extra config.
//
// Example:
//     // Create a RESTXMLService client from just a session.
//     svc := restxmlservice.New(mySession)
//
//     // Create a RESTXMLService client with additional configuration
//     svc := restxmlservice.New(mySession, aws.NewConfig().WithRegion("us-west-2"))
func New(p client.ConfigProvider, cfgs ...*aws.Config) *RESTXMLService {
	c := p.ClientConfig(EndpointsID, cfgs...)
	return newClient(*c.Config, c.Handlers, c.Endpoint, c.SigningRegion, c.SigningName)
}

// newClient creates, initializes and returns a new service client instance.
func newClient(cfg aws.Config, handlers request.Handlers, endpoint, signingRegion, signingName string) *RESTXMLService {
	svc := &RESTXMLService{
		Client: client.New(
			cfg,
			metadata.ClientInfo{
				ServiceName:   ServiceName,
				ServiceID:     ServiceID,
				SigningName:   signingName,
				SigningRegion: signingRegion,
				Endpoint:      endpoint,
				APIVersion:    "0000-00-00",
				JSONVersion:   "1.1",
				TargetPrefix:  "RESTXMLService_00000000",
			},
			handlers,
		),
	}

	// Handlers
	svc.Handlers.Sign.PushBackNamed(v4.SignRequestHandler)
	svc.Handlers.Build.PushBackNamed(restxml.BuildHandler)
	svc.Handlers.Unmarshal.PushBackNamed(restxml.UnmarshalHandler)
	svc.Handlers.UnmarshalMeta.PushBackNamed(restxml.UnmarshalMetaHandler)
	svc.Handlers.UnmarshalError.PushBackNamed(restxml.UnmarshalErrorHandler)

	svc.Handlers.UnmarshalStream.PushBackNamed(restxml.UnmarshalHandler)

	// Run custom client initialization if present
	if initClient != nil {
		initClient(svc.Client)
	}

	return svc
}

// newRequest creates a new request for a RESTXMLService operation and runs any
// custom request initialization.
func (c *RESTXMLService) newRequest(op *request.Operation, params, data interface{}) *request.Request {
	req := c.NewRequest(op, params, data)

	// Run custom request initialization if present
	if initRequest != nil {
		initRequest(req)
	}

	return req
}
