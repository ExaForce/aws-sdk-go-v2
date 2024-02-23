// Code generated by smithy-go-codegen DO NOT EDIT.

package networkmanager

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/networkmanager/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns a change set between the LIVE core network policy and a submitted
// policy.
func (c *Client) GetCoreNetworkChangeSet(ctx context.Context, params *GetCoreNetworkChangeSetInput, optFns ...func(*Options)) (*GetCoreNetworkChangeSetOutput, error) {
	if params == nil {
		params = &GetCoreNetworkChangeSetInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetCoreNetworkChangeSet", params, optFns, c.addOperationGetCoreNetworkChangeSetMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetCoreNetworkChangeSetOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetCoreNetworkChangeSetInput struct {

	// The ID of a core network.
	//
	// This member is required.
	CoreNetworkId *string

	// The ID of the policy version.
	//
	// This member is required.
	PolicyVersionId *int32

	// The maximum number of results to return.
	MaxResults *int32

	// The token for the next page of results.
	NextToken *string

	noSmithyDocumentSerde
}

type GetCoreNetworkChangeSetOutput struct {

	// Describes a core network changes.
	CoreNetworkChanges []types.CoreNetworkChange

	// The token for the next page of results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetCoreNetworkChangeSetMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetCoreNetworkChangeSet{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetCoreNetworkChangeSet{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetCoreNetworkChangeSet"); err != nil {
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
	if err = addOpGetCoreNetworkChangeSetValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetCoreNetworkChangeSet(options.Region), middleware.Before); err != nil {
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

// GetCoreNetworkChangeSetAPIClient is a client that implements the
// GetCoreNetworkChangeSet operation.
type GetCoreNetworkChangeSetAPIClient interface {
	GetCoreNetworkChangeSet(context.Context, *GetCoreNetworkChangeSetInput, ...func(*Options)) (*GetCoreNetworkChangeSetOutput, error)
}

var _ GetCoreNetworkChangeSetAPIClient = (*Client)(nil)

// GetCoreNetworkChangeSetPaginatorOptions is the paginator options for
// GetCoreNetworkChangeSet
type GetCoreNetworkChangeSetPaginatorOptions struct {
	// The maximum number of results to return.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// GetCoreNetworkChangeSetPaginator is a paginator for GetCoreNetworkChangeSet
type GetCoreNetworkChangeSetPaginator struct {
	options   GetCoreNetworkChangeSetPaginatorOptions
	client    GetCoreNetworkChangeSetAPIClient
	params    *GetCoreNetworkChangeSetInput
	nextToken *string
	firstPage bool
}

// NewGetCoreNetworkChangeSetPaginator returns a new
// GetCoreNetworkChangeSetPaginator
func NewGetCoreNetworkChangeSetPaginator(client GetCoreNetworkChangeSetAPIClient, params *GetCoreNetworkChangeSetInput, optFns ...func(*GetCoreNetworkChangeSetPaginatorOptions)) *GetCoreNetworkChangeSetPaginator {
	if params == nil {
		params = &GetCoreNetworkChangeSetInput{}
	}

	options := GetCoreNetworkChangeSetPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &GetCoreNetworkChangeSetPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *GetCoreNetworkChangeSetPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next GetCoreNetworkChangeSet page.
func (p *GetCoreNetworkChangeSetPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*GetCoreNetworkChangeSetOutput, error) {
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

	result, err := p.client.GetCoreNetworkChangeSet(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opGetCoreNetworkChangeSet(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetCoreNetworkChangeSet",
	}
}
