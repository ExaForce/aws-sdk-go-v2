// Code generated by smithy-go-codegen DO NOT EDIT.

package kinesisvideo

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns a list of tags associated with the specified stream. In the request,
// you must specify either the StreamName or the StreamARN .
func (c *Client) ListTagsForStream(ctx context.Context, params *ListTagsForStreamInput, optFns ...func(*Options)) (*ListTagsForStreamOutput, error) {
	if params == nil {
		params = &ListTagsForStreamInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListTagsForStream", params, optFns, c.addOperationListTagsForStreamMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListTagsForStreamOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListTagsForStreamInput struct {

	// If you specify this parameter and the result of a ListTagsForStream call is
	// truncated, the response includes a token that you can use in the next request to
	// fetch the next batch of tags.
	NextToken *string

	// The Amazon Resource Name (ARN) of the stream that you want to list tags for.
	StreamARN *string

	// The name of the stream that you want to list tags for.
	StreamName *string

	noSmithyDocumentSerde
}

type ListTagsForStreamOutput struct {

	// If you specify this parameter and the result of a ListTags call is truncated,
	// the response includes a token that you can use in the next request to fetch the
	// next set of tags.
	NextToken *string

	// A map of tag keys and values associated with the specified stream.
	Tags map[string]string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListTagsForStreamMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListTagsForStream{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListTagsForStream{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListTagsForStream"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListTagsForStream(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opListTagsForStream(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListTagsForStream",
	}
}
