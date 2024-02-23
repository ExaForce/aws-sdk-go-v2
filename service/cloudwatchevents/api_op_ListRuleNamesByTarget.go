// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudwatchevents

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists the rules for the specified target. You can see which of the rules in
// Amazon EventBridge can invoke a specific target in your account.
func (c *Client) ListRuleNamesByTarget(ctx context.Context, params *ListRuleNamesByTargetInput, optFns ...func(*Options)) (*ListRuleNamesByTargetOutput, error) {
	if params == nil {
		params = &ListRuleNamesByTargetInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListRuleNamesByTarget", params, optFns, c.addOperationListRuleNamesByTargetMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListRuleNamesByTargetOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListRuleNamesByTargetInput struct {

	// The Amazon Resource Name (ARN) of the target resource.
	//
	// This member is required.
	TargetArn *string

	// The name or ARN of the event bus to list rules for. If you omit this, the
	// default event bus is used.
	EventBusName *string

	// The maximum number of results to return.
	Limit *int32

	// The token returned by a previous call to retrieve the next set of results.
	NextToken *string

	noSmithyDocumentSerde
}

type ListRuleNamesByTargetOutput struct {

	// Indicates whether there are additional results to retrieve. If there are no
	// more results, the value is null.
	NextToken *string

	// The names of the rules that can invoke the given target.
	RuleNames []string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListRuleNamesByTargetMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListRuleNamesByTarget{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListRuleNamesByTarget{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListRuleNamesByTarget"); err != nil {
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
	if err = addOpListRuleNamesByTargetValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListRuleNamesByTarget(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opListRuleNamesByTarget(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListRuleNamesByTarget",
	}
}
