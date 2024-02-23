// Code generated by smithy-go-codegen DO NOT EDIT.

package personalize

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/personalize/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns a list of recommenders in a given Domain dataset group. When a Domain
// dataset group is not specified, all the recommenders associated with the account
// are listed. The response provides the properties for each recommender, including
// the Amazon Resource Name (ARN). For more information on recommenders, see
// CreateRecommender (https://docs.aws.amazon.com/personalize/latest/dg/API_CreateRecommender.html)
// .
func (c *Client) ListRecommenders(ctx context.Context, params *ListRecommendersInput, optFns ...func(*Options)) (*ListRecommendersOutput, error) {
	if params == nil {
		params = &ListRecommendersInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListRecommenders", params, optFns, c.addOperationListRecommendersMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListRecommendersOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListRecommendersInput struct {

	// The Amazon Resource Name (ARN) of the Domain dataset group to list the
	// recommenders for. When a Domain dataset group is not specified, all the
	// recommenders associated with the account are listed.
	DatasetGroupArn *string

	// The maximum number of recommenders to return.
	MaxResults *int32

	// A token returned from the previous call to ListRecommenders for getting the
	// next set of recommenders (if they exist).
	NextToken *string

	noSmithyDocumentSerde
}

type ListRecommendersOutput struct {

	// A token for getting the next set of recommenders (if they exist).
	NextToken *string

	// A list of the recommenders.
	Recommenders []types.RecommenderSummary

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListRecommendersMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListRecommenders{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListRecommenders{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListRecommenders"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListRecommenders(options.Region), middleware.Before); err != nil {
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

// ListRecommendersAPIClient is a client that implements the ListRecommenders
// operation.
type ListRecommendersAPIClient interface {
	ListRecommenders(context.Context, *ListRecommendersInput, ...func(*Options)) (*ListRecommendersOutput, error)
}

var _ ListRecommendersAPIClient = (*Client)(nil)

// ListRecommendersPaginatorOptions is the paginator options for ListRecommenders
type ListRecommendersPaginatorOptions struct {
	// The maximum number of recommenders to return.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListRecommendersPaginator is a paginator for ListRecommenders
type ListRecommendersPaginator struct {
	options   ListRecommendersPaginatorOptions
	client    ListRecommendersAPIClient
	params    *ListRecommendersInput
	nextToken *string
	firstPage bool
}

// NewListRecommendersPaginator returns a new ListRecommendersPaginator
func NewListRecommendersPaginator(client ListRecommendersAPIClient, params *ListRecommendersInput, optFns ...func(*ListRecommendersPaginatorOptions)) *ListRecommendersPaginator {
	if params == nil {
		params = &ListRecommendersInput{}
	}

	options := ListRecommendersPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListRecommendersPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListRecommendersPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListRecommenders page.
func (p *ListRecommendersPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListRecommendersOutput, error) {
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

	result, err := p.client.ListRecommenders(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListRecommenders(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListRecommenders",
	}
}
