// Code generated by smithy-go-codegen DO NOT EDIT.

package mediapackagev2

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/mediapackagev2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves all channel groups that are configured in AWS Elemental MediaPackage,
// including the channels and origin endpoints that are associated with it.
func (c *Client) ListChannelGroups(ctx context.Context, params *ListChannelGroupsInput, optFns ...func(*Options)) (*ListChannelGroupsOutput, error) {
	if params == nil {
		params = &ListChannelGroupsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListChannelGroups", params, optFns, c.addOperationListChannelGroupsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListChannelGroupsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListChannelGroupsInput struct {

	// The maximum number of results to return in the response.
	MaxResults *int32

	// The pagination token from the GET list request. Use the token to fetch the next
	// page of results.
	NextToken *string

	noSmithyDocumentSerde
}

type ListChannelGroupsOutput struct {

	// The objects being returned.
	Items []types.ChannelGroupListConfiguration

	// The pagination token from the GET list request. Use the token to fetch the next
	// page of results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListChannelGroupsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListChannelGroups{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListChannelGroups{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListChannelGroups"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListChannelGroups(options.Region), middleware.Before); err != nil {
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

// ListChannelGroupsAPIClient is a client that implements the ListChannelGroups
// operation.
type ListChannelGroupsAPIClient interface {
	ListChannelGroups(context.Context, *ListChannelGroupsInput, ...func(*Options)) (*ListChannelGroupsOutput, error)
}

var _ ListChannelGroupsAPIClient = (*Client)(nil)

// ListChannelGroupsPaginatorOptions is the paginator options for ListChannelGroups
type ListChannelGroupsPaginatorOptions struct {
	// The maximum number of results to return in the response.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListChannelGroupsPaginator is a paginator for ListChannelGroups
type ListChannelGroupsPaginator struct {
	options   ListChannelGroupsPaginatorOptions
	client    ListChannelGroupsAPIClient
	params    *ListChannelGroupsInput
	nextToken *string
	firstPage bool
}

// NewListChannelGroupsPaginator returns a new ListChannelGroupsPaginator
func NewListChannelGroupsPaginator(client ListChannelGroupsAPIClient, params *ListChannelGroupsInput, optFns ...func(*ListChannelGroupsPaginatorOptions)) *ListChannelGroupsPaginator {
	if params == nil {
		params = &ListChannelGroupsInput{}
	}

	options := ListChannelGroupsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListChannelGroupsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListChannelGroupsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListChannelGroups page.
func (p *ListChannelGroupsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListChannelGroupsOutput, error) {
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

	result, err := p.client.ListChannelGroups(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListChannelGroups(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListChannelGroups",
	}
}
