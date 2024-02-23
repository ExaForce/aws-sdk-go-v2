// Code generated by smithy-go-codegen DO NOT EDIT.

package nimble

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/nimble/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Add/update users with given persona to studio membership.
func (c *Client) PutStudioMembers(ctx context.Context, params *PutStudioMembersInput, optFns ...func(*Options)) (*PutStudioMembersOutput, error) {
	if params == nil {
		params = &PutStudioMembersInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutStudioMembers", params, optFns, c.addOperationPutStudioMembersMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutStudioMembersOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutStudioMembersInput struct {

	// The ID of the identity store.
	//
	// This member is required.
	IdentityStoreId *string

	// A list of members.
	//
	// This member is required.
	Members []types.NewStudioMember

	// The studio ID.
	//
	// This member is required.
	StudioId *string

	// Unique, case-sensitive identifier that you provide to ensure the idempotency of
	// the request. If you don’t specify a client token, the Amazon Web Services SDK
	// automatically generates a client token and uses it for the request to ensure
	// idempotency.
	ClientToken *string

	noSmithyDocumentSerde
}

type PutStudioMembersOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationPutStudioMembersMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpPutStudioMembers{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpPutStudioMembers{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "PutStudioMembers"); err != nil {
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
	if err = addIdempotencyToken_opPutStudioMembersMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpPutStudioMembersValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opPutStudioMembers(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpPutStudioMembers struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpPutStudioMembers) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpPutStudioMembers) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*PutStudioMembersInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *PutStudioMembersInput ")
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
func addIdempotencyToken_opPutStudioMembersMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpPutStudioMembers{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opPutStudioMembers(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "PutStudioMembers",
	}
}
