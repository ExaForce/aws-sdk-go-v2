// Code generated by smithy-go-codegen DO NOT EDIT.

package iotjobsdataplane

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/iotjobsdataplane/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates the status of a job execution.
func (c *Client) UpdateJobExecution(ctx context.Context, params *UpdateJobExecutionInput, optFns ...func(*Options)) (*UpdateJobExecutionOutput, error) {
	if params == nil {
		params = &UpdateJobExecutionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateJobExecution", params, optFns, c.addOperationUpdateJobExecutionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateJobExecutionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateJobExecutionInput struct {

	// The unique identifier assigned to this job when it was created.
	//
	// This member is required.
	JobId *string

	// The new status for the job execution (IN_PROGRESS, FAILED, SUCCESS, or
	// REJECTED). This must be specified on every update.
	//
	// This member is required.
	Status types.JobExecutionStatus

	// The name of the thing associated with the device.
	//
	// This member is required.
	ThingName *string

	// Optional. A number that identifies a particular job execution on a particular
	// device.
	ExecutionNumber *int64

	// Optional. The expected current version of the job execution. Each time you
	// update the job execution, its version is incremented. If the version of the job
	// execution stored in Jobs does not match, the update is rejected with a
	// VersionMismatch error, and an ErrorResponse that contains the current job
	// execution status data is returned. (This makes it unnecessary to perform a
	// separate DescribeJobExecution request in order to obtain the job execution
	// status data.)
	ExpectedVersion *int64

	// Optional. When set to true, the response contains the job document. The default
	// is false.
	IncludeJobDocument *bool

	// Optional. When included and set to true, the response contains the
	// JobExecutionState data. The default is false.
	IncludeJobExecutionState *bool

	// Optional. A collection of name/value pairs that describe the status of the job
	// execution. If not specified, the statusDetails are unchanged.
	StatusDetails map[string]string

	// Specifies the amount of time this device has to finish execution of this job.
	// If the job execution status is not set to a terminal state before this timer
	// expires, or before the timer is reset (by again calling UpdateJobExecution ,
	// setting the status to IN_PROGRESS and specifying a new timeout value in this
	// field) the job execution status will be automatically set to TIMED_OUT . Note
	// that setting or resetting this timeout has no effect on that job execution
	// timeout which may have been specified when the job was created ( CreateJob
	// using field timeoutConfig ).
	StepTimeoutInMinutes *int64

	noSmithyDocumentSerde
}

type UpdateJobExecutionOutput struct {

	// A JobExecutionState object.
	ExecutionState *types.JobExecutionState

	// The contents of the Job Documents.
	JobDocument *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateJobExecutionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateJobExecution{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateJobExecution{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "UpdateJobExecution"); err != nil {
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
	if err = addOpUpdateJobExecutionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateJobExecution(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateJobExecution(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "UpdateJobExecution",
	}
}
