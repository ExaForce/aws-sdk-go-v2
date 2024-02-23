// Code generated by smithy-go-codegen DO NOT EDIT.

package fis

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/fis/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a target account configuration for the experiment template. A target
// account configuration is required when accountTargeting of experimentOptions is
// set to multi-account . For more information, see experiment options (https://docs.aws.amazon.com/fis/latest/userguide/experiment-options.html)
// in the Fault Injection Simulator User Guide.
func (c *Client) CreateTargetAccountConfiguration(ctx context.Context, params *CreateTargetAccountConfigurationInput, optFns ...func(*Options)) (*CreateTargetAccountConfigurationOutput, error) {
	if params == nil {
		params = &CreateTargetAccountConfigurationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateTargetAccountConfiguration", params, optFns, c.addOperationCreateTargetAccountConfigurationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateTargetAccountConfigurationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateTargetAccountConfigurationInput struct {

	// The AWS account ID of the target account.
	//
	// This member is required.
	AccountId *string

	// The experiment template ID.
	//
	// This member is required.
	ExperimentTemplateId *string

	// The Amazon Resource Name (ARN) of an IAM role for the target account.
	//
	// This member is required.
	RoleArn *string

	// Unique, case-sensitive identifier that you provide to ensure the idempotency of
	// the request.
	ClientToken *string

	// The description of the target account.
	Description *string

	noSmithyDocumentSerde
}

type CreateTargetAccountConfigurationOutput struct {

	// Information about the target account configuration.
	TargetAccountConfiguration *types.TargetAccountConfiguration

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateTargetAccountConfigurationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateTargetAccountConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateTargetAccountConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateTargetAccountConfiguration"); err != nil {
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
	if err = addIdempotencyToken_opCreateTargetAccountConfigurationMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpCreateTargetAccountConfigurationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateTargetAccountConfiguration(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpCreateTargetAccountConfiguration struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreateTargetAccountConfiguration) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreateTargetAccountConfiguration) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CreateTargetAccountConfigurationInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CreateTargetAccountConfigurationInput ")
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
func addIdempotencyToken_opCreateTargetAccountConfigurationMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpCreateTargetAccountConfiguration{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCreateTargetAccountConfiguration(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateTargetAccountConfiguration",
	}
}
