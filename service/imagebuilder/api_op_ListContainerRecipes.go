// Code generated by smithy-go-codegen DO NOT EDIT.

package imagebuilder

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/imagebuilder/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns a list of container recipes.
func (c *Client) ListContainerRecipes(ctx context.Context, params *ListContainerRecipesInput, optFns ...func(*Options)) (*ListContainerRecipesOutput, error) {
	if params == nil {
		params = &ListContainerRecipesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListContainerRecipes", params, optFns, c.addOperationListContainerRecipesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListContainerRecipesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListContainerRecipesInput struct {

	// Use the following filters to streamline results:
	//   - containerType
	//   - name
	//   - parentImage
	//   - platform
	Filters []types.Filter

	// The maximum items to return in a request.
	MaxResults *int32

	// A token to specify where to start paginating. This is the nextToken from a
	// previously truncated response.
	NextToken *string

	// Returns container recipes belonging to the specified owner, that have been
	// shared with you. You can omit this field to return container recipes belonging
	// to your account.
	Owner types.Ownership

	noSmithyDocumentSerde
}

type ListContainerRecipesOutput struct {

	// The list of container recipes returned for the request.
	ContainerRecipeSummaryList []types.ContainerRecipeSummary

	// The next token used for paginated responses. When this field isn't empty, there
	// are additional elements that the service hasn't included in this request. Use
	// this token with the next request to retrieve additional objects.
	NextToken *string

	// The request ID that uniquely identifies this request.
	RequestId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListContainerRecipesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListContainerRecipes{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListContainerRecipes{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListContainerRecipes"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListContainerRecipes(options.Region), middleware.Before); err != nil {
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

// ListContainerRecipesAPIClient is a client that implements the
// ListContainerRecipes operation.
type ListContainerRecipesAPIClient interface {
	ListContainerRecipes(context.Context, *ListContainerRecipesInput, ...func(*Options)) (*ListContainerRecipesOutput, error)
}

var _ ListContainerRecipesAPIClient = (*Client)(nil)

// ListContainerRecipesPaginatorOptions is the paginator options for
// ListContainerRecipes
type ListContainerRecipesPaginatorOptions struct {
	// The maximum items to return in a request.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListContainerRecipesPaginator is a paginator for ListContainerRecipes
type ListContainerRecipesPaginator struct {
	options   ListContainerRecipesPaginatorOptions
	client    ListContainerRecipesAPIClient
	params    *ListContainerRecipesInput
	nextToken *string
	firstPage bool
}

// NewListContainerRecipesPaginator returns a new ListContainerRecipesPaginator
func NewListContainerRecipesPaginator(client ListContainerRecipesAPIClient, params *ListContainerRecipesInput, optFns ...func(*ListContainerRecipesPaginatorOptions)) *ListContainerRecipesPaginator {
	if params == nil {
		params = &ListContainerRecipesInput{}
	}

	options := ListContainerRecipesPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListContainerRecipesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListContainerRecipesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListContainerRecipes page.
func (p *ListContainerRecipesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListContainerRecipesOutput, error) {
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

	result, err := p.client.ListContainerRecipes(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListContainerRecipes(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListContainerRecipes",
	}
}
