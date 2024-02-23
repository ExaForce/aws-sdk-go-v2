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

// Finds new players to fill open slots in currently running game sessions. The
// backfill match process is essentially identical to the process of forming new
// matches. Backfill requests use the same matchmaker that was used to make the
// original match, and they provide matchmaking data for all players currently in
// the game session. FlexMatch uses this information to select new players so that
// backfilled match continues to meet the original match requirements. When using
// FlexMatch with Amazon GameLift managed hosting, you can request a backfill match
// from a client service by calling this operation with a GameSessions ID. You
// also have the option of making backfill requests directly from your game server.
// In response to a request, FlexMatch creates player sessions for the new players,
// updates the GameSession resource, and sends updated matchmaking data to the
// game server. You can request a backfill match at any point after a game session
// is started. Each game session can have only one active backfill request at a
// time; a subsequent request automatically replaces the earlier request. When
// using FlexMatch as a standalone component, request a backfill match by calling
// this operation without a game session identifier. As with newly formed matches,
// matchmaking results are returned in a matchmaking event so that your game can
// update the game session that is being backfilled. To request a backfill match,
// specify a unique ticket ID, the original matchmaking configuration, and
// matchmaking data for all current players in the game session being backfilled.
// Optionally, specify the GameSession ARN. If successful, a match backfill ticket
// is created and returned with status set to QUEUED. Track the status of backfill
// tickets using the same method for tracking tickets for new matches. Only game
// sessions created by FlexMatch are supported for match backfill. Learn more
// Backfill existing games with FlexMatch (https://docs.aws.amazon.com/gamelift/latest/flexmatchguide/match-backfill.html)
// Matchmaking events (https://docs.aws.amazon.com/gamelift/latest/flexmatchguide/match-events.html)
// (reference) How Amazon GameLift FlexMatch works (https://docs.aws.amazon.com/gamelift/latest/flexmatchguide/gamelift-match.html)
func (c *Client) StartMatchBackfill(ctx context.Context, params *StartMatchBackfillInput, optFns ...func(*Options)) (*StartMatchBackfillOutput, error) {
	if params == nil {
		params = &StartMatchBackfillInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "StartMatchBackfill", params, optFns, c.addOperationStartMatchBackfillMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*StartMatchBackfillOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StartMatchBackfillInput struct {

	// Name of the matchmaker to use for this request. You can use either the
	// configuration name or ARN value. The ARN of the matchmaker that was used with
	// the original game session is listed in the GameSession object, MatchmakerData
	// property.
	//
	// This member is required.
	ConfigurationName *string

	// Match information on all players that are currently assigned to the game
	// session. This information is used by the matchmaker to find new players and add
	// them to the existing game. You can include up to 199 Players in a
	// StartMatchBackfill request.
	//   - PlayerID, PlayerAttributes, Team -- This information is maintained in the
	//   GameSession object, MatchmakerData property, for all players who are currently
	//   assigned to the game session. The matchmaker data is in JSON syntax, formatted
	//   as a string. For more details, see Match Data (https://docs.aws.amazon.com/gamelift/latest/flexmatchguide/match-server.html#match-server-data)
	//   . The backfill request must specify the team membership for every player. Do not
	//   specify team if you are not using backfill.
	//   - LatencyInMs -- If the matchmaker uses player latency, include a latency
	//   value, in milliseconds, for the Region that the game session is currently in. Do
	//   not include latency values for any other Region.
	//
	// This member is required.
	Players []types.Player

	// A unique identifier for the game session. Use the game session ID. When using
	// FlexMatch as a standalone matchmaking solution, this parameter is not needed.
	GameSessionArn *string

	// A unique identifier for a matchmaking ticket. If no ticket ID is specified
	// here, Amazon GameLift will generate one in the form of a UUID. Use this
	// identifier to track the match backfill ticket status and retrieve match results.
	TicketId *string

	noSmithyDocumentSerde
}

type StartMatchBackfillOutput struct {

	// Ticket representing the backfill matchmaking request. This object includes the
	// information in the request, ticket status, and match results as generated during
	// the matchmaking process.
	MatchmakingTicket *types.MatchmakingTicket

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationStartMatchBackfillMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpStartMatchBackfill{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpStartMatchBackfill{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "StartMatchBackfill"); err != nil {
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
	if err = addOpStartMatchBackfillValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opStartMatchBackfill(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opStartMatchBackfill(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "StartMatchBackfill",
	}
}
