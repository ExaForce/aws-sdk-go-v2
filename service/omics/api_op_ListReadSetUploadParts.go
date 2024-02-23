// Code generated by smithy-go-codegen DO NOT EDIT.

package omics

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/omics/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// This operation will list all parts in a requested multipart upload for a
// sequence store.
func (c *Client) ListReadSetUploadParts(ctx context.Context, params *ListReadSetUploadPartsInput, optFns ...func(*Options)) (*ListReadSetUploadPartsOutput, error) {
	if params == nil {
		params = &ListReadSetUploadPartsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListReadSetUploadParts", params, optFns, c.addOperationListReadSetUploadPartsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListReadSetUploadPartsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListReadSetUploadPartsInput struct {

	// The source file for the upload part.
	//
	// This member is required.
	PartSource types.ReadSetPartSource

	// The Sequence Store ID used for the multipart uploads.
	//
	// This member is required.
	SequenceStoreId *string

	// The ID for the initiated multipart upload.
	//
	// This member is required.
	UploadId *string

	// Attributes used to filter for a specific subset of read set part uploads.
	Filter *types.ReadSetUploadPartListFilter

	// The maximum number of read set upload parts returned in a page.
	MaxResults *int32

	// Next token returned in the response of a previous ListReadSetUploadPartsRequest
	// call. Used to get the next page of results.
	NextToken *string

	noSmithyDocumentSerde
}

type ListReadSetUploadPartsOutput struct {

	// Next token returned in the response of a previous ListReadSetUploadParts call.
	// Used to get the next page of results.
	NextToken *string

	// An array of upload parts.
	Parts []types.ReadSetUploadPartListItem

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListReadSetUploadPartsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListReadSetUploadParts{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListReadSetUploadParts{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListReadSetUploadParts"); err != nil {
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
	if err = addEndpointPrefix_opListReadSetUploadPartsMiddleware(stack); err != nil {
		return err
	}
	if err = addOpListReadSetUploadPartsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListReadSetUploadParts(options.Region), middleware.Before); err != nil {
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

type endpointPrefix_opListReadSetUploadPartsMiddleware struct {
}

func (*endpointPrefix_opListReadSetUploadPartsMiddleware) ID() string {
	return "EndpointHostPrefix"
}

func (m *endpointPrefix_opListReadSetUploadPartsMiddleware) HandleFinalize(ctx context.Context, in middleware.FinalizeInput, next middleware.FinalizeHandler) (
	out middleware.FinalizeOutput, metadata middleware.Metadata, err error,
) {
	if smithyhttp.GetHostnameImmutable(ctx) || smithyhttp.IsEndpointHostPrefixDisabled(ctx) {
		return next.HandleFinalize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	req.URL.Host = "control-storage-" + req.URL.Host

	return next.HandleFinalize(ctx, in)
}
func addEndpointPrefix_opListReadSetUploadPartsMiddleware(stack *middleware.Stack) error {
	return stack.Finalize.Insert(&endpointPrefix_opListReadSetUploadPartsMiddleware{}, "ResolveEndpointV2", middleware.After)
}

// ListReadSetUploadPartsAPIClient is a client that implements the
// ListReadSetUploadParts operation.
type ListReadSetUploadPartsAPIClient interface {
	ListReadSetUploadParts(context.Context, *ListReadSetUploadPartsInput, ...func(*Options)) (*ListReadSetUploadPartsOutput, error)
}

var _ ListReadSetUploadPartsAPIClient = (*Client)(nil)

// ListReadSetUploadPartsPaginatorOptions is the paginator options for
// ListReadSetUploadParts
type ListReadSetUploadPartsPaginatorOptions struct {
	// The maximum number of read set upload parts returned in a page.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListReadSetUploadPartsPaginator is a paginator for ListReadSetUploadParts
type ListReadSetUploadPartsPaginator struct {
	options   ListReadSetUploadPartsPaginatorOptions
	client    ListReadSetUploadPartsAPIClient
	params    *ListReadSetUploadPartsInput
	nextToken *string
	firstPage bool
}

// NewListReadSetUploadPartsPaginator returns a new ListReadSetUploadPartsPaginator
func NewListReadSetUploadPartsPaginator(client ListReadSetUploadPartsAPIClient, params *ListReadSetUploadPartsInput, optFns ...func(*ListReadSetUploadPartsPaginatorOptions)) *ListReadSetUploadPartsPaginator {
	if params == nil {
		params = &ListReadSetUploadPartsInput{}
	}

	options := ListReadSetUploadPartsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListReadSetUploadPartsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListReadSetUploadPartsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListReadSetUploadParts page.
func (p *ListReadSetUploadPartsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListReadSetUploadPartsOutput, error) {
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

	result, err := p.client.ListReadSetUploadParts(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListReadSetUploadParts(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListReadSetUploadParts",
	}
}
