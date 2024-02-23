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

// This API modifies the attributes of a serverless cache.
func (c *Client) ModifyServerlessCache(ctx context.Context, params *ModifyServerlessCacheInput, optFns ...func(*Options)) (*ModifyServerlessCacheOutput, error) {
	if params == nil {
		params = &ModifyServerlessCacheInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ModifyServerlessCache", params, optFns, c.addOperationModifyServerlessCacheMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ModifyServerlessCacheOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ModifyServerlessCacheInput struct {

	// User-provided identifier for the serverless cache to be modified.
	//
	// This member is required.
	ServerlessCacheName *string

	// Modify the cache usage limit for the serverless cache.
	CacheUsageLimits *types.CacheUsageLimits

	// The daily time during which Elasticache begins taking a daily snapshot of the
	// serverless cache. Available for Redis only. The default is NULL, i.e. the
	// existing snapshot time configured for the cluster is not removed.
	DailySnapshotTime *string

	// User provided description for the serverless cache. Default = NULL, i.e. the
	// existing description is not removed/modified. The description has a maximum
	// length of 255 characters.
	Description *string

	// The identifier of the UserGroup to be removed from association with the Redis
	// serverless cache. Available for Redis only. Default is NULL.
	RemoveUserGroup *bool

	// The new list of VPC security groups to be associated with the serverless cache.
	// Populating this list means the current VPC security groups will be removed. This
	// security group is used to authorize traffic access for the VPC end-point
	// (private-link). Default = NULL - the existing list of VPC security groups is not
	// removed.
	SecurityGroupIds []string

	// The number of days for which Elasticache retains automatic snapshots before
	// deleting them. Available for Redis only. Default = NULL, i.e. the existing
	// snapshot-retention-limit will not be removed or modified. The maximum value
	// allowed is 35 days.
	SnapshotRetentionLimit *int32

	// The identifier of the UserGroup to be associated with the serverless cache.
	// Available for Redis only. Default is NULL - the existing UserGroup is not
	// removed.
	UserGroupId *string

	noSmithyDocumentSerde
}

type ModifyServerlessCacheOutput struct {

	// The response for the attempt to modify the serverless cache.
	ServerlessCache *types.ServerlessCache

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationModifyServerlessCacheMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsquery_serializeOpModifyServerlessCache{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpModifyServerlessCache{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ModifyServerlessCache"); err != nil {
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
	if err = addOpModifyServerlessCacheValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opModifyServerlessCache(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opModifyServerlessCache(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ModifyServerlessCache",
	}
}
