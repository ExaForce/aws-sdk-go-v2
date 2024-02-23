// Code generated by smithy-go-codegen DO NOT EDIT.

package pinpointsmsvoicev2

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/pinpointsmsvoicev2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Allows you to send a request that sends a voice message through Amazon
// Pinpoint. This operation uses Amazon Polly (http://aws.amazon.com/polly/) to
// convert a text script into a voice message.
func (c *Client) SendVoiceMessage(ctx context.Context, params *SendVoiceMessageInput, optFns ...func(*Options)) (*SendVoiceMessageOutput, error) {
	if params == nil {
		params = &SendVoiceMessageInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "SendVoiceMessage", params, optFns, c.addOperationSendVoiceMessageMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*SendVoiceMessageOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type SendVoiceMessageInput struct {

	// The destination phone number in E.164 format.
	//
	// This member is required.
	DestinationPhoneNumber *string

	// The origination identity to use for the voice call. This can be the
	// PhoneNumber, PhoneNumberId, PhoneNumberArn, PoolId, or PoolArn.
	//
	// This member is required.
	OriginationIdentity *string

	// The name of the configuration set to use. This can be either the
	// ConfigurationSetName or ConfigurationSetArn.
	ConfigurationSetName *string

	// You can specify custom data in this field. If you do, that data is logged to
	// the event destination.
	Context map[string]string

	// When set to true, the message is checked and validated, but isn't sent to the
	// end recipient.
	DryRun bool

	// The maximum amount to spend per voice message, in US dollars.
	MaxPricePerMinute *string

	// The text to convert to a voice message.
	MessageBody *string

	// Specifies if the MessageBody field contains text or speech synthesis markup
	// language (SSML) (https://docs.aws.amazon.com/polly/latest/dg/what-is.html) .
	//   - TEXT: This is the default value. When used the maximum character limit is
	//   3000.
	//   - SSML: When used the maximum character limit is 6000 including SSML tagging.
	MessageBodyTextType types.VoiceMessageBodyTextType

	// How long the voice message is valid for. By default this is 72 hours.
	TimeToLive *int32

	// The voice for the Amazon Polly (https://docs.aws.amazon.com/polly/latest/dg/what-is.html)
	// service to use. By default this is set to "MATTHEW".
	VoiceId types.VoiceId

	noSmithyDocumentSerde
}

type SendVoiceMessageOutput struct {

	// The unique identifier for the message.
	MessageId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationSendVoiceMessageMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpSendVoiceMessage{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpSendVoiceMessage{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "SendVoiceMessage"); err != nil {
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
	if err = addOpSendVoiceMessageValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opSendVoiceMessage(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opSendVoiceMessage(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "SendVoiceMessage",
	}
}
