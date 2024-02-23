// Code generated by smithy-go-codegen DO NOT EDIT.

package securityhub

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/securityhub/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Provides information about the associations for your configuration policies and
// self-managed behavior. Only the Security Hub delegated administrator can invoke
// this operation from the home Region.
func (c *Client) ListConfigurationPolicyAssociations(ctx context.Context, params *ListConfigurationPolicyAssociationsInput, optFns ...func(*Options)) (*ListConfigurationPolicyAssociationsOutput, error) {
	if params == nil {
		params = &ListConfigurationPolicyAssociationsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListConfigurationPolicyAssociations", params, optFns, c.addOperationListConfigurationPolicyAssociationsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListConfigurationPolicyAssociationsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListConfigurationPolicyAssociationsInput struct {

	// Options for filtering the ListConfigurationPolicyAssociations response. You can
	// filter by the Amazon Resource Name (ARN) or universally unique identifier (UUID)
	// of a configuration, AssociationType , or AssociationStatus .
	Filters *types.AssociationFilters

	// The maximum number of results that's returned by ListConfigurationPolicies in
	// each page of the response. When this parameter is used,
	// ListConfigurationPolicyAssociations returns the specified number of results in a
	// single page and a NextToken response element. You can see the remaining results
	// of the initial request by sending another ListConfigurationPolicyAssociations
	// request with the returned NextToken value. A valid range for MaxResults is
	// between 1 and 100.
	MaxResults *int32

	// The NextToken value that's returned from a previous paginated
	// ListConfigurationPolicyAssociations request where MaxResults was used but the
	// results exceeded the value of that parameter. Pagination continues from the end
	// of the previous response that returned the NextToken value. This value is null
	// when there are no more results to return.
	NextToken *string

	noSmithyDocumentSerde
}

type ListConfigurationPolicyAssociationsOutput struct {

	// An object that contains the details of each configuration policy association
	// that’s returned in a ListConfigurationPolicyAssociations request.
	ConfigurationPolicyAssociationSummaries []types.ConfigurationPolicyAssociationSummary

	// The NextToken value to include in the next ListConfigurationPolicyAssociations
	// request. When the results of a ListConfigurationPolicyAssociations request
	// exceed MaxResults , this value can be used to retrieve the next page of results.
	// This value is null when there are no more results to return.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListConfigurationPolicyAssociationsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListConfigurationPolicyAssociations{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListConfigurationPolicyAssociations{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListConfigurationPolicyAssociations"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListConfigurationPolicyAssociations(options.Region), middleware.Before); err != nil {
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

// ListConfigurationPolicyAssociationsAPIClient is a client that implements the
// ListConfigurationPolicyAssociations operation.
type ListConfigurationPolicyAssociationsAPIClient interface {
	ListConfigurationPolicyAssociations(context.Context, *ListConfigurationPolicyAssociationsInput, ...func(*Options)) (*ListConfigurationPolicyAssociationsOutput, error)
}

var _ ListConfigurationPolicyAssociationsAPIClient = (*Client)(nil)

// ListConfigurationPolicyAssociationsPaginatorOptions is the paginator options
// for ListConfigurationPolicyAssociations
type ListConfigurationPolicyAssociationsPaginatorOptions struct {
	// The maximum number of results that's returned by ListConfigurationPolicies in
	// each page of the response. When this parameter is used,
	// ListConfigurationPolicyAssociations returns the specified number of results in a
	// single page and a NextToken response element. You can see the remaining results
	// of the initial request by sending another ListConfigurationPolicyAssociations
	// request with the returned NextToken value. A valid range for MaxResults is
	// between 1 and 100.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListConfigurationPolicyAssociationsPaginator is a paginator for
// ListConfigurationPolicyAssociations
type ListConfigurationPolicyAssociationsPaginator struct {
	options   ListConfigurationPolicyAssociationsPaginatorOptions
	client    ListConfigurationPolicyAssociationsAPIClient
	params    *ListConfigurationPolicyAssociationsInput
	nextToken *string
	firstPage bool
}

// NewListConfigurationPolicyAssociationsPaginator returns a new
// ListConfigurationPolicyAssociationsPaginator
func NewListConfigurationPolicyAssociationsPaginator(client ListConfigurationPolicyAssociationsAPIClient, params *ListConfigurationPolicyAssociationsInput, optFns ...func(*ListConfigurationPolicyAssociationsPaginatorOptions)) *ListConfigurationPolicyAssociationsPaginator {
	if params == nil {
		params = &ListConfigurationPolicyAssociationsInput{}
	}

	options := ListConfigurationPolicyAssociationsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListConfigurationPolicyAssociationsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListConfigurationPolicyAssociationsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListConfigurationPolicyAssociations page.
func (p *ListConfigurationPolicyAssociationsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListConfigurationPolicyAssociationsOutput, error) {
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

	result, err := p.client.ListConfigurationPolicyAssociations(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListConfigurationPolicyAssociations(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListConfigurationPolicyAssociations",
	}
}
