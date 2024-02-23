// Code generated by smithy-go-codegen DO NOT EDIT.

package ecs

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/ecs/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists the account settings for a specified principal.
func (c *Client) ListAccountSettings(ctx context.Context, params *ListAccountSettingsInput, optFns ...func(*Options)) (*ListAccountSettingsOutput, error) {
	if params == nil {
		params = &ListAccountSettingsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListAccountSettings", params, optFns, c.addOperationListAccountSettingsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListAccountSettingsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListAccountSettingsInput struct {

	// Determines whether to return the effective settings. If true , the account
	// settings for the root user or the default setting for the principalArn are
	// returned. If false , the account settings for the principalArn are returned if
	// they're set. Otherwise, no account settings are returned.
	EffectiveSettings bool

	// The maximum number of account setting results returned by ListAccountSettings
	// in paginated output. When this parameter is used, ListAccountSettings only
	// returns maxResults results in a single page along with a nextToken response
	// element. The remaining results of the initial request can be seen by sending
	// another ListAccountSettings request with the returned nextToken value. This
	// value can be between 1 and 10. If this parameter isn't used, then
	// ListAccountSettings returns up to 10 results and a nextToken value if
	// applicable.
	MaxResults int32

	// The name of the account setting you want to list the settings for.
	Name types.SettingName

	// The nextToken value returned from a ListAccountSettings request indicating that
	// more results are available to fulfill the request and further calls will be
	// needed. If maxResults was provided, it's possible the number of results to be
	// fewer than maxResults . This token should be treated as an opaque identifier
	// that is only used to retrieve the next items in a list and not for other
	// programmatic purposes.
	NextToken *string

	// The ARN of the principal, which can be a user, role, or the root user. If this
	// field is omitted, the account settings are listed only for the authenticated
	// user. Federated users assume the account setting of the root user and can't have
	// explicit account settings set for them.
	PrincipalArn *string

	// The value of the account settings to filter results with. You must also specify
	// an account setting name to use this parameter.
	Value *string

	noSmithyDocumentSerde
}

type ListAccountSettingsOutput struct {

	// The nextToken value to include in a future ListAccountSettings request. When
	// the results of a ListAccountSettings request exceed maxResults , this value can
	// be used to retrieve the next page of results. This value is null when there are
	// no more results to return.
	NextToken *string

	// The account settings for the resource.
	Settings []types.Setting

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListAccountSettingsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListAccountSettings{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListAccountSettings{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListAccountSettings"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListAccountSettings(options.Region), middleware.Before); err != nil {
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

// ListAccountSettingsAPIClient is a client that implements the
// ListAccountSettings operation.
type ListAccountSettingsAPIClient interface {
	ListAccountSettings(context.Context, *ListAccountSettingsInput, ...func(*Options)) (*ListAccountSettingsOutput, error)
}

var _ ListAccountSettingsAPIClient = (*Client)(nil)

// ListAccountSettingsPaginatorOptions is the paginator options for
// ListAccountSettings
type ListAccountSettingsPaginatorOptions struct {
	// The maximum number of account setting results returned by ListAccountSettings
	// in paginated output. When this parameter is used, ListAccountSettings only
	// returns maxResults results in a single page along with a nextToken response
	// element. The remaining results of the initial request can be seen by sending
	// another ListAccountSettings request with the returned nextToken value. This
	// value can be between 1 and 10. If this parameter isn't used, then
	// ListAccountSettings returns up to 10 results and a nextToken value if
	// applicable.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListAccountSettingsPaginator is a paginator for ListAccountSettings
type ListAccountSettingsPaginator struct {
	options   ListAccountSettingsPaginatorOptions
	client    ListAccountSettingsAPIClient
	params    *ListAccountSettingsInput
	nextToken *string
	firstPage bool
}

// NewListAccountSettingsPaginator returns a new ListAccountSettingsPaginator
func NewListAccountSettingsPaginator(client ListAccountSettingsAPIClient, params *ListAccountSettingsInput, optFns ...func(*ListAccountSettingsPaginatorOptions)) *ListAccountSettingsPaginator {
	if params == nil {
		params = &ListAccountSettingsInput{}
	}

	options := ListAccountSettingsPaginatorOptions{}
	if params.MaxResults != 0 {
		options.Limit = params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListAccountSettingsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListAccountSettingsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListAccountSettings page.
func (p *ListAccountSettingsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListAccountSettingsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	params.MaxResults = p.options.Limit

	result, err := p.client.ListAccountSettings(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListAccountSettings(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListAccountSettings",
	}
}
