// Code generated by smithy-go-codegen DO NOT EDIT.

package migrationhubstrategy

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/migrationhubstrategy/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves detailed information about a specified server.
func (c *Client) GetServerDetails(ctx context.Context, params *GetServerDetailsInput, optFns ...func(*Options)) (*GetServerDetailsOutput, error) {
	if params == nil {
		params = &GetServerDetailsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetServerDetails", params, optFns, c.addOperationGetServerDetailsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetServerDetailsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetServerDetailsInput struct {

	// The ID of the server.
	//
	// This member is required.
	ServerId *string

	// The maximum number of items to include in the response. The maximum value is
	// 100.
	MaxResults *int32

	// The token from a previous call that you use to retrieve the next set of
	// results. For example, if a previous call to this action returned 100 items, but
	// you set maxResults to 10. You'll receive a set of 10 results along with a
	// token. You then use the returned token to retrieve the next set of 10.
	NextToken *string

	noSmithyDocumentSerde
}

type GetServerDetailsOutput struct {

	// The associated application group the server belongs to, as defined in AWS
	// Application Discovery Service.
	AssociatedApplications []types.AssociatedApplication

	// The token you use to retrieve the next set of results, or null if there are no
	// more results.
	NextToken *string

	// Detailed information about the server.
	ServerDetail *types.ServerDetail

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetServerDetailsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetServerDetails{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetServerDetails{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetServerDetails"); err != nil {
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
	if err = addOpGetServerDetailsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetServerDetails(options.Region), middleware.Before); err != nil {
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

// GetServerDetailsAPIClient is a client that implements the GetServerDetails
// operation.
type GetServerDetailsAPIClient interface {
	GetServerDetails(context.Context, *GetServerDetailsInput, ...func(*Options)) (*GetServerDetailsOutput, error)
}

var _ GetServerDetailsAPIClient = (*Client)(nil)

// GetServerDetailsPaginatorOptions is the paginator options for GetServerDetails
type GetServerDetailsPaginatorOptions struct {
	// The maximum number of items to include in the response. The maximum value is
	// 100.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// GetServerDetailsPaginator is a paginator for GetServerDetails
type GetServerDetailsPaginator struct {
	options   GetServerDetailsPaginatorOptions
	client    GetServerDetailsAPIClient
	params    *GetServerDetailsInput
	nextToken *string
	firstPage bool
}

// NewGetServerDetailsPaginator returns a new GetServerDetailsPaginator
func NewGetServerDetailsPaginator(client GetServerDetailsAPIClient, params *GetServerDetailsInput, optFns ...func(*GetServerDetailsPaginatorOptions)) *GetServerDetailsPaginator {
	if params == nil {
		params = &GetServerDetailsInput{}
	}

	options := GetServerDetailsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &GetServerDetailsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *GetServerDetailsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next GetServerDetails page.
func (p *GetServerDetailsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*GetServerDetailsOutput, error) {
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

	result, err := p.client.GetServerDetails(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opGetServerDetails(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetServerDetails",
	}
}
