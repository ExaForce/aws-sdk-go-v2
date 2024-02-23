// Code generated by smithy-go-codegen DO NOT EDIT.

package fsx

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/fsx/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Describes one or more Amazon FSx for NetApp ONTAP storage virtual machines
// (SVMs).
func (c *Client) DescribeStorageVirtualMachines(ctx context.Context, params *DescribeStorageVirtualMachinesInput, optFns ...func(*Options)) (*DescribeStorageVirtualMachinesOutput, error) {
	if params == nil {
		params = &DescribeStorageVirtualMachinesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeStorageVirtualMachines", params, optFns, c.addOperationDescribeStorageVirtualMachinesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeStorageVirtualMachinesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeStorageVirtualMachinesInput struct {

	// Enter a filter name:value pair to view a select set of SVMs.
	Filters []types.StorageVirtualMachineFilter

	// The maximum number of resources to return in the response. This value must be
	// an integer greater than zero.
	MaxResults *int32

	// (Optional) Opaque pagination token returned from a previous operation (String).
	// If present, this token indicates from what point you can continue processing the
	// request, where the previous NextToken value left off.
	NextToken *string

	// Enter the ID of one or more SVMs that you want to view.
	StorageVirtualMachineIds []string

	noSmithyDocumentSerde
}

type DescribeStorageVirtualMachinesOutput struct {

	// (Optional) Opaque pagination token returned from a previous operation (String).
	// If present, this token indicates from what point you can continue processing the
	// request, where the previous NextToken value left off.
	NextToken *string

	// Returned after a successful DescribeStorageVirtualMachines operation,
	// describing each SVM.
	StorageVirtualMachines []types.StorageVirtualMachine

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeStorageVirtualMachinesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeStorageVirtualMachines{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeStorageVirtualMachines{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeStorageVirtualMachines"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeStorageVirtualMachines(options.Region), middleware.Before); err != nil {
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

// DescribeStorageVirtualMachinesAPIClient is a client that implements the
// DescribeStorageVirtualMachines operation.
type DescribeStorageVirtualMachinesAPIClient interface {
	DescribeStorageVirtualMachines(context.Context, *DescribeStorageVirtualMachinesInput, ...func(*Options)) (*DescribeStorageVirtualMachinesOutput, error)
}

var _ DescribeStorageVirtualMachinesAPIClient = (*Client)(nil)

// DescribeStorageVirtualMachinesPaginatorOptions is the paginator options for
// DescribeStorageVirtualMachines
type DescribeStorageVirtualMachinesPaginatorOptions struct {
	// The maximum number of resources to return in the response. This value must be
	// an integer greater than zero.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeStorageVirtualMachinesPaginator is a paginator for
// DescribeStorageVirtualMachines
type DescribeStorageVirtualMachinesPaginator struct {
	options   DescribeStorageVirtualMachinesPaginatorOptions
	client    DescribeStorageVirtualMachinesAPIClient
	params    *DescribeStorageVirtualMachinesInput
	nextToken *string
	firstPage bool
}

// NewDescribeStorageVirtualMachinesPaginator returns a new
// DescribeStorageVirtualMachinesPaginator
func NewDescribeStorageVirtualMachinesPaginator(client DescribeStorageVirtualMachinesAPIClient, params *DescribeStorageVirtualMachinesInput, optFns ...func(*DescribeStorageVirtualMachinesPaginatorOptions)) *DescribeStorageVirtualMachinesPaginator {
	if params == nil {
		params = &DescribeStorageVirtualMachinesInput{}
	}

	options := DescribeStorageVirtualMachinesPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeStorageVirtualMachinesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeStorageVirtualMachinesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeStorageVirtualMachines page.
func (p *DescribeStorageVirtualMachinesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeStorageVirtualMachinesOutput, error) {
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

	result, err := p.client.DescribeStorageVirtualMachines(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opDescribeStorageVirtualMachines(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeStorageVirtualMachines",
	}
}
