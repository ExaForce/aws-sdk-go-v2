// Code generated by smithy-go-codegen DO NOT EDIT.

package appflow

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/appflow/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates an existing flow.
func (c *Client) UpdateFlow(ctx context.Context, params *UpdateFlowInput, optFns ...func(*Options)) (*UpdateFlowOutput, error) {
	if params == nil {
		params = &UpdateFlowInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateFlow", params, optFns, c.addOperationUpdateFlowMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateFlowOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateFlowInput struct {

	// The configuration that controls how Amazon AppFlow transfers data to the
	// destination connector.
	//
	// This member is required.
	DestinationFlowConfigList []types.DestinationFlowConfig

	// The specified name of the flow. Spaces are not allowed. Use underscores (_) or
	// hyphens (-) only.
	//
	// This member is required.
	FlowName *string

	// Contains information about the configuration of the source connector used in
	// the flow.
	//
	// This member is required.
	SourceFlowConfig *types.SourceFlowConfig

	// A list of tasks that Amazon AppFlow performs while transferring the data in the
	// flow run.
	//
	// This member is required.
	Tasks []types.Task

	// The trigger settings that determine how and when the flow runs.
	//
	// This member is required.
	TriggerConfig *types.TriggerConfig

	// The clientToken parameter is an idempotency token. It ensures that your
	// UpdateFlow request completes only once. You choose the value to pass. For
	// example, if you don't receive a response from your request, you can safely retry
	// the request with the same clientToken parameter value. If you omit a clientToken
	// value, the Amazon Web Services SDK that you are using inserts a value for you.
	// This way, the SDK can safely retry requests multiple times after a network
	// error. You must provide your own value for other use cases. If you specify input
	// parameters that differ from your first request, an error occurs. If you use a
	// different value for clientToken , Amazon AppFlow considers it a new call to
	// UpdateFlow . The token is active for 8 hours.
	ClientToken *string

	// A description of the flow.
	Description *string

	// Specifies the configuration that Amazon AppFlow uses when it catalogs the data
	// that's transferred by the associated flow. When Amazon AppFlow catalogs the data
	// from a flow, it stores metadata in a data catalog.
	MetadataCatalogConfig *types.MetadataCatalogConfig

	noSmithyDocumentSerde
}

type UpdateFlowOutput struct {

	// Indicates the current status of the flow.
	FlowStatus types.FlowStatus

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateFlowMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateFlow{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateFlow{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "UpdateFlow"); err != nil {
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
	if err = addIdempotencyToken_opUpdateFlowMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpUpdateFlowValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateFlow(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpUpdateFlow struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpUpdateFlow) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpUpdateFlow) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*UpdateFlowInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *UpdateFlowInput ")
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
func addIdempotencyToken_opUpdateFlowMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpUpdateFlow{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opUpdateFlow(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "UpdateFlow",
	}
}
