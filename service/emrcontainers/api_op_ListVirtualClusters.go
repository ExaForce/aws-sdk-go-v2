// Code generated by smithy-go-codegen DO NOT EDIT.

package emrcontainers

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/emrcontainers/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Lists information about the specified virtual cluster. Virtual cluster is a
// managed entity on Amazon EMR on EKS. You can create, describe, list and delete
// virtual clusters. They do not consume any additional resource in your system. A
// single virtual cluster maps to a single Kubernetes namespace. Given this
// relationship, you can model virtual clusters the same way you model Kubernetes
// namespaces to meet your requirements.
func (c *Client) ListVirtualClusters(ctx context.Context, params *ListVirtualClustersInput, optFns ...func(*Options)) (*ListVirtualClustersOutput, error) {
	if params == nil {
		params = &ListVirtualClustersInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListVirtualClusters", params, optFns, c.addOperationListVirtualClustersMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListVirtualClustersOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListVirtualClustersInput struct {

	// The container provider ID of the virtual cluster.
	ContainerProviderId *string

	// The container provider type of the virtual cluster. Amazon EKS is the only
	// supported type as of now.
	ContainerProviderType types.ContainerProviderType

	// The date and time after which the virtual clusters are created.
	CreatedAfter *time.Time

	// The date and time before which the virtual clusters are created.
	CreatedBefore *time.Time

	// The maximum number of virtual clusters that can be listed.
	MaxResults *int32

	// The token for the next set of virtual clusters to return.
	NextToken *string

	// The states of the requested virtual clusters.
	States []types.VirtualClusterState

	noSmithyDocumentSerde
}

type ListVirtualClustersOutput struct {

	// This output displays the token for the next set of virtual clusters.
	NextToken *string

	// This output lists the specified virtual clusters.
	VirtualClusters []types.VirtualCluster

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListVirtualClustersMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListVirtualClusters{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListVirtualClusters{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListVirtualClusters"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListVirtualClusters(options.Region), middleware.Before); err != nil {
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

// ListVirtualClustersAPIClient is a client that implements the
// ListVirtualClusters operation.
type ListVirtualClustersAPIClient interface {
	ListVirtualClusters(context.Context, *ListVirtualClustersInput, ...func(*Options)) (*ListVirtualClustersOutput, error)
}

var _ ListVirtualClustersAPIClient = (*Client)(nil)

// ListVirtualClustersPaginatorOptions is the paginator options for
// ListVirtualClusters
type ListVirtualClustersPaginatorOptions struct {
	// The maximum number of virtual clusters that can be listed.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListVirtualClustersPaginator is a paginator for ListVirtualClusters
type ListVirtualClustersPaginator struct {
	options   ListVirtualClustersPaginatorOptions
	client    ListVirtualClustersAPIClient
	params    *ListVirtualClustersInput
	nextToken *string
	firstPage bool
}

// NewListVirtualClustersPaginator returns a new ListVirtualClustersPaginator
func NewListVirtualClustersPaginator(client ListVirtualClustersAPIClient, params *ListVirtualClustersInput, optFns ...func(*ListVirtualClustersPaginatorOptions)) *ListVirtualClustersPaginator {
	if params == nil {
		params = &ListVirtualClustersInput{}
	}

	options := ListVirtualClustersPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListVirtualClustersPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListVirtualClustersPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListVirtualClusters page.
func (p *ListVirtualClustersPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListVirtualClustersOutput, error) {
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

	result, err := p.client.ListVirtualClusters(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListVirtualClusters(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListVirtualClusters",
	}
}
