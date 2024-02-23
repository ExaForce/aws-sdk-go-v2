// Code generated by smithy-go-codegen DO NOT EDIT.

package codebuild

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/codebuild/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Gets a list ARNs for the report groups in the current Amazon Web Services
// account.
func (c *Client) ListReportGroups(ctx context.Context, params *ListReportGroupsInput, optFns ...func(*Options)) (*ListReportGroupsOutput, error) {
	if params == nil {
		params = &ListReportGroupsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListReportGroups", params, optFns, c.addOperationListReportGroupsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListReportGroupsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListReportGroupsInput struct {

	// The maximum number of paginated report groups returned per response. Use
	// nextToken to iterate pages in the list of returned ReportGroup objects. The
	// default value is 100.
	MaxResults *int32

	// During a previous call, the maximum number of items that can be returned is the
	// value specified in maxResults . If there more items in the list, then a unique
	// string called a nextToken is returned. To get the next batch of items in the
	// list, call this operation again, adding the next token to the call. To get all
	// of the items in the list, keep calling this operation with each subsequent next
	// token that is returned, until no more next tokens are returned.
	NextToken *string

	// The criterion to be used to list build report groups. Valid values include:
	//   - CREATED_TIME : List based on when each report group was created.
	//   - LAST_MODIFIED_TIME : List based on when each report group was last changed.
	//   - NAME : List based on each report group's name.
	SortBy types.ReportGroupSortByType

	// Used to specify the order to sort the list of returned report groups. Valid
	// values are ASCENDING and DESCENDING .
	SortOrder types.SortOrderType

	noSmithyDocumentSerde
}

type ListReportGroupsOutput struct {

	// During a previous call, the maximum number of items that can be returned is the
	// value specified in maxResults . If there more items in the list, then a unique
	// string called a nextToken is returned. To get the next batch of items in the
	// list, call this operation again, adding the next token to the call. To get all
	// of the items in the list, keep calling this operation with each subsequent next
	// token that is returned, until no more next tokens are returned.
	NextToken *string

	// The list of ARNs for the report groups in the current Amazon Web Services
	// account.
	ReportGroups []string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListReportGroupsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListReportGroups{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListReportGroups{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListReportGroups"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListReportGroups(options.Region), middleware.Before); err != nil {
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

// ListReportGroupsAPIClient is a client that implements the ListReportGroups
// operation.
type ListReportGroupsAPIClient interface {
	ListReportGroups(context.Context, *ListReportGroupsInput, ...func(*Options)) (*ListReportGroupsOutput, error)
}

var _ ListReportGroupsAPIClient = (*Client)(nil)

// ListReportGroupsPaginatorOptions is the paginator options for ListReportGroups
type ListReportGroupsPaginatorOptions struct {
	// The maximum number of paginated report groups returned per response. Use
	// nextToken to iterate pages in the list of returned ReportGroup objects. The
	// default value is 100.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListReportGroupsPaginator is a paginator for ListReportGroups
type ListReportGroupsPaginator struct {
	options   ListReportGroupsPaginatorOptions
	client    ListReportGroupsAPIClient
	params    *ListReportGroupsInput
	nextToken *string
	firstPage bool
}

// NewListReportGroupsPaginator returns a new ListReportGroupsPaginator
func NewListReportGroupsPaginator(client ListReportGroupsAPIClient, params *ListReportGroupsInput, optFns ...func(*ListReportGroupsPaginatorOptions)) *ListReportGroupsPaginator {
	if params == nil {
		params = &ListReportGroupsInput{}
	}

	options := ListReportGroupsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListReportGroupsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListReportGroupsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListReportGroups page.
func (p *ListReportGroupsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListReportGroupsOutput, error) {
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

	result, err := p.client.ListReportGroups(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListReportGroups(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListReportGroups",
	}
}
