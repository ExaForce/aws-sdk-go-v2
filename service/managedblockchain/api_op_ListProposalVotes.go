// Code generated by smithy-go-codegen DO NOT EDIT.

package managedblockchain

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/managedblockchain/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns the list of votes for a specified proposal, including the value of each
// vote and the unique identifier of the member that cast the vote. Applies only to
// Hyperledger Fabric.
func (c *Client) ListProposalVotes(ctx context.Context, params *ListProposalVotesInput, optFns ...func(*Options)) (*ListProposalVotesOutput, error) {
	if params == nil {
		params = &ListProposalVotesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListProposalVotes", params, optFns, c.addOperationListProposalVotesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListProposalVotesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListProposalVotesInput struct {

	// The unique identifier of the network.
	//
	// This member is required.
	NetworkId *string

	// The unique identifier of the proposal.
	//
	// This member is required.
	ProposalId *string

	// The maximum number of votes to return.
	MaxResults *int32

	// The pagination token that indicates the next set of results to retrieve.
	NextToken *string

	noSmithyDocumentSerde
}

type ListProposalVotesOutput struct {

	// The pagination token that indicates the next set of results to retrieve.
	NextToken *string

	// The list of votes.
	ProposalVotes []types.VoteSummary

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListProposalVotesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListProposalVotes{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListProposalVotes{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListProposalVotes"); err != nil {
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
	if err = addOpListProposalVotesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListProposalVotes(options.Region), middleware.Before); err != nil {
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

// ListProposalVotesAPIClient is a client that implements the ListProposalVotes
// operation.
type ListProposalVotesAPIClient interface {
	ListProposalVotes(context.Context, *ListProposalVotesInput, ...func(*Options)) (*ListProposalVotesOutput, error)
}

var _ ListProposalVotesAPIClient = (*Client)(nil)

// ListProposalVotesPaginatorOptions is the paginator options for ListProposalVotes
type ListProposalVotesPaginatorOptions struct {
	// The maximum number of votes to return.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListProposalVotesPaginator is a paginator for ListProposalVotes
type ListProposalVotesPaginator struct {
	options   ListProposalVotesPaginatorOptions
	client    ListProposalVotesAPIClient
	params    *ListProposalVotesInput
	nextToken *string
	firstPage bool
}

// NewListProposalVotesPaginator returns a new ListProposalVotesPaginator
func NewListProposalVotesPaginator(client ListProposalVotesAPIClient, params *ListProposalVotesInput, optFns ...func(*ListProposalVotesPaginatorOptions)) *ListProposalVotesPaginator {
	if params == nil {
		params = &ListProposalVotesInput{}
	}

	options := ListProposalVotesPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListProposalVotesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListProposalVotesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListProposalVotes page.
func (p *ListProposalVotesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListProposalVotesOutput, error) {
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

	result, err := p.client.ListProposalVotes(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListProposalVotes(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListProposalVotes",
	}
}
