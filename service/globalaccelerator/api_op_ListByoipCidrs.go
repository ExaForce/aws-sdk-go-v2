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

// Lists the IP address ranges that were specified in calls to ProvisionByoipCidr (https://docs.aws.amazon.com/global-accelerator/latest/api/ProvisionByoipCidr.html)
// , including the current state and a history of state changes.
func (c *Client) ListByoipCidrs(ctx context.Context, params *ListByoipCidrsInput, optFns ...func(*Options)) (*ListByoipCidrsOutput, error) {
	if params == nil {
		params = &ListByoipCidrsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListByoipCidrs", params, optFns, c.addOperationListByoipCidrsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListByoipCidrsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListByoipCidrsInput struct {

	// The maximum number of results to return with a single call. To retrieve the
	// remaining results, make another call with the returned nextToken value.
	MaxResults *int32

	// The token for the next page of results.
	NextToken *string

	noSmithyDocumentSerde
}

type ListByoipCidrsOutput struct {

	// Information about your address ranges.
	ByoipCidrs []types.ByoipCidr

	// The token for the next page of results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListByoipCidrsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListByoipCidrs{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListByoipCidrs{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListByoipCidrs"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListByoipCidrs(options.Region), middleware.Before); err != nil {
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

// ListByoipCidrsAPIClient is a client that implements the ListByoipCidrs
// operation.
type ListByoipCidrsAPIClient interface {
	ListByoipCidrs(context.Context, *ListByoipCidrsInput, ...func(*Options)) (*ListByoipCidrsOutput, error)
}

var _ ListByoipCidrsAPIClient = (*Client)(nil)

// ListByoipCidrsPaginatorOptions is the paginator options for ListByoipCidrs
type ListByoipCidrsPaginatorOptions struct {
	// The maximum number of results to return with a single call. To retrieve the
	// remaining results, make another call with the returned nextToken value.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListByoipCidrsPaginator is a paginator for ListByoipCidrs
type ListByoipCidrsPaginator struct {
	options   ListByoipCidrsPaginatorOptions
	client    ListByoipCidrsAPIClient
	params    *ListByoipCidrsInput
	nextToken *string
	firstPage bool
}

// NewListByoipCidrsPaginator returns a new ListByoipCidrsPaginator
func NewListByoipCidrsPaginator(client ListByoipCidrsAPIClient, params *ListByoipCidrsInput, optFns ...func(*ListByoipCidrsPaginatorOptions)) *ListByoipCidrsPaginator {
	if params == nil {
		params = &ListByoipCidrsInput{}
	}

	options := ListByoipCidrsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListByoipCidrsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListByoipCidrsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListByoipCidrs page.
func (p *ListByoipCidrsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListByoipCidrsOutput, error) {
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

	result, err := p.client.ListByoipCidrs(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListByoipCidrs(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListByoipCidrs",
	}
}
