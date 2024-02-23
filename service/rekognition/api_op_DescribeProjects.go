// Code generated by smithy-go-codegen DO NOT EDIT.

package rekognition

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/rekognition/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Gets information about your Rekognition projects. This operation requires
// permissions to perform the rekognition:DescribeProjects action.
func (c *Client) DescribeProjects(ctx context.Context, params *DescribeProjectsInput, optFns ...func(*Options)) (*DescribeProjectsOutput, error) {
	if params == nil {
		params = &DescribeProjectsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeProjects", params, optFns, c.addOperationDescribeProjectsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeProjectsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeProjectsInput struct {

	// Specifies the type of customization to filter projects by. If no value is
	// specified, CUSTOM_LABELS is used as a default.
	Features []types.CustomizationFeature

	// The maximum number of results to return per paginated call. The largest value
	// you can specify is 100. If you specify a value greater than 100, a
	// ValidationException error occurs. The default value is 100.
	MaxResults *int32

	// If the previous response was incomplete (because there is more results to
	// retrieve), Rekognition returns a pagination token in the response. You can use
	// this pagination token to retrieve the next set of results.
	NextToken *string

	// A list of the projects that you want Rekognition to describe. If you don't
	// specify a value, the response includes descriptions for all the projects in your
	// AWS account.
	ProjectNames []string

	noSmithyDocumentSerde
}

type DescribeProjectsOutput struct {

	// If the previous response was incomplete (because there is more results to
	// retrieve), Amazon Rekognition returns a pagination token in the response. You
	// can use this pagination token to retrieve the next set of results.
	NextToken *string

	// A list of project descriptions. The list is sorted by the date and time the
	// projects are created.
	ProjectDescriptions []types.ProjectDescription

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeProjectsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeProjects{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeProjects{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeProjects"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeProjects(options.Region), middleware.Before); err != nil {
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

// DescribeProjectsAPIClient is a client that implements the DescribeProjects
// operation.
type DescribeProjectsAPIClient interface {
	DescribeProjects(context.Context, *DescribeProjectsInput, ...func(*Options)) (*DescribeProjectsOutput, error)
}

var _ DescribeProjectsAPIClient = (*Client)(nil)

// DescribeProjectsPaginatorOptions is the paginator options for DescribeProjects
type DescribeProjectsPaginatorOptions struct {
	// The maximum number of results to return per paginated call. The largest value
	// you can specify is 100. If you specify a value greater than 100, a
	// ValidationException error occurs. The default value is 100.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeProjectsPaginator is a paginator for DescribeProjects
type DescribeProjectsPaginator struct {
	options   DescribeProjectsPaginatorOptions
	client    DescribeProjectsAPIClient
	params    *DescribeProjectsInput
	nextToken *string
	firstPage bool
}

// NewDescribeProjectsPaginator returns a new DescribeProjectsPaginator
func NewDescribeProjectsPaginator(client DescribeProjectsAPIClient, params *DescribeProjectsInput, optFns ...func(*DescribeProjectsPaginatorOptions)) *DescribeProjectsPaginator {
	if params == nil {
		params = &DescribeProjectsInput{}
	}

	options := DescribeProjectsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeProjectsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeProjectsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeProjects page.
func (p *DescribeProjectsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeProjectsOutput, error) {
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

	result, err := p.client.DescribeProjects(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opDescribeProjects(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeProjects",
	}
}
