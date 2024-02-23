// Code generated by smithy-go-codegen DO NOT EDIT.

package codestarnotifications

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/codestarnotifications/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns a list of the notification rule targets for an Amazon Web Services
// account.
func (c *Client) ListTargets(ctx context.Context, params *ListTargetsInput, optFns ...func(*Options)) (*ListTargetsOutput, error) {
	if params == nil {
		params = &ListTargetsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListTargets", params, optFns, c.addOperationListTargetsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListTargetsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListTargetsInput struct {

	// The filters to use to return information by service or resource type. Valid
	// filters include target type, target address, and target status. A filter with
	// the same name can appear more than once when used with OR statements. Filters
	// with different names should be applied with AND statements.
	Filters []types.ListTargetsFilter

	// A non-negative integer used to limit the number of returned results. The
	// maximum number of results that can be returned is 100.
	MaxResults *int32

	// An enumeration token that, when provided in a request, returns the next batch
	// of the results.
	NextToken *string

	noSmithyDocumentSerde
}

type ListTargetsOutput struct {

	// An enumeration token that can be used in a request to return the next batch of
	// results.
	NextToken *string

	// The list of notification rule targets.
	Targets []types.TargetSummary

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListTargetsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListTargets{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListTargets{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListTargets"); err != nil {
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
	if err = addOpListTargetsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListTargets(options.Region), middleware.Before); err != nil {
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

// ListTargetsAPIClient is a client that implements the ListTargets operation.
type ListTargetsAPIClient interface {
	ListTargets(context.Context, *ListTargetsInput, ...func(*Options)) (*ListTargetsOutput, error)
}

var _ ListTargetsAPIClient = (*Client)(nil)

// ListTargetsPaginatorOptions is the paginator options for ListTargets
type ListTargetsPaginatorOptions struct {
	// A non-negative integer used to limit the number of returned results. The
	// maximum number of results that can be returned is 100.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListTargetsPaginator is a paginator for ListTargets
type ListTargetsPaginator struct {
	options   ListTargetsPaginatorOptions
	client    ListTargetsAPIClient
	params    *ListTargetsInput
	nextToken *string
	firstPage bool
}

// NewListTargetsPaginator returns a new ListTargetsPaginator
func NewListTargetsPaginator(client ListTargetsAPIClient, params *ListTargetsInput, optFns ...func(*ListTargetsPaginatorOptions)) *ListTargetsPaginator {
	if params == nil {
		params = &ListTargetsInput{}
	}

	options := ListTargetsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListTargetsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListTargetsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListTargets page.
func (p *ListTargetsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListTargetsOutput, error) {
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

	result, err := p.client.ListTargets(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListTargets(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListTargets",
	}
}
