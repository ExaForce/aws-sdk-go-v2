// Code generated by smithy-go-codegen DO NOT EDIT.

package memorydb

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/memorydb/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists available reserved node offerings.
func (c *Client) DescribeReservedNodesOfferings(ctx context.Context, params *DescribeReservedNodesOfferingsInput, optFns ...func(*Options)) (*DescribeReservedNodesOfferingsOutput, error) {
	if params == nil {
		params = &DescribeReservedNodesOfferingsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeReservedNodesOfferings", params, optFns, c.addOperationDescribeReservedNodesOfferingsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeReservedNodesOfferingsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeReservedNodesOfferingsInput struct {

	// Duration filter value, specified in years or seconds. Use this parameter to
	// show only reservations for a given duration.
	Duration *string

	// The maximum number of records to include in the response. If more records exist
	// than the specified MaxRecords value, a marker is included in the response so
	// that the remaining results can be retrieved.
	MaxResults *int32

	// An optional marker returned from a prior request. Use this marker for
	// pagination of results from this operation. If this parameter is specified, the
	// response includes only records beyond the marker, up to the value specified by
	// MaxRecords.
	NextToken *string

	// The node type for the reserved nodes. For more information, see Supported node
	// types (https://docs.aws.amazon.com/memorydb/latest/devguide/nodes.reserved.html#reserved-nodes-supported)
	// .
	NodeType *string

	// The offering type filter value. Use this parameter to show only the available
	// offerings matching the specified offering type. Valid values: "All
	// Upfront"|"Partial Upfront"| "No Upfront"
	OfferingType *string

	// The offering identifier filter value. Use this parameter to show only the
	// available offering that matches the specified reservation identifier.
	ReservedNodesOfferingId *string

	noSmithyDocumentSerde
}

type DescribeReservedNodesOfferingsOutput struct {

	// An optional marker returned from a prior request. Use this marker for
	// pagination of results from this operation. If this parameter is specified, the
	// response includes only records beyond the marker, up to the value specified by
	// MaxRecords.
	NextToken *string

	// Lists available reserved node offerings.
	ReservedNodesOfferings []types.ReservedNodesOffering

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeReservedNodesOfferingsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeReservedNodesOfferings{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeReservedNodesOfferings{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeReservedNodesOfferings"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeReservedNodesOfferings(options.Region), middleware.Before); err != nil {
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

// DescribeReservedNodesOfferingsAPIClient is a client that implements the
// DescribeReservedNodesOfferings operation.
type DescribeReservedNodesOfferingsAPIClient interface {
	DescribeReservedNodesOfferings(context.Context, *DescribeReservedNodesOfferingsInput, ...func(*Options)) (*DescribeReservedNodesOfferingsOutput, error)
}

var _ DescribeReservedNodesOfferingsAPIClient = (*Client)(nil)

// DescribeReservedNodesOfferingsPaginatorOptions is the paginator options for
// DescribeReservedNodesOfferings
type DescribeReservedNodesOfferingsPaginatorOptions struct {
	// The maximum number of records to include in the response. If more records exist
	// than the specified MaxRecords value, a marker is included in the response so
	// that the remaining results can be retrieved.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeReservedNodesOfferingsPaginator is a paginator for
// DescribeReservedNodesOfferings
type DescribeReservedNodesOfferingsPaginator struct {
	options   DescribeReservedNodesOfferingsPaginatorOptions
	client    DescribeReservedNodesOfferingsAPIClient
	params    *DescribeReservedNodesOfferingsInput
	nextToken *string
	firstPage bool
}

// NewDescribeReservedNodesOfferingsPaginator returns a new
// DescribeReservedNodesOfferingsPaginator
func NewDescribeReservedNodesOfferingsPaginator(client DescribeReservedNodesOfferingsAPIClient, params *DescribeReservedNodesOfferingsInput, optFns ...func(*DescribeReservedNodesOfferingsPaginatorOptions)) *DescribeReservedNodesOfferingsPaginator {
	if params == nil {
		params = &DescribeReservedNodesOfferingsInput{}
	}

	options := DescribeReservedNodesOfferingsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeReservedNodesOfferingsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeReservedNodesOfferingsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeReservedNodesOfferings page.
func (p *DescribeReservedNodesOfferingsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeReservedNodesOfferingsOutput, error) {
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

	result, err := p.client.DescribeReservedNodesOfferings(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opDescribeReservedNodesOfferings(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeReservedNodesOfferings",
	}
}
