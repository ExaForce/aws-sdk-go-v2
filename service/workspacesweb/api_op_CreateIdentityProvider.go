// Code generated by smithy-go-codegen DO NOT EDIT.

package workspacesweb

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/workspacesweb/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates an identity provider resource that is then associated with a web portal.
func (c *Client) CreateIdentityProvider(ctx context.Context, params *CreateIdentityProviderInput, optFns ...func(*Options)) (*CreateIdentityProviderOutput, error) {
	if params == nil {
		params = &CreateIdentityProviderInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateIdentityProvider", params, optFns, c.addOperationCreateIdentityProviderMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateIdentityProviderOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateIdentityProviderInput struct {

	// The identity provider details. The following list describes the provider detail
	// keys for each identity provider type.
	//   - For Google and Login with Amazon:
	//   - client_id
	//   - client_secret
	//   - authorize_scopes
	//   - For Facebook:
	//   - client_id
	//   - client_secret
	//   - authorize_scopes
	//   - api_version
	//   - For Sign in with Apple:
	//   - client_id
	//   - team_id
	//   - key_id
	//   - private_key
	//   - authorize_scopes
	//   - For OIDC providers:
	//   - client_id
	//   - client_secret
	//   - attributes_request_method
	//   - oidc_issuer
	//   - authorize_scopes
	//   - authorize_url if not available from discovery URL specified by oidc_issuer
	//   key
	//   - token_url if not available from discovery URL specified by oidc_issuer key
	//   - attributes_url if not available from discovery URL specified by oidc_issuer
	//   key
	//   - jwks_uri if not available from discovery URL specified by oidc_issuer key
	//   - For SAML providers:
	//   - MetadataFile OR MetadataURL
	//   - IDPSignout (boolean) optional
	//
	// This member is required.
	IdentityProviderDetails map[string]string

	// The identity provider name.
	//
	// This member is required.
	IdentityProviderName *string

	// The identity provider type.
	//
	// This member is required.
	IdentityProviderType types.IdentityProviderType

	// The ARN of the web portal.
	//
	// This member is required.
	PortalArn *string

	// A unique, case-sensitive identifier that you provide to ensure the idempotency
	// of the request. Idempotency ensures that an API request completes only once.
	// With an idempotent request, if the original request completes successfully,
	// subsequent retries with the same client token returns the result from the
	// original successful request. If you do not specify a client token, one is
	// automatically generated by the AWS SDK.
	ClientToken *string

	noSmithyDocumentSerde
}

type CreateIdentityProviderOutput struct {

	// The ARN of the identity provider.
	//
	// This member is required.
	IdentityProviderArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateIdentityProviderMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateIdentityProvider{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateIdentityProvider{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateIdentityProvider"); err != nil {
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
	if err = addIdempotencyToken_opCreateIdentityProviderMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpCreateIdentityProviderValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateIdentityProvider(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpCreateIdentityProvider struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreateIdentityProvider) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreateIdentityProvider) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CreateIdentityProviderInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CreateIdentityProviderInput ")
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
func addIdempotencyToken_opCreateIdentityProviderMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpCreateIdentityProvider{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCreateIdentityProvider(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateIdentityProvider",
	}
}
