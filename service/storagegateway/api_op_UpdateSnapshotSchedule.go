// Code generated by smithy-go-codegen DO NOT EDIT.

package storagegateway

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/storagegateway/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates a snapshot schedule configured for a gateway volume. This operation is
// only supported in the cached volume and stored volume gateway types. The default
// snapshot schedule for volume is once every 24 hours, starting at the creation
// time of the volume. You can use this API to change the snapshot schedule
// configured for the volume. In the request you must identify the gateway volume
// whose snapshot schedule you want to update, and the schedule information,
// including when you want the snapshot to begin on a day and the frequency (in
// hours) of snapshots.
func (c *Client) UpdateSnapshotSchedule(ctx context.Context, params *UpdateSnapshotScheduleInput, optFns ...func(*Options)) (*UpdateSnapshotScheduleOutput, error) {
	if params == nil {
		params = &UpdateSnapshotScheduleInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateSnapshotSchedule", params, optFns, c.addOperationUpdateSnapshotScheduleMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateSnapshotScheduleOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// A JSON object containing one or more of the following fields:
//   - UpdateSnapshotScheduleInput$Description
//   - UpdateSnapshotScheduleInput$RecurrenceInHours
//   - UpdateSnapshotScheduleInput$StartAt
//   - UpdateSnapshotScheduleInput$VolumeARN
type UpdateSnapshotScheduleInput struct {

	// Frequency of snapshots. Specify the number of hours between snapshots.
	//
	// This member is required.
	RecurrenceInHours *int32

	// The hour of the day at which the snapshot schedule begins represented as hh,
	// where hh is the hour (0 to 23). The hour of the day is in the time zone of the
	// gateway.
	//
	// This member is required.
	StartAt *int32

	// The Amazon Resource Name (ARN) of the volume. Use the ListVolumes operation to
	// return a list of gateway volumes.
	//
	// This member is required.
	VolumeARN *string

	// Optional description of the snapshot that overwrites the existing description.
	Description *string

	// A list of up to 50 tags that can be assigned to a snapshot. Each tag is a
	// key-value pair. Valid characters for key and value are letters, spaces, and
	// numbers representable in UTF-8 format, and the following special characters: + -
	// = . _ : / @. The maximum length of a tag's key is 128 characters, and the
	// maximum length for a tag's value is 256.
	Tags []types.Tag

	noSmithyDocumentSerde
}

// A JSON object containing the Amazon Resource Name (ARN) of the updated storage
// volume.
type UpdateSnapshotScheduleOutput struct {

	// The Amazon Resource Name (ARN) of the volume. Use the ListVolumes operation to
	// return a list of gateway volumes.
	VolumeARN *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateSnapshotScheduleMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateSnapshotSchedule{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateSnapshotSchedule{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "UpdateSnapshotSchedule"); err != nil {
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
	if err = addOpUpdateSnapshotScheduleValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateSnapshotSchedule(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateSnapshotSchedule(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "UpdateSnapshotSchedule",
	}
}
