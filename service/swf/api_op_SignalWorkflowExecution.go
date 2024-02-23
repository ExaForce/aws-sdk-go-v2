// Code generated by smithy-go-codegen DO NOT EDIT.

package swf

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Records a WorkflowExecutionSignaled event in the workflow execution history and
// creates a decision task for the workflow execution identified by the given
// domain, workflowId and runId. The event is recorded with the specified user
// defined signalName and input (if provided). If a runId isn't specified, then the
// WorkflowExecutionSignaled event is recorded in the history of the current open
// workflow with the matching workflowId in the domain. If the specified workflow
// execution isn't open, this method fails with UnknownResource . Access Control
// You can use IAM policies to control this action's access to Amazon SWF resources
// as follows:
//   - Use a Resource element with the domain name to limit the action to only
//     specified domains.
//   - Use an Action element to allow or deny permission to call this action.
//   - You cannot use an IAM policy to constrain this action's parameters.
//
// If the caller doesn't have sufficient permissions to invoke the action, or the
// parameter values fall outside the specified constraints, the action fails. The
// associated event attribute's cause parameter is set to OPERATION_NOT_PERMITTED .
// For details and example IAM policies, see Using IAM to Manage Access to Amazon
// SWF Workflows (https://docs.aws.amazon.com/amazonswf/latest/developerguide/swf-dev-iam.html)
// in the Amazon SWF Developer Guide.
func (c *Client) SignalWorkflowExecution(ctx context.Context, params *SignalWorkflowExecutionInput, optFns ...func(*Options)) (*SignalWorkflowExecutionOutput, error) {
	if params == nil {
		params = &SignalWorkflowExecutionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "SignalWorkflowExecution", params, optFns, c.addOperationSignalWorkflowExecutionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*SignalWorkflowExecutionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type SignalWorkflowExecutionInput struct {

	// The name of the domain containing the workflow execution to signal.
	//
	// This member is required.
	Domain *string

	// The name of the signal. This name must be meaningful to the target workflow.
	//
	// This member is required.
	SignalName *string

	// The workflowId of the workflow execution to signal.
	//
	// This member is required.
	WorkflowId *string

	// Data to attach to the WorkflowExecutionSignaled event in the target workflow
	// execution's history.
	Input *string

	// The runId of the workflow execution to signal.
	RunId *string

	noSmithyDocumentSerde
}

type SignalWorkflowExecutionOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationSignalWorkflowExecutionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpSignalWorkflowExecution{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpSignalWorkflowExecution{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "SignalWorkflowExecution"); err != nil {
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
	if err = addOpSignalWorkflowExecutionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opSignalWorkflowExecution(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opSignalWorkflowExecution(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "SignalWorkflowExecution",
	}
}
