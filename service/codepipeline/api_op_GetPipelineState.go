// Code generated by smithy-go-codegen DO NOT EDIT.

package codepipeline

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/codepipeline/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Returns information about the state of a pipeline, including the stages and
// actions. Values returned in the revisionId and revisionUrl fields indicate the
// source revision information, such as the commit ID, for the current state.
func (c *Client) GetPipelineState(ctx context.Context, params *GetPipelineStateInput, optFns ...func(*Options)) (*GetPipelineStateOutput, error) {
	if params == nil {
		params = &GetPipelineStateInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetPipelineState", params, optFns, c.addOperationGetPipelineStateMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetPipelineStateOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input of a GetPipelineState action.
type GetPipelineStateInput struct {

	// The name of the pipeline about which you want to get information.
	//
	// This member is required.
	Name *string

	noSmithyDocumentSerde
}

// Represents the output of a GetPipelineState action.
type GetPipelineStateOutput struct {

	// The date and time the pipeline was created, in timestamp format.
	Created *time.Time

	// The name of the pipeline for which you want to get the state.
	PipelineName *string

	// The version number of the pipeline. A newly created pipeline is always assigned
	// a version number of 1 .
	PipelineVersion *int32

	// A list of the pipeline stage output information, including stage name, state,
	// most recent run details, whether the stage is disabled, and other data.
	StageStates []types.StageState

	// The date and time the pipeline was last updated, in timestamp format.
	Updated *time.Time

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetPipelineStateMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetPipelineState{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetPipelineState{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetPipelineState"); err != nil {
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
	if err = addOpGetPipelineStateValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetPipelineState(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetPipelineState(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetPipelineState",
	}
}
