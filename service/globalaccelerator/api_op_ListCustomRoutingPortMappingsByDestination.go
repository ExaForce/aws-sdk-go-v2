// Code generated by smithy-go-codegen DO NOT EDIT.

package globalaccelerator

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/globalaccelerator/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// List the port mappings for a specific EC2 instance (destination) in a VPC
// subnet endpoint. The response is the mappings for one destination IP address.
// This is useful when your subnet endpoint has mappings that span multiple custom
// routing accelerators in your account, or for scenarios where you only want to
// list the port mappings for a specific destination instance.
func (c *Client) ListCustomRoutingPortMappingsByDestination(ctx context.Context, params *ListCustomRoutingPortMappingsByDestinationInput, optFns ...func(*Options)) (*ListCustomRoutingPortMappingsByDestinationOutput, error) {
	if params == nil {
		params = &ListCustomRoutingPortMappingsByDestinationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListCustomRoutingPortMappingsByDestination", params, optFns, c.addOperationListCustomRoutingPortMappingsByDestinationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListCustomRoutingPortMappingsByDestinationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListCustomRoutingPortMappingsByDestinationInput struct {

	// The endpoint IP address in a virtual private cloud (VPC) subnet for which you
	// want to receive back port mappings.
	//
	// This member is required.
	DestinationAddress *string

	// The ID for the virtual private cloud (VPC) subnet.
	//
	// This member is required.
	EndpointId *string

	// The number of destination port mappings that you want to return with this call.
	// The default value is 10.
	MaxResults *int32

	// The token for the next set of results. You receive this token from a previous
	// call.
	NextToken *string

	noSmithyDocumentSerde
}

type ListCustomRoutingPortMappingsByDestinationOutput struct {

	// The port mappings for the endpoint IP address that you specified in the request.
	DestinationPortMappings []types.DestinationPortMapping

	// The token for the next set of results. You receive this token from a previous
	// call.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListCustomRoutingPortMappingsByDestinationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListCustomRoutingPortMappingsByDestination{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListCustomRoutingPortMappingsByDestination{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListCustomRoutingPortMappingsByDestination"); err != nil {
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
	if err = addOpListCustomRoutingPortMappingsByDestinationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListCustomRoutingPortMappingsByDestination(options.Region), middleware.Before); err != nil {
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

// ListCustomRoutingPortMappingsByDestinationAPIClient is a client that implements
// the ListCustomRoutingPortMappingsByDestination operation.
type ListCustomRoutingPortMappingsByDestinationAPIClient interface {
	ListCustomRoutingPortMappingsByDestination(context.Context, *ListCustomRoutingPortMappingsByDestinationInput, ...func(*Options)) (*ListCustomRoutingPortMappingsByDestinationOutput, error)
}

var _ ListCustomRoutingPortMappingsByDestinationAPIClient = (*Client)(nil)

// ListCustomRoutingPortMappingsByDestinationPaginatorOptions is the paginator
// options for ListCustomRoutingPortMappingsByDestination
type ListCustomRoutingPortMappingsByDestinationPaginatorOptions struct {
	// The number of destination port mappings that you want to return with this call.
	// The default value is 10.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListCustomRoutingPortMappingsByDestinationPaginator is a paginator for
// ListCustomRoutingPortMappingsByDestination
type ListCustomRoutingPortMappingsByDestinationPaginator struct {
	options   ListCustomRoutingPortMappingsByDestinationPaginatorOptions
	client    ListCustomRoutingPortMappingsByDestinationAPIClient
	params    *ListCustomRoutingPortMappingsByDestinationInput
	nextToken *string
	firstPage bool
}

// NewListCustomRoutingPortMappingsByDestinationPaginator returns a new
// ListCustomRoutingPortMappingsByDestinationPaginator
func NewListCustomRoutingPortMappingsByDestinationPaginator(client ListCustomRoutingPortMappingsByDestinationAPIClient, params *ListCustomRoutingPortMappingsByDestinationInput, optFns ...func(*ListCustomRoutingPortMappingsByDestinationPaginatorOptions)) *ListCustomRoutingPortMappingsByDestinationPaginator {
	if params == nil {
		params = &ListCustomRoutingPortMappingsByDestinationInput{}
	}

	options := ListCustomRoutingPortMappingsByDestinationPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListCustomRoutingPortMappingsByDestinationPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListCustomRoutingPortMappingsByDestinationPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListCustomRoutingPortMappingsByDestination page.
func (p *ListCustomRoutingPortMappingsByDestinationPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListCustomRoutingPortMappingsByDestinationOutput, error) {
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

	result, err := p.client.ListCustomRoutingPortMappingsByDestination(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListCustomRoutingPortMappingsByDestination(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListCustomRoutingPortMappingsByDestination",
	}
}
