// Code generated by smithy-go-codegen DO NOT EDIT.

package clouddirectory

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/clouddirectory/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves all available parent paths for any object type such as node, leaf
// node, policy node, and index node objects. For more information about objects,
// see Directory Structure (https://docs.aws.amazon.com/clouddirectory/latest/developerguide/key_concepts_directorystructure.html)
// . Use this API to evaluate all parents for an object. The call returns all
// objects from the root of the directory up to the requested object. The API
// returns the number of paths based on user-defined MaxResults , in case there are
// multiple paths to the parent. The order of the paths and nodes returned is
// consistent among multiple API calls unless the objects are deleted or moved.
// Paths not leading to the directory root are ignored from the target object.
func (c *Client) ListObjectParentPaths(ctx context.Context, params *ListObjectParentPathsInput, optFns ...func(*Options)) (*ListObjectParentPathsOutput, error) {
	if params == nil {
		params = &ListObjectParentPathsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListObjectParentPaths", params, optFns, c.addOperationListObjectParentPathsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListObjectParentPathsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListObjectParentPathsInput struct {

	// The ARN of the directory to which the parent path applies.
	//
	// This member is required.
	DirectoryArn *string

	// The reference that identifies the object whose parent paths are listed.
	//
	// This member is required.
	ObjectReference *types.ObjectReference

	// The maximum number of items to be retrieved in a single call. This is an
	// approximate number.
	MaxResults *int32

	// The pagination token.
	NextToken *string

	noSmithyDocumentSerde
}

type ListObjectParentPathsOutput struct {

	// The pagination token.
	NextToken *string

	// Returns the path to the ObjectIdentifiers that are associated with the
	// directory.
	PathToObjectIdentifiersList []types.PathToObjectIdentifiers

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListObjectParentPathsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListObjectParentPaths{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListObjectParentPaths{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListObjectParentPaths"); err != nil {
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
	if err = addOpListObjectParentPathsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListObjectParentPaths(options.Region), middleware.Before); err != nil {
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

// ListObjectParentPathsAPIClient is a client that implements the
// ListObjectParentPaths operation.
type ListObjectParentPathsAPIClient interface {
	ListObjectParentPaths(context.Context, *ListObjectParentPathsInput, ...func(*Options)) (*ListObjectParentPathsOutput, error)
}

var _ ListObjectParentPathsAPIClient = (*Client)(nil)

// ListObjectParentPathsPaginatorOptions is the paginator options for
// ListObjectParentPaths
type ListObjectParentPathsPaginatorOptions struct {
	// The maximum number of items to be retrieved in a single call. This is an
	// approximate number.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListObjectParentPathsPaginator is a paginator for ListObjectParentPaths
type ListObjectParentPathsPaginator struct {
	options   ListObjectParentPathsPaginatorOptions
	client    ListObjectParentPathsAPIClient
	params    *ListObjectParentPathsInput
	nextToken *string
	firstPage bool
}

// NewListObjectParentPathsPaginator returns a new ListObjectParentPathsPaginator
func NewListObjectParentPathsPaginator(client ListObjectParentPathsAPIClient, params *ListObjectParentPathsInput, optFns ...func(*ListObjectParentPathsPaginatorOptions)) *ListObjectParentPathsPaginator {
	if params == nil {
		params = &ListObjectParentPathsInput{}
	}

	options := ListObjectParentPathsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListObjectParentPathsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListObjectParentPathsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListObjectParentPaths page.
func (p *ListObjectParentPathsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListObjectParentPathsOutput, error) {
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

	result, err := p.client.ListObjectParentPaths(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListObjectParentPaths(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListObjectParentPaths",
	}
}
