// Code generated by smithy-go-codegen DO NOT EDIT.

package mwaa

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/mwaa/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates an Amazon Managed Workflows for Apache Airflow (MWAA) environment.
func (c *Client) UpdateEnvironment(ctx context.Context, params *UpdateEnvironmentInput, optFns ...func(*Options)) (*UpdateEnvironmentOutput, error) {
	if params == nil {
		params = &UpdateEnvironmentInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateEnvironment", params, optFns, c.addOperationUpdateEnvironmentMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateEnvironmentOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateEnvironmentInput struct {

	// The name of your Amazon MWAA environment. For example, MyMWAAEnvironment .
	//
	// This member is required.
	Name *string

	// A list of key-value pairs containing the Apache Airflow configuration options
	// you want to attach to your environment. For more information, see Apache
	// Airflow configuration options (https://docs.aws.amazon.com/mwaa/latest/userguide/configuring-env-variables.html)
	// .
	AirflowConfigurationOptions map[string]string

	// The Apache Airflow version for your environment. To upgrade your environment,
	// specify a newer version of Apache Airflow supported by Amazon MWAA. Before you
	// upgrade an environment, make sure your requirements, DAGs, plugins, and other
	// resources used in your workflows are compatible with the new Apache Airflow
	// version. For more information about updating your resources, see Upgrading an
	// Amazon MWAA environment (https://docs.aws.amazon.com/mwaa/latest/userguide/upgrading-environment.html)
	// . Valid values: 1.10.12 , 2.0.2 , 2.2.2 , 2.4.3 , 2.5.1 , 2.6.3 , 2.7.2 .
	AirflowVersion *string

	// The relative path to the DAGs folder on your Amazon S3 bucket. For example, dags
	// . For more information, see Adding or updating DAGs (https://docs.aws.amazon.com/mwaa/latest/userguide/configuring-dag-folder.html)
	// .
	DagS3Path *string

	// The environment class type. Valid values: mw1.small , mw1.medium , mw1.large .
	// For more information, see Amazon MWAA environment class (https://docs.aws.amazon.com/mwaa/latest/userguide/environment-class.html)
	// .
	EnvironmentClass *string

	// The Amazon Resource Name (ARN) of the execution role in IAM that allows MWAA to
	// access Amazon Web Services resources in your environment. For example,
	// arn:aws:iam::123456789:role/my-execution-role . For more information, see
	// Amazon MWAA Execution role (https://docs.aws.amazon.com/mwaa/latest/userguide/mwaa-create-role.html)
	// .
	ExecutionRoleArn *string

	// The Apache Airflow log types to send to CloudWatch Logs.
	LoggingConfiguration *types.LoggingConfigurationInput

	// The maximum number of workers that you want to run in your environment. MWAA
	// scales the number of Apache Airflow workers up to the number you specify in the
	// MaxWorkers field. For example, 20 . When there are no more tasks running, and no
	// more in the queue, MWAA disposes of the extra workers leaving the one worker
	// that is included with your environment, or the number you specify in MinWorkers .
	MaxWorkers *int32

	// The minimum number of workers that you want to run in your environment. MWAA
	// scales the number of Apache Airflow workers up to the number you specify in the
	// MaxWorkers field. When there are no more tasks running, and no more in the
	// queue, MWAA disposes of the extra workers leaving the worker count you specify
	// in the MinWorkers field. For example, 2 .
	MinWorkers *int32

	// The VPC networking components used to secure and enable network traffic between
	// the Amazon Web Services resources for your environment. For more information,
	// see About networking on Amazon MWAA (https://docs.aws.amazon.com/mwaa/latest/userguide/networking-about.html)
	// .
	NetworkConfiguration *types.UpdateNetworkConfigurationInput

	// The version of the plugins.zip file on your Amazon S3 bucket. You must specify
	// a version each time a plugins.zip file is updated. For more information, see
	// How S3 Versioning works (https://docs.aws.amazon.com/AmazonS3/latest/userguide/versioning-workflows.html)
	// .
	PluginsS3ObjectVersion *string

	// The relative path to the plugins.zip file on your Amazon S3 bucket. For
	// example, plugins.zip . If specified, then the plugins.zip version is required.
	// For more information, see Installing custom plugins (https://docs.aws.amazon.com/mwaa/latest/userguide/configuring-dag-import-plugins.html)
	// .
	PluginsS3Path *string

	// The version of the requirements.txt file on your Amazon S3 bucket. You must
	// specify a version each time a requirements.txt file is updated. For more
	// information, see How S3 Versioning works (https://docs.aws.amazon.com/AmazonS3/latest/userguide/versioning-workflows.html)
	// .
	RequirementsS3ObjectVersion *string

	// The relative path to the requirements.txt file on your Amazon S3 bucket. For
	// example, requirements.txt . If specified, then a file version is required. For
	// more information, see Installing Python dependencies (https://docs.aws.amazon.com/mwaa/latest/userguide/working-dags-dependencies.html)
	// .
	RequirementsS3Path *string

	// The number of Apache Airflow schedulers to run in your Amazon MWAA environment.
	Schedulers *int32

	// The Amazon Resource Name (ARN) of the Amazon S3 bucket where your DAG code and
	// supporting files are stored. For example,
	// arn:aws:s3:::my-airflow-bucket-unique-name . For more information, see Create
	// an Amazon S3 bucket for Amazon MWAA (https://docs.aws.amazon.com/mwaa/latest/userguide/mwaa-s3-bucket.html)
	// .
	SourceBucketArn *string

	// The version of the startup shell script in your Amazon S3 bucket. You must
	// specify the version ID (https://docs.aws.amazon.com/AmazonS3/latest/userguide/versioning-workflows.html)
	// that Amazon S3 assigns to the file every time you update the script. Version IDs
	// are Unicode, UTF-8 encoded, URL-ready, opaque strings that are no more than
	// 1,024 bytes long. The following is an example:
	// 3sL4kqtJlcpXroDTDmJ+rmSpXd3dIbrHY+MTRCxf3vjVBH40Nr8X8gdRQBpUMLUo For more
	// information, see Using a startup script (https://docs.aws.amazon.com/mwaa/latest/userguide/using-startup-script.html)
	// .
	StartupScriptS3ObjectVersion *string

	// The relative path to the startup shell script in your Amazon S3 bucket. For
	// example, s3://mwaa-environment/startup.sh . Amazon MWAA runs the script as your
	// environment starts, and before running the Apache Airflow process. You can use
	// this script to install dependencies, modify Apache Airflow configuration
	// options, and set environment variables. For more information, see Using a
	// startup script (https://docs.aws.amazon.com/mwaa/latest/userguide/using-startup-script.html)
	// .
	StartupScriptS3Path *string

	// The Apache Airflow Web server access mode. For more information, see Apache
	// Airflow access modes (https://docs.aws.amazon.com/mwaa/latest/userguide/configuring-networking.html)
	// .
	WebserverAccessMode types.WebserverAccessMode

	// The day and time of the week in Coordinated Universal Time (UTC) 24-hour
	// standard time to start weekly maintenance updates of your environment in the
	// following format: DAY:HH:MM . For example: TUE:03:30 . You can specify a start
	// time in 30 minute increments only.
	WeeklyMaintenanceWindowStart *string

	noSmithyDocumentSerde
}

type UpdateEnvironmentOutput struct {

	// The Amazon Resource Name (ARN) of the Amazon MWAA environment. For example,
	// arn:aws:airflow:us-east-1:123456789012:environment/MyMWAAEnvironment .
	Arn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateEnvironmentMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateEnvironment{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateEnvironment{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "UpdateEnvironment"); err != nil {
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
	if err = addEndpointPrefix_opUpdateEnvironmentMiddleware(stack); err != nil {
		return err
	}
	if err = addOpUpdateEnvironmentValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateEnvironment(options.Region), middleware.Before); err != nil {
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

type endpointPrefix_opUpdateEnvironmentMiddleware struct {
}

func (*endpointPrefix_opUpdateEnvironmentMiddleware) ID() string {
	return "EndpointHostPrefix"
}

func (m *endpointPrefix_opUpdateEnvironmentMiddleware) HandleFinalize(ctx context.Context, in middleware.FinalizeInput, next middleware.FinalizeHandler) (
	out middleware.FinalizeOutput, metadata middleware.Metadata, err error,
) {
	if smithyhttp.GetHostnameImmutable(ctx) || smithyhttp.IsEndpointHostPrefixDisabled(ctx) {
		return next.HandleFinalize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	req.URL.Host = "api." + req.URL.Host

	return next.HandleFinalize(ctx, in)
}
func addEndpointPrefix_opUpdateEnvironmentMiddleware(stack *middleware.Stack) error {
	return stack.Finalize.Insert(&endpointPrefix_opUpdateEnvironmentMiddleware{}, "ResolveEndpointV2", middleware.After)
}

func newServiceMetadataMiddleware_opUpdateEnvironment(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "UpdateEnvironment",
	}
}
