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

// Gets a list of labeling jobs assigned to a specified work team.
func (c *Client) ListLabelingJobsForWorkteam(ctx context.Context, params *ListLabelingJobsForWorkteamInput, optFns ...func(*Options)) (*ListLabelingJobsForWorkteamOutput, error) {
	if params == nil {
		params = &ListLabelingJobsForWorkteamInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListLabelingJobsForWorkteam", params, optFns, c.addOperationListLabelingJobsForWorkteamMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListLabelingJobsForWorkteamOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListLabelingJobsForWorkteamInput struct {

	// The Amazon Resource Name (ARN) of the work team for which you want to see
	// labeling jobs for.
	//
	// This member is required.
	WorkteamArn *string

	// A filter that returns only labeling jobs created after the specified time
	// (timestamp).
	CreationTimeAfter *time.Time

	// A filter that returns only labeling jobs created before the specified time
	// (timestamp).
	CreationTimeBefore *time.Time

	// A filter the limits jobs to only the ones whose job reference code contains the
	// specified string.
	JobReferenceCodeContains *string

	// The maximum number of labeling jobs to return in each page of the response.
	MaxResults *int32

	// If the result of the previous ListLabelingJobsForWorkteam request was
	// truncated, the response includes a NextToken . To retrieve the next set of
	// labeling jobs, use the token in the next request.
	NextToken *string

	// The field to sort results by. The default is CreationTime .
	SortBy types.ListLabelingJobsForWorkteamSortByOptions

	// The sort order for results. The default is Ascending .
	SortOrder types.SortOrder

	noSmithyDocumentSerde
}

type ListLabelingJobsForWorkteamOutput struct {

	// An array of LabelingJobSummary objects, each describing a labeling job.
	//
	// This member is required.
	LabelingJobSummaryList []types.LabelingJobForWorkteamSummary

	// If the response is truncated, SageMaker returns this token. To retrieve the
	// next set of labeling jobs, use it in the subsequent request.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListLabelingJobsForWorkteamMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListLabelingJobsForWorkteam{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListLabelingJobsForWorkteam{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListLabelingJobsForWorkteam"); err != nil {
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
	if err = addOpListLabelingJobsForWorkteamValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListLabelingJobsForWorkteam(options.Region), middleware.Before); err != nil {
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

// ListLabelingJobsForWorkteamAPIClient is a client that implements the
// ListLabelingJobsForWorkteam operation.
type ListLabelingJobsForWorkteamAPIClient interface {
	ListLabelingJobsForWorkteam(context.Context, *ListLabelingJobsForWorkteamInput, ...func(*Options)) (*ListLabelingJobsForWorkteamOutput, error)
}

var _ ListLabelingJobsForWorkteamAPIClient = (*Client)(nil)

// ListLabelingJobsForWorkteamPaginatorOptions is the paginator options for
// ListLabelingJobsForWorkteam
type ListLabelingJobsForWorkteamPaginatorOptions struct {
	// The maximum number of labeling jobs to return in each page of the response.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListLabelingJobsForWorkteamPaginator is a paginator for
// ListLabelingJobsForWorkteam
type ListLabelingJobsForWorkteamPaginator struct {
	options   ListLabelingJobsForWorkteamPaginatorOptions
	client    ListLabelingJobsForWorkteamAPIClient
	params    *ListLabelingJobsForWorkteamInput
	nextToken *string
	firstPage bool
}

// NewListLabelingJobsForWorkteamPaginator returns a new
// ListLabelingJobsForWorkteamPaginator
func NewListLabelingJobsForWorkteamPaginator(client ListLabelingJobsForWorkteamAPIClient, params *ListLabelingJobsForWorkteamInput, optFns ...func(*ListLabelingJobsForWorkteamPaginatorOptions)) *ListLabelingJobsForWorkteamPaginator {
	if params == nil {
		params = &ListLabelingJobsForWorkteamInput{}
	}

	options := ListLabelingJobsForWorkteamPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListLabelingJobsForWorkteamPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListLabelingJobsForWorkteamPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListLabelingJobsForWorkteam page.
func (p *ListLabelingJobsForWorkteamPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListLabelingJobsForWorkteamOutput, error) {
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

	result, err := p.client.ListLabelingJobsForWorkteam(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListLabelingJobsForWorkteam(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListLabelingJobsForWorkteam",
	}
}
