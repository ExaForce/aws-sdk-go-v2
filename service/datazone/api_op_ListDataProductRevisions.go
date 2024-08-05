// Code generated by smithy-go-codegen DO NOT EDIT.

package datazone

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/datazone/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists data product revisions.
func (c *Client) ListDataProductRevisions(ctx context.Context, params *ListDataProductRevisionsInput, optFns ...func(*Options)) (*ListDataProductRevisionsOutput, error) {
	if params == nil {
		params = &ListDataProductRevisionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListDataProductRevisions", params, optFns, c.addOperationListDataProductRevisionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListDataProductRevisionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListDataProductRevisionsInput struct {

	// The ID of the domain of the data product revisions that you want to list.
	//
	// This member is required.
	DomainIdentifier *string

	// The ID of the data product revision.
	//
	// This member is required.
	Identifier *string

	// The maximum number of asset filters to return in a single call to
	// ListDataProductRevisions . When the number of data product revisions to be
	// listed is greater than the value of MaxResults , the response contains a
	// NextToken value that you can use in a subsequent call to
	// ListDataProductRevisions to list the next set of data product revisions.
	MaxResults *int32

	// When the number of data product revisions is greater than the default value for
	// the MaxResults parameter, or if you explicitly specify a value for MaxResults
	// that is less than the number of data product revisions, the response includes a
	// pagination token named NextToken . You can specify this NextToken value in a
	// subsequent call to ListDataProductRevisions to list the next set of data
	// product revisions.
	NextToken *string

	noSmithyDocumentSerde
}

type ListDataProductRevisionsOutput struct {

	// The results of the ListDataProductRevisions action.
	//
	// This member is required.
	Items []types.DataProductRevision

	// When the number of data product revisions is greater than the default value for
	// the MaxResults parameter, or if you explicitly specify a value for MaxResults
	// that is less than the number of data product revisions, the response includes a
	// pagination token named NextToken . You can specify this NextToken value in a
	// subsequent call to ListDataProductRevisions to list the next set of data
	// product revisions.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListDataProductRevisionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListDataProductRevisions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListDataProductRevisions{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListDataProductRevisions"); err != nil {
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
	if err = addOpListDataProductRevisionsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListDataProductRevisions(options.Region), middleware.Before); err != nil {
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

// ListDataProductRevisionsPaginatorOptions is the paginator options for
// ListDataProductRevisions
type ListDataProductRevisionsPaginatorOptions struct {
	// The maximum number of asset filters to return in a single call to
	// ListDataProductRevisions . When the number of data product revisions to be
	// listed is greater than the value of MaxResults , the response contains a
	// NextToken value that you can use in a subsequent call to
	// ListDataProductRevisions to list the next set of data product revisions.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListDataProductRevisionsPaginator is a paginator for ListDataProductRevisions
type ListDataProductRevisionsPaginator struct {
	options   ListDataProductRevisionsPaginatorOptions
	client    ListDataProductRevisionsAPIClient
	params    *ListDataProductRevisionsInput
	nextToken *string
	firstPage bool
}

// NewListDataProductRevisionsPaginator returns a new
// ListDataProductRevisionsPaginator
func NewListDataProductRevisionsPaginator(client ListDataProductRevisionsAPIClient, params *ListDataProductRevisionsInput, optFns ...func(*ListDataProductRevisionsPaginatorOptions)) *ListDataProductRevisionsPaginator {
	if params == nil {
		params = &ListDataProductRevisionsInput{}
	}

	options := ListDataProductRevisionsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListDataProductRevisionsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListDataProductRevisionsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListDataProductRevisions page.
func (p *ListDataProductRevisionsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListDataProductRevisionsOutput, error) {
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

	optFns = append([]func(*Options){
		addIsPaginatorUserAgent,
	}, optFns...)
	result, err := p.client.ListDataProductRevisions(ctx, &params, optFns...)
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

// ListDataProductRevisionsAPIClient is a client that implements the
// ListDataProductRevisions operation.
type ListDataProductRevisionsAPIClient interface {
	ListDataProductRevisions(context.Context, *ListDataProductRevisionsInput, ...func(*Options)) (*ListDataProductRevisionsOutput, error)
}

var _ ListDataProductRevisionsAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opListDataProductRevisions(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListDataProductRevisions",
	}
}
