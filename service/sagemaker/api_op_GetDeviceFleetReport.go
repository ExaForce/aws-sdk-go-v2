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

// Describes a fleet.
func (c *Client) GetDeviceFleetReport(ctx context.Context, params *GetDeviceFleetReportInput, optFns ...func(*Options)) (*GetDeviceFleetReportOutput, error) {
	if params == nil {
		params = &GetDeviceFleetReportInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetDeviceFleetReport", params, optFns, c.addOperationGetDeviceFleetReportMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetDeviceFleetReportOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetDeviceFleetReportInput struct {

	// The name of the fleet.
	//
	// This member is required.
	DeviceFleetName *string

	noSmithyDocumentSerde
}

type GetDeviceFleetReportOutput struct {

	// The Amazon Resource Name (ARN) of the device.
	//
	// This member is required.
	DeviceFleetArn *string

	// The name of the fleet.
	//
	// This member is required.
	DeviceFleetName *string

	// The versions of Edge Manager agent deployed on the fleet.
	AgentVersions []types.AgentVersion

	// Description of the fleet.
	Description *string

	// Status of devices.
	DeviceStats *types.DeviceStats

	// Status of model on device.
	ModelStats []types.EdgeModelStat

	// The output configuration for storing sample data collected by the fleet.
	OutputConfig *types.EdgeOutputConfig

	// Timestamp of when the report was generated.
	ReportGenerated *time.Time

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetDeviceFleetReportMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetDeviceFleetReport{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetDeviceFleetReport{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetDeviceFleetReport"); err != nil {
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
	if err = addOpGetDeviceFleetReportValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetDeviceFleetReport(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetDeviceFleetReport(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetDeviceFleetReport",
	}
}
