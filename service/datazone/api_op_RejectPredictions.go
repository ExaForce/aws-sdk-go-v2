// Code generated by smithy-go-codegen DO NOT EDIT.

package datazone

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/datazone/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Rejects automatically generated business-friendly metadata for your Amazon
// DataZone assets.
func (c *Client) RejectPredictions(ctx context.Context, params *RejectPredictionsInput, optFns ...func(*Options)) (*RejectPredictionsOutput, error) {
	if params == nil {
		params = &RejectPredictionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "RejectPredictions", params, optFns, c.addOperationRejectPredictionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*RejectPredictionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type RejectPredictionsInput struct {

	// The identifier of the Amazon DataZone domain.
	//
	// This member is required.
	DomainIdentifier *string

	// The identifier of the prediction.
	//
	// This member is required.
	Identifier *string

	// A unique, case-sensitive identifier that is provided to ensure the idempotency
	// of the request.
	ClientToken *string

	// Specifies the prediction (aka, the automatically generated piece of metadata)
	// and the target (for example, a column name) that can be rejected.
	RejectChoices []types.RejectChoice

	// Specifies the rule (or the conditions) under which a prediction can be rejected.
	RejectRule *types.RejectRule

	// The revision that is to be made to the asset.
	Revision *string

	noSmithyDocumentSerde
}

type RejectPredictionsOutput struct {

	// The ID of the asset.
	//
	// This member is required.
	AssetId *string

	// The revision that is to be made to the asset.
	//
	// This member is required.
	AssetRevision *string

	// The ID of the Amazon DataZone domain.
	//
	// This member is required.
	DomainId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationRejectPredictionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpRejectPredictions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpRejectPredictions{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "RejectPredictions"); err != nil {
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
	if err = addIdempotencyToken_opRejectPredictionsMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpRejectPredictionsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opRejectPredictions(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpRejectPredictions struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpRejectPredictions) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpRejectPredictions) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*RejectPredictionsInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *RejectPredictionsInput ")
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
func addIdempotencyToken_opRejectPredictionsMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpRejectPredictions{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opRejectPredictions(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "RejectPredictions",
	}
}
