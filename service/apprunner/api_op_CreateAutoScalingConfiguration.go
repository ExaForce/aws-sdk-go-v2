// Code generated by smithy-go-codegen DO NOT EDIT.

package apprunner

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/apprunner/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Create an App Runner automatic scaling configuration resource. App Runner
// requires this resource when you create or update App Runner services and you
// require non-default auto scaling settings. You can share an auto scaling
// configuration across multiple services. Create multiple revisions of a
// configuration by calling this action multiple times using the same
// AutoScalingConfigurationName . The call returns incremental
// AutoScalingConfigurationRevision values. When you create a service and configure
// an auto scaling configuration resource, the service uses the latest active
// revision of the auto scaling configuration by default. You can optionally
// configure the service to use a specific revision. Configure a higher MinSize to
// increase the spread of your App Runner service over more Availability Zones in
// the Amazon Web Services Region. The tradeoff is a higher minimal cost. Configure
// a lower MaxSize to control your cost. The tradeoff is lower responsiveness
// during peak demand.
func (c *Client) CreateAutoScalingConfiguration(ctx context.Context, params *CreateAutoScalingConfigurationInput, optFns ...func(*Options)) (*CreateAutoScalingConfigurationOutput, error) {
	if params == nil {
		params = &CreateAutoScalingConfigurationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateAutoScalingConfiguration", params, optFns, c.addOperationCreateAutoScalingConfigurationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateAutoScalingConfigurationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateAutoScalingConfigurationInput struct {

	// A name for the auto scaling configuration. When you use it for the first time
	// in an Amazon Web Services Region, App Runner creates revision number 1 of this
	// name. When you use the same name in subsequent calls, App Runner creates
	// incremental revisions of the configuration. Prior to the release of Auto scale
	// configuration enhancements (https://docs.aws.amazon.com/apprunner/latest/relnotes/release-2023-09-22-auto-scale-config.html)
	// , the name DefaultConfiguration was reserved. This restriction is no longer in
	// place. You can now manage DefaultConfiguration the same way you manage your
	// custom auto scaling configurations. This means you can do the following with the
	// DefaultConfiguration that App Runner provides:
	//   - Create new revisions of the DefaultConfiguration .
	//   - Delete the revisions of the DefaultConfiguration .
	//   - Delete the auto scaling configuration for which the App Runner
	//   DefaultConfiguration was created.
	//   - If you delete the auto scaling configuration you can create another custom
	//   auto scaling configuration with the same DefaultConfiguration name. The
	//   original DefaultConfiguration resource provided by App Runner remains in your
	//   account unless you make changes to it.
	//
	// This member is required.
	AutoScalingConfigurationName *string

	// The maximum number of concurrent requests that you want an instance to process.
	// If the number of concurrent requests exceeds this limit, App Runner scales up
	// your service. Default: 100
	MaxConcurrency *int32

	// The maximum number of instances that your service scales up to. At most MaxSize
	// instances actively serve traffic for your service. Default: 25
	MaxSize *int32

	// The minimum number of instances that App Runner provisions for your service.
	// The service always has at least MinSize provisioned instances. Some of them
	// actively serve traffic. The rest of them (provisioned and inactive instances)
	// are a cost-effective compute capacity reserve and are ready to be quickly
	// activated. You pay for memory usage of all the provisioned instances. You pay
	// for CPU usage of only the active subset. App Runner temporarily doubles the
	// number of provisioned instances during deployments, to maintain the same
	// capacity for both old and new code. Default: 1
	MinSize *int32

	// A list of metadata items that you can associate with your auto scaling
	// configuration resource. A tag is a key-value pair.
	Tags []types.Tag

	noSmithyDocumentSerde
}

type CreateAutoScalingConfigurationOutput struct {

	// A description of the App Runner auto scaling configuration that's created by
	// this request.
	//
	// This member is required.
	AutoScalingConfiguration *types.AutoScalingConfiguration

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateAutoScalingConfigurationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpCreateAutoScalingConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpCreateAutoScalingConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateAutoScalingConfiguration"); err != nil {
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
	if err = addOpCreateAutoScalingConfigurationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateAutoScalingConfiguration(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateAutoScalingConfiguration(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateAutoScalingConfiguration",
	}
}
