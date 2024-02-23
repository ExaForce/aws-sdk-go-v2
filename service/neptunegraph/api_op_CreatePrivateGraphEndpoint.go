// Code generated by smithy-go-codegen DO NOT EDIT.

package neptunegraph

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/neptunegraph/types"
	"github.com/aws/smithy-go/middleware"
	"github.com/aws/smithy-go/ptr"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Create a private graph endpoint to allow private access from to the graph from
// within a VPC. You can attach security groups to the private graph endpoint. VPC
// endpoint charges apply.
func (c *Client) CreatePrivateGraphEndpoint(ctx context.Context, params *CreatePrivateGraphEndpointInput, optFns ...func(*Options)) (*CreatePrivateGraphEndpointOutput, error) {
	if params == nil {
		params = &CreatePrivateGraphEndpointInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreatePrivateGraphEndpoint", params, optFns, c.addOperationCreatePrivateGraphEndpointMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreatePrivateGraphEndpointOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreatePrivateGraphEndpointInput struct {

	// The unique identifier of the Neptune Analytics graph.
	//
	// This member is required.
	GraphIdentifier *string

	// Subnets in which private graph endpoint ENIs are created.
	SubnetIds []string

	// The VPC in which the private graph endpoint needs to be created.
	VpcId *string

	// Security groups to be attached to the private graph endpoint..
	VpcSecurityGroupIds []string

	noSmithyDocumentSerde
}

func (in *CreatePrivateGraphEndpointInput) bindEndpointParams(p *EndpointParameters) {

	p.ApiType = ptr.String("ControlPlane")
}

type CreatePrivateGraphEndpointOutput struct {

	// Status of the private graph endpoint.
	//
	// This member is required.
	Status types.PrivateGraphEndpointStatus

	// Subnets in which the private graph endpoint ENIs are created.
	//
	// This member is required.
	SubnetIds []string

	// VPC in which the private graph endpoint is created.
	//
	// This member is required.
	VpcId *string

	// Endpoint ID of the prviate grpah endpoint.
	VpcEndpointId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreatePrivateGraphEndpointMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreatePrivateGraphEndpoint{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreatePrivateGraphEndpoint{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreatePrivateGraphEndpoint"); err != nil {
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
	if err = addOpCreatePrivateGraphEndpointValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreatePrivateGraphEndpoint(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreatePrivateGraphEndpoint(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreatePrivateGraphEndpoint",
	}
}
