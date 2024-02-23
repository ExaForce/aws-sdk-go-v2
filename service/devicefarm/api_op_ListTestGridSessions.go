// Code generated by smithy-go-codegen DO NOT EDIT.

package devicefarm

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/devicefarm/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Retrieves a list of sessions for a TestGridProject .
func (c *Client) ListTestGridSessions(ctx context.Context, params *ListTestGridSessionsInput, optFns ...func(*Options)) (*ListTestGridSessionsOutput, error) {
	if params == nil {
		params = &ListTestGridSessionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListTestGridSessions", params, optFns, c.addOperationListTestGridSessionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListTestGridSessionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListTestGridSessionsInput struct {

	// ARN of a TestGridProject .
	//
	// This member is required.
	ProjectArn *string

	// Return only sessions created after this time.
	CreationTimeAfter *time.Time

	// Return only sessions created before this time.
	CreationTimeBefore *time.Time

	// Return only sessions that ended after this time.
	EndTimeAfter *time.Time

	// Return only sessions that ended before this time.
	EndTimeBefore *time.Time

	// Return only this many results at a time.
	MaxResult *int32

	// Pagination token.
	NextToken *string

	// Return only sessions in this state.
	Status types.TestGridSessionStatus

	noSmithyDocumentSerde
}

type ListTestGridSessionsOutput struct {

	// Pagination token.
	NextToken *string

	// The sessions that match the criteria in a ListTestGridSessionsRequest .
	TestGridSessions []types.TestGridSession

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListTestGridSessionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListTestGridSessions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListTestGridSessions{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListTestGridSessions"); err != nil {
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
	if err = addOpListTestGridSessionsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListTestGridSessions(options.Region), middleware.Before); err != nil {
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

// ListTestGridSessionsAPIClient is a client that implements the
// ListTestGridSessions operation.
type ListTestGridSessionsAPIClient interface {
	ListTestGridSessions(context.Context, *ListTestGridSessionsInput, ...func(*Options)) (*ListTestGridSessionsOutput, error)
}

var _ ListTestGridSessionsAPIClient = (*Client)(nil)

// ListTestGridSessionsPaginatorOptions is the paginator options for
// ListTestGridSessions
type ListTestGridSessionsPaginatorOptions struct {
	// Return only this many results at a time.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListTestGridSessionsPaginator is a paginator for ListTestGridSessions
type ListTestGridSessionsPaginator struct {
	options   ListTestGridSessionsPaginatorOptions
	client    ListTestGridSessionsAPIClient
	params    *ListTestGridSessionsInput
	nextToken *string
	firstPage bool
}

// NewListTestGridSessionsPaginator returns a new ListTestGridSessionsPaginator
func NewListTestGridSessionsPaginator(client ListTestGridSessionsAPIClient, params *ListTestGridSessionsInput, optFns ...func(*ListTestGridSessionsPaginatorOptions)) *ListTestGridSessionsPaginator {
	if params == nil {
		params = &ListTestGridSessionsInput{}
	}

	options := ListTestGridSessionsPaginatorOptions{}
	if params.MaxResult != nil {
		options.Limit = *params.MaxResult
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListTestGridSessionsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListTestGridSessionsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListTestGridSessions page.
func (p *ListTestGridSessionsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListTestGridSessionsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.MaxResult = limit

	result, err := p.client.ListTestGridSessions(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListTestGridSessions(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListTestGridSessions",
	}
}
