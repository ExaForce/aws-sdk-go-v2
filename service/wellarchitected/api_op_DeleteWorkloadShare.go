// Code generated by smithy-go-codegen DO NOT EDIT.

package wellarchitected

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Delete a workload share.
func (c *Client) DeleteWorkloadShare(ctx context.Context, params *DeleteWorkloadShareInput, optFns ...func(*Options)) (*DeleteWorkloadShareOutput, error) {
	if params == nil {
		params = &DeleteWorkloadShareInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteWorkloadShare", params, optFns, c.addOperationDeleteWorkloadShareMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteWorkloadShareOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Input for Delete Workload Share
type DeleteWorkloadShareInput struct {

	// A unique case-sensitive string used to ensure that this request is idempotent
	// (executes only once). You should not reuse the same token for other requests. If
	// you retry a request with the same client request token and the same parameters
	// after the original request has completed successfully, the result of the
	// original request is returned. This token is listed as required, however, if you
	// do not specify it, the Amazon Web Services SDKs automatically generate one for
	// you. If you are not using the Amazon Web Services SDK or the CLI, you must
	// provide this token or the request will fail.
	//
	// This member is required.
	ClientRequestToken *string

	// The ID associated with the share.
	//
	// This member is required.
	ShareId *string

	// The ID assigned to the workload. This ID is unique within an Amazon Web
	// Services Region.
	//
	// This member is required.
	WorkloadId *string

	noSmithyDocumentSerde
}

type DeleteWorkloadShareOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeleteWorkloadShareMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDeleteWorkloadShare{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDeleteWorkloadShare{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DeleteWorkloadShare"); err != nil {
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
	if err = addIdempotencyToken_opDeleteWorkloadShareMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpDeleteWorkloadShareValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteWorkloadShare(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpDeleteWorkloadShare struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpDeleteWorkloadShare) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpDeleteWorkloadShare) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*DeleteWorkloadShareInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *DeleteWorkloadShareInput ")
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
func addIdempotencyToken_opDeleteWorkloadShareMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpDeleteWorkloadShare{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opDeleteWorkloadShare(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DeleteWorkloadShare",
	}
}
