// Code generated by smithy-go-codegen DO NOT EDIT.

package qapps

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/qapps/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists the library items for Amazon Q Apps that are published and available for
// users in your Amazon Web Services account.
func (c *Client) ListLibraryItems(ctx context.Context, params *ListLibraryItemsInput, optFns ...func(*Options)) (*ListLibraryItemsOutput, error) {
	if params == nil {
		params = &ListLibraryItemsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListLibraryItems", params, optFns, c.addOperationListLibraryItemsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListLibraryItemsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListLibraryItemsInput struct {

	// The unique identifier of the Amazon Q Business application environment instance.
	//
	// This member is required.
	InstanceId *string

	// Optional category to filter the library items by.
	CategoryId *string

	// The maximum number of library items to return in the response.
	Limit *int32

	// The token to request the next page of results.
	NextToken *string

	noSmithyDocumentSerde
}

type ListLibraryItemsOutput struct {

	// The list of library items meeting the request criteria.
	LibraryItems []types.LibraryItemMember

	// The token to use to request the next page of results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListLibraryItemsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListLibraryItems{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListLibraryItems{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListLibraryItems"); err != nil {
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
	if err = addTimeOffsetBuild(stack, c); err != nil {
		return err
	}
	if err = addUserAgentRetryMode(stack, options); err != nil {
		return err
	}
	if err = addOpListLibraryItemsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListLibraryItems(options.Region), middleware.Before); err != nil {
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

// ListLibraryItemsPaginatorOptions is the paginator options for ListLibraryItems
type ListLibraryItemsPaginatorOptions struct {
	// The maximum number of library items to return in the response.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListLibraryItemsPaginator is a paginator for ListLibraryItems
type ListLibraryItemsPaginator struct {
	options   ListLibraryItemsPaginatorOptions
	client    ListLibraryItemsAPIClient
	params    *ListLibraryItemsInput
	nextToken *string
	firstPage bool
}

// NewListLibraryItemsPaginator returns a new ListLibraryItemsPaginator
func NewListLibraryItemsPaginator(client ListLibraryItemsAPIClient, params *ListLibraryItemsInput, optFns ...func(*ListLibraryItemsPaginatorOptions)) *ListLibraryItemsPaginator {
	if params == nil {
		params = &ListLibraryItemsInput{}
	}

	options := ListLibraryItemsPaginatorOptions{}
	if params.Limit != nil {
		options.Limit = *params.Limit
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListLibraryItemsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListLibraryItemsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListLibraryItems page.
func (p *ListLibraryItemsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListLibraryItemsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.Limit = limit

	optFns = append([]func(*Options){
		addIsPaginatorUserAgent,
	}, optFns...)
	result, err := p.client.ListLibraryItems(ctx, &params, optFns...)
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

// ListLibraryItemsAPIClient is a client that implements the ListLibraryItems
// operation.
type ListLibraryItemsAPIClient interface {
	ListLibraryItems(context.Context, *ListLibraryItemsInput, ...func(*Options)) (*ListLibraryItemsOutput, error)
}

var _ ListLibraryItemsAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opListLibraryItems(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListLibraryItems",
	}
}
