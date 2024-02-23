// Code generated by smithy-go-codegen DO NOT EDIT.

package iotevents

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/iotevents/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves one or more analysis results of the detector model. After AWS IoT
// Events starts analyzing your detector model, you have up to 24 hours to retrieve
// the analysis results.
func (c *Client) GetDetectorModelAnalysisResults(ctx context.Context, params *GetDetectorModelAnalysisResultsInput, optFns ...func(*Options)) (*GetDetectorModelAnalysisResultsOutput, error) {
	if params == nil {
		params = &GetDetectorModelAnalysisResultsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetDetectorModelAnalysisResults", params, optFns, c.addOperationGetDetectorModelAnalysisResultsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetDetectorModelAnalysisResultsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetDetectorModelAnalysisResultsInput struct {

	// The ID of the analysis result that you want to retrieve.
	//
	// This member is required.
	AnalysisId *string

	// The maximum number of results to be returned per request.
	MaxResults *int32

	// The token that you can use to return the next set of results.
	NextToken *string

	noSmithyDocumentSerde
}

type GetDetectorModelAnalysisResultsOutput struct {

	// Contains information about one or more analysis results.
	AnalysisResults []types.AnalysisResult

	// The token that you can use to return the next set of results, or null if there
	// are no more results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetDetectorModelAnalysisResultsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetDetectorModelAnalysisResults{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetDetectorModelAnalysisResults{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetDetectorModelAnalysisResults"); err != nil {
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
	if err = addOpGetDetectorModelAnalysisResultsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetDetectorModelAnalysisResults(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetDetectorModelAnalysisResults(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetDetectorModelAnalysisResults",
	}
}
