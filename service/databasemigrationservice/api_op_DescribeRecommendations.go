// Code generated by smithy-go-codegen DO NOT EDIT.

package databasemigrationservice

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/databasemigrationservice/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns a paginated list of target engine recommendations for your source
// databases.
func (c *Client) DescribeRecommendations(ctx context.Context, params *DescribeRecommendationsInput, optFns ...func(*Options)) (*DescribeRecommendationsOutput, error) {
	if params == nil {
		params = &DescribeRecommendationsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeRecommendations", params, optFns, c.addOperationDescribeRecommendationsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeRecommendationsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeRecommendationsInput struct {

	// Filters applied to the target engine recommendations described in the form of
	// key-value pairs.
	Filters []types.Filter

	// The maximum number of records to include in the response. If more records exist
	// than the specified MaxRecords value, Fleet Advisor includes a pagination token
	// in the response so that you can retrieve the remaining results.
	MaxRecords *int32

	// Specifies the unique pagination token that makes it possible to display the
	// next page of results. If this parameter is specified, the response includes only
	// records beyond the marker, up to the value specified by MaxRecords . If
	// NextToken is returned by a previous response, there are more results available.
	// The value of NextToken is a unique pagination token for each page. Make the
	// call again using the returned token to retrieve the next page. Keep all other
	// arguments unchanged.
	NextToken *string

	noSmithyDocumentSerde
}

type DescribeRecommendationsOutput struct {

	// The unique pagination token returned for you to pass to a subsequent request.
	// Fleet Advisor returns this token when the number of records in the response is
	// greater than the MaxRecords value. To retrieve the next page, make the call
	// again using the returned token and keeping all other arguments unchanged.
	NextToken *string

	// The list of recommendations of target engines that Fleet Advisor created for
	// the source database.
	Recommendations []types.Recommendation

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeRecommendationsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeRecommendations{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeRecommendations{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeRecommendations"); err != nil {
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
	if err = addOpDescribeRecommendationsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeRecommendations(options.Region), middleware.Before); err != nil {
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

// DescribeRecommendationsAPIClient is a client that implements the
// DescribeRecommendations operation.
type DescribeRecommendationsAPIClient interface {
	DescribeRecommendations(context.Context, *DescribeRecommendationsInput, ...func(*Options)) (*DescribeRecommendationsOutput, error)
}

var _ DescribeRecommendationsAPIClient = (*Client)(nil)

// DescribeRecommendationsPaginatorOptions is the paginator options for
// DescribeRecommendations
type DescribeRecommendationsPaginatorOptions struct {
	// The maximum number of records to include in the response. If more records exist
	// than the specified MaxRecords value, Fleet Advisor includes a pagination token
	// in the response so that you can retrieve the remaining results.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeRecommendationsPaginator is a paginator for DescribeRecommendations
type DescribeRecommendationsPaginator struct {
	options   DescribeRecommendationsPaginatorOptions
	client    DescribeRecommendationsAPIClient
	params    *DescribeRecommendationsInput
	nextToken *string
	firstPage bool
}

// NewDescribeRecommendationsPaginator returns a new
// DescribeRecommendationsPaginator
func NewDescribeRecommendationsPaginator(client DescribeRecommendationsAPIClient, params *DescribeRecommendationsInput, optFns ...func(*DescribeRecommendationsPaginatorOptions)) *DescribeRecommendationsPaginator {
	if params == nil {
		params = &DescribeRecommendationsInput{}
	}

	options := DescribeRecommendationsPaginatorOptions{}
	if params.MaxRecords != nil {
		options.Limit = *params.MaxRecords
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeRecommendationsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeRecommendationsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeRecommendations page.
func (p *DescribeRecommendationsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeRecommendationsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.MaxRecords = limit

	result, err := p.client.DescribeRecommendations(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opDescribeRecommendations(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeRecommendations",
	}
}
