// Code generated by smithy-go-codegen DO NOT EDIT.

package connectcampaigns

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/connectcampaigns/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates dials requests for the specified campaign Amazon Connect account. This
// API is idempotent.
func (c *Client) PutDialRequestBatch(ctx context.Context, params *PutDialRequestBatchInput, optFns ...func(*Options)) (*PutDialRequestBatchOutput, error) {
	if params == nil {
		params = &PutDialRequestBatchInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutDialRequestBatch", params, optFns, c.addOperationPutDialRequestBatchMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutDialRequestBatchOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// PutDialRequestBatchRequest
type PutDialRequestBatchInput struct {

	// A list of dial requests.
	//
	// This member is required.
	DialRequests []types.DialRequest

	// Identifier representing a Campaign
	//
	// This member is required.
	Id *string

	noSmithyDocumentSerde
}

// PutDialRequestBatchResponse
type PutDialRequestBatchOutput struct {

	// A list of failed requests.
	FailedRequests []types.FailedRequest

	// A list of successful requests identified by the unique client token.
	SuccessfulRequests []types.SuccessfulRequest

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationPutDialRequestBatchMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpPutDialRequestBatch{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpPutDialRequestBatch{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "PutDialRequestBatch"); err != nil {
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
	if err = addOpPutDialRequestBatchValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opPutDialRequestBatch(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opPutDialRequestBatch(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "PutDialRequestBatch",
	}
}
