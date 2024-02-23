// Code generated by smithy-go-codegen DO NOT EDIT.

package resourcegroupstaggingapi

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/resourcegroupstaggingapi/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns a table that shows counts of resources that are noncompliant with their
// tag policies. For more information on tag policies, see Tag Policies (https://docs.aws.amazon.com/organizations/latest/userguide/orgs_manage_policies_tag-policies.html)
// in the Organizations User Guide. You can call this operation only from the
// organization's management account and from the us-east-1 Region. This operation
// supports pagination, where the response can be sent in multiple pages. You
// should check the PaginationToken response parameter to determine if there are
// additional results available to return. Repeat the query, passing the
// PaginationToken response parameter value as an input to the next request until
// you recieve a null value. A null value for PaginationToken indicates that there
// are no more results waiting to be returned.
func (c *Client) GetComplianceSummary(ctx context.Context, params *GetComplianceSummaryInput, optFns ...func(*Options)) (*GetComplianceSummaryOutput, error) {
	if params == nil {
		params = &GetComplianceSummaryInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetComplianceSummary", params, optFns, c.addOperationGetComplianceSummaryMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetComplianceSummaryOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetComplianceSummaryInput struct {

	// Specifies a list of attributes to group the counts of noncompliant resources
	// by. If supplied, the counts are sorted by those attributes.
	GroupBy []types.GroupByAttribute

	// Specifies the maximum number of results to be returned in each page. A query
	// can return fewer than this maximum, even if there are more results still to
	// return. You should always check the PaginationToken response value to see if
	// there are more results. You can specify a minimum of 1 and a maximum value of
	// 100.
	MaxResults *int32

	// Specifies a PaginationToken response value from a previous request to indicate
	// that you want the next page of results. Leave this parameter empty in your
	// initial request.
	PaginationToken *string

	// Specifies a list of Amazon Web Services Regions to limit the output to. If you
	// use this parameter, the count of returned noncompliant resources includes only
	// resources in the specified Regions.
	RegionFilters []string

	// Specifies that you want the response to include information for only resources
	// of the specified types. The format of each resource type is
	// service[:resourceType] . For example, specifying a resource type of ec2 returns
	// all Amazon EC2 resources (which includes EC2 instances). Specifying a resource
	// type of ec2:instance returns only EC2 instances. The string for each service
	// name and resource type is the same as that embedded in a resource's Amazon
	// Resource Name (ARN). Consult the Amazon Web Services General Reference (https://docs.aws.amazon.com/general/latest/gr/)
	// for the following:
	//   - For a list of service name strings, see Amazon Web Services Service
	//   Namespaces (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html#genref-aws-service-namespaces)
	//   .
	//   - For resource type strings, see Example ARNs (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html#arns-syntax)
	//   .
	//   - For more information about ARNs, see Amazon Resource Names (ARNs) and
	//   Amazon Web Services Service Namespaces (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html)
	//   .
	// You can specify multiple resource types by using a comma separated array. The
	// array can include up to 100 items. Note that the length constraint requirement
	// applies to each resource type filter.
	ResourceTypeFilters []string

	// Specifies that you want the response to include information for only resources
	// that have tags with the specified tag keys. If you use this parameter, the count
	// of returned noncompliant resources includes only resources that have the
	// specified tag keys.
	TagKeyFilters []string

	// Specifies target identifiers (usually, specific account IDs) to limit the
	// output by. If you use this parameter, the count of returned noncompliant
	// resources includes only resources with the specified target IDs.
	TargetIdFilters []string

	noSmithyDocumentSerde
}

type GetComplianceSummaryOutput struct {

	// A string that indicates that there is more data available than this response
	// contains. To receive the next part of the response, specify this response value
	// as the PaginationToken value in the request for the next page.
	PaginationToken *string

	// A table that shows counts of noncompliant resources.
	SummaryList []types.Summary

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetComplianceSummaryMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetComplianceSummary{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetComplianceSummary{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetComplianceSummary"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetComplianceSummary(options.Region), middleware.Before); err != nil {
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

// GetComplianceSummaryAPIClient is a client that implements the
// GetComplianceSummary operation.
type GetComplianceSummaryAPIClient interface {
	GetComplianceSummary(context.Context, *GetComplianceSummaryInput, ...func(*Options)) (*GetComplianceSummaryOutput, error)
}

var _ GetComplianceSummaryAPIClient = (*Client)(nil)

// GetComplianceSummaryPaginatorOptions is the paginator options for
// GetComplianceSummary
type GetComplianceSummaryPaginatorOptions struct {
	// Specifies the maximum number of results to be returned in each page. A query
	// can return fewer than this maximum, even if there are more results still to
	// return. You should always check the PaginationToken response value to see if
	// there are more results. You can specify a minimum of 1 and a maximum value of
	// 100.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// GetComplianceSummaryPaginator is a paginator for GetComplianceSummary
type GetComplianceSummaryPaginator struct {
	options   GetComplianceSummaryPaginatorOptions
	client    GetComplianceSummaryAPIClient
	params    *GetComplianceSummaryInput
	nextToken *string
	firstPage bool
}

// NewGetComplianceSummaryPaginator returns a new GetComplianceSummaryPaginator
func NewGetComplianceSummaryPaginator(client GetComplianceSummaryAPIClient, params *GetComplianceSummaryInput, optFns ...func(*GetComplianceSummaryPaginatorOptions)) *GetComplianceSummaryPaginator {
	if params == nil {
		params = &GetComplianceSummaryInput{}
	}

	options := GetComplianceSummaryPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &GetComplianceSummaryPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.PaginationToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *GetComplianceSummaryPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next GetComplianceSummary page.
func (p *GetComplianceSummaryPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*GetComplianceSummaryOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.PaginationToken = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.MaxResults = limit

	result, err := p.client.GetComplianceSummary(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.PaginationToken

	if p.options.StopOnDuplicateToken &&
		prevToken != nil &&
		p.nextToken != nil &&
		*prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

func newServiceMetadataMiddleware_opGetComplianceSummary(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetComplianceSummary",
	}
}
