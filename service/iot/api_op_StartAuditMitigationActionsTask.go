// Code generated by smithy-go-codegen DO NOT EDIT.

package iot

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/iot/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Starts a task that applies a set of mitigation actions to the specified target.
// Requires permission to access the StartAuditMitigationActionsTask (https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions)
// action.
func (c *Client) StartAuditMitigationActionsTask(ctx context.Context, params *StartAuditMitigationActionsTaskInput, optFns ...func(*Options)) (*StartAuditMitigationActionsTaskOutput, error) {
	if params == nil {
		params = &StartAuditMitigationActionsTaskInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "StartAuditMitigationActionsTask", params, optFns, c.addOperationStartAuditMitigationActionsTaskMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*StartAuditMitigationActionsTaskOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StartAuditMitigationActionsTaskInput struct {

	// For an audit check, specifies which mitigation actions to apply. Those actions
	// must be defined in your Amazon Web Services accounts.
	//
	// This member is required.
	AuditCheckToActionsMapping map[string][]string

	// Each audit mitigation task must have a unique client request token. If you try
	// to start a new task with the same token as a task that already exists, an
	// exception occurs. If you omit this value, a unique client request token is
	// generated automatically.
	//
	// This member is required.
	ClientRequestToken *string

	// Specifies the audit findings to which the mitigation actions are applied. You
	// can apply them to a type of audit check, to all findings from an audit, or to a
	// specific set of findings.
	//
	// This member is required.
	Target *types.AuditMitigationActionsTaskTarget

	// A unique identifier for the task. You can use this identifier to check the
	// status of the task or to cancel it.
	//
	// This member is required.
	TaskId *string

	noSmithyDocumentSerde
}

type StartAuditMitigationActionsTaskOutput struct {

	// The unique identifier for the audit mitigation task. This matches the taskId
	// that you specified in the request.
	TaskId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationStartAuditMitigationActionsTaskMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpStartAuditMitigationActionsTask{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpStartAuditMitigationActionsTask{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "StartAuditMitigationActionsTask"); err != nil {
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
	if err = addIdempotencyToken_opStartAuditMitigationActionsTaskMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpStartAuditMitigationActionsTaskValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opStartAuditMitigationActionsTask(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpStartAuditMitigationActionsTask struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpStartAuditMitigationActionsTask) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpStartAuditMitigationActionsTask) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*StartAuditMitigationActionsTaskInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *StartAuditMitigationActionsTaskInput ")
	}

	if input.ClientRequestToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.ClientRequestToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opStartAuditMitigationActionsTaskMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpStartAuditMitigationActionsTask{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opStartAuditMitigationActionsTask(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "StartAuditMitigationActionsTask",
	}
}
