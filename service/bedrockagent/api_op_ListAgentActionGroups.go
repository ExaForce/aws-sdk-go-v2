// Code generated by smithy-go-codegen DO NOT EDIT.

package bedrockagent

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/bedrockagent/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists an Action Group for existing Amazon Bedrock Agent Version
func (c *Client) ListAgentActionGroups(ctx context.Context, params *ListAgentActionGroupsInput, optFns ...func(*Options)) (*ListAgentActionGroupsOutput, error) {
	if params == nil {
		params = &ListAgentActionGroupsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListAgentActionGroups", params, optFns, c.addOperationListAgentActionGroupsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListAgentActionGroupsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// List Action Groups Request
type ListAgentActionGroupsInput struct {

	// Id generated at the server side when an Agent is Listed
	//
	// This member is required.
	AgentId *string

	// Id generated at the server side when an Agent is Listed
	//
	// This member is required.
	AgentVersion *string

	// Max Results.
	MaxResults *int32

	// Opaque continuation token of previous paginated response.
	NextToken *string

	noSmithyDocumentSerde
}

// List Action Groups Response
type ListAgentActionGroupsOutput struct {

	// List of ActionGroup Summaries
	//
	// This member is required.
	ActionGroupSummaries []types.ActionGroupSummary

	// Opaque continuation token of previous paginated response.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListAgentActionGroupsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListAgentActionGroups{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListAgentActionGroups{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListAgentActionGroups"); err != nil {
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
	if err = addOpListAgentActionGroupsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListAgentActionGroups(options.Region), middleware.Before); err != nil {
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

// ListAgentActionGroupsAPIClient is a client that implements the
// ListAgentActionGroups operation.
type ListAgentActionGroupsAPIClient interface {
	ListAgentActionGroups(context.Context, *ListAgentActionGroupsInput, ...func(*Options)) (*ListAgentActionGroupsOutput, error)
}

var _ ListAgentActionGroupsAPIClient = (*Client)(nil)

// ListAgentActionGroupsPaginatorOptions is the paginator options for
// ListAgentActionGroups
type ListAgentActionGroupsPaginatorOptions struct {
	// Max Results.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListAgentActionGroupsPaginator is a paginator for ListAgentActionGroups
type ListAgentActionGroupsPaginator struct {
	options   ListAgentActionGroupsPaginatorOptions
	client    ListAgentActionGroupsAPIClient
	params    *ListAgentActionGroupsInput
	nextToken *string
	firstPage bool
}

// NewListAgentActionGroupsPaginator returns a new ListAgentActionGroupsPaginator
func NewListAgentActionGroupsPaginator(client ListAgentActionGroupsAPIClient, params *ListAgentActionGroupsInput, optFns ...func(*ListAgentActionGroupsPaginatorOptions)) *ListAgentActionGroupsPaginator {
	if params == nil {
		params = &ListAgentActionGroupsInput{}
	}

	options := ListAgentActionGroupsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListAgentActionGroupsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListAgentActionGroupsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListAgentActionGroups page.
func (p *ListAgentActionGroupsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListAgentActionGroupsOutput, error) {
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

	result, err := p.client.ListAgentActionGroups(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListAgentActionGroups(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListAgentActionGroups",
	}
}
