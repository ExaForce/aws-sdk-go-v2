// Code generated by smithy-go-codegen DO NOT EDIT.

package sagemaker

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/sagemaker/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Update an Amazon SageMaker Model Card. You cannot update both model card
// content and model card status in a single call.
func (c *Client) UpdateModelCard(ctx context.Context, params *UpdateModelCardInput, optFns ...func(*Options)) (*UpdateModelCardOutput, error) {
	if params == nil {
		params = &UpdateModelCardInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateModelCard", params, optFns, c.addOperationUpdateModelCardMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateModelCardOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateModelCardInput struct {

	// The name or Amazon Resource Name (ARN) of the model card to update.
	//
	// This member is required.
	ModelCardName *string

	// The updated model card content. Content must be in model card JSON schema (https://docs.aws.amazon.com/sagemaker/latest/dg/model-cards.html#model-cards-json-schema)
	// and provided as a string. When updating model card content, be sure to include
	// the full content and not just updated content.
	Content *string

	// The approval status of the model card within your organization. Different
	// organizations might have different criteria for model card review and approval.
	//   - Draft : The model card is a work in progress.
	//   - PendingReview : The model card is pending review.
	//   - Approved : The model card is approved.
	//   - Archived : The model card is archived. No more updates should be made to the
	//   model card, but it can still be exported.
	ModelCardStatus types.ModelCardStatus

	noSmithyDocumentSerde
}

type UpdateModelCardOutput struct {

	// The Amazon Resource Name (ARN) of the updated model card.
	//
	// This member is required.
	ModelCardArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateModelCardMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateModelCard{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateModelCard{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "UpdateModelCard"); err != nil {
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
	if err = addOpUpdateModelCardValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateModelCard(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateModelCard(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "UpdateModelCard",
	}
}
