// Code generated by smithy-go-codegen DO NOT EDIT.

package supportapp

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/supportapp/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists the Slack workspace configurations for an Amazon Web Services account.
func (c *Client) ListSlackWorkspaceConfigurations(ctx context.Context, params *ListSlackWorkspaceConfigurationsInput, optFns ...func(*Options)) (*ListSlackWorkspaceConfigurationsOutput, error) {
	if params == nil {
		params = &ListSlackWorkspaceConfigurationsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListSlackWorkspaceConfigurations", params, optFns, c.addOperationListSlackWorkspaceConfigurationsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListSlackWorkspaceConfigurationsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListSlackWorkspaceConfigurationsInput struct {

	// If the results of a search are large, the API only returns a portion of the
	// results and includes a nextToken pagination token in the response. To retrieve
	// the next batch of results, reissue the search request and include the returned
	// token. When the API returns the last set of results, the response doesn't
	// include a pagination token value.
	NextToken *string

	noSmithyDocumentSerde
}

type ListSlackWorkspaceConfigurationsOutput struct {

	// The point where pagination should resume when the response returns only partial
	// results.
	NextToken *string

	// The configurations for a Slack workspace.
	SlackWorkspaceConfigurations []types.SlackWorkspaceConfiguration

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListSlackWorkspaceConfigurationsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListSlackWorkspaceConfigurations{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListSlackWorkspaceConfigurations{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListSlackWorkspaceConfigurations"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListSlackWorkspaceConfigurations(options.Region), middleware.Before); err != nil {
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

// ListSlackWorkspaceConfigurationsAPIClient is a client that implements the
// ListSlackWorkspaceConfigurations operation.
type ListSlackWorkspaceConfigurationsAPIClient interface {
	ListSlackWorkspaceConfigurations(context.Context, *ListSlackWorkspaceConfigurationsInput, ...func(*Options)) (*ListSlackWorkspaceConfigurationsOutput, error)
}

var _ ListSlackWorkspaceConfigurationsAPIClient = (*Client)(nil)

// ListSlackWorkspaceConfigurationsPaginatorOptions is the paginator options for
// ListSlackWorkspaceConfigurations
type ListSlackWorkspaceConfigurationsPaginatorOptions struct {
	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListSlackWorkspaceConfigurationsPaginator is a paginator for
// ListSlackWorkspaceConfigurations
type ListSlackWorkspaceConfigurationsPaginator struct {
	options   ListSlackWorkspaceConfigurationsPaginatorOptions
	client    ListSlackWorkspaceConfigurationsAPIClient
	params    *ListSlackWorkspaceConfigurationsInput
	nextToken *string
	firstPage bool
}

// NewListSlackWorkspaceConfigurationsPaginator returns a new
// ListSlackWorkspaceConfigurationsPaginator
func NewListSlackWorkspaceConfigurationsPaginator(client ListSlackWorkspaceConfigurationsAPIClient, params *ListSlackWorkspaceConfigurationsInput, optFns ...func(*ListSlackWorkspaceConfigurationsPaginatorOptions)) *ListSlackWorkspaceConfigurationsPaginator {
	if params == nil {
		params = &ListSlackWorkspaceConfigurationsInput{}
	}

	options := ListSlackWorkspaceConfigurationsPaginatorOptions{}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListSlackWorkspaceConfigurationsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListSlackWorkspaceConfigurationsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListSlackWorkspaceConfigurations page.
func (p *ListSlackWorkspaceConfigurationsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListSlackWorkspaceConfigurationsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	result, err := p.client.ListSlackWorkspaceConfigurations(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListSlackWorkspaceConfigurations(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListSlackWorkspaceConfigurations",
	}
}
