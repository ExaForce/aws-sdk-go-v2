// Code generated by smithy-go-codegen DO NOT EDIT.

package route53recoverycontrolconfig

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns an array of all Amazon Route 53 health checks associated with a
// specific routing control.
func (c *Client) ListAssociatedRoute53HealthChecks(ctx context.Context, params *ListAssociatedRoute53HealthChecksInput, optFns ...func(*Options)) (*ListAssociatedRoute53HealthChecksOutput, error) {
	if params == nil {
		params = &ListAssociatedRoute53HealthChecksInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListAssociatedRoute53HealthChecks", params, optFns, c.addOperationListAssociatedRoute53HealthChecksMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListAssociatedRoute53HealthChecksOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListAssociatedRoute53HealthChecksInput struct {

	// The Amazon Resource Name (ARN) of the routing control.
	//
	// This member is required.
	RoutingControlArn *string

	// The number of objects that you want to return with this call.
	MaxResults *int32

	// The token that identifies which batch of results you want to see.
	NextToken *string

	noSmithyDocumentSerde
}

type ListAssociatedRoute53HealthChecksOutput struct {

	// Identifiers for the health checks.
	HealthCheckIds []string

	// Next token for listing health checks.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListAssociatedRoute53HealthChecksMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListAssociatedRoute53HealthChecks{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListAssociatedRoute53HealthChecks{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListAssociatedRoute53HealthChecks"); err != nil {
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
	if err = addOpListAssociatedRoute53HealthChecksValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListAssociatedRoute53HealthChecks(options.Region), middleware.Before); err != nil {
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

// ListAssociatedRoute53HealthChecksAPIClient is a client that implements the
// ListAssociatedRoute53HealthChecks operation.
type ListAssociatedRoute53HealthChecksAPIClient interface {
	ListAssociatedRoute53HealthChecks(context.Context, *ListAssociatedRoute53HealthChecksInput, ...func(*Options)) (*ListAssociatedRoute53HealthChecksOutput, error)
}

var _ ListAssociatedRoute53HealthChecksAPIClient = (*Client)(nil)

// ListAssociatedRoute53HealthChecksPaginatorOptions is the paginator options for
// ListAssociatedRoute53HealthChecks
type ListAssociatedRoute53HealthChecksPaginatorOptions struct {
	// The number of objects that you want to return with this call.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListAssociatedRoute53HealthChecksPaginator is a paginator for
// ListAssociatedRoute53HealthChecks
type ListAssociatedRoute53HealthChecksPaginator struct {
	options   ListAssociatedRoute53HealthChecksPaginatorOptions
	client    ListAssociatedRoute53HealthChecksAPIClient
	params    *ListAssociatedRoute53HealthChecksInput
	nextToken *string
	firstPage bool
}

// NewListAssociatedRoute53HealthChecksPaginator returns a new
// ListAssociatedRoute53HealthChecksPaginator
func NewListAssociatedRoute53HealthChecksPaginator(client ListAssociatedRoute53HealthChecksAPIClient, params *ListAssociatedRoute53HealthChecksInput, optFns ...func(*ListAssociatedRoute53HealthChecksPaginatorOptions)) *ListAssociatedRoute53HealthChecksPaginator {
	if params == nil {
		params = &ListAssociatedRoute53HealthChecksInput{}
	}

	options := ListAssociatedRoute53HealthChecksPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListAssociatedRoute53HealthChecksPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListAssociatedRoute53HealthChecksPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListAssociatedRoute53HealthChecks page.
func (p *ListAssociatedRoute53HealthChecksPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListAssociatedRoute53HealthChecksOutput, error) {
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

	result, err := p.client.ListAssociatedRoute53HealthChecks(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListAssociatedRoute53HealthChecks(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListAssociatedRoute53HealthChecks",
	}
}
