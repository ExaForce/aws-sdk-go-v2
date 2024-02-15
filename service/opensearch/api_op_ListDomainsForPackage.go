// Code generated by smithy-go-codegen DO NOT EDIT.

package opensearch

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/opensearch/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists all Amazon OpenSearch Service domains associated with a given package.
// For more information, see Custom packages for Amazon OpenSearch Service (https://docs.aws.amazon.com/opensearch-service/latest/developerguide/custom-packages.html)
// .
func (c *Client) ListDomainsForPackage(ctx context.Context, params *ListDomainsForPackageInput, optFns ...func(*Options)) (*ListDomainsForPackageOutput, error) {
	if params == nil {
		params = &ListDomainsForPackageInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListDomainsForPackage", params, optFns, c.addOperationListDomainsForPackageMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListDomainsForPackageOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Container for the request parameters to the ListDomainsForPackage operation.
type ListDomainsForPackageInput struct {

	// The unique identifier of the package for which to list associated domains.
	//
	// This member is required.
	PackageID *string

	// An optional parameter that specifies the maximum number of results to return.
	// You can use nextToken to get the next page of results.
	MaxResults int32

	// If your initial ListDomainsForPackage operation returns a nextToken , you can
	// include the returned nextToken in subsequent ListDomainsForPackage operations,
	// which returns results in the next page.
	NextToken *string

	noSmithyDocumentSerde
}

// Container for the response parameters to the ListDomainsForPackage operation.
type ListDomainsForPackageOutput struct {

	// Information about all domains associated with a package.
	DomainPackageDetailsList []types.DomainPackageDetails

	// When nextToken is returned, there are more results available. The value of
	// nextToken is a unique pagination token for each page. Send the request again
	// using the returned token to retrieve the next page.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListDomainsForPackageMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListDomainsForPackage{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListDomainsForPackage{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListDomainsForPackage"); err != nil {
		return fmt.Errorf("add protocol finalizers: %v", err)
	}

	if err = addlegacyEndpointContextSetter(stack, options); err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
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
	if err = addOpListDomainsForPackageValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListDomainsForPackage(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecursionDetection(stack); err != nil {
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

// ListDomainsForPackageAPIClient is a client that implements the
// ListDomainsForPackage operation.
type ListDomainsForPackageAPIClient interface {
	ListDomainsForPackage(context.Context, *ListDomainsForPackageInput, ...func(*Options)) (*ListDomainsForPackageOutput, error)
}

var _ ListDomainsForPackageAPIClient = (*Client)(nil)

// ListDomainsForPackagePaginatorOptions is the paginator options for
// ListDomainsForPackage
type ListDomainsForPackagePaginatorOptions struct {
	// An optional parameter that specifies the maximum number of results to return.
	// You can use nextToken to get the next page of results.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListDomainsForPackagePaginator is a paginator for ListDomainsForPackage
type ListDomainsForPackagePaginator struct {
	options   ListDomainsForPackagePaginatorOptions
	client    ListDomainsForPackageAPIClient
	params    *ListDomainsForPackageInput
	nextToken *string
	firstPage bool
}

// NewListDomainsForPackagePaginator returns a new ListDomainsForPackagePaginator
func NewListDomainsForPackagePaginator(client ListDomainsForPackageAPIClient, params *ListDomainsForPackageInput, optFns ...func(*ListDomainsForPackagePaginatorOptions)) *ListDomainsForPackagePaginator {
	if params == nil {
		params = &ListDomainsForPackageInput{}
	}

	options := ListDomainsForPackagePaginatorOptions{}
	if params.MaxResults != 0 {
		options.Limit = params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListDomainsForPackagePaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListDomainsForPackagePaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListDomainsForPackage page.
func (p *ListDomainsForPackagePaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListDomainsForPackageOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	params.MaxResults = p.options.Limit

	result, err := p.client.ListDomainsForPackage(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListDomainsForPackage(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListDomainsForPackage",
	}
}
