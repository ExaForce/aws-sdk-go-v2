// Code generated by smithy-go-codegen DO NOT EDIT.

package kinesisvideo

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/kinesisvideo/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Gets the ImageGenerationConfiguration for a given Kinesis video stream.
func (c *Client) DescribeImageGenerationConfiguration(ctx context.Context, params *DescribeImageGenerationConfigurationInput, optFns ...func(*Options)) (*DescribeImageGenerationConfigurationOutput, error) {
	if params == nil {
		params = &DescribeImageGenerationConfigurationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeImageGenerationConfiguration", params, optFns, c.addOperationDescribeImageGenerationConfigurationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeImageGenerationConfigurationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeImageGenerationConfigurationInput struct {

	// The Amazon Resource Name (ARN) of the Kinesis video stream from which to
	// retrieve the image generation configuration. You must specify either the
	// StreamName or the StreamARN .
	StreamARN *string

	// The name of the stream from which to retrieve the image generation
	// configuration. You must specify either the StreamName or the StreamARN .
	StreamName *string

	noSmithyDocumentSerde
}

type DescribeImageGenerationConfigurationOutput struct {

	// The structure that contains the information required for the Kinesis video
	// stream (KVS) images delivery. If this structure is null, the configuration will
	// be deleted from the stream.
	ImageGenerationConfiguration *types.ImageGenerationConfiguration

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeImageGenerationConfigurationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribeImageGenerationConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeImageGenerationConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeImageGenerationConfiguration"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeImageGenerationConfiguration(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeImageGenerationConfiguration(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeImageGenerationConfiguration",
	}
}
