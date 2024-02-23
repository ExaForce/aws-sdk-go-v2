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

// Lists multipart read set uploads and for in progress uploads. Once the upload
// is completed, a read set is created and the upload will no longer be returned in
// the respone.
func (c *Client) ListMultipartReadSetUploads(ctx context.Context, params *ListMultipartReadSetUploadsInput, optFns ...func(*Options)) (*ListMultipartReadSetUploadsOutput, error) {
	if params == nil {
		params = &ListMultipartReadSetUploadsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListMultipartReadSetUploads", params, optFns, c.addOperationListMultipartReadSetUploadsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListMultipartReadSetUploadsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListMultipartReadSetUploadsInput struct {

	// The Sequence Store ID used for the multipart uploads.
	//
	// This member is required.
	SequenceStoreId *string

	// The maximum number of multipart uploads returned in a page.
	MaxResults *int32

	// Next token returned in the response of a previous ListMultipartReadSetUploads
	// call. Used to get the next page of results.
	NextToken *string

	noSmithyDocumentSerde
}

type ListMultipartReadSetUploadsOutput struct {

	// Next token returned in the response of a previous ListMultipartReadSetUploads
	// call. Used to get the next page of results.
	NextToken *string

	// An array of multipart uploads.
	Uploads []types.MultipartReadSetUploadListItem

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListMultipartReadSetUploadsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListMultipartReadSetUploads{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListMultipartReadSetUploads{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListMultipartReadSetUploads"); err != nil {
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
	if err = addEndpointPrefix_opListMultipartReadSetUploadsMiddleware(stack); err != nil {
		return err
	}
	if err = addOpListMultipartReadSetUploadsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListMultipartReadSetUploads(options.Region), middleware.Before); err != nil {
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

type endpointPrefix_opListMultipartReadSetUploadsMiddleware struct {
}

func (*endpointPrefix_opListMultipartReadSetUploadsMiddleware) ID() string {
	return "EndpointHostPrefix"
}

func (m *endpointPrefix_opListMultipartReadSetUploadsMiddleware) HandleFinalize(ctx context.Context, in middleware.FinalizeInput, next middleware.FinalizeHandler) (
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
func addEndpointPrefix_opListMultipartReadSetUploadsMiddleware(stack *middleware.Stack) error {
	return stack.Finalize.Insert(&endpointPrefix_opListMultipartReadSetUploadsMiddleware{}, "ResolveEndpointV2", middleware.After)
}

// ListMultipartReadSetUploadsAPIClient is a client that implements the
// ListMultipartReadSetUploads operation.
type ListMultipartReadSetUploadsAPIClient interface {
	ListMultipartReadSetUploads(context.Context, *ListMultipartReadSetUploadsInput, ...func(*Options)) (*ListMultipartReadSetUploadsOutput, error)
}

var _ ListMultipartReadSetUploadsAPIClient = (*Client)(nil)

// ListMultipartReadSetUploadsPaginatorOptions is the paginator options for
// ListMultipartReadSetUploads
type ListMultipartReadSetUploadsPaginatorOptions struct {
	// The maximum number of multipart uploads returned in a page.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListMultipartReadSetUploadsPaginator is a paginator for
// ListMultipartReadSetUploads
type ListMultipartReadSetUploadsPaginator struct {
	options   ListMultipartReadSetUploadsPaginatorOptions
	client    ListMultipartReadSetUploadsAPIClient
	params    *ListMultipartReadSetUploadsInput
	nextToken *string
	firstPage bool
}

// NewListMultipartReadSetUploadsPaginator returns a new
// ListMultipartReadSetUploadsPaginator
func NewListMultipartReadSetUploadsPaginator(client ListMultipartReadSetUploadsAPIClient, params *ListMultipartReadSetUploadsInput, optFns ...func(*ListMultipartReadSetUploadsPaginatorOptions)) *ListMultipartReadSetUploadsPaginator {
	if params == nil {
		params = &ListMultipartReadSetUploadsInput{}
	}

	options := ListMultipartReadSetUploadsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListMultipartReadSetUploadsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListMultipartReadSetUploadsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListMultipartReadSetUploads page.
func (p *ListMultipartReadSetUploadsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListMultipartReadSetUploadsOutput, error) {
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

	result, err := p.client.ListMultipartReadSetUploads(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListMultipartReadSetUploads(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListMultipartReadSetUploads",
	}
}
