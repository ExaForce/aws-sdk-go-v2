// Code generated by smithy-go-codegen DO NOT EDIT.

package chatbot

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/chatbot/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates a Chime Webhook Configuration
func (c *Client) UpdateChimeWebhookConfiguration(ctx context.Context, params *UpdateChimeWebhookConfigurationInput, optFns ...func(*Options)) (*UpdateChimeWebhookConfigurationOutput, error) {
	if params == nil {
		params = &UpdateChimeWebhookConfigurationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateChimeWebhookConfiguration", params, optFns, c.addOperationUpdateChimeWebhookConfigurationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateChimeWebhookConfigurationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateChimeWebhookConfigurationInput struct {

	// The ARN of the ChimeWebhookConfiguration to update.
	//
	// This member is required.
	ChatConfigurationArn *string

	// The ARN of the IAM role that defines the permissions for AWS Chatbot. This is a
	// user-defined role that AWS Chatbot will assume. This is not the service-linked
	// role. For more information, see IAM Policies for AWS Chatbot.
	IamRoleArn *string

	// Logging levels include ERROR, INFO, or NONE.
	LoggingLevel *string

	// The ARNs of the SNS topics that deliver notifications to AWS Chatbot.
	SnsTopicArns []string

	// Description of the webhook. Recommend using the convention RoomName/WebhookName
	// . See Chime setup tutorial for more details:
	// https://docs.aws.amazon.com/chatbot/latest/adminguide/chime-setup.html.
	WebhookDescription *string

	// URL for the Chime webhook.
	WebhookUrl *string

	noSmithyDocumentSerde
}

type UpdateChimeWebhookConfigurationOutput struct {

	// Chime webhook configuration.
	WebhookConfiguration *types.ChimeWebhookConfiguration

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateChimeWebhookConfigurationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateChimeWebhookConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateChimeWebhookConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "UpdateChimeWebhookConfiguration"); err != nil {
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
	if err = addOpUpdateChimeWebhookConfigurationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateChimeWebhookConfiguration(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateChimeWebhookConfiguration(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "UpdateChimeWebhookConfiguration",
	}
}
