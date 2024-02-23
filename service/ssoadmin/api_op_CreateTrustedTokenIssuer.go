// Code generated by smithy-go-codegen DO NOT EDIT.

package ssoadmin

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/ssoadmin/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a connection to a trusted token issuer in an instance of IAM Identity
// Center. A trusted token issuer enables trusted identity propagation to be used
// with applications that authenticate outside of Amazon Web Services. This trusted
// token issuer describes an external identity provider (IdP) that can generate
// claims or assertions in the form of access tokens for a user. Applications
// enabled for IAM Identity Center can use these tokens for authentication.
func (c *Client) CreateTrustedTokenIssuer(ctx context.Context, params *CreateTrustedTokenIssuerInput, optFns ...func(*Options)) (*CreateTrustedTokenIssuerOutput, error) {
	if params == nil {
		params = &CreateTrustedTokenIssuerInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateTrustedTokenIssuer", params, optFns, c.addOperationCreateTrustedTokenIssuerMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateTrustedTokenIssuerOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateTrustedTokenIssuerInput struct {

	// Specifies the ARN of the instance of IAM Identity Center to contain the new
	// trusted token issuer configuration.
	//
	// This member is required.
	InstanceArn *string

	// Specifies the name of the new trusted token issuer configuration.
	//
	// This member is required.
	Name *string

	// Specifies settings that apply to the new trusted token issuer configuration.
	// The settings that are available depend on what TrustedTokenIssuerType you
	// specify.
	//
	// This member is required.
	TrustedTokenIssuerConfiguration types.TrustedTokenIssuerConfiguration

	// Specifies the type of the new trusted token issuer.
	//
	// This member is required.
	TrustedTokenIssuerType types.TrustedTokenIssuerType

	// Specifies a unique, case-sensitive ID that you provide to ensure the
	// idempotency of the request. This lets you safely retry the request without
	// accidentally performing the same operation a second time. Passing the same value
	// to a later call to an operation requires that you also pass the same value for
	// all other parameters. We recommend that you use a UUID type of value. (https://wikipedia.org/wiki/Universally_unique_identifier)
	// . If you don't provide this value, then Amazon Web Services generates a random
	// one for you. If you retry the operation with the same ClientToken , but with
	// different parameters, the retry fails with an IdempotentParameterMismatch error.
	ClientToken *string

	// Specifies tags to be attached to the new trusted token issuer configuration.
	Tags []types.Tag

	noSmithyDocumentSerde
}

type CreateTrustedTokenIssuerOutput struct {

	// The ARN of the new trusted token issuer configuration.
	TrustedTokenIssuerArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateTrustedTokenIssuerMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateTrustedTokenIssuer{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateTrustedTokenIssuer{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateTrustedTokenIssuer"); err != nil {
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
	if err = addIdempotencyToken_opCreateTrustedTokenIssuerMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpCreateTrustedTokenIssuerValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateTrustedTokenIssuer(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpCreateTrustedTokenIssuer struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreateTrustedTokenIssuer) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreateTrustedTokenIssuer) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CreateTrustedTokenIssuerInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CreateTrustedTokenIssuerInput ")
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
func addIdempotencyToken_opCreateTrustedTokenIssuerMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpCreateTrustedTokenIssuer{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCreateTrustedTokenIssuer(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateTrustedTokenIssuer",
	}
}
