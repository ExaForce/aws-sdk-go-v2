// Code generated by smithy-go-codegen DO NOT EDIT.

package frauddetector

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Gets all the elements in the specified list.
func (c *Client) GetListElements(ctx context.Context, params *GetListElementsInput, optFns ...func(*Options)) (*GetListElementsOutput, error) {
	if params == nil {
		params = &GetListElementsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetListElements", params, optFns, c.addOperationGetListElementsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetListElementsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetListElementsInput struct {

	// The name of the list.
	//
	// This member is required.
	Name *string

	// The maximum number of objects to return for the request.
	MaxResults *int32

	// The next token for the subsequent request.
	NextToken *string

	noSmithyDocumentSerde
}

type GetListElementsOutput struct {

	// The list elements.
	Elements []string

	// The next page token.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetListElementsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetListElements{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetListElements{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetListElements"); err != nil {
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
	if err = addOpGetListElementsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetListElements(options.Region), middleware.Before); err != nil {
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

// GetListElementsAPIClient is a client that implements the GetListElements
// operation.
type GetListElementsAPIClient interface {
	GetListElements(context.Context, *GetListElementsInput, ...func(*Options)) (*GetListElementsOutput, error)
}

var _ GetListElementsAPIClient = (*Client)(nil)

// GetListElementsPaginatorOptions is the paginator options for GetListElements
type GetListElementsPaginatorOptions struct {
	// The maximum number of objects to return for the request.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// GetListElementsPaginator is a paginator for GetListElements
type GetListElementsPaginator struct {
	options   GetListElementsPaginatorOptions
	client    GetListElementsAPIClient
	params    *GetListElementsInput
	nextToken *string
	firstPage bool
}

// NewGetListElementsPaginator returns a new GetListElementsPaginator
func NewGetListElementsPaginator(client GetListElementsAPIClient, params *GetListElementsInput, optFns ...func(*GetListElementsPaginatorOptions)) *GetListElementsPaginator {
	if params == nil {
		params = &GetListElementsInput{}
	}

	options := GetListElementsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &GetListElementsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *GetListElementsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next GetListElements page.
func (p *GetListElementsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*GetListElementsOutput, error) {
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

	result, err := p.client.GetListElements(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opGetListElements(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetListElements",
	}
}
