// Code generated by smithy-go-codegen DO NOT EDIT.

package eventbridge

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/eventbridge/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a global endpoint. Global endpoints improve your application's
// availability by making it regional-fault tolerant. To do this, you define a
// primary and secondary Region with event buses in each Region. You also create a
// Amazon Route 53 health check that will tell EventBridge to route events to the
// secondary Region when an "unhealthy" state is encountered and events will be
// routed back to the primary Region when the health check reports a "healthy"
// state.
func (c *Client) CreateEndpoint(ctx context.Context, params *CreateEndpointInput, optFns ...func(*Options)) (*CreateEndpointOutput, error) {
	if params == nil {
		params = &CreateEndpointInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateEndpoint", params, optFns, c.addOperationCreateEndpointMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateEndpointOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateEndpointInput struct {

	// Define the event buses used. The names of the event buses must be identical in
	// each Region.
	//
	// This member is required.
	EventBuses []types.EndpointEventBus

	// The name of the global endpoint. For example,
	// "Name":"us-east-2-custom_bus_A-endpoint" .
	//
	// This member is required.
	Name *string

	// Configure the routing policy, including the health check and secondary Region..
	//
	// This member is required.
	RoutingConfig *types.RoutingConfig

	// A description of the global endpoint.
	Description *string

	// Enable or disable event replication. The default state is ENABLED which means
	// you must supply a RoleArn . If you don't have a RoleArn or you don't want event
	// replication enabled, set the state to DISABLED .
	ReplicationConfig *types.ReplicationConfig

	// The ARN of the role used for replication.
	RoleArn *string

	noSmithyDocumentSerde
}

type CreateEndpointOutput struct {

	// The ARN of the endpoint that was created by this request.
	Arn *string

	// The event buses used by this request.
	EventBuses []types.EndpointEventBus

	// The name of the endpoint that was created by this request.
	Name *string

	// Whether event replication was enabled or disabled by this request.
	ReplicationConfig *types.ReplicationConfig

	// The ARN of the role used by event replication for this request.
	RoleArn *string

	// The routing configuration defined by this request.
	RoutingConfig *types.RoutingConfig

	// The state of the endpoint that was created by this request.
	State types.EndpointState

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateEndpointMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateEndpoint{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateEndpoint{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateEndpoint"); err != nil {
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
	if err = addOpCreateEndpointValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateEndpoint(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateEndpoint(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateEndpoint",
	}
}
