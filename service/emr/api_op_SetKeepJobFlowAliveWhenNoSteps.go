// Code generated by smithy-go-codegen DO NOT EDIT.

package emr

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// You can use the SetKeepJobFlowAliveWhenNoSteps to configure a cluster (job
// flow) to terminate after the step execution, i.e., all your steps are executed.
// If you want a transient cluster that shuts down after the last of the current
// executing steps are completed, you can configure SetKeepJobFlowAliveWhenNoSteps
// to false. If you want a long running cluster, configure
// SetKeepJobFlowAliveWhenNoSteps to true. For more information, see Managing
// Cluster Termination (https://docs.aws.amazon.com/emr/latest/ManagementGuide/UsingEMR_TerminationProtection.html)
// in the Amazon EMR Management Guide.
func (c *Client) SetKeepJobFlowAliveWhenNoSteps(ctx context.Context, params *SetKeepJobFlowAliveWhenNoStepsInput, optFns ...func(*Options)) (*SetKeepJobFlowAliveWhenNoStepsOutput, error) {
	if params == nil {
		params = &SetKeepJobFlowAliveWhenNoStepsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "SetKeepJobFlowAliveWhenNoSteps", params, optFns, c.addOperationSetKeepJobFlowAliveWhenNoStepsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*SetKeepJobFlowAliveWhenNoStepsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type SetKeepJobFlowAliveWhenNoStepsInput struct {

	// A list of strings that uniquely identify the clusters to protect. This
	// identifier is returned by RunJobFlow (https://docs.aws.amazon.com/emr/latest/APIReference/API_RunJobFlow.html)
	// and can also be obtained from DescribeJobFlows (https://docs.aws.amazon.com/emr/latest/APIReference/API_DescribeJobFlows.html)
	// .
	//
	// This member is required.
	JobFlowIds []string

	// A Boolean that indicates whether to terminate the cluster after all steps are
	// executed.
	//
	// This member is required.
	KeepJobFlowAliveWhenNoSteps *bool

	noSmithyDocumentSerde
}

type SetKeepJobFlowAliveWhenNoStepsOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationSetKeepJobFlowAliveWhenNoStepsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpSetKeepJobFlowAliveWhenNoSteps{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpSetKeepJobFlowAliveWhenNoSteps{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "SetKeepJobFlowAliveWhenNoSteps"); err != nil {
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
	if err = addOpSetKeepJobFlowAliveWhenNoStepsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opSetKeepJobFlowAliveWhenNoSteps(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opSetKeepJobFlowAliveWhenNoSteps(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "SetKeepJobFlowAliveWhenNoSteps",
	}
}
