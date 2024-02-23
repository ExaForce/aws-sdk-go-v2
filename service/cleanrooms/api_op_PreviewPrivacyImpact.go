// Code generated by smithy-go-codegen DO NOT EDIT.

package cleanrooms

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/cleanrooms/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// An estimate of the number of aggregation functions that the member who can
// query can run given epsilon and noise parameters.
func (c *Client) PreviewPrivacyImpact(ctx context.Context, params *PreviewPrivacyImpactInput, optFns ...func(*Options)) (*PreviewPrivacyImpactOutput, error) {
	if params == nil {
		params = &PreviewPrivacyImpactInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PreviewPrivacyImpact", params, optFns, c.addOperationPreviewPrivacyImpactMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PreviewPrivacyImpactOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PreviewPrivacyImpactInput struct {

	// A unique identifier for one of your memberships for a collaboration. Accepts a
	// membership ID.
	//
	// This member is required.
	MembershipIdentifier *string

	// Specifies the desired epsilon and noise parameters to preview.
	//
	// This member is required.
	Parameters types.PreviewPrivacyImpactParametersInput

	noSmithyDocumentSerde
}

type PreviewPrivacyImpactOutput struct {

	// An estimate of the number of aggregation functions that the member who can
	// query can run given the epsilon and noise parameters. This does not change the
	// privacy budget.
	//
	// This member is required.
	PrivacyImpact types.PrivacyImpact

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationPreviewPrivacyImpactMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpPreviewPrivacyImpact{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpPreviewPrivacyImpact{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "PreviewPrivacyImpact"); err != nil {
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
	if err = addOpPreviewPrivacyImpactValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opPreviewPrivacyImpact(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opPreviewPrivacyImpact(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "PreviewPrivacyImpact",
	}
}
