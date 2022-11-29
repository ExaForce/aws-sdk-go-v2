// Code generated by smithy-go-codegen DO NOT EDIT.

package opensearchserverless

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Deletes a security configuration for OpenSearch Serverless. For more
// information, see SAML authentication for Amazon OpenSearch Serverless
// (https://docs.aws.amazon.com/opensearch-service/latest/developerguide/serverless-saml.html).
func (c *Client) DeleteSecurityConfig(ctx context.Context, params *DeleteSecurityConfigInput, optFns ...func(*Options)) (*DeleteSecurityConfigOutput, error) {
	if params == nil {
		params = &DeleteSecurityConfigInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteSecurityConfig", params, optFns, c.addOperationDeleteSecurityConfigMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteSecurityConfigOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteSecurityConfigInput struct {

	// The security configuration identifier. For SAML the ID will be
	// saml/<accountId>/<idpProviderName>. For example, saml/123456789123/OKTADev.
	//
	// This member is required.
	Id *string

	// Unique, case-sensitive identifier to ensure idempotency of the request.
	ClientToken *string

	noSmithyDocumentSerde
}

type DeleteSecurityConfigOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeleteSecurityConfigMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpDeleteSecurityConfig{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpDeleteSecurityConfig{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addIdempotencyToken_opDeleteSecurityConfigMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpDeleteSecurityConfigValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteSecurityConfig(options.Region), middleware.Before); err != nil {
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
	return nil
}

type idempotencyToken_initializeOpDeleteSecurityConfig struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpDeleteSecurityConfig) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpDeleteSecurityConfig) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*DeleteSecurityConfigInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *DeleteSecurityConfigInput ")
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
func addIdempotencyToken_opDeleteSecurityConfigMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpDeleteSecurityConfig{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opDeleteSecurityConfig(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "aoss",
		OperationName: "DeleteSecurityConfig",
	}
}
