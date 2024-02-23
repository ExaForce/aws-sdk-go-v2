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

// Gets a list of PipeLineExecutionStep objects.
func (c *Client) ListPipelineExecutionSteps(ctx context.Context, params *ListPipelineExecutionStepsInput, optFns ...func(*Options)) (*ListPipelineExecutionStepsOutput, error) {
	if params == nil {
		params = &ListPipelineExecutionStepsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListPipelineExecutionSteps", params, optFns, c.addOperationListPipelineExecutionStepsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListPipelineExecutionStepsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListPipelineExecutionStepsInput struct {

	// The maximum number of pipeline execution steps to return in the response.
	MaxResults *int32

	// If the result of the previous ListPipelineExecutionSteps request was truncated,
	// the response includes a NextToken . To retrieve the next set of pipeline
	// execution steps, use the token in the next request.
	NextToken *string

	// The Amazon Resource Name (ARN) of the pipeline execution.
	PipelineExecutionArn *string

	// The field by which to sort results. The default is CreatedTime .
	SortOrder types.SortOrder

	noSmithyDocumentSerde
}

type ListPipelineExecutionStepsOutput struct {

	// If the result of the previous ListPipelineExecutionSteps request was truncated,
	// the response includes a NextToken . To retrieve the next set of pipeline
	// execution steps, use the token in the next request.
	NextToken *string

	// A list of PipeLineExecutionStep objects. Each PipeLineExecutionStep consists of
	// StepName, StartTime, EndTime, StepStatus, and Metadata. Metadata is an object
	// with properties for each job that contains relevant information about the job
	// created by the step.
	PipelineExecutionSteps []types.PipelineExecutionStep

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListPipelineExecutionStepsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListPipelineExecutionSteps{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListPipelineExecutionSteps{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListPipelineExecutionSteps"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListPipelineExecutionSteps(options.Region), middleware.Before); err != nil {
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

// ListPipelineExecutionStepsAPIClient is a client that implements the
// ListPipelineExecutionSteps operation.
type ListPipelineExecutionStepsAPIClient interface {
	ListPipelineExecutionSteps(context.Context, *ListPipelineExecutionStepsInput, ...func(*Options)) (*ListPipelineExecutionStepsOutput, error)
}

var _ ListPipelineExecutionStepsAPIClient = (*Client)(nil)

// ListPipelineExecutionStepsPaginatorOptions is the paginator options for
// ListPipelineExecutionSteps
type ListPipelineExecutionStepsPaginatorOptions struct {
	// The maximum number of pipeline execution steps to return in the response.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListPipelineExecutionStepsPaginator is a paginator for
// ListPipelineExecutionSteps
type ListPipelineExecutionStepsPaginator struct {
	options   ListPipelineExecutionStepsPaginatorOptions
	client    ListPipelineExecutionStepsAPIClient
	params    *ListPipelineExecutionStepsInput
	nextToken *string
	firstPage bool
}

// NewListPipelineExecutionStepsPaginator returns a new
// ListPipelineExecutionStepsPaginator
func NewListPipelineExecutionStepsPaginator(client ListPipelineExecutionStepsAPIClient, params *ListPipelineExecutionStepsInput, optFns ...func(*ListPipelineExecutionStepsPaginatorOptions)) *ListPipelineExecutionStepsPaginator {
	if params == nil {
		params = &ListPipelineExecutionStepsInput{}
	}

	options := ListPipelineExecutionStepsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListPipelineExecutionStepsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListPipelineExecutionStepsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListPipelineExecutionSteps page.
func (p *ListPipelineExecutionStepsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListPipelineExecutionStepsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.MaxResults = limit

	result, err := p.client.ListPipelineExecutionSteps(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.NextToken

	if p.options.StopOnDuplicateToken &&
		prevToken != nil &&
		p.nextToken != nil &&
		*prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

func newServiceMetadataMiddleware_opListPipelineExecutionSteps(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListPipelineExecutionSteps",
	}
}
