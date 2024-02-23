// Code generated by smithy-go-codegen DO NOT EDIT.

package qbusiness

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/qbusiness/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates an Amazon Q web experience.
func (c *Client) UpdateWebExperience(ctx context.Context, params *UpdateWebExperienceInput, optFns ...func(*Options)) (*UpdateWebExperienceOutput, error) {
	if params == nil {
		params = &UpdateWebExperienceInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateWebExperience", params, optFns, c.addOperationUpdateWebExperienceMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateWebExperienceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateWebExperienceInput struct {

	// The identifier of the Amazon Q application attached to the web experience.
	//
	// This member is required.
	ApplicationId *string

	// The identifier of the Amazon Q web experience.
	//
	// This member is required.
	WebExperienceId *string

	// The authentication configuration of the Amazon Q web experience.
	AuthenticationConfiguration types.WebExperienceAuthConfiguration

	// Determines whether sample prompts are enabled in the web experience for an end
	// user.
	SamplePromptsControlMode types.WebExperienceSamplePromptsControlMode

	// The subtitle of the Amazon Q web experience.
	Subtitle *string

	// The title of the Amazon Q web experience.
	Title *string

	// A customized welcome message for an end user in an Amazon Q web experience.
	WelcomeMessage *string

	noSmithyDocumentSerde
}

type UpdateWebExperienceOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateWebExperienceMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateWebExperience{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateWebExperience{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "UpdateWebExperience"); err != nil {
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
	if err = addOpUpdateWebExperienceValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateWebExperience(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateWebExperience(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "UpdateWebExperience",
	}
}
