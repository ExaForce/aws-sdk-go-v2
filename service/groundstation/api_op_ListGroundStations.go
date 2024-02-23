// Code generated by smithy-go-codegen DO NOT EDIT.

package groundstation

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/groundstation/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns a list of ground stations.
func (c *Client) ListGroundStations(ctx context.Context, params *ListGroundStationsInput, optFns ...func(*Options)) (*ListGroundStationsOutput, error) {
	if params == nil {
		params = &ListGroundStationsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListGroundStations", params, optFns, c.addOperationListGroundStationsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListGroundStationsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListGroundStationsInput struct {

	// Maximum number of ground stations returned.
	MaxResults *int32

	// Next token that can be supplied in the next call to get the next page of ground
	// stations.
	NextToken *string

	// Satellite ID to retrieve on-boarded ground stations.
	SatelliteId *string

	noSmithyDocumentSerde
}

type ListGroundStationsOutput struct {

	// List of ground stations.
	GroundStationList []types.GroundStationData

	// Next token that can be supplied in the next call to get the next page of ground
	// stations.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListGroundStationsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListGroundStations{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListGroundStations{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListGroundStations"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListGroundStations(options.Region), middleware.Before); err != nil {
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

// ListGroundStationsAPIClient is a client that implements the ListGroundStations
// operation.
type ListGroundStationsAPIClient interface {
	ListGroundStations(context.Context, *ListGroundStationsInput, ...func(*Options)) (*ListGroundStationsOutput, error)
}

var _ ListGroundStationsAPIClient = (*Client)(nil)

// ListGroundStationsPaginatorOptions is the paginator options for
// ListGroundStations
type ListGroundStationsPaginatorOptions struct {
	// Maximum number of ground stations returned.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListGroundStationsPaginator is a paginator for ListGroundStations
type ListGroundStationsPaginator struct {
	options   ListGroundStationsPaginatorOptions
	client    ListGroundStationsAPIClient
	params    *ListGroundStationsInput
	nextToken *string
	firstPage bool
}

// NewListGroundStationsPaginator returns a new ListGroundStationsPaginator
func NewListGroundStationsPaginator(client ListGroundStationsAPIClient, params *ListGroundStationsInput, optFns ...func(*ListGroundStationsPaginatorOptions)) *ListGroundStationsPaginator {
	if params == nil {
		params = &ListGroundStationsInput{}
	}

	options := ListGroundStationsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListGroundStationsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListGroundStationsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListGroundStations page.
func (p *ListGroundStationsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListGroundStationsOutput, error) {
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

	result, err := p.client.ListGroundStations(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListGroundStations(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListGroundStations",
	}
}
