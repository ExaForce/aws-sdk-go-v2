// Code generated by smithy-go-codegen DO NOT EDIT.

package deadline

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/deadline/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists tasks for a job.
func (c *Client) ListTasks(ctx context.Context, params *ListTasksInput, optFns ...func(*Options)) (*ListTasksOutput, error) {
	if params == nil {
		params = &ListTasksInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListTasks", params, optFns, c.addOperationListTasksMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListTasksOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListTasksInput struct {

	// The farm ID connected to the tasks.
	//
	// This member is required.
	FarmId *string

	// The job ID for the tasks.
	//
	// This member is required.
	JobId *string

	// The queue ID connected to the tasks.
	//
	// This member is required.
	QueueId *string

	// The step ID for the tasks.
	//
	// This member is required.
	StepId *string

	// The maximum number of results to return. Use this parameter with NextToken to
	// get results as a set of sequential pages.
	MaxResults *int32

	// The token for the next set of results, or null to start from the beginning.
	NextToken *string

	noSmithyDocumentSerde
}

type ListTasksOutput struct {

	// Tasks for the job.
	//
	// This member is required.
	Tasks []types.TaskSummary

	// If Deadline Cloud returns nextToken , then there are more results available. The
	// value of nextToken is a unique pagination token for each page. To retrieve the
	// next page, call the operation again using the returned token. Keep all other
	// arguments unchanged. If no results remain, then nextToken is set to null . Each
	// pagination token expires after 24 hours. If you provide a token that isn't
	// valid, then you receive an HTTP 400 ValidationException error.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListTasksMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListTasks{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListTasks{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListTasks"); err != nil {
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
	if err = addEndpointPrefix_opListTasksMiddleware(stack); err != nil {
		return err
	}
	if err = addOpListTasksValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListTasks(options.Region), middleware.Before); err != nil {
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

type endpointPrefix_opListTasksMiddleware struct {
}

func (*endpointPrefix_opListTasksMiddleware) ID() string {
	return "EndpointHostPrefix"
}

func (m *endpointPrefix_opListTasksMiddleware) HandleFinalize(ctx context.Context, in middleware.FinalizeInput, next middleware.FinalizeHandler) (
	out middleware.FinalizeOutput, metadata middleware.Metadata, err error,
) {
	if smithyhttp.GetHostnameImmutable(ctx) || smithyhttp.IsEndpointHostPrefixDisabled(ctx) {
		return next.HandleFinalize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	req.URL.Host = "management." + req.URL.Host

	return next.HandleFinalize(ctx, in)
}
func addEndpointPrefix_opListTasksMiddleware(stack *middleware.Stack) error {
	return stack.Finalize.Insert(&endpointPrefix_opListTasksMiddleware{}, "ResolveEndpointV2", middleware.After)
}

// ListTasksAPIClient is a client that implements the ListTasks operation.
type ListTasksAPIClient interface {
	ListTasks(context.Context, *ListTasksInput, ...func(*Options)) (*ListTasksOutput, error)
}

var _ ListTasksAPIClient = (*Client)(nil)

// ListTasksPaginatorOptions is the paginator options for ListTasks
type ListTasksPaginatorOptions struct {
	// The maximum number of results to return. Use this parameter with NextToken to
	// get results as a set of sequential pages.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListTasksPaginator is a paginator for ListTasks
type ListTasksPaginator struct {
	options   ListTasksPaginatorOptions
	client    ListTasksAPIClient
	params    *ListTasksInput
	nextToken *string
	firstPage bool
}

// NewListTasksPaginator returns a new ListTasksPaginator
func NewListTasksPaginator(client ListTasksAPIClient, params *ListTasksInput, optFns ...func(*ListTasksPaginatorOptions)) *ListTasksPaginator {
	if params == nil {
		params = &ListTasksInput{}
	}

	options := ListTasksPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListTasksPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListTasksPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListTasks page.
func (p *ListTasksPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListTasksOutput, error) {
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

	result, err := p.client.ListTasks(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListTasks(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListTasks",
	}
}
