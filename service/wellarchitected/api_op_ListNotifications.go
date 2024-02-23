// Code generated by smithy-go-codegen DO NOT EDIT.

package wellarchitected

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/wellarchitected/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// List lens notifications.
func (c *Client) ListNotifications(ctx context.Context, params *ListNotificationsInput, optFns ...func(*Options)) (*ListNotificationsOutput, error) {
	if params == nil {
		params = &ListNotificationsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListNotifications", params, optFns, c.addOperationListNotificationsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListNotificationsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListNotificationsInput struct {

	// The maximum number of results to return for this request.
	MaxResults *int32

	// The token to use to retrieve the next set of results.
	NextToken *string

	// The ARN for the related resource for the notification. Only one of WorkloadID
	// or ResourceARN should be specified.
	ResourceArn *string

	// The ID assigned to the workload. This ID is unique within an Amazon Web
	// Services Region.
	WorkloadId *string

	noSmithyDocumentSerde
}

type ListNotificationsOutput struct {

	// The token to use to retrieve the next set of results.
	NextToken *string

	// List of lens notification summaries in a workload.
	NotificationSummaries []types.NotificationSummary

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListNotificationsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListNotifications{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListNotifications{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListNotifications"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListNotifications(options.Region), middleware.Before); err != nil {
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

// ListNotificationsAPIClient is a client that implements the ListNotifications
// operation.
type ListNotificationsAPIClient interface {
	ListNotifications(context.Context, *ListNotificationsInput, ...func(*Options)) (*ListNotificationsOutput, error)
}

var _ ListNotificationsAPIClient = (*Client)(nil)

// ListNotificationsPaginatorOptions is the paginator options for ListNotifications
type ListNotificationsPaginatorOptions struct {
	// The maximum number of results to return for this request.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListNotificationsPaginator is a paginator for ListNotifications
type ListNotificationsPaginator struct {
	options   ListNotificationsPaginatorOptions
	client    ListNotificationsAPIClient
	params    *ListNotificationsInput
	nextToken *string
	firstPage bool
}

// NewListNotificationsPaginator returns a new ListNotificationsPaginator
func NewListNotificationsPaginator(client ListNotificationsAPIClient, params *ListNotificationsInput, optFns ...func(*ListNotificationsPaginatorOptions)) *ListNotificationsPaginator {
	if params == nil {
		params = &ListNotificationsInput{}
	}

	options := ListNotificationsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListNotificationsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListNotificationsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListNotifications page.
func (p *ListNotificationsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListNotificationsOutput, error) {
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

	result, err := p.client.ListNotifications(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListNotifications(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListNotifications",
	}
}
