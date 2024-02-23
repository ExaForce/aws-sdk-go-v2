// Code generated by smithy-go-codegen DO NOT EDIT.

package wisdom

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Removes a URI template from a knowledge base.
func (c *Client) RemoveKnowledgeBaseTemplateUri(ctx context.Context, params *RemoveKnowledgeBaseTemplateUriInput, optFns ...func(*Options)) (*RemoveKnowledgeBaseTemplateUriOutput, error) {
	if params == nil {
		params = &RemoveKnowledgeBaseTemplateUriInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "RemoveKnowledgeBaseTemplateUri", params, optFns, c.addOperationRemoveKnowledgeBaseTemplateUriMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*RemoveKnowledgeBaseTemplateUriOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type RemoveKnowledgeBaseTemplateUriInput struct {

	// The identifier of the knowledge base. This should not be a QUICK_RESPONSES type
	// knowledge base if you're storing Wisdom Content resource to it. Can be either
	// the ID or the ARN. URLs cannot contain the ARN.
	//
	// This member is required.
	KnowledgeBaseId *string

	noSmithyDocumentSerde
}

type RemoveKnowledgeBaseTemplateUriOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationRemoveKnowledgeBaseTemplateUriMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpRemoveKnowledgeBaseTemplateUri{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpRemoveKnowledgeBaseTemplateUri{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "RemoveKnowledgeBaseTemplateUri"); err != nil {
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
	if err = addOpRemoveKnowledgeBaseTemplateUriValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opRemoveKnowledgeBaseTemplateUri(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opRemoveKnowledgeBaseTemplateUri(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "RemoveKnowledgeBaseTemplateUri",
	}
}
