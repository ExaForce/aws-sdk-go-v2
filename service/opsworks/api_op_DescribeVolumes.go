// Code generated by smithy-go-codegen DO NOT EDIT.

package opsworks

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/opsworks/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Describes an instance's Amazon EBS volumes. This call accepts only one
// resource-identifying parameter. Required Permissions: To use this action, an IAM
// user must have a Show, Deploy, or Manage permissions level for the stack, or an
// attached policy that explicitly grants permissions. For more information about
// user permissions, see Managing User Permissions (https://docs.aws.amazon.com/opsworks/latest/userguide/opsworks-security-users.html)
// .
func (c *Client) DescribeVolumes(ctx context.Context, params *DescribeVolumesInput, optFns ...func(*Options)) (*DescribeVolumesOutput, error) {
	if params == nil {
		params = &DescribeVolumesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeVolumes", params, optFns, c.addOperationDescribeVolumesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeVolumesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeVolumesInput struct {

	// The instance ID. If you use this parameter, DescribeVolumes returns
	// descriptions of the volumes associated with the specified instance.
	InstanceId *string

	// The RAID array ID. If you use this parameter, DescribeVolumes returns
	// descriptions of the volumes associated with the specified RAID array.
	RaidArrayId *string

	// A stack ID. The action describes the stack's registered Amazon EBS volumes.
	StackId *string

	// Am array of volume IDs. If you use this parameter, DescribeVolumes returns
	// descriptions of the specified volumes. Otherwise, it returns a description of
	// every volume.
	VolumeIds []string

	noSmithyDocumentSerde
}

// Contains the response to a DescribeVolumes request.
type DescribeVolumesOutput struct {

	// An array of volume IDs.
	Volumes []types.Volume

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeVolumesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeVolumes{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeVolumes{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeVolumes"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeVolumes(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeVolumes(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeVolumes",
	}
}
