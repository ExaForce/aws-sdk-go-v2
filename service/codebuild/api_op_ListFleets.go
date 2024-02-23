// Code generated by smithy-go-codegen DO NOT EDIT.

package codebuild

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/codebuild/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Gets a list of compute fleet names with each compute fleet name representing a
// single compute fleet.
func (c *Client) ListFleets(ctx context.Context, params *ListFleetsInput, optFns ...func(*Options)) (*ListFleetsOutput, error) {
	if params == nil {
		params = &ListFleetsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListFleets", params, optFns, c.addOperationListFleetsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListFleetsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListFleetsInput struct {

	// The maximum number of paginated compute fleets returned per response. Use
	// nextToken to iterate pages in the list of returned compute fleets.
	MaxResults *int32

	// During a previous call, if there are more than 100 items in the list, only the
	// first 100 items are returned, along with a unique string called a nextToken. To
	// get the next batch of items in the list, call this operation again, adding the
	// next token to the call. To get all of the items in the list, keep calling this
	// operation with each subsequent next token that is returned, until no more next
	// tokens are returned.
	NextToken *string

	// The criterion to be used to list compute fleet names. Valid values include:
	//   - CREATED_TIME : List based on when each compute fleet was created.
	//   - LAST_MODIFIED_TIME : List based on when information about each compute fleet
	//   was last changed.
	//   - NAME : List based on each compute fleet's name.
	// Use sortOrder to specify in what order to list the compute fleet names based on
	// the preceding criteria.
	SortBy types.FleetSortByType

	// The order in which to list compute fleets. Valid values include:
	//   - ASCENDING : List in ascending order.
	//   - DESCENDING : List in descending order.
	// Use sortBy to specify the criterion to be used to list compute fleet names.
	SortOrder types.SortOrderType

	noSmithyDocumentSerde
}

type ListFleetsOutput struct {

	// The list of compute fleet names.
	Fleets []string

	// If there are more than 100 items in the list, only the first 100 items are
	// returned, along with a unique string called a nextToken. To get the next batch
	// of items in the list, call this operation again, adding the next token to the
	// call.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListFleetsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListFleets{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListFleets{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListFleets"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListFleets(options.Region), middleware.Before); err != nil {
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

// ListFleetsAPIClient is a client that implements the ListFleets operation.
type ListFleetsAPIClient interface {
	ListFleets(context.Context, *ListFleetsInput, ...func(*Options)) (*ListFleetsOutput, error)
}

var _ ListFleetsAPIClient = (*Client)(nil)

// ListFleetsPaginatorOptions is the paginator options for ListFleets
type ListFleetsPaginatorOptions struct {
	// The maximum number of paginated compute fleets returned per response. Use
	// nextToken to iterate pages in the list of returned compute fleets.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListFleetsPaginator is a paginator for ListFleets
type ListFleetsPaginator struct {
	options   ListFleetsPaginatorOptions
	client    ListFleetsAPIClient
	params    *ListFleetsInput
	nextToken *string
	firstPage bool
}

// NewListFleetsPaginator returns a new ListFleetsPaginator
func NewListFleetsPaginator(client ListFleetsAPIClient, params *ListFleetsInput, optFns ...func(*ListFleetsPaginatorOptions)) *ListFleetsPaginator {
	if params == nil {
		params = &ListFleetsInput{}
	}

	options := ListFleetsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListFleetsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListFleetsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListFleets page.
func (p *ListFleetsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListFleetsOutput, error) {
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

	result, err := p.client.ListFleets(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListFleets(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListFleets",
	}
}
