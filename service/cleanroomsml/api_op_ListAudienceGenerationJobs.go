// Code generated by smithy-go-codegen DO NOT EDIT.

package cleanroomsml

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/cleanroomsml/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns a list of audience generation jobs.
func (c *Client) ListAudienceGenerationJobs(ctx context.Context, params *ListAudienceGenerationJobsInput, optFns ...func(*Options)) (*ListAudienceGenerationJobsOutput, error) {
	if params == nil {
		params = &ListAudienceGenerationJobsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListAudienceGenerationJobs", params, optFns, c.addOperationListAudienceGenerationJobsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListAudienceGenerationJobsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListAudienceGenerationJobsInput struct {

	// The identifier of the collaboration that contains the audience generation jobs
	// that you are interested in.
	CollaborationId *string

	// The Amazon Resource Name (ARN) of the configured audience model that was used
	// for the audience generation jobs that you are interested in.
	ConfiguredAudienceModelArn *string

	// The maximum size of the results that is returned per call.
	MaxResults *int32

	// The token value retrieved from a previous call to access the next page of
	// results.
	NextToken *string

	noSmithyDocumentSerde
}

type ListAudienceGenerationJobsOutput struct {

	// The audience generation jobs that match the request.
	//
	// This member is required.
	AudienceGenerationJobs []types.AudienceGenerationJobSummary

	// The token value retrieved from a previous call to access the next page of
	// results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListAudienceGenerationJobsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListAudienceGenerationJobs{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListAudienceGenerationJobs{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListAudienceGenerationJobs"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListAudienceGenerationJobs(options.Region), middleware.Before); err != nil {
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

// ListAudienceGenerationJobsAPIClient is a client that implements the
// ListAudienceGenerationJobs operation.
type ListAudienceGenerationJobsAPIClient interface {
	ListAudienceGenerationJobs(context.Context, *ListAudienceGenerationJobsInput, ...func(*Options)) (*ListAudienceGenerationJobsOutput, error)
}

var _ ListAudienceGenerationJobsAPIClient = (*Client)(nil)

// ListAudienceGenerationJobsPaginatorOptions is the paginator options for
// ListAudienceGenerationJobs
type ListAudienceGenerationJobsPaginatorOptions struct {
	// The maximum size of the results that is returned per call.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListAudienceGenerationJobsPaginator is a paginator for
// ListAudienceGenerationJobs
type ListAudienceGenerationJobsPaginator struct {
	options   ListAudienceGenerationJobsPaginatorOptions
	client    ListAudienceGenerationJobsAPIClient
	params    *ListAudienceGenerationJobsInput
	nextToken *string
	firstPage bool
}

// NewListAudienceGenerationJobsPaginator returns a new
// ListAudienceGenerationJobsPaginator
func NewListAudienceGenerationJobsPaginator(client ListAudienceGenerationJobsAPIClient, params *ListAudienceGenerationJobsInput, optFns ...func(*ListAudienceGenerationJobsPaginatorOptions)) *ListAudienceGenerationJobsPaginator {
	if params == nil {
		params = &ListAudienceGenerationJobsInput{}
	}

	options := ListAudienceGenerationJobsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListAudienceGenerationJobsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListAudienceGenerationJobsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListAudienceGenerationJobs page.
func (p *ListAudienceGenerationJobsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListAudienceGenerationJobsOutput, error) {
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

	result, err := p.client.ListAudienceGenerationJobs(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListAudienceGenerationJobs(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListAudienceGenerationJobs",
	}
}
