// Code generated by smithy-go-codegen DO NOT EDIT.

package qbusiness

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/qbusiness/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists the Amazon Q indices you have created.
func (c *Client) ListIndices(ctx context.Context, params *ListIndicesInput, optFns ...func(*Options)) (*ListIndicesOutput, error) {
	if params == nil {
		params = &ListIndicesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListIndices", params, optFns, c.addOperationListIndicesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListIndicesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListIndicesInput struct {

	// The identifier of the Amazon Q application connected to the index.
	//
	// This member is required.
	ApplicationId *string

	// The maximum number of indices to return.
	MaxResults *int32

	// If the maxResults response was incomplete because there is more data to
	// retrieve, Amazon Q returns a pagination token in the response. You can use this
	// pagination token to retrieve the next set of Amazon Q indices.
	NextToken *string

	noSmithyDocumentSerde
}

type ListIndicesOutput struct {

	// An array of information on the items in one or more indexes.
	Indices []types.Index

	// If the response is truncated, Amazon Q returns this token that you can use in
	// the subsequent request to retrieve the next set of indexes.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListIndicesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListIndices{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListIndices{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListIndices"); err != nil {
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
	if err = addOpListIndicesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListIndices(options.Region), middleware.Before); err != nil {
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

// ListIndicesAPIClient is a client that implements the ListIndices operation.
type ListIndicesAPIClient interface {
	ListIndices(context.Context, *ListIndicesInput, ...func(*Options)) (*ListIndicesOutput, error)
}

var _ ListIndicesAPIClient = (*Client)(nil)

// ListIndicesPaginatorOptions is the paginator options for ListIndices
type ListIndicesPaginatorOptions struct {
	// The maximum number of indices to return.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListIndicesPaginator is a paginator for ListIndices
type ListIndicesPaginator struct {
	options   ListIndicesPaginatorOptions
	client    ListIndicesAPIClient
	params    *ListIndicesInput
	nextToken *string
	firstPage bool
}

// NewListIndicesPaginator returns a new ListIndicesPaginator
func NewListIndicesPaginator(client ListIndicesAPIClient, params *ListIndicesInput, optFns ...func(*ListIndicesPaginatorOptions)) *ListIndicesPaginator {
	if params == nil {
		params = &ListIndicesInput{}
	}

	options := ListIndicesPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListIndicesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListIndicesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListIndices page.
func (p *ListIndicesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListIndicesOutput, error) {
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

	result, err := p.client.ListIndices(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListIndices(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListIndices",
	}
}
