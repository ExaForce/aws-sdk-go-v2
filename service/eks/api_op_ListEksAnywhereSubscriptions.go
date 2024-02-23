// Code generated by smithy-go-codegen DO NOT EDIT.

package eks

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/eks/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Displays the full description of the subscription.
func (c *Client) ListEksAnywhereSubscriptions(ctx context.Context, params *ListEksAnywhereSubscriptionsInput, optFns ...func(*Options)) (*ListEksAnywhereSubscriptionsOutput, error) {
	if params == nil {
		params = &ListEksAnywhereSubscriptionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListEksAnywhereSubscriptions", params, optFns, c.addOperationListEksAnywhereSubscriptionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListEksAnywhereSubscriptionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListEksAnywhereSubscriptionsInput struct {

	// An array of subscription statuses to filter on.
	IncludeStatus []types.EksAnywhereSubscriptionStatus

	// The maximum number of cluster results returned by ListEksAnywhereSubscriptions
	// in paginated output. When you use this parameter, ListEksAnywhereSubscriptions
	// returns only maxResults results in a single page along with a nextToken response
	// element. You can see the remaining results of the initial request by sending
	// another ListEksAnywhereSubscriptions request with the returned nextToken value.
	// This value can be between 1 and 100. If you don't use this parameter,
	// ListEksAnywhereSubscriptions returns up to 10 results and a nextToken value if
	// applicable.
	MaxResults *int32

	// The nextToken value returned from a previous paginated
	// ListEksAnywhereSubscriptions request where maxResults was used and the results
	// exceeded the value of that parameter. Pagination continues from the end of the
	// previous results that returned the nextToken value.
	NextToken *string

	noSmithyDocumentSerde
}

type ListEksAnywhereSubscriptionsOutput struct {

	// The nextToken value to include in a future ListEksAnywhereSubscriptions
	// request. When the results of a ListEksAnywhereSubscriptions request exceed
	// maxResults, you can use this value to retrieve the next page of results. This
	// value is null when there are no more results to return.
	NextToken *string

	// A list of all subscription objects in the region, filtered by includeStatus and
	// paginated by nextToken and maxResults.
	Subscriptions []types.EksAnywhereSubscription

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListEksAnywhereSubscriptionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListEksAnywhereSubscriptions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListEksAnywhereSubscriptions{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListEksAnywhereSubscriptions"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListEksAnywhereSubscriptions(options.Region), middleware.Before); err != nil {
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

// ListEksAnywhereSubscriptionsAPIClient is a client that implements the
// ListEksAnywhereSubscriptions operation.
type ListEksAnywhereSubscriptionsAPIClient interface {
	ListEksAnywhereSubscriptions(context.Context, *ListEksAnywhereSubscriptionsInput, ...func(*Options)) (*ListEksAnywhereSubscriptionsOutput, error)
}

var _ ListEksAnywhereSubscriptionsAPIClient = (*Client)(nil)

// ListEksAnywhereSubscriptionsPaginatorOptions is the paginator options for
// ListEksAnywhereSubscriptions
type ListEksAnywhereSubscriptionsPaginatorOptions struct {
	// The maximum number of cluster results returned by ListEksAnywhereSubscriptions
	// in paginated output. When you use this parameter, ListEksAnywhereSubscriptions
	// returns only maxResults results in a single page along with a nextToken response
	// element. You can see the remaining results of the initial request by sending
	// another ListEksAnywhereSubscriptions request with the returned nextToken value.
	// This value can be between 1 and 100. If you don't use this parameter,
	// ListEksAnywhereSubscriptions returns up to 10 results and a nextToken value if
	// applicable.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListEksAnywhereSubscriptionsPaginator is a paginator for
// ListEksAnywhereSubscriptions
type ListEksAnywhereSubscriptionsPaginator struct {
	options   ListEksAnywhereSubscriptionsPaginatorOptions
	client    ListEksAnywhereSubscriptionsAPIClient
	params    *ListEksAnywhereSubscriptionsInput
	nextToken *string
	firstPage bool
}

// NewListEksAnywhereSubscriptionsPaginator returns a new
// ListEksAnywhereSubscriptionsPaginator
func NewListEksAnywhereSubscriptionsPaginator(client ListEksAnywhereSubscriptionsAPIClient, params *ListEksAnywhereSubscriptionsInput, optFns ...func(*ListEksAnywhereSubscriptionsPaginatorOptions)) *ListEksAnywhereSubscriptionsPaginator {
	if params == nil {
		params = &ListEksAnywhereSubscriptionsInput{}
	}

	options := ListEksAnywhereSubscriptionsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListEksAnywhereSubscriptionsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListEksAnywhereSubscriptionsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListEksAnywhereSubscriptions page.
func (p *ListEksAnywhereSubscriptionsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListEksAnywhereSubscriptionsOutput, error) {
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

	result, err := p.client.ListEksAnywhereSubscriptions(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListEksAnywhereSubscriptions(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListEksAnywhereSubscriptions",
	}
}
