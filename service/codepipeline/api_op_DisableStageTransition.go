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

// Prevents artifacts in a pipeline from transitioning to the next stage in the
// pipeline.
func (c *Client) DisableStageTransition(ctx context.Context, params *DisableStageTransitionInput, optFns ...func(*Options)) (*DisableStageTransitionOutput, error) {
	if params == nil {
		params = &DisableStageTransitionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DisableStageTransition", params, optFns, c.addOperationDisableStageTransitionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DisableStageTransitionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input of a DisableStageTransition action.
type DisableStageTransitionInput struct {

	// The name of the pipeline in which you want to disable the flow of artifacts
	// from one stage to another.
	//
	// This member is required.
	PipelineName *string

	// The reason given to the user that a stage is disabled, such as waiting for
	// manual approval or manual tests. This message is displayed in the pipeline
	// console UI.
	//
	// This member is required.
	Reason *string

	// The name of the stage where you want to disable the inbound or outbound
	// transition of artifacts.
	//
	// This member is required.
	StageName *string

	// Specifies whether artifacts are prevented from transitioning into the stage and
	// being processed by the actions in that stage (inbound), or prevented from
	// transitioning from the stage after they have been processed by the actions in
	// that stage (outbound).
	//
	// This member is required.
	TransitionType types.StageTransitionType

	noSmithyDocumentSerde
}

type DisableStageTransitionOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDisableStageTransitionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDisableStageTransition{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDisableStageTransition{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DisableStageTransition"); err != nil {
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
	if err = addOpDisableStageTransitionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDisableStageTransition(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDisableStageTransition(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DisableStageTransition",
	}
}
