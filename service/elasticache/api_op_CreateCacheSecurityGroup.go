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

// Creates a new cache security group. Use a cache security group to control
// access to one or more clusters. Cache security groups are only used when you are
// creating a cluster outside of an Amazon Virtual Private Cloud (Amazon VPC). If
// you are creating a cluster inside of a VPC, use a cache subnet group instead.
// For more information, see CreateCacheSubnetGroup (https://docs.aws.amazon.com/AmazonElastiCache/latest/APIReference/API_CreateCacheSubnetGroup.html)
// .
func (c *Client) CreateCacheSecurityGroup(ctx context.Context, params *CreateCacheSecurityGroupInput, optFns ...func(*Options)) (*CreateCacheSecurityGroupOutput, error) {
	if params == nil {
		params = &CreateCacheSecurityGroupInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateCacheSecurityGroup", params, optFns, c.addOperationCreateCacheSecurityGroupMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateCacheSecurityGroupOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input of a CreateCacheSecurityGroup operation.
type CreateCacheSecurityGroupInput struct {

	// A name for the cache security group. This value is stored as a lowercase
	// string. Constraints: Must contain no more than 255 alphanumeric characters.
	// Cannot be the word "Default". Example: mysecuritygroup
	//
	// This member is required.
	CacheSecurityGroupName *string

	// A description for the cache security group.
	//
	// This member is required.
	Description *string

	// A list of tags to be added to this resource. A tag is a key-value pair. A tag
	// key must be accompanied by a tag value, although null is accepted.
	Tags []types.Tag

	noSmithyDocumentSerde
}

type CreateCacheSecurityGroupOutput struct {

	// Represents the output of one of the following operations:
	//   - AuthorizeCacheSecurityGroupIngress
	//   - CreateCacheSecurityGroup
	//   - RevokeCacheSecurityGroupIngress
	CacheSecurityGroup *types.CacheSecurityGroup

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateCacheSecurityGroupMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsquery_serializeOpCreateCacheSecurityGroup{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpCreateCacheSecurityGroup{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateCacheSecurityGroup"); err != nil {
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
	if err = addOpCreateCacheSecurityGroupValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateCacheSecurityGroup(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateCacheSecurityGroup(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateCacheSecurityGroup",
	}
}
