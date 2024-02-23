// Code generated by smithy-go-codegen DO NOT EDIT.

package controltower

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/controltower/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns the landing zone ARN for the landing zone deployed in your managed
// account. This API also creates an ARN for existing accounts that do not yet have
// a landing zone ARN. Returns one landing zone ARN.
func (c *Client) ListLandingZones(ctx context.Context, params *ListLandingZonesInput, optFns ...func(*Options)) (*ListLandingZonesOutput, error) {
	if params == nil {
		params = &ListLandingZonesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListLandingZones", params, optFns, c.addOperationListLandingZonesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListLandingZonesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListLandingZonesInput struct {

	// The maximum number of returned landing zone ARNs, which is one.
	MaxResults *int32

	// The token to continue the list from a previous API call with the same
	// parameters.
	NextToken *string

	noSmithyDocumentSerde
}

type ListLandingZonesOutput struct {

	// The ARN of the landing zone.
	//
	// This member is required.
	LandingZones []types.LandingZoneSummary

	// Retrieves the next page of results. If the string is empty, the response is the
	// end of the results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListLandingZonesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListLandingZones{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListLandingZones{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListLandingZones"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListLandingZones(options.Region), middleware.Before); err != nil {
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

// ListLandingZonesAPIClient is a client that implements the ListLandingZones
// operation.
type ListLandingZonesAPIClient interface {
	ListLandingZones(context.Context, *ListLandingZonesInput, ...func(*Options)) (*ListLandingZonesOutput, error)
}

var _ ListLandingZonesAPIClient = (*Client)(nil)

// ListLandingZonesPaginatorOptions is the paginator options for ListLandingZones
type ListLandingZonesPaginatorOptions struct {
	// The maximum number of returned landing zone ARNs, which is one.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListLandingZonesPaginator is a paginator for ListLandingZones
type ListLandingZonesPaginator struct {
	options   ListLandingZonesPaginatorOptions
	client    ListLandingZonesAPIClient
	params    *ListLandingZonesInput
	nextToken *string
	firstPage bool
}

// NewListLandingZonesPaginator returns a new ListLandingZonesPaginator
func NewListLandingZonesPaginator(client ListLandingZonesAPIClient, params *ListLandingZonesInput, optFns ...func(*ListLandingZonesPaginatorOptions)) *ListLandingZonesPaginator {
	if params == nil {
		params = &ListLandingZonesInput{}
	}

	options := ListLandingZonesPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListLandingZonesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListLandingZonesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListLandingZones page.
func (p *ListLandingZonesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListLandingZonesOutput, error) {
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

	result, err := p.client.ListLandingZones(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListLandingZones(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListLandingZones",
	}
}
