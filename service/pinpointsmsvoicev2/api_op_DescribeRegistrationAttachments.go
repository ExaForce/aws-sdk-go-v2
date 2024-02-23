// Code generated by smithy-go-codegen DO NOT EDIT.

package pinpointsmsvoicev2

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/pinpointsmsvoicev2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves the specified registration attachments or all registration
// attachments associated with your Amazon Web Services account.
func (c *Client) DescribeRegistrationAttachments(ctx context.Context, params *DescribeRegistrationAttachmentsInput, optFns ...func(*Options)) (*DescribeRegistrationAttachmentsOutput, error) {
	if params == nil {
		params = &DescribeRegistrationAttachmentsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeRegistrationAttachments", params, optFns, c.addOperationDescribeRegistrationAttachmentsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeRegistrationAttachmentsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeRegistrationAttachmentsInput struct {

	// An array of RegistrationAttachmentFilter objects to filter the results.
	Filters []types.RegistrationAttachmentFilter

	// The maximum number of results to return per each request.
	MaxResults *int32

	// The token to be used for the next set of paginated results. You don't need to
	// supply a value for this field in the initial request.
	NextToken *string

	// The unique identifier of registration attachments to find. This is an array of
	// RegistrationAttachmentId.
	RegistrationAttachmentIds []string

	noSmithyDocumentSerde
}

type DescribeRegistrationAttachmentsOutput struct {

	// An array of RegistrationAttachments objects that contain the details for the
	// requested registration attachments.
	//
	// This member is required.
	RegistrationAttachments []types.RegistrationAttachmentsInformation

	// The token to be used for the next set of paginated results. You don't need to
	// supply a value for this field in the initial request.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeRegistrationAttachmentsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpDescribeRegistrationAttachments{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpDescribeRegistrationAttachments{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeRegistrationAttachments"); err != nil {
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
	if err = addOpDescribeRegistrationAttachmentsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeRegistrationAttachments(options.Region), middleware.Before); err != nil {
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

// DescribeRegistrationAttachmentsAPIClient is a client that implements the
// DescribeRegistrationAttachments operation.
type DescribeRegistrationAttachmentsAPIClient interface {
	DescribeRegistrationAttachments(context.Context, *DescribeRegistrationAttachmentsInput, ...func(*Options)) (*DescribeRegistrationAttachmentsOutput, error)
}

var _ DescribeRegistrationAttachmentsAPIClient = (*Client)(nil)

// DescribeRegistrationAttachmentsPaginatorOptions is the paginator options for
// DescribeRegistrationAttachments
type DescribeRegistrationAttachmentsPaginatorOptions struct {
	// The maximum number of results to return per each request.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeRegistrationAttachmentsPaginator is a paginator for
// DescribeRegistrationAttachments
type DescribeRegistrationAttachmentsPaginator struct {
	options   DescribeRegistrationAttachmentsPaginatorOptions
	client    DescribeRegistrationAttachmentsAPIClient
	params    *DescribeRegistrationAttachmentsInput
	nextToken *string
	firstPage bool
}

// NewDescribeRegistrationAttachmentsPaginator returns a new
// DescribeRegistrationAttachmentsPaginator
func NewDescribeRegistrationAttachmentsPaginator(client DescribeRegistrationAttachmentsAPIClient, params *DescribeRegistrationAttachmentsInput, optFns ...func(*DescribeRegistrationAttachmentsPaginatorOptions)) *DescribeRegistrationAttachmentsPaginator {
	if params == nil {
		params = &DescribeRegistrationAttachmentsInput{}
	}

	options := DescribeRegistrationAttachmentsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeRegistrationAttachmentsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeRegistrationAttachmentsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeRegistrationAttachments page.
func (p *DescribeRegistrationAttachmentsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeRegistrationAttachmentsOutput, error) {
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

	result, err := p.client.DescribeRegistrationAttachments(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opDescribeRegistrationAttachments(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeRegistrationAttachments",
	}
}
