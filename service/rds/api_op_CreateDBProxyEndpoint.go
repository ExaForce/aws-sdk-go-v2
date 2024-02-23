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

// Creates a DBProxyEndpoint . Only applies to proxies that are associated with
// Aurora DB clusters. You can use DB proxy endpoints to specify read/write or
// read-only access to the DB cluster. You can also use DB proxy endpoints to
// access a DB proxy through a different VPC than the proxy's default VPC.
func (c *Client) CreateDBProxyEndpoint(ctx context.Context, params *CreateDBProxyEndpointInput, optFns ...func(*Options)) (*CreateDBProxyEndpointOutput, error) {
	if params == nil {
		params = &CreateDBProxyEndpointInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateDBProxyEndpoint", params, optFns, c.addOperationCreateDBProxyEndpointMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateDBProxyEndpointOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateDBProxyEndpointInput struct {

	// The name of the DB proxy endpoint to create.
	//
	// This member is required.
	DBProxyEndpointName *string

	// The name of the DB proxy associated with the DB proxy endpoint that you create.
	//
	// This member is required.
	DBProxyName *string

	// The VPC subnet IDs for the DB proxy endpoint that you create. You can specify a
	// different set of subnet IDs than for the original DB proxy.
	//
	// This member is required.
	VpcSubnetIds []string

	// A list of tags. For more information, see Tagging Amazon RDS Resources (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_Tagging.html)
	// in the Amazon RDS User Guide.
	Tags []types.Tag

	// The role of the DB proxy endpoint. The role determines whether the endpoint can
	// be used for read/write or only read operations. The default is READ_WRITE . The
	// only role that proxies for RDS for Microsoft SQL Server support is READ_WRITE .
	TargetRole types.DBProxyEndpointTargetRole

	// The VPC security group IDs for the DB proxy endpoint that you create. You can
	// specify a different set of security group IDs than for the original DB proxy.
	// The default is the default security group for the VPC.
	VpcSecurityGroupIds []string

	noSmithyDocumentSerde
}

type CreateDBProxyEndpointOutput struct {

	// The DBProxyEndpoint object that is created by the API operation. The DB proxy
	// endpoint that you create might provide capabilities such as read/write or
	// read-only operations, or using a different VPC than the proxy's default VPC.
	DBProxyEndpoint *types.DBProxyEndpoint

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateDBProxyEndpointMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsquery_serializeOpCreateDBProxyEndpoint{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpCreateDBProxyEndpoint{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateDBProxyEndpoint"); err != nil {
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
	if err = addOpCreateDBProxyEndpointValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateDBProxyEndpoint(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateDBProxyEndpoint(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateDBProxyEndpoint",
	}
}
