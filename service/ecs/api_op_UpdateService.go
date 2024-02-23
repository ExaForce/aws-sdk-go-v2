// Code generated by smithy-go-codegen DO NOT EDIT.

package ecs

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/ecs/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Modifies the parameters of a service. For services using the rolling update ( ECS
// ) you can update the desired count, deployment configuration, network
// configuration, load balancers, service registries, enable ECS managed tags
// option, propagate tags option, task placement constraints and strategies, and
// task definition. When you update any of these parameters, Amazon ECS starts new
// tasks with the new configuration. You can attach Amazon EBS volumes to Amazon
// ECS tasks by configuring the volume when starting or running a task, or when
// creating or updating a service. For more infomation, see Amazon EBS volumes (https://docs.aws.amazon.com/AmazonECS/latest/developerguide/ebs-volumes.html#ebs-volume-types)
// in the Amazon Elastic Container Service Developer Guide. You can update your
// volume configurations and trigger a new deployment. volumeConfigurations is
// only supported for REPLICA service and not DAEMON service. If you leave
// volumeConfigurations null , it doesn't trigger a new deployment. For more
// infomation on volumes, see Amazon EBS volumes (https://docs.aws.amazon.com/AmazonECS/latest/developerguide/ebs-volumes.html#ebs-volume-types)
// in the Amazon Elastic Container Service Developer Guide. For services using the
// blue/green ( CODE_DEPLOY ) deployment controller, only the desired count,
// deployment configuration, health check grace period, task placement constraints
// and strategies, enable ECS managed tags option, and propagate tags can be
// updated using this API. If the network configuration, platform version, task
// definition, or load balancer need to be updated, create a new CodeDeploy
// deployment. For more information, see CreateDeployment (https://docs.aws.amazon.com/codedeploy/latest/APIReference/API_CreateDeployment.html)
// in the CodeDeploy API Reference. For services using an external deployment
// controller, you can update only the desired count, task placement constraints
// and strategies, health check grace period, enable ECS managed tags option, and
// propagate tags option, using this API. If the launch type, load balancer,
// network configuration, platform version, or task definition need to be updated,
// create a new task set For more information, see CreateTaskSet . You can add to
// or subtract from the number of instantiations of a task definition in a service
// by specifying the cluster that the service is running in and a new desiredCount
// parameter. You can attach Amazon EBS volumes to Amazon ECS tasks by configuring
// the volume when starting or running a task, or when creating or updating a
// service. For more infomation, see Amazon EBS volumes (https://docs.aws.amazon.com/AmazonECS/latest/developerguide/ebs-volumes.html#ebs-volume-types)
// in the Amazon Elastic Container Service Developer Guide. If you have updated the
// container image of your application, you can create a new task definition with
// that image and deploy it to your service. The service scheduler uses the minimum
// healthy percent and maximum percent parameters (in the service's deployment
// configuration) to determine the deployment strategy. If your updated Docker
// image uses the same tag as what is in the existing task definition for your
// service (for example, my_image:latest ), you don't need to create a new revision
// of your task definition. You can update the service using the forceNewDeployment
// option. The new tasks launched by the deployment pull the current image/tag
// combination from your repository when they start. You can also update the
// deployment configuration of a service. When a deployment is triggered by
// updating the task definition of a service, the service scheduler uses the
// deployment configuration parameters, minimumHealthyPercent and maximumPercent ,
// to determine the deployment strategy.
//   - If minimumHealthyPercent is below 100%, the scheduler can ignore
//     desiredCount temporarily during a deployment. For example, if desiredCount is
//     four tasks, a minimum of 50% allows the scheduler to stop two existing tasks
//     before starting two new tasks. Tasks for services that don't use a load balancer
//     are considered healthy if they're in the RUNNING state. Tasks for services
//     that use a load balancer are considered healthy if they're in the RUNNING
//     state and are reported as healthy by the load balancer.
//   - The maximumPercent parameter represents an upper limit on the number of
//     running tasks during a deployment. You can use it to define the deployment batch
//     size. For example, if desiredCount is four tasks, a maximum of 200% starts
//     four new tasks before stopping the four older tasks (provided that the cluster
//     resources required to do this are available).
//
// When UpdateService stops a task during a deployment, the equivalent of docker
// stop is issued to the containers running in the task. This results in a SIGTERM
// and a 30-second timeout. After this, SIGKILL is sent and the containers are
// forcibly stopped. If the container handles the SIGTERM gracefully and exits
// within 30 seconds from receiving it, no SIGKILL is sent. When the service
// scheduler launches new tasks, it determines task placement in your cluster with
// the following logic.
//   - Determine which of the container instances in your cluster can support your
//     service's task definition. For example, they have the required CPU, memory,
//     ports, and container instance attributes.
//   - By default, the service scheduler attempts to balance tasks across
//     Availability Zones in this manner even though you can choose a different
//     placement strategy.
//   - Sort the valid container instances by the fewest number of running tasks
//     for this service in the same Availability Zone as the instance. For example, if
//     zone A has one running service task and zones B and C each have zero, valid
//     container instances in either zone B or C are considered optimal for placement.
//   - Place the new service task on a valid container instance in an optimal
//     Availability Zone (based on the previous steps), favoring container instances
//     with the fewest number of running tasks for this service.
//
// When the service scheduler stops running tasks, it attempts to maintain balance
// across the Availability Zones in your cluster using the following logic:
//   - Sort the container instances by the largest number of running tasks for
//     this service in the same Availability Zone as the instance. For example, if zone
//     A has one running service task and zones B and C each have two, container
//     instances in either zone B or C are considered optimal for termination.
//   - Stop the task on a container instance in an optimal Availability Zone
//     (based on the previous steps), favoring container instances with the largest
//     number of running tasks for this service.
//
// You must have a service-linked role when you update any of the following
// service properties:
//   - loadBalancers ,
//   - serviceRegistries
//
// For more information about the role see the CreateService request parameter role (https://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_CreateService.html#ECS-CreateService-request-role)
// .
func (c *Client) UpdateService(ctx context.Context, params *UpdateServiceInput, optFns ...func(*Options)) (*UpdateServiceOutput, error) {
	if params == nil {
		params = &UpdateServiceInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateService", params, optFns, c.addOperationUpdateServiceMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateServiceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateServiceInput struct {

	// The name of the service to update.
	//
	// This member is required.
	Service *string

	// The capacity provider strategy to update the service to use. if the service
	// uses the default capacity provider strategy for the cluster, the service can be
	// updated to use one or more capacity providers as opposed to the default capacity
	// provider strategy. However, when a service is using a capacity provider strategy
	// that's not the default capacity provider strategy, the service can't be updated
	// to use the cluster's default capacity provider strategy. A capacity provider
	// strategy consists of one or more capacity providers along with the base and
	// weight to assign to them. A capacity provider must be associated with the
	// cluster to be used in a capacity provider strategy. The
	// PutClusterCapacityProviders API is used to associate a capacity provider with a
	// cluster. Only capacity providers with an ACTIVE or UPDATING status can be used.
	// If specifying a capacity provider that uses an Auto Scaling group, the capacity
	// provider must already be created. New capacity providers can be created with the
	// CreateCapacityProvider API operation. To use a Fargate capacity provider,
	// specify either the FARGATE or FARGATE_SPOT capacity providers. The Fargate
	// capacity providers are available to all accounts and only need to be associated
	// with a cluster to be used. The PutClusterCapacityProviders API operation is
	// used to update the list of available capacity providers for a cluster after the
	// cluster is created.
	CapacityProviderStrategy []types.CapacityProviderStrategyItem

	// The short name or full Amazon Resource Name (ARN) of the cluster that your
	// service runs on. If you do not specify a cluster, the default cluster is
	// assumed.
	Cluster *string

	// Optional deployment parameters that control how many tasks run during the
	// deployment and the ordering of stopping and starting tasks.
	DeploymentConfiguration *types.DeploymentConfiguration

	// The number of instantiations of the task to place and keep running in your
	// service.
	DesiredCount *int32

	// Determines whether to turn on Amazon ECS managed tags for the tasks in the
	// service. For more information, see Tagging Your Amazon ECS Resources (https://docs.aws.amazon.com/AmazonECS/latest/developerguide/ecs-using-tags.html)
	// in the Amazon Elastic Container Service Developer Guide. Only tasks launched
	// after the update will reflect the update. To update the tags on all tasks, set
	// forceNewDeployment to true , so that Amazon ECS starts new tasks with the
	// updated tags.
	EnableECSManagedTags *bool

	// If true , this enables execute command functionality on all task containers. If
	// you do not want to override the value that was set when the service was created,
	// you can set this to null when performing this action.
	EnableExecuteCommand *bool

	// Determines whether to force a new deployment of the service. By default,
	// deployments aren't forced. You can use this option to start a new deployment
	// with no service definition changes. For example, you can update a service's
	// tasks to use a newer Docker image with the same image/tag combination (
	// my_image:latest ) or to roll Fargate tasks onto a newer platform version.
	ForceNewDeployment bool

	// The period of time, in seconds, that the Amazon ECS service scheduler ignores
	// unhealthy Elastic Load Balancing target health checks after a task has first
	// started. This is only valid if your service is configured to use a load
	// balancer. If your service's tasks take a while to start and respond to Elastic
	// Load Balancing health checks, you can specify a health check grace period of up
	// to 2,147,483,647 seconds. During that time, the Amazon ECS service scheduler
	// ignores the Elastic Load Balancing health check status. This grace period can
	// prevent the ECS service scheduler from marking tasks as unhealthy and stopping
	// them before they have time to come up.
	HealthCheckGracePeriodSeconds *int32

	// A list of Elastic Load Balancing load balancer objects. It contains the load
	// balancer name, the container name, and the container port to access from the
	// load balancer. The container name is as it appears in a container definition.
	// When you add, update, or remove a load balancer configuration, Amazon ECS starts
	// new tasks with the updated Elastic Load Balancing configuration, and then stops
	// the old tasks when the new tasks are running. For services that use rolling
	// updates, you can add, update, or remove Elastic Load Balancing target groups.
	// You can update from a single target group to multiple target groups and from
	// multiple target groups to a single target group. For services that use
	// blue/green deployments, you can update Elastic Load Balancing target groups by
	// using CreateDeployment (https://docs.aws.amazon.com/codedeploy/latest/APIReference/API_CreateDeployment.html)
	// through CodeDeploy. Note that multiple target groups are not supported for
	// blue/green deployments. For more information see Register multiple target
	// groups with a service (https://docs.aws.amazon.com/AmazonECS/latest/developerguide/register-multiple-targetgroups.html)
	// in the Amazon Elastic Container Service Developer Guide. For services that use
	// the external deployment controller, you can add, update, or remove load
	// balancers by using CreateTaskSet (https://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_CreateTaskSet.html)
	// . Note that multiple target groups are not supported for external deployments.
	// For more information see Register multiple target groups with a service (https://docs.aws.amazon.com/AmazonECS/latest/developerguide/register-multiple-targetgroups.html)
	// in the Amazon Elastic Container Service Developer Guide. You can remove existing
	// loadBalancers by passing an empty list.
	LoadBalancers []types.LoadBalancer

	// An object representing the network configuration for the service.
	NetworkConfiguration *types.NetworkConfiguration

	// An array of task placement constraint objects to update the service to use. If
	// no value is specified, the existing placement constraints for the service will
	// remain unchanged. If this value is specified, it will override any existing
	// placement constraints defined for the service. To remove all existing placement
	// constraints, specify an empty array. You can specify a maximum of 10 constraints
	// for each task. This limit includes constraints in the task definition and those
	// specified at runtime.
	PlacementConstraints []types.PlacementConstraint

	// The task placement strategy objects to update the service to use. If no value
	// is specified, the existing placement strategy for the service will remain
	// unchanged. If this value is specified, it will override the existing placement
	// strategy defined for the service. To remove an existing placement strategy,
	// specify an empty object. You can specify a maximum of five strategy rules for
	// each service.
	PlacementStrategy []types.PlacementStrategy

	// The platform version that your tasks in the service run on. A platform version
	// is only specified for tasks using the Fargate launch type. If a platform version
	// is not specified, the LATEST platform version is used. For more information,
	// see Fargate Platform Versions (https://docs.aws.amazon.com/AmazonECS/latest/developerguide/platform_versions.html)
	// in the Amazon Elastic Container Service Developer Guide.
	PlatformVersion *string

	// Determines whether to propagate the tags from the task definition or the
	// service to the task. If no value is specified, the tags aren't propagated. Only
	// tasks launched after the update will reflect the update. To update the tags on
	// all tasks, set forceNewDeployment to true , so that Amazon ECS starts new tasks
	// with the updated tags.
	PropagateTags types.PropagateTags

	// The configuration for this service to discover and connect to services, and be
	// discovered by, and connected from, other services within a namespace. Tasks that
	// run in a namespace can use short names to connect to services in the namespace.
	// Tasks can connect to services across all of the clusters in the namespace. Tasks
	// connect through a managed proxy container that collects logs and metrics for
	// increased visibility. Only the tasks that Amazon ECS services create are
	// supported with Service Connect. For more information, see Service Connect (https://docs.aws.amazon.com/AmazonECS/latest/developerguide/service-connect.html)
	// in the Amazon Elastic Container Service Developer Guide.
	ServiceConnectConfiguration *types.ServiceConnectConfiguration

	// The details for the service discovery registries to assign to this service. For
	// more information, see Service Discovery (https://docs.aws.amazon.com/AmazonECS/latest/developerguide/service-discovery.html)
	// . When you add, update, or remove the service registries configuration, Amazon
	// ECS starts new tasks with the updated service registries configuration, and then
	// stops the old tasks when the new tasks are running. You can remove existing
	// serviceRegistries by passing an empty list.
	ServiceRegistries []types.ServiceRegistry

	// The family and revision ( family:revision ) or full ARN of the task definition
	// to run in your service. If a revision is not specified, the latest ACTIVE
	// revision is used. If you modify the task definition with UpdateService , Amazon
	// ECS spawns a task with the new version of the task definition and then stops an
	// old task after the new version is running.
	TaskDefinition *string

	// The details of the volume that was configuredAtLaunch . You can configure the
	// size, volumeType, IOPS, throughput, snapshot and encryption in
	// ServiceManagedEBSVolumeConfiguration (https://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_ServiceManagedEBSVolumeConfiguration.html)
	// . The name of the volume must match the name from the task definition. If set
	// to null, no new deployment is triggered. Otherwise, if this configuration
	// differs from the existing one, it triggers a new deployment.
	VolumeConfigurations []types.ServiceVolumeConfiguration

	noSmithyDocumentSerde
}

type UpdateServiceOutput struct {

	// The full description of your service following the update call.
	Service *types.Service

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateServiceMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateService{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateService{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "UpdateService"); err != nil {
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
	if err = addOpUpdateServiceValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateService(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateService(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "UpdateService",
	}
}
