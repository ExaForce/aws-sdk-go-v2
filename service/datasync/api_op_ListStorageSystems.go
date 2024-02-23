// Code generated by smithy-go-codegen DO NOT EDIT.

package datasync

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/datasync/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists the on-premises storage systems that you're using with DataSync Discovery.
func (c *Client) ListStorageSystems(ctx context.Context, params *ListStorageSystemsInput, optFns ...func(*Options)) (*ListStorageSystemsOutput, error) {
	if params == nil {
		params = &ListStorageSystemsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListStorageSystems", params, optFns, c.addOperationListStorageSystemsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListStorageSystemsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListStorageSystemsInput struct {

	// Specifies how many results you want in the response.
	MaxResults *int32

	// Specifies an opaque string that indicates the position to begin the next list
	// of results in the response.
	NextToken *string

	noSmithyDocumentSerde
}

type ListStorageSystemsOutput struct {

	// The opaque string that indicates the position to begin the next list of results
	// in the response.
	NextToken *string

	// The Amazon Resource Names ARNs) of the on-premises storage systems that you're
	// using with DataSync Discovery.
	StorageSystems []types.StorageSystemListEntry

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListStorageSystemsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListStorageSystems{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListStorageSystems{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListStorageSystems"); err != nil {
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
	if err = addEndpointPrefix_opListStorageSystemsMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListStorageSystems(options.Region), middleware.Before); err != nil {
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

type endpointPrefix_opListStorageSystemsMiddleware struct {
}

func (*endpointPrefix_opListStorageSystemsMiddleware) ID() string {
	return "EndpointHostPrefix"
}

func (m *endpointPrefix_opListStorageSystemsMiddleware) HandleFinalize(ctx context.Context, in middleware.FinalizeInput, next middleware.FinalizeHandler) (
	out middleware.FinalizeOutput, metadata middleware.Metadata, err error,
) {
	if smithyhttp.GetHostnameImmutable(ctx) || smithyhttp.IsEndpointHostPrefixDisabled(ctx) {
		return next.HandleFinalize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	req.URL.Host = "discovery-" + req.URL.Host

	return next.HandleFinalize(ctx, in)
}
func addEndpointPrefix_opListStorageSystemsMiddleware(stack *middleware.Stack) error {
	return stack.Finalize.Insert(&endpointPrefix_opListStorageSystemsMiddleware{}, "ResolveEndpointV2", middleware.After)
}

// ListStorageSystemsAPIClient is a client that implements the ListStorageSystems
// operation.
type ListStorageSystemsAPIClient interface {
	ListStorageSystems(context.Context, *ListStorageSystemsInput, ...func(*Options)) (*ListStorageSystemsOutput, error)
}

var _ ListStorageSystemsAPIClient = (*Client)(nil)

// ListStorageSystemsPaginatorOptions is the paginator options for
// ListStorageSystems
type ListStorageSystemsPaginatorOptions struct {
	// Specifies how many results you want in the response.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListStorageSystemsPaginator is a paginator for ListStorageSystems
type ListStorageSystemsPaginator struct {
	options   ListStorageSystemsPaginatorOptions
	client    ListStorageSystemsAPIClient
	params    *ListStorageSystemsInput
	nextToken *string
	firstPage bool
}

// NewListStorageSystemsPaginator returns a new ListStorageSystemsPaginator
func NewListStorageSystemsPaginator(client ListStorageSystemsAPIClient, params *ListStorageSystemsInput, optFns ...func(*ListStorageSystemsPaginatorOptions)) *ListStorageSystemsPaginator {
	if params == nil {
		params = &ListStorageSystemsInput{}
	}

	options := ListStorageSystemsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListStorageSystemsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListStorageSystemsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListStorageSystems page.
func (p *ListStorageSystemsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListStorageSystemsOutput, error) {
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

	result, err := p.client.ListStorageSystems(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListStorageSystems(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListStorageSystems",
	}
}
