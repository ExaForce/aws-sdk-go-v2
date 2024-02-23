// Code generated by smithy-go-codegen DO NOT EDIT.

package chimesdkmessaging

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/chimesdkmessaging/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Sends a message to a particular channel that the member is a part of. The
// x-amz-chime-bearer request header is mandatory. Use the ARN of the
// AppInstanceUser or AppInstanceBot that makes the API call as the value in the
// header. Also, STANDARD messages can be up to 4KB in size and contain metadata.
// Metadata is arbitrary, and you can use it in a variety of ways, such as
// containing a link to an attachment. CONTROL messages are limited to 30 bytes
// and do not contain metadata.
func (c *Client) SendChannelMessage(ctx context.Context, params *SendChannelMessageInput, optFns ...func(*Options)) (*SendChannelMessageOutput, error) {
	if params == nil {
		params = &SendChannelMessageInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "SendChannelMessage", params, optFns, c.addOperationSendChannelMessageMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*SendChannelMessageOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type SendChannelMessageInput struct {

	// The ARN of the channel.
	//
	// This member is required.
	ChannelArn *string

	// The ARN of the AppInstanceUser or AppInstanceBot that makes the API call.
	//
	// This member is required.
	ChimeBearer *string

	// The Idempotency token for each client request.
	//
	// This member is required.
	ClientRequestToken *string

	// The content of the channel message.
	//
	// This member is required.
	Content *string

	// Boolean that controls whether the message is persisted on the back end.
	// Required.
	//
	// This member is required.
	Persistence types.ChannelMessagePersistenceType

	// The type of message, STANDARD or CONTROL . STANDARD messages can be up to 4KB
	// in size and contain metadata. Metadata is arbitrary, and you can use it in a
	// variety of ways, such as containing a link to an attachment. CONTROL messages
	// are limited to 30 bytes and do not contain metadata.
	//
	// This member is required.
	Type types.ChannelMessageType

	// The content type of the channel message.
	ContentType *string

	// The attributes for the message, used for message filtering along with a
	// FilterRule defined in the PushNotificationPreferences .
	MessageAttributes map[string]types.MessageAttributeValue

	// The optional metadata for each message.
	Metadata *string

	// The push notification configuration of the message.
	PushNotification *types.PushNotificationConfiguration

	// The ID of the SubChannel in the request.
	SubChannelId *string

	// The target of a message. Must be a member of the channel, such as another user,
	// a bot, or the sender. Only the target and the sender can view targeted messages.
	// Only users who can see targeted messages can take actions on them. However,
	// administrators can delete targeted messages that they can’t see.
	Target []types.Target

	noSmithyDocumentSerde
}

type SendChannelMessageOutput struct {

	// The ARN of the channel.
	ChannelArn *string

	// The ID string assigned to each message.
	MessageId *string

	// The status of the channel message.
	Status *types.ChannelMessageStatusStructure

	// The ID of the SubChannel in the response.
	SubChannelId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationSendChannelMessageMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpSendChannelMessage{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpSendChannelMessage{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "SendChannelMessage"); err != nil {
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
	if err = addIdempotencyToken_opSendChannelMessageMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpSendChannelMessageValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opSendChannelMessage(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpSendChannelMessage struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpSendChannelMessage) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpSendChannelMessage) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*SendChannelMessageInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *SendChannelMessageInput ")
	}

	if input.ClientRequestToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.ClientRequestToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opSendChannelMessageMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpSendChannelMessage{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opSendChannelMessage(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "SendChannelMessage",
	}
}
