// Code generated by smithy-go-codegen DO NOT EDIT.

package kendra

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/kendra/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a block list to exlcude certain queries from suggestions. Any query
// that contains words or phrases specified in the block list is blocked or
// filtered out from being shown as a suggestion. You need to provide the file
// location of your block list text file in your S3 bucket. In your text file,
// enter each block word or phrase on a separate line. For information on the
// current quota limits for block lists, see Quotas for Amazon Kendra (https://docs.aws.amazon.com/kendra/latest/dg/quotas.html)
// . CreateQuerySuggestionsBlockList is currently not supported in the Amazon Web
// Services GovCloud (US-West) region. For an example of creating a block list for
// query suggestions using the Python SDK, see Query suggestions block list (https://docs.aws.amazon.com/kendra/latest/dg/query-suggestions.html#query-suggestions-blocklist)
// .
func (c *Client) CreateQuerySuggestionsBlockList(ctx context.Context, params *CreateQuerySuggestionsBlockListInput, optFns ...func(*Options)) (*CreateQuerySuggestionsBlockListOutput, error) {
	if params == nil {
		params = &CreateQuerySuggestionsBlockListInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateQuerySuggestionsBlockList", params, optFns, c.addOperationCreateQuerySuggestionsBlockListMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateQuerySuggestionsBlockListOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateQuerySuggestionsBlockListInput struct {

	// The identifier of the index you want to create a query suggestions block list
	// for.
	//
	// This member is required.
	IndexId *string

	// A name for the block list. For example, the name 'offensive-words', which
	// includes all offensive words that could appear in user queries and need to be
	// blocked from suggestions.
	//
	// This member is required.
	Name *string

	// The Amazon Resource Name (ARN) of an IAM role with permission to access your S3
	// bucket that contains the block list text file. For more information, see IAM
	// access roles for Amazon Kendra (https://docs.aws.amazon.com/kendra/latest/dg/iam-roles.html)
	// .
	//
	// This member is required.
	RoleArn *string

	// The S3 path to your block list text file in your S3 bucket. Each block word or
	// phrase should be on a separate line in a text file. For information on the
	// current quota limits for block lists, see Quotas for Amazon Kendra (https://docs.aws.amazon.com/kendra/latest/dg/quotas.html)
	// .
	//
	// This member is required.
	SourceS3Path *types.S3Path

	// A token that you provide to identify the request to create a query suggestions
	// block list.
	ClientToken *string

	// A description for the block list. For example, the description "List of all
	// offensive words that can appear in user queries and need to be blocked from
	// suggestions."
	Description *string

	// A list of key-value pairs that identify or categorize the block list. Tag keys
	// and values can consist of Unicode letters, digits, white space, and any of the
	// following symbols: _ . : / = + - @.
	Tags []types.Tag

	noSmithyDocumentSerde
}

type CreateQuerySuggestionsBlockListOutput struct {

	// The identifier of the block list.
	Id *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateQuerySuggestionsBlockListMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateQuerySuggestionsBlockList{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateQuerySuggestionsBlockList{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateQuerySuggestionsBlockList"); err != nil {
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
	if err = addIdempotencyToken_opCreateQuerySuggestionsBlockListMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpCreateQuerySuggestionsBlockListValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateQuerySuggestionsBlockList(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpCreateQuerySuggestionsBlockList struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreateQuerySuggestionsBlockList) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreateQuerySuggestionsBlockList) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CreateQuerySuggestionsBlockListInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CreateQuerySuggestionsBlockListInput ")
	}

	if input.ClientToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.ClientToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opCreateQuerySuggestionsBlockListMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpCreateQuerySuggestionsBlockList{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCreateQuerySuggestionsBlockList(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateQuerySuggestionsBlockList",
	}
}
