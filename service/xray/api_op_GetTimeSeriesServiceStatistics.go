// Code generated by smithy-go-codegen DO NOT EDIT.

package xray

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/xray/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Get an aggregation of service statistics defined by a specific time range.
func (c *Client) GetTimeSeriesServiceStatistics(ctx context.Context, params *GetTimeSeriesServiceStatisticsInput, optFns ...func(*Options)) (*GetTimeSeriesServiceStatisticsOutput, error) {
	if params == nil {
		params = &GetTimeSeriesServiceStatisticsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetTimeSeriesServiceStatistics", params, optFns, c.addOperationGetTimeSeriesServiceStatisticsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetTimeSeriesServiceStatisticsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetTimeSeriesServiceStatisticsInput struct {

	// The end of the time frame for which to aggregate statistics.
	//
	// This member is required.
	EndTime *time.Time

	// The start of the time frame for which to aggregate statistics.
	//
	// This member is required.
	StartTime *time.Time

	// A filter expression defining entities that will be aggregated for statistics.
	// Supports ID, service, and edge functions. If no selector expression is
	// specified, edge statistics are returned.
	EntitySelectorExpression *string

	// The forecasted high and low fault count values. Forecast enabled requests
	// require the EntitySelectorExpression ID be provided.
	ForecastStatistics *bool

	// The Amazon Resource Name (ARN) of the group for which to pull statistics from.
	GroupARN *string

	// The case-sensitive name of the group for which to pull statistics from.
	GroupName *string

	// Pagination token.
	NextToken *string

	// Aggregation period in seconds.
	Period *int32

	noSmithyDocumentSerde
}

type GetTimeSeriesServiceStatisticsOutput struct {

	// A flag indicating whether or not a group's filter expression has been
	// consistent, or if a returned aggregation might show statistics from an older
	// version of the group's filter expression.
	ContainsOldGroupVersions bool

	// Pagination token.
	NextToken *string

	// The collection of statistics.
	TimeSeriesServiceStatistics []types.TimeSeriesServiceStatistics

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetTimeSeriesServiceStatisticsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetTimeSeriesServiceStatistics{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetTimeSeriesServiceStatistics{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetTimeSeriesServiceStatistics"); err != nil {
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
	if err = addOpGetTimeSeriesServiceStatisticsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetTimeSeriesServiceStatistics(options.Region), middleware.Before); err != nil {
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

// GetTimeSeriesServiceStatisticsAPIClient is a client that implements the
// GetTimeSeriesServiceStatistics operation.
type GetTimeSeriesServiceStatisticsAPIClient interface {
	GetTimeSeriesServiceStatistics(context.Context, *GetTimeSeriesServiceStatisticsInput, ...func(*Options)) (*GetTimeSeriesServiceStatisticsOutput, error)
}

var _ GetTimeSeriesServiceStatisticsAPIClient = (*Client)(nil)

// GetTimeSeriesServiceStatisticsPaginatorOptions is the paginator options for
// GetTimeSeriesServiceStatistics
type GetTimeSeriesServiceStatisticsPaginatorOptions struct {
	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// GetTimeSeriesServiceStatisticsPaginator is a paginator for
// GetTimeSeriesServiceStatistics
type GetTimeSeriesServiceStatisticsPaginator struct {
	options   GetTimeSeriesServiceStatisticsPaginatorOptions
	client    GetTimeSeriesServiceStatisticsAPIClient
	params    *GetTimeSeriesServiceStatisticsInput
	nextToken *string
	firstPage bool
}

// NewGetTimeSeriesServiceStatisticsPaginator returns a new
// GetTimeSeriesServiceStatisticsPaginator
func NewGetTimeSeriesServiceStatisticsPaginator(client GetTimeSeriesServiceStatisticsAPIClient, params *GetTimeSeriesServiceStatisticsInput, optFns ...func(*GetTimeSeriesServiceStatisticsPaginatorOptions)) *GetTimeSeriesServiceStatisticsPaginator {
	if params == nil {
		params = &GetTimeSeriesServiceStatisticsInput{}
	}

	options := GetTimeSeriesServiceStatisticsPaginatorOptions{}

	for _, fn := range optFns {
		fn(&options)
	}

	return &GetTimeSeriesServiceStatisticsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *GetTimeSeriesServiceStatisticsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next GetTimeSeriesServiceStatistics page.
func (p *GetTimeSeriesServiceStatisticsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*GetTimeSeriesServiceStatisticsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	result, err := p.client.GetTimeSeriesServiceStatistics(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opGetTimeSeriesServiceStatistics(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetTimeSeriesServiceStatistics",
	}
}
