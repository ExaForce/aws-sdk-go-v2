// Code generated by smithy-go-codegen DO NOT EDIT.

package bedrock

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/bedrock/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// List the provisioned capacities. For more information, see Provisioned
// throughput (https://docs.aws.amazon.com/bedrock/latest/userguide/what-is-service.html)
// in the Bedrock User Guide.
func (c *Client) ListProvisionedModelThroughputs(ctx context.Context, params *ListProvisionedModelThroughputsInput, optFns ...func(*Options)) (*ListProvisionedModelThroughputsOutput, error) {
	if params == nil {
		params = &ListProvisionedModelThroughputsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListProvisionedModelThroughputs", params, optFns, c.addOperationListProvisionedModelThroughputsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListProvisionedModelThroughputsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListProvisionedModelThroughputsInput struct {

	// Return provisioned capacities created after the specified time.
	CreationTimeAfter *time.Time

	// Return provisioned capacities created before the specified time.
	CreationTimeBefore *time.Time

	// THe maximum number of results to return in the response.
	MaxResults *int32

	// Return the list of provisioned capacities where their model ARN is equal to
	// this parameter.
	ModelArnEquals *string

	// Return the list of provisioned capacities if their name contains these
	// characters.
	NameContains *string

	// Continuation token from the previous response, for Amazon Bedrock to list the
	// next set of results.
	NextToken *string

	// The field to sort by in the returned list of provisioned capacities.
	SortBy types.SortByProvisionedModels

	// The sort order of the results.
	SortOrder types.SortOrder

	// Return the list of provisioned capacities that match the specified status.
	StatusEquals types.ProvisionedModelStatus

	noSmithyDocumentSerde
}

type ListProvisionedModelThroughputsOutput struct {

	// Continuation token for the next request to list the next set of results.
	NextToken *string

	// List of summaries, one for each provisioned throughput in the response.
	ProvisionedModelSummaries []types.ProvisionedModelSummary

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListProvisionedModelThroughputsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListProvisionedModelThroughputs{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListProvisionedModelThroughputs{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListProvisionedModelThroughputs"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListProvisionedModelThroughputs(options.Region), middleware.Before); err != nil {
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

// ListProvisionedModelThroughputsAPIClient is a client that implements the
// ListProvisionedModelThroughputs operation.
type ListProvisionedModelThroughputsAPIClient interface {
	ListProvisionedModelThroughputs(context.Context, *ListProvisionedModelThroughputsInput, ...func(*Options)) (*ListProvisionedModelThroughputsOutput, error)
}

var _ ListProvisionedModelThroughputsAPIClient = (*Client)(nil)

// ListProvisionedModelThroughputsPaginatorOptions is the paginator options for
// ListProvisionedModelThroughputs
type ListProvisionedModelThroughputsPaginatorOptions struct {
	// THe maximum number of results to return in the response.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListProvisionedModelThroughputsPaginator is a paginator for
// ListProvisionedModelThroughputs
type ListProvisionedModelThroughputsPaginator struct {
	options   ListProvisionedModelThroughputsPaginatorOptions
	client    ListProvisionedModelThroughputsAPIClient
	params    *ListProvisionedModelThroughputsInput
	nextToken *string
	firstPage bool
}

// NewListProvisionedModelThroughputsPaginator returns a new
// ListProvisionedModelThroughputsPaginator
func NewListProvisionedModelThroughputsPaginator(client ListProvisionedModelThroughputsAPIClient, params *ListProvisionedModelThroughputsInput, optFns ...func(*ListProvisionedModelThroughputsPaginatorOptions)) *ListProvisionedModelThroughputsPaginator {
	if params == nil {
		params = &ListProvisionedModelThroughputsInput{}
	}

	options := ListProvisionedModelThroughputsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListProvisionedModelThroughputsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListProvisionedModelThroughputsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListProvisionedModelThroughputs page.
func (p *ListProvisionedModelThroughputsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListProvisionedModelThroughputsOutput, error) {
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

	result, err := p.client.ListProvisionedModelThroughputs(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListProvisionedModelThroughputs(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListProvisionedModelThroughputs",
	}
}
