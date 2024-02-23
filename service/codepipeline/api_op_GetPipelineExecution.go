// Code generated by smithy-go-codegen DO NOT EDIT.

package codepipeline

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/codepipeline/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns information about an execution of a pipeline, including details about
// artifacts, the pipeline execution ID, and the name, version, and status of the
// pipeline.
func (c *Client) GetPipelineExecution(ctx context.Context, params *GetPipelineExecutionInput, optFns ...func(*Options)) (*GetPipelineExecutionOutput, error) {
	if params == nil {
		params = &GetPipelineExecutionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetPipelineExecution", params, optFns, c.addOperationGetPipelineExecutionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetPipelineExecutionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input of a GetPipelineExecution action.
type GetPipelineExecutionInput struct {

	// The ID of the pipeline execution about which you want to get execution details.
	//
	// This member is required.
	PipelineExecutionId *string

	// The name of the pipeline about which you want to get execution details.
	//
	// This member is required.
	PipelineName *string

	noSmithyDocumentSerde
}

// Represents the output of a GetPipelineExecution action.
type GetPipelineExecutionOutput struct {

	// Represents information about the execution of a pipeline.
	PipelineExecution *types.PipelineExecution

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetPipelineExecutionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetPipelineExecution{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetPipelineExecution{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetPipelineExecution"); err != nil {
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
	if err = addOpGetPipelineExecutionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetPipelineExecution(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetPipelineExecution(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetPipelineExecution",
	}
}
