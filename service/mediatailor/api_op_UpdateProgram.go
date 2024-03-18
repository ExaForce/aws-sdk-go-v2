// Code generated by smithy-go-codegen DO NOT EDIT.

package mediatailor

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/mediatailor/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Updates a program within a channel.
func (c *Client) UpdateProgram(ctx context.Context, params *UpdateProgramInput, optFns ...func(*Options)) (*UpdateProgramOutput, error) {
	if params == nil {
		params = &UpdateProgramInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateProgram", params, optFns, c.addOperationUpdateProgramMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateProgramOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateProgramInput struct {

	// The name of the channel for this Program.
	//
	// This member is required.
	ChannelName *string

	// The name of the Program.
	//
	// This member is required.
	ProgramName *string

	// The schedule configuration settings.
	//
	// This member is required.
	ScheduleConfiguration *types.UpdateProgramScheduleConfiguration

	// The ad break configuration settings.
	AdBreaks []types.AdBreak

	// The list of AudienceMedia defined in program.
	AudienceMedia []types.AudienceMedia

	noSmithyDocumentSerde
}

type UpdateProgramOutput struct {

	// The ad break configuration settings.
	AdBreaks []types.AdBreak

	// The ARN to assign to the program.
	Arn *string

	// The list of AudienceMedia defined in program.
	AudienceMedia []types.AudienceMedia

	// The name to assign to the channel for this program.
	ChannelName *string

	// The clip range configuration settings.
	ClipRange *types.ClipRange

	// The time the program was created.
	CreationTime *time.Time

	// The duration of the live program in milliseconds.
	DurationMillis *int64

	// The name of the LiveSource for this Program.
	LiveSourceName *string

	// The name to assign to this program.
	ProgramName *string

	// The scheduled start time for this Program.
	ScheduledStartTime *time.Time

	// The name to assign to the source location for this program.
	SourceLocationName *string

	// The name that's used to refer to a VOD source.
	VodSourceName *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateProgramMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateProgram{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateProgram{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "UpdateProgram"); err != nil {
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
	if err = addOpUpdateProgramValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateProgram(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateProgram(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "UpdateProgram",
	}
}
