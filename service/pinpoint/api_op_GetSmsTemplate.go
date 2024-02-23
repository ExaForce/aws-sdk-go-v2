// Code generated by smithy-go-codegen DO NOT EDIT.

package pinpoint

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/pinpoint/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves the content and settings of a message template for messages that are
// sent through the SMS channel.
func (c *Client) GetSmsTemplate(ctx context.Context, params *GetSmsTemplateInput, optFns ...func(*Options)) (*GetSmsTemplateOutput, error) {
	if params == nil {
		params = &GetSmsTemplateInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetSmsTemplate", params, optFns, c.addOperationGetSmsTemplateMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetSmsTemplateOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetSmsTemplateInput struct {

	// The name of the message template. A template name must start with an
	// alphanumeric character and can contain a maximum of 128 characters. The
	// characters can be alphanumeric characters, underscores (_), or hyphens (-).
	// Template names are case sensitive.
	//
	// This member is required.
	TemplateName *string

	// The unique identifier for the version of the message template to update,
	// retrieve information about, or delete. To retrieve identifiers and other
	// information for all the versions of a template, use the Template Versions
	// resource. If specified, this value must match the identifier for an existing
	// template version. If specified for an update operation, this value must match
	// the identifier for the latest existing version of the template. This restriction
	// helps ensure that race conditions don't occur. If you don't specify a value for
	// this parameter, Amazon Pinpoint does the following:
	//   - For a get operation, retrieves information about the active version of the
	//   template.
	//   - For an update operation, saves the updates to (overwrites) the latest
	//   existing version of the template, if the create-new-version parameter isn't used
	//   or is set to false.
	//   - For a delete operation, deletes the template, including all versions of the
	//   template.
	Version *string

	noSmithyDocumentSerde
}

type GetSmsTemplateOutput struct {

	// Provides information about the content and settings for a message template that
	// can be used in text messages that are sent through the SMS channel.
	//
	// This member is required.
	SMSTemplateResponse *types.SMSTemplateResponse

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetSmsTemplateMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetSmsTemplate{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetSmsTemplate{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetSmsTemplate"); err != nil {
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
	if err = addOpGetSmsTemplateValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetSmsTemplate(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetSmsTemplate(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetSmsTemplate",
	}
}
