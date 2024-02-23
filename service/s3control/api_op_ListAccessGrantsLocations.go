// Code generated by smithy-go-codegen DO NOT EDIT.

package s3control

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	s3controlcust "github.com/aws/aws-sdk-go-v2/service/s3control/internal/customizations"
	"github.com/aws/aws-sdk-go-v2/service/s3control/types"
	smithy "github.com/aws/smithy-go"
	"github.com/aws/smithy-go/middleware"
	"github.com/aws/smithy-go/ptr"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"strings"
)

// Returns a list of the locations registered in your S3 Access Grants instance.
// Permissions You must have the s3:ListAccessGrantsLocations permission to use
// this operation.
func (c *Client) ListAccessGrantsLocations(ctx context.Context, params *ListAccessGrantsLocationsInput, optFns ...func(*Options)) (*ListAccessGrantsLocationsOutput, error) {
	if params == nil {
		params = &ListAccessGrantsLocationsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListAccessGrantsLocations", params, optFns, c.addOperationListAccessGrantsLocationsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListAccessGrantsLocationsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListAccessGrantsLocationsInput struct {

	// The ID of the Amazon Web Services account that is making this request.
	//
	// This member is required.
	AccountId *string

	// The S3 path to the location that you are registering. The location scope can be
	// the default S3 location s3:// , the S3 path to a bucket s3:// , or the S3 path
	// to a bucket and prefix s3:/// . A prefix in S3 is a string of characters at the
	// beginning of an object key name used to organize the objects that you store in
	// your S3 buckets. For example, object key names that start with the engineering/
	// prefix or object key names that start with the marketing/campaigns/ prefix.
	LocationScope *string

	// The maximum number of access grants that you would like returned in the List
	// Access Grants response. If the results include the pagination token NextToken ,
	// make another call using the NextToken to determine if there are more results.
	MaxResults int32

	// A pagination token to request the next page of results. Pass this value into a
	// subsequent List Access Grants Locations request in order to retrieve the next
	// page of results.
	NextToken *string

	noSmithyDocumentSerde
}

func (in *ListAccessGrantsLocationsInput) bindEndpointParams(p *EndpointParameters) {
	p.AccountId = in.AccountId
	p.RequiresAccountId = ptr.Bool(true)
}

type ListAccessGrantsLocationsOutput struct {

	// A container for a list of registered locations in an S3 Access Grants instance.
	AccessGrantsLocationsList []types.ListAccessGrantsLocationsEntry

	// A pagination token to request the next page of results. Pass this value into a
	// subsequent List Access Grants Locations request in order to retrieve the next
	// page of results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListAccessGrantsLocationsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestxml_serializeOpListAccessGrantsLocations{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpListAccessGrantsLocations{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListAccessGrantsLocations"); err != nil {
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
	if err = s3controlcust.AddUpdateOutpostARN(stack); err != nil {
		return err
	}
	if err = addSetLegacyContextSigningOptionsMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddContentChecksumMiddleware(stack); err != nil {
		return err
	}
	if err = addEndpointPrefix_opListAccessGrantsLocationsMiddleware(stack); err != nil {
		return err
	}
	if err = addOpListAccessGrantsLocationsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListAccessGrantsLocations(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addMetadataRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addRecursionDetection(stack); err != nil {
		return err
	}
	if err = addListAccessGrantsLocationsUpdateEndpoint(stack, options); err != nil {
		return err
	}
	if err = addStashOperationInput(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = v4.AddContentSHA256HeaderMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	if err = addDisableHTTPSMiddleware(stack, options); err != nil {
		return err
	}
	if err = s3controlcust.AddDisableHostPrefixMiddleware(stack); err != nil {
		return err
	}
	return nil
}

type endpointPrefix_opListAccessGrantsLocationsMiddleware struct {
}

func (*endpointPrefix_opListAccessGrantsLocationsMiddleware) ID() string {
	return "EndpointHostPrefix"
}

func (m *endpointPrefix_opListAccessGrantsLocationsMiddleware) HandleFinalize(ctx context.Context, in middleware.FinalizeInput, next middleware.FinalizeHandler) (
	out middleware.FinalizeOutput, metadata middleware.Metadata, err error,
) {
	if smithyhttp.GetHostnameImmutable(ctx) || smithyhttp.IsEndpointHostPrefixDisabled(ctx) {
		return next.HandleFinalize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	opaqueInput := getOperationInput(ctx)
	input, ok := opaqueInput.(*ListAccessGrantsLocationsInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input type %T", opaqueInput)
	}

	var prefix strings.Builder
	if input.AccountId == nil {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("AccountId forms part of the endpoint host and so may not be nil")}
	} else if !smithyhttp.ValidHostLabel(*input.AccountId) {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("AccountId forms part of the endpoint host and so must match \"[a-zA-Z0-9-]{1,63}\", but was \"%s\"", *input.AccountId)}
	} else {
		prefix.WriteString(*input.AccountId)
	}
	prefix.WriteString(".")
	req.URL.Host = prefix.String() + req.URL.Host

	return next.HandleFinalize(ctx, in)
}
func addEndpointPrefix_opListAccessGrantsLocationsMiddleware(stack *middleware.Stack) error {
	return stack.Finalize.Insert(&endpointPrefix_opListAccessGrantsLocationsMiddleware{}, "ResolveEndpointV2", middleware.After)
}

// ListAccessGrantsLocationsAPIClient is a client that implements the
// ListAccessGrantsLocations operation.
type ListAccessGrantsLocationsAPIClient interface {
	ListAccessGrantsLocations(context.Context, *ListAccessGrantsLocationsInput, ...func(*Options)) (*ListAccessGrantsLocationsOutput, error)
}

var _ ListAccessGrantsLocationsAPIClient = (*Client)(nil)

// ListAccessGrantsLocationsPaginatorOptions is the paginator options for
// ListAccessGrantsLocations
type ListAccessGrantsLocationsPaginatorOptions struct {
	// The maximum number of access grants that you would like returned in the List
	// Access Grants response. If the results include the pagination token NextToken ,
	// make another call using the NextToken to determine if there are more results.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListAccessGrantsLocationsPaginator is a paginator for ListAccessGrantsLocations
type ListAccessGrantsLocationsPaginator struct {
	options   ListAccessGrantsLocationsPaginatorOptions
	client    ListAccessGrantsLocationsAPIClient
	params    *ListAccessGrantsLocationsInput
	nextToken *string
	firstPage bool
}

// NewListAccessGrantsLocationsPaginator returns a new
// ListAccessGrantsLocationsPaginator
func NewListAccessGrantsLocationsPaginator(client ListAccessGrantsLocationsAPIClient, params *ListAccessGrantsLocationsInput, optFns ...func(*ListAccessGrantsLocationsPaginatorOptions)) *ListAccessGrantsLocationsPaginator {
	if params == nil {
		params = &ListAccessGrantsLocationsInput{}
	}

	options := ListAccessGrantsLocationsPaginatorOptions{}
	if params.MaxResults != 0 {
		options.Limit = params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListAccessGrantsLocationsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListAccessGrantsLocationsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListAccessGrantsLocations page.
func (p *ListAccessGrantsLocationsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListAccessGrantsLocationsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	params.MaxResults = p.options.Limit

	result, err := p.client.ListAccessGrantsLocations(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListAccessGrantsLocations(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListAccessGrantsLocations",
	}
}

func copyListAccessGrantsLocationsInputForUpdateEndpoint(params interface{}) (interface{}, error) {
	input, ok := params.(*ListAccessGrantsLocationsInput)
	if !ok {
		return nil, fmt.Errorf("expect *ListAccessGrantsLocationsInput type, got %T", params)
	}
	cpy := *input
	return &cpy, nil
}
func (in *ListAccessGrantsLocationsInput) copy() interface{} {
	v := *in
	return &v
}
func backFillListAccessGrantsLocationsAccountID(input interface{}, v string) error {
	in := input.(*ListAccessGrantsLocationsInput)
	if in.AccountId != nil {
		if !strings.EqualFold(*in.AccountId, v) {
			return fmt.Errorf("error backfilling account id")
		}
		return nil
	}
	in.AccountId = &v
	return nil
}
func addListAccessGrantsLocationsUpdateEndpoint(stack *middleware.Stack, options Options) error {
	return s3controlcust.UpdateEndpoint(stack, s3controlcust.UpdateEndpointOptions{
		Accessor: s3controlcust.UpdateEndpointParameterAccessor{GetARNInput: nopGetARNAccessor,
			BackfillAccountID: nopBackfillAccountIDAccessor,
			GetOutpostIDInput: nopGetOutpostIDFromInput,
			UpdateARNField:    nopSetARNAccessor,
			CopyInput:         copyListAccessGrantsLocationsInputForUpdateEndpoint,
		},
		EndpointResolver:        options.EndpointResolver,
		EndpointResolverOptions: options.EndpointOptions,
		UseARNRegion:            options.UseARNRegion,
	})
}
