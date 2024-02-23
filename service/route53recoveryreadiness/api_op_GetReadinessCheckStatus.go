// Code generated by smithy-go-codegen DO NOT EDIT.

package route53recoveryreadiness

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/route53recoveryreadiness/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Gets the readiness status for an individual readiness check. To see the overall
// readiness status for a recovery group, that considers the readiness status for
// all the readiness checks in a recovery group, use
// GetRecoveryGroupReadinessSummary.
func (c *Client) GetReadinessCheckStatus(ctx context.Context, params *GetReadinessCheckStatusInput, optFns ...func(*Options)) (*GetReadinessCheckStatusOutput, error) {
	if params == nil {
		params = &GetReadinessCheckStatusInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetReadinessCheckStatus", params, optFns, c.addOperationGetReadinessCheckStatusMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetReadinessCheckStatusOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetReadinessCheckStatusInput struct {

	// Name of a readiness check.
	//
	// This member is required.
	ReadinessCheckName *string

	// The number of objects that you want to return with this call.
	MaxResults *int32

	// The token that identifies which batch of results you want to see.
	NextToken *string

	noSmithyDocumentSerde
}

type GetReadinessCheckStatusOutput struct {

	// Top level messages for readiness check status
	Messages []types.Message

	// The token that identifies which batch of results you want to see.
	NextToken *string

	// The readiness at rule level.
	Readiness types.Readiness

	// Summary of the readiness of resources.
	Resources []types.ResourceResult

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetReadinessCheckStatusMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetReadinessCheckStatus{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetReadinessCheckStatus{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetReadinessCheckStatus"); err != nil {
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
	if err = addOpGetReadinessCheckStatusValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetReadinessCheckStatus(options.Region), middleware.Before); err != nil {
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

// GetReadinessCheckStatusAPIClient is a client that implements the
// GetReadinessCheckStatus operation.
type GetReadinessCheckStatusAPIClient interface {
	GetReadinessCheckStatus(context.Context, *GetReadinessCheckStatusInput, ...func(*Options)) (*GetReadinessCheckStatusOutput, error)
}

var _ GetReadinessCheckStatusAPIClient = (*Client)(nil)

// GetReadinessCheckStatusPaginatorOptions is the paginator options for
// GetReadinessCheckStatus
type GetReadinessCheckStatusPaginatorOptions struct {
	// The number of objects that you want to return with this call.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// GetReadinessCheckStatusPaginator is a paginator for GetReadinessCheckStatus
type GetReadinessCheckStatusPaginator struct {
	options   GetReadinessCheckStatusPaginatorOptions
	client    GetReadinessCheckStatusAPIClient
	params    *GetReadinessCheckStatusInput
	nextToken *string
	firstPage bool
}

// NewGetReadinessCheckStatusPaginator returns a new
// GetReadinessCheckStatusPaginator
func NewGetReadinessCheckStatusPaginator(client GetReadinessCheckStatusAPIClient, params *GetReadinessCheckStatusInput, optFns ...func(*GetReadinessCheckStatusPaginatorOptions)) *GetReadinessCheckStatusPaginator {
	if params == nil {
		params = &GetReadinessCheckStatusInput{}
	}

	options := GetReadinessCheckStatusPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &GetReadinessCheckStatusPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *GetReadinessCheckStatusPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next GetReadinessCheckStatus page.
func (p *GetReadinessCheckStatusPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*GetReadinessCheckStatusOutput, error) {
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

	result, err := p.client.GetReadinessCheckStatus(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opGetReadinessCheckStatus(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetReadinessCheckStatus",
	}
}
