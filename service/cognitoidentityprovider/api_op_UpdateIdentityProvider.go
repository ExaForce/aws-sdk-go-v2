// Code generated by smithy-go-codegen DO NOT EDIT.

package cognitoidentityprovider

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates IdP information for a user pool. Amazon Cognito evaluates Identity and
// Access Management (IAM) policies in requests for this API operation. For this
// operation, you must use IAM credentials to authorize requests, and you must
// grant yourself the corresponding IAM permission in a policy. Learn more
//   - Signing Amazon Web Services API Requests (https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_aws-signing.html)
//   - Using the Amazon Cognito user pools API and user pool endpoints (https://docs.aws.amazon.com/cognito/latest/developerguide/user-pools-API-operations.html)
func (c *Client) UpdateIdentityProvider(ctx context.Context, params *UpdateIdentityProviderInput, optFns ...func(*Options)) (*UpdateIdentityProviderOutput, error) {
	if params == nil {
		params = &UpdateIdentityProviderInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateIdentityProvider", params, optFns, c.addOperationUpdateIdentityProviderMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateIdentityProviderOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateIdentityProviderInput struct {

	// The IdP name.
	//
	// This member is required.
	ProviderName *string

	// The user pool ID.
	//
	// This member is required.
	UserPoolId *string

	// The IdP attribute mapping to be changed.
	AttributeMapping map[string]string

	// A list of IdP identifiers.
	IdpIdentifiers []string

	// The scopes, URLs, and identifiers for your external identity provider. The
	// following examples describe the provider detail keys for each IdP type. These
	// values and their schema are subject to change. Social IdP authorize_scopes
	// values must match the values listed here. OpenID Connect (OIDC) Amazon Cognito
	// accepts the following elements when it can't discover endpoint URLs from
	// oidc_issuer : attributes_url , authorize_url , jwks_uri , token_url . Create or
	// update request: "ProviderDetails": { "attributes_request_method": "GET",
	// "attributes_url": "https://auth.example.com/userInfo", "authorize_scopes":
	// "openid profile email", "authorize_url": "https://auth.example.com/authorize",
	// "client_id": "1example23456789", "client_secret": "provider-app-client-secret",
	// "jwks_uri": "https://auth.example.com/.well-known/jwks.json", "oidc_issuer":
	// "https://auth.example.com", "token_url": "https://example.com/token" } Describe
	// response: "ProviderDetails": { "attributes_request_method": "GET",
	// "attributes_url": "https://auth.example.com/userInfo",
	// "attributes_url_add_attributes": "false", "authorize_scopes": "openid profile
	// email", "authorize_url": "https://auth.example.com/authorize", "client_id":
	// "1example23456789", "client_secret": "provider-app-client-secret", "jwks_uri":
	// "https://auth.example.com/.well-known/jwks.json", "oidc_issuer":
	// "https://auth.example.com", "token_url": "https://example.com/token" } SAML
	// Create or update request with Metadata URL: "ProviderDetails": { "IDPInit":
	// "true", "IDPSignout": "true", "EncryptedResponses" : "true", "MetadataURL":
	// "https://auth.example.com/sso/saml/metadata", "RequestSigningAlgorithm":
	// "rsa-sha256" } Create or update request with Metadata file: "ProviderDetails":
	// { "IDPInit": "true", "IDPSignout": "true", "EncryptedResponses" : "true",
	// "MetadataFile": "[metadata XML]", "RequestSigningAlgorithm": "rsa-sha256" } The
	// value of MetadataFile must be the plaintext metadata document with all quote
	// (") characters escaped by backslashes. Describe response: "ProviderDetails": {
	// "IDPInit": "true", "IDPSignout": "true", "EncryptedResponses" : "true",
	// "ActiveEncryptionCertificate": "[certificate]", "MetadataURL":
	// "https://auth.example.com/sso/saml/metadata", "RequestSigningAlgorithm":
	// "rsa-sha256", "SLORedirectBindingURI": "https://auth.example.com/slo/saml",
	// "SSORedirectBindingURI": "https://auth.example.com/sso/saml" } LoginWithAmazon
	// Create or update request: "ProviderDetails": { "authorize_scopes": "profile
	// postal_code", "client_id": "amzn1.application-oa2-client.1example23456789",
	// "client_secret": "provider-app-client-secret" Describe response:
	// "ProviderDetails": { "attributes_url": "https://api.amazon.com/user/profile",
	// "attributes_url_add_attributes": "false", "authorize_scopes": "profile
	// postal_code", "authorize_url": "https://www.amazon.com/ap/oa", "client_id":
	// "amzn1.application-oa2-client.1example23456789", "client_secret":
	// "provider-app-client-secret", "token_request_method": "POST", "token_url":
	// "https://api.amazon.com/auth/o2/token" } Google Create or update request:
	// "ProviderDetails": { "authorize_scopes": "email profile openid", "client_id":
	// "1example23456789.apps.googleusercontent.com", "client_secret":
	// "provider-app-client-secret" } Describe response: "ProviderDetails": {
	// "attributes_url": "https://people.googleapis.com/v1/people/me?personFields=",
	// "attributes_url_add_attributes": "true", "authorize_scopes": "email profile
	// openid", "authorize_url": "https://accounts.google.com/o/oauth2/v2/auth",
	// "client_id": "1example23456789.apps.googleusercontent.com", "client_secret":
	// "provider-app-client-secret", "oidc_issuer": "https://accounts.google.com",
	// "token_request_method": "POST", "token_url":
	// "https://www.googleapis.com/oauth2/v4/token" } SignInWithApple Create or update
	// request: "ProviderDetails": { "authorize_scopes": "email name", "client_id":
	// "com.example.cognito", "private_key": "1EXAMPLE", "key_id": "2EXAMPLE",
	// "team_id": "3EXAMPLE" } Describe response: "ProviderDetails": {
	// "attributes_url_add_attributes": "false", "authorize_scopes": "email name",
	// "authorize_url": "https://appleid.apple.com/auth/authorize", "client_id":
	// "com.example.cognito", "key_id": "1EXAMPLE", "oidc_issuer":
	// "https://appleid.apple.com", "team_id": "2EXAMPLE", "token_request_method":
	// "POST", "token_url": "https://appleid.apple.com/auth/token" } Facebook Create or
	// update request: "ProviderDetails": { "api_version": "v17.0",
	// "authorize_scopes": "public_profile, email", "client_id": "1example23456789",
	// "client_secret": "provider-app-client-secret" } Describe response:
	// "ProviderDetails": { "api_version": "v17.0", "attributes_url":
	// "https://graph.facebook.com/v17.0/me?fields=", "attributes_url_add_attributes":
	// "true", "authorize_scopes": "public_profile, email", "authorize_url":
	// "https://www.facebook.com/v17.0/dialog/oauth", "client_id": "1example23456789",
	// "client_secret": "provider-app-client-secret", "token_request_method": "GET",
	// "token_url": "https://graph.facebook.com/v17.0/oauth/access_token" }
	ProviderDetails map[string]string

	noSmithyDocumentSerde
}

type UpdateIdentityProviderOutput struct {

	// The identity provider details.
	//
	// This member is required.
	IdentityProvider *types.IdentityProviderType

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateIdentityProviderMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateIdentityProvider{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateIdentityProvider{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "UpdateIdentityProvider"); err != nil {
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
	if err = addOpUpdateIdentityProviderValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateIdentityProvider(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateIdentityProvider(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "UpdateIdentityProvider",
	}
}
