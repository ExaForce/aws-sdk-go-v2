// Code generated by smithy-go-codegen DO NOT EDIT.

package mediatailor

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/mediatailor/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Provides details about a specific video on demand (VOD) source in a specific
// source location.
func (c *Client) DescribeVodSource(ctx context.Context, params *DescribeVodSourceInput, optFns ...func(*Options)) (*DescribeVodSourceOutput, error) {
	if params == nil {
		params = &DescribeVodSourceInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeVodSource", params, optFns, c.addOperationDescribeVodSourceMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeVodSourceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeVodSourceInput struct {

	// The name of the source location associated with this VOD Source.
	//
	// This member is required.
	SourceLocationName *string

	// The name of the VOD Source.
	//
	// This member is required.
	VodSourceName *string

	noSmithyDocumentSerde
}

type DescribeVodSourceOutput struct {

	// The ad break opportunities within the VOD source.
	AdBreakOpportunities []types.AdBreakOpportunity

	// The ARN of the VOD source.
	Arn *string

	// The timestamp that indicates when the VOD source was created.
	CreationTime *time.Time

	// The HTTP package configurations.
	HttpPackageConfigurations []types.HttpPackageConfiguration

	// The last modified time of the VOD source.
	LastModifiedTime *time.Time

	// The name of the source location associated with the VOD source.
	SourceLocationName *string

	// The tags assigned to the VOD source. Tags are key-value pairs that you can
	// associate with Amazon resources to help with organization, access control, and
	// cost tracking. For more information, see Tagging AWS Elemental MediaTailor
	// Resources (https://docs.aws.amazon.com/mediatailor/latest/ug/tagging.html) .
	Tags map[string]string

	// The name of the VOD source.
	VodSourceName *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeVodSourceMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribeVodSource{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeVodSource{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeVodSource"); err != nil {
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
	if err = addOpDescribeVodSourceValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeVodSource(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeVodSource(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeVodSource",
	}
}
