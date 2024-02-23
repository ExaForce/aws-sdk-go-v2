// Code generated by smithy-go-codegen DO NOT EDIT.

package kinesisanalyticsv2

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/kinesisanalyticsv2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Adds an Amazon CloudWatch log stream to monitor application configuration
// errors.
func (c *Client) AddApplicationCloudWatchLoggingOption(ctx context.Context, params *AddApplicationCloudWatchLoggingOptionInput, optFns ...func(*Options)) (*AddApplicationCloudWatchLoggingOptionOutput, error) {
	if params == nil {
		params = &AddApplicationCloudWatchLoggingOptionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "AddApplicationCloudWatchLoggingOption", params, optFns, c.addOperationAddApplicationCloudWatchLoggingOptionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*AddApplicationCloudWatchLoggingOptionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type AddApplicationCloudWatchLoggingOptionInput struct {

	// The Kinesis Data Analytics application name.
	//
	// This member is required.
	ApplicationName *string

	// Provides the Amazon CloudWatch log stream Amazon Resource Name (ARN).
	//
	// This member is required.
	CloudWatchLoggingOption *types.CloudWatchLoggingOption

	// A value you use to implement strong concurrency for application updates. You
	// must provide the CurrentApplicationVersionId or the ConditionalToken . You get
	// the application's current ConditionalToken using DescribeApplication . For
	// better concurrency support, use the ConditionalToken parameter instead of
	// CurrentApplicationVersionId .
	ConditionalToken *string

	// The version ID of the Kinesis Data Analytics application. You must provide the
	// CurrentApplicationVersionId or the ConditionalToken .You can retrieve the
	// application version ID using DescribeApplication . For better concurrency
	// support, use the ConditionalToken parameter instead of
	// CurrentApplicationVersionId .
	CurrentApplicationVersionId *int64

	noSmithyDocumentSerde
}

type AddApplicationCloudWatchLoggingOptionOutput struct {

	// The application's ARN.
	ApplicationARN *string

	// The new version ID of the Kinesis Data Analytics application. Kinesis Data
	// Analytics updates the ApplicationVersionId each time you change the CloudWatch
	// logging options.
	ApplicationVersionId *int64

	// The descriptions of the current CloudWatch logging options for the Kinesis Data
	// Analytics application.
	CloudWatchLoggingOptionDescriptions []types.CloudWatchLoggingOptionDescription

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationAddApplicationCloudWatchLoggingOptionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpAddApplicationCloudWatchLoggingOption{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpAddApplicationCloudWatchLoggingOption{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "AddApplicationCloudWatchLoggingOption"); err != nil {
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
	if err = addOpAddApplicationCloudWatchLoggingOptionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opAddApplicationCloudWatchLoggingOption(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opAddApplicationCloudWatchLoggingOption(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "AddApplicationCloudWatchLoggingOption",
	}
}
