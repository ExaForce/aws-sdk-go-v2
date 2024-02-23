// Code generated by smithy-go-codegen DO NOT EDIT.

package datasync

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/datasync/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Provides details about how an DataSync transfer location for an S3 bucket is
// configured.
func (c *Client) DescribeLocationS3(ctx context.Context, params *DescribeLocationS3Input, optFns ...func(*Options)) (*DescribeLocationS3Output, error) {
	if params == nil {
		params = &DescribeLocationS3Input{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeLocationS3", params, optFns, c.addOperationDescribeLocationS3Middlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeLocationS3Output)
	out.ResultMetadata = metadata
	return out, nil
}

// DescribeLocationS3Request
type DescribeLocationS3Input struct {

	// Specifies the Amazon Resource Name (ARN) of the Amazon S3 location.
	//
	// This member is required.
	LocationArn *string

	noSmithyDocumentSerde
}

// DescribeLocationS3Response
type DescribeLocationS3Output struct {

	// The ARNs of the DataSync agents deployed on your Outpost when using working
	// with Amazon S3 on Outposts. For more information, see Deploy your DataSync
	// agent on Outposts (https://docs.aws.amazon.com/datasync/latest/userguide/deploy-agents.html#outposts-agent)
	// .
	AgentArns []string

	// The time that the Amazon S3 location was created.
	CreationTime *time.Time

	// The ARN of the Amazon S3 location.
	LocationArn *string

	// The URL of the Amazon S3 location that was described.
	LocationUri *string

	// Specifies the Amazon Resource Name (ARN) of the Identity and Access Management
	// (IAM) role that DataSync uses to access your S3 bucket. For more information,
	// see Accessing S3 buckets (https://docs.aws.amazon.com/datasync/latest/userguide/create-s3-location.html#create-s3-location-access)
	// .
	S3Config *types.S3Config

	// When Amazon S3 is a destination location, this is the storage class that you
	// chose for your objects. Some storage classes have behaviors that can affect your
	// Amazon S3 storage costs. For more information, see Storage class considerations
	// with Amazon S3 transfers (https://docs.aws.amazon.com/datasync/latest/userguide/create-s3-location.html#using-storage-classes)
	// .
	S3StorageClass types.S3StorageClass

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeLocationS3Middlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeLocationS3{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeLocationS3{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeLocationS3"); err != nil {
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
	if err = addOpDescribeLocationS3ValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeLocationS3(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeLocationS3(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeLocationS3",
	}
}
