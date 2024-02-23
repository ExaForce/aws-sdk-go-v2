// Code generated by smithy-go-codegen DO NOT EDIT.

package datapipeline

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Deletes a pipeline, its pipeline definition, and its run history. AWS Data
// Pipeline attempts to cancel instances associated with the pipeline that are
// currently being processed by task runners. Deleting a pipeline cannot be undone.
// You cannot query or restore a deleted pipeline. To temporarily pause a pipeline
// instead of deleting it, call SetStatus with the status set to PAUSE on
// individual components. Components that are paused by SetStatus can be resumed.
// POST / HTTP/1.1 Content-Type: application/x-amz-json-1.1 X-Amz-Target:
// DataPipeline.DeletePipeline Content-Length: 50 Host:
// datapipeline.us-east-1.amazonaws.com X-Amz-Date: Mon, 12 Nov 2012 17:49:52 GMT
// Authorization: AuthParams {"pipelineId": "df-06372391ZG65EXAMPLE"}
// x-amzn-RequestId: b7a88c81-0754-11e2-af6f-6bc7a6be60d9 Content-Type:
// application/x-amz-json-1.1 Content-Length: 0 Date: Mon, 12 Nov 2012 17:50:53 GMT
// Unexpected response: 200, OK, undefined
func (c *Client) DeletePipeline(ctx context.Context, params *DeletePipelineInput, optFns ...func(*Options)) (*DeletePipelineOutput, error) {
	if params == nil {
		params = &DeletePipelineInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeletePipeline", params, optFns, c.addOperationDeletePipelineMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeletePipelineOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Contains the parameters for DeletePipeline.
type DeletePipelineInput struct {

	// The ID of the pipeline.
	//
	// This member is required.
	PipelineId *string

	noSmithyDocumentSerde
}

type DeletePipelineOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeletePipelineMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDeletePipeline{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeletePipeline{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DeletePipeline"); err != nil {
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
	if err = addOpDeletePipelineValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeletePipeline(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeletePipeline(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DeletePipeline",
	}
}
