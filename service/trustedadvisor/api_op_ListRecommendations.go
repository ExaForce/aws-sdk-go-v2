// Code generated by smithy-go-codegen DO NOT EDIT.

package trustedadvisor

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/trustedadvisor/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// List a filterable set of Recommendations
func (c *Client) ListRecommendations(ctx context.Context, params *ListRecommendationsInput, optFns ...func(*Options)) (*ListRecommendationsOutput, error) {
	if params == nil {
		params = &ListRecommendationsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListRecommendations", params, optFns, c.addOperationListRecommendationsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListRecommendationsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListRecommendationsInput struct {

	// After the last update of the Recommendation
	AfterLastUpdatedAt *time.Time

	// The aws service associated with the Recommendation
	AwsService *string

	// Before the last update of the Recommendation
	BeforeLastUpdatedAt *time.Time

	// The check identifier of the Recommendation
	CheckIdentifier *string

	// The maximum number of results to return per page.
	MaxResults *int32

	// The token for the next set of results. Use the value returned in the previous
	// response in the next request to retrieve the next set of results.
	NextToken *string

	// The pillar of the Recommendation
	Pillar types.RecommendationPillar

	// The source of the Recommendation
	Source types.RecommendationSource

	// The status of the Recommendation
	Status types.RecommendationStatus

	// The type of the Recommendation
	Type types.RecommendationType

	noSmithyDocumentSerde
}

type ListRecommendationsOutput struct {

	// The list of Recommendations
	//
	// This member is required.
	RecommendationSummaries []types.RecommendationSummary

	// The token for the next set of results. Use the value returned in the previous
	// response in the next request to retrieve the next set of results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListRecommendationsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListRecommendations{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListRecommendations{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListRecommendations"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListRecommendations(options.Region), middleware.Before); err != nil {
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

// ListRecommendationsAPIClient is a client that implements the
// ListRecommendations operation.
type ListRecommendationsAPIClient interface {
	ListRecommendations(context.Context, *ListRecommendationsInput, ...func(*Options)) (*ListRecommendationsOutput, error)
}

var _ ListRecommendationsAPIClient = (*Client)(nil)

// ListRecommendationsPaginatorOptions is the paginator options for
// ListRecommendations
type ListRecommendationsPaginatorOptions struct {
	// The maximum number of results to return per page.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListRecommendationsPaginator is a paginator for ListRecommendations
type ListRecommendationsPaginator struct {
	options   ListRecommendationsPaginatorOptions
	client    ListRecommendationsAPIClient
	params    *ListRecommendationsInput
	nextToken *string
	firstPage bool
}

// NewListRecommendationsPaginator returns a new ListRecommendationsPaginator
func NewListRecommendationsPaginator(client ListRecommendationsAPIClient, params *ListRecommendationsInput, optFns ...func(*ListRecommendationsPaginatorOptions)) *ListRecommendationsPaginator {
	if params == nil {
		params = &ListRecommendationsInput{}
	}

	options := ListRecommendationsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListRecommendationsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListRecommendationsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListRecommendations page.
func (p *ListRecommendationsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListRecommendationsOutput, error) {
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

	result, err := p.client.ListRecommendations(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListRecommendations(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListRecommendations",
	}
}
