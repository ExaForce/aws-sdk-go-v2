// Code generated by smithy-go-codegen DO NOT EDIT.

package datapipeline

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/datapipeline/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Validates the specified pipeline and starts processing pipeline tasks. If the
// pipeline does not pass validation, activation fails. If you need to pause the
// pipeline to investigate an issue with a component, such as a data source or
// script, call DeactivatePipeline . To activate a finished pipeline, modify the
// end date for the pipeline and then activate it. POST / HTTP/1.1 Content-Type:
// application/x-amz-json-1.1 X-Amz-Target: DataPipeline.ActivatePipeline
// Content-Length: 39 Host: datapipeline.us-east-1.amazonaws.com X-Amz-Date: Mon,
// 12 Nov 2012 17:49:52 GMT Authorization: AuthParams {"pipelineId":
// "df-06372391ZG65EXAMPLE"} HTTP/1.1 200 x-amzn-RequestId:
// ee19d5bf-074e-11e2-af6f-6bc7a6be60d9 Content-Type: application/x-amz-json-1.1
// Content-Length: 2 Date: Mon, 12 Nov 2012 17:50:53 GMT {}
func (c *Client) ActivatePipeline(ctx context.Context, params *ActivatePipelineInput, optFns ...func(*Options)) (*ActivatePipelineOutput, error) {
	if params == nil {
		params = &ActivatePipelineInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ActivatePipeline", params, optFns, c.addOperationActivatePipelineMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ActivatePipelineOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Contains the parameters for ActivatePipeline.
type ActivatePipelineInput struct {

	// The ID of the pipeline.
	//
	// This member is required.
	PipelineId *string

	// A list of parameter values to pass to the pipeline at activation.
	ParameterValues []types.ParameterValue

	// The date and time to resume the pipeline. By default, the pipeline resumes from
	// the last completed execution.
	StartTimestamp *time.Time

	noSmithyDocumentSerde
}

// Contains the output of ActivatePipeline.
type ActivatePipelineOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationActivatePipelineMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpActivatePipeline{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpActivatePipeline{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ActivatePipeline"); err != nil {
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
	if err = addOpActivatePipelineValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opActivatePipeline(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opActivatePipeline(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ActivatePipeline",
	}
}
