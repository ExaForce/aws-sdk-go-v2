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

// Lists parent objects that are associated with a given object in pagination
// fashion.
func (c *Client) ListObjectParents(ctx context.Context, params *ListObjectParentsInput, optFns ...func(*Options)) (*ListObjectParentsOutput, error) {
	if params == nil {
		params = &ListObjectParentsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListObjectParents", params, optFns, c.addOperationListObjectParentsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListObjectParentsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListObjectParentsInput struct {

	// The Amazon Resource Name (ARN) that is associated with the Directory where the
	// object resides. For more information, see arns .
	//
	// This member is required.
	DirectoryArn *string

	// The reference that identifies the object for which parent objects are being
	// listed.
	//
	// This member is required.
	ObjectReference *types.ObjectReference

	// Represents the manner and timing in which the successful write or update of an
	// object is reflected in a subsequent read operation of that same object.
	ConsistencyLevel types.ConsistencyLevel

	// When set to True, returns all ListObjectParentsResponse$ParentLinks . There
	// could be multiple links between a parent-child pair.
	IncludeAllLinksToEachParent bool

	// The maximum number of items to be retrieved in a single call. This is an
	// approximate number.
	MaxResults *int32

	// The pagination token.
	NextToken *string

	noSmithyDocumentSerde
}

type ListObjectParentsOutput struct {

	// The pagination token.
	NextToken *string

	// Returns a list of parent reference and LinkName Tuples.
	ParentLinks []types.ObjectIdentifierAndLinkNameTuple

	// The parent structure, which is a map with key as the ObjectIdentifier and
	// LinkName as the value.
	Parents map[string]string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListObjectParentsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListObjectParents{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListObjectParents{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListObjectParents"); err != nil {
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
	if err = addOpListObjectParentsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListObjectParents(options.Region), middleware.Before); err != nil {
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

// ListObjectParentsAPIClient is a client that implements the ListObjectParents
// operation.
type ListObjectParentsAPIClient interface {
	ListObjectParents(context.Context, *ListObjectParentsInput, ...func(*Options)) (*ListObjectParentsOutput, error)
}

var _ ListObjectParentsAPIClient = (*Client)(nil)

// ListObjectParentsPaginatorOptions is the paginator options for ListObjectParents
type ListObjectParentsPaginatorOptions struct {
	// The maximum number of items to be retrieved in a single call. This is an
	// approximate number.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListObjectParentsPaginator is a paginator for ListObjectParents
type ListObjectParentsPaginator struct {
	options   ListObjectParentsPaginatorOptions
	client    ListObjectParentsAPIClient
	params    *ListObjectParentsInput
	nextToken *string
	firstPage bool
}

// NewListObjectParentsPaginator returns a new ListObjectParentsPaginator
func NewListObjectParentsPaginator(client ListObjectParentsAPIClient, params *ListObjectParentsInput, optFns ...func(*ListObjectParentsPaginatorOptions)) *ListObjectParentsPaginator {
	if params == nil {
		params = &ListObjectParentsInput{}
	}

	options := ListObjectParentsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListObjectParentsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListObjectParentsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListObjectParents page.
func (p *ListObjectParentsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListObjectParentsOutput, error) {
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

	result, err := p.client.ListObjectParents(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListObjectParents(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListObjectParents",
	}
}
