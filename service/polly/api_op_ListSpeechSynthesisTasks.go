// Code generated by smithy-go-codegen DO NOT EDIT.

package polly

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/polly/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns a list of SpeechSynthesisTask objects ordered by their creation date.
// This operation can filter the tasks by their status, for example, allowing users
// to list only tasks that are completed.
func (c *Client) ListSpeechSynthesisTasks(ctx context.Context, params *ListSpeechSynthesisTasksInput, optFns ...func(*Options)) (*ListSpeechSynthesisTasksOutput, error) {
	if params == nil {
		params = &ListSpeechSynthesisTasksInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListSpeechSynthesisTasks", params, optFns, c.addOperationListSpeechSynthesisTasksMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListSpeechSynthesisTasksOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListSpeechSynthesisTasksInput struct {

	// Maximum number of speech synthesis tasks returned in a List operation.
	MaxResults *int32

	// The pagination token to use in the next request to continue the listing of
	// speech synthesis tasks.
	NextToken *string

	// Status of the speech synthesis tasks returned in a List operation
	Status types.TaskStatus

	noSmithyDocumentSerde
}

type ListSpeechSynthesisTasksOutput struct {

	// An opaque pagination token returned from the previous List operation in this
	// request. If present, this indicates where to continue the listing.
	NextToken *string

	// List of SynthesisTask objects that provides information from the specified task
	// in the list request, including output format, creation time, task status, and so
	// on.
	SynthesisTasks []types.SynthesisTask

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListSpeechSynthesisTasksMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListSpeechSynthesisTasks{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListSpeechSynthesisTasks{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListSpeechSynthesisTasks"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListSpeechSynthesisTasks(options.Region), middleware.Before); err != nil {
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

// ListSpeechSynthesisTasksAPIClient is a client that implements the
// ListSpeechSynthesisTasks operation.
type ListSpeechSynthesisTasksAPIClient interface {
	ListSpeechSynthesisTasks(context.Context, *ListSpeechSynthesisTasksInput, ...func(*Options)) (*ListSpeechSynthesisTasksOutput, error)
}

var _ ListSpeechSynthesisTasksAPIClient = (*Client)(nil)

// ListSpeechSynthesisTasksPaginatorOptions is the paginator options for
// ListSpeechSynthesisTasks
type ListSpeechSynthesisTasksPaginatorOptions struct {
	// Maximum number of speech synthesis tasks returned in a List operation.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListSpeechSynthesisTasksPaginator is a paginator for ListSpeechSynthesisTasks
type ListSpeechSynthesisTasksPaginator struct {
	options   ListSpeechSynthesisTasksPaginatorOptions
	client    ListSpeechSynthesisTasksAPIClient
	params    *ListSpeechSynthesisTasksInput
	nextToken *string
	firstPage bool
}

// NewListSpeechSynthesisTasksPaginator returns a new
// ListSpeechSynthesisTasksPaginator
func NewListSpeechSynthesisTasksPaginator(client ListSpeechSynthesisTasksAPIClient, params *ListSpeechSynthesisTasksInput, optFns ...func(*ListSpeechSynthesisTasksPaginatorOptions)) *ListSpeechSynthesisTasksPaginator {
	if params == nil {
		params = &ListSpeechSynthesisTasksInput{}
	}

	options := ListSpeechSynthesisTasksPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListSpeechSynthesisTasksPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListSpeechSynthesisTasksPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListSpeechSynthesisTasks page.
func (p *ListSpeechSynthesisTasksPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListSpeechSynthesisTasksOutput, error) {
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

	result, err := p.client.ListSpeechSynthesisTasks(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListSpeechSynthesisTasks(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListSpeechSynthesisTasks",
	}
}
