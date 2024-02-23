// Code generated by smithy-go-codegen DO NOT EDIT.

package comprehendmedical

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/comprehendmedical/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// InferRxNorm detects medications as entities listed in a patient record and
// links to the normalized concept identifiers in the RxNorm database from the
// National Library of Medicine. Amazon Comprehend Medical only detects medical
// entities in English language texts.
func (c *Client) InferRxNorm(ctx context.Context, params *InferRxNormInput, optFns ...func(*Options)) (*InferRxNormOutput, error) {
	if params == nil {
		params = &InferRxNormInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "InferRxNorm", params, optFns, c.addOperationInferRxNormMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*InferRxNormOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type InferRxNormInput struct {

	// The input text used for analysis.
	//
	// This member is required.
	Text *string

	noSmithyDocumentSerde
}

type InferRxNormOutput struct {

	// The medication entities detected in the text linked to RxNorm concepts. If the
	// action is successful, the service sends back an HTTP 200 response, as well as
	// the entities detected.
	//
	// This member is required.
	Entities []types.RxNormEntity

	// The version of the model used to analyze the documents, in the format n.n.n You
	// can use this information to track the model used for a particular batch of
	// documents.
	ModelVersion *string

	// If the result of the previous request to InferRxNorm was truncated, include the
	// PaginationToken to fetch the next page of medication entities.
	PaginationToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationInferRxNormMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpInferRxNorm{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpInferRxNorm{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "InferRxNorm"); err != nil {
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
	if err = addOpInferRxNormValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opInferRxNorm(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opInferRxNorm(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "InferRxNorm",
	}
}
