// Code generated by smithy-go-codegen DO NOT EDIT.

package pipes

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/pipes/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Get the pipes associated with this account. For more information about pipes,
// see Amazon EventBridge Pipes (https://docs.aws.amazon.com/eventbridge/latest/userguide/eb-pipes.html)
// in the Amazon EventBridge User Guide.
func (c *Client) ListPipes(ctx context.Context, params *ListPipesInput, optFns ...func(*Options)) (*ListPipesOutput, error) {
	if params == nil {
		params = &ListPipesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListPipes", params, optFns, c.addOperationListPipesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListPipesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListPipesInput struct {

	// The state the pipe is in.
	CurrentState types.PipeState

	// The state the pipe should be in.
	DesiredState types.RequestedPipeState

	// The maximum number of pipes to include in the response.
	Limit *int32

	// A value that will return a subset of the pipes associated with this account.
	// For example, "NamePrefix": "ABC" will return all endpoints with "ABC" in the
	// name.
	NamePrefix *string

	// If nextToken is returned, there are more results available. The value of
	// nextToken is a unique pagination token for each page. Make the call again using
	// the returned token to retrieve the next page. Keep all other arguments
	// unchanged. Each pagination token expires after 24 hours. Using an expired
	// pagination token will return an HTTP 400 InvalidToken error.
	NextToken *string

	// The prefix matching the pipe source.
	SourcePrefix *string

	// The prefix matching the pipe target.
	TargetPrefix *string

	noSmithyDocumentSerde
}

type ListPipesOutput struct {

	// If nextToken is returned, there are more results available. The value of
	// nextToken is a unique pagination token for each page. Make the call again using
	// the returned token to retrieve the next page. Keep all other arguments
	// unchanged. Each pagination token expires after 24 hours. Using an expired
	// pagination token will return an HTTP 400 InvalidToken error.
	NextToken *string

	// The pipes returned by the call.
	Pipes []types.Pipe

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListPipesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListPipes{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListPipes{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListPipes"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListPipes(options.Region), middleware.Before); err != nil {
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

// ListPipesAPIClient is a client that implements the ListPipes operation.
type ListPipesAPIClient interface {
	ListPipes(context.Context, *ListPipesInput, ...func(*Options)) (*ListPipesOutput, error)
}

var _ ListPipesAPIClient = (*Client)(nil)

// ListPipesPaginatorOptions is the paginator options for ListPipes
type ListPipesPaginatorOptions struct {
	// The maximum number of pipes to include in the response.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListPipesPaginator is a paginator for ListPipes
type ListPipesPaginator struct {
	options   ListPipesPaginatorOptions
	client    ListPipesAPIClient
	params    *ListPipesInput
	nextToken *string
	firstPage bool
}

// NewListPipesPaginator returns a new ListPipesPaginator
func NewListPipesPaginator(client ListPipesAPIClient, params *ListPipesInput, optFns ...func(*ListPipesPaginatorOptions)) *ListPipesPaginator {
	if params == nil {
		params = &ListPipesInput{}
	}

	options := ListPipesPaginatorOptions{}
	if params.Limit != nil {
		options.Limit = *params.Limit
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListPipesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListPipesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListPipes page.
func (p *ListPipesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListPipesOutput, error) {
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

	result, err := p.client.ListPipes(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListPipes(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListPipes",
	}
}
