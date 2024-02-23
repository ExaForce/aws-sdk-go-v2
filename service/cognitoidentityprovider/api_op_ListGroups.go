// Code generated by smithy-go-codegen DO NOT EDIT.

package cognitoidentityprovider

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists the groups associated with a user pool. Amazon Cognito evaluates Identity
// and Access Management (IAM) policies in requests for this API operation. For
// this operation, you must use IAM credentials to authorize requests, and you must
// grant yourself the corresponding IAM permission in a policy. Learn more
//   - Signing Amazon Web Services API Requests (https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_aws-signing.html)
//   - Using the Amazon Cognito user pools API and user pool endpoints (https://docs.aws.amazon.com/cognito/latest/developerguide/user-pools-API-operations.html)
func (c *Client) ListGroups(ctx context.Context, params *ListGroupsInput, optFns ...func(*Options)) (*ListGroupsOutput, error) {
	if params == nil {
		params = &ListGroupsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListGroups", params, optFns, c.addOperationListGroupsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListGroupsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListGroupsInput struct {

	// The user pool ID for the user pool.
	//
	// This member is required.
	UserPoolId *string

	// The limit of the request to list groups.
	Limit *int32

	// An identifier that was returned from the previous call to this operation, which
	// can be used to return the next set of items in the list.
	NextToken *string

	noSmithyDocumentSerde
}

type ListGroupsOutput struct {

	// The group objects for the groups.
	Groups []types.GroupType

	// An identifier that was returned from the previous call to this operation, which
	// can be used to return the next set of items in the list.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListGroupsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListGroups{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListGroups{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListGroups"); err != nil {
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
	if err = addOpListGroupsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListGroups(options.Region), middleware.Before); err != nil {
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

// ListGroupsAPIClient is a client that implements the ListGroups operation.
type ListGroupsAPIClient interface {
	ListGroups(context.Context, *ListGroupsInput, ...func(*Options)) (*ListGroupsOutput, error)
}

var _ ListGroupsAPIClient = (*Client)(nil)

// ListGroupsPaginatorOptions is the paginator options for ListGroups
type ListGroupsPaginatorOptions struct {
	// The limit of the request to list groups.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListGroupsPaginator is a paginator for ListGroups
type ListGroupsPaginator struct {
	options   ListGroupsPaginatorOptions
	client    ListGroupsAPIClient
	params    *ListGroupsInput
	nextToken *string
	firstPage bool
}

// NewListGroupsPaginator returns a new ListGroupsPaginator
func NewListGroupsPaginator(client ListGroupsAPIClient, params *ListGroupsInput, optFns ...func(*ListGroupsPaginatorOptions)) *ListGroupsPaginator {
	if params == nil {
		params = &ListGroupsInput{}
	}

	options := ListGroupsPaginatorOptions{}
	if params.Limit != nil {
		options.Limit = *params.Limit
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListGroupsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListGroupsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListGroups page.
func (p *ListGroupsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListGroupsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.Limit = limit

	result, err := p.client.ListGroups(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListGroups(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListGroups",
	}
}
