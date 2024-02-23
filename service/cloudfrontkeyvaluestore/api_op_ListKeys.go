// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudfrontkeyvaluestore

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/cloudfrontkeyvaluestore/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns a list of key value pairs.
func (c *Client) ListKeys(ctx context.Context, params *ListKeysInput, optFns ...func(*Options)) (*ListKeysOutput, error) {
	if params == nil {
		params = &ListKeysInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListKeys", params, optFns, c.addOperationListKeysMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListKeysOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListKeysInput struct {

	// The Amazon Resource Name (ARN) of the Key Value Store.
	//
	// This member is required.
	KvsARN *string

	// Maximum number of results that are returned per call. The default is 10 and
	// maximum allowed page is 50.
	MaxResults *int32

	// If nextToken is returned in the response, there are more results available.
	// Make the next call using the returned token to retrieve the next page.
	NextToken *string

	noSmithyDocumentSerde
}

func (in *ListKeysInput) bindEndpointParams(p *EndpointParameters) {
	p.KvsARN = in.KvsARN

}

type ListKeysOutput struct {

	// Key value pairs
	Items []types.ListKeysResponseListItem

	// If nextToken is returned in the response, there are more results available.
	// Make the next call using the returned token to retrieve the next page.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListKeysMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListKeys{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListKeys{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListKeys"); err != nil {
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
	if err = addOpListKeysValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListKeys(options.Region), middleware.Before); err != nil {
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

// ListKeysAPIClient is a client that implements the ListKeys operation.
type ListKeysAPIClient interface {
	ListKeys(context.Context, *ListKeysInput, ...func(*Options)) (*ListKeysOutput, error)
}

var _ ListKeysAPIClient = (*Client)(nil)

// ListKeysPaginatorOptions is the paginator options for ListKeys
type ListKeysPaginatorOptions struct {
	// Maximum number of results that are returned per call. The default is 10 and
	// maximum allowed page is 50.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListKeysPaginator is a paginator for ListKeys
type ListKeysPaginator struct {
	options   ListKeysPaginatorOptions
	client    ListKeysAPIClient
	params    *ListKeysInput
	nextToken *string
	firstPage bool
}

// NewListKeysPaginator returns a new ListKeysPaginator
func NewListKeysPaginator(client ListKeysAPIClient, params *ListKeysInput, optFns ...func(*ListKeysPaginatorOptions)) *ListKeysPaginator {
	if params == nil {
		params = &ListKeysInput{}
	}

	options := ListKeysPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListKeysPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListKeysPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListKeys page.
func (p *ListKeysPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListKeysOutput, error) {
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

	result, err := p.client.ListKeys(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListKeys(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListKeys",
	}
}
