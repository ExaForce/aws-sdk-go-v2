// Code generated by smithy-go-codegen DO NOT EDIT.

package cloud9

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/cloud9/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates an Cloud9 development environment, launches an Amazon Elastic Compute
// Cloud (Amazon EC2) instance, and then connects from the instance to the
// environment.
func (c *Client) CreateEnvironmentEC2(ctx context.Context, params *CreateEnvironmentEC2Input, optFns ...func(*Options)) (*CreateEnvironmentEC2Output, error) {
	if params == nil {
		params = &CreateEnvironmentEC2Input{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateEnvironmentEC2", params, optFns, c.addOperationCreateEnvironmentEC2Middlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateEnvironmentEC2Output)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateEnvironmentEC2Input struct {

	// The identifier for the Amazon Machine Image (AMI) that's used to create the EC2
	// instance. To choose an AMI for the instance, you must specify a valid AMI alias
	// or a valid Amazon EC2 Systems Manager (SSM) path. From December 04, 2023, you
	// will be required to include the imageId parameter for the CreateEnvironmentEC2
	// action. This change will be reflected across all direct methods of communicating
	// with the API, such as Amazon Web Services SDK, Amazon Web Services CLI and
	// Amazon Web Services CloudFormation. This change will only affect direct API
	// consumers, and not Cloud9 console users. We recommend using Amazon Linux 2023 as
	// the AMI to create your environment as it is fully supported. Since Ubuntu 18.04
	// has ended standard support as of May 31, 2023, we recommend you choose Ubuntu
	// 22.04. AMI aliases
	//   - Amazon Linux 2: amazonlinux-2-x86_64
	//   - Amazon Linux 2023 (recommended): amazonlinux-2023-x86_64
	//   - Ubuntu 18.04: ubuntu-18.04-x86_64
	//   - Ubuntu 22.04: ubuntu-22.04-x86_64
	// SSM paths
	//   - Amazon Linux 2: resolve:ssm:/aws/service/cloud9/amis/amazonlinux-2-x86_64
	//   - Amazon Linux 2023 (recommended):
	//   resolve:ssm:/aws/service/cloud9/amis/amazonlinux-2023-x86_64
	//   - Ubuntu 18.04: resolve:ssm:/aws/service/cloud9/amis/ubuntu-18.04-x86_64
	//   - Ubuntu 22.04: resolve:ssm:/aws/service/cloud9/amis/ubuntu-22.04-x86_64
	//
	// This member is required.
	ImageId *string

	// The type of instance to connect to the environment (for example, t2.micro ).
	//
	// This member is required.
	InstanceType *string

	// The name of the environment to create. This name is visible to other IAM users
	// in the same Amazon Web Services account.
	//
	// This member is required.
	Name *string

	// The number of minutes until the running instance is shut down after the
	// environment has last been used.
	AutomaticStopTimeMinutes *int32

	// A unique, case-sensitive string that helps Cloud9 to ensure this operation
	// completes no more than one time. For more information, see Client Tokens (https://docs.aws.amazon.com/AWSEC2/latest/APIReference/Run_Instance_Idempotency.html)
	// in the Amazon EC2 API Reference.
	ClientRequestToken *string

	// The connection type used for connecting to an Amazon EC2 environment. Valid
	// values are CONNECT_SSH (default) and CONNECT_SSM (connected through Amazon EC2
	// Systems Manager). For more information, see Accessing no-ingress EC2 instances
	// with Amazon EC2 Systems Manager (https://docs.aws.amazon.com/cloud9/latest/user-guide/ec2-ssm.html)
	// in the Cloud9 User Guide.
	ConnectionType types.ConnectionType

	// The description of the environment to create.
	Description *string

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation . Otherwise, it is
	// UnauthorizedOperation .
	DryRun *bool

	// The Amazon Resource Name (ARN) of the environment owner. This ARN can be the
	// ARN of any IAM principal. If this value is not specified, the ARN defaults to
	// this environment's creator.
	OwnerArn *string

	// The ID of the subnet in Amazon VPC that Cloud9 will use to communicate with the
	// Amazon EC2 instance.
	SubnetId *string

	// An array of key-value pairs that will be associated with the new Cloud9
	// development environment.
	Tags []types.Tag

	noSmithyDocumentSerde
}

type CreateEnvironmentEC2Output struct {

	// The ID of the environment that was created.
	EnvironmentId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateEnvironmentEC2Middlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateEnvironmentEC2{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateEnvironmentEC2{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateEnvironmentEC2"); err != nil {
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
	if err = addOpCreateEnvironmentEC2ValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateEnvironmentEC2(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateEnvironmentEC2(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateEnvironmentEC2",
	}
}
