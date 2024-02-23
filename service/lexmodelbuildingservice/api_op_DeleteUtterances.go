// Code generated by smithy-go-codegen DO NOT EDIT.

package lexmodelbuildingservice

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Deletes stored utterances. Amazon Lex stores the utterances that users send to
// your bot. Utterances are stored for 15 days for use with the GetUtterancesView
// operation, and then stored indefinitely for use in improving the ability of your
// bot to respond to user input. Use the DeleteUtterances operation to manually
// delete stored utterances for a specific user. When you use the DeleteUtterances
// operation, utterances stored for improving your bot's ability to respond to user
// input are deleted immediately. Utterances stored for use with the
// GetUtterancesView operation are deleted after 15 days. This operation requires
// permissions for the lex:DeleteUtterances action.
func (c *Client) DeleteUtterances(ctx context.Context, params *DeleteUtterancesInput, optFns ...func(*Options)) (*DeleteUtterancesOutput, error) {
	if params == nil {
		params = &DeleteUtterancesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteUtterances", params, optFns, c.addOperationDeleteUtterancesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteUtterancesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteUtterancesInput struct {

	// The name of the bot that stored the utterances.
	//
	// This member is required.
	BotName *string

	// The unique identifier for the user that made the utterances. This is the user
	// ID that was sent in the PostContent (http://docs.aws.amazon.com/lex/latest/dg/API_runtime_PostContent.html)
	// or PostText (http://docs.aws.amazon.com/lex/latest/dg/API_runtime_PostText.html)
	// operation request that contained the utterance.
	//
	// This member is required.
	UserId *string

	noSmithyDocumentSerde
}

type DeleteUtterancesOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeleteUtterancesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDeleteUtterances{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDeleteUtterances{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DeleteUtterances"); err != nil {
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
	if err = addOpDeleteUtterancesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteUtterances(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeleteUtterances(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DeleteUtterances",
	}
}
