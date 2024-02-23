// Code generated by smithy-go-codegen DO NOT EDIT.

package datazone

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/datazone/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Creates a subscription target in Amazon DataZone.
func (c *Client) CreateSubscriptionTarget(ctx context.Context, params *CreateSubscriptionTargetInput, optFns ...func(*Options)) (*CreateSubscriptionTargetOutput, error) {
	if params == nil {
		params = &CreateSubscriptionTargetInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateSubscriptionTarget", params, optFns, c.addOperationCreateSubscriptionTargetMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateSubscriptionTargetOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateSubscriptionTargetInput struct {

	// The asset types that can be included in the subscription target.
	//
	// This member is required.
	ApplicableAssetTypes []string

	// The authorized principals of the subscription target.
	//
	// This member is required.
	AuthorizedPrincipals []string

	// The ID of the Amazon DataZone domain in which subscription target is created.
	//
	// This member is required.
	DomainIdentifier *string

	// The ID of the environment in which subscription target is created.
	//
	// This member is required.
	EnvironmentIdentifier *string

	// The manage access role that is used to create the subscription target.
	//
	// This member is required.
	ManageAccessRole *string

	// The name of the subscription target.
	//
	// This member is required.
	Name *string

	// The configuration of the subscription target.
	//
	// This member is required.
	SubscriptionTargetConfig []types.SubscriptionTargetForm

	// The type of the subscription target.
	//
	// This member is required.
	Type *string

	// A unique, case-sensitive identifier that is provided to ensure the idempotency
	// of the request.
	ClientToken *string

	// The provider of the subscription target.
	Provider *string

	noSmithyDocumentSerde
}

type CreateSubscriptionTargetOutput struct {

	// The asset types that can be included in the subscription target.
	//
	// This member is required.
	ApplicableAssetTypes []string

	// The authorised principals of the subscription target.
	//
	// This member is required.
	AuthorizedPrincipals []string

	// The timestamp of when the subscription target was created.
	//
	// This member is required.
	CreatedAt *time.Time

	// The Amazon DataZone user who created the subscription target.
	//
	// This member is required.
	CreatedBy *string

	// The ID of the Amazon DataZone domain in which the subscription target was
	// created.
	//
	// This member is required.
	DomainId *string

	// The ID of the environment in which the subscription target was created.
	//
	// This member is required.
	EnvironmentId *string

	// The ID of the subscription target.
	//
	// This member is required.
	Id *string

	// The manage access role with which the subscription target was created.
	//
	// This member is required.
	ManageAccessRole *string

	// The name of the subscription target.
	//
	// This member is required.
	Name *string

	// ???
	//
	// This member is required.
	ProjectId *string

	// The provider of the subscription target.
	//
	// This member is required.
	Provider *string

	// The configuration of the subscription target.
	//
	// This member is required.
	SubscriptionTargetConfig []types.SubscriptionTargetForm

	// The type of the subscription target.
	//
	// This member is required.
	Type *string

	// The timestamp of when the subscription target was updated.
	UpdatedAt *time.Time

	// The Amazon DataZone user who updated the subscription target.
	UpdatedBy *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateSubscriptionTargetMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateSubscriptionTarget{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateSubscriptionTarget{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateSubscriptionTarget"); err != nil {
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
	if err = addIdempotencyToken_opCreateSubscriptionTargetMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpCreateSubscriptionTargetValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateSubscriptionTarget(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpCreateSubscriptionTarget struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreateSubscriptionTarget) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreateSubscriptionTarget) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CreateSubscriptionTargetInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CreateSubscriptionTargetInput ")
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
func addIdempotencyToken_opCreateSubscriptionTargetMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpCreateSubscriptionTarget{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCreateSubscriptionTarget(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateSubscriptionTarget",
	}
}
