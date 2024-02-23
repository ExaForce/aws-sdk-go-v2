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

// Retrieves information of a SageMaker HyperPod cluster.
func (c *Client) DescribeCluster(ctx context.Context, params *DescribeClusterInput, optFns ...func(*Options)) (*DescribeClusterOutput, error) {
	if params == nil {
		params = &DescribeClusterInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeCluster", params, optFns, c.addOperationDescribeClusterMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeClusterOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeClusterInput struct {

	// The string name or the Amazon Resource Name (ARN) of the SageMaker HyperPod
	// cluster.
	//
	// This member is required.
	ClusterName *string

	noSmithyDocumentSerde
}

type DescribeClusterOutput struct {

	// The Amazon Resource Name (ARN) of the SageMaker HyperPod cluster.
	//
	// This member is required.
	ClusterArn *string

	// The status of the SageMaker HyperPod cluster.
	//
	// This member is required.
	ClusterStatus types.ClusterStatus

	// The instance groups of the SageMaker HyperPod cluster.
	//
	// This member is required.
	InstanceGroups []types.ClusterInstanceGroupDetails

	// The name of the SageMaker HyperPod cluster.
	ClusterName *string

	// The time when the SageMaker Cluster is created.
	CreationTime *time.Time

	// The failure message of the SageMaker HyperPod cluster.
	FailureMessage *string

	// Specifies an Amazon Virtual Private Cloud (VPC) that your SageMaker jobs,
	// hosted models, and compute resources have access to. You can control access to
	// and from your resources by configuring a VPC. For more information, see Give
	// SageMaker Access to Resources in your Amazon VPC (https://docs.aws.amazon.com/sagemaker/latest/dg/infrastructure-give-access.html)
	// .
	VpcConfig *types.VpcConfig

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeClusterMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeCluster{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeCluster{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeCluster"); err != nil {
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
	if err = addOpDescribeClusterValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeCluster(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeCluster(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeCluster",
	}
}
