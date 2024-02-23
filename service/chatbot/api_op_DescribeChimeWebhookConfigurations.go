// Code generated by smithy-go-codegen DO NOT EDIT.

package chatbot

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/chatbot/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists Chime Webhook Configurations optionally filtered by ChatConfigurationArn
func (c *Client) DescribeChimeWebhookConfigurations(ctx context.Context, params *DescribeChimeWebhookConfigurationsInput, optFns ...func(*Options)) (*DescribeChimeWebhookConfigurationsOutput, error) {
	if params == nil {
		params = &DescribeChimeWebhookConfigurationsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeChimeWebhookConfigurations", params, optFns, c.addOperationDescribeChimeWebhookConfigurationsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeChimeWebhookConfigurationsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeChimeWebhookConfigurationsInput struct {

	// An optional ARN of a ChimeWebhookConfiguration to describe.
	ChatConfigurationArn *string

	// The maximum number of results to include in the response. If more results exist
	// than the specified MaxResults value, a token is included in the response so that
	// the remaining results can be retrieved.
	MaxResults *int32

	// An optional token returned from a prior request. Use this token for pagination
	// of results from this action. If this parameter is specified, the response
	// includes only results beyond the token, up to the value specified by MaxResults.
	NextToken *string

	noSmithyDocumentSerde
}

type DescribeChimeWebhookConfigurationsOutput struct {

	// An optional token returned from a prior request. Use this token for pagination
	// of results from this action. If this parameter is specified, the response
	// includes only results beyond the token, up to the value specified by MaxResults.
	NextToken *string

	// A list of Chime webhooks associated with the account.
	WebhookConfigurations []types.ChimeWebhookConfiguration

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeChimeWebhookConfigurationsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribeChimeWebhookConfigurations{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeChimeWebhookConfigurations{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeChimeWebhookConfigurations"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeChimeWebhookConfigurations(options.Region), middleware.Before); err != nil {
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

// DescribeChimeWebhookConfigurationsAPIClient is a client that implements the
// DescribeChimeWebhookConfigurations operation.
type DescribeChimeWebhookConfigurationsAPIClient interface {
	DescribeChimeWebhookConfigurations(context.Context, *DescribeChimeWebhookConfigurationsInput, ...func(*Options)) (*DescribeChimeWebhookConfigurationsOutput, error)
}

var _ DescribeChimeWebhookConfigurationsAPIClient = (*Client)(nil)

// DescribeChimeWebhookConfigurationsPaginatorOptions is the paginator options for
// DescribeChimeWebhookConfigurations
type DescribeChimeWebhookConfigurationsPaginatorOptions struct {
	// The maximum number of results to include in the response. If more results exist
	// than the specified MaxResults value, a token is included in the response so that
	// the remaining results can be retrieved.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeChimeWebhookConfigurationsPaginator is a paginator for
// DescribeChimeWebhookConfigurations
type DescribeChimeWebhookConfigurationsPaginator struct {
	options   DescribeChimeWebhookConfigurationsPaginatorOptions
	client    DescribeChimeWebhookConfigurationsAPIClient
	params    *DescribeChimeWebhookConfigurationsInput
	nextToken *string
	firstPage bool
}

// NewDescribeChimeWebhookConfigurationsPaginator returns a new
// DescribeChimeWebhookConfigurationsPaginator
func NewDescribeChimeWebhookConfigurationsPaginator(client DescribeChimeWebhookConfigurationsAPIClient, params *DescribeChimeWebhookConfigurationsInput, optFns ...func(*DescribeChimeWebhookConfigurationsPaginatorOptions)) *DescribeChimeWebhookConfigurationsPaginator {
	if params == nil {
		params = &DescribeChimeWebhookConfigurationsInput{}
	}

	options := DescribeChimeWebhookConfigurationsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeChimeWebhookConfigurationsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeChimeWebhookConfigurationsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeChimeWebhookConfigurations page.
func (p *DescribeChimeWebhookConfigurationsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeChimeWebhookConfigurationsOutput, error) {
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

	result, err := p.client.DescribeChimeWebhookConfigurations(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opDescribeChimeWebhookConfigurations(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeChimeWebhookConfigurations",
	}
}
