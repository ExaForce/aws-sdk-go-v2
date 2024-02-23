// Code generated by smithy-go-codegen DO NOT EDIT.

package lexruntimev2

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/lexruntimev2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"io"
)

// Creates a new session or modifies an existing session with an Amazon Lex V2
// bot. Use this operation to enable your application to set the state of the bot.
func (c *Client) PutSession(ctx context.Context, params *PutSessionInput, optFns ...func(*Options)) (*PutSessionOutput, error) {
	if params == nil {
		params = &PutSessionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutSession", params, optFns, c.addOperationPutSessionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutSessionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutSessionInput struct {

	// The alias identifier of the bot that receives the session data.
	//
	// This member is required.
	BotAliasId *string

	// The identifier of the bot that receives the session data.
	//
	// This member is required.
	BotId *string

	// The locale where the session is in use.
	//
	// This member is required.
	LocaleId *string

	// The identifier of the session that receives the session data.
	//
	// This member is required.
	SessionId *string

	// Sets the state of the session with the user. You can use this to set the
	// current intent, attributes, context, and dialog action. Use the dialog action to
	// determine the next step that Amazon Lex V2 should use in the conversation with
	// the user.
	//
	// This member is required.
	SessionState *types.SessionState

	// A list of messages to send to the user. Messages are sent in the order that
	// they are defined in the list.
	Messages []types.Message

	// Request-specific information passed between Amazon Lex V2 and the client
	// application. The namespace x-amz-lex: is reserved for special attributes. Don't
	// create any request attributes with the prefix x-amz-lex: .
	RequestAttributes map[string]string

	// The message that Amazon Lex V2 returns in the response can be either text or
	// speech depending on the value of this parameter.
	//   - If the value is text/plain; charset=utf-8 , Amazon Lex V2 returns text in
	//   the response.
	ResponseContentType *string

	noSmithyDocumentSerde
}

type PutSessionOutput struct {

	// If the requested content type was audio, the audio version of the message to
	// convey to the user.
	AudioStream io.ReadCloser

	// The type of response. Same as the type specified in the responseContentType
	// field in the request.
	ContentType *string

	// A list of messages that were last sent to the user. The messages are ordered
	// based on how you return the messages from you Lambda function or the order that
	// the messages are defined in the bot.
	Messages *string

	// A base-64-encoded gzipped field that provides request-specific information
	// passed between the client application and Amazon Lex V2. These are the same as
	// the requestAttribute parameter in the call to the PutSession operation.
	RequestAttributes *string

	// The identifier of the session that received the data.
	SessionId *string

	// A base-64-encoded gzipped field that represents the current state of the dialog
	// between the user and the bot. Use this to determine the progress of the
	// conversation and what the next action may be.
	SessionState *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationPutSessionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpPutSession{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpPutSession{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "PutSession"); err != nil {
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
	if err = addSetLegacyContextSigningOptionsMiddleware(stack); err != nil {
		return err
	}
	if err = addOpPutSessionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opPutSession(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opPutSession(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "PutSession",
	}
}
