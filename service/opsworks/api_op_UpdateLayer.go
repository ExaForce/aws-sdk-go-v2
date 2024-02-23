// Code generated by smithy-go-codegen DO NOT EDIT.

package opsworks

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/opsworks/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates a specified layer. Required Permissions: To use this action, an IAM
// user must have a Manage permissions level for the stack, or an attached policy
// that explicitly grants permissions. For more information on user permissions,
// see Managing User Permissions (https://docs.aws.amazon.com/opsworks/latest/userguide/opsworks-security-users.html)
// .
func (c *Client) UpdateLayer(ctx context.Context, params *UpdateLayerInput, optFns ...func(*Options)) (*UpdateLayerOutput, error) {
	if params == nil {
		params = &UpdateLayerInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateLayer", params, optFns, c.addOperationUpdateLayerMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateLayerOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateLayerInput struct {

	// The layer ID.
	//
	// This member is required.
	LayerId *string

	// One or more user-defined key/value pairs to be added to the stack attributes.
	Attributes map[string]string

	// Whether to automatically assign an Elastic IP address (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/elastic-ip-addresses-eip.html)
	// to the layer's instances. For more information, see How to Edit a Layer (https://docs.aws.amazon.com/opsworks/latest/userguide/workinglayers-basics-edit.html)
	// .
	AutoAssignElasticIps *bool

	// For stacks that are running in a VPC, whether to automatically assign a public
	// IP address to the layer's instances. For more information, see How to Edit a
	// Layer (https://docs.aws.amazon.com/opsworks/latest/userguide/workinglayers-basics-edit.html)
	// .
	AutoAssignPublicIps *bool

	// Specifies CloudWatch Logs configuration options for the layer. For more
	// information, see CloudWatchLogsLogStream .
	CloudWatchLogsConfiguration *types.CloudWatchLogsConfiguration

	// The ARN of an IAM profile to be used for all of the layer's EC2 instances. For
	// more information about IAM ARNs, see Using Identifiers (https://docs.aws.amazon.com/IAM/latest/UserGuide/Using_Identifiers.html)
	// .
	CustomInstanceProfileArn *string

	// A JSON-formatted string containing custom stack configuration and deployment
	// attributes to be installed on the layer's instances. For more information, see
	// Using Custom JSON (https://docs.aws.amazon.com/opsworks/latest/userguide/workingcookbook-json-override.html)
	// .
	CustomJson *string

	// A LayerCustomRecipes object that specifies the layer's custom recipes.
	CustomRecipes *types.Recipes

	// An array containing the layer's custom security group IDs.
	CustomSecurityGroupIds []string

	// Whether to disable auto healing for the layer.
	EnableAutoHealing *bool

	// Whether to install operating system and package updates when the instance
	// boots. The default value is true . To control when updates are installed, set
	// this value to false . You must then update your instances manually by using
	// CreateDeployment to run the update_dependencies stack command or manually
	// running yum (Amazon Linux) or apt-get (Ubuntu) on the instances. We strongly
	// recommend using the default value of true , to ensure that your instances have
	// the latest security updates.
	InstallUpdatesOnBoot *bool

	//
	LifecycleEventConfiguration *types.LifecycleEventConfiguration

	// The layer name, which is used by the console.
	Name *string

	// An array of Package objects that describe the layer's packages.
	Packages []string

	// For custom layers only, use this parameter to specify the layer's short name,
	// which is used internally by AWS OpsWorks Stacks and by Chef. The short name is
	// also used as the name for the directory where your app files are installed. It
	// can have a maximum of 200 characters and must be in the following format:
	// /\A[a-z0-9\-\_\.]+\Z/. The built-in layers' short names are defined by AWS
	// OpsWorks Stacks. For more information, see the Layer Reference (https://docs.aws.amazon.com/opsworks/latest/userguide/layers.html)
	Shortname *string

	// Whether to use Amazon EBS-optimized instances.
	UseEbsOptimizedInstances *bool

	// A VolumeConfigurations object that describes the layer's Amazon EBS volumes.
	VolumeConfigurations []types.VolumeConfiguration

	noSmithyDocumentSerde
}

type UpdateLayerOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateLayerMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateLayer{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateLayer{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "UpdateLayer"); err != nil {
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
	if err = addOpUpdateLayerValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateLayer(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateLayer(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "UpdateLayer",
	}
}
