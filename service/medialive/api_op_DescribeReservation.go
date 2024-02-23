// Code generated by smithy-go-codegen DO NOT EDIT.

package medialive

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/medialive/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Get details for a reservation.
func (c *Client) DescribeReservation(ctx context.Context, params *DescribeReservationInput, optFns ...func(*Options)) (*DescribeReservationOutput, error) {
	if params == nil {
		params = &DescribeReservationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeReservation", params, optFns, c.addOperationDescribeReservationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeReservationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Placeholder documentation for DescribeReservationRequest
type DescribeReservationInput struct {

	// Unique reservation ID, e.g. '1234567'
	//
	// This member is required.
	ReservationId *string

	noSmithyDocumentSerde
}

// Placeholder documentation for DescribeReservationResponse
type DescribeReservationOutput struct {

	// Unique reservation ARN, e.g.
	// 'arn:aws:medialive:us-west-2:123456789012:reservation:1234567'
	Arn *string

	// Number of reserved resources
	Count *int32

	// Currency code for usagePrice and fixedPrice in ISO-4217 format, e.g. 'USD'
	CurrencyCode *string

	// Lease duration, e.g. '12'
	Duration *int32

	// Units for duration, e.g. 'MONTHS'
	DurationUnits types.OfferingDurationUnits

	// Reservation UTC end date and time in ISO-8601 format, e.g. '2019-03-01T00:00:00'
	End *string

	// One-time charge for each reserved resource, e.g. '0.0' for a NO_UPFRONT offering
	FixedPrice *float64

	// User specified reservation name
	Name *string

	// Offering description, e.g. 'HD AVC output at 10-20 Mbps, 30 fps, and standard
	// VQ in US West (Oregon)'
	OfferingDescription *string

	// Unique offering ID, e.g. '87654321'
	OfferingId *string

	// Offering type, e.g. 'NO_UPFRONT'
	OfferingType types.OfferingType

	// AWS region, e.g. 'us-west-2'
	Region *string

	// Renewal settings for the reservation
	RenewalSettings *types.RenewalSettings

	// Unique reservation ID, e.g. '1234567'
	ReservationId *string

	// Resource configuration details
	ResourceSpecification *types.ReservationResourceSpecification

	// Reservation UTC start date and time in ISO-8601 format, e.g.
	// '2018-03-01T00:00:00'
	Start *string

	// Current state of reservation, e.g. 'ACTIVE'
	State types.ReservationState

	// A collection of key-value pairs
	Tags map[string]string

	// Recurring usage charge for each reserved resource, e.g. '157.0'
	UsagePrice *float64

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeReservationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribeReservation{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeReservation{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeReservation"); err != nil {
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
	if err = addOpDescribeReservationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeReservation(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeReservation(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeReservation",
	}
}
