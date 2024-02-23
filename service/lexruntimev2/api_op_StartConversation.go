// Code generated by smithy-go-codegen DO NOT EDIT.

package lexruntimev2

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/protocol/eventstream/eventstreamapi"
	"github.com/aws/aws-sdk-go-v2/service/lexruntimev2/types"
	"github.com/aws/smithy-go/middleware"
	smithysync "github.com/aws/smithy-go/sync"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"sync"
	"time"
)

// Starts an HTTP/2 bidirectional event stream that enables you to send audio,
// text, or DTMF input in real time. After your application starts a conversation,
// users send input to Amazon Lex V2 as a stream of events. Amazon Lex V2 processes
// the incoming events and responds with streaming text or audio events. Audio
// input must be in the following format: audio/lpcm sample-rate=8000
// sample-size-bits=16 channel-count=1; is-big-endian=false . If the optional
// post-fulfillment response is specified, the messages are returned as follows.
// For more information, see PostFulfillmentStatusSpecification (https://docs.aws.amazon.com/lexv2/latest/dg/API_PostFulfillmentStatusSpecification.html)
// .
//   - Success message - Returned if the Lambda function completes successfully
//     and the intent state is fulfilled or ready fulfillment if the message is
//     present.
//   - Failed message - The failed message is returned if the Lambda function
//     throws an exception or if the Lambda function returns a failed intent state
//     without a message.
//   - Timeout message - If you don't configure a timeout message and a timeout,
//     and the Lambda function doesn't return within 30 seconds, the timeout message is
//     returned. If you configure a timeout, the timeout message is returned when the
//     period times out.
//
// For more information, see Completion message (https://docs.aws.amazon.com/lexv2/latest/dg/streaming-progress.html#progress-complete.html)
// . If the optional update message is configured, it is played at the specified
// frequency while the Lambda function is running and the update message state is
// active. If the fulfillment update message is not active, the Lambda function
// runs with a 30 second timeout. For more information, see Update message  (https://docs.aws.amazon.com/lexv2/latest/dg/streaming-progress.html#progress-update.html)
// The StartConversation operation is supported only in the following SDKs:
//   - AWS SDK for C++ (https://docs.aws.amazon.com/goto/SdkForCpp/runtime.lex.v2-2020-08-07/StartConversation)
//   - AWS SDK for Java V2 (https://docs.aws.amazon.com/goto/SdkForJavaV2/runtime.lex.v2-2020-08-07/StartConversation)
//   - AWS SDK for Ruby V3 (https://docs.aws.amazon.com/goto/SdkForRubyV3/runtime.lex.v2-2020-08-07/StartConversation)
func (c *Client) StartConversation(ctx context.Context, params *StartConversationInput, optFns ...func(*Options)) (*StartConversationOutput, error) {
	if params == nil {
		params = &StartConversationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "StartConversation", params, optFns, c.addOperationStartConversationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*StartConversationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StartConversationInput struct {

	// The alias identifier in use for the bot that processes the request.
	//
	// This member is required.
	BotAliasId *string

	// The identifier of the bot to process the request.
	//
	// This member is required.
	BotId *string

	// The locale where the session is in use.
	//
	// This member is required.
	LocaleId *string

	// The identifier of the user session that is having the conversation.
	//
	// This member is required.
	SessionId *string

	// The conversation type that you are using the Amazon Lex V2. If the conversation
	// mode is AUDIO you can send both audio and DTMF information. If the mode is TEXT
	// you can only send text.
	ConversationMode types.ConversationMode

	noSmithyDocumentSerde
}

type StartConversationOutput struct {
	eventStream *StartConversationEventStream

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

// GetStream returns the type to interact with the event stream.
func (o *StartConversationOutput) GetStream() *StartConversationEventStream {
	return o.eventStream
}

func (c *Client) addOperationStartConversationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpStartConversation{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpStartConversation{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "StartConversation"); err != nil {
		return fmt.Errorf("add protocol finalizers: %v", err)
	}

	if err = addlegacyEndpointContextSetter(stack, options); err != nil {
		return err
	}
	if err = addEventStreamStartConversationMiddleware(stack, options); err != nil {
		return err
	}
	if err = smithyhttp.AddRequireMinimumProtocol(stack, 2, 0); err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = addClientRequestID(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addStreamingEventsPayload(stack); err != nil {
		return err
	}
	if err = addContentSHA256Header(stack); err != nil {
		return err
	}
	if err = addRetry(stack, options); err != nil {
		return err
	}
	if err = addRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = addRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack, options); err != nil {
		return err
	}
	if err = eventstreamapi.AddInitializeStreamWriter(stack); err != nil {
		return err
	}
	if err = addSetLegacyContextSigningOptionsMiddleware(stack); err != nil {
		return err
	}
	if err = addOpStartConversationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opStartConversation(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRecursionDetection(stack); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	if err = addDisableHTTPSMiddleware(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opStartConversation(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "StartConversation",
	}
}

// StartConversationEventStream provides the event stream handling for the StartConversation operation.
//
// For testing and mocking the event stream this type should be initialized via
// the NewStartConversationEventStream constructor function. Using the functional options
// to pass in nested mock behavior.
type StartConversationEventStream struct {
	// StartConversationRequestEventStreamWriter is the EventStream writer for the
	// StartConversationRequestEventStream events. This value is automatically set by
	// the SDK when the API call is made Use this member when unit testing your code
	// with the SDK to mock out the EventStream Writer.
	//
	// Must not be nil.
	Writer StartConversationRequestEventStreamWriter

	// StartConversationResponseEventStreamReader is the EventStream reader for the
	// StartConversationResponseEventStream events. This value is automatically set by
	// the SDK when the API call is made Use this member when unit testing your code
	// with the SDK to mock out the EventStream Reader.
	//
	// Must not be nil.
	Reader StartConversationResponseEventStreamReader

	done      chan struct{}
	closeOnce sync.Once
	err       *smithysync.OnceErr
}

// NewStartConversationEventStream initializes an StartConversationEventStream.
// This function should only be used for testing and mocking the StartConversationEventStream
// stream within your application.
//
// The Writer member must be set before writing events to the stream.
//
// The Reader member must be set before reading events from the stream.
func NewStartConversationEventStream(optFns ...func(*StartConversationEventStream)) *StartConversationEventStream {
	es := &StartConversationEventStream{
		done: make(chan struct{}),
		err:  smithysync.NewOnceErr(),
	}
	for _, fn := range optFns {
		fn(es)
	}
	return es
}

// Send writes the event to the stream blocking until the event is written.
// Returns an error if the event was not written.
func (es *StartConversationEventStream) Send(ctx context.Context, event types.StartConversationRequestEventStream) error {
	return es.Writer.Send(ctx, event)
}

// Events returns a channel to read events from.
func (es *StartConversationEventStream) Events() <-chan types.StartConversationResponseEventStream {
	return es.Reader.Events()
}

// Close closes the stream. This will also cause the stream to be closed.
// Close must be called when done using the stream API. Not calling Close
// may result in resource leaks.
//
// Will close the underlying EventStream writer and reader, and no more events can be
// sent or received.
func (es *StartConversationEventStream) Close() error {
	es.closeOnce.Do(es.safeClose)
	return es.Err()
}

func (es *StartConversationEventStream) safeClose() {
	close(es.done)

	t := time.NewTicker(time.Second)
	defer t.Stop()
	writeCloseDone := make(chan error)
	go func() {
		if err := es.Writer.Close(); err != nil {
			es.err.SetError(err)
		}
		close(writeCloseDone)
	}()
	select {
	case <-t.C:
	case <-writeCloseDone:
	}

	es.Reader.Close()
}

// Err returns any error that occurred while reading or writing EventStream Events
// from the service API's response. Returns nil if there were no errors.
func (es *StartConversationEventStream) Err() error {
	if err := es.err.Err(); err != nil {
		return err
	}

	if err := es.Writer.Err(); err != nil {
		return err
	}

	if err := es.Reader.Err(); err != nil {
		return err
	}

	return nil
}

func (es *StartConversationEventStream) waitStreamClose() {
	type errorSet interface {
		ErrorSet() <-chan struct{}
	}

	var inputErrCh <-chan struct{}
	if v, ok := es.Writer.(errorSet); ok {
		inputErrCh = v.ErrorSet()
	}

	var outputErrCh <-chan struct{}
	if v, ok := es.Reader.(errorSet); ok {
		outputErrCh = v.ErrorSet()
	}
	var outputClosedCh <-chan struct{}
	if v, ok := es.Reader.(interface{ Closed() <-chan struct{} }); ok {
		outputClosedCh = v.Closed()
	}

	select {
	case <-es.done:
	case <-inputErrCh:
		es.err.SetError(es.Writer.Err())
		es.Close()

	case <-outputErrCh:
		es.err.SetError(es.Reader.Err())
		es.Close()

	case <-outputClosedCh:
		if err := es.Reader.Err(); err != nil {
			es.err.SetError(es.Reader.Err())
		}
		es.Close()

	}
}
