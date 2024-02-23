// Code generated by smithy-go-codegen DO NOT EDIT.

package securityhub

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/securityhub/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns a list of the standards that are currently enabled.
func (c *Client) GetEnabledStandards(ctx context.Context, params *GetEnabledStandardsInput, optFns ...func(*Options)) (*GetEnabledStandardsOutput, error) {
	if params == nil {
		params = &GetEnabledStandardsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetEnabledStandards", params, optFns, c.addOperationGetEnabledStandardsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetEnabledStandardsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetEnabledStandardsInput struct {

	// The maximum number of results to return in the response.
	MaxResults *int32

	// The token that is required for pagination. On your first call to the
	// GetEnabledStandards operation, set the value of this parameter to NULL . For
	// subsequent calls to the operation, to continue listing data, set the value of
	// this parameter to the value returned from the previous response.
	NextToken *string

	// The list of the standards subscription ARNs for the standards to retrieve.
	StandardsSubscriptionArns []string

	noSmithyDocumentSerde
}

type GetEnabledStandardsOutput struct {

	// The pagination token to use to request the next page of results.
	NextToken *string

	// The list of StandardsSubscriptions objects that include information about the
	// enabled standards.
	StandardsSubscriptions []types.StandardsSubscription

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetEnabledStandardsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetEnabledStandards{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetEnabledStandards{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetEnabledStandards"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetEnabledStandards(options.Region), middleware.Before); err != nil {
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

// GetEnabledStandardsAPIClient is a client that implements the
// GetEnabledStandards operation.
type GetEnabledStandardsAPIClient interface {
	GetEnabledStandards(context.Context, *GetEnabledStandardsInput, ...func(*Options)) (*GetEnabledStandardsOutput, error)
}

var _ GetEnabledStandardsAPIClient = (*Client)(nil)

// GetEnabledStandardsPaginatorOptions is the paginator options for
// GetEnabledStandards
type GetEnabledStandardsPaginatorOptions struct {
	// The maximum number of results to return in the response.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// GetEnabledStandardsPaginator is a paginator for GetEnabledStandards
type GetEnabledStandardsPaginator struct {
	options   GetEnabledStandardsPaginatorOptions
	client    GetEnabledStandardsAPIClient
	params    *GetEnabledStandardsInput
	nextToken *string
	firstPage bool
}

// NewGetEnabledStandardsPaginator returns a new GetEnabledStandardsPaginator
func NewGetEnabledStandardsPaginator(client GetEnabledStandardsAPIClient, params *GetEnabledStandardsInput, optFns ...func(*GetEnabledStandardsPaginatorOptions)) *GetEnabledStandardsPaginator {
	if params == nil {
		params = &GetEnabledStandardsInput{}
	}

	options := GetEnabledStandardsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &GetEnabledStandardsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *GetEnabledStandardsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next GetEnabledStandards page.
func (p *GetEnabledStandardsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*GetEnabledStandardsOutput, error) {
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

	result, err := p.client.GetEnabledStandards(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opGetEnabledStandards(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetEnabledStandards",
	}
}
