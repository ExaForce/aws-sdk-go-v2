// Code generated by smithy-go-codegen DO NOT EDIT.

package sagemaker

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/sagemaker/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Notifies the pipeline that the execution of a callback step succeeded and
// provides a list of the step's output parameters. When a callback step is run,
// the pipeline generates a callback token and includes the token in a message sent
// to Amazon Simple Queue Service (Amazon SQS).
func (c *Client) SendPipelineExecutionStepSuccess(ctx context.Context, params *SendPipelineExecutionStepSuccessInput, optFns ...func(*Options)) (*SendPipelineExecutionStepSuccessOutput, error) {
	if params == nil {
		params = &SendPipelineExecutionStepSuccessInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "SendPipelineExecutionStepSuccess", params, optFns, c.addOperationSendPipelineExecutionStepSuccessMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*SendPipelineExecutionStepSuccessOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type SendPipelineExecutionStepSuccessInput struct {

	// The pipeline generated token from the Amazon SQS queue.
	//
	// This member is required.
	CallbackToken *string

	// A unique, case-sensitive identifier that you provide to ensure the idempotency
	// of the operation. An idempotent operation completes no more than one time.
	ClientRequestToken *string

	// A list of the output parameters of the callback step.
	OutputParameters []types.OutputParameter

	noSmithyDocumentSerde
}

type SendPipelineExecutionStepSuccessOutput struct {

	// The Amazon Resource Name (ARN) of the pipeline execution.
	PipelineExecutionArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationSendPipelineExecutionStepSuccessMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpSendPipelineExecutionStepSuccess{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpSendPipelineExecutionStepSuccess{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "SendPipelineExecutionStepSuccess"); err != nil {
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
	if err = addIdempotencyToken_opSendPipelineExecutionStepSuccessMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpSendPipelineExecutionStepSuccessValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opSendPipelineExecutionStepSuccess(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpSendPipelineExecutionStepSuccess struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpSendPipelineExecutionStepSuccess) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpSendPipelineExecutionStepSuccess) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*SendPipelineExecutionStepSuccessInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *SendPipelineExecutionStepSuccessInput ")
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
func addIdempotencyToken_opSendPipelineExecutionStepSuccessMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpSendPipelineExecutionStepSuccess{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opSendPipelineExecutionStepSuccess(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "SendPipelineExecutionStepSuccess",
	}
}
