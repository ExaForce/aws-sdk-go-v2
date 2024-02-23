// Code generated by smithy-go-codegen DO NOT EDIT.

package medialive

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/medialive/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// List input devices that are currently being transferred. List input devices
// that you are transferring from your AWS account or input devices that another
// AWS account is transferring to you.
func (c *Client) ListInputDeviceTransfers(ctx context.Context, params *ListInputDeviceTransfersInput, optFns ...func(*Options)) (*ListInputDeviceTransfersOutput, error) {
	if params == nil {
		params = &ListInputDeviceTransfersInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListInputDeviceTransfers", params, optFns, c.addOperationListInputDeviceTransfersMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListInputDeviceTransfersOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Placeholder documentation for ListInputDeviceTransfersRequest
type ListInputDeviceTransfersInput struct {

	// Placeholder documentation for __string
	//
	// This member is required.
	TransferType *string

	// Placeholder documentation for MaxResults
	MaxResults *int32

	// Placeholder documentation for __string
	NextToken *string

	noSmithyDocumentSerde
}

// Placeholder documentation for ListInputDeviceTransfersResponse
type ListInputDeviceTransfersOutput struct {

	// The list of devices that you are transferring or are being transferred to you.
	InputDeviceTransfers []types.TransferringInputDeviceSummary

	// A token to get additional list results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListInputDeviceTransfersMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListInputDeviceTransfers{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListInputDeviceTransfers{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListInputDeviceTransfers"); err != nil {
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
	if err = addOpListInputDeviceTransfersValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListInputDeviceTransfers(options.Region), middleware.Before); err != nil {
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

// ListInputDeviceTransfersAPIClient is a client that implements the
// ListInputDeviceTransfers operation.
type ListInputDeviceTransfersAPIClient interface {
	ListInputDeviceTransfers(context.Context, *ListInputDeviceTransfersInput, ...func(*Options)) (*ListInputDeviceTransfersOutput, error)
}

var _ ListInputDeviceTransfersAPIClient = (*Client)(nil)

// ListInputDeviceTransfersPaginatorOptions is the paginator options for
// ListInputDeviceTransfers
type ListInputDeviceTransfersPaginatorOptions struct {
	// Placeholder documentation for MaxResults
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListInputDeviceTransfersPaginator is a paginator for ListInputDeviceTransfers
type ListInputDeviceTransfersPaginator struct {
	options   ListInputDeviceTransfersPaginatorOptions
	client    ListInputDeviceTransfersAPIClient
	params    *ListInputDeviceTransfersInput
	nextToken *string
	firstPage bool
}

// NewListInputDeviceTransfersPaginator returns a new
// ListInputDeviceTransfersPaginator
func NewListInputDeviceTransfersPaginator(client ListInputDeviceTransfersAPIClient, params *ListInputDeviceTransfersInput, optFns ...func(*ListInputDeviceTransfersPaginatorOptions)) *ListInputDeviceTransfersPaginator {
	if params == nil {
		params = &ListInputDeviceTransfersInput{}
	}

	options := ListInputDeviceTransfersPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListInputDeviceTransfersPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListInputDeviceTransfersPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListInputDeviceTransfers page.
func (p *ListInputDeviceTransfersPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListInputDeviceTransfersOutput, error) {
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

	result, err := p.client.ListInputDeviceTransfers(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListInputDeviceTransfers(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListInputDeviceTransfers",
	}
}
