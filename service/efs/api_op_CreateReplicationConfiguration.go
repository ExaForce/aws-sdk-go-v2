// Code generated by smithy-go-codegen DO NOT EDIT.

package efs

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/efs/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Creates a replication configuration that replicates an existing EFS file system
// to a new, read-only file system. For more information, see Amazon EFS
// replication (https://docs.aws.amazon.com/efs/latest/ug/efs-replication.html) in
// the Amazon EFS User Guide. The replication configuration specifies the
// following:
//   - Source file system – The EFS file system that you want replicated. The
//     source file system cannot be a destination file system in an existing
//     replication configuration.
//   - Amazon Web Services Region – The Amazon Web Services Region in which the
//     destination file system is created. Amazon EFS replication is available in all
//     Amazon Web Services Regions in which EFS is available. The Region must be
//     enabled. For more information, see Managing Amazon Web Services Regions (https://docs.aws.amazon.com/general/latest/gr/rande-manage.html#rande-manage-enable)
//     in the Amazon Web Services General Reference Reference Guide.
//   - Destination file system configuration – The configuration of the
//     destination file system to which the source file system will be replicated.
//     There can only be one destination file system in a replication configuration.
//     Parameters for the replication configuration include:
//   - File system ID – The ID of the destination file system for the replication.
//     If no ID is provided, then EFS creates a new file system with the default
//     settings. For existing file systems, the file system's replication overwrite
//     protection must be disabled. For more information, see Replicating to an
//     existing file system (https://docs.aws.amazon.com/efs/latest/ug/efs-replication#replicate-existing-destination)
//     .
//   - Availability Zone – If you want the destination file system to use One Zone
//     storage, you must specify the Availability Zone to create the file system in.
//     For more information, see EFS file system types (https://docs.aws.amazon.com/efs/latest/ug/storage-classes.html)
//     in the Amazon EFS User Guide.
//   - Encryption – All destination file systems are created with encryption at
//     rest enabled. You can specify the Key Management Service (KMS) key that is used
//     to encrypt the destination file system. If you don't specify a KMS key, your
//     service-managed KMS key for Amazon EFS is used. After the file system is
//     created, you cannot change the KMS key.
//
// After the file system is created, you cannot change the KMS key. For new
// destination file systems, the following properties are set by default:
//
//   - Performance mode - The destination file system's performance mode matches
//     that of the source file system, unless the destination file system uses EFS One
//     Zone storage. In that case, the General Purpose performance mode is used. The
//     performance mode cannot be changed.
//
//   - Throughput mode - The destination file system's throughput mode matches
//     that of the source file system. After the file system is created, you can modify
//     the throughput mode.
//
//   - Lifecycle management – Lifecycle management is not enabled on the
//     destination file system. After the destination file system is created, you can
//     enable lifecycle management.
//
//   - Automatic backups – Automatic daily backups are enabled on the destination
//     file system. After the file system is created, you can change this setting.
//
// For more information, see Amazon EFS replication (https://docs.aws.amazon.com/efs/latest/ug/efs-replication.html)
// in the Amazon EFS User Guide.
func (c *Client) CreateReplicationConfiguration(ctx context.Context, params *CreateReplicationConfigurationInput, optFns ...func(*Options)) (*CreateReplicationConfigurationOutput, error) {
	if params == nil {
		params = &CreateReplicationConfigurationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateReplicationConfiguration", params, optFns, c.addOperationCreateReplicationConfigurationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateReplicationConfigurationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateReplicationConfigurationInput struct {

	// An array of destination configuration objects. Only one destination
	// configuration object is supported.
	//
	// This member is required.
	Destinations []types.DestinationToCreate

	// Specifies the Amazon EFS file system that you want to replicate. This file
	// system cannot already be a source or destination file system in another
	// replication configuration.
	//
	// This member is required.
	SourceFileSystemId *string

	noSmithyDocumentSerde
}

// Describes the replication configuration for a specific file system.
type CreateReplicationConfigurationOutput struct {

	// Describes when the replication configuration was created.
	//
	// This member is required.
	CreationTime *time.Time

	// An array of destination objects. Only one destination object is supported.
	//
	// This member is required.
	Destinations []types.Destination

	// The Amazon Resource Name (ARN) of the original source EFS file system in the
	// replication configuration.
	//
	// This member is required.
	OriginalSourceFileSystemArn *string

	// The Amazon Resource Name (ARN) of the current source file system in the
	// replication configuration.
	//
	// This member is required.
	SourceFileSystemArn *string

	// The ID of the source Amazon EFS file system that is being replicated.
	//
	// This member is required.
	SourceFileSystemId *string

	// The Amazon Web Services Region in which the source EFS file system is located.
	//
	// This member is required.
	SourceFileSystemRegion *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateReplicationConfigurationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateReplicationConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateReplicationConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateReplicationConfiguration"); err != nil {
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
	if err = addOpCreateReplicationConfigurationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateReplicationConfiguration(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateReplicationConfiguration(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateReplicationConfiguration",
	}
}
