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

// A description of the fleet the device belongs to.
func (c *Client) DescribeDeviceFleet(ctx context.Context, params *DescribeDeviceFleetInput, optFns ...func(*Options)) (*DescribeDeviceFleetOutput, error) {
	if params == nil {
		params = &DescribeDeviceFleetInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeDeviceFleet", params, optFns, c.addOperationDescribeDeviceFleetMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeDeviceFleetOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeDeviceFleetInput struct {

	// The name of the fleet.
	//
	// This member is required.
	DeviceFleetName *string

	noSmithyDocumentSerde
}

type DescribeDeviceFleetOutput struct {

	// Timestamp of when the device fleet was created.
	//
	// This member is required.
	CreationTime *time.Time

	// The The Amazon Resource Name (ARN) of the fleet.
	//
	// This member is required.
	DeviceFleetArn *string

	// The name of the fleet.
	//
	// This member is required.
	DeviceFleetName *string

	// Timestamp of when the device fleet was last updated.
	//
	// This member is required.
	LastModifiedTime *time.Time

	// The output configuration for storing sampled data.
	//
	// This member is required.
	OutputConfig *types.EdgeOutputConfig

	// A description of the fleet.
	Description *string

	// The Amazon Resource Name (ARN) alias created in Amazon Web Services Internet of
	// Things (IoT).
	IotRoleAlias *string

	// The Amazon Resource Name (ARN) that has access to Amazon Web Services Internet
	// of Things (IoT).
	RoleArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeDeviceFleetMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeDeviceFleet{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeDeviceFleet{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeDeviceFleet"); err != nil {
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
	if err = addOpDescribeDeviceFleetValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeDeviceFleet(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeDeviceFleet(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeDeviceFleet",
	}
}
