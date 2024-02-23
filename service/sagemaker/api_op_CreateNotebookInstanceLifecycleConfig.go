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

// Creates a lifecycle configuration that you can associate with a notebook
// instance. A lifecycle configuration is a collection of shell scripts that run
// when you create or start a notebook instance. Each lifecycle configuration
// script has a limit of 16384 characters. The value of the $PATH environment
// variable that is available to both scripts is /sbin:bin:/usr/sbin:/usr/bin .
// View Amazon CloudWatch Logs for notebook instance lifecycle configurations in
// log group /aws/sagemaker/NotebookInstances in log stream
// [notebook-instance-name]/[LifecycleConfigHook] . Lifecycle configuration scripts
// cannot run for longer than 5 minutes. If a script runs for longer than 5
// minutes, it fails and the notebook instance is not created or started. For
// information about notebook instance lifestyle configurations, see Step 2.1:
// (Optional) Customize a Notebook Instance (https://docs.aws.amazon.com/sagemaker/latest/dg/notebook-lifecycle-config.html)
// .
func (c *Client) CreateNotebookInstanceLifecycleConfig(ctx context.Context, params *CreateNotebookInstanceLifecycleConfigInput, optFns ...func(*Options)) (*CreateNotebookInstanceLifecycleConfigOutput, error) {
	if params == nil {
		params = &CreateNotebookInstanceLifecycleConfigInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateNotebookInstanceLifecycleConfig", params, optFns, c.addOperationCreateNotebookInstanceLifecycleConfigMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateNotebookInstanceLifecycleConfigOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateNotebookInstanceLifecycleConfigInput struct {

	// The name of the lifecycle configuration.
	//
	// This member is required.
	NotebookInstanceLifecycleConfigName *string

	// A shell script that runs only once, when you create a notebook instance. The
	// shell script must be a base64-encoded string.
	OnCreate []types.NotebookInstanceLifecycleHook

	// A shell script that runs every time you start a notebook instance, including
	// when you create the notebook instance. The shell script must be a base64-encoded
	// string.
	OnStart []types.NotebookInstanceLifecycleHook

	noSmithyDocumentSerde
}

type CreateNotebookInstanceLifecycleConfigOutput struct {

	// The Amazon Resource Name (ARN) of the lifecycle configuration.
	NotebookInstanceLifecycleConfigArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateNotebookInstanceLifecycleConfigMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateNotebookInstanceLifecycleConfig{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateNotebookInstanceLifecycleConfig{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateNotebookInstanceLifecycleConfig"); err != nil {
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
	if err = addOpCreateNotebookInstanceLifecycleConfigValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateNotebookInstanceLifecycleConfig(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateNotebookInstanceLifecycleConfig(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateNotebookInstanceLifecycleConfig",
	}
}
