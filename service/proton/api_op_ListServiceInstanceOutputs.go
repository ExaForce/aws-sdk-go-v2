// Code generated by smithy-go-codegen DO NOT EDIT.

package proton

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/proton/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Get a list service of instance Infrastructure as Code (IaC) outputs.
func (c *Client) ListServiceInstanceOutputs(ctx context.Context, params *ListServiceInstanceOutputsInput, optFns ...func(*Options)) (*ListServiceInstanceOutputsOutput, error) {
	if params == nil {
		params = &ListServiceInstanceOutputsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListServiceInstanceOutputs", params, optFns, c.addOperationListServiceInstanceOutputsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListServiceInstanceOutputsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListServiceInstanceOutputsInput struct {

	// The name of the service instance whose outputs you want.
	//
	// This member is required.
	ServiceInstanceName *string

	// The name of the service that serviceInstanceName is associated to.
	//
	// This member is required.
	ServiceName *string

	// The ID of the deployment whose outputs you want.
	DeploymentId *string

	// A token that indicates the location of the next output in the array of outputs,
	// after the list of outputs that was previously requested.
	NextToken *string

	noSmithyDocumentSerde
}

type ListServiceInstanceOutputsOutput struct {

	// An array of service instance Infrastructure as Code (IaC) outputs.
	//
	// This member is required.
	Outputs []types.Output

	// A token that indicates the location of the next output in the array of outputs,
	// after the current requested list of outputs.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListServiceInstanceOutputsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpListServiceInstanceOutputs{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpListServiceInstanceOutputs{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListServiceInstanceOutputs"); err != nil {
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
	if err = addOpListServiceInstanceOutputsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListServiceInstanceOutputs(options.Region), middleware.Before); err != nil {
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

// ListServiceInstanceOutputsAPIClient is a client that implements the
// ListServiceInstanceOutputs operation.
type ListServiceInstanceOutputsAPIClient interface {
	ListServiceInstanceOutputs(context.Context, *ListServiceInstanceOutputsInput, ...func(*Options)) (*ListServiceInstanceOutputsOutput, error)
}

var _ ListServiceInstanceOutputsAPIClient = (*Client)(nil)

// ListServiceInstanceOutputsPaginatorOptions is the paginator options for
// ListServiceInstanceOutputs
type ListServiceInstanceOutputsPaginatorOptions struct {
	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListServiceInstanceOutputsPaginator is a paginator for
// ListServiceInstanceOutputs
type ListServiceInstanceOutputsPaginator struct {
	options   ListServiceInstanceOutputsPaginatorOptions
	client    ListServiceInstanceOutputsAPIClient
	params    *ListServiceInstanceOutputsInput
	nextToken *string
	firstPage bool
}

// NewListServiceInstanceOutputsPaginator returns a new
// ListServiceInstanceOutputsPaginator
func NewListServiceInstanceOutputsPaginator(client ListServiceInstanceOutputsAPIClient, params *ListServiceInstanceOutputsInput, optFns ...func(*ListServiceInstanceOutputsPaginatorOptions)) *ListServiceInstanceOutputsPaginator {
	if params == nil {
		params = &ListServiceInstanceOutputsInput{}
	}

	options := ListServiceInstanceOutputsPaginatorOptions{}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListServiceInstanceOutputsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListServiceInstanceOutputsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListServiceInstanceOutputs page.
func (p *ListServiceInstanceOutputsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListServiceInstanceOutputsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	result, err := p.client.ListServiceInstanceOutputs(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListServiceInstanceOutputs(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListServiceInstanceOutputs",
	}
}
