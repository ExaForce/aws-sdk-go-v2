// Code generated by smithy-go-codegen DO NOT EDIT.

package billingconductor

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/billingconductor/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a custom line item that can be used to create a one-time fixed charge
// that can be applied to a single billing group for the current or previous
// billing period. The one-time fixed charge is either a fee or discount.
func (c *Client) CreateCustomLineItem(ctx context.Context, params *CreateCustomLineItemInput, optFns ...func(*Options)) (*CreateCustomLineItemOutput, error) {
	if params == nil {
		params = &CreateCustomLineItemInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateCustomLineItem", params, optFns, c.addOperationCreateCustomLineItemMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateCustomLineItemOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateCustomLineItemInput struct {

	// The Amazon Resource Name (ARN) that references the billing group where the
	// custom line item applies to.
	//
	// This member is required.
	BillingGroupArn *string

	// A CustomLineItemChargeDetails that describes the charge details for a custom
	// line item.
	//
	// This member is required.
	ChargeDetails *types.CustomLineItemChargeDetails

	// The description of the custom line item. This is shown on the Bills page in
	// association with the charge value.
	//
	// This member is required.
	Description *string

	// The name of the custom line item.
	//
	// This member is required.
	Name *string

	// The Amazon Web Services account in which this custom line item will be applied
	// to.
	AccountId *string

	// A time range for which the custom line item is effective.
	BillingPeriodRange *types.CustomLineItemBillingPeriodRange

	// The token that is needed to support idempotency. Idempotency isn't currently
	// supported, but will be implemented in a future update.
	ClientToken *string

	// A map that contains tag keys and tag values that are attached to a custom line
	// item.
	Tags map[string]string

	noSmithyDocumentSerde
}

type CreateCustomLineItemOutput struct {

	// The Amazon Resource Name (ARN) of the created custom line item.
	Arn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateCustomLineItemMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateCustomLineItem{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateCustomLineItem{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateCustomLineItem"); err != nil {
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
	if err = addIdempotencyToken_opCreateCustomLineItemMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpCreateCustomLineItemValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateCustomLineItem(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpCreateCustomLineItem struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreateCustomLineItem) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreateCustomLineItem) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CreateCustomLineItemInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CreateCustomLineItemInput ")
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
func addIdempotencyToken_opCreateCustomLineItemMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpCreateCustomLineItem{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCreateCustomLineItem(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateCustomLineItem",
	}
}
