// Code generated by smithy-go-codegen DO NOT EDIT.

package appstream

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Creates a URL to start a create app block builder streaming session.
func (c *Client) CreateAppBlockBuilderStreamingURL(ctx context.Context, params *CreateAppBlockBuilderStreamingURLInput, optFns ...func(*Options)) (*CreateAppBlockBuilderStreamingURLOutput, error) {
	if params == nil {
		params = &CreateAppBlockBuilderStreamingURLInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateAppBlockBuilderStreamingURL", params, optFns, c.addOperationCreateAppBlockBuilderStreamingURLMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateAppBlockBuilderStreamingURLOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateAppBlockBuilderStreamingURLInput struct {

	// The name of the app block builder.
	//
	// This member is required.
	AppBlockBuilderName *string

	// The time that the streaming URL will be valid, in seconds. Specify a value
	// between 1 and 604800 seconds. The default is 3600 seconds.
	Validity *int64

	noSmithyDocumentSerde
}

type CreateAppBlockBuilderStreamingURLOutput struct {

	// The elapsed time, in seconds after the Unix epoch, when this URL expires.
	Expires *time.Time

	// The URL to start the streaming session.
	StreamingURL *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateAppBlockBuilderStreamingURLMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateAppBlockBuilderStreamingURL{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateAppBlockBuilderStreamingURL{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateAppBlockBuilderStreamingURL"); err != nil {
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
	if err = addOpCreateAppBlockBuilderStreamingURLValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateAppBlockBuilderStreamingURL(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateAppBlockBuilderStreamingURL(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateAppBlockBuilderStreamingURL",
	}
}
