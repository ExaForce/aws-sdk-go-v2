// Code generated by smithy-go-codegen DO NOT EDIT.

package sagemaker

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/sagemaker/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Returns information about an AutoML job created by calling CreateAutoMLJobV2 (https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_CreateAutoMLJobV2.html)
// or CreateAutoMLJob (https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_CreateAutoMLJob.html)
// .
func (c *Client) DescribeAutoMLJobV2(ctx context.Context, params *DescribeAutoMLJobV2Input, optFns ...func(*Options)) (*DescribeAutoMLJobV2Output, error) {
	if params == nil {
		params = &DescribeAutoMLJobV2Input{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeAutoMLJobV2", params, optFns, c.addOperationDescribeAutoMLJobV2Middlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeAutoMLJobV2Output)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeAutoMLJobV2Input struct {

	// Requests information about an AutoML job V2 using its unique name.
	//
	// This member is required.
	AutoMLJobName *string

	noSmithyDocumentSerde
}

type DescribeAutoMLJobV2Output struct {

	// Returns the Amazon Resource Name (ARN) of the AutoML job V2.
	//
	// This member is required.
	AutoMLJobArn *string

	// Returns an array of channel objects describing the input data and their
	// location.
	//
	// This member is required.
	AutoMLJobInputDataConfig []types.AutoMLJobChannel

	// Returns the name of the AutoML job V2.
	//
	// This member is required.
	AutoMLJobName *string

	// Returns the secondary status of the AutoML job V2.
	//
	// This member is required.
	AutoMLJobSecondaryStatus types.AutoMLJobSecondaryStatus

	// Returns the status of the AutoML job V2.
	//
	// This member is required.
	AutoMLJobStatus types.AutoMLJobStatus

	// Returns the creation time of the AutoML job V2.
	//
	// This member is required.
	CreationTime *time.Time

	// Returns the job's last modified time.
	//
	// This member is required.
	LastModifiedTime *time.Time

	// Returns the job's output data config.
	//
	// This member is required.
	OutputDataConfig *types.AutoMLOutputDataConfig

	// The ARN of the IAM role that has read permission to the input data location and
	// write permission to the output data location in Amazon S3.
	//
	// This member is required.
	RoleArn *string

	// The artifacts that are generated during an AutoML job.
	AutoMLJobArtifacts *types.AutoMLJobArtifacts

	// Returns the job's objective.
	AutoMLJobObjective *types.AutoMLJobObjective

	// Returns the configuration settings of the problem type set for the AutoML job
	// V2.
	AutoMLProblemTypeConfig types.AutoMLProblemTypeConfig

	// Returns the name of the problem type configuration set for the AutoML job V2.
	AutoMLProblemTypeConfigName types.AutoMLProblemTypeConfigName

	// Information about the candidate produced by an AutoML training job V2,
	// including its status, steps, and other properties.
	BestCandidate *types.AutoMLCandidate

	// Returns the configuration settings of how the data are split into train and
	// validation datasets.
	DataSplitConfig *types.AutoMLDataSplitConfig

	// Returns the end time of the AutoML job V2.
	EndTime *time.Time

	// Returns the reason for the failure of the AutoML job V2, when applicable.
	FailureReason *string

	// Indicates whether the model was deployed automatically to an endpoint and the
	// name of that endpoint if deployed automatically.
	ModelDeployConfig *types.ModelDeployConfig

	// Provides information about endpoint for the model deployment.
	ModelDeployResult *types.ModelDeployResult

	// Returns a list of reasons for partial failures within an AutoML job V2.
	PartialFailureReasons []types.AutoMLPartialFailureReason

	// Returns the resolved attributes used by the AutoML job V2.
	ResolvedAttributes *types.AutoMLResolvedAttributes

	// Returns the security configuration for traffic encryption or Amazon VPC
	// settings.
	SecurityConfig *types.AutoMLSecurityConfig

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeAutoMLJobV2Middlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeAutoMLJobV2{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeAutoMLJobV2{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeAutoMLJobV2"); err != nil {
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
	if err = addOpDescribeAutoMLJobV2ValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeAutoMLJobV2(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeAutoMLJobV2(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeAutoMLJobV2",
	}
}
