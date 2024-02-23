// Code generated by smithy-go-codegen DO NOT EDIT.

package opensearch

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/opensearch/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// A list of maintenance actions for the domain.
func (c *Client) ListDomainMaintenances(ctx context.Context, params *ListDomainMaintenancesInput, optFns ...func(*Options)) (*ListDomainMaintenancesOutput, error) {
	if params == nil {
		params = &ListDomainMaintenancesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListDomainMaintenances", params, optFns, c.addOperationListDomainMaintenancesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListDomainMaintenancesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Container for the parameters to the ListDomainMaintenances operation.
type ListDomainMaintenancesInput struct {

	// The name of the domain.
	//
	// This member is required.
	DomainName *string

	// The name of the action.
	Action types.MaintenanceType

	// An optional parameter that specifies the maximum number of results to return.
	// You can use nextToken to get the next page of results.
	MaxResults int32

	// If your initial ListDomainMaintenances operation returns a nextToken , include
	// the returned nextToken in subsequent ListDomainMaintenances operations, which
	// returns results in the next page.
	NextToken *string

	// The status of the action.
	Status types.MaintenanceStatus

	noSmithyDocumentSerde
}

// The result of a ListDomainMaintenances request that contains information about
// the requested actions.
type ListDomainMaintenancesOutput struct {

	// A list of the submitted maintenance actions.
	DomainMaintenances []types.DomainMaintenanceDetails

	// When nextToken is returned, there are more results available. The value of
	// nextToken is a unique pagination token for each page. Send the request again
	// using the returned token to retrieve the next page.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListDomainMaintenancesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListDomainMaintenances{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListDomainMaintenances{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListDomainMaintenances"); err != nil {
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
	if err = addOpListDomainMaintenancesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListDomainMaintenances(options.Region), middleware.Before); err != nil {
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

// ListDomainMaintenancesAPIClient is a client that implements the
// ListDomainMaintenances operation.
type ListDomainMaintenancesAPIClient interface {
	ListDomainMaintenances(context.Context, *ListDomainMaintenancesInput, ...func(*Options)) (*ListDomainMaintenancesOutput, error)
}

var _ ListDomainMaintenancesAPIClient = (*Client)(nil)

// ListDomainMaintenancesPaginatorOptions is the paginator options for
// ListDomainMaintenances
type ListDomainMaintenancesPaginatorOptions struct {
	// An optional parameter that specifies the maximum number of results to return.
	// You can use nextToken to get the next page of results.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListDomainMaintenancesPaginator is a paginator for ListDomainMaintenances
type ListDomainMaintenancesPaginator struct {
	options   ListDomainMaintenancesPaginatorOptions
	client    ListDomainMaintenancesAPIClient
	params    *ListDomainMaintenancesInput
	nextToken *string
	firstPage bool
}

// NewListDomainMaintenancesPaginator returns a new ListDomainMaintenancesPaginator
func NewListDomainMaintenancesPaginator(client ListDomainMaintenancesAPIClient, params *ListDomainMaintenancesInput, optFns ...func(*ListDomainMaintenancesPaginatorOptions)) *ListDomainMaintenancesPaginator {
	if params == nil {
		params = &ListDomainMaintenancesInput{}
	}

	options := ListDomainMaintenancesPaginatorOptions{}
	if params.MaxResults != 0 {
		options.Limit = params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListDomainMaintenancesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListDomainMaintenancesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListDomainMaintenances page.
func (p *ListDomainMaintenancesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListDomainMaintenancesOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	params.MaxResults = p.options.Limit

	result, err := p.client.ListDomainMaintenances(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListDomainMaintenances(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListDomainMaintenances",
	}
}
