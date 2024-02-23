// Code generated by smithy-go-codegen DO NOT EDIT.

package wellarchitected

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/wellarchitected/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// List all milestones for an existing workload.
func (c *Client) ListMilestones(ctx context.Context, params *ListMilestonesInput, optFns ...func(*Options)) (*ListMilestonesOutput, error) {
	if params == nil {
		params = &ListMilestonesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListMilestones", params, optFns, c.addOperationListMilestonesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListMilestonesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Input to list all milestones for a workload.
type ListMilestonesInput struct {

	// The ID assigned to the workload. This ID is unique within an Amazon Web
	// Services Region.
	//
	// This member is required.
	WorkloadId *string

	// The maximum number of results to return for this request.
	MaxResults *int32

	// The token to use to retrieve the next set of results.
	NextToken *string

	noSmithyDocumentSerde
}

// Output of a list milestones call.
type ListMilestonesOutput struct {

	// A list of milestone summaries.
	MilestoneSummaries []types.MilestoneSummary

	// The token to use to retrieve the next set of results.
	NextToken *string

	// The ID assigned to the workload. This ID is unique within an Amazon Web
	// Services Region.
	WorkloadId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListMilestonesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListMilestones{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListMilestones{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListMilestones"); err != nil {
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
	if err = addOpListMilestonesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListMilestones(options.Region), middleware.Before); err != nil {
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

// ListMilestonesAPIClient is a client that implements the ListMilestones
// operation.
type ListMilestonesAPIClient interface {
	ListMilestones(context.Context, *ListMilestonesInput, ...func(*Options)) (*ListMilestonesOutput, error)
}

var _ ListMilestonesAPIClient = (*Client)(nil)

// ListMilestonesPaginatorOptions is the paginator options for ListMilestones
type ListMilestonesPaginatorOptions struct {
	// The maximum number of results to return for this request.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListMilestonesPaginator is a paginator for ListMilestones
type ListMilestonesPaginator struct {
	options   ListMilestonesPaginatorOptions
	client    ListMilestonesAPIClient
	params    *ListMilestonesInput
	nextToken *string
	firstPage bool
}

// NewListMilestonesPaginator returns a new ListMilestonesPaginator
func NewListMilestonesPaginator(client ListMilestonesAPIClient, params *ListMilestonesInput, optFns ...func(*ListMilestonesPaginatorOptions)) *ListMilestonesPaginator {
	if params == nil {
		params = &ListMilestonesInput{}
	}

	options := ListMilestonesPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListMilestonesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListMilestonesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListMilestones page.
func (p *ListMilestonesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListMilestonesOutput, error) {
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

	result, err := p.client.ListMilestones(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListMilestones(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListMilestones",
	}
}
