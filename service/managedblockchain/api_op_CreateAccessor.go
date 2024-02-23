// Code generated by smithy-go-codegen DO NOT EDIT.

package managedblockchain

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/managedblockchain/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a new accessor for use with Amazon Managed Blockchain service that
// supports token based access. The accessor contains information required for
// token based access.
func (c *Client) CreateAccessor(ctx context.Context, params *CreateAccessorInput, optFns ...func(*Options)) (*CreateAccessorOutput, error) {
	if params == nil {
		params = &CreateAccessorInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateAccessor", params, optFns, c.addOperationCreateAccessorMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateAccessorOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateAccessorInput struct {

	// The type of accessor. Currently, accessor type is restricted to BILLING_TOKEN .
	//
	// This member is required.
	AccessorType types.AccessorType

	// This is a unique, case-sensitive identifier that you provide to ensure the
	// idempotency of the operation. An idempotent operation completes no more than
	// once. This identifier is required only if you make a service request directly
	// using an HTTP client. It is generated automatically if you use an Amazon Web
	// Services SDK or the Amazon Web Services CLI.
	//
	// This member is required.
	ClientRequestToken *string

	// The blockchain network that the Accessor token is created for. We recommend
	// using the appropriate networkType value for the blockchain network that you are
	// creating the Accessor token for. You cannnot use the value
	// ETHEREUM_MAINNET_AND_GOERLI to specify a networkType for your Accessor token.
	// The default value of ETHEREUM_MAINNET_AND_GOERLI is only applied:
	//   - when the CreateAccessor action does not set a networkType .
	//   - to all existing Accessor tokens that were created before the networkType
	//   property was introduced.
	NetworkType types.AccessorNetworkType

	// Tags to assign to the Accessor. Each tag consists of a key and an optional
	// value. You can specify multiple key-value pairs in a single request with an
	// overall maximum of 50 tags allowed per resource. For more information about
	// tags, see Tagging Resources (https://docs.aws.amazon.com/managed-blockchain/latest/ethereum-dev/tagging-resources.html)
	// in the Amazon Managed Blockchain Ethereum Developer Guide, or Tagging Resources (https://docs.aws.amazon.com/managed-blockchain/latest/hyperledger-fabric-dev/tagging-resources.html)
	// in the Amazon Managed Blockchain Hyperledger Fabric Developer Guide.
	Tags map[string]string

	noSmithyDocumentSerde
}

type CreateAccessorOutput struct {

	// The unique identifier of the accessor.
	AccessorId *string

	// The billing token is a property of the Accessor. Use this token to when making
	// calls to the blockchain network. The billing token is used to track your
	// accessor token for billing requests.
	BillingToken *string

	// The blockchain network that the accessor token is created for.
	NetworkType types.AccessorNetworkType

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateAccessorMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateAccessor{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateAccessor{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateAccessor"); err != nil {
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
	if err = addIdempotencyToken_opCreateAccessorMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpCreateAccessorValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateAccessor(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpCreateAccessor struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreateAccessor) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreateAccessor) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CreateAccessorInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CreateAccessorInput ")
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
func addIdempotencyToken_opCreateAccessorMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpCreateAccessor{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCreateAccessor(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateAccessor",
	}
}
