// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package restxmlservice

import (
	"bytes"
	"fmt"
	"io"
	"sync"
	"sync/atomic"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/awsutil"
	"github.com/aws/aws-sdk-go/aws/client"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/private/protocol"
	"github.com/aws/aws-sdk-go/private/protocol/eventstream"
	"github.com/aws/aws-sdk-go/private/protocol/eventstream/eventstreamapi"
	"github.com/aws/aws-sdk-go/private/protocol/rest"
	"github.com/aws/aws-sdk-go/private/protocol/restxml"
)

const opEmptyStream = "EmptyStream"

// EmptyStreamRequest generates a "aws/request.Request" representing the
// client's request for the EmptyStream operation. The "output" return
// value will be populated with the request's response once the request completes
// successfully.
//
// Use "Send" method on the returned Request to send the API call to the service.
// the "output" return value is not valid until after Send returns without error.
//
// See EmptyStream for more information on using the EmptyStream
// API call, and error handling.
//
// This method is useful when you want to inject custom logic or configuration
// into the SDK's request lifecycle. Such as custom headers, or retry logic.
//
//
//    // Example sending a request using the EmptyStreamRequest method.
//    req, resp := client.EmptyStreamRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
//
// See also, https://docs.aws.amazon.com/goto/WebAPI/RESTXMLService-0000-00-00/EmptyStream
func (c *RESTXMLService) EmptyStreamRequest(input *EmptyStreamInput) (req *request.Request, output *EmptyStreamOutput) {
	op := &request.Operation{
		Name:       opEmptyStream,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &EmptyStreamInput{}
	}

	output = &EmptyStreamOutput{}
	req = c.newRequest(op, input, output)
	req.Handlers.Send.Swap(client.LogHTTPResponseHandler.Name, client.LogHTTPResponseHeaderHandler)
	req.Handlers.Unmarshal.Swap(restxml.UnmarshalHandler.Name, rest.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBack(output.runEventStreamLoop)
	return
}

// EmptyStream API operation for REST XML Service.
//
// Returns awserr.Error for service API and SDK errors. Use runtime type assertions
// with awserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the AWS API reference guide for REST XML Service's
// API operation EmptyStream for usage and error information.
// See also, https://docs.aws.amazon.com/goto/WebAPI/RESTXMLService-0000-00-00/EmptyStream
func (c *RESTXMLService) EmptyStream(input *EmptyStreamInput) (*EmptyStreamOutput, error) {
	req, out := c.EmptyStreamRequest(input)
	return out, req.Send()
}

// EmptyStreamWithContext is the same as EmptyStream with the addition of
// the ability to pass a context and additional request options.
//
// See EmptyStream for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *RESTXMLService) EmptyStreamWithContext(ctx aws.Context, input *EmptyStreamInput, opts ...request.Option) (*EmptyStreamOutput, error) {
	req, out := c.EmptyStreamRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opGetEventStream = "GetEventStream"

// GetEventStreamRequest generates a "aws/request.Request" representing the
// client's request for the GetEventStream operation. The "output" return
// value will be populated with the request's response once the request completes
// successfully.
//
// Use "Send" method on the returned Request to send the API call to the service.
// the "output" return value is not valid until after Send returns without error.
//
// See GetEventStream for more information on using the GetEventStream
// API call, and error handling.
//
// This method is useful when you want to inject custom logic or configuration
// into the SDK's request lifecycle. Such as custom headers, or retry logic.
//
//
//    // Example sending a request using the GetEventStreamRequest method.
//    req, resp := client.GetEventStreamRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
//
// See also, https://docs.aws.amazon.com/goto/WebAPI/RESTXMLService-0000-00-00/GetEventStream
func (c *RESTXMLService) GetEventStreamRequest(input *GetEventStreamInput) (req *request.Request, output *GetEventStreamOutput) {
	op := &request.Operation{
		Name:       opGetEventStream,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetEventStreamInput{}
	}

	output = &GetEventStreamOutput{}
	req = c.newRequest(op, input, output)
	req.Handlers.Send.Swap(client.LogHTTPResponseHandler.Name, client.LogHTTPResponseHeaderHandler)
	req.Handlers.Unmarshal.Swap(restxml.UnmarshalHandler.Name, rest.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBack(output.runEventStreamLoop)
	return
}

// GetEventStream API operation for REST XML Service.
//
// Returns awserr.Error for service API and SDK errors. Use runtime type assertions
// with awserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the AWS API reference guide for REST XML Service's
// API operation GetEventStream for usage and error information.
// See also, https://docs.aws.amazon.com/goto/WebAPI/RESTXMLService-0000-00-00/GetEventStream
func (c *RESTXMLService) GetEventStream(input *GetEventStreamInput) (*GetEventStreamOutput, error) {
	req, out := c.GetEventStreamRequest(input)
	return out, req.Send()
}

// GetEventStreamWithContext is the same as GetEventStream with the addition of
// the ability to pass a context and additional request options.
//
// See GetEventStream for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *RESTXMLService) GetEventStreamWithContext(ctx aws.Context, input *GetEventStreamInput, opts ...request.Option) (*GetEventStreamOutput, error) {
	req, out := c.GetEventStreamRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type EmptyEvent struct {
	_ struct{} `locationName:"EmptyEvent" type:"structure"`
}

// String returns the string representation
func (s EmptyEvent) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s EmptyEvent) GoString() string {
	return s.String()
}

// The EmptyEvent is and event in the EventStream group of events.
func (s *EmptyEvent) eventEventStream() {}

// UnmarshalEvent unmarshals the EventStream Message into the EmptyEvent value.
// This method is only used internally within the SDK's EventStream handling.
func (s *EmptyEvent) UnmarshalEvent(
	payloadUnmarshaler protocol.PayloadUnmarshaler,
	msg eventstream.Message,
) error {
	return nil
}

// EmptyStreamEventStream provides handling of EventStreams for
// the EmptyStream API.
//
// Use this type to receive EmptyEventStream events. The events
// can be read from the Events channel member.
//
// The events that can be received are:
//
type EmptyStreamEventStream struct {
	// Reader is the EventStream reader for the EmptyEventStream
	// events. This value is automatically set by the SDK when the API call is made
	// Use this member when unit testing your code with the SDK to mock out the
	// EventStream Reader.
	//
	// Must not be nil.
	Reader EmptyStreamEventStreamReader

	// StreamCloser is the io.Closer for the EventStream connection. For HTTP
	// EventStream this is the response Body. The stream will be closed when
	// the Close method of the EventStream is called.
	StreamCloser io.Closer
}

// Close closes the EventStream. This will also cause the Events channel to be
// closed. You can use the closing of the Events channel to terminate your
// application's read from the API's EventStream.
//
// Will close the underlying EventStream reader. For EventStream over HTTP
// connection this will also close the HTTP connection.
//
// Close must be called when done using the EventStream API. Not calling Close
// may result in resource leaks.
func (es *EmptyStreamEventStream) Close() (err error) {
	es.Reader.Close()
	return es.Err()
}

// Err returns any error that occurred while reading EventStream Events from
// the service API's response. Returns nil if there were no errors.
func (es *EmptyStreamEventStream) Err() error {
	if err := es.Reader.Err(); err != nil {
		return err
	}
	es.StreamCloser.Close()

	return nil
}

// Events returns a channel to read EventStream Events from the
// EmptyStream API.
//
// These events are:
//
func (es *EmptyStreamEventStream) Events() <-chan EmptyEventStreamEvent {
	return es.Reader.Events()
}

// EmptyEventStreamEvent groups together all EventStream
// events read from the EmptyStream API.
//
// These events are:
//
type EmptyEventStreamEvent interface {
	eventEmptyEventStream()
}

// EmptyStreamEventStreamReader provides the interface for reading EventStream
// Events from the EmptyStream API. The
// default implementation for this interface will be EmptyStreamEventStream.
//
// The reader's Close method must allow multiple concurrent calls.
//
// These events are:
//
type EmptyStreamEventStreamReader interface {
	// Returns a channel of events as they are read from the event stream.
	Events() <-chan EmptyEventStreamEvent

	// Close will close the underlying event stream reader. For event stream over
	// HTTP this will also close the HTTP connection.
	Close() error

	// Returns any error that has occurred while reading from the event stream.
	Err() error
}

type readEmptyStreamEventStream struct {
	eventReader *eventstreamapi.EventReader
	stream      chan EmptyEventStreamEvent
	errVal      atomic.Value

	done      chan struct{}
	closeOnce sync.Once
}

func newReadEmptyStreamEventStream(
	reader io.ReadCloser,
	unmarshalers request.HandlerList,
	logger aws.Logger,
	logLevel aws.LogLevelType,
) *readEmptyStreamEventStream {
	r := &readEmptyStreamEventStream{
		stream: make(chan EmptyEventStreamEvent),
		done:   make(chan struct{}),
	}

	r.eventReader = eventstreamapi.NewEventReader(
		reader,
		protocol.HandlerPayloadUnmarshal{
			Unmarshalers: unmarshalers,
		},
		r.unmarshalerForEventType,
	)
	r.eventReader.UseLogger(logger, logLevel)

	return r
}

// Close will close the underlying event stream reader. For EventStream over
// HTTP this will also close the HTTP connection.
func (r *readEmptyStreamEventStream) Close() error {
	r.closeOnce.Do(r.safeClose)

	return r.Err()
}

func (r *readEmptyStreamEventStream) safeClose() {
	close(r.done)
	err := r.eventReader.Close()
	if err != nil {
		r.errVal.Store(err)
	}
}

func (r *readEmptyStreamEventStream) Err() error {
	if v := r.errVal.Load(); v != nil {
		return v.(error)
	}

	return nil
}

func (r *readEmptyStreamEventStream) Events() <-chan EmptyEventStreamEvent {
	return r.stream
}

func (r *readEmptyStreamEventStream) readEventStream() {
	defer close(r.stream)

	for {
		event, err := r.eventReader.ReadEvent()
		if err != nil {
			if err == io.EOF {
				return
			}
			select {
			case <-r.done:
				// If closed already ignore the error
				return
			default:
			}
			r.errVal.Store(err)
			return
		}

		select {
		case r.stream <- event.(EmptyEventStreamEvent):
		case <-r.done:
			return
		}
	}
}

func (r *readEmptyStreamEventStream) unmarshalerForEventType(
	eventType string,
) (eventstreamapi.Unmarshaler, error) {
	switch eventType {
	default:
		return nil, awserr.New(
			request.ErrCodeSerialization,
			fmt.Sprintf("unknown event type name, %s, for EmptyStreamEventStream", eventType),
			nil,
		)
	}
}

type EmptyStreamInput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s EmptyStreamInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s EmptyStreamInput) GoString() string {
	return s.String()
}

type EmptyStreamOutput struct {
	_ struct{} `type:"structure"`

	// Use EventStream to use the API's stream.
	EventStream *EmptyStreamEventStream `type:"structure"`
}

// String returns the string representation
func (s EmptyStreamOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s EmptyStreamOutput) GoString() string {
	return s.String()
}

// SetEventStream sets the EventStream field's value.
func (s *EmptyStreamOutput) SetEventStream(v *EmptyStreamEventStream) *EmptyStreamOutput {
	s.EventStream = v
	return s
}

func (s *EmptyStreamOutput) runEventStreamLoop(r *request.Request) {
	if r.Error != nil {
		return
	}
	reader := newReadEmptyStreamEventStream(
		r.HTTPResponse.Body,
		r.Handlers.UnmarshalStream,
		r.Config.Logger,
		r.Config.LogLevel.Value(),
	)
	go reader.readEventStream()

	eventStream := &EmptyStreamEventStream{
		StreamCloser: r.HTTPResponse.Body,
		Reader:       reader,
	}
	s.EventStream = eventStream
}

type ExceptionEvent struct {
	_ struct{} `locationName:"ExceptionEvent" type:"structure"`

	IntVal *int64 `type:"integer"`

	Message_ *string `locationName:"message" type:"string"`
}

// String returns the string representation
func (s ExceptionEvent) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s ExceptionEvent) GoString() string {
	return s.String()
}

// The ExceptionEvent is and event in the EventStream group of events.
func (s *ExceptionEvent) eventEventStream() {}

// UnmarshalEvent unmarshals the EventStream Message into the ExceptionEvent value.
// This method is only used internally within the SDK's EventStream handling.
func (s *ExceptionEvent) UnmarshalEvent(
	payloadUnmarshaler protocol.PayloadUnmarshaler,
	msg eventstream.Message,
) error {
	if err := payloadUnmarshaler.UnmarshalPayload(
		bytes.NewReader(msg.Payload), s,
	); err != nil {
		return err
	}
	return nil
}

// Code returns the exception type name.
func (s ExceptionEvent) Code() string {
	return "ExceptionEvent"
}

// Message returns the exception's message.
func (s ExceptionEvent) Message() string {
	return *s.Message_
}

// OrigErr always returns nil, satisfies awserr.Error interface.
func (s ExceptionEvent) OrigErr() error {
	return nil
}

func (s ExceptionEvent) Error() string {
	return fmt.Sprintf("%s: %s", s.Code(), s.Message())
}

type ExplicitPayloadEvent struct {
	_ struct{} `locationName:"ExplicitPayloadEvent" type:"structure" payload:"NestedVal"`

	LongVal *int64 `location:"header" type:"long"`

	NestedVal *NestedShape `locationName:"NestedVal" type:"structure"`

	StringVal *string `location:"header" type:"string"`
}

// String returns the string representation
func (s ExplicitPayloadEvent) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s ExplicitPayloadEvent) GoString() string {
	return s.String()
}

// SetLongVal sets the LongVal field's value.
func (s *ExplicitPayloadEvent) SetLongVal(v int64) *ExplicitPayloadEvent {
	s.LongVal = &v
	return s
}

// SetNestedVal sets the NestedVal field's value.
func (s *ExplicitPayloadEvent) SetNestedVal(v *NestedShape) *ExplicitPayloadEvent {
	s.NestedVal = v
	return s
}

// SetStringVal sets the StringVal field's value.
func (s *ExplicitPayloadEvent) SetStringVal(v string) *ExplicitPayloadEvent {
	s.StringVal = &v
	return s
}

// The ExplicitPayloadEvent is and event in the EventStream group of events.
func (s *ExplicitPayloadEvent) eventEventStream() {}

// UnmarshalEvent unmarshals the EventStream Message into the ExplicitPayloadEvent value.
// This method is only used internally within the SDK's EventStream handling.
func (s *ExplicitPayloadEvent) UnmarshalEvent(
	payloadUnmarshaler protocol.PayloadUnmarshaler,
	msg eventstream.Message,
) error {
	if hv := msg.Headers.Get("LongVal"); hv != nil {
		v := hv.Get().(int64)
		s.LongVal = &v
	}
	if hv := msg.Headers.Get("StringVal"); hv != nil {
		v := hv.Get().(string)
		s.StringVal = &v
	}
	if err := payloadUnmarshaler.UnmarshalPayload(
		bytes.NewReader(msg.Payload), s,
	); err != nil {
		return err
	}
	return nil
}

// GetEventStreamEventStream provides handling of EventStreams for
// the GetEventStream API.
//
// Use this type to receive EventStream events. The events
// can be read from the Events channel member.
//
// The events that can be received are:
//
//     * EmptyEvent
//     * ExplicitPayloadEvent
//     * HeaderOnlyEvent
//     * ImplicitPayloadEvent
//     * PayloadOnlyEvent
//     * PayloadOnlyBlobEvent
//     * PayloadOnlyStringEvent
type GetEventStreamEventStream struct {
	// Reader is the EventStream reader for the EventStream
	// events. This value is automatically set by the SDK when the API call is made
	// Use this member when unit testing your code with the SDK to mock out the
	// EventStream Reader.
	//
	// Must not be nil.
	Reader GetEventStreamEventStreamReader

	// StreamCloser is the io.Closer for the EventStream connection. For HTTP
	// EventStream this is the response Body. The stream will be closed when
	// the Close method of the EventStream is called.
	StreamCloser io.Closer
}

// Close closes the EventStream. This will also cause the Events channel to be
// closed. You can use the closing of the Events channel to terminate your
// application's read from the API's EventStream.
//
// Will close the underlying EventStream reader. For EventStream over HTTP
// connection this will also close the HTTP connection.
//
// Close must be called when done using the EventStream API. Not calling Close
// may result in resource leaks.
func (es *GetEventStreamEventStream) Close() (err error) {
	es.Reader.Close()
	return es.Err()
}

// Err returns any error that occurred while reading EventStream Events from
// the service API's response. Returns nil if there were no errors.
func (es *GetEventStreamEventStream) Err() error {
	if err := es.Reader.Err(); err != nil {
		return err
	}
	es.StreamCloser.Close()

	return nil
}

// Events returns a channel to read EventStream Events from the
// GetEventStream API.
//
// These events are:
//
//     * EmptyEvent
//     * ExplicitPayloadEvent
//     * HeaderOnlyEvent
//     * ImplicitPayloadEvent
//     * PayloadOnlyEvent
//     * PayloadOnlyBlobEvent
//     * PayloadOnlyStringEvent
func (es *GetEventStreamEventStream) Events() <-chan EventStreamEvent {
	return es.Reader.Events()
}

// EventStreamEvent groups together all EventStream
// events read from the GetEventStream API.
//
// These events are:
//
//     * EmptyEvent
//     * ExplicitPayloadEvent
//     * HeaderOnlyEvent
//     * ImplicitPayloadEvent
//     * PayloadOnlyEvent
//     * PayloadOnlyBlobEvent
//     * PayloadOnlyStringEvent
type EventStreamEvent interface {
	eventEventStream()
}

// GetEventStreamEventStreamReader provides the interface for reading EventStream
// Events from the GetEventStream API. The
// default implementation for this interface will be GetEventStreamEventStream.
//
// The reader's Close method must allow multiple concurrent calls.
//
// These events are:
//
//     * EmptyEvent
//     * ExplicitPayloadEvent
//     * HeaderOnlyEvent
//     * ImplicitPayloadEvent
//     * PayloadOnlyEvent
//     * PayloadOnlyBlobEvent
//     * PayloadOnlyStringEvent
type GetEventStreamEventStreamReader interface {
	// Returns a channel of events as they are read from the event stream.
	Events() <-chan EventStreamEvent

	// Close will close the underlying event stream reader. For event stream over
	// HTTP this will also close the HTTP connection.
	Close() error

	// Returns any error that has occurred while reading from the event stream.
	Err() error
}

type readGetEventStreamEventStream struct {
	eventReader *eventstreamapi.EventReader
	stream      chan EventStreamEvent
	errVal      atomic.Value

	done      chan struct{}
	closeOnce sync.Once
}

func newReadGetEventStreamEventStream(
	reader io.ReadCloser,
	unmarshalers request.HandlerList,
	logger aws.Logger,
	logLevel aws.LogLevelType,
) *readGetEventStreamEventStream {
	r := &readGetEventStreamEventStream{
		stream: make(chan EventStreamEvent),
		done:   make(chan struct{}),
	}

	r.eventReader = eventstreamapi.NewEventReader(
		reader,
		protocol.HandlerPayloadUnmarshal{
			Unmarshalers: unmarshalers,
		},
		r.unmarshalerForEventType,
	)
	r.eventReader.UseLogger(logger, logLevel)

	return r
}

// Close will close the underlying event stream reader. For EventStream over
// HTTP this will also close the HTTP connection.
func (r *readGetEventStreamEventStream) Close() error {
	r.closeOnce.Do(r.safeClose)

	return r.Err()
}

func (r *readGetEventStreamEventStream) safeClose() {
	close(r.done)
	err := r.eventReader.Close()
	if err != nil {
		r.errVal.Store(err)
	}
}

func (r *readGetEventStreamEventStream) Err() error {
	if v := r.errVal.Load(); v != nil {
		return v.(error)
	}

	return nil
}

func (r *readGetEventStreamEventStream) Events() <-chan EventStreamEvent {
	return r.stream
}

func (r *readGetEventStreamEventStream) readEventStream() {
	defer close(r.stream)

	for {
		event, err := r.eventReader.ReadEvent()
		if err != nil {
			if err == io.EOF {
				return
			}
			select {
			case <-r.done:
				// If closed already ignore the error
				return
			default:
			}
			r.errVal.Store(err)
			return
		}

		select {
		case r.stream <- event.(EventStreamEvent):
		case <-r.done:
			return
		}
	}
}

func (r *readGetEventStreamEventStream) unmarshalerForEventType(
	eventType string,
) (eventstreamapi.Unmarshaler, error) {
	switch eventType {
	case "Empty":
		return &EmptyEvent{}, nil

	case "ExplicitPayload":
		return &ExplicitPayloadEvent{}, nil

	case "Headers":
		return &HeaderOnlyEvent{}, nil

	case "ImplicitPayload":
		return &ImplicitPayloadEvent{}, nil

	case "PayloadOnly":
		return &PayloadOnlyEvent{}, nil

	case "PayloadOnlyBlob":
		return &PayloadOnlyBlobEvent{}, nil

	case "PayloadOnlyString":
		return &PayloadOnlyStringEvent{}, nil

	case "Exception":
		return &ExceptionEvent{}, nil
	default:
		return nil, awserr.New(
			request.ErrCodeSerialization,
			fmt.Sprintf("unknown event type name, %s, for GetEventStreamEventStream", eventType),
			nil,
		)
	}
}

type GetEventStreamInput struct {
	_ struct{} `type:"structure"`

	InputVal *string `type:"string"`
}

// String returns the string representation
func (s GetEventStreamInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s GetEventStreamInput) GoString() string {
	return s.String()
}

// SetInputVal sets the InputVal field's value.
func (s *GetEventStreamInput) SetInputVal(v string) *GetEventStreamInput {
	s.InputVal = &v
	return s
}

type GetEventStreamOutput struct {
	_ struct{} `type:"structure"`

	// Use EventStream to use the API's stream.
	EventStream *GetEventStreamEventStream `type:"structure"`

	IntVal *int64 `type:"integer"`

	StrVal *string `type:"string"`
}

// String returns the string representation
func (s GetEventStreamOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s GetEventStreamOutput) GoString() string {
	return s.String()
}

// SetEventStream sets the EventStream field's value.
func (s *GetEventStreamOutput) SetEventStream(v *GetEventStreamEventStream) *GetEventStreamOutput {
	s.EventStream = v
	return s
}

// SetIntVal sets the IntVal field's value.
func (s *GetEventStreamOutput) SetIntVal(v int64) *GetEventStreamOutput {
	s.IntVal = &v
	return s
}

// SetStrVal sets the StrVal field's value.
func (s *GetEventStreamOutput) SetStrVal(v string) *GetEventStreamOutput {
	s.StrVal = &v
	return s
}

func (s *GetEventStreamOutput) runEventStreamLoop(r *request.Request) {
	if r.Error != nil {
		return
	}
	reader := newReadGetEventStreamEventStream(
		r.HTTPResponse.Body,
		r.Handlers.UnmarshalStream,
		r.Config.Logger,
		r.Config.LogLevel.Value(),
	)
	go reader.readEventStream()

	eventStream := &GetEventStreamEventStream{
		StreamCloser: r.HTTPResponse.Body,
		Reader:       reader,
	}
	s.EventStream = eventStream
}

type HeaderOnlyEvent struct {
	_ struct{} `locationName:"HeaderOnlyEvent" type:"structure"`

	// BlobVal is automatically base64 encoded/decoded by the SDK.
	BlobVal []byte `location:"header" type:"blob"`

	BoolVal *bool `location:"header" type:"boolean"`

	ByteVal *int64 `location:"header" type:"byte"`

	IntegerVal *int64 `location:"header" type:"integer"`

	LongVal *int64 `location:"header" type:"long"`

	ShortVal *int64 `location:"header" type:"short"`

	StringVal *string `location:"header" type:"string"`

	TimeVal *time.Time `location:"header" type:"timestamp"`
}

// String returns the string representation
func (s HeaderOnlyEvent) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s HeaderOnlyEvent) GoString() string {
	return s.String()
}

// SetBlobVal sets the BlobVal field's value.
func (s *HeaderOnlyEvent) SetBlobVal(v []byte) *HeaderOnlyEvent {
	s.BlobVal = v
	return s
}

// SetBoolVal sets the BoolVal field's value.
func (s *HeaderOnlyEvent) SetBoolVal(v bool) *HeaderOnlyEvent {
	s.BoolVal = &v
	return s
}

// SetByteVal sets the ByteVal field's value.
func (s *HeaderOnlyEvent) SetByteVal(v int64) *HeaderOnlyEvent {
	s.ByteVal = &v
	return s
}

// SetIntegerVal sets the IntegerVal field's value.
func (s *HeaderOnlyEvent) SetIntegerVal(v int64) *HeaderOnlyEvent {
	s.IntegerVal = &v
	return s
}

// SetLongVal sets the LongVal field's value.
func (s *HeaderOnlyEvent) SetLongVal(v int64) *HeaderOnlyEvent {
	s.LongVal = &v
	return s
}

// SetShortVal sets the ShortVal field's value.
func (s *HeaderOnlyEvent) SetShortVal(v int64) *HeaderOnlyEvent {
	s.ShortVal = &v
	return s
}

// SetStringVal sets the StringVal field's value.
func (s *HeaderOnlyEvent) SetStringVal(v string) *HeaderOnlyEvent {
	s.StringVal = &v
	return s
}

// SetTimeVal sets the TimeVal field's value.
func (s *HeaderOnlyEvent) SetTimeVal(v time.Time) *HeaderOnlyEvent {
	s.TimeVal = &v
	return s
}

// The HeaderOnlyEvent is and event in the EventStream group of events.
func (s *HeaderOnlyEvent) eventEventStream() {}

// UnmarshalEvent unmarshals the EventStream Message into the HeaderOnlyEvent value.
// This method is only used internally within the SDK's EventStream handling.
func (s *HeaderOnlyEvent) UnmarshalEvent(
	payloadUnmarshaler protocol.PayloadUnmarshaler,
	msg eventstream.Message,
) error {
	if hv := msg.Headers.Get("BlobVal"); hv != nil {
		v := hv.Get().([]byte)
		s.BlobVal = v
	}
	if hv := msg.Headers.Get("BoolVal"); hv != nil {
		v := hv.Get().(bool)
		s.BoolVal = &v
	}
	if hv := msg.Headers.Get("ByteVal"); hv != nil {
		v := hv.Get().(int8)
		m := int64(v)
		s.ByteVal = &m
	}
	if hv := msg.Headers.Get("IntegerVal"); hv != nil {
		v := hv.Get().(int32)
		m := int64(v)
		s.IntegerVal = &m
	}
	if hv := msg.Headers.Get("LongVal"); hv != nil {
		v := hv.Get().(int64)
		s.LongVal = &v
	}
	if hv := msg.Headers.Get("ShortVal"); hv != nil {
		v := hv.Get().(int16)
		m := int64(v)
		s.ShortVal = &m
	}
	if hv := msg.Headers.Get("StringVal"); hv != nil {
		v := hv.Get().(string)
		s.StringVal = &v
	}
	if hv := msg.Headers.Get("TimeVal"); hv != nil {
		v := hv.Get().(time.Time)
		s.TimeVal = &v
	}
	return nil
}

type ImplicitPayloadEvent struct {
	_ struct{} `locationName:"ImplicitPayloadEvent" type:"structure"`

	ByteVal *int64 `location:"header" type:"byte"`

	IntegerVal *int64 `type:"integer"`

	ShortVal *int64 `type:"short"`
}

// String returns the string representation
func (s ImplicitPayloadEvent) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s ImplicitPayloadEvent) GoString() string {
	return s.String()
}

// SetByteVal sets the ByteVal field's value.
func (s *ImplicitPayloadEvent) SetByteVal(v int64) *ImplicitPayloadEvent {
	s.ByteVal = &v
	return s
}

// SetIntegerVal sets the IntegerVal field's value.
func (s *ImplicitPayloadEvent) SetIntegerVal(v int64) *ImplicitPayloadEvent {
	s.IntegerVal = &v
	return s
}

// SetShortVal sets the ShortVal field's value.
func (s *ImplicitPayloadEvent) SetShortVal(v int64) *ImplicitPayloadEvent {
	s.ShortVal = &v
	return s
}

// The ImplicitPayloadEvent is and event in the EventStream group of events.
func (s *ImplicitPayloadEvent) eventEventStream() {}

// UnmarshalEvent unmarshals the EventStream Message into the ImplicitPayloadEvent value.
// This method is only used internally within the SDK's EventStream handling.
func (s *ImplicitPayloadEvent) UnmarshalEvent(
	payloadUnmarshaler protocol.PayloadUnmarshaler,
	msg eventstream.Message,
) error {
	if hv := msg.Headers.Get("ByteVal"); hv != nil {
		v := hv.Get().(int8)
		m := int64(v)
		s.ByteVal = &m
	}
	if err := payloadUnmarshaler.UnmarshalPayload(
		bytes.NewReader(msg.Payload), s,
	); err != nil {
		return err
	}
	return nil
}

type NestedShape struct {
	_ struct{} `type:"structure"`

	IntVal *int64 `type:"integer"`

	StrVal *string `type:"string"`
}

// String returns the string representation
func (s NestedShape) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s NestedShape) GoString() string {
	return s.String()
}

// SetIntVal sets the IntVal field's value.
func (s *NestedShape) SetIntVal(v int64) *NestedShape {
	s.IntVal = &v
	return s
}

// SetStrVal sets the StrVal field's value.
func (s *NestedShape) SetStrVal(v string) *NestedShape {
	s.StrVal = &v
	return s
}

type PayloadOnlyBlobEvent struct {
	_ struct{} `locationName:"PayloadOnlyBlobEvent" type:"structure" payload:"BlobPayload"`

	// BlobPayload is automatically base64 encoded/decoded by the SDK.
	BlobPayload []byte `type:"blob"`
}

// String returns the string representation
func (s PayloadOnlyBlobEvent) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s PayloadOnlyBlobEvent) GoString() string {
	return s.String()
}

// SetBlobPayload sets the BlobPayload field's value.
func (s *PayloadOnlyBlobEvent) SetBlobPayload(v []byte) *PayloadOnlyBlobEvent {
	s.BlobPayload = v
	return s
}

// The PayloadOnlyBlobEvent is and event in the EventStream group of events.
func (s *PayloadOnlyBlobEvent) eventEventStream() {}

// UnmarshalEvent unmarshals the EventStream Message into the PayloadOnlyBlobEvent value.
// This method is only used internally within the SDK's EventStream handling.
func (s *PayloadOnlyBlobEvent) UnmarshalEvent(
	payloadUnmarshaler protocol.PayloadUnmarshaler,
	msg eventstream.Message,
) error {
	s.BlobPayload = make([]byte, len(msg.Payload))
	copy(s.BlobPayload, msg.Payload)
	return nil
}

type PayloadOnlyEvent struct {
	_ struct{} `locationName:"PayloadOnlyEvent" type:"structure" payload:"NestedVal"`

	NestedVal *NestedShape `locationName:"NestedVal" type:"structure"`
}

// String returns the string representation
func (s PayloadOnlyEvent) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s PayloadOnlyEvent) GoString() string {
	return s.String()
}

// SetNestedVal sets the NestedVal field's value.
func (s *PayloadOnlyEvent) SetNestedVal(v *NestedShape) *PayloadOnlyEvent {
	s.NestedVal = v
	return s
}

// The PayloadOnlyEvent is and event in the EventStream group of events.
func (s *PayloadOnlyEvent) eventEventStream() {}

// UnmarshalEvent unmarshals the EventStream Message into the PayloadOnlyEvent value.
// This method is only used internally within the SDK's EventStream handling.
func (s *PayloadOnlyEvent) UnmarshalEvent(
	payloadUnmarshaler protocol.PayloadUnmarshaler,
	msg eventstream.Message,
) error {
	if err := payloadUnmarshaler.UnmarshalPayload(
		bytes.NewReader(msg.Payload), s,
	); err != nil {
		return err
	}
	return nil
}

type PayloadOnlyStringEvent struct {
	_ struct{} `locationName:"PayloadOnlyStringEvent" type:"structure" payload:"StringPayload"`

	StringPayload *string `locationName:"StringPayload" type:"string"`
}

// String returns the string representation
func (s PayloadOnlyStringEvent) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s PayloadOnlyStringEvent) GoString() string {
	return s.String()
}

// SetStringPayload sets the StringPayload field's value.
func (s *PayloadOnlyStringEvent) SetStringPayload(v string) *PayloadOnlyStringEvent {
	s.StringPayload = &v
	return s
}

// The PayloadOnlyStringEvent is and event in the EventStream group of events.
func (s *PayloadOnlyStringEvent) eventEventStream() {}

// UnmarshalEvent unmarshals the EventStream Message into the PayloadOnlyStringEvent value.
// This method is only used internally within the SDK's EventStream handling.
func (s *PayloadOnlyStringEvent) UnmarshalEvent(
	payloadUnmarshaler protocol.PayloadUnmarshaler,
	msg eventstream.Message,
) error {
	s.StringPayload = aws.String(string(msg.Payload))
	return nil
}
