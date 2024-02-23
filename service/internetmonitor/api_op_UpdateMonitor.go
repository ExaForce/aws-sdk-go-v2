// Code generated by smithy-go-codegen DO NOT EDIT.

package internetmonitor

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/internetmonitor/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates a monitor. You can update a monitor to change the percentage of traffic
// to monitor or the maximum number of city-networks (locations and ASNs), to add
// or remove resources, or to change the status of the monitor. Note that you can't
// change the name of a monitor. The city-network maximum that you choose is the
// limit, but you only pay for the number of city-networks that are actually
// monitored. For more information, see Choosing a city-network maximum value (https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/IMCityNetworksMaximum.html)
// in the Amazon CloudWatch User Guide.
func (c *Client) UpdateMonitor(ctx context.Context, params *UpdateMonitorInput, optFns ...func(*Options)) (*UpdateMonitorOutput, error) {
	if params == nil {
		params = &UpdateMonitorInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateMonitor", params, optFns, c.addOperationUpdateMonitorMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateMonitorOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateMonitorInput struct {

	// The name of the monitor.
	//
	// This member is required.
	MonitorName *string

	// A unique, case-sensitive string of up to 64 ASCII characters that you specify
	// to make an idempotent API request. You should not reuse the same client token
	// for other API requests.
	ClientToken *string

	// The list of health score thresholds. A threshold percentage for health scores,
	// along with other configuration information, determines when Internet Monitor
	// creates a health event when there's an internet issue that affects your
	// application end users. For more information, see Change health event thresholds (https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch-IM-overview.html#IMUpdateThresholdFromOverview)
	// in the Internet Monitor section of the CloudWatch User Guide.
	HealthEventsConfig *types.HealthEventsConfig

	// Publish internet measurements for Internet Monitor to another location, such as
	// an Amazon S3 bucket. The measurements are also published to Amazon CloudWatch
	// Logs.
	InternetMeasurementsLogDelivery *types.InternetMeasurementsLogDelivery

	// The maximum number of city-networks to monitor for your application. A
	// city-network is the location (city) where clients access your application
	// resources from and the ASN or network provider, such as an internet service
	// provider (ISP), that clients access the resources through. Setting this limit
	// can help control billing costs.
	MaxCityNetworksToMonitor *int32

	// The resources to include in a monitor, which you provide as a set of Amazon
	// Resource Names (ARNs). Resources can be VPCs, NLBs, Amazon CloudFront
	// distributions, or Amazon WorkSpaces directories. You can add a combination of
	// VPCs and CloudFront distributions, or you can add WorkSpaces directories, or you
	// can add NLBs. You can't add NLBs or WorkSpaces directories together with any
	// other resources. If you add only Amazon Virtual Private Clouds resources, at
	// least one VPC must have an Internet Gateway attached to it, to make sure that it
	// has internet connectivity.
	ResourcesToAdd []string

	// The resources to remove from a monitor, which you provide as a set of Amazon
	// Resource Names (ARNs).
	ResourcesToRemove []string

	// The status for a monitor. The accepted values for Status with the UpdateMonitor
	// API call are the following: ACTIVE and INACTIVE . The following values are not
	// accepted: PENDING , and ERROR .
	Status types.MonitorConfigState

	// The percentage of the internet-facing traffic for your application that you
	// want to monitor with this monitor. If you set a city-networks maximum, that
	// limit overrides the traffic percentage that you set. To learn more, see
	// Choosing an application traffic percentage to monitor  (https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/IMTrafficPercentage.html)
	// in the Amazon CloudWatch Internet Monitor section of the CloudWatch User Guide.
	TrafficPercentageToMonitor *int32

	noSmithyDocumentSerde
}

type UpdateMonitorOutput struct {

	// The Amazon Resource Name (ARN) of the monitor.
	//
	// This member is required.
	MonitorArn *string

	// The status of a monitor.
	//
	// This member is required.
	Status types.MonitorConfigState

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateMonitorMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateMonitor{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateMonitor{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "UpdateMonitor"); err != nil {
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
	if err = addIdempotencyToken_opUpdateMonitorMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpUpdateMonitorValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateMonitor(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpUpdateMonitor struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpUpdateMonitor) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpUpdateMonitor) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*UpdateMonitorInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *UpdateMonitorInput ")
	}

	if input.ClientToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.ClientToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opUpdateMonitorMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpUpdateMonitor{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opUpdateMonitor(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "UpdateMonitor",
	}
}
