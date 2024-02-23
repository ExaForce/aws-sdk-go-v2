// Code generated by smithy-go-codegen DO NOT EDIT.

package codecommit

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/codecommit/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Posts a comment on a pull request.
func (c *Client) PostCommentForPullRequest(ctx context.Context, params *PostCommentForPullRequestInput, optFns ...func(*Options)) (*PostCommentForPullRequestOutput, error) {
	if params == nil {
		params = &PostCommentForPullRequestInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PostCommentForPullRequest", params, optFns, c.addOperationPostCommentForPullRequestMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PostCommentForPullRequestOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PostCommentForPullRequestInput struct {

	// The full commit ID of the commit in the source branch that is the current tip
	// of the branch for the pull request when you post the comment.
	//
	// This member is required.
	AfterCommitId *string

	// The full commit ID of the commit in the destination branch that was the tip of
	// the branch at the time the pull request was created.
	//
	// This member is required.
	BeforeCommitId *string

	// The content of your comment on the change.
	//
	// This member is required.
	Content *string

	// The system-generated ID of the pull request. To get this ID, use
	// ListPullRequests .
	//
	// This member is required.
	PullRequestId *string

	// The name of the repository where you want to post a comment on a pull request.
	//
	// This member is required.
	RepositoryName *string

	// A unique, client-generated idempotency token that, when provided in a request,
	// ensures the request cannot be repeated with a changed parameter. If a request is
	// received with the same parameters and a token is included, the request returns
	// information about the initial request that used that token.
	ClientRequestToken *string

	// The location of the change where you want to post your comment. If no location
	// is provided, the comment is posted as a general comment on the pull request
	// difference between the before commit ID and the after commit ID.
	Location *types.Location

	noSmithyDocumentSerde
}

type PostCommentForPullRequestOutput struct {

	// In the directionality of the pull request, the blob ID of the after blob.
	AfterBlobId *string

	// The full commit ID of the commit in the destination branch where the pull
	// request is merged.
	AfterCommitId *string

	// In the directionality of the pull request, the blob ID of the before blob.
	BeforeBlobId *string

	// The full commit ID of the commit in the source branch used to create the pull
	// request, or in the case of an updated pull request, the full commit ID of the
	// commit used to update the pull request.
	BeforeCommitId *string

	// The content of the comment you posted.
	Comment *types.Comment

	// The location of the change where you posted your comment.
	Location *types.Location

	// The system-generated ID of the pull request.
	PullRequestId *string

	// The name of the repository where you posted a comment on a pull request.
	RepositoryName *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationPostCommentForPullRequestMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpPostCommentForPullRequest{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpPostCommentForPullRequest{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "PostCommentForPullRequest"); err != nil {
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
	if err = addIdempotencyToken_opPostCommentForPullRequestMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpPostCommentForPullRequestValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opPostCommentForPullRequest(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpPostCommentForPullRequest struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpPostCommentForPullRequest) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpPostCommentForPullRequest) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*PostCommentForPullRequestInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *PostCommentForPullRequestInput ")
	}

	if input.ClientRequestToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.ClientRequestToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opPostCommentForPullRequestMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpPostCommentForPullRequest{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opPostCommentForPullRequest(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "PostCommentForPullRequest",
	}
}
