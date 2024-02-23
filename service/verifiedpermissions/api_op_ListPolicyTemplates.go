// Code generated by smithy-go-codegen DO NOT EDIT.

package verifiedpermissions

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/verifiedpermissions/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns a paginated list of all policy templates in the specified policy store.
func (c *Client) ListPolicyTemplates(ctx context.Context, params *ListPolicyTemplatesInput, optFns ...func(*Options)) (*ListPolicyTemplatesOutput, error) {
	if params == nil {
		params = &ListPolicyTemplatesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListPolicyTemplates", params, optFns, c.addOperationListPolicyTemplatesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListPolicyTemplatesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListPolicyTemplatesInput struct {

	// Specifies the ID of the policy store that contains the policy templates you
	// want to list.
	//
	// This member is required.
	PolicyStoreId *string

	// Specifies the total number of results that you want included in each response.
	// If additional items exist beyond the number you specify, the NextToken response
	// element is returned with a value (not null). Include the specified value as the
	// NextToken request parameter in the next call to the operation to get the next
	// set of results. Note that the service might return fewer results than the
	// maximum even when there are more results available. You should check NextToken
	// after every operation to ensure that you receive all of the results. If you do
	// not specify this parameter, the operation defaults to 10 policy templates per
	// response. You can specify a maximum of 50 policy templates per response.
	MaxResults *int32

	// Specifies that you want to receive the next page of results. Valid only if you
	// received a NextToken response in the previous request. If you did, it indicates
	// that more output is available. Set this parameter to the value provided by the
	// previous call's NextToken response to request the next page of results.
	NextToken *string

	noSmithyDocumentSerde
}

type ListPolicyTemplatesOutput struct {

	// The list of the policy templates in the specified policy store.
	//
	// This member is required.
	PolicyTemplates []types.PolicyTemplateItem

	// If present, this value indicates that more output is available than is included
	// in the current response. Use this value in the NextToken request parameter in a
	// subsequent call to the operation to get the next part of the output. You should
	// repeat this until the NextToken response element comes back as null . This
	// indicates that this is the last page of results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListPolicyTemplatesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpListPolicyTemplates{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpListPolicyTemplates{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListPolicyTemplates"); err != nil {
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
	if err = addOpListPolicyTemplatesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListPolicyTemplates(options.Region), middleware.Before); err != nil {
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

// ListPolicyTemplatesAPIClient is a client that implements the
// ListPolicyTemplates operation.
type ListPolicyTemplatesAPIClient interface {
	ListPolicyTemplates(context.Context, *ListPolicyTemplatesInput, ...func(*Options)) (*ListPolicyTemplatesOutput, error)
}

var _ ListPolicyTemplatesAPIClient = (*Client)(nil)

// ListPolicyTemplatesPaginatorOptions is the paginator options for
// ListPolicyTemplates
type ListPolicyTemplatesPaginatorOptions struct {
	// Specifies the total number of results that you want included in each response.
	// If additional items exist beyond the number you specify, the NextToken response
	// element is returned with a value (not null). Include the specified value as the
	// NextToken request parameter in the next call to the operation to get the next
	// set of results. Note that the service might return fewer results than the
	// maximum even when there are more results available. You should check NextToken
	// after every operation to ensure that you receive all of the results. If you do
	// not specify this parameter, the operation defaults to 10 policy templates per
	// response. You can specify a maximum of 50 policy templates per response.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListPolicyTemplatesPaginator is a paginator for ListPolicyTemplates
type ListPolicyTemplatesPaginator struct {
	options   ListPolicyTemplatesPaginatorOptions
	client    ListPolicyTemplatesAPIClient
	params    *ListPolicyTemplatesInput
	nextToken *string
	firstPage bool
}

// NewListPolicyTemplatesPaginator returns a new ListPolicyTemplatesPaginator
func NewListPolicyTemplatesPaginator(client ListPolicyTemplatesAPIClient, params *ListPolicyTemplatesInput, optFns ...func(*ListPolicyTemplatesPaginatorOptions)) *ListPolicyTemplatesPaginator {
	if params == nil {
		params = &ListPolicyTemplatesInput{}
	}

	options := ListPolicyTemplatesPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListPolicyTemplatesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListPolicyTemplatesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListPolicyTemplates page.
func (p *ListPolicyTemplatesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListPolicyTemplatesOutput, error) {
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

	result, err := p.client.ListPolicyTemplates(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListPolicyTemplates(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListPolicyTemplates",
	}
}
