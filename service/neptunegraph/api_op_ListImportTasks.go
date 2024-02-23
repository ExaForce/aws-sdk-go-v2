// Code generated by smithy-go-codegen DO NOT EDIT.

package neptunegraph

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/neptunegraph/types"
	"github.com/aws/smithy-go/middleware"
	"github.com/aws/smithy-go/ptr"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists import tasks.
func (c *Client) ListImportTasks(ctx context.Context, params *ListImportTasksInput, optFns ...func(*Options)) (*ListImportTasksOutput, error) {
	if params == nil {
		params = &ListImportTasksInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListImportTasks", params, optFns, c.addOperationListImportTasksMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListImportTasksOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListImportTasksInput struct {

	// The total number of records to return in the command's output. If the total
	// number of records available is more than the value specified, nextToken is
	// provided in the command's output. To resume pagination, provide the nextToken
	// output value in the nextToken argument of a subsequent command. Do not use the
	// nextToken response element directly outside of the Amazon CLI.
	MaxResults *int32

	// Pagination token used to paginate output. When this value is provided as input,
	// the service returns results from where the previous response left off. When this
	// value is present in output, it indicates that there are more results to
	// retrieve.
	NextToken *string

	noSmithyDocumentSerde
}

func (in *ListImportTasksInput) bindEndpointParams(p *EndpointParameters) {

	p.ApiType = ptr.String("ControlPlane")
}

type ListImportTasksOutput struct {

	// The requested list of import tasks.
	//
	// This member is required.
	Tasks []types.ImportTaskSummary

	// Pagination token used to paginate output. When this value is provided as input,
	// the service returns results from where the previous response left off. When this
	// value is present in output, it indicates that there are more results to
	// retrieve.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListImportTasksMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListImportTasks{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListImportTasks{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListImportTasks"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListImportTasks(options.Region), middleware.Before); err != nil {
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

// ListImportTasksAPIClient is a client that implements the ListImportTasks
// operation.
type ListImportTasksAPIClient interface {
	ListImportTasks(context.Context, *ListImportTasksInput, ...func(*Options)) (*ListImportTasksOutput, error)
}

var _ ListImportTasksAPIClient = (*Client)(nil)

// ListImportTasksPaginatorOptions is the paginator options for ListImportTasks
type ListImportTasksPaginatorOptions struct {
	// The total number of records to return in the command's output. If the total
	// number of records available is more than the value specified, nextToken is
	// provided in the command's output. To resume pagination, provide the nextToken
	// output value in the nextToken argument of a subsequent command. Do not use the
	// nextToken response element directly outside of the Amazon CLI.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListImportTasksPaginator is a paginator for ListImportTasks
type ListImportTasksPaginator struct {
	options   ListImportTasksPaginatorOptions
	client    ListImportTasksAPIClient
	params    *ListImportTasksInput
	nextToken *string
	firstPage bool
}

// NewListImportTasksPaginator returns a new ListImportTasksPaginator
func NewListImportTasksPaginator(client ListImportTasksAPIClient, params *ListImportTasksInput, optFns ...func(*ListImportTasksPaginatorOptions)) *ListImportTasksPaginator {
	if params == nil {
		params = &ListImportTasksInput{}
	}

	options := ListImportTasksPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListImportTasksPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListImportTasksPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListImportTasks page.
func (p *ListImportTasksPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListImportTasksOutput, error) {
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

	result, err := p.client.ListImportTasks(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListImportTasks(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListImportTasks",
	}
}
