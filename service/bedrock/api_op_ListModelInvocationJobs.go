// Code generated by smithy-go-codegen DO NOT EDIT.

package bedrock

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/bedrock/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Lists all batch inference jobs in the account. For more information, see [View details about a batch inference job].
//
// [View details about a batch inference job]: https://docs.aws.amazon.com/bedrock/latest/userguide/batch-inference-manage.html#batch-inference-view
func (c *Client) ListModelInvocationJobs(ctx context.Context, params *ListModelInvocationJobsInput, optFns ...func(*Options)) (*ListModelInvocationJobsOutput, error) {
	if params == nil {
		params = &ListModelInvocationJobsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListModelInvocationJobs", params, optFns, c.addOperationListModelInvocationJobsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListModelInvocationJobsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListModelInvocationJobsInput struct {

	// The maximum number of results to return. If there are more results than the
	// number that you specify, a nextToken value is returned. Use the nextToken in a
	// request to return the next batch of results.
	MaxResults *int32

	// Specify a string to filter for batch inference jobs whose names contain the
	// string.
	NameContains *string

	// If there were more results than the value you specified in the maxResults field
	// in a previous ListModelInvocationJobs request, the response would have returned
	// a nextToken value. To see the next batch of results, send the nextToken value
	// in another request.
	NextToken *string

	// An attribute by which to sort the results.
	SortBy types.SortJobsBy

	// Specifies whether to sort the results by ascending or descending order.
	SortOrder types.SortOrder

	// Specify a status to filter for batch inference jobs whose statuses match the
	// string you specify.
	StatusEquals types.ModelInvocationJobStatus

	// Specify a time to filter for batch inference jobs that were submitted after the
	// time you specify.
	SubmitTimeAfter *time.Time

	// Specify a time to filter for batch inference jobs that were submitted before
	// the time you specify.
	SubmitTimeBefore *time.Time

	noSmithyDocumentSerde
}

type ListModelInvocationJobsOutput struct {

	// A list of items, each of which contains a summary about a batch inference job.
	InvocationJobSummaries []types.ModelInvocationJobSummary

	// If there are more results than can fit in the response, a nextToken is
	// returned. Use the nextToken in a request to return the next batch of results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListModelInvocationJobsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListModelInvocationJobs{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListModelInvocationJobs{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListModelInvocationJobs"); err != nil {
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
	if err = addTimeOffsetBuild(stack, c); err != nil {
		return err
	}
	if err = addUserAgentRetryMode(stack, options); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListModelInvocationJobs(options.Region), middleware.Before); err != nil {
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

// ListModelInvocationJobsPaginatorOptions is the paginator options for
// ListModelInvocationJobs
type ListModelInvocationJobsPaginatorOptions struct {
	// The maximum number of results to return. If there are more results than the
	// number that you specify, a nextToken value is returned. Use the nextToken in a
	// request to return the next batch of results.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListModelInvocationJobsPaginator is a paginator for ListModelInvocationJobs
type ListModelInvocationJobsPaginator struct {
	options   ListModelInvocationJobsPaginatorOptions
	client    ListModelInvocationJobsAPIClient
	params    *ListModelInvocationJobsInput
	nextToken *string
	firstPage bool
}

// NewListModelInvocationJobsPaginator returns a new
// ListModelInvocationJobsPaginator
func NewListModelInvocationJobsPaginator(client ListModelInvocationJobsAPIClient, params *ListModelInvocationJobsInput, optFns ...func(*ListModelInvocationJobsPaginatorOptions)) *ListModelInvocationJobsPaginator {
	if params == nil {
		params = &ListModelInvocationJobsInput{}
	}

	options := ListModelInvocationJobsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListModelInvocationJobsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListModelInvocationJobsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListModelInvocationJobs page.
func (p *ListModelInvocationJobsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListModelInvocationJobsOutput, error) {
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

	optFns = append([]func(*Options){
		addIsPaginatorUserAgent,
	}, optFns...)
	result, err := p.client.ListModelInvocationJobs(ctx, &params, optFns...)
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

// ListModelInvocationJobsAPIClient is a client that implements the
// ListModelInvocationJobs operation.
type ListModelInvocationJobsAPIClient interface {
	ListModelInvocationJobs(context.Context, *ListModelInvocationJobsInput, ...func(*Options)) (*ListModelInvocationJobsOutput, error)
}

var _ ListModelInvocationJobsAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opListModelInvocationJobs(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListModelInvocationJobs",
	}
}
