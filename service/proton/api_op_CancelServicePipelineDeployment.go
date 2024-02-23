// Code generated by smithy-go-codegen DO NOT EDIT.

package proton

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/proton/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Attempts to cancel a service pipeline deployment on an UpdateServicePipeline
// action, if the deployment is IN_PROGRESS . For more information, see Update a
// service pipeline (https://docs.aws.amazon.com/proton/latest/userguide/ag-svc-pipeline-update.html)
// in the Proton User guide. The following list includes potential cancellation
// scenarios.
//   - If the cancellation attempt succeeds, the resulting deployment state is
//     CANCELLED .
//   - If the cancellation attempt fails, the resulting deployment state is FAILED
//     .
//   - If the current UpdateServicePipeline action succeeds before the cancellation
//     attempt starts, the resulting deployment state is SUCCEEDED and the
//     cancellation attempt has no effect.
func (c *Client) CancelServicePipelineDeployment(ctx context.Context, params *CancelServicePipelineDeploymentInput, optFns ...func(*Options)) (*CancelServicePipelineDeploymentOutput, error) {
	if params == nil {
		params = &CancelServicePipelineDeploymentInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CancelServicePipelineDeployment", params, optFns, c.addOperationCancelServicePipelineDeploymentMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CancelServicePipelineDeploymentOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CancelServicePipelineDeploymentInput struct {

	// The name of the service with the service pipeline deployment to cancel.
	//
	// This member is required.
	ServiceName *string

	noSmithyDocumentSerde
}

type CancelServicePipelineDeploymentOutput struct {

	// The service pipeline detail data that's returned by Proton.
	//
	// This member is required.
	Pipeline *types.ServicePipeline

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCancelServicePipelineDeploymentMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpCancelServicePipelineDeployment{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpCancelServicePipelineDeployment{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CancelServicePipelineDeployment"); err != nil {
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
	if err = addOpCancelServicePipelineDeploymentValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCancelServicePipelineDeployment(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCancelServicePipelineDeployment(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CancelServicePipelineDeployment",
	}
}
