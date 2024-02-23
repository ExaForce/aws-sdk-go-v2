// Code generated by smithy-go-codegen DO NOT EDIT.

package sagemaker

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/sagemaker/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates an inference component, which is a SageMaker hosting object that you
// can use to deploy a model to an endpoint. In the inference component settings,
// you specify the model, the endpoint, and how the model utilizes the resources
// that the endpoint hosts. You can optimize resource utilization by tailoring how
// the required CPU cores, accelerators, and memory are allocated. You can deploy
// multiple inference components to an endpoint, where each inference component
// contains one model and the resource utilization needs for that individual model.
// After you deploy an inference component, you can directly invoke the associated
// model when you use the InvokeEndpoint API action.
func (c *Client) CreateInferenceComponent(ctx context.Context, params *CreateInferenceComponentInput, optFns ...func(*Options)) (*CreateInferenceComponentOutput, error) {
	if params == nil {
		params = &CreateInferenceComponentInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateInferenceComponent", params, optFns, c.addOperationCreateInferenceComponentMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateInferenceComponentOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateInferenceComponentInput struct {

	// The name of an existing endpoint where you host the inference component.
	//
	// This member is required.
	EndpointName *string

	// A unique name to assign to the inference component.
	//
	// This member is required.
	InferenceComponentName *string

	// Runtime settings for a model that is deployed with an inference component.
	//
	// This member is required.
	RuntimeConfig *types.InferenceComponentRuntimeConfig

	// Details about the resources to deploy with this inference component, including
	// the model, container, and compute resources.
	//
	// This member is required.
	Specification *types.InferenceComponentSpecification

	// The name of an existing production variant where you host the inference
	// component.
	//
	// This member is required.
	VariantName *string

	// A list of key-value pairs associated with the model. For more information, see
	// Tagging Amazon Web Services resources (https://docs.aws.amazon.com/general/latest/gr/aws_tagging.html)
	// in the Amazon Web Services General Reference.
	Tags []types.Tag

	noSmithyDocumentSerde
}

type CreateInferenceComponentOutput struct {

	// The Amazon Resource Name (ARN) of the inference component.
	//
	// This member is required.
	InferenceComponentArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateInferenceComponentMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateInferenceComponent{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateInferenceComponent{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateInferenceComponent"); err != nil {
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
	if err = addOpCreateInferenceComponentValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateInferenceComponent(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateInferenceComponent(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateInferenceComponent",
	}
}
