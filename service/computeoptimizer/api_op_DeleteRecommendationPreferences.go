// Code generated by smithy-go-codegen DO NOT EDIT.

package computeoptimizer

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/computeoptimizer/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Deletes a recommendation preference, such as enhanced infrastructure metrics.
// For more information, see Activating enhanced infrastructure metrics (https://docs.aws.amazon.com/compute-optimizer/latest/ug/enhanced-infrastructure-metrics.html)
// in the Compute Optimizer User Guide.
func (c *Client) DeleteRecommendationPreferences(ctx context.Context, params *DeleteRecommendationPreferencesInput, optFns ...func(*Options)) (*DeleteRecommendationPreferencesOutput, error) {
	if params == nil {
		params = &DeleteRecommendationPreferencesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteRecommendationPreferences", params, optFns, c.addOperationDeleteRecommendationPreferencesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteRecommendationPreferencesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteRecommendationPreferencesInput struct {

	// The name of the recommendation preference to delete.
	//
	// This member is required.
	RecommendationPreferenceNames []types.RecommendationPreferenceName

	// The target resource type of the recommendation preference to delete. The
	// Ec2Instance option encompasses standalone instances and instances that are part
	// of Auto Scaling groups. The AutoScalingGroup option encompasses only instances
	// that are part of an Auto Scaling group. The valid values for this parameter are
	// Ec2Instance and AutoScalingGroup .
	//
	// This member is required.
	ResourceType types.ResourceType

	// An object that describes the scope of the recommendation preference to delete.
	// You can delete recommendation preferences that are created at the organization
	// level (for management accounts of an organization only), account level, and
	// resource level. For more information, see Activating enhanced infrastructure
	// metrics (https://docs.aws.amazon.com/compute-optimizer/latest/ug/enhanced-infrastructure-metrics.html)
	// in the Compute Optimizer User Guide.
	Scope *types.Scope

	noSmithyDocumentSerde
}

type DeleteRecommendationPreferencesOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeleteRecommendationPreferencesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpDeleteRecommendationPreferences{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpDeleteRecommendationPreferences{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DeleteRecommendationPreferences"); err != nil {
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
	if err = addOpDeleteRecommendationPreferencesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteRecommendationPreferences(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeleteRecommendationPreferences(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DeleteRecommendationPreferences",
	}
}
