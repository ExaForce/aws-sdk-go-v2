// Code generated by smithy-go-codegen DO NOT EDIT.

package backupgateway

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/backupgateway/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// This action sets the bandwidth rate limit schedule for a specified gateway. By
// default, gateways do not have a bandwidth rate limit schedule, which means no
// bandwidth rate limiting is in effect. Use this to initiate a gateway's bandwidth
// rate limit schedule.
func (c *Client) PutBandwidthRateLimitSchedule(ctx context.Context, params *PutBandwidthRateLimitScheduleInput, optFns ...func(*Options)) (*PutBandwidthRateLimitScheduleOutput, error) {
	if params == nil {
		params = &PutBandwidthRateLimitScheduleInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutBandwidthRateLimitSchedule", params, optFns, c.addOperationPutBandwidthRateLimitScheduleMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutBandwidthRateLimitScheduleOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutBandwidthRateLimitScheduleInput struct {

	// An array containing bandwidth rate limit schedule intervals for a gateway. When
	// no bandwidth rate limit intervals have been scheduled, the array is empty.
	//
	// This member is required.
	BandwidthRateLimitIntervals []types.BandwidthRateLimitInterval

	// The Amazon Resource Name (ARN) of the gateway. Use the ListGateways (https://docs.aws.amazon.com/aws-backup/latest/devguide/API_BGW_ListGateways.html)
	// operation to return a list of gateways for your account and Amazon Web Services
	// Region.
	//
	// This member is required.
	GatewayArn *string

	noSmithyDocumentSerde
}

type PutBandwidthRateLimitScheduleOutput struct {

	// The Amazon Resource Name (ARN) of the gateway. Use the ListGateways (https://docs.aws.amazon.com/aws-backup/latest/devguide/API_BGW_ListGateways.html)
	// operation to return a list of gateways for your account and Amazon Web Services
	// Region.
	GatewayArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationPutBandwidthRateLimitScheduleMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpPutBandwidthRateLimitSchedule{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpPutBandwidthRateLimitSchedule{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "PutBandwidthRateLimitSchedule"); err != nil {
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
	if err = addOpPutBandwidthRateLimitScheduleValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opPutBandwidthRateLimitSchedule(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opPutBandwidthRateLimitSchedule(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "PutBandwidthRateLimitSchedule",
	}
}
