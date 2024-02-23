// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudfront

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/cloudfront/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// List streaming distributions.
func (c *Client) ListStreamingDistributions(ctx context.Context, params *ListStreamingDistributionsInput, optFns ...func(*Options)) (*ListStreamingDistributionsOutput, error) {
	if params == nil {
		params = &ListStreamingDistributionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListStreamingDistributions", params, optFns, c.addOperationListStreamingDistributionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListStreamingDistributionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The request to list your streaming distributions.
type ListStreamingDistributionsInput struct {

	// The value that you provided for the Marker request parameter.
	Marker *string

	// The value that you provided for the MaxItems request parameter.
	MaxItems *int32

	noSmithyDocumentSerde
}

// The returned result of the corresponding request.
type ListStreamingDistributionsOutput struct {

	// The StreamingDistributionList type.
	StreamingDistributionList *types.StreamingDistributionList

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListStreamingDistributionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestxml_serializeOpListStreamingDistributions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpListStreamingDistributions{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListStreamingDistributions"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListStreamingDistributions(options.Region), middleware.Before); err != nil {
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

// ListStreamingDistributionsAPIClient is a client that implements the
// ListStreamingDistributions operation.
type ListStreamingDistributionsAPIClient interface {
	ListStreamingDistributions(context.Context, *ListStreamingDistributionsInput, ...func(*Options)) (*ListStreamingDistributionsOutput, error)
}

var _ ListStreamingDistributionsAPIClient = (*Client)(nil)

// ListStreamingDistributionsPaginatorOptions is the paginator options for
// ListStreamingDistributions
type ListStreamingDistributionsPaginatorOptions struct {
	// The value that you provided for the MaxItems request parameter.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListStreamingDistributionsPaginator is a paginator for
// ListStreamingDistributions
type ListStreamingDistributionsPaginator struct {
	options   ListStreamingDistributionsPaginatorOptions
	client    ListStreamingDistributionsAPIClient
	params    *ListStreamingDistributionsInput
	nextToken *string
	firstPage bool
}

// NewListStreamingDistributionsPaginator returns a new
// ListStreamingDistributionsPaginator
func NewListStreamingDistributionsPaginator(client ListStreamingDistributionsAPIClient, params *ListStreamingDistributionsInput, optFns ...func(*ListStreamingDistributionsPaginatorOptions)) *ListStreamingDistributionsPaginator {
	if params == nil {
		params = &ListStreamingDistributionsInput{}
	}

	options := ListStreamingDistributionsPaginatorOptions{}
	if params.MaxItems != nil {
		options.Limit = *params.MaxItems
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListStreamingDistributionsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.Marker,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListStreamingDistributionsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListStreamingDistributions page.
func (p *ListStreamingDistributionsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListStreamingDistributionsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.Marker = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.MaxItems = limit

	result, err := p.client.ListStreamingDistributions(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = nil
	if result.StreamingDistributionList != nil {
		p.nextToken = result.StreamingDistributionList.NextMarker
	}

	if p.options.StopOnDuplicateToken &&
		prevToken != nil &&
		p.nextToken != nil &&
		*prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

func newServiceMetadataMiddleware_opListStreamingDistributions(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListStreamingDistributions",
	}
}
