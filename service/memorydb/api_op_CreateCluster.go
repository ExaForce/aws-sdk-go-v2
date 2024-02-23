// Code generated by smithy-go-codegen DO NOT EDIT.

package memorydb

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/memorydb/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a cluster. All nodes in the cluster run the same protocol-compliant
// engine software.
func (c *Client) CreateCluster(ctx context.Context, params *CreateClusterInput, optFns ...func(*Options)) (*CreateClusterOutput, error) {
	if params == nil {
		params = &CreateClusterInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateCluster", params, optFns, c.addOperationCreateClusterMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateClusterOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateClusterInput struct {

	// The name of the Access Control List to associate with the cluster.
	//
	// This member is required.
	ACLName *string

	// The name of the cluster. This value must be unique as it also serves as the
	// cluster identifier.
	//
	// This member is required.
	ClusterName *string

	// The compute and memory capacity of the nodes in the cluster.
	//
	// This member is required.
	NodeType *string

	// When set to true, the cluster will automatically receive minor engine version
	// upgrades after launch.
	AutoMinorVersionUpgrade *bool

	// Enables data tiering. Data tiering is only supported for clusters using the
	// r6gd node type. This parameter must be set when using r6gd nodes. For more
	// information, see Data tiering (https://docs.aws.amazon.com/memorydb/latest/devguide/data-tiering.html)
	// .
	DataTiering *bool

	// An optional description of the cluster.
	Description *string

	// The version number of the Redis engine to be used for the cluster.
	EngineVersion *string

	// The ID of the KMS key used to encrypt the cluster.
	KmsKeyId *string

	// Specifies the weekly time range during which maintenance on the cluster is
	// performed. It is specified as a range in the format ddd:hh24:mi-ddd:hh24:mi (24H
	// Clock UTC). The minimum maintenance window is a 60 minute period. Valid values
	// for ddd are:
	//   - sun
	//   - mon
	//   - tue
	//   - wed
	//   - thu
	//   - fri
	//   - sat
	// Example: sun:23:00-mon:01:30
	MaintenanceWindow *string

	// The number of replicas to apply to each shard. The default value is 1. The
	// maximum is 5.
	NumReplicasPerShard *int32

	// The number of shards the cluster will contain. The default value is 1.
	NumShards *int32

	// The name of the parameter group associated with the cluster.
	ParameterGroupName *string

	// The port number on which each of the nodes accepts connections.
	Port *int32

	// A list of security group names to associate with this cluster.
	SecurityGroupIds []string

	// A list of Amazon Resource Names (ARN) that uniquely identify the RDB snapshot
	// files stored in Amazon S3. The snapshot files are used to populate the new
	// cluster. The Amazon S3 object name in the ARN cannot contain any commas.
	SnapshotArns []string

	// The name of a snapshot from which to restore data into the new cluster. The
	// snapshot status changes to restoring while the new cluster is being created.
	SnapshotName *string

	// The number of days for which MemoryDB retains automatic snapshots before
	// deleting them. For example, if you set SnapshotRetentionLimit to 5, a snapshot
	// that was taken today is retained for 5 days before being deleted.
	SnapshotRetentionLimit *int32

	// The daily time range (in UTC) during which MemoryDB begins taking a daily
	// snapshot of your shard. Example: 05:00-09:00 If you do not specify this
	// parameter, MemoryDB automatically chooses an appropriate time range.
	SnapshotWindow *string

	// The Amazon Resource Name (ARN) of the Amazon Simple Notification Service (SNS)
	// topic to which notifications are sent.
	SnsTopicArn *string

	// The name of the subnet group to be used for the cluster.
	SubnetGroupName *string

	// A flag to enable in-transit encryption on the cluster.
	TLSEnabled *bool

	// A list of tags to be added to this resource. Tags are comma-separated key,value
	// pairs (e.g. Key=myKey, Value=myKeyValue. You can include multiple tags as shown
	// following: Key=myKey, Value=myKeyValue Key=mySecondKey, Value=mySecondKeyValue.
	Tags []types.Tag

	noSmithyDocumentSerde
}

type CreateClusterOutput struct {

	// The newly-created cluster.
	Cluster *types.Cluster

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateClusterMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateCluster{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateCluster{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateCluster"); err != nil {
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
	if err = addOpCreateClusterValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateCluster(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateCluster(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateCluster",
	}
}
