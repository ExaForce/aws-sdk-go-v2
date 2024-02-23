// Code generated by smithy-go-codegen DO NOT EDIT.

package rds

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/rds/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Promotes the specified secondary DB cluster to be the primary DB cluster in the
// global database cluster to fail over or switch over a global database.
// Switchover operations were previously called "managed planned failovers."
// Although this operation can be used either to fail over or to switch over a
// global database cluster, its intended use is for global database failover. To
// switch over a global database cluster, we recommend that you use the
// SwitchoverGlobalCluster operation instead. How you use this operation depends on
// whether you are failing over or switching over your global database cluster:
//   - Failing over - Specify the AllowDataLoss parameter and don't specify the
//     Switchover parameter.
//   - Switching over - Specify the Switchover parameter or omit it, but don't
//     specify the AllowDataLoss parameter.
//
// About failing over and switching over While failing over and switching over a
// global database cluster both change the primary DB cluster, you use these
// operations for different reasons:
//   - Failing over - Use this operation to respond to an unplanned event, such as
//     a Regional disaster in the primary Region. Failing over can result in a loss of
//     write transaction data that wasn't replicated to the chosen secondary before the
//     failover event occurred. However, the recovery process that promotes a DB
//     instance on the chosen seconday DB cluster to be the primary writer DB instance
//     guarantees that the data is in a transactionally consistent state. For more
//     information about failing over an Amazon Aurora global database, see
//     Performing managed failovers for Aurora global databases (https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/aurora-global-database-disaster-recovery.html#aurora-global-database-failover.managed-unplanned)
//     in the Amazon Aurora User Guide.
//   - Switching over - Use this operation on a healthy global database cluster
//     for planned events, such as Regional rotation or to fail back to the original
//     primary DB cluster after a failover operation. With this operation, there is no
//     data loss. For more information about switching over an Amazon Aurora global
//     database, see Performing switchovers for Aurora global databases (https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/aurora-global-database-disaster-recovery.html#aurora-global-database-disaster-recovery.managed-failover)
//     in the Amazon Aurora User Guide.
func (c *Client) FailoverGlobalCluster(ctx context.Context, params *FailoverGlobalClusterInput, optFns ...func(*Options)) (*FailoverGlobalClusterOutput, error) {
	if params == nil {
		params = &FailoverGlobalClusterInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "FailoverGlobalCluster", params, optFns, c.addOperationFailoverGlobalClusterMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*FailoverGlobalClusterOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type FailoverGlobalClusterInput struct {

	// The identifier of the global database cluster (Aurora global database) this
	// operation should apply to. The identifier is the unique key assigned by the user
	// when the Aurora global database is created. In other words, it's the name of the
	// Aurora global database. Constraints:
	//   - Must match the identifier of an existing global database cluster.
	//
	// This member is required.
	GlobalClusterIdentifier *string

	// The identifier of the secondary Aurora DB cluster that you want to promote to
	// the primary for the global database cluster. Use the Amazon Resource Name (ARN)
	// for the identifier so that Aurora can locate the cluster in its Amazon Web
	// Services Region.
	//
	// This member is required.
	TargetDbClusterIdentifier *string

	// Specifies whether to allow data loss for this global database cluster
	// operation. Allowing data loss triggers a global failover operation. If you don't
	// specify AllowDataLoss , the global database cluster operation defaults to a
	// switchover. Constraints:
	//   - Can't be specified together with the Switchover parameter.
	AllowDataLoss *bool

	// Specifies whether to switch over this global database cluster. Constraints:
	//   - Can't be specified together with the AllowDataLoss parameter.
	Switchover *bool

	noSmithyDocumentSerde
}

type FailoverGlobalClusterOutput struct {

	// A data type representing an Aurora global database.
	GlobalCluster *types.GlobalCluster

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationFailoverGlobalClusterMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsquery_serializeOpFailoverGlobalCluster{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpFailoverGlobalCluster{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "FailoverGlobalCluster"); err != nil {
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
	if err = addOpFailoverGlobalClusterValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opFailoverGlobalCluster(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opFailoverGlobalCluster(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "FailoverGlobalCluster",
	}
}
