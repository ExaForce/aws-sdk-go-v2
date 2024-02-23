// Code generated by smithy-go-codegen DO NOT EDIT.

package gamelift

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/gamelift/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// This operation is used with the Amazon GameLift FleetIQ solution and game
// server groups. Creates a new game server resource and notifies Amazon GameLift
// FleetIQ that the game server is ready to host gameplay and players. This
// operation is called by a game server process that is running on an instance in a
// game server group. Registering game servers enables Amazon GameLift FleetIQ to
// track available game servers and enables game clients and services to claim a
// game server for a new game session. To register a game server, identify the game
// server group and instance where the game server is running, and provide a unique
// identifier for the game server. You can also include connection and game server
// data. Once a game server is successfully registered, it is put in status
// AVAILABLE . A request to register a game server may fail if the instance it is
// running on is in the process of shutting down as part of instance balancing or
// scale-down activity. Learn more Amazon GameLift FleetIQ Guide (https://docs.aws.amazon.com/gamelift/latest/fleetiqguide/gsg-intro.html)
func (c *Client) RegisterGameServer(ctx context.Context, params *RegisterGameServerInput, optFns ...func(*Options)) (*RegisterGameServerOutput, error) {
	if params == nil {
		params = &RegisterGameServerInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "RegisterGameServer", params, optFns, c.addOperationRegisterGameServerMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*RegisterGameServerOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type RegisterGameServerInput struct {

	// A unique identifier for the game server group where the game server is running.
	//
	// This member is required.
	GameServerGroupName *string

	// A custom string that uniquely identifies the game server to register. Game
	// server IDs are developer-defined and must be unique across all game server
	// groups in your Amazon Web Services account.
	//
	// This member is required.
	GameServerId *string

	// The unique identifier for the instance where the game server is running. This
	// ID is available in the instance metadata. EC2 instance IDs use a 17-character
	// format, for example: i-1234567890abcdef0 .
	//
	// This member is required.
	InstanceId *string

	// Information that is needed to make inbound client connections to the game
	// server. This might include the IP address and port, DNS name, and other
	// information.
	ConnectionInfo *string

	// A set of custom game server properties, formatted as a single string value.
	// This data is passed to a game client or service when it requests information on
	// game servers.
	GameServerData *string

	noSmithyDocumentSerde
}

type RegisterGameServerOutput struct {

	// Object that describes the newly registered game server.
	GameServer *types.GameServer

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationRegisterGameServerMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpRegisterGameServer{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpRegisterGameServer{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "RegisterGameServer"); err != nil {
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
	if err = addOpRegisterGameServerValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opRegisterGameServer(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opRegisterGameServer(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "RegisterGameServer",
	}
}
