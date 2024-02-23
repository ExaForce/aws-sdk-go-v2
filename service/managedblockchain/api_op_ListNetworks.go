// Code generated by smithy-go-codegen DO NOT EDIT.

package managedblockchain

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/managedblockchain/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns information about the networks in which the current Amazon Web Services
// account participates. Applies to Hyperledger Fabric and Ethereum.
func (c *Client) ListNetworks(ctx context.Context, params *ListNetworksInput, optFns ...func(*Options)) (*ListNetworksOutput, error) {
	if params == nil {
		params = &ListNetworksInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListNetworks", params, optFns, c.addOperationListNetworksMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListNetworksOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListNetworksInput struct {

	// An optional framework specifier. If provided, only networks of this framework
	// type are listed.
	Framework types.Framework

	// The maximum number of networks to list.
	MaxResults *int32

	// The name of the network.
	Name *string

	// The pagination token that indicates the next set of results to retrieve.
	NextToken *string

	// An optional status specifier. If provided, only networks currently in this
	// status are listed. Applies only to Hyperledger Fabric.
	Status types.NetworkStatus

	noSmithyDocumentSerde
}

type ListNetworksOutput struct {

	// An array of NetworkSummary objects that contain configuration properties for
	// each network.
	Networks []types.NetworkSummary

	// The pagination token that indicates the next set of results to retrieve.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListNetworksMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListNetworks{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListNetworks{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListNetworks"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListNetworks(options.Region), middleware.Before); err != nil {
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

// ListNetworksAPIClient is a client that implements the ListNetworks operation.
type ListNetworksAPIClient interface {
	ListNetworks(context.Context, *ListNetworksInput, ...func(*Options)) (*ListNetworksOutput, error)
}

var _ ListNetworksAPIClient = (*Client)(nil)

// ListNetworksPaginatorOptions is the paginator options for ListNetworks
type ListNetworksPaginatorOptions struct {
	// The maximum number of networks to list.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListNetworksPaginator is a paginator for ListNetworks
type ListNetworksPaginator struct {
	options   ListNetworksPaginatorOptions
	client    ListNetworksAPIClient
	params    *ListNetworksInput
	nextToken *string
	firstPage bool
}

// NewListNetworksPaginator returns a new ListNetworksPaginator
func NewListNetworksPaginator(client ListNetworksAPIClient, params *ListNetworksInput, optFns ...func(*ListNetworksPaginatorOptions)) *ListNetworksPaginator {
	if params == nil {
		params = &ListNetworksInput{}
	}

	options := ListNetworksPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListNetworksPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListNetworksPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListNetworks page.
func (p *ListNetworksPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListNetworksOutput, error) {
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

	result, err := p.client.ListNetworks(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListNetworks(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListNetworks",
	}
}
