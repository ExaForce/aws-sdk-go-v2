// Code generated by smithy-go-codegen DO NOT EDIT.

package servicecatalog

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/servicecatalog/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Gets information about the products to which the caller has access.
func (c *Client) SearchProducts(ctx context.Context, params *SearchProductsInput, optFns ...func(*Options)) (*SearchProductsOutput, error) {
	if params == nil {
		params = &SearchProductsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "SearchProducts", params, optFns, c.addOperationSearchProductsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*SearchProductsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type SearchProductsInput struct {

	// The language code.
	//   - jp - Japanese
	//   - zh - Chinese
	AcceptLanguage *string

	// The search filters. If no search filters are specified, the output includes all
	// products to which the caller has access.
	Filters map[string][]string

	// The maximum number of items to return with this call.
	PageSize int32

	// The page token for the next set of results. To retrieve the first set of
	// results, use null.
	PageToken *string

	// The sort field. If no value is specified, the results are not sorted.
	SortBy types.ProductViewSortBy

	// The sort order. If no value is specified, the results are not sorted.
	SortOrder types.SortOrder

	noSmithyDocumentSerde
}

type SearchProductsOutput struct {

	// The page token to use to retrieve the next set of results. If there are no
	// additional results, this value is null.
	NextPageToken *string

	// The product view aggregations.
	ProductViewAggregations map[string][]types.ProductViewAggregationValue

	// Information about the product views.
	ProductViewSummaries []types.ProductViewSummary

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationSearchProductsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpSearchProducts{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpSearchProducts{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "SearchProducts"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opSearchProducts(options.Region), middleware.Before); err != nil {
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

// SearchProductsAPIClient is a client that implements the SearchProducts
// operation.
type SearchProductsAPIClient interface {
	SearchProducts(context.Context, *SearchProductsInput, ...func(*Options)) (*SearchProductsOutput, error)
}

var _ SearchProductsAPIClient = (*Client)(nil)

// SearchProductsPaginatorOptions is the paginator options for SearchProducts
type SearchProductsPaginatorOptions struct {
	// The maximum number of items to return with this call.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// SearchProductsPaginator is a paginator for SearchProducts
type SearchProductsPaginator struct {
	options   SearchProductsPaginatorOptions
	client    SearchProductsAPIClient
	params    *SearchProductsInput
	nextToken *string
	firstPage bool
}

// NewSearchProductsPaginator returns a new SearchProductsPaginator
func NewSearchProductsPaginator(client SearchProductsAPIClient, params *SearchProductsInput, optFns ...func(*SearchProductsPaginatorOptions)) *SearchProductsPaginator {
	if params == nil {
		params = &SearchProductsInput{}
	}

	options := SearchProductsPaginatorOptions{}
	if params.PageSize != 0 {
		options.Limit = params.PageSize
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &SearchProductsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.PageToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *SearchProductsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next SearchProducts page.
func (p *SearchProductsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*SearchProductsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.PageToken = p.nextToken

	params.PageSize = p.options.Limit

	result, err := p.client.SearchProducts(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.NextPageToken

	if p.options.StopOnDuplicateToken &&
		prevToken != nil &&
		p.nextToken != nil &&
		*prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

func newServiceMetadataMiddleware_opSearchProducts(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "SearchProducts",
	}
}
