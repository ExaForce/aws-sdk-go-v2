// Code generated by smithy-go-codegen DO NOT EDIT.

package kendra

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/kendra/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Gets information on the settings of query suggestions for an index. This is
// used to check the current settings applied to query suggestions.
// DescribeQuerySuggestionsConfig is currently not supported in the Amazon Web
// Services GovCloud (US-West) region.
func (c *Client) DescribeQuerySuggestionsConfig(ctx context.Context, params *DescribeQuerySuggestionsConfigInput, optFns ...func(*Options)) (*DescribeQuerySuggestionsConfigOutput, error) {
	if params == nil {
		params = &DescribeQuerySuggestionsConfigInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeQuerySuggestionsConfig", params, optFns, c.addOperationDescribeQuerySuggestionsConfigMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeQuerySuggestionsConfigOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeQuerySuggestionsConfigInput struct {

	// The identifier of the index with query suggestions that you want to get
	// information on.
	//
	// This member is required.
	IndexId *string

	noSmithyDocumentSerde
}

type DescribeQuerySuggestionsConfigOutput struct {

	// Configuration information for the document fields/attributes that you want to
	// base query suggestions on.
	AttributeSuggestionsConfig *types.AttributeSuggestionsDescribeConfig

	// TRUE to use all queries, otherwise use only queries that include user
	// information to generate the query suggestions.
	IncludeQueriesWithoutUserInformation *bool

	// The Unix timestamp when query suggestions for an index was last cleared. After
	// you clear suggestions, Amazon Kendra learns new suggestions based on new queries
	// added to the query log from the time you cleared suggestions. Amazon Kendra only
	// considers re-occurences of a query from the time you cleared suggestions.
	LastClearTime *time.Time

	// The Unix timestamp when query suggestions for an index was last updated. Amazon
	// Kendra automatically updates suggestions every 24 hours, after you change a
	// setting or after you apply a block list (https://docs.aws.amazon.com/kendra/latest/dg/query-suggestions.html#query-suggestions-blocklist)
	// .
	LastSuggestionsBuildTime *time.Time

	// The minimum number of unique users who must search a query in order for the
	// query to be eligible to suggest to your users.
	MinimumNumberOfQueryingUsers *int32

	// The minimum number of times a query must be searched in order for the query to
	// be eligible to suggest to your users.
	MinimumQueryCount *int32

	// Whether query suggestions are currently in ENABLED mode or LEARN_ONLY mode. By
	// default, Amazon Kendra enables query suggestions. LEARN_ONLY turns off query
	// suggestions for your users. You can change the mode using the
	// UpdateQuerySuggestionsConfig (https://docs.aws.amazon.com/kendra/latest/dg/API_UpdateQuerySuggestionsConfig.html)
	// API.
	Mode types.Mode

	// How recent your queries are in your query log time window (in days).
	QueryLogLookBackWindowInDays *int32

	// Whether the status of query suggestions settings is currently ACTIVE or UPDATING
	// . Active means the current settings apply and Updating means your changed
	// settings are in the process of applying.
	Status types.QuerySuggestionsStatus

	// The current total count of query suggestions for an index. This count can
	// change when you update your query suggestions settings, if you filter out
	// certain queries from suggestions using a block list, and as the query log
	// accumulates more queries for Amazon Kendra to learn from. If the count is much
	// lower than you expected, it could be because Amazon Kendra needs more queries in
	// the query history to learn from or your current query suggestions settings are
	// too strict.
	TotalSuggestionsCount *int32

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeQuerySuggestionsConfigMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeQuerySuggestionsConfig{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeQuerySuggestionsConfig{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeQuerySuggestionsConfig"); err != nil {
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
	if err = addOpDescribeQuerySuggestionsConfigValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeQuerySuggestionsConfig(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeQuerySuggestionsConfig(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeQuerySuggestionsConfig",
	}
}
