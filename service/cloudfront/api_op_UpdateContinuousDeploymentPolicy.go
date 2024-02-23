// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudfront

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/cloudfront/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates a continuous deployment policy. You can update a continuous deployment
// policy to enable or disable it, to change the percentage of traffic that it
// sends to the staging distribution, or to change the staging distribution that it
// sends traffic to. When you update a continuous deployment policy configuration,
// all the fields are updated with the values that are provided in the request. You
// cannot update some fields independent of others. To update a continuous
// deployment policy configuration:
//   - Use GetContinuousDeploymentPolicyConfig to get the current configuration.
//   - Locally modify the fields in the continuous deployment policy configuration
//     that you want to update.
//   - Use UpdateContinuousDeploymentPolicy , providing the entire continuous
//     deployment policy configuration, including the fields that you modified and
//     those that you didn't.
func (c *Client) UpdateContinuousDeploymentPolicy(ctx context.Context, params *UpdateContinuousDeploymentPolicyInput, optFns ...func(*Options)) (*UpdateContinuousDeploymentPolicyOutput, error) {
	if params == nil {
		params = &UpdateContinuousDeploymentPolicyInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateContinuousDeploymentPolicy", params, optFns, c.addOperationUpdateContinuousDeploymentPolicyMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateContinuousDeploymentPolicyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateContinuousDeploymentPolicyInput struct {

	// The continuous deployment policy configuration.
	//
	// This member is required.
	ContinuousDeploymentPolicyConfig *types.ContinuousDeploymentPolicyConfig

	// The identifier of the continuous deployment policy that you are updating.
	//
	// This member is required.
	Id *string

	// The current version ( ETag value) of the continuous deployment policy that you
	// are updating.
	IfMatch *string

	noSmithyDocumentSerde
}

type UpdateContinuousDeploymentPolicyOutput struct {

	// A continuous deployment policy.
	ContinuousDeploymentPolicy *types.ContinuousDeploymentPolicy

	// The version identifier for the current version of the continuous deployment
	// policy.
	ETag *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateContinuousDeploymentPolicyMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestxml_serializeOpUpdateContinuousDeploymentPolicy{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpUpdateContinuousDeploymentPolicy{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "UpdateContinuousDeploymentPolicy"); err != nil {
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
	if err = addOpUpdateContinuousDeploymentPolicyValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateContinuousDeploymentPolicy(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateContinuousDeploymentPolicy(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "UpdateContinuousDeploymentPolicy",
	}
}
