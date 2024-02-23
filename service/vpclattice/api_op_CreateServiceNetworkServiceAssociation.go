// Code generated by smithy-go-codegen DO NOT EDIT.

package vpclattice

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/vpclattice/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Associates a service with a service network. You can't use this operation if
// the service and service network are already associated or if there is a
// disassociation or deletion in progress. If the association fails, you can retry
// the operation by deleting the association and recreating it. You cannot
// associate a service and service network that are shared with a caller. The
// caller must own either the service or the service network. As a result of this
// operation, the association is created in the service network account and the
// association owner account.
func (c *Client) CreateServiceNetworkServiceAssociation(ctx context.Context, params *CreateServiceNetworkServiceAssociationInput, optFns ...func(*Options)) (*CreateServiceNetworkServiceAssociationOutput, error) {
	if params == nil {
		params = &CreateServiceNetworkServiceAssociationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateServiceNetworkServiceAssociation", params, optFns, c.addOperationCreateServiceNetworkServiceAssociationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateServiceNetworkServiceAssociationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateServiceNetworkServiceAssociationInput struct {

	// The ID or Amazon Resource Name (ARN) of the service.
	//
	// This member is required.
	ServiceIdentifier *string

	// The ID or Amazon Resource Name (ARN) of the service network. You must use the
	// ARN if the resources specified in the operation are in different accounts.
	//
	// This member is required.
	ServiceNetworkIdentifier *string

	// A unique, case-sensitive identifier that you provide to ensure the idempotency
	// of the request. If you retry a request that completed successfully using the
	// same client token and parameters, the retry succeeds without performing any
	// actions. If the parameters aren't identical, the retry fails.
	ClientToken *string

	// The tags for the association.
	Tags map[string]string

	noSmithyDocumentSerde
}

type CreateServiceNetworkServiceAssociationOutput struct {

	// The Amazon Resource Name (ARN) of the association.
	Arn *string

	// The account that created the association.
	CreatedBy *string

	// The custom domain name of the service.
	CustomDomainName *string

	// The DNS name of the service.
	DnsEntry *types.DnsEntry

	// The ID of the association.
	Id *string

	// The operation's status.
	Status types.ServiceNetworkServiceAssociationStatus

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateServiceNetworkServiceAssociationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateServiceNetworkServiceAssociation{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateServiceNetworkServiceAssociation{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateServiceNetworkServiceAssociation"); err != nil {
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
	if err = addIdempotencyToken_opCreateServiceNetworkServiceAssociationMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpCreateServiceNetworkServiceAssociationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateServiceNetworkServiceAssociation(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpCreateServiceNetworkServiceAssociation struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreateServiceNetworkServiceAssociation) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreateServiceNetworkServiceAssociation) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CreateServiceNetworkServiceAssociationInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CreateServiceNetworkServiceAssociationInput ")
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
func addIdempotencyToken_opCreateServiceNetworkServiceAssociationMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpCreateServiceNetworkServiceAssociation{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCreateServiceNetworkServiceAssociation(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateServiceNetworkServiceAssociation",
	}
}
