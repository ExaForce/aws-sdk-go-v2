// Code generated by smithy-go-codegen DO NOT EDIT.

package deadline

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/deadline/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates a step.
func (c *Client) UpdateStep(ctx context.Context, params *UpdateStepInput, optFns ...func(*Options)) (*UpdateStepOutput, error) {
	if params == nil {
		params = &UpdateStepInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateStep", params, optFns, c.addOperationUpdateStepMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateStepOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateStepInput struct {

	// The farm ID to update.
	//
	// This member is required.
	FarmId *string

	// The job ID to update.
	//
	// This member is required.
	JobId *string

	// The queue ID to update.
	//
	// This member is required.
	QueueId *string

	// The step ID to update.
	//
	// This member is required.
	StepId *string

	// The task status to update the step's tasks to.
	//
	// This member is required.
	TargetTaskRunStatus types.StepTargetTaskRunStatus

	// The unique token which the server uses to recognize retries of the same request.
	ClientToken *string

	noSmithyDocumentSerde
}

type UpdateStepOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateStepMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateStep{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateStep{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "UpdateStep"); err != nil {
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
	if err = addEndpointPrefix_opUpdateStepMiddleware(stack); err != nil {
		return err
	}
	if err = addIdempotencyToken_opUpdateStepMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpUpdateStepValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateStep(options.Region), middleware.Before); err != nil {
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

type endpointPrefix_opUpdateStepMiddleware struct {
}

func (*endpointPrefix_opUpdateStepMiddleware) ID() string {
	return "EndpointHostPrefix"
}

func (m *endpointPrefix_opUpdateStepMiddleware) HandleFinalize(ctx context.Context, in middleware.FinalizeInput, next middleware.FinalizeHandler) (
	out middleware.FinalizeOutput, metadata middleware.Metadata, err error,
) {
	if smithyhttp.GetHostnameImmutable(ctx) || smithyhttp.IsEndpointHostPrefixDisabled(ctx) {
		return next.HandleFinalize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	req.URL.Host = "management." + req.URL.Host

	return next.HandleFinalize(ctx, in)
}
func addEndpointPrefix_opUpdateStepMiddleware(stack *middleware.Stack) error {
	return stack.Finalize.Insert(&endpointPrefix_opUpdateStepMiddleware{}, "ResolveEndpointV2", middleware.After)
}

type idempotencyToken_initializeOpUpdateStep struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpUpdateStep) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpUpdateStep) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*UpdateStepInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *UpdateStepInput ")
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
func addIdempotencyToken_opUpdateStepMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpUpdateStep{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opUpdateStep(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "UpdateStep",
	}
}
