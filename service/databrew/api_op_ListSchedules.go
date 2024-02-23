// Code generated by smithy-go-codegen DO NOT EDIT.

package databrew

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/databrew/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists the DataBrew schedules that are defined.
func (c *Client) ListSchedules(ctx context.Context, params *ListSchedulesInput, optFns ...func(*Options)) (*ListSchedulesOutput, error) {
	if params == nil {
		params = &ListSchedulesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListSchedules", params, optFns, c.addOperationListSchedulesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListSchedulesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListSchedulesInput struct {

	// The name of the job that these schedules apply to.
	JobName *string

	// The maximum number of results to return in this request.
	MaxResults *int32

	// The token returned by a previous call to retrieve the next set of results.
	NextToken *string

	noSmithyDocumentSerde
}

type ListSchedulesOutput struct {

	// A list of schedules that are defined.
	//
	// This member is required.
	Schedules []types.Schedule

	// A token that you can use in a subsequent call to retrieve the next set of
	// results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListSchedulesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListSchedules{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListSchedules{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListSchedules"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListSchedules(options.Region), middleware.Before); err != nil {
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

// ListSchedulesAPIClient is a client that implements the ListSchedules operation.
type ListSchedulesAPIClient interface {
	ListSchedules(context.Context, *ListSchedulesInput, ...func(*Options)) (*ListSchedulesOutput, error)
}

var _ ListSchedulesAPIClient = (*Client)(nil)

// ListSchedulesPaginatorOptions is the paginator options for ListSchedules
type ListSchedulesPaginatorOptions struct {
	// The maximum number of results to return in this request.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListSchedulesPaginator is a paginator for ListSchedules
type ListSchedulesPaginator struct {
	options   ListSchedulesPaginatorOptions
	client    ListSchedulesAPIClient
	params    *ListSchedulesInput
	nextToken *string
	firstPage bool
}

// NewListSchedulesPaginator returns a new ListSchedulesPaginator
func NewListSchedulesPaginator(client ListSchedulesAPIClient, params *ListSchedulesInput, optFns ...func(*ListSchedulesPaginatorOptions)) *ListSchedulesPaginator {
	if params == nil {
		params = &ListSchedulesInput{}
	}

	options := ListSchedulesPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListSchedulesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListSchedulesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListSchedules page.
func (p *ListSchedulesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListSchedulesOutput, error) {
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

	result, err := p.client.ListSchedules(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListSchedules(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListSchedules",
	}
}
