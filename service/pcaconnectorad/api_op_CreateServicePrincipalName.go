// Code generated by smithy-go-codegen DO NOT EDIT.

package pcaconnectorad

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a service principal name (SPN) for the service account in Active
// Directory. Kerberos authentication uses SPNs to associate a service instance
// with a service sign-in account.
func (c *Client) CreateServicePrincipalName(ctx context.Context, params *CreateServicePrincipalNameInput, optFns ...func(*Options)) (*CreateServicePrincipalNameOutput, error) {
	if params == nil {
		params = &CreateServicePrincipalNameInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateServicePrincipalName", params, optFns, c.addOperationCreateServicePrincipalNameMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateServicePrincipalNameOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateServicePrincipalNameInput struct {

	// The Amazon Resource Name (ARN) that was returned when you called CreateConnector (https://docs.aws.amazon.com/pca-connector-ad/latest/APIReference/API_CreateConnector.html)
	// .
	//
	// This member is required.
	ConnectorArn *string

	// The Amazon Resource Name (ARN) that was returned when you called
	// CreateDirectoryRegistration (https://docs.aws.amazon.com/pca-connector-ad/latest/APIReference/API_CreateDirectoryRegistration.html)
	// .
	//
	// This member is required.
	DirectoryRegistrationArn *string

	// Idempotency token.
	ClientToken *string

	noSmithyDocumentSerde
}

type CreateServicePrincipalNameOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateServicePrincipalNameMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateServicePrincipalName{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateServicePrincipalName{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateServicePrincipalName"); err != nil {
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
	if err = addIdempotencyToken_opCreateServicePrincipalNameMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpCreateServicePrincipalNameValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateServicePrincipalName(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpCreateServicePrincipalName struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreateServicePrincipalName) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreateServicePrincipalName) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CreateServicePrincipalNameInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CreateServicePrincipalNameInput ")
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
func addIdempotencyToken_opCreateServicePrincipalNameMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpCreateServicePrincipalName{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCreateServicePrincipalName(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateServicePrincipalName",
	}
}
