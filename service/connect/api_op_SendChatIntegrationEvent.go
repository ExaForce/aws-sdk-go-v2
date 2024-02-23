// Code generated by smithy-go-codegen DO NOT EDIT.

package connect

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/connect/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Processes chat integration events from Amazon Web Services or external
// integrations to Amazon Connect. A chat integration event includes:
//   - SourceId, DestinationId, and Subtype: a set of identifiers, uniquely
//     representing a chat
//   - ChatEvent: details of the chat action to perform such as sending a message,
//     event, or disconnecting from a chat
//
// When a chat integration event is sent with chat identifiers that do not map to
// an active chat contact, a new chat contact is also created before handling chat
// action. Access to this API is currently restricted to Amazon Pinpoint for
// supporting SMS integration.
func (c *Client) SendChatIntegrationEvent(ctx context.Context, params *SendChatIntegrationEventInput, optFns ...func(*Options)) (*SendChatIntegrationEventOutput, error) {
	if params == nil {
		params = &SendChatIntegrationEventInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "SendChatIntegrationEvent", params, optFns, c.addOperationSendChatIntegrationEventMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*SendChatIntegrationEventOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type SendChatIntegrationEventInput struct {

	// Chat system identifier, used in part to uniquely identify chat. This is
	// associated with the Amazon Connect instance and flow to be used to start chats.
	// For SMS, this is the phone number destination of inbound SMS messages
	// represented by an Amazon Pinpoint phone number ARN.
	//
	// This member is required.
	DestinationId *string

	// Chat integration event payload
	//
	// This member is required.
	Event *types.ChatEvent

	// External identifier of chat customer participant, used in part to uniquely
	// identify a chat. For SMS, this is the E164 phone number of the chat customer
	// participant.
	//
	// This member is required.
	SourceId *string

	// Contact properties to apply when starting a new chat. If the integration event
	// is handled with an existing chat, this is ignored.
	NewSessionDetails *types.NewSessionDetails

	// Classification of a channel. This is used in part to uniquely identify chat.
	// Valid value: ["connect:sms"]
	Subtype *string

	noSmithyDocumentSerde
}

type SendChatIntegrationEventOutput struct {

	// Identifier of chat contact used to handle integration event. This may be null
	// if the integration event is not valid without an already existing chat contact.
	InitialContactId *string

	// Whether handling the integration event resulted in creating a new chat or
	// acting on existing chat.
	NewChatCreated *bool

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationSendChatIntegrationEventMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpSendChatIntegrationEvent{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpSendChatIntegrationEvent{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "SendChatIntegrationEvent"); err != nil {
		return fmt.Errorf("add protocol finalizers: %v", err)
	}

	if err = addlegacyEndpointContextSetter(stack, options); err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = addClientRequestID(stack); err != nil {
		return err
	}
	if err = addComputeContentLength(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addComputePayloadSHA256(stack); err != nil {
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
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addSetLegacyContextSigningOptionsMiddleware(stack); err != nil {
		return err
	}
	if err = addOpSendChatIntegrationEventValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opSendChatIntegrationEvent(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opSendChatIntegrationEvent(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "SendChatIntegrationEvent",
	}
}
