// Code generated by smithy-go-codegen DO NOT EDIT.

package elasticache

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/elasticache/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a serverless cache.
func (c *Client) CreateServerlessCache(ctx context.Context, params *CreateServerlessCacheInput, optFns ...func(*Options)) (*CreateServerlessCacheOutput, error) {
	if params == nil {
		params = &CreateServerlessCacheInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateServerlessCache", params, optFns, c.addOperationCreateServerlessCacheMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateServerlessCacheOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateServerlessCacheInput struct {

	// The name of the cache engine to be used for creating the serverless cache.
	//
	// This member is required.
	Engine *string

	// User-provided identifier for the serverless cache. This parameter is stored as
	// a lowercase string.
	//
	// This member is required.
	ServerlessCacheName *string

	// Sets the cache usage limits for storage and ElastiCache Processing Units for
	// the cache.
	CacheUsageLimits *types.CacheUsageLimits

	// The daily time that snapshots will be created from the new serverless cache. By
	// default this number is populated with 0, i.e. no snapshots will be created on an
	// automatic daily basis. Available for Redis only.
	DailySnapshotTime *string

	// User-provided description for the serverless cache. The default is NULL, i.e.
	// if no description is provided then an empty string will be returned. The maximum
	// length is 255 characters.
	Description *string

	// ARN of the customer managed key for encrypting the data at rest. If no KMS key
	// is provided, a default service key is used.
	KmsKeyId *string

	// The version of the cache engine that will be used to create the serverless
	// cache.
	MajorEngineVersion *string

	// A list of the one or more VPC security groups to be associated with the
	// serverless cache. The security group will authorize traffic access for the VPC
	// end-point (private-link). If no other information is given this will be the
	// VPC’s Default Security Group that is associated with the cluster VPC end-point.
	SecurityGroupIds []string

	// The ARN(s) of the snapshot that the new serverless cache will be created from.
	// Available for Redis only.
	SnapshotArnsToRestore []string

	// The number of snapshots that will be retained for the serverless cache that is
	// being created. As new snapshots beyond this limit are added, the oldest
	// snapshots will be deleted on a rolling basis. Available for Redis only.
	SnapshotRetentionLimit *int32

	// A list of the identifiers of the subnets where the VPC endpoint for the
	// serverless cache will be deployed. All the subnetIds must belong to the same
	// VPC.
	SubnetIds []string

	// The list of tags (key, value) pairs to be added to the serverless cache
	// resource. Default is NULL.
	Tags []types.Tag

	// The identifier of the UserGroup to be associated with the serverless cache.
	// Available for Redis only. Default is NULL.
	UserGroupId *string

	noSmithyDocumentSerde
}

type CreateServerlessCacheOutput struct {

	// The response for the attempt to create the serverless cache.
	ServerlessCache *types.ServerlessCache

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateServerlessCacheMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsquery_serializeOpCreateServerlessCache{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpCreateServerlessCache{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateServerlessCache"); err != nil {
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
	if err = addOpCreateServerlessCacheValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateServerlessCache(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateServerlessCache(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateServerlessCache",
	}
}
