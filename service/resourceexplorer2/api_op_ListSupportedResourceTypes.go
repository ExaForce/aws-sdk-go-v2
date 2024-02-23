// Code generated by smithy-go-codegen DO NOT EDIT.

package resourceexplorer2

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/resourceexplorer2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves a list of all resource types currently supported by Amazon Web
// Services Resource Explorer.
func (c *Client) ListSupportedResourceTypes(ctx context.Context, params *ListSupportedResourceTypesInput, optFns ...func(*Options)) (*ListSupportedResourceTypesOutput, error) {
	if params == nil {
		params = &ListSupportedResourceTypesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListSupportedResourceTypes", params, optFns, c.addOperationListSupportedResourceTypesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListSupportedResourceTypesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListSupportedResourceTypesInput struct {

	// The maximum number of results that you want included on each page of the
	// response. If you do not include this parameter, it defaults to a value
	// appropriate to the operation. If additional items exist beyond those included in
	// the current response, the NextToken response element is present and has a value
	// (is not null). Include that value as the NextToken request parameter in the
	// next call to the operation to get the next part of the results. An API operation
	// can return fewer results than the maximum even when there are more results
	// available. You should check NextToken after every operation to ensure that you
	// receive all of the results.
	MaxResults *int32

	// The parameter for receiving additional results if you receive a NextToken
	// response in a previous request. A NextToken response indicates that more output
	// is available. Set this parameter to the value of the previous call's NextToken
	// response to indicate where the output should continue from. The pagination
	// tokens expire after 24 hours.
	NextToken *string

	noSmithyDocumentSerde
}

type ListSupportedResourceTypesOutput struct {

	// If present, indicates that more output is available than is included in the
	// current response. Use this value in the NextToken request parameter in a
	// subsequent call to the operation to get the next part of the output. You should
	// repeat this until the NextToken response element comes back as null . The
	// pagination tokens expire after 24 hours.
	NextToken *string

	// The list of resource types supported by Resource Explorer.
	ResourceTypes []types.SupportedResourceType

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListSupportedResourceTypesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListSupportedResourceTypes{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListSupportedResourceTypes{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListSupportedResourceTypes"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListSupportedResourceTypes(options.Region), middleware.Before); err != nil {
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

// ListSupportedResourceTypesAPIClient is a client that implements the
// ListSupportedResourceTypes operation.
type ListSupportedResourceTypesAPIClient interface {
	ListSupportedResourceTypes(context.Context, *ListSupportedResourceTypesInput, ...func(*Options)) (*ListSupportedResourceTypesOutput, error)
}

var _ ListSupportedResourceTypesAPIClient = (*Client)(nil)

// ListSupportedResourceTypesPaginatorOptions is the paginator options for
// ListSupportedResourceTypes
type ListSupportedResourceTypesPaginatorOptions struct {
	// The maximum number of results that you want included on each page of the
	// response. If you do not include this parameter, it defaults to a value
	// appropriate to the operation. If additional items exist beyond those included in
	// the current response, the NextToken response element is present and has a value
	// (is not null). Include that value as the NextToken request parameter in the
	// next call to the operation to get the next part of the results. An API operation
	// can return fewer results than the maximum even when there are more results
	// available. You should check NextToken after every operation to ensure that you
	// receive all of the results.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListSupportedResourceTypesPaginator is a paginator for
// ListSupportedResourceTypes
type ListSupportedResourceTypesPaginator struct {
	options   ListSupportedResourceTypesPaginatorOptions
	client    ListSupportedResourceTypesAPIClient
	params    *ListSupportedResourceTypesInput
	nextToken *string
	firstPage bool
}

// NewListSupportedResourceTypesPaginator returns a new
// ListSupportedResourceTypesPaginator
func NewListSupportedResourceTypesPaginator(client ListSupportedResourceTypesAPIClient, params *ListSupportedResourceTypesInput, optFns ...func(*ListSupportedResourceTypesPaginatorOptions)) *ListSupportedResourceTypesPaginator {
	if params == nil {
		params = &ListSupportedResourceTypesInput{}
	}

	options := ListSupportedResourceTypesPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListSupportedResourceTypesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListSupportedResourceTypesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListSupportedResourceTypes page.
func (p *ListSupportedResourceTypesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListSupportedResourceTypesOutput, error) {
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

	result, err := p.client.ListSupportedResourceTypes(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListSupportedResourceTypes(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListSupportedResourceTypes",
	}
}
