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

// Returns a list of image build versions.
func (c *Client) ListImageBuildVersions(ctx context.Context, params *ListImageBuildVersionsInput, optFns ...func(*Options)) (*ListImageBuildVersionsOutput, error) {
	if params == nil {
		params = &ListImageBuildVersionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListImageBuildVersions", params, optFns, c.addOperationListImageBuildVersionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListImageBuildVersionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListImageBuildVersionsInput struct {

	// The Amazon Resource Name (ARN) of the image whose build versions you want to
	// retrieve.
	//
	// This member is required.
	ImageVersionArn *string

	// Use the following filters to streamline results:
	//   - name
	//   - osVersion
	//   - platform
	//   - type
	//   - version
	Filters []types.Filter

	// The maximum items to return in a request.
	MaxResults *int32

	// A token to specify where to start paginating. This is the nextToken from a
	// previously truncated response.
	NextToken *string

	noSmithyDocumentSerde
}

type ListImageBuildVersionsOutput struct {

	// The list of image build versions.
	ImageSummaryList []types.ImageSummary

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

func (c *Client) addOperationListImageBuildVersionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListImageBuildVersions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListImageBuildVersions{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListImageBuildVersions"); err != nil {
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
	if err = addOpListImageBuildVersionsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListImageBuildVersions(options.Region), middleware.Before); err != nil {
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

// ListImageBuildVersionsAPIClient is a client that implements the
// ListImageBuildVersions operation.
type ListImageBuildVersionsAPIClient interface {
	ListImageBuildVersions(context.Context, *ListImageBuildVersionsInput, ...func(*Options)) (*ListImageBuildVersionsOutput, error)
}

var _ ListImageBuildVersionsAPIClient = (*Client)(nil)

// ListImageBuildVersionsPaginatorOptions is the paginator options for
// ListImageBuildVersions
type ListImageBuildVersionsPaginatorOptions struct {
	// The maximum items to return in a request.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListImageBuildVersionsPaginator is a paginator for ListImageBuildVersions
type ListImageBuildVersionsPaginator struct {
	options   ListImageBuildVersionsPaginatorOptions
	client    ListImageBuildVersionsAPIClient
	params    *ListImageBuildVersionsInput
	nextToken *string
	firstPage bool
}

// NewListImageBuildVersionsPaginator returns a new ListImageBuildVersionsPaginator
func NewListImageBuildVersionsPaginator(client ListImageBuildVersionsAPIClient, params *ListImageBuildVersionsInput, optFns ...func(*ListImageBuildVersionsPaginatorOptions)) *ListImageBuildVersionsPaginator {
	if params == nil {
		params = &ListImageBuildVersionsInput{}
	}

	options := ListImageBuildVersionsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListImageBuildVersionsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListImageBuildVersionsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListImageBuildVersions page.
func (p *ListImageBuildVersionsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListImageBuildVersionsOutput, error) {
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

	result, err := p.client.ListImageBuildVersions(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListImageBuildVersions(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListImageBuildVersions",
	}
}
