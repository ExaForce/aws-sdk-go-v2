// Code generated by smithy-go-codegen DO NOT EDIT.

package bedrockagent

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/bedrockagent/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// List ingestion jobs
func (c *Client) ListIngestionJobs(ctx context.Context, params *ListIngestionJobsInput, optFns ...func(*Options)) (*ListIngestionJobsOutput, error) {
	if params == nil {
		params = &ListIngestionJobsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListIngestionJobs", params, optFns, c.addOperationListIngestionJobsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListIngestionJobsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListIngestionJobsInput struct {

	// Identifier for a resource.
	//
	// This member is required.
	DataSourceId *string

	// Identifier for a resource.
	//
	// This member is required.
	KnowledgeBaseId *string

	// List of IngestionJobFilters
	Filters []types.IngestionJobFilter

	// Max Results.
	MaxResults *int32

	// Opaque continuation token of previous paginated response.
	NextToken *string

	// Sorts the response returned by ListIngestionJobs operation.
	SortBy *types.IngestionJobSortBy

	noSmithyDocumentSerde
}

type ListIngestionJobsOutput struct {

	// List of IngestionJobSummaries
	//
	// This member is required.
	IngestionJobSummaries []types.IngestionJobSummary

	// Opaque continuation token of previous paginated response.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListIngestionJobsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListIngestionJobs{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListIngestionJobs{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListIngestionJobs"); err != nil {
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
	if err = addOpListIngestionJobsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListIngestionJobs(options.Region), middleware.Before); err != nil {
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

// ListIngestionJobsAPIClient is a client that implements the ListIngestionJobs
// operation.
type ListIngestionJobsAPIClient interface {
	ListIngestionJobs(context.Context, *ListIngestionJobsInput, ...func(*Options)) (*ListIngestionJobsOutput, error)
}

var _ ListIngestionJobsAPIClient = (*Client)(nil)

// ListIngestionJobsPaginatorOptions is the paginator options for ListIngestionJobs
type ListIngestionJobsPaginatorOptions struct {
	// Max Results.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListIngestionJobsPaginator is a paginator for ListIngestionJobs
type ListIngestionJobsPaginator struct {
	options   ListIngestionJobsPaginatorOptions
	client    ListIngestionJobsAPIClient
	params    *ListIngestionJobsInput
	nextToken *string
	firstPage bool
}

// NewListIngestionJobsPaginator returns a new ListIngestionJobsPaginator
func NewListIngestionJobsPaginator(client ListIngestionJobsAPIClient, params *ListIngestionJobsInput, optFns ...func(*ListIngestionJobsPaginatorOptions)) *ListIngestionJobsPaginator {
	if params == nil {
		params = &ListIngestionJobsInput{}
	}

	options := ListIngestionJobsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListIngestionJobsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListIngestionJobsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListIngestionJobs page.
func (p *ListIngestionJobsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListIngestionJobsOutput, error) {
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

	result, err := p.client.ListIngestionJobs(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListIngestionJobs(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListIngestionJobs",
	}
}
