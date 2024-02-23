// Code generated by smithy-go-codegen DO NOT EDIT.

package arczonalshift

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/arczonalshift/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Cancel a zonal shift in Amazon Route 53 Application Recovery Controller. To
// cancel the zonal shift, specify the zonal shift ID. A zonal shift can be one
// that you've started for a resource in your Amazon Web Services account in an
// Amazon Web Services Region, or it can be a zonal shift started by a practice run
// with zonal autoshift.
func (c *Client) CancelZonalShift(ctx context.Context, params *CancelZonalShiftInput, optFns ...func(*Options)) (*CancelZonalShiftOutput, error) {
	if params == nil {
		params = &CancelZonalShiftInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CancelZonalShift", params, optFns, c.addOperationCancelZonalShiftMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CancelZonalShiftOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CancelZonalShiftInput struct {

	// The internally-generated identifier of a zonal shift.
	//
	// This member is required.
	ZonalShiftId *string

	noSmithyDocumentSerde
}

type CancelZonalShiftOutput struct {

	// The Availability Zone that traffic is moved away from for a resource when you
	// start a zonal shift. Until the zonal shift expires or you cancel it, traffic for
	// the resource is instead moved to other Availability Zones in the Amazon Web
	// Services Region.
	//
	// This member is required.
	AwayFrom *string

	// A comment that you enter about the zonal shift. Only the latest comment is
	// retained; no comment history is maintained. A new comment overwrites any
	// existing comment string.
	//
	// This member is required.
	Comment *string

	// The expiry time (expiration time) for a customer-started zonal shift. A zonal
	// shift is temporary and must be set to expire when you start the zonal shift. You
	// can initially set a zonal shift to expire in a maximum of three days (72 hours).
	// However, you can update a zonal shift to set a new expiration at any time. When
	// you start a zonal shift, you specify how long you want it to be active, which
	// Route 53 ARC converts to an expiry time (expiration time). You can cancel a
	// zonal shift when you're ready to restore traffic to the Availability Zone, or
	// just wait for it to expire. Or you can update the zonal shift to specify another
	// length of time to expire in.
	//
	// This member is required.
	ExpiryTime *time.Time

	// The identifier for the resource to shift away traffic for. The identifier is
	// the Amazon Resource Name (ARN) for the resource. At this time, supported
	// resources are Network Load Balancers and Application Load Balancers with
	// cross-zone load balancing turned off.
	//
	// This member is required.
	ResourceIdentifier *string

	// The time (UTC) when the zonal shift starts.
	//
	// This member is required.
	StartTime *time.Time

	// A status for a zonal shift. The Status for a zonal shift can have one of the
	// following values:
	//   - ACTIVE: The zonal shift has been started and active.
	//   - EXPIRED: The zonal shift has expired (the expiry time was exceeded).
	//   - CANCELED: The zonal shift was canceled.
	//
	// This member is required.
	Status types.ZonalShiftStatus

	// The identifier of a zonal shift.
	//
	// This member is required.
	ZonalShiftId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCancelZonalShiftMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCancelZonalShift{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCancelZonalShift{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CancelZonalShift"); err != nil {
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
	if err = addOpCancelZonalShiftValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCancelZonalShift(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCancelZonalShift(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CancelZonalShift",
	}
}
