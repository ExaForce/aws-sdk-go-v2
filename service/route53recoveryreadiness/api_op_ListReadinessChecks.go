// Code generated by smithy-go-codegen DO NOT EDIT.

package route53recoveryreadiness

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/route53recoveryreadiness/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists the readiness checks for an account.
func (c *Client) ListReadinessChecks(ctx context.Context, params *ListReadinessChecksInput, optFns ...func(*Options)) (*ListReadinessChecksOutput, error) {
	if params == nil {
		params = &ListReadinessChecksInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListReadinessChecks", params, optFns, c.addOperationListReadinessChecksMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListReadinessChecksOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListReadinessChecksInput struct {

	// The number of objects that you want to return with this call.
	MaxResults *int32

	// The token that identifies which batch of results you want to see.
	NextToken *string

	noSmithyDocumentSerde
}

type ListReadinessChecksOutput struct {

	// The token that identifies which batch of results you want to see.
	NextToken *string

	// A list of readiness checks associated with the account.
	ReadinessChecks []types.ReadinessCheckOutput

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListReadinessChecksMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListReadinessChecks{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListReadinessChecks{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListReadinessChecks"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListReadinessChecks(options.Region), middleware.Before); err != nil {
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

// ListReadinessChecksAPIClient is a client that implements the
// ListReadinessChecks operation.
type ListReadinessChecksAPIClient interface {
	ListReadinessChecks(context.Context, *ListReadinessChecksInput, ...func(*Options)) (*ListReadinessChecksOutput, error)
}

var _ ListReadinessChecksAPIClient = (*Client)(nil)

// ListReadinessChecksPaginatorOptions is the paginator options for
// ListReadinessChecks
type ListReadinessChecksPaginatorOptions struct {
	// The number of objects that you want to return with this call.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListReadinessChecksPaginator is a paginator for ListReadinessChecks
type ListReadinessChecksPaginator struct {
	options   ListReadinessChecksPaginatorOptions
	client    ListReadinessChecksAPIClient
	params    *ListReadinessChecksInput
	nextToken *string
	firstPage bool
}

// NewListReadinessChecksPaginator returns a new ListReadinessChecksPaginator
func NewListReadinessChecksPaginator(client ListReadinessChecksAPIClient, params *ListReadinessChecksInput, optFns ...func(*ListReadinessChecksPaginatorOptions)) *ListReadinessChecksPaginator {
	if params == nil {
		params = &ListReadinessChecksInput{}
	}

	options := ListReadinessChecksPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListReadinessChecksPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListReadinessChecksPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListReadinessChecks page.
func (p *ListReadinessChecksPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListReadinessChecksOutput, error) {
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

	result, err := p.client.ListReadinessChecks(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListReadinessChecks(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListReadinessChecks",
	}
}
