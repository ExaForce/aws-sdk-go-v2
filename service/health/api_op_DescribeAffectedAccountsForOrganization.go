// Code generated by smithy-go-codegen DO NOT EDIT.

package health

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/health/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns a list of accounts in the organization from Organizations that are
// affected by the provided event. For more information about the different types
// of Health events, see Event (https://docs.aws.amazon.com/health/latest/APIReference/API_Event.html)
// . Before you can call this operation, you must first enable Health to work with
// Organizations. To do this, call the EnableHealthServiceAccessForOrganization (https://docs.aws.amazon.com/health/latest/APIReference/API_EnableHealthServiceAccessForOrganization.html)
// operation from your organization's management account. This API operation uses
// pagination. Specify the nextToken parameter in the next request to return more
// results.
func (c *Client) DescribeAffectedAccountsForOrganization(ctx context.Context, params *DescribeAffectedAccountsForOrganizationInput, optFns ...func(*Options)) (*DescribeAffectedAccountsForOrganizationOutput, error) {
	if params == nil {
		params = &DescribeAffectedAccountsForOrganizationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeAffectedAccountsForOrganization", params, optFns, c.addOperationDescribeAffectedAccountsForOrganizationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeAffectedAccountsForOrganizationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeAffectedAccountsForOrganizationInput struct {

	// The unique identifier for the event. The event ARN has the
	// arn:aws:health:event-region::event/SERVICE/EVENT_TYPE_CODE/EVENT_TYPE_PLUS_ID
	// format. For example, an event ARN might look like the following:
	// arn:aws:health:us-east-1::event/EC2/EC2_INSTANCE_RETIREMENT_SCHEDULED/EC2_INSTANCE_RETIREMENT_SCHEDULED_ABC123-DEF456
	//
	// This member is required.
	EventArn *string

	// The maximum number of items to return in one batch, between 10 and 100,
	// inclusive.
	MaxResults *int32

	// If the results of a search are large, only a portion of the results are
	// returned, and a nextToken pagination token is returned in the response. To
	// retrieve the next batch of results, reissue the search request and include the
	// returned token. When all results have been returned, the response does not
	// contain a pagination token value.
	NextToken *string

	noSmithyDocumentSerde
}

type DescribeAffectedAccountsForOrganizationOutput struct {

	// A JSON set of elements of the affected accounts.
	AffectedAccounts []string

	// This parameter specifies if the Health event is a public Amazon Web Service
	// event or an account-specific event.
	//   - If the eventScopeCode value is PUBLIC , then the affectedAccounts value is
	//   always empty.
	//   - If the eventScopeCode value is ACCOUNT_SPECIFIC , then the affectedAccounts
	//   value lists the affected Amazon Web Services accounts in your organization. For
	//   example, if an event affects a service such as Amazon Elastic Compute Cloud and
	//   you have Amazon Web Services accounts that use that service, those account IDs
	//   appear in the response.
	//   - If the eventScopeCode value is NONE , then the eventArn that you specified
	//   in the request is invalid or doesn't exist.
	EventScopeCode types.EventScopeCode

	// If the results of a search are large, only a portion of the results are
	// returned, and a nextToken pagination token is returned in the response. To
	// retrieve the next batch of results, reissue the search request and include the
	// returned token. When all results have been returned, the response does not
	// contain a pagination token value.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeAffectedAccountsForOrganizationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeAffectedAccountsForOrganization{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeAffectedAccountsForOrganization{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeAffectedAccountsForOrganization"); err != nil {
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
	if err = addOpDescribeAffectedAccountsForOrganizationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeAffectedAccountsForOrganization(options.Region), middleware.Before); err != nil {
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

// DescribeAffectedAccountsForOrganizationAPIClient is a client that implements
// the DescribeAffectedAccountsForOrganization operation.
type DescribeAffectedAccountsForOrganizationAPIClient interface {
	DescribeAffectedAccountsForOrganization(context.Context, *DescribeAffectedAccountsForOrganizationInput, ...func(*Options)) (*DescribeAffectedAccountsForOrganizationOutput, error)
}

var _ DescribeAffectedAccountsForOrganizationAPIClient = (*Client)(nil)

// DescribeAffectedAccountsForOrganizationPaginatorOptions is the paginator
// options for DescribeAffectedAccountsForOrganization
type DescribeAffectedAccountsForOrganizationPaginatorOptions struct {
	// The maximum number of items to return in one batch, between 10 and 100,
	// inclusive.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeAffectedAccountsForOrganizationPaginator is a paginator for
// DescribeAffectedAccountsForOrganization
type DescribeAffectedAccountsForOrganizationPaginator struct {
	options   DescribeAffectedAccountsForOrganizationPaginatorOptions
	client    DescribeAffectedAccountsForOrganizationAPIClient
	params    *DescribeAffectedAccountsForOrganizationInput
	nextToken *string
	firstPage bool
}

// NewDescribeAffectedAccountsForOrganizationPaginator returns a new
// DescribeAffectedAccountsForOrganizationPaginator
func NewDescribeAffectedAccountsForOrganizationPaginator(client DescribeAffectedAccountsForOrganizationAPIClient, params *DescribeAffectedAccountsForOrganizationInput, optFns ...func(*DescribeAffectedAccountsForOrganizationPaginatorOptions)) *DescribeAffectedAccountsForOrganizationPaginator {
	if params == nil {
		params = &DescribeAffectedAccountsForOrganizationInput{}
	}

	options := DescribeAffectedAccountsForOrganizationPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeAffectedAccountsForOrganizationPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeAffectedAccountsForOrganizationPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeAffectedAccountsForOrganization page.
func (p *DescribeAffectedAccountsForOrganizationPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeAffectedAccountsForOrganizationOutput, error) {
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

	result, err := p.client.DescribeAffectedAccountsForOrganization(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opDescribeAffectedAccountsForOrganization(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeAffectedAccountsForOrganization",
	}
}
