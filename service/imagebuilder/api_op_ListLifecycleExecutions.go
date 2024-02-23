// Code generated by smithy-go-codegen DO NOT EDIT.

package imagebuilder

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/imagebuilder/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Get the lifecycle runtime history for the specified resource.
func (c *Client) ListLifecycleExecutions(ctx context.Context, params *ListLifecycleExecutionsInput, optFns ...func(*Options)) (*ListLifecycleExecutionsOutput, error) {
	if params == nil {
		params = &ListLifecycleExecutionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListLifecycleExecutions", params, optFns, c.addOperationListLifecycleExecutionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListLifecycleExecutionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListLifecycleExecutionsInput struct {

	// The Amazon Resource Name (ARN) of the resource for which to get a list of
	// lifecycle runtime instances.
	//
	// This member is required.
	ResourceArn *string

	// The maximum items to return in a request.
	MaxResults *int32

	// A token to specify where to start paginating. This is the nextToken from a
	// previously truncated response.
	NextToken *string

	noSmithyDocumentSerde
}

type ListLifecycleExecutionsOutput struct {

	// A list of lifecycle runtime instances for the specified resource.
	LifecycleExecutions []types.LifecycleExecution

	// The next token used for paginated responses. When this field isn't empty, there
	// are additional elements that the service hasn't included in this request. Use
	// this token with the next request to retrieve additional objects.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListLifecycleExecutionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListLifecycleExecutions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListLifecycleExecutions{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListLifecycleExecutions"); err != nil {
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
	if err = addOpListLifecycleExecutionsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListLifecycleExecutions(options.Region), middleware.Before); err != nil {
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

// ListLifecycleExecutionsAPIClient is a client that implements the
// ListLifecycleExecutions operation.
type ListLifecycleExecutionsAPIClient interface {
	ListLifecycleExecutions(context.Context, *ListLifecycleExecutionsInput, ...func(*Options)) (*ListLifecycleExecutionsOutput, error)
}

var _ ListLifecycleExecutionsAPIClient = (*Client)(nil)

// ListLifecycleExecutionsPaginatorOptions is the paginator options for
// ListLifecycleExecutions
type ListLifecycleExecutionsPaginatorOptions struct {
	// The maximum items to return in a request.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListLifecycleExecutionsPaginator is a paginator for ListLifecycleExecutions
type ListLifecycleExecutionsPaginator struct {
	options   ListLifecycleExecutionsPaginatorOptions
	client    ListLifecycleExecutionsAPIClient
	params    *ListLifecycleExecutionsInput
	nextToken *string
	firstPage bool
}

// NewListLifecycleExecutionsPaginator returns a new
// ListLifecycleExecutionsPaginator
func NewListLifecycleExecutionsPaginator(client ListLifecycleExecutionsAPIClient, params *ListLifecycleExecutionsInput, optFns ...func(*ListLifecycleExecutionsPaginatorOptions)) *ListLifecycleExecutionsPaginator {
	if params == nil {
		params = &ListLifecycleExecutionsInput{}
	}

	options := ListLifecycleExecutionsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListLifecycleExecutionsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListLifecycleExecutionsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListLifecycleExecutions page.
func (p *ListLifecycleExecutionsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListLifecycleExecutionsOutput, error) {
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

	result, err := p.client.ListLifecycleExecutions(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListLifecycleExecutions(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListLifecycleExecutions",
	}
}
