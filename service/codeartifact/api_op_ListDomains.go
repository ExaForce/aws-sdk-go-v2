// Code generated by smithy-go-codegen DO NOT EDIT.

package codeartifact

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/codeartifact/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns a list of DomainSummary (https://docs.aws.amazon.com/codeartifact/latest/APIReference/API_PackageVersionDescription.html)
// objects for all domains owned by the Amazon Web Services account that makes this
// call. Each returned DomainSummary object contains information about a domain.
func (c *Client) ListDomains(ctx context.Context, params *ListDomainsInput, optFns ...func(*Options)) (*ListDomainsOutput, error) {
	if params == nil {
		params = &ListDomainsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListDomains", params, optFns, c.addOperationListDomainsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListDomainsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListDomainsInput struct {

	// The maximum number of results to return per page.
	MaxResults *int32

	// The token for the next set of results. Use the value returned in the previous
	// response in the next request to retrieve the next set of results.
	NextToken *string

	noSmithyDocumentSerde
}

type ListDomainsOutput struct {

	// The returned list of DomainSummary (https://docs.aws.amazon.com/codeartifact/latest/APIReference/API_DomainSummary.html)
	// objects.
	Domains []types.DomainSummary

	// The token for the next set of results. Use the value returned in the previous
	// response in the next request to retrieve the next set of results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListDomainsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListDomains{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListDomains{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListDomains"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListDomains(options.Region), middleware.Before); err != nil {
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

// ListDomainsAPIClient is a client that implements the ListDomains operation.
type ListDomainsAPIClient interface {
	ListDomains(context.Context, *ListDomainsInput, ...func(*Options)) (*ListDomainsOutput, error)
}

var _ ListDomainsAPIClient = (*Client)(nil)

// ListDomainsPaginatorOptions is the paginator options for ListDomains
type ListDomainsPaginatorOptions struct {
	// The maximum number of results to return per page.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListDomainsPaginator is a paginator for ListDomains
type ListDomainsPaginator struct {
	options   ListDomainsPaginatorOptions
	client    ListDomainsAPIClient
	params    *ListDomainsInput
	nextToken *string
	firstPage bool
}

// NewListDomainsPaginator returns a new ListDomainsPaginator
func NewListDomainsPaginator(client ListDomainsAPIClient, params *ListDomainsInput, optFns ...func(*ListDomainsPaginatorOptions)) *ListDomainsPaginator {
	if params == nil {
		params = &ListDomainsInput{}
	}

	options := ListDomainsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListDomainsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListDomainsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListDomains page.
func (p *ListDomainsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListDomainsOutput, error) {
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

	result, err := p.client.ListDomains(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListDomains(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListDomains",
	}
}
