// Code generated by smithy-go-codegen DO NOT EDIT.

package sagemaker

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/sagemaker/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Gets a list of past alerts in a model monitoring schedule.
func (c *Client) ListMonitoringAlertHistory(ctx context.Context, params *ListMonitoringAlertHistoryInput, optFns ...func(*Options)) (*ListMonitoringAlertHistoryOutput, error) {
	if params == nil {
		params = &ListMonitoringAlertHistoryInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListMonitoringAlertHistory", params, optFns, c.addOperationListMonitoringAlertHistoryMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListMonitoringAlertHistoryOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListMonitoringAlertHistoryInput struct {

	// A filter that returns only alerts created on or after the specified time.
	CreationTimeAfter *time.Time

	// A filter that returns only alerts created on or before the specified time.
	CreationTimeBefore *time.Time

	// The maximum number of results to display. The default is 100.
	MaxResults *int32

	// The name of a monitoring alert.
	MonitoringAlertName *string

	// The name of a monitoring schedule.
	MonitoringScheduleName *string

	// If the result of the previous ListMonitoringAlertHistory request was truncated,
	// the response includes a NextToken . To retrieve the next set of alerts in the
	// history, use the token in the next request.
	NextToken *string

	// The field used to sort results. The default is CreationTime .
	SortBy types.MonitoringAlertHistorySortKey

	// The sort order, whether Ascending or Descending , of the alert history. The
	// default is Descending .
	SortOrder types.SortOrder

	// A filter that retrieves only alerts with a specific status.
	StatusEquals types.MonitoringAlertStatus

	noSmithyDocumentSerde
}

type ListMonitoringAlertHistoryOutput struct {

	// An alert history for a model monitoring schedule.
	MonitoringAlertHistory []types.MonitoringAlertHistorySummary

	// If the response is truncated, SageMaker returns this token. To retrieve the
	// next set of alerts, use it in the subsequent request.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListMonitoringAlertHistoryMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListMonitoringAlertHistory{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListMonitoringAlertHistory{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListMonitoringAlertHistory"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListMonitoringAlertHistory(options.Region), middleware.Before); err != nil {
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

// ListMonitoringAlertHistoryAPIClient is a client that implements the
// ListMonitoringAlertHistory operation.
type ListMonitoringAlertHistoryAPIClient interface {
	ListMonitoringAlertHistory(context.Context, *ListMonitoringAlertHistoryInput, ...func(*Options)) (*ListMonitoringAlertHistoryOutput, error)
}

var _ ListMonitoringAlertHistoryAPIClient = (*Client)(nil)

// ListMonitoringAlertHistoryPaginatorOptions is the paginator options for
// ListMonitoringAlertHistory
type ListMonitoringAlertHistoryPaginatorOptions struct {
	// The maximum number of results to display. The default is 100.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListMonitoringAlertHistoryPaginator is a paginator for
// ListMonitoringAlertHistory
type ListMonitoringAlertHistoryPaginator struct {
	options   ListMonitoringAlertHistoryPaginatorOptions
	client    ListMonitoringAlertHistoryAPIClient
	params    *ListMonitoringAlertHistoryInput
	nextToken *string
	firstPage bool
}

// NewListMonitoringAlertHistoryPaginator returns a new
// ListMonitoringAlertHistoryPaginator
func NewListMonitoringAlertHistoryPaginator(client ListMonitoringAlertHistoryAPIClient, params *ListMonitoringAlertHistoryInput, optFns ...func(*ListMonitoringAlertHistoryPaginatorOptions)) *ListMonitoringAlertHistoryPaginator {
	if params == nil {
		params = &ListMonitoringAlertHistoryInput{}
	}

	options := ListMonitoringAlertHistoryPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListMonitoringAlertHistoryPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListMonitoringAlertHistoryPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListMonitoringAlertHistory page.
func (p *ListMonitoringAlertHistoryPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListMonitoringAlertHistoryOutput, error) {
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

	result, err := p.client.ListMonitoringAlertHistory(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListMonitoringAlertHistory(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListMonitoringAlertHistory",
	}
}
