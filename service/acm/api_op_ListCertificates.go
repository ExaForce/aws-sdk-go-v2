// Code generated by smithy-go-codegen DO NOT EDIT.

package acm

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/acm/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves a list of certificate ARNs and domain names. You can request that
// only certificates that match a specific status be listed. You can also filter by
// specific attributes of the certificate. Default filtering returns only RSA_2048
// certificates. For more information, see Filters .
func (c *Client) ListCertificates(ctx context.Context, params *ListCertificatesInput, optFns ...func(*Options)) (*ListCertificatesOutput, error) {
	if params == nil {
		params = &ListCertificatesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListCertificates", params, optFns, c.addOperationListCertificatesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListCertificatesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListCertificatesInput struct {

	// Filter the certificate list by status value.
	CertificateStatuses []types.CertificateStatus

	// Filter the certificate list. For more information, see the Filters structure.
	Includes *types.Filters

	// Use this parameter when paginating results to specify the maximum number of
	// items to return in the response. If additional items exist beyond the number you
	// specify, the NextToken element is sent in the response. Use this NextToken
	// value in a subsequent request to retrieve additional items.
	MaxItems *int32

	// Use this parameter only when paginating results and only in a subsequent
	// request after you receive a response with truncated results. Set it to the value
	// of NextToken from the response you just received.
	NextToken *string

	// Specifies the field to sort results by. If you specify SortBy , you must also
	// specify SortOrder .
	SortBy types.SortBy

	// Specifies the order of sorted results. If you specify SortOrder , you must also
	// specify SortBy .
	SortOrder types.SortOrder

	noSmithyDocumentSerde
}

type ListCertificatesOutput struct {

	// A list of ACM certificates.
	CertificateSummaryList []types.CertificateSummary

	// When the list is truncated, this value is present and contains the value to use
	// for the NextToken parameter in a subsequent pagination request.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListCertificatesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListCertificates{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListCertificates{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListCertificates"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListCertificates(options.Region), middleware.Before); err != nil {
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

// ListCertificatesAPIClient is a client that implements the ListCertificates
// operation.
type ListCertificatesAPIClient interface {
	ListCertificates(context.Context, *ListCertificatesInput, ...func(*Options)) (*ListCertificatesOutput, error)
}

var _ ListCertificatesAPIClient = (*Client)(nil)

// ListCertificatesPaginatorOptions is the paginator options for ListCertificates
type ListCertificatesPaginatorOptions struct {
	// Use this parameter when paginating results to specify the maximum number of
	// items to return in the response. If additional items exist beyond the number you
	// specify, the NextToken element is sent in the response. Use this NextToken
	// value in a subsequent request to retrieve additional items.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListCertificatesPaginator is a paginator for ListCertificates
type ListCertificatesPaginator struct {
	options   ListCertificatesPaginatorOptions
	client    ListCertificatesAPIClient
	params    *ListCertificatesInput
	nextToken *string
	firstPage bool
}

// NewListCertificatesPaginator returns a new ListCertificatesPaginator
func NewListCertificatesPaginator(client ListCertificatesAPIClient, params *ListCertificatesInput, optFns ...func(*ListCertificatesPaginatorOptions)) *ListCertificatesPaginator {
	if params == nil {
		params = &ListCertificatesInput{}
	}

	options := ListCertificatesPaginatorOptions{}
	if params.MaxItems != nil {
		options.Limit = *params.MaxItems
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListCertificatesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListCertificatesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListCertificates page.
func (p *ListCertificatesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListCertificatesOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.MaxItems = limit

	result, err := p.client.ListCertificates(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListCertificates(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListCertificates",
	}
}
