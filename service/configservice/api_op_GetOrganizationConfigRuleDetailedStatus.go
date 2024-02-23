// Code generated by smithy-go-codegen DO NOT EDIT.

package configservice

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/configservice/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns detailed status for each member account within an organization for a
// given organization Config rule.
func (c *Client) GetOrganizationConfigRuleDetailedStatus(ctx context.Context, params *GetOrganizationConfigRuleDetailedStatusInput, optFns ...func(*Options)) (*GetOrganizationConfigRuleDetailedStatusOutput, error) {
	if params == nil {
		params = &GetOrganizationConfigRuleDetailedStatusInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetOrganizationConfigRuleDetailedStatus", params, optFns, c.addOperationGetOrganizationConfigRuleDetailedStatusMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetOrganizationConfigRuleDetailedStatusOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetOrganizationConfigRuleDetailedStatusInput struct {

	// The name of your organization Config rule for which you want status details for
	// member accounts.
	//
	// This member is required.
	OrganizationConfigRuleName *string

	// A StatusDetailFilters object.
	Filters *types.StatusDetailFilters

	// The maximum number of OrganizationConfigRuleDetailedStatus returned on each
	// page. If you do not specify a number, Config uses the default. The default is
	// 100.
	Limit int32

	// The nextToken string returned on a previous page that you use to get the next
	// page of results in a paginated response.
	NextToken *string

	noSmithyDocumentSerde
}

type GetOrganizationConfigRuleDetailedStatusOutput struct {

	// The nextToken string returned on a previous page that you use to get the next
	// page of results in a paginated response.
	NextToken *string

	// A list of MemberAccountStatus objects.
	OrganizationConfigRuleDetailedStatus []types.MemberAccountStatus

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetOrganizationConfigRuleDetailedStatusMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetOrganizationConfigRuleDetailedStatus{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetOrganizationConfigRuleDetailedStatus{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetOrganizationConfigRuleDetailedStatus"); err != nil {
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
	if err = addOpGetOrganizationConfigRuleDetailedStatusValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetOrganizationConfigRuleDetailedStatus(options.Region), middleware.Before); err != nil {
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

// GetOrganizationConfigRuleDetailedStatusAPIClient is a client that implements
// the GetOrganizationConfigRuleDetailedStatus operation.
type GetOrganizationConfigRuleDetailedStatusAPIClient interface {
	GetOrganizationConfigRuleDetailedStatus(context.Context, *GetOrganizationConfigRuleDetailedStatusInput, ...func(*Options)) (*GetOrganizationConfigRuleDetailedStatusOutput, error)
}

var _ GetOrganizationConfigRuleDetailedStatusAPIClient = (*Client)(nil)

// GetOrganizationConfigRuleDetailedStatusPaginatorOptions is the paginator
// options for GetOrganizationConfigRuleDetailedStatus
type GetOrganizationConfigRuleDetailedStatusPaginatorOptions struct {
	// The maximum number of OrganizationConfigRuleDetailedStatus returned on each
	// page. If you do not specify a number, Config uses the default. The default is
	// 100.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// GetOrganizationConfigRuleDetailedStatusPaginator is a paginator for
// GetOrganizationConfigRuleDetailedStatus
type GetOrganizationConfigRuleDetailedStatusPaginator struct {
	options   GetOrganizationConfigRuleDetailedStatusPaginatorOptions
	client    GetOrganizationConfigRuleDetailedStatusAPIClient
	params    *GetOrganizationConfigRuleDetailedStatusInput
	nextToken *string
	firstPage bool
}

// NewGetOrganizationConfigRuleDetailedStatusPaginator returns a new
// GetOrganizationConfigRuleDetailedStatusPaginator
func NewGetOrganizationConfigRuleDetailedStatusPaginator(client GetOrganizationConfigRuleDetailedStatusAPIClient, params *GetOrganizationConfigRuleDetailedStatusInput, optFns ...func(*GetOrganizationConfigRuleDetailedStatusPaginatorOptions)) *GetOrganizationConfigRuleDetailedStatusPaginator {
	if params == nil {
		params = &GetOrganizationConfigRuleDetailedStatusInput{}
	}

	options := GetOrganizationConfigRuleDetailedStatusPaginatorOptions{}
	if params.Limit != 0 {
		options.Limit = params.Limit
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &GetOrganizationConfigRuleDetailedStatusPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *GetOrganizationConfigRuleDetailedStatusPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next GetOrganizationConfigRuleDetailedStatus page.
func (p *GetOrganizationConfigRuleDetailedStatusPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*GetOrganizationConfigRuleDetailedStatusOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	params.Limit = p.options.Limit

	result, err := p.client.GetOrganizationConfigRuleDetailedStatus(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opGetOrganizationConfigRuleDetailedStatus(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetOrganizationConfigRuleDetailedStatus",
	}
}
