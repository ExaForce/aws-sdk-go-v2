// Code generated by smithy-go-codegen DO NOT EDIT.

package datasync

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/datasync/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns a list of source and destination locations. If you have more locations
// than are returned in a response (that is, the response returns only a truncated
// list of your agents), the response contains a token that you can specify in your
// next request to fetch the next page of locations.
func (c *Client) ListLocations(ctx context.Context, params *ListLocationsInput, optFns ...func(*Options)) (*ListLocationsOutput, error) {
	if params == nil {
		params = &ListLocationsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListLocations", params, optFns, c.addOperationListLocationsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListLocationsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// ListLocationsRequest
type ListLocationsInput struct {

	// You can use API filters to narrow down the list of resources returned by
	// ListLocations . For example, to retrieve all tasks on a specific source
	// location, you can use ListLocations with filter name LocationType S3 and
	// Operator Equals .
	Filters []types.LocationFilter

	// The maximum number of locations to return.
	MaxResults *int32

	// An opaque string that indicates the position at which to begin the next list of
	// locations.
	NextToken *string

	noSmithyDocumentSerde
}

// ListLocationsResponse
type ListLocationsOutput struct {

	// An array that contains a list of locations.
	Locations []types.LocationListEntry

	// An opaque string that indicates the position at which to begin returning the
	// next list of locations.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListLocationsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListLocations{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListLocations{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListLocations"); err != nil {
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
	if err = addOpListLocationsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListLocations(options.Region), middleware.Before); err != nil {
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

// ListLocationsAPIClient is a client that implements the ListLocations operation.
type ListLocationsAPIClient interface {
	ListLocations(context.Context, *ListLocationsInput, ...func(*Options)) (*ListLocationsOutput, error)
}

var _ ListLocationsAPIClient = (*Client)(nil)

// ListLocationsPaginatorOptions is the paginator options for ListLocations
type ListLocationsPaginatorOptions struct {
	// The maximum number of locations to return.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListLocationsPaginator is a paginator for ListLocations
type ListLocationsPaginator struct {
	options   ListLocationsPaginatorOptions
	client    ListLocationsAPIClient
	params    *ListLocationsInput
	nextToken *string
	firstPage bool
}

// NewListLocationsPaginator returns a new ListLocationsPaginator
func NewListLocationsPaginator(client ListLocationsAPIClient, params *ListLocationsInput, optFns ...func(*ListLocationsPaginatorOptions)) *ListLocationsPaginator {
	if params == nil {
		params = &ListLocationsInput{}
	}

	options := ListLocationsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListLocationsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListLocationsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListLocations page.
func (p *ListLocationsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListLocationsOutput, error) {
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

	result, err := p.client.ListLocations(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListLocations(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListLocations",
	}
}
